package user

import (
	"database/sql"
	 
	"github.com/canalesb93/user-server/internal/database"
)

// UserRepository is a struct that represents a repository for managing users in the database.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository returns a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// SaveUser saves the given user to the database.
func (ur *UserRepository) SaveUser(user *User) error {
	if user.ID == 0 {
        // Generate a new random ID for the user.
        user.ID = database.GenerateRandomID()
    }

	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
        INSERT INTO users (id, name, email)
        VALUES (?, ?, ?)
        ON CONFLICT (id) DO UPDATE SET name=?, email=?
    `)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Name, user.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetUserByEmail retrieves the user with the given email from the database.
func (ur *UserRepository) GetUserByEmail(email string) (*User, error) {
	stmt, err := ur.db.Prepare(`
        SELECT id, name, email FROM users WHERE email=?
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	user, err := ur.scanUser(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves the user with the given ID from the database.
func (ur *UserRepository) GetUserByID(id int64) (*User, error) {
	stmt, err := ur.db.Prepare(`
        SELECT * FROM users WHERE id=?
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	user, err := ur.scanUser(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves all users from the database.
func (ur *UserRepository) GetAllUsers() ([]*User, error) {
	rows, err := ur.db.Query(`
        SELECT * FROM users
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user, err := ur.scanUserFromRows(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// scanUser is a helper function that scans a single row from a query result into a User struct.
func (ur *UserRepository) scanUser(row *sql.Row) (*User, error) {
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// scanUser is a helper function that scans a single row from a query result into a User struct.
func (ur *UserRepository) scanUserFromRows(rows *sql.Rows) (*User, error) {
	var user User
	err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}