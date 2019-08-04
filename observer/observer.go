package observer

type Observer struct {
	onNextHandler     func(interface{})
	onCompleteHandler func()
	onErrorHandler    func(err error)
}

func New() Observer {
	return DefaultObserver
}

func (ob Observer) OnNextHandler(onNext func(item interface{})) Observer {
	ob.onNextHandler = onNext
	return ob
}

func (ob Observer) OnErrorHandler(onError func(error)) Observer {
	ob.onErrorHandler = onError
	return ob
}

var DefaultObserver = Observer{
	onNextHandler: func(i interface{}) {

	},
	onCompleteHandler: func() {

	},
	onErrorHandler: func(err error) {

	},
}

func (ob Observer) apply(item interface{}) {
	switch item := item.(type) {
	case error:
		ob.onErrorHandler(item)
	default:
		ob.onNextHandler(item)

	}
}

func (ob Observer) OnNext(item interface{}) {
	switch item := item.(type) {
	case error:
		return
	default:
		if ob.onNextHandler != nil {
			ob.onNextHandler(item)
		}

	}
}

func (ob Observer) OnError(item interface{}) {
	switch item := item.(type) {
	case error:
		if ob.onErrorHandler != nil {
			ob.onErrorHandler(item)
		}

	}

}

func (ob Observer) OnDone() {
	if ob.onCompleteHandler != nil {
		ob.onCompleteHandler()
	}

}
