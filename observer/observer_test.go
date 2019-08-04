package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateObserver(t *testing.T) {
	assert := assert.New(t)
	observer := New()

	assert.IsType(Observer{}, observer)
	assert.NotNil(observer.onErrorHandler)
	assert.NotNil(observer.onNextHandler)
	assert.NotNil(observer.onCompleteHandler)

}

func TestObserver_OnNext(t *testing.T) {
	assert := assert.New(t)
	observer := New()
	const OnNextValue = "next"
	onNextTest := ""

	observer.onNextHandler = func(item interface{}) {
		if item, ok := item.(string); ok {
			onNextTest = item
		}
	}

	observer.OnNext(OnNextValue)
	assert.Equal(OnNextValue, onNextTest)
}
