package container

import "kueku/internal/domain/cake"


// Repository . 
type Repository interface {
	CakeRepository() cake.Repository
}
