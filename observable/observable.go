package observable

import (
	"goRxExample/iterable"
	"goRxExample/observer"
	"goRxExample/subscription"
)

type Observable <-chan interface{}

func From(iterable iterable.Iterable) Observable {
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

	return Observable(source)

}

func (observable Observable) Subscribe(observer observer.Observer) <-chan subscription.Subscription {
	done := make(chan subscription.Subscription)
	sub := subscription.New()
	sub.Subscribe()
	go func() {
		for item := range observable {
			switch item := item.(type) {
			case error:
				observer.OnError(item)
				sub.Error = item
				sub.UnSubscribe()
				break

			default:
				observer.OnNext(item)
			}
		}
		if sub.Error == nil {
			observer.OnDone()
		}

		done <- sub
	}()

	return done
}
