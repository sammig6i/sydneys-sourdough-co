from typing import List

from fastapi import FastAPI
from pydantic import BaseModel

from .model import EmbeddingModel

app = FastAPI()
model = EmbeddingModel()


class TextInput(BaseModel):
  texts: List[str]


@app.post("/embed")
async def create_embeddings(input_data: TextInput):
  embeddings = model.get_embeddings(input_data.texts)
  return {"embeddings": embeddings.tolist()}


@app.get("/health")
async def health_check():
  return {"status": "healthy"}
