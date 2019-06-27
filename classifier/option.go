package classifier

//
//type Option struct {
//	Name        string
//	Description string
//	Version     string
//	Debug       bool
//	Dir         string
//	Storage     string
//	BatchSize   int
//	Topic       string
//	Partition   int
//}

//func NewOption(name, description, version string, debug bool) *Option {
//	return &Option{
//		Name:        name,
//		description: description,
//		version:     version,
//		debug:       debug,
//	}
//}
//
//func (c *Option) Name() string {
//	return c.name
//}
//
//func (c *Option) Version() string {
//	return c.version
//}
//
//func (c *Option) Debug() bool {
//	return c.debug
//}

func (c *Option) Validate() error {
	// Source directory
	//if len(c.Dir) < 1 {
	//	return hippo.ErrorRequiredOption
	//}
	//
	//abs, err := filepath.Abs(c.Dir)
	//if err != nil {
	//	return hippo.ErrorDirectoryNotFound
	//}
	//c.Dir = abs
	//
	//if _, err := os.Stat(c.Dir); os.IsNotExist(err) {
	//	return hippo.ErrorDirectoryNotFound
	//}
	//
	//// Storage directory
	//if len(c.Storage) < 1 {
	//	return hippo.ErrorRequiredOption
	//}
	//
	//abs, err = filepath.Abs(c.Storage)
	//if err != nil {
	//	return hippo.ErrorDirectoryNotFound
	//}
	//c.Storage = abs
	//
	//if _, err := os.Stat(c.Storage); os.IsNotExist(err) {
	//	return hippo.ErrorDirectoryNotFound
	//}

	return nil
}
