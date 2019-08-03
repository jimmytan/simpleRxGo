package main

import (
	"fmt"
	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {

	score := 9

	onNext := handlers.NextFunc(func(item interface{}) {
		if num, ok := item.(int); ok {
			score += num
		}
	})

	onDone := handlers.DoneFunc(func() {
		score *= 2
	})

	watcher := observer.New(onNext, onDone)

	it, _ := iterable.New([]interface{}{1, 2, 3, 4, 5})
	source := observable.From(it)
	sub := source.Subscribe(watcher)
	<-sub
	fmt.Println(score)

}
