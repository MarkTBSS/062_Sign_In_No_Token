package middlewaresUsecases

import _pkgMiddlewaresMiddlewaresRepositories "github.com/MarkTBSS/062_Sign_In_No_Token/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
