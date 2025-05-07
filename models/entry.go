package models

// The Entry struct represents a time entry
type Entry struct {
    Customer string
    Process  string
    Task     string
    Minutes  int
    Date     string
}

// The Customer struct represents a customer
type Customer struct {
	Name string
}

type Process struct {
    Customer string
    Process string
}
