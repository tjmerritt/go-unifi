package unifi

import (
	"time"
)

type QueryISPMetricsSiteReq struct {
	BeginTimestamp time.Time `json:"beginTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
	HostID         string    `json:"hostId"`
	SiteID         string    `json:"siteId"`
}

type QueryISPMetricsReq struct {
	Sites []QueryISPMetricsSiteReq `json:"sites"`
}

func (u *Context) QueryISPMetrics(sampleRate time.Duration, req *QueryISPMetricsReq) ([]MetricData, Status) {
	data := []MetricData{}
	params := map[string]string{
		"type": convertSampleRate(sampleRate),
	}
	status := u.Post("/ea/isp-metrics/{type}/query", params, req, &data)
	return data, status
}
