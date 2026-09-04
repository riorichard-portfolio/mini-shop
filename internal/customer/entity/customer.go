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
) Customer {
	return Customer{
		id:             ID,
		email:          email,
		hashedPassword: hashedPassword,
	}
}

func (c *Customer) ID() string             { return c.id }
func (c *Customer) Email() string          { return c.email }
func (c *Customer) HashedPassword() string { return c.hashedPassword }
