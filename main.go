package main

import (
	"context"
	"fmt"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func main() {
	const queueName = "<YOUR-QUEUE-NAME>"
	const connectionString = "<YOUR-ENDPOINT-STRING>"

	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connectionString))
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	qm := ns.NewQueueManager()
	qe, err := qm.Get(ctx, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if qe == nil {
		_, err := qm.Put(ctx, queueName)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	q, err := ns.NewQueue(ctx, queueName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = q.Send(context.Background(), servicebus.NewMessageFromString("Hello World!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(q.Name)

}
