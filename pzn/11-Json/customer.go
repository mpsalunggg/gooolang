package main

type Customer struct {
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Age        int      `json:"age"`
	Married    bool     `json:"married"`
	Hobbies    []string `json:"hobbies"`
}
