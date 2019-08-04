package subscription

import "time"

type Subscription struct {
	SubscriptAt   time.Time
	UnSubscriptAt time.Time
	Error         error
}

func (sub Subscription) Subscribe() {
	sub.SubscriptAt = time.Now()
}

func (sub Subscription) UnSubscribe() {
	sub.UnSubscriptAt = time.Now()
}

func New() Subscription {
	return Subscription{}
}
