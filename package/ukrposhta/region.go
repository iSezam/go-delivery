package ukrposhta

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Regions(name, nameEn string) (regions []RegionEntity, statusCode int, err error) {
	var entities struct {
		XMLName xml.Name       `xml:"Entries"`
		Entries []RegionEntity `xml:"Entry"`
	}
	var URL string
	if nameEn == "" {
		URL = fmt.Sprintf("https://ukrposhta.ua/address-classifier/get_regions_by_region_ua?region_name=%s", url.QueryEscape(name))
	} else {
		URL = fmt.Sprintf("https://ukrposhta.ua/address-classifier/get_regions_by_region_ua?region_name=%s&region_name_en=%s", url.QueryEscape(name), url.QueryEscape(nameEn))
	}

	resp, err := http.Get(URL)
	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode == http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(b, &entities)
		resp.Body.Close()

		regions = entities.Entries
	}
	return regions, resp.StatusCode, nil
}

type RegionID struct {
	RegionID string `xml:"REGION_ID,omitempty"`
}

type RegionEntity struct {
	XMLName       xml.Name `xml:"Entry"`
	RegionID      string   `xml:"REGION_ID"`
	RegionUA      string   `xml:"REGION_UA,omitempty"`
	RegionEN      string   `xml:"REGION_EN,omitempty"`
	RegionRU      string   `xml:"REGION_RU,omitempty"`
	RegionKATOTTG string   `xml:"REGION_KATOTTG,omitempty"`
	RegionKOATUU  string   `xml:"REGION_KOATUU,omitempty"`
}
