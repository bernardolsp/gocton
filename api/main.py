from fastapi import FastAPI
from pydantic import BaseModel, TypeAdapter
from typing import List, Dict

app = FastAPI()

class Step(BaseModel):
    run: str

class Job(BaseModel):
    steps: List[Step]

class TaskFile(BaseModel):
    version: str
    jobs: Dict[str, Job]
    workflow: List[str]

adapter = TypeAdapter(TaskFile)

@app.get("/")
def main():
    print("Hello! I guess")

@app.post("/workflow")
def post_workflow():
    print("workflow")
