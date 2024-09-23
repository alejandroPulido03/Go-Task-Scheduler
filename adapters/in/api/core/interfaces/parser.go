package interfaces

type Parser interface {
	Parse(T interface{}) (interface{}, error)
}