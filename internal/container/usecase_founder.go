package container

import "kueku/internal/usecase"

// Usecases .
type Usecases interface {
	CakeUsecase() usecase.CakeUsecase
}
