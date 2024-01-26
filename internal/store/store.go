package store

import (
	"fmt"

	"github.com/tasks-hub/users-service/internal/config"
)

// NewUserStoreFactory creates a UserStore based on the specified database type
func NewUserStoreFactory(databaseType string, cfg config.Database) (UserStore, error) {
	switch databaseType {
	case "postgres":
		return NewPostgresUserStore(cfg)
	case "in-memory":
		return NewInMemoryUserStore(), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}
