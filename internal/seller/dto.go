package seller

type TokenPayload struct {
	SellerID string
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
