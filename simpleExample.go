package main

import (
	"errors"
	"fmt"
	"goRxExample/iterable"
	"goRxExample/observable"
	"goRxExample/observer"
)

func main() {
	onNext := func(item interface{}) {
		fmt.Println(item)
	}

	onError := func(err error) {
		fmt.Printf("have error %v\n", err)
	}
	observer := observer.New().OnNextHandler(onNext).OnErrorHandler(onError)

	iter := iterable.New([]interface{}{1, 2, 4, 5, errors.New("doom")})

	obser := observable.From(iter)

	sub := obser.Subscribe(observer)
	<-sub
}
