package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/canalesb93/user-server/internal/user"
)

// getUsersHandler retrieves a list of all users.
func (s *Server) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	repo := user.NewUserRepository(s.db.Conn)
	users, err := repo.GetAllUsers()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// createUserHandler creates a new user.
func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u *user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := user.NewUserRepository(s.db.Conn)
	if err := repo.SaveUser(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// getUserHandler retrieves a user by their ID.
func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := user.NewUserRepository(s.db.Conn)
	u, err := repo.GetUserByID( int64(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
