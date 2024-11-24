package main

import (
	"fmt"
	"net/http"
	"deploy/handlers"
	"deploy/database"
	"path/filepath"
	"os"
	"time"
	"log"
)


func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := handlers.Store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error getting session: %v", err)
			return
		}

		if session.IsNew {
			session.Values["user_id"] = handlers.GenerateUniqueID()
			session.Values["is_authenticated"] = false
			err := session.Save(r, w)
			if err != nil {
				http.Error(w, "Failed to save session", http.StatusInternalServerError)
				log.Printf("Error saving session: %v", err)
				return
			}
			log.Println("New session created for user")
		}
		next.ServeHTTP(w, r)
	})
}


func main() {
	database.InitDB()
	defer database.DB.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.Redirect(w, r, "/shop", http.StatusFound)
            return
        }
		staticFilePath := filepath.Join("static", r.URL.Path)
		htmlFilePath := filepath.Join("static", r.URL.Path) + ".html"
        if _, err := os.Stat(staticFilePath); err == nil {
            http.ServeFile(w, r, staticFilePath)
            return
        } else if _, err := os.Stat(htmlFilePath); err == nil {

			http.ServeFile(w, r, htmlFilePath)
			return
		}

        http.NotFound(w, r)
    })

	clearDatabase()


    mux.Handle("/cart", sessionMiddleware(http.HandlerFunc(handlers.CartHandler)))
    mux.Handle("/shop", sessionMiddleware(http.HandlerFunc(handlers.ShopHandler)))
	mux.Handle("/orders", sessionMiddleware(http.HandlerFunc(handlers.OrdersHandler)))
	mux.Handle("/orders_history", sessionMiddleware(http.HandlerFunc(handlers.OrderHistoryHandler)))
	mux.HandleFunc("/add_to_cart", handlers.AddToCart)
	mux.HandleFunc("/remove_cart", handlers.RemoveCartHandler)
	mux.HandleFunc("/place_order", handlers.PlaceOrder)
	mux.HandleFunc("/update_profile", handlers.UpdateProfile)
	mux.HandleFunc("/update_address", handlers.UpdateAddressHandler)






	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}


func clearDatabase() {
	database.DB.Exec("DELETE FROM cart")
	database.DB.Exec("DELETE FROM user_profile")
	database.DB.Exec("DELETE FROM orders")

	println("Database fully cleaned at", time.Now().Format(time.RFC3339))

}