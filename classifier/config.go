package classifier

// Yaml config
type Config struct {
	Server struct {
		Timezone string `json:"timezone"`
	}
	Kafka struct {
		Broker []string `json:"broker"`
	}
}
