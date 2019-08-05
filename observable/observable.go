package observable

import (
	"goRxExample/iterable"
	"goRxExample/observable/function"
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

func (ob Observable) Subscribe(observer observer.Observer) <-chan subscription.Subscription {
	done := make(chan subscription.Subscription)
	sub := subscription.New()
	sub.Subscribe()
	go func() {
		for item := range ob {
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

func (ob Observable) Map(mapper function.MapperFunction) Observable {
	out := make(chan interface{})
	go func() {
		for item := range ob {
			out <- mapper(item)
		}
		close(out)
	}()
	return Observable(out)
}

func (ob Observable) Filter(filter function.FilterFunction) Observable {
	out := make(chan interface{})
	go func() {
		for item := range ob {
			if isValid := filter(item); isValid {
				out <- item
			}

		}

		close(out)
	}()
	return Observable(out)
}

func (ob Observable) Distinct(key function.KeyFunction) Observable {
	out := make(chan interface{})
	go func() {
		keySets := make(map[interface{}]bool)
		for item := range ob {
			keyValue := key(item)
			_, ok := keySets[keyValue]
			if !ok {
				out <- item
				keySets[keyValue] = true
			}
		}
		close(out)
	}()

	return Observable(out)
}
