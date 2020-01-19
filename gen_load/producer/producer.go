package producer

import (
	"fmt"
	"io"
	"sync"
	"time"

	generator "github.com/andreyst/agents-test-bench/gen_load/generator"
	serializer "github.com/andreyst/agents-test-bench/gen_load/serializer"
	"github.com/andreyst/agents-test-bench/gen_load/transport"
)

// Producer generates messages, serializes them with serializer and sends them via transport. When deadline is hit, producer stops.
type Producer struct {
	endpoint string
	// TODO: Pointer
	transport    io.Writer
	deadline     time.Time
	generator    *generator.Generator
	waitGroup    *sync.WaitGroup
	messagesSent int64
	serializer   serializer.Serializer
}

// NewProducer creates a new Producer with specified endpoint.
func NewProducer(trans io.Writer, ser serializer.Serializer, deadline time.Time, rateBytesPerSec int64, waitGroup *sync.WaitGroup) (producer *Producer, err error) {
	producer = new(Producer)
	producer.transport = trans
	producer.serializer = ser
	producer.deadline = deadline
	producer.generator = generator.NewGenerator(rateBytesPerSec)
	producer.waitGroup = waitGroup
	return producer, nil
}

// NewTCPProducer creates a new TCP producer with specified endpoint.
func NewTCPProducer(ser serializer.Serializer, endpoint string, deadline time.Time, rateBytesPerSec int64, waitGroup *sync.WaitGroup) (producer *Producer, err error) {
	trans, err := transport.NewTCPTransport(endpoint)
	if err != nil {
		return nil, err
	}

	producer, err = NewProducer(trans, ser, deadline, rateBytesPerSec, waitGroup)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

// NewFluentBitTCPProducer creates a new FluentBitTcpProducer with specified endpoint.
func NewFluentBitTCPProducer(endpoint string, deadline time.Time, rateBytesPerSec int64, waitGroup *sync.WaitGroup) (producer *Producer, err error) {
	ser := new(serializer.FluentBitMessageSerializer)
	producer, err = NewTCPProducer(ser, endpoint, deadline, rateBytesPerSec, waitGroup)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

// Produce starts a producing cycle.
func (producer *Producer) Produce() {
	defer producer.waitGroup.Done()

	for {
		if time.Now().After(producer.deadline) {
			// TODO: Move to debug output/better statistics output
			fmt.Printf("DEADLINE: %d messages sent", producer.messagesSent)
			return
		}

		p := make([]byte, 128)
		_, err := producer.generator.Read(p)
		if err, ok := err.(*generator.RateLimitedError); ok {
			time.Sleep(err.LimitTTL)
			continue
		} else if err != nil {
			fmt.Printf(err.Error())
			break
		}

		msg := producer.serializer.Serialize(p)
		_, err = producer.transport.Write(msg)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		// TODO: Move to debug output
		// fmt.Printf("Written bytes: %d\n", n)
		producer.messagesSent++
	}
}
