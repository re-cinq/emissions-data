package compute

type AvergaeWattsOptions struct {
	MinWatts          float64
	MaxWatts          float64
	AvgCPUUtilization float64
}

func AvergaeWatts(o AvergaeWattsOptions) float64 {
	return o.MinWatts + o.AvgCPUUtilization*(o.MaxWatts-o.MinWatts)
}

type ComputeWattHoursOptions struct {
	AverageWatts float64
	//vCPU Hours
	CPUHours float64
}

func ComputeWattHours(o ComputeWattHoursOptions) float64 {
	return o.AverageWatts * o.CPUHours
}
