package unifi

import (
	"time"
)

type Host struct {
	HostID       string `json:"hostId"`
	HostName     string `json:"hostName"`
	UpdatedAt    time.Time
	UpdatedAtStr string   `json:"updatedAt"`
	Devices      []Device `json:"devices"`
}

type Device struct {
	AdoptionTime    time.Time
	AdoptionTimeStr string `json:"adoptionTime"`
	FirmwareStatus  string `json:"firmwareStatus"`
	DeviceID        string `json:"id"`
	IP              string `json:"ip"`
	IsConsole       bool   `json:"isConsole"`
	IsManaged       bool   `json:"isManaged"`
	MAC             string `json:"mac"`
	Model           string `json:"model"`
	Name            string `json:"name"`
	Note            any    `json:"note"`
	ProductLine     string `json:"productLine"`
	ShortName       string `json:"shortname"`
	StartupTime     time.Time
	StartupTimeStr  string `json:"startupTime"`
	Status          string `json:"status"`
	UpdateAvailable any    `json:"updateAvailable"`
	Vesion          string `json:"version"`
	UIDB            struct {
		GUID   string `json:"guid"`
		ID     string `json:"id"`
		Images struct {
			Default   string `json:"default"`
			NoPadding string `json:"nopadding"`
			Topology  string `json:"topology"`
		} `json:"images"`
	} `json:"uidb"`
}

func (u *Context) ListDevices(nextToken string, hosts []string, lastProcessed time.Time) ([]Host, Status) {
	resp := []Host{}
	params := map[string]string{
		"nextToken": nextToken,
	}
	status := u.Get("/ea/devices", params, &resp)
	return resp, status
}
