package domain

// Login content struct deifinition
type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

// Login usecase methods deifinition
type LoginUsecase interface {
	AuthenticateUser(login Login) error
}

// Login repository methods deifinition
type LoginRepository interface {
	GetLoginByCpf(cpf string) (Login, error)
}
