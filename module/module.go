package module

type Address struct {
	UniqueId        string `json:"uniqueId"`
	Address         string `json:"address"`
	CorpAddressCode string `json:"corpAddressCode"`
	PickUpLocation  string `json:"pickUpLocation"`
}

type Corp struct {
	UniqueId                  string    `json:"uniqueId"`
	UseCloset                 bool      `json:"useCloset"`
	Name                      string    `json:"name"`
	Namespace                 string    `json:"namespace"`
	PriceVisible              bool      `json:"priceVisible"`
	ShowPrice                 bool      `json:"showPrice"`
	PriceLimit                float64   `json:"priceLimit"`
	PriceLimitInCent          uint64    `json:"priceLimitInCent"`
	AcceptCashPaymentToMeican bool      `json:"acceptCashPaymentToMeican"`
	AlwaysOpen                bool      `json:"alwaysOpen"`
	AddressList               []Address `json:"addressList"`
}

type UserTab struct {
	Corp         Corp   `json:"corp"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Name         string `json:"name"`
	LastUsedTime uint64 `json:"lastUsedTime"`
	UniqueId     string `json:"uniqueId"`
}

type OpeningTime struct {
	UniqueId         string `json:"uniqueId"`
	Name             string `json:"name"`
	OpenTime         string `json:"openTime"`
	CloseTime        string `json:"closeTime"`
	DefaultAlarmTime string `json:"defaultAlarmTime"`
	PostboxOpenTime  string `json:"postboxOpenTime"`
}

type Dish struct {
	Name                 string `json:"name"`
	PriceInCent          uint64 `json:"priceInCent"`
	OriginalPriceInCent  uint64 `json:"originalPriceInCent"`
	IsSection            bool   `json:"isSection"`
	ActionRequiredLevel  string `json:"actionRequiredLevel"`
	ActionRequiredReason string `json:"actionRequiredReason"`
	Id                   uint64 `json:"id"`
}

type DishItem struct {
	Dish  Dish   `json:"dish"`
	Count uint32 `json:"count"`
}

type RestaurantItem struct {
	UniqueId     string     `json:"uniqueId"`
	DishItemList []DishItem `json:"dishItemList"`
}

type CorpOrderUser struct {
	IsLegacyPay                   bool             `json:"isLegacyPay"`
	PayStatus                     string           `json:"payStatus"`
	RestaurantItemList            []RestaurantItem `json:"restaurantItemList"`
	Corp                          Corp             `json:"corp"`
	ReadyToDelete                 bool             `json:"readyToDelete"`
	ActionRequiredLevel           string           `json:"actionRequiredLevel"`
	CorpOrderStatus               string           `json:"corpOrderStatus"`
	ShowPrice                     bool             `json:"showPrice"`
	UnpaidUserToMeicanPrice       string           `json:"unpaidUserToMeicanPrice"`
	UnpaidUserToMeicanPriceInCent uint64           `json:"unpaidUserToMeicanPriceInCent"`
	PaidUserToMeicanPrice         string           `json:"paidUserToMeicanPrice"`
	PaidUserToMeicanPriceInCent   uint64           `json:"paidUserToMeicanPriceInCent"`
	Timestamp                     uint64           `json:"timestamp"`
	UniqueId                      string           `json:"uniqueId"`
}

type CalendarItem struct {
	TargetTime    uint64        `json:"targetTime"`
	Title         string        `json:"title"`
	UserTab       UserTab       `json:"userTab"`
	OpeningTime   OpeningTime   `json:"openingTime"`
	CorpOrderUser CorpOrderUser `json:"corpOrderUser"`
	Status        string        `json:"status"`
	Reason        string        `json:"reason"`
}

type DateListItem struct {
	Date             string         `json:"date"`
	CalendarItemList []CalendarItem `json:"calendarItemList"`
}

type OrderRsp struct {
	StartDate string         `json:"startDate"`
	EndDate   string         `json:"endDate"`
	DateList  []DateListItem `json:"dateList"`
}

type Restaurant struct {
	UniqueId           string  `json:"uniqueId"`
	Name               string  `json:"name"`
	Tel                string  `json:"tel"`
	Rating             uint64  `json:"rating"`
	MinimumOrder       uint64  `json:"minimumOrder"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Warning            string  `json:"warning"`
	OpeningTime        string  `json:"openingTime"`
	OnlinePayment      bool    `json:"onlinePayment"`
	Open               bool    `json:"open"`
	AvailableDishCount uint64  `json:"availableDishCount"`
	DishLimit          uint64  `json:"dishLimit"`
	RestaurantStatus   uint32  `json:"restaurantStatus"`
	RemarkEnabled      bool    `json:"remarkEnabled"`
}

type RestaurantRsp struct {
	NoMore         bool         `json:"noMore"`
	TargetTime     string       `json:"targetTime"`
	RestaurantList []Restaurant `json:"restaurantList"`
}

type DishRsp struct {
	AdditionalInfo                AdditionalInfo `json:"additionalInfo"`
	Assessment                    Assessment     `json:"assessment"`
	AvailableDishCount            uint64         `json:"availableDishCount"`
	BusinessLicenseURL            string         `json:"businessLicenseUrl"`
	CorpRestaurantID              uint64         `json:"corpRestaurantId"`
	DeliveryFeeInCent             interface{}    `json:"deliveryFeeInCent"`
	DeliveryRange                 interface{}    `json:"deliveryRange"`
	DeliveryRangeMeter            interface{}    `json:"deliveryRangeMeter"`
	DishLimit                     uint64         `json:"dishLimit"`
	DishList                      []DishList     `json:"dishList"`
	Latitude                      float64        `json:"latitude"`
	Longitude                     float64        `json:"longitude"`
	MinimumOrder                  interface{}    `json:"minimumOrder"`
	MyRegularDishIDList           []interface{}  `json:"myRegularDishIdList"`
	Name                          string         `json:"name"`
	OnlinePayment                 bool           `json:"onlinePayment"`
	Open                          bool           `json:"open"`
	OpeningTime                   string         `json:"openingTime"`
	OthersRegularDishIDList       []interface{}  `json:"othersRegularDishIdList"`
	OthersRegularDishIDListSource string         `json:"othersRegularDishIdListSource"`
	Rating                        uint64         `json:"rating"`
	RemarkEnabled                 bool           `json:"remarkEnabled"`
	RestaurantID                  uint64         `json:"restaurantId"`
	RestaurantStatus              uint64         `json:"restaurantStatus"`
	SanitationCertificateURL      string         `json:"sanitationCertificateUrl"`
	SectionList                   []SectionList  `json:"sectionList"`
	ShowPrice                     bool           `json:"showPrice"`
	TargetTime                    string         `json:"targetTime"`
	Tel                           string         `json:"tel"`
	UniqueID                      string         `json:"uniqueId"`
	Warning                       string         `json:"warning"`
}

type AdditionalInfo struct {
	Address        string `json:"address"`
	AssessDate     string `json:"assessDate"`
	AssessEndDate  string `json:"assessEndDate"`
	BusinessType   string `json:"businessType"`
	CityName       string `json:"cityName"`
	CityURL        string `json:"cityUrl"`
	CompanyName    string `json:"companyName"`
	District       string `json:"district"`
	Level          string `json:"level"`
	LicenseNumber  string `json:"licenseNumber"`
	Representative string `json:"representative"`
}

type Fields struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Assessment struct {
	AssessmentIconURL string   `json:"assessmentIconUrl"`
	Fields            []Fields `json:"fields"`
}

type DishList struct {
	DishSectionID       uint64 `json:"dishSectionId"`
	ID                  uint64 `json:"id"`
	IsSection           bool   `json:"isSection"`
	Name                string `json:"name"`
	OriginalPriceInCent uint64 `json:"originalPriceInCent"`
	PriceInCent         uint64 `json:"priceInCent"`
	PriceString         string `json:"priceString"`
}

type SectionList struct {
	ID         uint64   `json:"id"`
	DishIDList []uint64 `json:"dishIdList"`
	Name       string   `json:"name"`
}

type AddOrderRsp struct {
	Message string `json:"message"`
	Order   Order  `json:"order"`
	Status  string `json:"status"`
}
type Order struct {
	UniqueID string `json:"uniqueId"`
}

type CandidateRestAndDish struct {
	RestaurantName string
	DishName       string

	TabUniqueId     string
	AddressUniqueId string
	TargetTime      string
	DishId          uint64

	Address     string
	PriceString string

	Title string
}
