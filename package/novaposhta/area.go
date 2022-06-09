package novaposhta

type Area struct {
	Ref           string `json:"Ref"`
	AreasCenter   string `json:"AreasCenter,omitempty"`
	DescriptionRu string `json:"DescriptionRu,omitempty"`
	Description   string `json:"Description,omitempty"`
}
