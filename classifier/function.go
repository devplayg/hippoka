package classifier

import (
	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

import (
//	"github.com/davecgh/go-spew/spew"
//	"github.com/devplayg/hippo"
//	log "github.com/sirupsen/logrus"
//	"io"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//	"regexp"
//	"time"
)

var fetchInterval = time.Duration(2000 * time.Millisecond) // option 으로 처리해야함

func ReadConfig(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func readDir(dir string, batchSize int) ([]os.FileInfo, error) {
	files := make([]os.FileInfo, 0)
	idx := 0
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Error(err)
			return nil
		}
		if err == nil && !f.IsDir() && f != nil && f.Mode().IsRegular() && f.Size() >= 0 {
			idx++
			files = append(files, f)
		}

		if idx == batchSize {
			//spew.Dump(files)
			return io.EOF
		}

		return nil
	})
	return files, err
}

//
//var imageFileNamePattern = `^(\w+)_(\d{14})_(\w+).*$` // CM000001_20180410223302_CC
//var imageFileNameRegex *regexp.Regexp

//
//func init() {
//	imageFileNameRegex = regexp.MustCompile(imageFileNamePattern)
//}
//
//func getEventFromFileName(name string) (*Event, error) {
//	match := imageFileNameRegex.FindStringSubmatch(name)
//	if len(match) == 4 {
//		t, err := time.ParseInLocation("20060102150405", match[2], loc)
//		if err != nil {
//			return nil, hippo.ErrorInvalidTimeFormat
//		}
//		return NewEvent(match[1], t, match[3]), nil
//	}
//
//	return nil, hippo.ErrorInvalidDataFormat
//}
//

//
//func processFiles(files []os.FileInfo) error {
//	for _, f := range files {
//		e, _ := getEventFromFileName(f.Name())
//		spew.Dump(e)
//	}
//	return nil
//}
//
//func ReadConfig(path string) (*Config, error) {
//	b, err := ioutil.ReadFile(path)
//	if err != nil {
//		return nil, err
//	}
//
//	config := &Config{}
//	err = yaml.Unmarshal(b, config)
//	if err != nil {
//		return nil, err
//	}
//
//	return config, nil
//}
//
