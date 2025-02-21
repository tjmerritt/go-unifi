package unifi

import (
	"time"
)

type MetricData struct {
	MetricType string         `json:"metricType"`
	Periods    []MetricPeriod `json:"periods"`
	HostID     string         `json:"hostId"`
	SiteID     string         `json:"siteId"`
}

type MetricPeriod struct {
	Data       PeriodData `json:"data"`
	MetricTime TimeString `json:"metricTime"`
	Version    string     `json:"version"`
}

type PeriodData struct {
	WAN WANMetrics `json:"wan"`
}

type WANMetrics struct {
	AvgLatency   int     `json:"avgLatency"`
	DownloadKbps int     `json:"download_kbps"`
	Downtime     float64 `json:"downtime"`
	ISPAsn       string  `json:"ispAsn"`
	ISPName      string  `json:"ispName"`
	MaxLatency   int     `json:"maxLatency"`
	PacketLoss   float64 `json:"packetLoss"`
	UploadKbps   int     `json:"upload_kbps"`
	Uptime       float64 `json:"uptime"`
}

func (u *Context) GetISPMetrics(sampleRate time.Duration, begin, end time.Time) ([]MetricData, Status) {
	data := []MetricData{}
	params := map[string]string{
		"sampleRate": convertSampleRate(sampleRate),
		"begin":      begin.Format(time.DateTime),
		"end":        end.Format(time.DateTime),
	}
	status := u.Get("/ea/isp-metrics/{type}", params, &data)
	return data, status
}

func convertSampleRate(d time.Duration) string {
	switch d {
	case 5 * time.Minute:
		return "5m"
	case time.Hour:
		return "1h"
	default:
		return "5m"
	}
}
