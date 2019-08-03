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
