package formats

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"strings"
)

type Entry struct {
	Term string `json:"term" xml:"term"`
	Def  string `json:"definition" xml:"definition"`
}

func GetJSON(data []byte) ([]Entry, error) {
	var d []Entry
	err := json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func GetXML(data []byte) ([]Entry, error) {
	var d []Entry
	err := xml.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func GetSimple(data []byte) ([]Entry, error) {
	var line []string
	var entry Entry
	var entries []Entry

	rd := bytes.NewReader(data)
	sc := bufio.NewScanner(rd)

	for sc.Scan() {
		line = strings.Split(sc.Text(), ", ")
		entry.Term = strings.ReplaceAll(line[0], "\"", "")
		entry.Def = strings.ReplaceAll(line[1], "\"", "")
		entries = append(entries, entry)
	}

	return entries, nil
}
