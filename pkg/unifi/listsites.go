package unifi

type Sites struct {
	HostID     string     `json:"hostId"`
	IsOwner    bool       `json:"isOwner"`
	Meta       Meta       `json:"meta"`
	Permission string     `json:"permission"`
	SiteID     string     `json:"siteId"`
	Statistics Statistics `json:"statistics"`
}

type Meta struct {
	Description string `json:"desc"`
	GatewayMAC  string `json:"gatewayMac"`
	Name        string `json:"name"`
	Timezone    string `json:"timezone"`
}

type Statistics struct {
	Counts         Counters        `json:"counts"`
	InternetIssues []InternetIssue `json:"internetIssues"`
	ISPInfo        ISPInfo         `json:"ispInfo"`
	Percentages    Percentages     `json:"percentages"`
}

type Counters struct {
	CriticalNotification int `json:"criticalNotification"`
	GatewayDevice        int `json:"gatewayDevice"`
	GuestClient          int `json:"guestClient"`
	LANConfiguration     int `json:"lanConfiguration"`
	OfflineDevice        int `json:"offlineDevice"`
	OfflineGatewayDevice int `json:"offlineGatewayDevice"`
	OfflineWifiDevice    int `json:"offlineWifiDevice"`
	OfflineWiredDevice   int `json:"offlineWiredDevice"`
	PendingUpdateDevice  int `json:"pendingUpdateDevice"`
	TotalDevice          int `json:"totalDevice"`
	WANConfiguration     int `json:"wanConfiguration"`
	WIFIClient           int `json:"wifiClient"`
	WIFIConfiguration    int `json:"wifiConfiguration"`
	WIFIDevice           int `json:"wifiDevice"`
	WiredClient          int `json:"wiredClient"`
	WiredDevice          int `json:"wiredDevice"`
}

type InternetIssue struct {
	HighLatency  bool `json:"highLatency"`
	Index        int  `json:"index"`
	LatencyAvgMS int  `json:"latencyAvgMs"`
	LatencyMaxMS int  `json:"latencyMaxMs"`
}

type ISPInfo struct {
	Name         string `json:"name"`
	Organization string `json:"organization"`
}

type Percentages struct {
	TXRetry   float64 `json:"txRetry"`
	WANUptime float64 `json:"wanUptime"`
}

func (u *Context) ListSites(nextToken string) ([]Sites, Status) {
	sites := []Sites{}
	params := map[string]string{
		"nextToken": nextToken,
	}
	status := u.Get("/ea/sites", params, &sites)
	return sites, status
}
