package main

import (
	"encoding/json"
	"flag"
	"github.com/whosonfirst/go-whosonfirst-csv"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-log"
	_ "github.com/whosonfirst/go-whosonfirst-mmdb"
	"github.com/whosonfirst/go-whosonfirst-spr"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	var concordances = flag.String("concordances", "", "")
	var data_root = flag.String("data-root", "/usr/local/data", "")
	var repo = flag.String("repo", "whosonfirst-data", "")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	fh, err := os.Open(*concordances)

	if err != nil {
		logger.Fatal("failed to open %s, because %s", *concordances, err)
	}

	lookup := make(map[int64]spr.StandardPlacesResult)

	root := filepath.Join(*data_root, *repo)
	data := filepath.Join(root, "data")

	reader, err := csv.NewDictReader(fh)

	if err != nil {
		logger.Fatal("failed to open csv reader because %s", err)
	}

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			logger.Fatal("failed to process %s because %s", *concordances, err)
		}

		str_gnid, ok := row["gn:id"]

		if !ok {
			logger.Fatal("missing gn:id key")
		}

		gnid, err := strconv.ParseInt(str_gnid, 10, 64)

		if err != nil {
			logger.Fatal("failed to parse %s because %s", str_gnid, err)
		}

		str_wofid, ok := row["wof:id"]

		if !ok {
			logger.Fatal("missing wof:id key")
		}

		if str_wofid == "-1" {
			continue
		}

		wofid, err := strconv.ParseInt(str_wofid, 10, 64)

		if err != nil {
			logger.Fatal("failed to parse %s because %s", str_wofid, err)
		}

		abs_path, err := uri.Id2AbsPath(data, wofid)

		if err != nil {
			logger.Fatal("failed to determine absolute path for %d because %s", wofid, err)
		}

		f, err := feature.LoadWOFFeatureFromFile(abs_path)

		if err != nil {
			logger.Fatal("failed to load %s because %s", abs_path, err)
		}

		s, err := f.SPR()

		if err != nil {
			logger.Fatal("failed to create SPR for %s because %s", abs_path, err)
		}

		lookup[gnid] = s
	}

	enc, err := json.Marshal(lookup)

	if err != nil {
		logger.Fatal("failed to marshal lookup because %s", err)
	}

	writer := os.Stdout
	writer.Write(enc)
}
