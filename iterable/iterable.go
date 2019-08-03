package iterable

import "errors"

type Iterator interface {
	Next() (interface{}, error)
}

type Iterable <-chan interface{}

func (iter Iterable) Next() (interface{}, error) {
	if next, ok := <-iter; ok {
		return next, nil
	}

	return nil, errors.New("error")

}

func New(any interface{}) (Iterable, error) {
	switch any := any.(type) {
	case []interface{}:
		data := make(chan interface{}, len(any))
		go func() {
			for _, val := range any {
				data <- val
			}
			close(data)
		}()
		return Iterable(data), nil
	case chan interface{}:
		return Iterable(any), nil

	default:
		data := make(chan interface{}, 1)
		go func() {
			data <- any
		}()
		return Iterable(data), nil

	}
}
