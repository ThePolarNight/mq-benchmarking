package test

import (
	"log"

	"github.com/ThePolarNight/mq-benchmarking/benchmark"
)

type Tester struct {
	Name         string
	MessageSize  int
	MessageCount int
	TestLatency  bool
	benchmark.MessageSender
	benchmark.MessageReceiver
}

func (tester Tester) Test() {
	log.Printf("Begin %s test", tester.Name)
	tester.Setup()
	defer tester.Teardown()

	if tester.TestLatency {
		tester.testLatency()
	} else {
		tester.testThroughput()
	}

	log.Printf("End %s test", tester.Name)
}

func (tester Tester) testThroughput() {
	receiver := benchmark.NewReceiveEndpoint(tester, tester.MessageCount)
	sender := &benchmark.SendEndpoint{MessageSender: tester}
	sender.TestThroughput(tester.MessageSize, tester.MessageCount)
	receiver.WaitForCompletion()
}

func (tester Tester) testLatency() {
	receiver := benchmark.NewReceiveEndpoint(tester, tester.MessageCount)
	sender := &benchmark.SendEndpoint{MessageSender: tester}
	sender.TestLatency(tester.MessageSize, tester.MessageCount)
	receiver.WaitForCompletion()
}
