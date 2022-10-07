package main

import (
	"flag"
	"log"
	"os"

	"github.com/lukas-blaha/quickle/pkg/db"
)

func main() {
	lesson := flag.String("name", "", "name of the study set you want to import")
	f := flag.String("file", "", "path to the file you want to import")
	fileFormat := flag.String("format", "", "file format <simple|json|xml>")
	flag.Parse()

	if *lesson == "" || *f == "" || *fileFormat == "" {
		flag.Usage()
		os.Exit(1)
	}

	data := Data{
		Path:   *f,
		Format: *fileFormat,
	}

	// connect to DB
	conn := db.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	if err := db.NewTable(conn, *lesson); err != nil {
		log.Fatal(err)
	}

	d, err := data.LoadData()
	if err != nil {
		log.Fatal(err)
	}

	// parse data
	data.ParseData(d)

	// import to db
	db.ImportData(conn, *lesson, data.Entries)
}
