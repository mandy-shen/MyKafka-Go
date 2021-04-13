package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "neu-kafka"
	brokerAddress = "localhost:9092"
)

func main() {

	if len(os.Args) != 2 || !strings.HasPrefix(os.Args[1], "add-row") {
		fmt.Printf("wrong input. example: \"add-row neu-student, My Name, ID0000001\"\n")
		return
	}

	// "add-row neu-student, My Name, ID0000001"
	// _msg := "add-row neu-student, My Name, ID0000001"
	_msg := os.Args[1]

	go cmd("zookeeper-server-start.sh", "zookeeper.properties", "zookeeper")

	seconds := 0.5
	time.Sleep(time.Duration(seconds) * time.Second)

	go cmd("kafka-server-start.sh", "server.properties", "kafka")

	seconds = 6
	time.Sleep(time.Duration(seconds) * time.Second)

	ctx := context.Background()
	produce(ctx)

	send(_msg, ctx)
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

func cmd(sh, prop, name string) {
	cmd := exec.Command("/kafka_2.12-2.7.0/bin/"+sh, "-daemon", "/kafka_2.12-2.7.0/config/"+prop)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(name+".Run() failed with %s\n", err)
	}
	fmt.Printf(name + " is running...\n")
}
