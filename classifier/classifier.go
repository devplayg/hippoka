package classifier

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/devplayg/hippo"
	"log"
	"sync/atomic"
	"time"
)

var loc = time.UTC

type Classifier struct {
	engine    *hippo.Engine
	DB        *sql.DB
	config    *Config
	Name      string
	Version   string
	Debug     bool
	BatchSize int
	Topic     string
	Partition int
}

func (c *Classifier) Start() error {
	err := c.loadConfig()
	if err != nil {
		return err
	}

	//err = c.initDatabase()
	//if err != nil {
	//    return err
	//}

	err = c.setTimezone()
	if err != nil {
		return err
	}

	c.classify()

	//hippo.WaitForSignals()

	return nil
}

func (c *Classifier) Stop() error {
	//println("classifier is stopping")
	//err := c.DB.Close()
	//if err != nil {
	//    return err
	//}
	return nil
}
func (c *Classifier) SetEngine(e *hippo.Engine) {
	c.engine = e
}

func (c *Classifier) GetName() string {
	return c.Name
}

func (c *Classifier) GetVersion() string {
	return c.Version
}

func (c *Classifier) IsDebug() bool {
	return c.Debug
}

func (c *Classifier) classify() {
	config := sarama.NewConfig()
	if c.Partition >= 0 {
		config.Producer.Partitioner = sarama.NewManualPartitioner
	} else {
		config.Producer.Partitioner = sarama.NewHashPartitioner
	}
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(c.config.Kafka.Broker, config)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err := producer.Close(); err != nil {
			c.engine.ErrChan <- err
		}
	}()

	var messageNo int64 = 1
	e := NewFakeEvent(messageNo)
	b, _ := json.Marshal(e)

	message := &sarama.ProducerMessage{
		Topic:     c.Topic,
		Partition: int32(c.Partition),
		Value:     sarama.ByteEncoder(b),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		c.engine.ErrChan <- err
	}
	fmt.Printf("[%s] p=%d, o=%d\n", c.Topic, partition, offset)
	atomic.AddInt64(&messageNo, 1)

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
}

func (c *Classifier) loadConfig() error {
	config, err := ReadConfig("config.yaml")
	if err != nil {
		return err
	}
	c.config = config
	return nil
}

func (c *Classifier) setTimezone() error {
	if len(c.config.Server.Timezone) > 0 {
		customLoc, err := time.LoadLocation(c.config.Server.Timezone)
		if err != nil {
			return err
		}
		loc = customLoc
	}

	return nil
}

//
//func (c *Classifier) initDatabase() error {
//    connStr := fmt.Sprintf(
//        "%s:%s@tcp(%s:%s)/%s?allowAllFiles=true&charset=utf8&parseTime=true&loc=%s",
//        c.config.Datasource.Username,
//        c.config.Datasource.Password,
//        c.config.Datasource.Hostname,
//        c.config.Datasource.Port,
//        c.config.Datasource.Database,
//        strings.Replace(c.config.Server.Timezone, "/", "%2F", -1),
//    )
//
//    db, _ := sql.Open("mysql", connStr)
//    err := db.Ping()
//    if err != nil {
//        return err
//    }
//
//    db.SetMaxIdleConns(3)
//    db.SetMaxOpenConns(3)
//    c.DB = db
//    return nil
//}

//
//func (c *Classifier) insertData(filename string) error {
//	query := `
//		LOAD DATA LOCAL INFILE %q
//		INTO TABLE log_xxxx
//		FIELDS TERMINATED BY '\t'
//		LINES TERMINATED BY '\n'
//	`
//	query = fmt.Sprintf(query, filepath.ToSlash(filename))
//	rs, err := c.engine.DB.Exec(query)
//	if err != nil {
//		return err
//	}
//	rowsAffected, _ := rs.RowsAffected()
//	log.Debugf("table=%s, affected_rows=%d", "data", rowsAffected)
//	return nil
//}
