package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// User representa um usuário no sistema
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Storage interface define os métodos para armazenamento de usuários
type Storage interface {
	Create(user User) (User, error)
	GetByID(id int) (User, error)
	GetAll() ([]User, error)
	Update(id int, user User) (User, error)
	Delete(id int) error
}

// MemoryStorage é uma implementação em memória do Storage
type MemoryStorage struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

// NewMemoryStorage cria uma nova instância de MemoryStorage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *MemoryStorage) Create(user User) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = s.nextID
	s.users[user.ID] = user
	s.nextID++

	return user, nil
}

func (s *MemoryStorage) GetByID(id int) (User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *MemoryStorage) GetAll() ([]User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return users, nil
}

func (s *MemoryStorage) Update(id int, user User) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return User{}, fmt.Errorf("user not found")
	}

	user.ID = id
	s.users[id] = user

	return user, nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("user not found")
	}

	delete(s.users, id)
	return nil
}

// Handlers

func handleCreateUser(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newUser, err := store.Create(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}
}

func handleGetUser(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := store.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func handleGetAllUsers(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := store.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func handleUpdateUser(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid user ID", http.StatusBadRequest)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedUser, err := store.Update(id, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedUser)
	}
}

func handleDeleteUser(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid user ID", http.StatusBadRequest)
			return
		}

		if err := store.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	store := NewMemoryStorage()

	mux := http.NewServeMux()

	// Rotas da API
	mux.HandleFunc("POST /users", handleCreateUser(store))
	mux.HandleFunc("GET /users", handleGetAllUsers(store))
	mux.HandleFunc("GET /users/{id}", handleGetUser(store))
	mux.HandleFunc("PUT /users/{id}", handleUpdateUser(store))
	mux.HandleFunc("DELETE /users/{id}", handleDeleteUser(store))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
