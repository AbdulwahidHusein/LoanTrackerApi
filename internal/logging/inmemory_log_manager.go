package logging

import (
	"LoanTrackerApi/internal/entity"
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InMemoryLogManager struct {
	mu   sync.Mutex
	logs []entity.Log
}

func NewInMemoryLogManager() *InMemoryLogManager {
	return &InMemoryLogManager{}
}

func (lm *InMemoryLogManager) AddLog(ctx context.Context, category, message string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	log := entity.Log{
		ID:        generateID(),
		Timestamp: time.Now(),
		Category:  category,
		Message:   message,
	}
	lm.logs = append(lm.logs, log)
	return nil
}

func (lm *InMemoryLogManager) GetLogs(ctx context.Context) ([]entity.Log, error) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	return lm.logs, nil
}

// Generate a unique ID for logs
func generateID() string {
	return primitive.NewObjectID().Hex()
}
