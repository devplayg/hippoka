package generator

import (
	"github.com/devplayg/hippo"
	"time"
)

var loc = time.UTC

type Generator struct {
	engine *hippo.Engine

	Name    string
	Version string
	Debug   bool

	Topic     string
	Partition int
	Brokers   []string

	Size     int
	Max      int64
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
	loop := false
	var i int64 = 1
	for {
		event := NewFakeEvent(i)
		c.produce(event)

		if loop == false {
			return
		}
	}
}

//func (g *Generator) classify() {
//config := sarama.NewConfig()
//config.Producer.Partitioner = sarama.NewManualPartitioner
//if c.Partition < 0 {
//	config.Producer.Partitioner = sarama.NewHashPartitioner
//	//config.Producer.Partitioner = sarama.NewRandomPartitioner
//}
//config.Producer.RequiredAcks = sarama.WaitForAll
//config.Producer.Return.Successes = true
//producer, err := sarama.NewSyncProducer(c.Brokers, config)
//if err != nil {
//	log.Fatal(err)
//	return
//}
//defer func() {
//	if err := producer.Close(); err != nil {
//		c.engine.ErrChan <- err
//	}
//}()
//
//var messageNo int64 = 1
//e := NewFakeEvent(messageNo)
//b, _ := json.Marshal(e)
//
//message := &sarama.ProducerMessage{
//	Topic:     g.Topic,
//	Partition: int32(g.Partition),
//	Value:     sarama.ByteEncoder(b),
//}
//partition, offset, err := producer.SendMessage(message)
//if err != nil {
//	c.engine.ErrChan <- err
//}
//fmt.Printf("[%s] p=%d, o=%d\n", c.Topic, partition, offset)
//atomic.AddInt64(&messageNo, 1)

//e := NewFakeEvent()
//spew.Dump(e)
//println("hello")
//for {
//    c.engine.ErrChan<-errors.New("test from classify")
//	//files, err := readDir(option.Dir, option.BatchSize)
//	//if err != nil && err != io.EOF {
//	//	log.Error(err)
//	//	time.Sleep(fetchInterval)
//	//	continue
//	//}
//	//
//	//// Handle files
//	//err = processFiles(files)
//	//if err != nil {
//	//	log.Error(err)
//	//	time.Sleep(fetchInterval)
//	//	continue
//	//}
//
//	// Sleep
//	time.Sleep(fetchInterval)
//}
//}
