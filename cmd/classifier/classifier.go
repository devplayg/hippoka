package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/devplayg/hippo"
	"github.com/devplayg/hippoka/classifier"
	"github.com/spf13/pflag"
	"os"
	"runtime"
	"strings"
)

const (
	appName        = "classifier"
	appDescription = "Data Classifier"
	appVersion     = "1.0.1"
)

var (
	option *classifier.Option

	fs        = pflag.NewFlagSet(appName, pflag.ContinueOnError)
	debug     = fs.Bool("debug", false, "Debug")
	cpu       = fs.IntP("cpu", "c", 0, "CPU Count")
	dir       = fs.StringP("dir", "d", "", "Source directory")
	batchSize = fs.UintP("batch", "b", 1, "Batch size")
	brokers   = fs.String("bootstrap-server", "", "*REQUIRED: comma-separated Kafka servers to connect to")
	topic     = fs.String("topic", "", "*REQUIRED: Kafka topic")
	partition = fs.IntP("partition", "p", -1, "Partition")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(*cpu)

	fs.Usage = hippo.Usage(fs, appDescription, appVersion)
	_ = fs.Parse(os.Args[1:])

	spew.Dump(*topic)

	if len(*topic) < 1 {
		fs.Usage()
		os.Exit(1)
	}
}

func main() {
	classifier := &classifier.Classifier{
		Name:      appName,
		Version:   appVersion,
		Debug:     *debug,
		BatchSize: int(*batchSize),
		Topic:     *topic,
		Partition: *partition,
		Brokers:   strings.Split(*brokers, ","),
	}
	engine := hippo.NewEngine(classifier)
	err := engine.Start()
	if err != nil {
		panic(err)
	}
	defer engine.Stop()
}
