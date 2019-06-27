package generator

import (
	"github.com/Shopify/sarama"
	"github.com/devplayg/hippo"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

var (
	loc         = time.UTC
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type Generator struct {
	engine *hippo.Engine

	Name    string
	Version string
	Debug   bool

	Topic     string
	Partition int32
	Brokers   []string

	Count    int64
	DataSize int64
	Interval int
}

func (g *Generator) Start() error {
	g.generate()
	return nil
}

func (g *Generator) Stop() error {
	return nil
}
func (g *Generator) SetEngine(e *hippo.Engine) {
	g.engine = e
}

func (g *Generator) GetName() string {
	return g.Name
}

func (g *Generator) GetVersion() string {
	return g.Version
}

func (g *Generator) IsDebug() bool {
	return g.Debug
}

func (g *Generator) generate() {
	// Create config
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewManualPartitioner
	if g.Partition < 0 {
		config.Producer.Partitioner = sarama.NewHashPartitioner
		//config.Producer.Partitioner = sarama.NewRandomPartitioner
	}
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// Start producer
	producer, err := sarama.NewSyncProducer(g.Brokers, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			g.engine.ErrChan <- err
		}
	}()

	data := randString(g.DataSize)
	var i int64 = 1
	for {
		e := NewEvent(i, data)
		msg := NewMessage(e.encoded)
		m := NewFakeMessage(g.Topic, g.Partition, *msg)
		partition, offset, err := producer.SendMessage(m)
		if err != nil {
			g.engine.ErrChan <- err
			continue
		}
		log.WithFields(log.Fields{
			"topic":     g.Topic,
			"partition": partition,
			"offset":    offset,
		}).Debug()
		i++
		if g.Count > 0 && i > g.Count {
			return
		}
		//time.Sleep(time.Duration(g.Interval) * time.Millisecond)
	}
}

func randString(n int64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
