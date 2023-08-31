package common

type Ptr[T any] interface {
	PtrFields() []any
	*T
}
