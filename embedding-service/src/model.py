import torch
import torch.nn.functional as F
from transformers import AutoModel, AutoTokenizer


class EmbeddingModel:
  def __init__(self):
    self.model_name = "sentence-transformers/all-MiniLM-L6-v2"
    self.tokenizer = AutoTokenizer.from_pretrained(self.model_name)
    self.model = AutoModel.from_pretrained(self.model_name)

  def mean_pooling(self, model_output, attention_mask):
    token_embeddings = model_output[0]
    input_mask_expanded = attention_mask.unsqueeze(-1).expand(token_embeddings.size()).float()
    return torch.sum(token_embeddings * input_mask_expanded, 1) / torch.clamp(input_mask_expanded.sum(1), min=1e-9)

  def get_embeddings(self, texts):
    encoded_input = self.tokenizer(texts, padding=True, truncation=True, return_tensors="pt")

    with torch.no_grad():
      model_output = self.model(**encoded_input)

    sentence_embeddings = self.mean_pooling(model_output, encoded_input["attention_mask"])

    sentence_embeddings = F.normalize(sentence_embeddings, p=2, dim=1)

    return sentence_embeddings
