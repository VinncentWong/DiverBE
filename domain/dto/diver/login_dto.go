package diver

type LoginDto struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
