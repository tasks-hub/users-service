package store

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tasks-hub/users-service/internal/config"
	"github.com/tasks-hub/users-service/internal/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
func (p *PostgresUserStore) CreateUser(userInput *entities.CreateUserInput) (*entities.User, error) {
	var user entities.User
	query, args, err := sq.
		Insert("users").
		Columns("username", "email", "password").
		Values(userInput.Username, userInput.Email, userInput.Password).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building SQL query for user creation: %v", err)
	}

	err = p.db.QueryRowx(query, args...).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	return &user, nil
}

// GetUserByID retrieves a user from the data store based on the user ID
func (p *PostgresUserStore) GetUserByID(userID string) (*entities.User, error) {
	var user entities.User
	query, args, err := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"id": userID, "deleted_at": nil}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building SQL query for user retrieval by ID: %v", err)
	}

	err = p.db.Get(&user, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %v", err)
	}
	return &user, nil
}

// GetUserByEmail retrieves a user from the data store based on the user email
func (p *PostgresUserStore) GetUserByEmail(userCredentials *entities.UserCredentials) (*entities.User, error) {
	var user entities.User
	query, args, err := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"email": userCredentials.Email, "deleted_at": nil}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building SQL query for user retrieval by ID: %v", err)
	}

	err = p.db.Get(&user, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %v", err)
		}
		return nil, fmt.Errorf("error getting user by ID: %v", err)
	}
	return &user, nil
}

func (p *PostgresUserStore) UpdateUser(user *entities.User) error {
	updateBuilder := sq.Update("users")

	if user.Username != "" {
		updateBuilder = updateBuilder.Set("username", user.Username)
	}
	if user.Email != "" {
		updateBuilder = updateBuilder.Set("email", user.Email)
	}
	if string(user.Password) != "" {
		updateBuilder = updateBuilder.Set("password", user.Password)
	}

	updateBuilder = updateBuilder.
		Where(sq.Eq{"id": user.ID, "deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building SQL query: %v", err)
	}

	stmt, err := p.db.Preparex(query)
	if err != nil {
		return fmt.Errorf("error preparing statement for user update: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

// DeleteUser deletes a user from the data store based on the user ID (soft delete)
func (p *PostgresUserStore) DeleteUser(userID string) error {
	query, args, err := sq.
		Update("users").
		Set("deleted_at", "NOW()").
		Where(sq.Eq{"id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("error building SQL query for soft delete by ID: %v", err)
	}

	_, err = p.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error performing soft delete by ID: %v", err)
	}
	return nil
}
