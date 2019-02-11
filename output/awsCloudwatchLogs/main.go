package awsCloudwatchLogs

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/gabrielperezs/sytoco/lib"
)

const (
	// All details about the API:
	// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#CloudWatchLogs.PutLogEvents

	// * The maximum number of log events in a batch is 10,000.
	awsLogsBatchRecords = 10000
	// The maximum batch size is 1,048,576 bytes, and this size is calculated
	// as the sum of all event messages in UTF-8, plus 26 bytes for each log
	// event.
	awsLogsBatchSize       = 1048576 // bytes
	awsLogsRecordExtraSize = 26      // bytes

	inputTimeout = 15 * time.Second
)

type Config struct {
	Profile       string
	Region        string
	LogGroupName  string
	LogStreamName string
}

type LogClient struct {
	sess              *session.Session
	svc               *cloudwatchlogs.CloudWatchLogs
	buffer            []*cloudwatchlogs.InputLogEvent
	bufferSize        int
	chAppend          chan lib.Record
	chFlush           chan bool
	chTimeout         *time.Timer
	nextSequenceToken *string
	logGroupName      string
	logStreamName     string
}

func New(cfg *Config) (*LogClient, error) {
	l := &LogClient{
		chAppend:      make(chan lib.Record, awsLogsBatchRecords),
		chFlush:       make(chan bool, 1),
		chTimeout:     time.NewTimer(inputTimeout),
		logGroupName:  cfg.LogGroupName,
		logStreamName: cfg.LogStreamName,
	}
	var err error
	l.sess, err = session.NewSessionWithOptions(session.Options{
		Profile:           cfg.Profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	l.svc = cloudwatchlogs.New(l.sess, &aws.Config{Region: aws.String(cfg.Region)})

	_, err = l.svc.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(l.logGroupName),
		LogStreamName: aws.String(l.logStreamName),
	})
	if err != nil {
		if !strings.Contains(err.Error(), "ResourceAlreadyExistsException") {
			log.Printf("sytoco/awsCloudwatchLogs ERROR CreateLogStream: %s", err)
		}
	}

	resp, err := l.svc.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String(l.logGroupName), // Required
		LogStreamNamePrefix: aws.String(l.logStreamName),
		Descending:          aws.Bool(true),
		Limit:               aws.Int64(1),
	})
	if err != nil {
		log.Printf("sytoco/awsCloudwatchLogs ERROR DescribeLogStreams: %s", err.Error())
		return nil, err
	}

	if len(resp.LogStreams) == 0 {
		log.Print("sytoco/awsCloudwatchLogs ERROR: empty results requesting for logStreams information")
		return nil, err
	}

	if resp.LogStreams[0].UploadSequenceToken != nil {
		l.nextSequenceToken = aws.String("")
		*l.nextSequenceToken = *resp.LogStreams[0].UploadSequenceToken
	}

	go l.listener()
	return l, nil
}

func (l *LogClient) listener() {
	for {
		select {
		case r := <-l.chAppend:
			logAwsFormat, size, err := l.formatRecord(r)
			if err != nil || l.bufferSize+size >= awsLogsBatchSize {
				l.Flush()
				break
			}
			l.buffer = append(l.buffer, logAwsFormat)
			l.bufferSize += size
		case <-l.chFlush:
			l.Flush()
		case <-l.chTimeout.C:
			l.Flush()
		}
	}
}

func (l *LogClient) formatRecord(r lib.Record) (*cloudwatchlogs.InputLogEvent, int, error) {
	epochNano, err := strconv.ParseInt(r.REALTIME_TIMESTAMP, 10, 64)
	if err != nil {
		return nil, 0, err
	}

	msgSize := len([]byte(r.MESSAGE)) + awsLogsRecordExtraSize
	logAwsFormat := &cloudwatchlogs.InputLogEvent{
		Message:   aws.String(r.HOSTNAME + " " + r.SYSLOG_IDENTIFIER + " [" + r.PID + "]: " + r.MESSAGE),
		Timestamp: aws.Int64(epochNano / int64(time.Microsecond)),
	}
	return logAwsFormat, msgSize, nil
}

func (l *LogClient) Write(r lib.Record) {
	select {
	case l.chAppend <- r:
	default:
	}
}

func (l *LogClient) Flush() {
	if !l.chTimeout.Stop() {
		select {
		case <-l.chTimeout.C:
		default:
		}
	}
	defer l.chTimeout.Reset(inputTimeout)

	if len(l.buffer) == 0 {
		return
	}

	output, err := l.svc.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogEvents:     l.buffer,
		LogGroupName:  aws.String(l.logGroupName),
		LogStreamName: aws.String(l.logStreamName),
		SequenceToken: l.nextSequenceToken,
	})
	if err != nil {
		log.Printf("sytoco/awsCloudwatchLogs ERROR: %s", err)
		return
	}
	l.nextSequenceToken = output.NextSequenceToken
	l.buffer = l.buffer[:0]
	l.bufferSize = 0
}

func (l *LogClient) Exit() {
	close(l.chAppend)
	l.chTimeout.Stop()
	l.chFlush <- true
}
