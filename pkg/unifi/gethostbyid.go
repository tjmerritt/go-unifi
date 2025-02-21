package unifi

import (
	"encoding/json"
	"time"

	"github.com/araddon/dateparse"
)

type TimeString struct {
	String string    `json:"time"`
	Time   time.Time `json:"-"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ts *TimeString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	ts.String = str
	parsed, err := dateparse.ParseAny(str)
	if err == nil {
		ts.Time = parsed
	}
	return nil
}

type HostInfo struct {
	HardwareID                string        `json:"hardwareId"`
	ID                        string        `json:"id"`
	IPAddress                 string        `json:"ipAddress"`
	IsBlocked                 bool          `json:"isBlocked"`
	LastConnectionStateChange TimeString    `json:"lastConnectionStateChange"`
	LatestBackupTime          TimeString    `json:"latestBackupTime"`
	Owner                     bool          `json:"owner"`
	RegistrationTime          TimeString    `json:"registrationTime"`
	ReportedState             ReportedState `json:"reportedState"`
	Type                      string        `json:"type"`
	UserData                  UserData      `json:"userData"`
}

type ReportedState struct {
	AnonID                     string         `json:"anonid"`
	Apps                       []App          `json:"apps"`
	AvailableChannels          []string       `json:"availableChannels"`
	ConsolesOnSameLocalNetwork any            `json:"consolesOnSameLocalNetwork"`
	ControllerUUID             string         `json:"controller_uuid"`
	Controllers                []Controller   `json:"controllers"`
	Country                    int            `json:"country"`
	DeviceErrorCode            any            `json:"deviceErrorCode"`
	DeviceState                string         `json:"deviceState"`
	DeviceStateLastChanged     int            `json:"deviceStateLastChanged"`
	DirectConnectDomain        string         `json:"directConnectDomain"`
	Features                   Features       `json:"features"`
	FirmwareUpdate             FirmwareUpdate `json:"firmwareUpdate"`
	Hardware                   Hardware       `json:"hardware"`
	HostType                   int            `json:"host_type"`
	Hostname                   string         `json:"hostname"`
	InternetIssues5min         InternetIssues `json:"internetIssues5min"`
	IP                         string         `json:"ip"`
	IPAddrs                    []string       `json:"ipAddrs"`
	IsStacked                  bool           `json:"isStacked"`
	Location                   Location       `json:"location"`
	MAC                        string         `json:"mac"`
	MgmtPort                   int            `json:"mgmt_port"`
	Name                       string         `json:"name"`
	ReleaseChannel             string         `json:"releaseChannel"`
	State                      string         `json:"state"`
	Timezone                   string         `json:"timezone"`
	UIDB                       UIDB           `json:"uidb"`
	UnadoptedUnifiOSDevices    any            `json:"unadoptedUnifiOSDevices"`
	Version                    string         `json:"version"`
}

type App struct {
	ControllerStatus string `json:"controllerStatus"`
	Name             string `json:"name"`
	Port             int    `json:"port"`
	SwaiVersion      int    `json:"swaiVersion"`
	Type             string `json:"type"`
	Version          string `json:"version"`
}

type Controller struct {
	Abridged                bool     `json:"abridged"`
	ControllerStatus        string   `json:"controllerStatus"`
	Features                Features `json:"features,omitempty"`
	InitialDeviceListSynced bool     `json:"initialDeviceListSynced"`
	InstallState            string   `json:"installState"`
	Installable             bool     `json:"installable"`
	IsConfigured            bool     `json:"isConfigured"`
	IsGeofencingEnabled     bool     `json:"isGeofencingEnabled"`
	IsInstalled             bool     `json:"isInstalled"`
	IsRunning               bool     `json:"isRunning"`
	Name                    string   `json:"name"`
	Port                    int      `json:"port"`
	ReleaseChannel          string   `json:"releaseChannel"`
	Required                bool     `json:"required"`
	RestorePercentage       int      `json:"restorePercentage"`
	State                   string   `json:"state"`
	Status                  string   `json:"status"`
	StatusMessage           string   `json:"statusMessage"`
	SwaiVersion             int      `json:"swaiVersion"`
	Type                    string   `json:"type"`
	UIVersion               string   `json:"uiVersion"`
	UnadoptedDevices        any      `json:"unadoptedDevices"`
	Updatable               bool     `json:"updatable"`
	UpdateAvailable         any      `json:"updateAvailable"`
	Version                 string   `json:"version"`
}

type Features struct {
	Cloud                        CloudFeatures      `json:"cloud"`
	CloudBackup                  bool               `json:"cloudBackup"`
	DeviceList                   DeviceListFeatures `json:"deviceList"`
	DirectRemoteConnection       bool               `json:"directRemoteConnection"`
	HasGateway                   bool               `json:"hasGateway"`
	HasLCM                       bool               `json:"hasLCM"`
	HasLED                       bool               `json:"hasLED"`
	InfoApis                     InfoApis           `json:"infoApis"`
	IsAutomaticFailoverAvailable bool               `json:"isAutomaticFailoverAvailable"`
	MFA                          bool               `json:"mfa"`
	Notifications                bool               `json:"notifications"`
	SharedTokens                 bool               `json:"sharedTokens"`
	SupportForm                  bool               `json:"supportForm"`
	Teleport                     bool               `json:"teleport"`
	TeleportState                string             `json:"teleportState"`
	UIDService                   bool               `json:"uidService"`
}

type CloudFeatures struct {
	ApplicationEvents     bool `json:"applicationEvents"`
	ApplicationEventsHttp bool `json:"applicationEventsHttp"`
}

type DeviceListFeatures struct {
	AutolinkDevices bool `json:"autolinkDevices"`
	PartialUpdates  bool `json:"partialUpdates"`
	UCP4Events      bool `json:"ucp4Events"`
}

type InfoApis struct {
	FirmwareUpdate bool `json:"firmwareUpdate"`
}

type FirmwareUpdate struct {
	LatestAvailableVersion any `json:"latestAvailableVersion"`
}

type Hardware struct {
	BOM             string `json:"bom"`
	CPUID           string `json:"cpu.id"`
	DebianCodename  string `json:"debianCodename"`
	FirmwareVersion string `json:"firmwareVersion"`
	HWRev           int    `json:"hwrev"`
	MAC             string `json:"mac"`
	Name            string `json:"name"`
	QRID            string `json:"qrid"`
	Reboot          string `json:"reboot"`
	SerialNo        string `json:"serialno"`
	ShortName       string `json:"shortname"`
	Subtype         string `json:"subtype"`
	SysID           int    `json:"sysid"`
	Upgrade         string `json:"upgrade"`
	UUID            string `json:"uuid"`
}

type InternetIssues struct {
	Periods []Period `json:"periods"`
}

type Period struct {
	Index int `json:"index"`
}

type Location struct {
	Lat    float64 `json:"lat"`
	Long   float64 `json:"long"`
	Radius int     `json:"radius"`
	Text   string  `json:"text"`
}

type UIDB struct {
	GUID   string     `json:"guid"`
	ID     string     `json:"id"`
	Images UIDBImages `json:"images"`
}

type UIDBImages struct {
	Default   string `json:"default"`
	NoPadding string `json:"nopadding"`
	Topology  string `json:"topology"`
}

type UserData struct {
	Apps                []string             `json:"apps"`
	ConsoleGroupMembers []ConsoleGroupMember `json:"consoleGroupMembers"`
	Controllers         []string             `json:"controllers"`
	Email               string               `json:"email"`
	Features            UserFeatures         `json:"features"`
	FullName            string               `json:"fullName"`
	LocalID             string               `json:"localId"`
	Permissions         Permissions          `json:"permissions"`
	Role                string               `json:"role"`
	RoleID              string               `json:"roleId"`
	Status              string               `json:"status"`
}

type ConsoleGroupMember struct {
	MAC            string         `json:"mac"`
	Role           string         `json:"role"`
	RoleAttributes RoleAttributes `json:"roleAttributes"`
	SysID          int            `json:"sysId"`
}

type RoleAttributes struct {
	Applications              Applications `json:"applications"`
	CandidateRoles            []string     `json:"candidateRoles"`
	ConnectedState            string       `json:"connectedState"`
	ConnectedStateLastChanged TimeString   `json:"connectedStateLastChanged"`
}

type Applications struct {
	Access     ApplicationStatus `json:"access"`
	Connect    ApplicationStatus `json:"connect"`
	Innerspace ApplicationStatus `json:"innerspace"`
	Network    ApplicationStatus `json:"network"`
	Protect    ApplicationStatus `json:"protect"`
	Talk       ApplicationStatus `json:"talk"`
}

type ApplicationStatus struct {
	Owned     bool `json:"owned"`
	Required  bool `json:"required"`
	Supported bool `json:"supported"`
}

type UserFeatures struct {
	DeviceGroups       bool              `json:"deviceGroups"`
	Floorplan          FloorplanFeatures `json:"floorplan"`
	ManageApplications bool              `json:"manageApplications"`
	Notifications      bool              `json:"notifications"`
	Pion               bool              `json:"pion"`
	WebRTC             WebRTCFeatures    `json:"webrtc"`
}

type FloorplanFeatures struct {
	CanEdit bool `json:"canEdit"`
	CanView bool `json:"canView"`
}

type WebRTCFeatures struct {
	ICERestart   bool `json:"iceRestart"`
	MediaStreams bool `json:"mediaStreams"`
	TwoWayAudio  bool `json:"twoWayAudio"`
}

type Permissions struct {
	NetworkManagement        []string `json:"network.management"`
	ProtectManagement        []string `json:"protect.management"`
	SystemManagementLocation []string `json:"system.management.location"`
	SystemManagementUser     []string `json:"system.management.user"`
}

func (u *Context) GetHostByID(hostID string) (HostInfo, Status) {
	data := HostInfo{}
	params := map[string]string{
		"hostID": hostID,
	}
	status := u.Get("/ea/hosts/{hostID}", params, &data)
	if status.Failed() {
		return HostInfo{}, status
	}
	return data, status
}
