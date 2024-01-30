package store

import (
	"fmt"

	"github.com/tasks-hub/users-service/internal/config"
)

// NewUserStoreFactory creates a UserStore based on the specified database type
func NewUserStoreFactory(cfg config.Config) (UserStore, error) {
	switch cfg.DatabaseType {
	case "postgres":
		return NewPostgresUserStore(cfg)
	case "in-memory":
		return NewInMemoryUserStore(), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.DatabaseType)
	}
}
