package ovirt

import (
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/model/base"
)

//
// Errors
var NotFound = libmodel.NotFound

type InvalidRefError = base.InvalidRefError

const (
	MaxDetail = base.MaxDetail
)

//
// Types
type Model = base.Model
type ListOptions = base.ListOptions
type Concern = base.Concern
type Ref = base.Ref

//
// Base oVirt model.
type Base struct {
	// Managed object ID.
	ID string `sql:"pk"`
	// Name
	Name string `sql:"d0,index(name)"`
	// Revision
	Description string `sql:"d0"`
	// Revision
	Revision int64 `sql:"incremented,d0,index(revision)"`
}

//
// Get the PK.
func (m *Base) Pk() string {
	return m.ID
}

//
// String representation.
func (m *Base) String() string {
	return m.ID
}

type DataCenter struct {
	Base
}

type Cluster struct {
	Base
	DataCenter    string `sql:"d0,index(dataCenter)"`
	HaReservation bool   `sql:""`
	KsmEnabled    bool   `sql:""`
}

type Network struct {
	Base
	DataCenter string   `sql:"d0,index(dataCenter)"`
	VLan       string   `sql:""`
	Usages     []string `sql:""`
	Profiles   []string `sql:""`
}

type NICProfile struct {
	Base
	Network       string `sql:"d0,index(network)"`
	PortMirroring bool   `sql:""`
	NetworkFilter string `sql:""`
	QoS           string `sql:""`
}

type DiskProfile struct {
	Base
	StorageDomain string `sql:"d0,index(storageDomain)"`
	QoS           string `sql:""`
}

type StorageDomain struct {
	Base
	DataCenter string `sql:"d0,index(dataCenter)"`
	Type       string `sql:""`
	Storage    struct {
		Type string
	} `sql:""`
	Available int64 `sql:""`
	Used      int64 `sql:""`
}

type Host struct {
	Base
	Cluster            string              `sql:"d0,index(cluster)"`
	ProductName        string              `sql:""`
	ProductVersion     string              `sql:""`
	InMaintenance      bool                `sql:""`
	CpuSockets         int16               `sql:""`
	CpuCores           int16               `sql:""`
	NetworkAttachments []NetworkAttachment `sql:""`
	NICs               []HostNIC           `sql:""`
}

type NetworkAttachment struct {
	ID      string `json:"id"`
	Network string `json:"network"`
}

type HostNIC struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LinkSpeed int64  `json:"linkSpeed"`
	MTU       int64  `json:"mtu"`
	VLan      string `json:"vlan"`
}

type VM struct {
	Base
	Cluster                     string           `sql:"d0,index(cluster)"`
	Host                        string           `sql:"d0,index(host)"`
	RevisionValidated           int64            `sql:"d0,index(revisionValidated)" eq:"-"`
	PolicyVersion               int              `sql:"d0,index(policyVersion)" eq:"-"`
	GuestName                   string           `sql:""`
	CpuSockets                  int16            `sql:""`
	CpuCores                    int16            `sql:""`
	CpuAffinity                 []CpuPinning     `sql:""`
	CpuShares                   int16            `sql:""`
	Memory                      int64            `sql:""`
	BalloonedMemory             bool             `sql:""`
	BIOS                        string           `sql:""`
	Display                     string           `sql:""`
	IOThreads                   int16            `sql:""`
	StorageErrorResumeBehaviour string           `sql:""`
	HaEnabled                   bool             `sql:""`
	UsbEnabled                  bool             `sql:""`
	BootMenuEnabled             bool             `sql:""`
	PlacementPolicyAffinity     string           `sql:""`
	Timezone                    string           `sql:""`
	Status                      string           `sql:""`
	Stateless                   string           `sql:""`
	HasIllegalImages            bool             `sql:""`
	NumaNodeAffinity            []string         `sql:""`
	LeaseStorageDomain          string           `sql:""`
	DiskAttachments             []DiskAttachment `sql:""`
	NICs                        []NIC            `sql:""`
	HostDevices                 []HostDevice     `sql:""`
	CDROMs                      []CDROM          `sql:""`
	WatchDogs                   []WatchDog       `sql:""`
	Properties                  []Property       `sql:""`
	Snapshots                   []Snapshot       `sql:""`
	Concerns                    []Concern        `sql:"" eq:"-"`
}

//
// Determine if current revision has been validated.
func (m *VM) Validated() bool {
	return m.RevisionValidated == m.Revision
}

type Snapshot struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	Type          string `json:"type"`
	PersistMemory bool   `json:"persistMemory"`
}

type DiskAttachment struct {
	ID              string `json:"id"`
	Interface       string `json:"interface"`
	SCSIReservation bool   `json:"scsiReservation"`
	Disk            string `json:"disk"`
}

type NIC struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Interface  string      `json:"interface"`
	Plugged    bool        `json:"plugged"`
	IpAddress  []IpAddress `json:"ipAddress"`
	Profile    string      `json:"profile"`
	Properties []Property  `json:"properties"`
}

type IpAddress struct {
	Address string `json:"address"`
	Version string `json:"version"`
}

type CpuPinning struct {
	Set int32 `json:"set"`
	Cpu int32 `json:"cpu"`
}

type HostDevice struct {
	Capability string `json:"capability"`
	Product    string `json:"product"`
	Vendor     string `json:"vendor"`
}

type CDROM struct {
	ID   string `json:"id"`
	File string `json:"file,omitempty"`
}

type WatchDog struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Model  string `json:"model"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Disk struct {
	Base
	Shared          bool   `sql:""`
	Profile         string `sql:"index(profile)"`
	StorageDomain   string `sql:"index(storageDomain)"`
	Status          string `sql:""`
	ActualSize      int64  `sql:""`
	Backup          string `sql:""`
	StorageType     string `sql:""`
	ProvisionedSize int64  `sql:""`
}
