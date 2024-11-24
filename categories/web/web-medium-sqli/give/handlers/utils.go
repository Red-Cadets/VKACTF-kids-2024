package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func readFlag() string {
	data, err := os.ReadFile("flag")
	if err != nil {
		return "Congratulations!!"
	}
	return string(data)
}

func InitSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading secret_key")
	}
	secret_key := os.Getenv("SECRET_KEY")
	return secret_key
}

var secret_key = InitSecretKey()
var Store = sessions.NewCookieStore([]byte(secret_key))

func init() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   2592000,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
}

func GenerateUniqueID() string {
	return uuid.New().String()
}

func GetUserIDfromCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := Store.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
		return "", fmt.Errorf("error retrieving session: %v", err)
	}

	userID, ok := session.Values["user_id"].(string)
	if !ok || userID == "" {
		http.Error(w, "User ID not found in session", http.StatusUnauthorized)
		return "", fmt.Errorf("user ID not found in session or invalid")
	}
	return userID, nil
}
