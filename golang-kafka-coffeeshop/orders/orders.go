package orders

const (
	StatusPending   = "PENDING"
	StatusConfirmed = "CONFIRMED"
	StatusPreparing = "PREPARING"
	StatusReady     = "READY"
	StatusPickedUp  = "PICKED_UP"
	StatusCancelled = "CANCELLED"
	StatusCompleted = "COMPLETED"
)

type Order struct {
	OrderID       string      `json:"order_id"`
	OrderStatus   string      `json:"order_status"`
	Items         []OrderItem `json:"items"`
	TotalAmount   float64     `json:"total_amount"`
	CustomerName  string      `json:"customer_name"`
	CustomerEmail string      `json:"customer_email"`
	CustomerPhone string      `json:"customer_phone"`
	CreatedAt     string      `json:"created_at"`
}

type OrderItem struct {
	ItemID     string  `json:"item_id"`
	ItemName   string  `json:"item_name"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}

func NewOrder() {

}
