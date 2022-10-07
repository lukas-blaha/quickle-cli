package main

import (
	"io"
	"log"
	"os"

	"github.com/lukas-blaha/quickle/pkg/formats"
)

type Data struct {
	Path    string
	Format  string
	Entries []formats.Entry
}

func (d Data) LoadData() ([]byte, error) {
	f, err := os.Open(d.Path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *Data) ParseData(data []byte) {
	var err error
	switch {
	case d.Format == "simple":
		if d.Entries, err = formats.GetSimple(data); err != nil {
			log.Fatal(err)
		}
	case d.Format == "json":
		if d.Entries, err = formats.GetJSON(data); err != nil {
			log.Fatal(err)
		}
	case d.Format == "xml":
		if d.Entries, err = formats.GetXML(data); err != nil {
			log.Fatal(err)
		}
	}
}
