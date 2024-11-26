package bybit

type BybitResponse struct {
	RetCode int            `json:"ret_code"`
	RetMsg  string         `json:"ret_msg"`
	Result  BybitResult    `json:"result"`
	ExtCode string         `json:"ext_code"`
	ExtInfo map[string]any `json:"ext_info"`
	TimeNow string         `json:"time_now"`
}

type BybitResult struct {
	Count int         `json:"count"`
	Items []BybitItem `json:"items"`
}

type BybitItem struct {
	ID                 string             `json:"id"`
	AccountID          string             `json:"accountId"`
	UserID             string             `json:"userId"`
	NickName           string             `json:"nickName"`
	TokenID            string             `json:"tokenId"`
	TokenName          string             `json:"tokenName"`
	CurrencyID         string             `json:"currencyId"`
	Side               int                `json:"side"`
	PriceType          int                `json:"priceType"`
	Price              string             `json:"price"`
	Premium            string             `json:"premium"`
	LastQuantity       string             `json:"lastQuantity"`
	Quantity           string             `json:"quantity"`
	FrozenQuantity     string             `json:"frozenQuantity"`
	ExecutedQuantity   string             `json:"executedQuantity"`
	MinAmount          string             `json:"minAmount"`
	MaxAmount          string             `json:"maxAmount"`
	Remark             string             `json:"remark"`
	Status             int                `json:"status"`
	CreateDate         string             `json:"createDate"`
	Payments           []string           `json:"payments"`
	OrderNum           int                `json:"orderNum"`
	FinishNum          int                `json:"finishNum"`
	RecentOrderNum     int                `json:"recentOrderNum"`
	RecentExecuteRate  int                `json:"recentExecuteRate"`
	Fee                string             `json:"fee"`
	IsOnline           bool               `json:"isOnline"`
	LastLogoutTime     string             `json:"lastLogoutTime"`
	Blocked            string             `json:"blocked"`
	MakerContact       bool               `json:"makerContact"`
	SymbolInfo         SymbolInfo         `json:"symbolInfo"`
	TradingPreferences TradingPreferences `json:"tradingPreferenceSet"`
	Version            int                `json:"version"`
	AuthStatus         int                `json:"authStatus"`
	Recommend          bool               `json:"recommend"`
	RecommendTag       string             `json:"recommendTag"`
	AuthTag            []string           `json:"authTag"`
	UserType           string             `json:"userType"`
	ItemType           string             `json:"itemType"`
	PaymentPeriod      int                `json:"paymentPeriod"`
}

type SymbolInfo struct {
	ID                    string       `json:"id"`
	ExchangeID            string       `json:"exchangeId"`
	OrgID                 string       `json:"orgId"`
	TokenID               string       `json:"tokenId"`
	CurrencyID            string       `json:"currencyId"`
	Status                int          `json:"status"`
	LowerLimitAlarm       float64      `json:"lowerLimitAlarm"`
	UpperLimitAlarm       float64      `json:"upperLimitAlarm"`
	ItemDownRange         string       `json:"itemDownRange"`
	ItemUpRange           string       `json:"itemUpRange"`
	CurrencyMinQuote      string       `json:"currencyMinQuote"`
	CurrencyMaxQuote      string       `json:"currencyMaxQuote"`
	CurrencyLowerMaxQuote string       `json:"currencyLowerMaxQuote"`
	TokenMinQuote         string       `json:"tokenMinQuote"`
	TokenMaxQuote         string       `json:"tokenMaxQuote"`
	KycCurrencyLimit      string       `json:"kycCurrencyLimit"`
	ItemSideLimit         int          `json:"itemSideLimit"`
	BuyFeeRate            string       `json:"buyFeeRate"`
	SellFeeRate           string       `json:"sellFeeRate"`
	OrderAutoCancelMinute int          `json:"orderAutoCancelMinute"`
	OrderFinishMinute     int          `json:"orderFinishMinute"`
	TradeSide             int          `json:"tradeSide"`
	Currency              CurrencyInfo `json:"currency"`
	Token                 TokenInfo    `json:"token"`
	BuyAd                 *any         `json:"buyAd"`
	SellAd                *any         `json:"sellAd"`
}

type CurrencyInfo struct {
	ID         string `json:"id"`
	ExchangeID string `json:"exchangeId"`
	OrgID      string `json:"orgId"`
	CurrencyID string `json:"currencyId"`
	Scale      int    `json:"scale"`
}

type TokenInfo struct {
	ID         string `json:"id"`
	ExchangeID string `json:"exchangeId"`
	OrgID      string `json:"orgId"`
	TokenID    string `json:"tokenId"`
	Scale      int    `json:"scale"`
	Sequence   int    `json:"sequence"`
}

type TradingPreferences struct {
	HasUnPostAd               int    `json:"hasUnPostAd"`
	IsKyc                     int    `json:"isKyc"`
	IsEmail                   int    `json:"isEmail"`
	IsMobile                  int    `json:"isMobile"`
	HasRegisterTime           int    `json:"hasRegisterTime"`
	RegisterTimeThreshold     int    `json:"registerTimeThreshold"`
	OrderFinishNumberDay30    int    `json:"orderFinishNumberDay30"`
	CompleteRateDay30         string `json:"completeRateDay30"`
	NationalLimit             string `json:"nationalLimit"`
	HasOrderFinishNumberDay30 int    `json:"hasOrderFinishNumberDay30"`
	HasCompleteRateDay30      int    `json:"hasCompleteRateDay30"`
	HasNationalLimit          int    `json:"hasNationalLimit"`
}
