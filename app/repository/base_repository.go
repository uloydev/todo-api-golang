package repository

type BaseRepository[T any, Filter any] interface {
	Insert(entity T) T

	FindAll(*Filter) []T

	FindById(id uint) T
}
