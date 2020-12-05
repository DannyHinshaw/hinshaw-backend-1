package schemas

// Application users
type AppUser struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Customers to be managed.
type Customer struct {
	Id        string `json:"id"`
	Address   string `json:"address"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Customer credit scores.
type CreditScore struct {
	Equifax    int    `json:"equifax"`
	Experian   int    `json:"experian"`
	TransUnion int    `json:"trans_union"`
	CustomerId string `json:"customer_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
