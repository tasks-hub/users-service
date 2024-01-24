package store

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tasks-hub/users-service/internal/config"
	"github.com/tasks-hub/users-service/internal/entities"
)

type PostgresUserStore struct {
	db *sqlx.DB
}

func NewPostgresUserStore(cfg config.Database) (*PostgresUserStore, error) {
	postgresHost, err := os.ReadFile(cfg.PostgresHostFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading Postgres Host file: %v", err)
	}

	postgresDB, err := os.ReadFile(cfg.PostgresDBFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading Postgres DB file: %v", err)
	}

	postgresUser, err := os.ReadFile(cfg.PostgresUserFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading Postgres User file: %v", err)
	}

	postgresPassword, err := os.ReadFile(cfg.PostgresPasswordFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading Postgres Password file: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		string(postgresHost),
		string(postgresDB),
		string(postgresUser),
		string(postgresPassword),
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresUserStore{
		db: db,
	}, nil
}

// CreateUser creates a new user in the data store and returns the ID of the created user
func (p *PostgresUserStore) CreateUser(user *entities.CreateUserInput) (string, error) {
	var userID string
	stmt, err := p.db.Preparex(`
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`)
	if err != nil {
		return userID, fmt.Errorf("error preparing statement for user creation: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowx(user.Username, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return userID, fmt.Errorf("error creating user: %v", err)
	}
	return userID, nil
}

// GetUserByID retrieves a user from the data store based on the user ID
func (p *PostgresUserStore) GetUserByID(userID string) (*entities.User, error) {
	var user entities.User
	stmt, err := p.db.Preparex(`
		SELECT * FROM users
		WHERE id = $1
	`)
	if err != nil {
		return nil, fmt.Errorf("error preparing statement for user retrieval by ID: %v", err)
	}
	defer stmt.Close()

	err = stmt.Get(&user, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %v", err)
	}
	return &user, nil
}

// UpdateUser updates the information of an existing user in the data store
func (p *PostgresUserStore) UpdateUser(user *entities.User) error {
	stmt, err := p.db.Preparex(`
		UPDATE users
		SET username=$1, email=$2, password=$3
		WHERE id=$4
	`)
	if err != nil {
		return fmt.Errorf("error preparing statement for user update: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

// DeleteUser deletes a user from the data store based on the user ID
func (p *PostgresUserStore) DeleteUser(userID int) error {
	stmt, err := p.db.Preparex(`
		DELETE FROM users
		WHERE id = $1
	`)
	if err != nil {
		return fmt.Errorf("error preparing statement for user deletion by ID: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("error deleting user by ID: %v", err)
	}
	return nil
}
