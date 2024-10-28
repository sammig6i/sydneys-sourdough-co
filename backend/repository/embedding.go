package repository

import (
	"fmt"

	"github.com/sammig6i/sydneys-sourdough-co/pkg/embedding"
)

var embeddingClient *embedding.Client

func InitEmbeddingClient(client *embedding.Client) {
	embeddingClient = client
}

func getEmbedding(text string) ([]float32, error) {
	embeddings, err := embeddingClient.GetEmbeddings([]string{text})
	if err != nil {
		return make([]float32, 384), fmt.Errorf("error getting embeddings, using zero vector: %w", err)
	}
	return embeddings, nil
}
