package novaposhta

type T struct {
	Ref                            string `json:"Ref"`
	SiteKey                        string `json:"SiteKey"`
	Description                    string `json:"Description"`
	DescriptionRu                  string `json:"DescriptionRu"`
	ShortAddress                   string `json:"ShortAddress"`
	ShortAddressRu                 string `json:"ShortAddressRu"`
	Phone                          string `json:"Phone"`
	TypeOfWarehouse                string `json:"TypeOfWarehouse"`
	Number                         string `json:"Number"`
	CityRef                        string `json:"CityRef"`
	CityDescription                string `json:"CityDescription"`
	CityDescriptionRu              string `json:"CityDescriptionRu"`
	SettlementRef                  string `json:"SettlementRef"`
	SettlementDescription          string `json:"SettlementDescription"`
	SettlementAreaDescription      string `json:"SettlementAreaDescription"`
	SettlementRegionsDescription   string `json:"SettlementRegionsDescription"`
	SettlementTypeDescription      string `json:"SettlementTypeDescription"`
	SettlementTypeDescriptionRu    string `json:"SettlementTypeDescriptionRu"`
	Longitude                      string `json:"Longitude"`
	Latitude                       string `json:"Latitude"`
	PostFinance                    string `json:"PostFinance"`
	BicycleParking                 string `json:"BicycleParking"`
	PaymentAccess                  string `json:"PaymentAccess"`
	POSTerminal                    string `json:"POSTerminal"`
	InternationalShipping          string `json:"InternationalShipping"`
	SelfServiceWorkplacesCount     string `json:"SelfServiceWorkplacesCount"`
	TotalMaxWeightAllowed          string `json:"TotalMaxWeightAllowed"`
	PlaceMaxWeightAllowed          string `json:"PlaceMaxWeightAllowed"`
	SendingLimitationsOnDimensions struct {
		Width  int `json:"Width"`
		Height int `json:"Height"`
		Length int `json:"Length"`
	} `json:"SendingLimitationsOnDimensions"`
	ReceivingLimitationsOnDimensions struct {
		Width  int `json:"Width"`
		Height int `json:"Height"`
		Length int `json:"Length"`
	} `json:"ReceivingLimitationsOnDimensions"`
	Reception struct {
		Monday    string `json:"Monday"`
		Tuesday   string `json:"Tuesday"`
		Wednesday string `json:"Wednesday"`
		Thursday  string `json:"Thursday"`
		Friday    string `json:"Friday"`
		Saturday  string `json:"Saturday"`
		Sunday    string `json:"Sunday"`
	} `json:"Reception"`
	Delivery struct {
		Monday    string `json:"Monday"`
		Tuesday   string `json:"Tuesday"`
		Wednesday string `json:"Wednesday"`
		Thursday  string `json:"Thursday"`
		Friday    string `json:"Friday"`
		Saturday  string `json:"Saturday"`
		Sunday    string `json:"Sunday"`
	} `json:"Delivery"`
	Schedule struct {
		Monday    string `json:"Monday"`
		Tuesday   string `json:"Tuesday"`
		Wednesday string `json:"Wednesday"`
		Thursday  string `json:"Thursday"`
		Friday    string `json:"Friday"`
		Saturday  string `json:"Saturday"`
		Sunday    string `json:"Sunday"`
	} `json:"Schedule"`
	DistrictCode        string `json:"DistrictCode"`
	WarehouseStatus     string `json:"WarehouseStatus"`
	WarehouseStatusDate string `json:"WarehouseStatusDate"`
	CategoryOfWarehouse string `json:"CategoryOfWarehouse"`
	RegionCity          string `json:"RegionCity"`
	WarehouseForAgent   string `json:"WarehouseForAgent"`
	MaxDeclaredCost     string `json:"MaxDeclaredCost"`
	DenyToSelect        string `json:"DenyToSelect"`
	PostMachineType     string `json:"PostMachineType"`
	OnlyReceivingParcel string `json:"OnlyReceivingParcel"`
	WarehouseIndex      string `json:"WarehouseIndex"`
}
