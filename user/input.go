package user

type RegisterUserInput struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginInput struct {
	Email    string
	Password string
}
