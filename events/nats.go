package events

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/working/go-clean-architecture/domain"
)

type NatsEventStore struct {
	nc                      *nats.Conn
	meowCreatedSubscription *nats.Subscription
	meowCreatedChan         chan MeowCreatedMessage
}

func NewNats(natsURL string) (*NatsEventStore, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	return &NatsEventStore{nc: nc}, nil
}

func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.meowCreatedSubscription != nil {
		if err := e.meowCreatedSubscription.Unsubscribe(); err != nil {
			log.Fatal(err)
		}
	}
	close(e.meowCreatedChan)
}

func (e *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (e *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

func (e *NatsEventStore) PublishMeowCreated(meow domain.Meow) error {
	m := MeowCreatedMessage{meow.Id, meow.Body, meow.CreatedAt}
	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventStore) OnMeowCreated(f func(MeowCreatedMessage)) (err error) {
	m := MeowCreatedMessage{}
	e.meowCreatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		if err := e.readMessage(msg.Data, &m); err != nil {
			log.Fatal(err)
		}
		f(m)
	})
	return

}

func (e *NatsEventStore) SubscribeMeowCreated() (<-chan MeowCreatedMessage, error) {
	m := MeowCreatedMessage{}
	e.meowCreatedChan = make(chan MeowCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)

	var err error
	e.meowCreatedSubscription, err = e.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}

	//Decode Message
	go func() {
		for {
			select {
			case msg := <-ch:
				if err := e.readMessage(msg.Data, &m); err != nil {
					log.Fatal(err)
				}
				e.meowCreatedChan <- m
			}
		}

	}()
	return (<-chan MeowCreatedMessage)(e.meowCreatedChan), nil

}
