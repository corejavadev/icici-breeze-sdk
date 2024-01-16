package breeze_models

import "time"

type NseFnoScript struct {
	Token                    int       `csv:"Token"`
	InstrumentName           string    `csv:"InstrumentName"`
	ShortName                string    `csv:"ShortName"`
	Series                   string    `csv:"Series"`
	ExpiryDate               time.Time `csv:"ExpiryDate"`
	StrikePrice              int       `csv:"StrikePrice"`
	OptionType               string    `csv:"OptionType"`
	CALevel                  int       `csv:"CALevel"`
	PermittedToTrade         int       `csv:"PermittedToTrade"`
	IssueCapital             float64   `csv:"IssueCapital"`
	WarningQty               int       `csv:"WarningQty"`
	FreezeQty                int       `csv:"FreezeQty"`
	CreditRating             string    `csv:"CreditRating"`
	NormalMarketStatus       string    `csv:"NormalMarketStatus"`
	OddLotMarketStatus       string    `csv:"OddLotMarketStatus"`
	SpotMarketStatus         string    `csv:"SpotMarketStatus"`
	AuctionMarketStatus      string    `csv:"AuctionMarketStatus"`
	NormalMarketEligibility  string    `csv:"NormalMarketEligibility"`
	OddLotMarketEligibility  string    `csv:"OddLotMarketEligibility"`
	SpotMarketEligibility    string    `csv:"SpotMarketEligibility"`
	AuctionMarketEligibility string    `csv:"AuctionMarketEligibility"`
	IssueRate                float64   `csv:"IssueRate"`
	IssueStartDate           time.Time `csv:"IssueStartDate"`
	InterestPaymentDate      time.Time `csv:"InterestPaymentDate"`
	IssueMaturityDate        time.Time `csv:"IssueMaturityDate"`
	MarginPercentage         float64   `csv:"MarginPercentage"`
	MinimumLotQty            int       `csv:"MinimumLotQty"`
	LotSize                  int       `csv:"LotSize"`
	TickSize                 float64   `csv:"TickSize"`
	CompanyName              string    `csv:"CompanyName"`
	ListingDate              time.Time `csv:"ListingDate"`
	ExpulsionDate            time.Time `csv:"ExpulsionDate"`
	ReAdmissionDate          time.Time `csv:"ReAdmissionDate"`
	RecordDate               time.Time `csv:"RecordDate"`
	LowPriceRange            float64   `csv:"LowPriceRange"`
	HighPriceRange           float64   `csv:"HighPriceRange"`
	SecurityExpiryDate       time.Time `csv:"SecurityExpiryDate"`
	NoDeliveryStartDate      time.Time `csv:"NoDeliveryStartDate"`
	NoDeliveryEndDate        time.Time `csv:"NoDeliveryEndDate"`
	MF                       int       `csv:"MF"`
	AON                      int       `csv:"AON"`
	ParticipantInMarketIndex string    `csv:"ParticipantInMarketIndex"`
	BookClsStartDate         time.Time `csv:"BookClsStartDate"`
	BookClsEndDate           time.Time `csv:"BookClsEndDate"`
	ExcerciseStartDate       time.Time `csv:"ExcerciseStartDate"`
	ExcerciseEndDate         time.Time `csv:"ExcerciseEndDate"`
	OldToken                 int       `csv:"OldToken"`
	AssetInstrument          string    `csv:"AssetInstrument"`
	AssetName                string    `csv:"AssetName"`
	AssetToken               int       `csv:"AssetToken"`
	IntrinsicValue           float64   `csv:"IntrinsicValue"`
	ExtrinsicValue           float64   `csv:"ExtrinsicValue"`
	ExcerciseStyle           string    `csv:"ExcerciseStyle"`
	EGM                      int       `csv:"EGM"`
	AGM                      int       `csv:"AGM"`
	Interest                 float64   `csv:"Interest"`
	Bonus                    int       `csv:"Bonus"`
	Rights                   int       `csv:"Rights"`
	Dividends                float64   `csv:"Dividends"`
	ExAllowed                int       `csv:"ExAllowed"`
	ExRejectionAllowed       int       `csv:"ExRejectionAllowed"`
	PlAllowed                int       `csv:"PlAllowed"`
	IsThisAsset              int       `csv:"IsThisAsset"`
	IsCorpAdjusted           int       `csv:"IsCorpAdjusted"`
	LocalUpdateDatetime      time.Time `csv:"LocalUpdateDatetime"`
	DeleteFlag               int       `csv:"DeleteFlag"`
	Remarks                  string    `csv:"Remarks"`
	BasePrice                float64   `csv:"BasePrice"`
	ExchangeCode             string    `csv:"ExchangeCode"`
}