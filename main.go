package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"

	"github.com/gabrielperezs/sytoco/input/systemdJournald"
	"github.com/gabrielperezs/sytoco/output/awsCloudwatchLogs"
)

var (
	config     *Config
	configFile string
	done       = make(chan bool, 1)
)

type Config struct {
	AwsCloudwatchLogs *awsCloudwatchLogs.Config
	SystemdJournald   *systemdJournald.Config
}

func main() {

	flag.StringVar(&configFile, "config", "/usr/local/etc/systoco.conf", "Sytoco file config (toml format)")
	flag.Parse()

	config := &Config{
		AwsCloudwatchLogs: &awsCloudwatchLogs.Config{},
	}
	if _, err := toml.DecodeFile(configFile, config); err != nil {
		log.Panicf("Reading config file: %s", err)
	}

	signals()

	logCli, err := awsCloudwatchLogs.New(config.AwsCloudwatchLogs)
	if err != nil {
		log.Panicf("Starting awsCloudwatchLogs: %s", err)
	}

	s := systemdJournald.New(config.SystemdJournald, logCli)

	<-done
	s.Exit()
	logCli.Exit()
}

func signals() {
	chSign := make(chan os.Signal, 10)
	signal.Notify(chSign, syscall.SIGHUP, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGKILL, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			switch <-chSign {
			default:
				log.Printf("Closing...")
				done <- true
			}
		}
	}()
}
