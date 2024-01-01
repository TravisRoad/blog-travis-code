package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// User represents a user in the system
type User struct {
	Username string
	Password string
}

// AuthManager manages user authentication state
type AuthManager struct {
	users    map[string]User
	LoggedIn map[string]bool
	mutex    sync.Mutex
}

// NewAuthManager creates a new AuthManager
func NewAuthManager() *AuthManager {
	return &AuthManager{
		users:    make(map[string]User),
		LoggedIn: make(map[string]bool),
	}
}

// RegisterUser registers a new user
func (am *AuthManager) RegisterUser(username, password string) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	am.users[username] = User{Username: username, Password: password}
}

// Login authenticates a user and marks them as logged in
func (am *AuthManager) Login(username, password string) bool {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	user, exists := am.users[username]
	if !exists || user.Password != password {
		return false
	}

	am.LoggedIn[username] = true
	return true
}

// IsLoggedIn checks if a user is logged in
func (am *AuthManager) IsLoggedIn(username string) bool {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	return am.LoggedIn[username]
}

// Logout logs out a user
func (am *AuthManager) Logout(username string) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	delete(am.LoggedIn, username)
}

func handleLogin(am *AuthManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("failed to read request body", err)
		}
		u := User{}
		if err := json.Unmarshal(b, &u); err != nil {
			slog.Error("failed to unmarshal user", err)
		}

		if am.Login(u.Username, u.Password) {
			cookie := http.Cookie{
				Name:    "session",
				Value:   u.Username,
				Path:    "/",
				Expires: time.Now().Add(24 * time.Hour),
			}
			http.SetCookie(w, &cookie)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Login successful")
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid credentials")
		}

	}
}

// handleIsLoggedIn handles the check login status request
func handleIsLoggedIn(am *AuthManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if am.IsLoggedIn(cookie.Value) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			u := User{Username: cookie.Value}
			jsonString, _ := json.Marshal(u)
			w.Write(jsonString)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "User is not logged in")
		}
	}
}

// handleLogout handles the logout request
func handleLogout(am *AuthManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		am.Logout(cookie.Value)

		cookie.Expires = time.Unix(0, 0)

		w.WriteHeader(http.StatusOK)
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "Logout successful")
	}
}

func main() {
	authManager := NewAuthManager()

	// Register a sample user
	authManager.RegisterUser("user1", "password123")
	authManager.RegisterUser("admin", "admin")

	// Define API endpoints
	http.HandleFunc("/api/login", handleLogin(authManager))
	http.HandleFunc("/api/islogin", handleIsLoggedIn(authManager))
	http.HandleFunc("/api/logout", handleLogout(authManager))

	// Start the server
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
