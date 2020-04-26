package user

import "github.com/syafdia/go-clean-arch/internal/usecase"

type GetUser struct {
}

func NewGetUserUseCase() usecase.UseCase {
	return &GetUser{}
}

func (s *GetUser) Execute(input usecase.Input) (usecase.Output, error) {
	return nil, nil
}
