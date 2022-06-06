package ukrposhta

import "encoding/xml"

type WarehouseEntry struct {
	XMLName   xml.Name `xml:"Entry"`
	ID        string   `xml:"ID"`
	Postcode  string   `xml:"POSTINDEX,omitempty"`
	Name      string   `xml:"PO_LONG,omitempty"`
	ShortName string   `xml:"PO_SHORT,omitempty"`
	Address   string   `xml:"ADDRESS,omitempty"`
	Phone     string   `xml:"PHONE,omitempty"`
	DistrictID
	RegionID
	CityID
}
