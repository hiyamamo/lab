package gitlab

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func newConfigService() *configService {
	return &configService{}
}

type configService struct {
}

func (s *configService) Save(filename string, c *Config) error {
	err := os.MkdirAll(filepath.Dir(filename), 0771)
	if err != nil {
		return err
	}

	w, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer w.Close()

	return encode(w, c)
}

func encode(w io.Writer, c *Config) error {
	d, err := yaml.Marshal(c)

	if err != nil {
		return err
	}

	n, err := w.Write(d)
	if err == nil && n < len(d) {
		err = io.ErrShortWrite
	}

	return err
}

func (s *configService) Load(filename string, c *Config) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	return decode(r, c)
}

func decode(r io.Reader, c *Config) error {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(d, c)

	if err != nil {
		return err
	}

	return nil
}
