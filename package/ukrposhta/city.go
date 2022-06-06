package ukrposhta

import "encoding/xml"

type CityID struct {
	CityID string `xml:"CITY_ID,omitempty"`
}

type CityEntry struct {
	XMLName         xml.Name `xml:"Entry"`
	CityID          string   `xml:"CITY_ID"`
	CityUA          string   `xml:"CITY_UA,omitempty"`
	CityEN          string   `xml:"CITY_EN,omitempty"`
	CityRU          string   `xml:"CITY_RU,omitempty"`
	CityKOATUU      string   `xml:"CITY_KOATUU,omitempty"`
	CityKATOTTG     string   `xml:"CITY_KATOTTG,omitempty"`
	CityTypeUA      string   `xml:"CITYTYPE_UA,omitempty"`
	CityTypeEN      string   `xml:"CITYTYPE_EN,omitempty"`
	CityTypeRU      string   `xml:"CITYTYPE_RU,omitempty"`
	ShortCityTypeUA string   `xml:"SHORTCITYTYPE_UA,omitempty"`
	ShortCityTypeEN string   `xml:"SHORTCITYTYPE_EN,omitempty"`
	ShortCityTypeRU string   `xml:"SHORTCITYTYPE_RU,omitempty"`
	Population      string   `xml:"POPULATION,omitempty"`
	Longitude       string   `xml:"LONGITUDE,omitempty"`
	Latitude        string   `xml:"LATTITUDE,omitempty"`
	RegionID
	DistrictID
}
