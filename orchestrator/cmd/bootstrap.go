package main

import (
	"log"
	"time"

	"github.com/bernardolsp/gocton/orchestrator/internal/parser"
	"github.com/bernardolsp/gocton/orchestrator/pkg/communication"
)

func bootstrap_messager() *communication.Messager {
	var m communication.Messager
	log.Printf("Waiting %v for services to start...", initialDelay)
	time.Sleep(initialDelay)
	m, err := init_messager(m)
	if err != nil {
		log.Fatal("failed initializing rabbit")
	}
	messager := &communication.Messager{
		Channel: m.Channel,
		Context: m.Context,
		Queue:   m.Queue,
	}
	return messager
}

func bootstrap_parser(messager *communication.Messager) {
	parser := &parser.Parser{
		Messager: *messager,
	}
	wg.Add(1)
	go func() {
		parser.Initialize()
		wg.Done()
	}()
	log.Println("This executes after the initialize :-)")
	workflow, err := parser.ParseTaskFile("./example_taskfile.yml")
	if err != nil {
		log.Fatalf("Failed to parse task file: %v", err)
	}
	parser.PrintJobsAndSteps(workflow)
	wg.Wait()
}

func init_messager(m communication.Messager) (communication.Messager, error) {
	for i := 0; i < retryAmount; i++ {
		var err error
		log.Printf("Attempt %d to initialize RabbitMQ", i+1)
		m, err = communication.Initialize(&communication.Communicator{
			Type:             "rabbit",
			ConnectionString: "amqp://guest:guest@rabbitmq:5672/", // todo: grab via env
		})
		if err == nil {
			log.Printf("Successfully initialized RabbitMQ on attempt %d", i+1)
			return m, err
		}
		log.Printf("Attempt %d failed: %v. Retrying...", i+1, err)
		if i < retryAmount-1 { // Don't sleep after the last attempt
			time.Sleep(retryDelay)
		} else {
			log.Fatalf("could not connect reliably to rabbit")
		}
	}
	return m, nil
}
