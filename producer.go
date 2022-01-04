package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"context"
	"bufio"
	"os"
	"strings"
)

func Producer(ctx context.Context, tops string, msg string) {
	topic := tops
	var broker []string 
	broker = append(broker,"localhost:9092")
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: broker,
		Topic:   topic,
	})
	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	fmt.Println("Publisher_publishing message : ")
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte("My key"),
		// create an arbitrary message payload for the value
		Value: []byte(msg),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func main() {
	ctx := context.Background()
	flag := 0


	var topic_name string = "test"
	for{
		fmt.Println(" ------------------------ ")
		fmt.Println("\n# To Enter topic name type ===> /topic <topic_name>")

		if flag != 0{
			fmt.Println("# To message in ", topic_name, " topic type ===> /msg <your message> ")
		}
		
		msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		
		
		if cmd == "/topic" {
			topic_name = strings.TrimSpace(args[1])
			flag = 1
		} else if cmd == "/msg" {
			msg = strings.TrimSpace(strings.Join(args[1:], " "))
			Producer(ctx, topic_name, msg)
		} else {
			fmt.Println("Unknonw Command Entered !")
		}


	}


}


