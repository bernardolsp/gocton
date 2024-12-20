package parser

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bernardolsp/gocton/orchestrator/pkg/communication"
	"gopkg.in/yaml.v3"
)

type Parser struct {
	Messager communication.Messager
}

type TaskFile struct {
	Version  string         `yaml:"version"`
	Jobs     map[string]Job `yaml:"jobs"`
	Workflow []string       `yaml:"workflow"`
}

type Job struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Run string `yaml:"run"`
}

type JobMessage struct {
	Index   int      `json:"index"`
	JobName string   `json:"job_name"`
	Steps   []string `json:"steps"`
}

func (p Parser) Initialize() {
	p.ListenAndParse()
}

func (p Parser) ParseTaskFile(filePath string) (*TaskFile, error) {
	// Read the file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// initialize the struct
	var taskFile TaskFile

	// unmarshal yaml into TaskFile
	err = yaml.Unmarshal(file, &taskFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Workflow: %v", taskFile.Workflow)
	return &taskFile, nil
}

func (p Parser) PrintJobsAndSteps(taskFile *TaskFile) {
	// Iterate through the workflow to maintain the order of jobs
	for index, jobName := range taskFile.Workflow {
		job, exists := taskFile.Jobs[jobName]
		if !exists {
			log.Printf("Warning: Job '%s' mentioned in workflow but not defined", jobName)
			continue
		}

		steps := make([]string, len(job.Steps))
		for i, step := range job.Steps {
			steps[i] = step.Run
		}

		jobMessage := JobMessage{
			Index:   index,
			JobName: jobName,
			Steps:   steps,
		}
		message, err := json.Marshal(jobMessage)
		if err != nil {
			log.Fatalf("Error marshalling job message, %v", err)
		}
		err = p.Messager.SendMessage(p.Messager.Queue, message)
		if err != nil {
			log.Fatalf("Error sending message to messager, %v", err)
		}
	}
}

func (p *Parser) ListenAndParse() {
	// This function will start a Messenger Processor
	msgs, err := p.Messager.Channel.Consume(
		p.Messager.Queue, // queue
		"this",           // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Fatal("Failed to register a consumer. err:", err)
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages.")
	<-forever
}
