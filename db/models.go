package db

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Customer struct {
	Id        string `json:"id"`
	Address   string `json:"address"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

type CreditScore struct {
	Equifax    int    `json:"equifax"`
	Experian   int    `json:"experian"`
	TransUnion int    `json:"trans_union"`
	CustomerId string `json:"customer_id"`
}
