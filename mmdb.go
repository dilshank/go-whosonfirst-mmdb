package mmdb

import (
	"encoding/json"
	"errors"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-spr"
)

func ResultToWOFStandardPlacesResult(r interface{}) (spr.StandardPlacesResult, error) {

	i := r.(map[string]interface{})
	str_body, ok := i["spr"]

	if !ok {
		return nil, errors.New("Result is missing a 'spr' attribute")
	}

	body := []byte(str_body.(string))

	var s feature.WOFStandardPlacesResult
	err := json.Unmarshal(body, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
