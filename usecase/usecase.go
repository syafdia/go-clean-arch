package usecase

import "context"

type Input interface{}

type Output interface{}

type UseCase interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
