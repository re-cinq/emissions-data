package v2

type Wattage struct {
	Percentage int
	Wattage    float64
}

type Instance struct {
	Kind                 string
	VCPU                 int
	MemoryGB             float64
	GPUCount             int
	GPUMemoryGB          float64
	StorageInfoAndSizeGB string
	StorageType          string
	PkgWatt              []Wattage
	RAMWatt              []Wattage
	GPUWatt              []Wattage
	TotalWatt            []Wattage
	DeltaFullMachine     float64
	EmbodiedHourlyGCO2e  float64
	Platform
}

type Platform struct {
	Architecture        string
	HardwareInformation string
	VCPU                int
	MemoryGB            float64
	StorageDriveCount   int
	GPUCount            int
	GPUName             string
	MemoryScope3        float64
	StorageScope3       float64
	GPUScope3           float64
	CPUScope3           float64
	TotalScope3         float64
}
