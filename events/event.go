package events

import (
	"github.com/working/go-clean-architecture/domain"
)

type EventStore interface {
	Close()
	PublishMeowCreated(meow domain.Meow) error
	SubscribeMeowCreated() (<-chan MeowCreatedMessage, error)
	OnMeowCreated(f func(MeowCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishMeowCreated(meow domain.Meow) error {
	return impl.PublishMeowCreated(meow)
}

func SubscribeMeowCreated() (<-chan MeowCreatedMessage, error) {
	return impl.SubscribeMeowCreated()
}

func OnMeowCreated(f func(MeowCreatedMessage)) error {
	return impl.OnMeowCreated(f)
}
