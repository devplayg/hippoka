package main

import (
	"github.com/devplayg/hippo"
	"github.com/devplayg/hippoka/generator"
	"github.com/spf13/pflag"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	appName        = "generator"
	appDescription = "Data Generator"
	appVersion     = "1.0.0"
)

var (
	fs        = pflag.NewFlagSet(appName, pflag.ContinueOnError)
	debug     = fs.BoolP("debug", "d", false, "Debug")
	brokers   = fs.String("bootstrap-server", "", "*REQUIRED: comma-separated Kafka servers to connect to")
	topic     = fs.String("topic", "", "*REQUIRED: Kafka topic")
	partition = fs.Int32P("partition", "p", -1, "Partition")

	// Data
	dataSize = fs.Int64P("size", "s", 10e3, "Data size")
	count    = fs.Int64P("count", "c", 0, "Message count to send")
	interval = fs.IntP("interval", "i", 100, "Interval(ms)")
	//max      = fs.Int64P("max", "m", -1, "Max count to send")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	fs.Usage = hippo.Usage(fs, appDescription, appVersion)
	_ = fs.Parse(os.Args[1:])

	if len(*topic) < 1 || len(*brokers) < 1 {
		fs.Usage()
		os.Exit(1)
	}
}

func main() {
	generator := &generator.Generator{
		Name:    appName,
		Version: appVersion,
		Debug:   *debug,

		Topic:     *topic,
		Partition: *partition,
		Brokers:   strings.Split(*brokers, ","),

		Count:    *count,
		DataSize: *dataSize,
		Interval: *interval,
	}
	engine := hippo.NewEngine(generator)
	err := engine.Start()
	if err != nil {
		panic(err)
	}
	defer engine.Stop()
}
