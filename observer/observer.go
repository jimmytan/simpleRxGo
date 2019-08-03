package observer

type Observer struct {
	OnNextHandler     func(interface{})
	OnCompleteHandler func()
	OnErrorHandler    func(err error)
}

func (ob Observer) apply(item interface{}) {
	switch item := item.(type) {
	case error:
		ob.OnErrorHandler(item)
	default:
		ob.OnNextHandler(item)

	}
}
