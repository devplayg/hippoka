package generator

import "github.com/bxcodec/faker"

type Event struct {
	No       int64
	UnixTime int64  `faker:"unix_time"`
	IP       string `faker:"ipv4"`
	Category int    `faker:"boundary_start=1, boundary_end=10"`
	Checksum string `faker:"len=64"`
}

func NewFakeEvent(no int64) *Event {
	e := &Event{}
	err := faker.FakeData(e)
	if err != nil {
		panic(err)
	}
	return e
}
