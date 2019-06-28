package collector

type Collector struct {
	//UserAgent string
	MaxDepth          int
	DisallowedDomains []string
	disallowedDomains map[string]bool
	AllowedDomains    []string
	allowedDomains    map[string]bool

	//DisallowedURLFilters []*regexp.Regexp
	//URLFilters []*regexp.Regexp
	//AllowURLRevisit bool
	MaxBodySize int
	//CacheDir string
	//IgnoreRobotsTxt bool
	//Async bool
	//ParseHTTPErrorResponse bool
	//ID uint32
	//DetectCharset bool
	//RedirectHandler func(req *http.Request, via []*http.Request) error
	//CheckHead         bool
	//store             storage.Storage
	//debugger          debug.Debugger
	//robotsMap         map[sstring]*robotstxt.RobotsData
	//htmlCallbacks     []*htmlCallbackContainer
	//xmlCallbacks      []*xmlCallbackContainer
	//requestCallbacks  []RequestCallback
	//responseCallbacks []ResponseCallback
	//errorCallbacks    []ErrorCallback
	//scrapedCallbacks  []ScrapedCallback
	//requestCount      uint32
	//responseCount     uint32
	//backend           *httpBackend
	//wg                *sync.WaitGroup
	//lock              *sync.RWMutex
}

func (c *Collector) Start(url string) error {
	c.visit(url)
}

func (c *Collector) visit(url string) {

}
