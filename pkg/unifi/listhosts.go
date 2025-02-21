package unifi

import (
	"strings"
	"time"
)

func (u *Context) ListHosts(nextToken string, hosts []string, lastProcessed time.Time) ([]HostInfo, Status) {
	data := []HostInfo{}
	params := map[string]string{
		"nextToken":     nextToken,
		"hosts":         strings.Join(hosts, ","),
		"lastProcessed": lastProcessed.Format(time.DateTime),
	}
	status := u.Get("/ea/hosts/{hostID}", params, &data)
	if status.Failed() {
		return []HostInfo{}, status
	}
	return data, status
}
