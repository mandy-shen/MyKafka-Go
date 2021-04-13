package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "neu-kafka"
	brokerAddress = "kafka:9092"
)

func main() {

	if len(os.Args) != 2 || !strings.HasPrefix(os.Args[1], "add-row") {
		fmt.Printf("wrong input. example: \"add-row neu-student, My Name, ID0000001\"\n")
		return
	}

	// "add-row neu-student, My Name, ID0000001"
	// _msg := "add-row neu-student, My Name, ID0000001"
	_msg := os.Args[1]

	ctx := context.Background()
	produce(ctx)

	send(_msg, ctx)
	//consume(ctx)
}

func produce(ctx context.Context) {
	conn, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
}

func send(_msg string, ctx context.Context) {
	tmp := _msg[strings.Index(_msg, " "):]
	data := strings.Split(tmp, ",")
	msg := fmt.Sprintf("Object : %s\nName : %s\nUser ID: %s", strings.TrimSpace(data[0]), strings.TrimSpace(data[1]), strings.TrimSpace(data[2]))

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		//Logger:  log.New(os.Stdout, "[kafka writer] ", 0),
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(msg),
		Value: []byte(msg),
	})
	if err != nil {
		panic(err.Error())
	}
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	fmt.Println(string(msg))
}

func consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		//Logger:  log.New(os.Stdout, "[kafka reader] ", 0),
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("received: ", string(msg.Value))
	}
}
