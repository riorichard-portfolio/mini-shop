package seller

type Seller struct {
	id             string
	email          string
	hashedPassword string
}

func New(
	ID string,
	email string,
	hashedPassword string,
) *Seller {
	return &Seller{
		id:             ID,
		email:          email,
		hashedPassword: hashedPassword,
	}
}

func (s *Seller) ID() string             { return s.id }
func (s *Seller) Email() string          { return s.email }
func (s *Seller) HashedPassword() string { return s.hashedPassword }
