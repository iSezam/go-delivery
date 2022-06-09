package novaposhta

type City struct {
	Ref                         string `json:"Ref"`
	Description                 string `json:"Description,omitempty"`
	DescriptionRu               string `json:"DescriptionRu,omitempty"`
	Delivery1                   string `json:"Delivery1,omitempty"`
	Delivery2                   string `json:"Delivery2,omitempty"`
	Delivery3                   string `json:"Delivery3,omitempty"`
	Delivery4                   string `json:"Delivery4,omitempty"`
	Delivery5                   string `json:"Delivery5,omitempty"`
	Delivery6                   string `json:"Delivery6,omitempty"`
	Delivery7                   string `json:"Delivery7,omitempty"`
	Area                        string `json:"Area,omitempty"`
	SettlementType              string `json:"SettlementType,omitempty"`
	IsBranch                    string `json:"IsBranch,omitempty"`
	Conglomerates               string `json:"Conglomerates,omitempty"`
	CityID                      string `json:"CityID,omitempty"`
	SettlementTypeDescriptionRu string `json:"SettlementTypeDescriptionRu,omitempty"`
	SettlementTypeDescription   string `json:"SettlementTypeDescription,omitempty"`
}
