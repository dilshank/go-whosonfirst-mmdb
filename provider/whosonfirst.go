package provider

import (
	"encoding/json"
	"errors"
	"github.com/oschwald/maxminddb-golang"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-spr"
	"net"
)

type WOFProvider struct {
	iplookup.Provider
	db      *maxminddb.Reader
}

func NewWOFProvider(path string) (iplookup.Provider, error) {

	db, err := maxminddb.Open(path)

	if err != nil {
		return nil, err
	}

	pr := WOFProvider{
		db:      db,
	}

	return &pr, nil
}

func (pr *WOFProvider) QueryString(str_addr string) (spr.StandardPlacesResult, error) {

	addr := net.ParseIP(str_addr)
	return pr.Query(addr)
}

func (pr *WOFProvider) Query(addr net.IP) (spr.StandardPlacesResult, error) {

	var r interface{}
	err := pr.db.Lookup(addr, &r)

	if err != nil {
		return nil, err
	}

	return pr.resultToWOFStandardPlacesResult(r)
}

func resultToWOFStandardPlacesResult(r interface{}) (spr.StandardPlacesResult, error) {

	i := r.(map[string]interface{})
	str_body, ok := i["spr"]

	if !ok {
		return nil, errors.New("Result is missing a 'spr' attribute")
	}

	body := []byte(str_body.(string))

	// see this... it's not clear to me that this is what we actually want to
	// use as an SPR or more specfically we may want a struct thingy with extra
	// properties for JSON-ify-ing - we're still trying to figure that out
	// (20170828/thisisaaronland)

	var s feature.WOFStandardPlacesResult
	err := json.Unmarshal(body, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
