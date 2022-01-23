package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"pack.ag/amqp"
)

func main() {
	// Create client
	client, err := amqp.Dial("amqp://localhost:5672",
		amqp.ConnSASLPlain("admin", "admin"),
	)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()

	// Send a message
	{
		// Create a sender
		sender1, err := session.NewSender(
			amqp.LinkTargetAddress("/queue-name1"),
		)
		sender2, err := session.NewSender(
			amqp.LinkTargetAddress("/queue-name2"),
		)
		if err != nil {
			log.Fatal("Creating sender link:", err)
		}

		ctx, cancel := context.WithTimeout(ctx, 120*time.Second)

		// Send message
		for i := 0 ; i< 100 ; i++{
			time.Sleep(1*time.Second)
			err = sender1.Send(ctx, amqp.NewMessage([]byte(fmt.Sprintf("test%v",i+1))))
			if err != nil {
				log.Fatal("Sending message:", err)
			}
			err = sender2.Send(ctx, amqp.NewMessage([]byte(fmt.Sprintf("test%v",i+1))))
			if err != nil {
				log.Fatal("Sending message:", err)
			}

		}

		sender1.Close(ctx)
		sender2.Close(ctx)
		cancel()
	}

}