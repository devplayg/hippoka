package main

import (
	"github.com/devplayg/hippo"
	"github.com/devplayg/hippoka/classifier"
	"github.com/spf13/pflag"
	"os"
	"runtime"
)

const (
	appName        = "classifier"
	appDescription = "Data Classifier"
	appVersion     = "1.0"
)

var (
	option *classifier.Option

	fs        = pflag.NewFlagSet(appName, pflag.ContinueOnError)
	debug     = fs.Bool("debug", false, "Debug")
	cpu       = fs.IntP("cpu", "c", 0, "CPU Count")
	dir       = fs.StringP("dir", "d", "", "Source directory (required)")
	batchSize = fs.UintP("batch", "b", 1, "Batch size")
	topic     = fs.StringP("topic", "t", "my-topic", "Topic")
	partition = fs.IntP("partition", "p", 1, "Partition")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(*cpu)

	fs.Usage = hippo.Usage(fs, appDescription, appVersion)
	_ = fs.Parse(os.Args[1:])

	if len(*dir) < 1 {
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
	}
	engine := hippo.NewEngine(classifier)
	err := engine.Start()
	if err != nil {
		panic(err)
	}
	defer engine.Stop()
}
