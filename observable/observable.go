package observable

import (
	"goRxExample/iterable"
	"goRxExample/observer"
	"goRxExample/subscription"
)

type Observable <-chan interface{}

func From(iterable iterable.Iterable) (Observable, error) {
	source := make(chan interface{})

	go func() {
		for {
			val, err := iterable.Next()
			if err != nil {
				break
			}

			source <- val
		}
		close(source)
	}()

	return Observable(source), nil

}

func (observable Observable) Subscribe(observer observer.Observer) <-chan subscription.Subscription {

}
