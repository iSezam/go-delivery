package ukrposhta

import "encoding/xml"

type StreetID struct {
	StreetID string `xml:"STREET_ID,omitempty"`
}

type StreetEntry struct {
	XMLName      xml.Name `xml:"Entry"`
	StreetID     string   `xml:"STREET_ID"`
	StreetUA     string   `xml:"STREET_UA,omitempty"`
	StreetEN     string   `xml:"STREET_EN,omitempty"`
	StreetRU     string   `xml:"STREET_RU,omitempty"`
	StreetTypeUA string   `xml:"STREETTYPE_UA,omitempty"`
	StreetTypeEN string   `xml:"STREETTYPE_EN,omitempty"`
	StreetTypeRU string   `xml:"STREETTYPE_RU,omitempty"`
	RegionID
	DistrictID
	CityID
}
