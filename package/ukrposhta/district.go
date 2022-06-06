package ukrposhta

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Districts(regionID, name string) (districts []DistrictEntity, statusCode int, err error) {
	var entities struct {
		XMLName xml.Name         `xml:"Entries"`
		Entries []DistrictEntity `xml:"Entry"`
	}

	URL := fmt.Sprintf("https://ukrposhta.ua/address-classifier/get_districts_by_region_id_and_district_ua?region_id=%s&district_ua=%s", regionID, url.QueryEscape(name))

	resp, err := http.Get(URL)
	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode == http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(b, &entities)
		resp.Body.Close()

		districts = entities.Entries
	}
	return districts, resp.StatusCode, nil
}

type DistrictID struct {
	DistrictID string `xml:"DISTRICT_ID,omitempty"`
}

type DistrictEntity struct {
	XMLName         xml.Name `xml:"Entry"`
	DistrictID      string   `xml:"DISTRICT_ID"`
	DistrictUA      string   `xml:"DISTRICT_UA,omitempty"`
	DistrictEN      string   `xml:"DISTRICT_EN,omitempty"`
	DistrictRU      string   `xml:"DISTRICT_RU,omitempty"`
	DistrictKOATUU  string   `xml:"DISTRICT_KOATUU,omitempty"`
	DistrictKATOTTG string   `xml:"DISTRICT_KATOTTG,omitempty"`
	NewDistrictUA   string   `xml:"NEW_DISTRICT_UA,omitempty"`
	RegionID
}
