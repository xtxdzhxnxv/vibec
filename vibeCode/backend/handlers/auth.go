package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vibecode/backend/database"
	"vibecode/backend/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Printf("Получены данные пользователя: %s, %s", user.Username, user.Password)

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = database.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
	log.Printf("Someone registered account: %s", user.Username)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Printf("полученны данные пользователя для входа: %s, %s", user.Username, user.Password)

	query := "SELECT id, username FROM users WHERE username = ? AND password = ?"
	err = database.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID, &user.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged in successfully"))
	log.Printf("Someone logged in : %s", user.Username)

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM users WHERE username = ? AND password = ?"
	_, err = database.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account deleted successfully"))
	log.Printf("Account deleted successfully: %s", user.Username)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET username = ?, password = ? WHERE id = ?"
	_, err = database.DB.Exec(query, user.Username, user.Password, user.ID)
	if err != nil {
		http.Error(w, "Failed to update account", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account updated successfully"))
	log.Printf("Account updated successfully: %s, %s", user.Username, user.Password)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := "SELECT id, username FROM users"
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch accounts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var accounts []models.User
	for rows.Next() {
		var account models.User
		err = rows.Scan(&account.ID, &account.Username)
		if err != nil {
			http.Error(w, "Failed to scan account", http.StatusInternalServerError)
			return
		}
		accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "Failed to fetch accounts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
	log.Print("Someone requested all accounts")
}
