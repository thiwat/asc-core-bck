package order

type OrderStatus string

const (
	PendingPayment OrderStatus = "pending_payment"
	Paid                       = "paid"
	Cancelled                  = "cancelled"
	Completed                  = "completed"
)
