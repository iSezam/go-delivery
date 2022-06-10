package novaposhta

type Settlement struct {
	Ref                         string   `json:"Ref"`
	SettlementType              string   `json:"SettlementType,omitempty"`
	Latitude                    string   `json:"Latitude,omitempty"`
	Longitude                   string   `json:"Longitude,omitempty"`
	Description                 string   `json:"Description,omitempty"`
	DescriptionRu               string   `json:"DescriptionRu,omitempty"`
	SettlementTypeDescription   string   `json:"SettlementTypeDescription,omitempty"`
	SettlementTypeDescriptionRu string   `json:"SettlementTypeDescriptionRu,omitempty"`
	Region                      string   `json:"Region,omitempty"`
	RegionsDescription          string   `json:"RegionsDescription,omitempty"`
	RegionsDescriptionRu        string   `json:"RegionsDescriptionRu,omitempty"`
	Area                        string   `json:"Area,omitempty"`
	AreaDescription             string   `json:"AreaDescription,omitempty"`
	AreaDescriptionRu           string   `json:"AreaDescriptionRu,omitempty"`
	Index1                      string   `json:"Index1,omitempty"`
	Index2                      string   `json:"Index2,omitempty"`
	IndexCOATSU1                string   `json:"IndexCOATSU1,omitempty"`
	Delivery1                   string   `json:"Delivery1,omitempty"`
	Delivery2                   string   `json:"Delivery2,omitempty"`
	Delivery3                   string   `json:"Delivery3,omitempty"`
	Delivery4                   string   `json:"Delivery4,omitempty"`
	Delivery5                   string   `json:"Delivery5,omitempty"`
	Delivery6                   string   `json:"Delivery6,omitempty"`
	Delivery7                   string   `json:"Delivery7,omitempty"`
	Warehouse                   string   `json:"Warehouse,omitempty"`
	Conglomerates               []string `json:"Conglomerates,omitempty"`
}
