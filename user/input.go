package user

type RegisterUserInput struct {
	Name       string
	Occupation string
	Email      string
	Password   string
}

type LoginInput struct {
	Email    string
	Password string
}
