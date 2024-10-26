### gocton

Gocton is a ci-cd tool written in go and python.

The repository has `api`, `orchestrator` and `worker` folders.

The `api` folder contains code for the api server - this is a python fastapi server that listens for webhooks from github and triggers the orchestrator.

The `orchestrator` folder contains code for the orchestrator - this is a go program that listens for webhooks in order to build tasks.

The `worker` folder contains code for the worker - this is a go program that listens for tasks and executes them. They register themselves with the orchestrator and are executed by the orchestrator.

