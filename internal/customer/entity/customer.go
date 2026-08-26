package entity

type Customer struct {
	id             string
	email          string
	hashedPassword string
}

func NewCustomer(
	ID string,
	email string,
	hashedPassword string,
) *Customer {
	return &Customer{
		id:             ID,
		email:          email,
		hashedPassword: hashedPassword,
	}
}

func (s *Customer) ID() string             { return s.id }
func (s *Customer) Email() string          { return s.email }
func (s *Customer) HashedPassword() string { return s.hashedPassword }
