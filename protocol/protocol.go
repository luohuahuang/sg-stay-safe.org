package protocol

type CheckInEvent struct {
	AnonymousId string `json:"anonymous_id"`
	SiteId      string `json:"site_id"`
	Timestamp   int64  `json:"timestamp"`
}

type GeneralResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ViolationEvent struct {
	CheckInEvent    CheckInEvent
	GeneralResponse GeneralResponse
}

type CheckInResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BanSiteRequest struct {
	SiteId    string `json:"site_id"`
	ShouldBan bool   `json:"should_ban"`
}

type BanUserRequest struct {
	AnonymousId string `json:"anonymous_id"`
	ShouldBan   bool   `json:"should_ban"`
}

type Rule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsEnabled   bool   `json:"is_enabled"`
	Value       int    `json:"value"`
}

type ViolationEmail struct {
	Recipient   string `json:"recipient"`
	AnonymousId string `json:"anonymous_id"`
	Site        Site   `json:"site"`
}

type Site struct {
	Name        string `json:"name"`
	SiteId      string `json:"site_id"`
	Address     string `json:"address"`
	PostalCode  string `json:"postal_code"`
	Description string `json:"description"`
	Capacity    string `json:"capacity"`
	ShouldBan   bool   `json:"should_ban"`
}
