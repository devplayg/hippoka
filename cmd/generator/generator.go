package main

import (
	"github.com/devplayg/hippo"
	"github.com/devplayg/hippoka/generator"
	"github.com/spf13/pflag"
	"os"
	"runtime"
	"strings"
)

const (
	appName        = "generator"
	appDescription = "Data Generator"
	appVersion     = "1.0.0"
)

var (
	fs        = pflag.NewFlagSet(appName, pflag.ContinueOnError)
	debug     = fs.Bool("debug", false, "Debug")
	brokers   = fs.String("bootstrap-server", "", "*REQUIRED: comma-separated Kafka servers to connect to")
	topic     = fs.String("topic", "", "*REQUIRED: Kafka topic")
	partition = fs.IntP("partition", "p", -1, "Partition")

	// Data
	size     = fs.IntP("size", "s", 1, "Message size")
	interval = fs.IntP("interval", "i", 100, "Interval(ms)")
	max      = fs.Int64P("max", "m", 10000, "Message size")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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

		Size:     *size,
		Max:      *max,
		Interval: *interval,
	}
	engine := hippo.NewEngine(generator)
	err := engine.Start()
	if err != nil {
		panic(err)
	}
	defer engine.Stop()
}
