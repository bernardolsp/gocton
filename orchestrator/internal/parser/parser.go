package parser

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

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

func ParseTaskFile(filePath string) (*TaskFile, error) {
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

func PrintJobsAndSteps(taskFile *TaskFile) {
	// Iterate through the workflow to maintain the order of jobs
	for _, jobName := range taskFile.Workflow {
		job, exists := taskFile.Jobs[jobName]
		if !exists {
			log.Printf("Warning: Job '%s' mentioned in workflow but not defined", jobName)
			continue
		}

		steps := make([]string, len(job.Steps))
		for i, step := range job.Steps {
			steps[i] = step.Run
		}

		log.Printf("Job: %s. Steps: %v", jobName, steps)
	}
}
