package handlers

import (
	"context"
	"database/sql"
	"deploy/database"
	"deploy/models"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var tmpl *template.Template

var tmplOrders = template.Must(template.ParseFiles("static/orders.html"))
var tmplOrderHistory = template.Must(template.ParseFiles("static/orders_history.html"))

func ShopHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	rows, err := database.DB.Query("SELECT id, name, price, image_url FROM items")
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.ImageURL); err != nil {
			http.Error(w, "Error scanning item", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	if len(items) == 0 {
		log.Println("No products found")
	}

	err = tmpl.Execute(w, struct {
		Items []models.Item
	}{
		Items: items,
	})
	if err != nil {
		log.Println(err, nil)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	itemIDStr := r.URL.Query().Get("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var name string
	var image_url string
	var quantity int
	database.DB.QueryRow("SELECT name, image_url FROM items WHERE id = $1", itemID).Scan(&name, &image_url)
	err = database.DB.QueryRow("SELECT quantity FROM cart WHERE user_id = $1 AND item_id = $2", userID, itemID).Scan(&quantity)
	allowedItem := itemID == 1 || itemID == 2 || itemID == 3

	if err == sql.ErrNoRows && allowedItem {
		_, err = database.DB.Exec("INSERT INTO cart (user_id, item_id, name, image_url, quantity) VALUES ($1, $2, $3, $4, 1)", userID, itemID, name, image_url)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
	} else if err == nil {
		_, err = database.DB.Exec("UPDATE cart SET quantity = quantity + 1 WHERE user_id = $1 AND item_id = $2", userID, itemID)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"success":     true,
		"newQuantity": quantity + 1,
	}
	json.NewEncoder(w).Encode(response)

}

func RemoveCartHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	itemIDStr := r.URL.Query().Get("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var quantity int
	err = database.DB.QueryRow("SELECT quantity FROM cart WHERE item_id = $1 AND user_id = $2", itemID, userID).Scan(&quantity)
	if quantity == 1 {
		_, err = database.DB.Exec("DELETE FROM cart WHERE item_id = $1 AND user_id = $2", itemID, userID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Delete item error", http.StatusInternalServerError)
		}
	}
	if err == nil {
		_, err = database.DB.Exec("UPDATE cart SET quantity = quantity - 1 WHERE item_id = $1 AND user_id = $2", itemID, userID)
		if err != nil {
			http.Error(w, "Database error remove", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Database error remove", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"success":     true,
		"newQuantity": quantity - 1,
	}
	json.NewEncoder(w).Encode(response)
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	order := models.Order{
		Name:      r.FormValue("name"),
		Phone:     r.FormValue("phone"),
		Address:   r.FormValue("address"),
		Comment:   r.FormValue("comment"),
		PromoCode: r.FormValue("promo_code"),
	}

	rows, err := database.DB.Query("SELECT item_id, name, image_url, quantity FROM cart WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, "Failed to load data from cart", http.StatusInternalServerError)
	}
	defer rows.Close()

	var cart []models.Cart
	for rows.Next() {
		var item models.Cart
		if err := rows.Scan(&item.ItemID, &item.Name, &item.ImageURL, &item.Quantity); err != nil {
			http.Error(w, "Failed read from cart", http.StatusInternalServerError)
			return
		}
		cart = append(cart, item)
	}

	itemsJSON, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, "Failed download items", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO orders (user_id, name, phone, address, comment, promo_code, items, is_delivered, delivery_time) VALUES ($1, $2, $3, $4, $5, $6, $7, 0, 5)",
		userID, order.Name, order.Phone, order.Address, order.Comment, order.PromoCode, itemsJSON)
	if err != nil {
		http.Error(w, "Failed to place order", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("DELETE FROM cart WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, "Failed clear cart", http.StatusInternalServerError)
		return
	}

	go func() {
		time.Sleep(5 * time.Second)
		database.DB.Exec("UPDATE orders SET is_delivered = 1 WHERE id = currval('orders_id_seq') AND user_id = $1", userID) //check
	}()

	http.Redirect(w, r, "/orders", http.StatusFound)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	profile := models.Profile{
		Name:    r.FormValue("name"),
		Address: r.FormValue("address"),
		Phone:   r.FormValue("phone"),
	}

	_, err = database.DB.Exec(`
		INSERT INTO user_profile (user_id, name, address, phone)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id)
		DO UPDATE SET
			name = EXCLUDED.name,
			address = EXCLUDED.address,
			phone = EXCLUDED.phone
	`, userID, profile.Name, profile.Address, profile.Phone)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/cart", http.StatusFound)
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var profile models.Profile
	err = database.DB.QueryRow("SELECT name, address, phone FROM user_profile WHERE user_id = $1", userID).Scan(
		&profile.Name, &profile.Address, &profile.Phone,
	)

	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Failed to load profile data", http.StatusInternalServerError)
		return
	}

	rows, err := database.DB.Query("SELECT item_id, name, image_url, quantity FROM cart WHERE user_id = $1", userID)
	if err != nil {
		http.Error(w, "Error fetching cart items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cart []models.Cart
	for rows.Next() {
		var cartItem models.Cart
		if err := rows.Scan(&cartItem.ItemID, &cartItem.Name, &cartItem.ImageURL, &cartItem.Quantity); err != nil {
			http.Error(w, "Error processing cart item", http.StatusInternalServerError)
			return
		}
		cart = append(cart, cartItem)
	}

	tmpl, err := template.ParseFiles("./static/cart.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct {
		Profile models.Profile
		Cart    []models.Cart
	}{
		Profile: profile,
		Cart:    cart,
	})
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	rows, err := database.DB.Query("SELECT id, name, phone, address, comment, promo_code, items, is_delivered, delivery_time FROM orders WHERE is_delivered = 0 AND user_id = $1", userID)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Error loading active orders", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		var itemsData []byte
		if err := rows.Scan(&order.ID, &order.Name, &order.Phone, &order.Address, &order.Comment, &order.PromoCode, &itemsData, &order.IsDelivered, &order.DeliveryTime); err != nil {
			log.Println("Scan Error:", err)
			http.Error(w, "Error reading completed order", http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(itemsData, &order.Items); err != nil {
			log.Println("JSON decoding error:", err)
			return
		}
		orders = append(orders, order)
	}

	if err := tmplOrders.Execute(w, orders); err != nil {
		log.Println("Error when rendering template", err)
	}
}

func OrderHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	rows, err := database.DB.Query("SELECT id, name, phone, address, comment, promo_code, items, is_delivered, delivery_time FROM orders WHERE is_delivered = 1 AND user_id = $1", userID)
	if err != nil {
		http.Error(w, "Error loading completed orders", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var completedOrders []models.Order
	for rows.Next() {
		var order models.Order
		var itemsData []byte
		if err := rows.Scan(&order.ID, &order.Name, &order.Phone, &order.Address, &order.Comment, &order.PromoCode, &itemsData, &order.IsDelivered, &order.DeliveryTime); err != nil {
			log.Println("Scan Error:", err)
			http.Error(w, "Error reading completed order", http.StatusInternalServerError)
			return
		}

		if order.DeliveryTime > 5 {
			order.Comment = readFlag()
		}

		if err := json.Unmarshal(itemsData, &order.Items); err != nil {
			log.Println("JSON decoding error:", err)
			return
		}
		completedOrders = append(completedOrders, order)
	}

	if err := tmplOrderHistory.Execute(w, completedOrders); err != nil {
		log.Println("Error when rendering order history template:", err)
	}
}

func UpdateAddressHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserIDfromCookie(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	orderID := r.FormValue("order_id")
	newAddress := r.FormValue("new_address")

	query := fmt.Sprintf("UPDATE orders SET address = '%s' \n", newAddress)
	query = query + "WHERE id = $1 AND is_delivered = 0 AND user_id = $2"
	_, err = database.DB.ExecContext(ctx, query, orderID, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	response := map[string]bool{"success": true}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

	http.Redirect(w, r, "/orders", http.StatusSeeOther)
}
