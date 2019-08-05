package function

type (
	MapperFunction func(any interface{}) interface{}

	FilterFunction func(any interface{}) bool

	KeyFunction func(any interface{}) interface{}
)
