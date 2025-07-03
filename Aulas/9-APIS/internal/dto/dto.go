package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type GetJWTInput struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}