package magazine

// Subscriber struct
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

// Employee struct
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

// Address struct
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
