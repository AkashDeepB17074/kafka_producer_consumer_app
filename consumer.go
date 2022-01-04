package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
	"bufio"
	"strings"
	"os"
)


func consumer(topic string){
	fmt.Println("Subscriber is running now")
	var broker []string 
	broker = append(broker,"localhost:9092")
	conf := kafka.ReaderConfig{
		Brokers:       	broker,
		Topic:          topic,
		GroupID:        "",
		MaxBytes:       10e6,
		CommitInterval: 30 * time.Second,
		RetentionTime:  time.Minute,
	}
	reader := kafka.NewReader(conf)
	reader.SetOffset(kafka.LastOffset)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			// logger.InfoLog("Error Occured while reading message")
			fmt.Println(err)
			fmt.Println("Error occured while reading messages")
			continue
		}
		message := string(m.Value)
		fmt.Println("[",time.Now().Format(time.RFC822),"] : " , message)
	}
}

func main(){
	fmt.Println("# Enter Topic Name have to show")
	topic, _ := bufio.NewReader(os.Stdin).ReadString('\n')	
	topic = strings.Trim(topic, "\r\n")
	consumer(topic)
}