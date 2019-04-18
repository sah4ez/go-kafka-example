package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaAddr  = flag.String("kafka", "127.0.0.1:9092", "kafka address")
	topic      = flag.String("topic", "test", "kafka topic")
	partition  = flag.Int("partition", 0, "topic partition")
	subscriber = flag.Bool("subscriber", false, "subscriber on topic")
	publisher  = flag.Bool("publisher", false, "publishser to topic")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println(*subscriber, *publisher)

	if *subscriber && *publisher {
		fmt.Println("should be selecte type application -subscriber or -publisher")
		os.Exit(2)
	}

	if *subscriber {
		fmt.Println("start subscriber")
		// make a new reader that consumes from topic-A
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{*kafkaAddr},
			GroupID:  fmt.Sprintf("g-%d", *partition),
			Topic:    *topic,
			MaxWait:  100 * time.Millisecond,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		})
		for {
			ctx := context.Background()
			m, err := r.ReadMessage(ctx)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			r.CommitMessages(ctx, m)
		}

		r.Close()
	}

	if *publisher {
		fmt.Println("start publisher")
		// make a writer that produces to topic-A, using the least-bytes distribution
		w := kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{*kafkaAddr},
			Topic:   *topic,
		})

		for i := 0; i < 1000; i++ {
			msgs := []kafka.Message{}
			for j := 0; j < 1000; j++ {
				msg := kafka.Message{
					Key:   []byte(fmt.Sprintf("k-%d-%d", i, j)),
					Value: []byte(fmt.Sprintf("v-%d-%d", i, j)),
					Time:  time.Now(),
				}
				msgs = append(msgs, msg)
			}
			err := w.WriteMessages(context.Background(), msgs...)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(i)
		}

		time.Sleep(10 * time.Second)
		w.Close()
	}
}
