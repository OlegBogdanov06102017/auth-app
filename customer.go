package authapp

type Customer struct {
	ID    int    `json:"id"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}
