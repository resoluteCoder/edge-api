package models

// EdgeDevice is the entity that represents and Edge Device
// It is a combination of the data of a Device owned by Inventory API
// and the Device data saved on Edge API
type EdgeDevice struct {
	*Device
	DeviceName string // TODO: Can be deleted when we start saving the name field on the Device object
	LastSeen   string // TODO: Can be deleted when we start saving lastseen field on the Device object
	// Booted status is referring to the LastDeployment of this device
	// TODO: Needs to be rethinked when we get to the greenbot epic
	Booted  bool
	Account string
}

// DeviceDetails is a Device with Image and Update transactions
// It contains data from multiple tables on the database
type DeviceDetails struct {
	Device             EdgeDevice           `json:"Device,omitempty"`
	Image              *ImageInfo           `json:"ImageInfo"`
	UpdateTransactions *[]UpdateTransaction `json:"UpdateTransactions,omitempty"`
	DevicesGroups      *[]DeviceGroup       `json:"DevicesGroups,omitempty"`
	Updating           *bool                `json:"DeviceUpdating,omitempty"`
}

// DeviceDetailsList is the list of devices with details from Inventory and Edge API
type DeviceDetailsList struct {
	Total   int             `json:"total"`
	Count   int             `json:"count"`
	Devices []DeviceDetails `json:"data"`
}

// DeviceViewList is the list of devices for a given account, formatted for the UI
type DeviceViewList struct {
	Total   int          `json:"total"`
	Devices []DeviceView `json:"devices"`
}

// DeviceView is the device information needed for the UI
type DeviceView struct {
	DeviceID        uint   `json:"DeviceID"`
	DeviceName      string `json:"DeviceName"`
	DeviceUUID      string `json:"DeviceUUID"`
	DeviceGroupID   uint   `json:"DeviceGroupID"`
	DeviceGroupName string `json:"DeviceGroupName"`
	ImageID         uint   `json:"ImageID"`
	ImageName       string `json:"ImageName"`
	LastSeen        string `json:"LastSeen"`
	UpdateAvailable bool   `json:"UpdateAvailable"`
	Status          string `json:"Status"`
	ImageSetID      uint   `json:"ImageSetID"`
}

// Device is a record of Edge Devices referenced by their UUID as per the
// cloud.redhat.com Inventory.
//
//	Connected refers to the devices Cloud Connector state, 0 is unavailable
//	and 1 is reachable.
//
// THEEDGE-1921 created 2 temporary indexes to address production issue
type Device struct {
	Model
	UUID              string              `gorm:"index" json:"UUID"`
	AvailableHash     string              `json:"AvailableHash,omitempty"`
	RHCClientID       string              `json:"RHCClientID"`
	Connected         bool                `gorm:"default:true" json:"Connected"`
	Name              string              `json:"Name"`
	LastSeen          EdgeAPITime         `json:"LastSeen"`
	CurrentHash       string              `json:"CurrentHash,omitempty"`
	Account           string              `gorm:"index" json:"Account"`
	ImageID           uint                `json:"ImageID"`
	UpdateAvailable   bool                `json:"UpdateAvailable"`
	DevicesGroups     []DeviceGroup       `faker:"-" gorm:"many2many:device_groups_devices;" json:"DevicesGroups"`
	UpdateTransaction []UpdateTransaction `faker:"-" gorm:"many2many:updatetransaction_devices;" json:"UpdateTransaction"`
}
