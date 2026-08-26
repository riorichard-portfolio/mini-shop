package customer

type TokenPayload struct {
	CustomerID string
}

type RegisterInput struct {
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	AccessToken string
}
