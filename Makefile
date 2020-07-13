
all: ffjson build

ffjson:
	ffjson lib/record.go

build:
	go build