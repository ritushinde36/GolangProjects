package order_operations

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ritushinde36/GolangProjects/golang-kafka-coffeeshop/kafka_operations"
	"github.com/ritushinde36/GolangProjects/golang-kafka-coffeeshop/orders"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder orders.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderInBytes, err := json.Marshal(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = kafka_operations.ProduceMessage("t_orders", orderInBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Order created and sent to Kafka successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
