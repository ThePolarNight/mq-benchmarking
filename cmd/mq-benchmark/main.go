package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ThePolarNight/mq-benchmarking/benchmark"
	"github.com/ThePolarNight/mq-benchmarking/benchmark/mq"
	"github.com/ThePolarNight/mq-benchmarking/test"
)

func newTester(subject string, testLatency bool, msgCount, msgSize int) *test.Tester {
	var messageSender benchmark.MessageSender
	var messageReceiver benchmark.MessageReceiver

	switch subject {
	case "inproc":
		inproc := mq.NewInproc(msgCount, testLatency)
		messageSender = inproc
		messageReceiver = inproc
	case "zeromq":
		zeromq := mq.NewZeromq(msgCount, testLatency)
		messageSender = zeromq
		messageReceiver = zeromq
	case "kafka":
		kafka := mq.NewKafka(msgCount, testLatency)
		messageSender = kafka
		messageReceiver = kafka
	case "rabbitmq":
		rabbitmq := mq.NewRabbitmq(msgCount, testLatency)
		messageSender = rabbitmq
		messageReceiver = rabbitmq
	case "nsq":
		nsq := mq.NewNsq(msgCount, testLatency)
		messageSender = nsq
		messageReceiver = nsq
	case "redis":
		redis := mq.NewRedis(msgCount, testLatency)
		messageSender = redis
		messageReceiver = redis
	case "activemq":
		activemq := mq.NewActivemq(msgCount, testLatency)
		messageSender = activemq
		messageReceiver = activemq
	default:
		return nil
	}

	return &test.Tester{
		subject,
		msgSize,
		msgCount,
		testLatency,
		messageSender,
		messageReceiver,
	}
}

func parseArgs(usage string) (string, bool, int, int) {

	if len(os.Args) < 2 {
		log.Print(usage)
		os.Exit(1)
	}

	test := os.Args[1]
	messageCount := 1000000
	messageSize := 1000
	testLatency := false

	if len(os.Args) > 2 {
		latency, err := strconv.ParseBool(os.Args[2])
		if err != nil {
			log.Print(usage)
			os.Exit(1)
		}
		testLatency = latency
	}

	if len(os.Args) > 3 {
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Print(usage)
			os.Exit(1)
		}
		messageCount = count
	}

	if len(os.Args) > 4 {
		size, err := strconv.Atoi(os.Args[4])
		if err != nil {
			log.Print(usage)
			os.Exit(1)
		}
		messageSize = size
	}

	return test, testLatency, messageCount, messageSize
}

func main() {
	usage := fmt.Sprintf(
		"usage: %s "+
			"{"+
			"inproc|"+
			"zeromq|"+
			"nanomsg|"+
			"kestrel|"+
			"kafka|"+
			"rabbitmq|"+
			"nsq|"+
			"redis|"+
			"activemq"+
			"} "+
			"[test_latency] [num_messages] [message_size]",
		os.Args[0])

	tester := newTester(parseArgs(usage))
	if tester == nil {
		log.Println(usage)
		os.Exit(1)
	}

	tester.Test()
}
