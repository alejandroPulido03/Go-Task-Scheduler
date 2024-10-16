package dtos

type Parser interface{
	ToEntity() (interface{}, error)
	ToRaw() (interface{}, error)
}
