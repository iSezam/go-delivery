package model

import "encoding/xml"

type HouseEntry struct {
	XMLName       xml.Name `xml:"Entry"`
	StreetID      string   `xml:"STREET_ID"`
	Postcode      string   `xml:"POSTCODE"`
	HouseNumberUA string   `xml:"HOUSENUMBER_UA"`
}
