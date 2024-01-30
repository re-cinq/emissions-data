package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	diff "github.com/r3labs/diff/v3"
	"github.com/stretchr/testify/assert"

	v2 "github.com/re-cinq/emissions-data/pkg/types/v2"
)

func TestGetInstanceData(t *testing.T) {
	testPath := "testdata/test-input-AWS-EC2-Dataset.csv"

	expected := []v2.Instance{
		{
			Kind:                 "a1.medium",
			VCPU:                 1,
			MemoryGB:             2,
			StorageInfoAndSizeGB: "EBS-Only",
			StorageType:          "EBS",
			GPUCount:             0,
			GPUMemoryGB:          0,
			PkgWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    0.29,
				},
				{
					Percentage: 10,
					Wattage:    0.80,
				},
				{
					Percentage: 50,
					Wattage:    1.88,
				},
				{
					Percentage: 100,
					Wattage:    2.55,
				},
			},
			RAMWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    0.40,
				},
				{
					Percentage: 10,
					Wattage:    0.60,
				},
				{
					Percentage: 50,
					Wattage:    0.80,
				},
				{
					Percentage: 100,
					Wattage:    1.20,
				},
			},
			GPUWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    0,
				},
				{
					Percentage: 10,
					Wattage:    0,
				},
				{
					Percentage: 50,
					Wattage:    0,
				},
				{
					Percentage: 100,
					Wattage:    0,
				},
			},
			TotalWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    1.2,
				},
				{
					Percentage: 10,
					Wattage:    1.9,
				},
				{
					Percentage: 50,
					Wattage:    3.2,
				},
				{
					Percentage: 100,
					Wattage:    4.2,
				},
			},
			DeltaFullMachine:    0.5,
			EmbodiedHourlyGCO2e: 1.8,
			Platform: v2.Platform{
				VCPU:                16,
				Architecture:        "Graviton",
				MemoryGB:            32,
				StorageDriveCount:   0,
				GPUCount:            0,
				GPUName:             "N/A",
				MemoryScope3:        22.2,
				StorageScope3:       0,
				GPUScope3:           0,
				CPUScope3:           0,
				TotalScope3:         1022.2,
				HardwareInformation: "AWS Graviton (ARM)",
			},
		},
		{
			Kind:                 "c3.xlarge",
			VCPU:                 4,
			MemoryGB:             7.5,
			StorageInfoAndSizeGB: "2 x 40 (SSD)",
			StorageType:          "SSD",
			GPUCount:             0,
			GPUMemoryGB:          0,
			PkgWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    2.77,
				},
				{
					Percentage: 10,
					Wattage:    7.92,
				},
				{
					Percentage: 50,
					Wattage:    16.29,
				},
				{
					Percentage: 100,
					Wattage:    22.30,
				},
			},
			RAMWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    1.50,
				},
				{
					Percentage: 10,
					Wattage:    2.25,
				},
				{
					Percentage: 50,
					Wattage:    3,
				},
				{
					Percentage: 100,
					Wattage:    4.5,
				},
			},
			GPUWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    0,
				},
				{
					Percentage: 10,
					Wattage:    0,
				},
				{
					Percentage: 50,
					Wattage:    0,
				},
				{
					Percentage: 100,
					Wattage:    0,
				},
			},
			TotalWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    8.9,
				},
				{
					Percentage: 10,
					Wattage:    14.8,
				},
				{
					Percentage: 50,
					Wattage:    23.9,
				},
				{
					Percentage: 100,
					Wattage:    31.4,
				},
			},
			DeltaFullMachine:    4.6,
			EmbodiedHourlyGCO2e: 3.9,
			Platform: v2.Platform{
				VCPU:                40,
				Architecture:        "Xeon E5-2680 v2",
				MemoryGB:            60,
				StorageDriveCount:   2,
				GPUCount:            0,
				GPUName:             "N/A",
				MemoryScope3:        61,
				StorageScope3:       200,
				GPUScope3:           0,
				CPUScope3:           100,
				TotalScope3:         1361,
				HardwareInformation: "2.8 GHz Intel Xeon E5-2680v2 (Ivy Bridge) processor",
			},
		},
		{
			Kind:                 "c5ad.4xlarge",
			VCPU:                 16,
			MemoryGB:             32,
			StorageInfoAndSizeGB: "2 x 300 NVMe SSD",
			StorageType:          "SSD",
			GPUCount:             0,
			GPUMemoryGB:          0,
			PkgWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    5.42,
				},
				{
					Percentage: 10,
					Wattage:    14.85,
				},
				{
					Percentage: 50,
					Wattage:    35.11,
				},
				{
					Percentage: 100,
					Wattage:    47.53,
				},
			},
			RAMWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    6.40,
				},
				{
					Percentage: 10,
					Wattage:    9.60,
				},
				{
					Percentage: 50,
					Wattage:    12.80,
				},
				{
					Percentage: 100,
					Wattage:    19.20,
				},
			},
			GPUWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    0,
				},
				{
					Percentage: 10,
					Wattage:    0,
				},
				{
					Percentage: 50,
					Wattage:    0,
				},
				{
					Percentage: 100,
					Wattage:    0,
				},
			},
			TotalWatt: []v2.Wattage{
				{
					Percentage: 0,
					Wattage:    21.2,
				},
				{
					Percentage: 10,
					Wattage:    33.8,
				},
				{
					Percentage: 50,
					Wattage:    57.2,
				},
				{
					Percentage: 100,
					Wattage:    76.1,
				},
			},
			DeltaFullMachine:    9.3,
			EmbodiedHourlyGCO2e: 7.0,
			Platform: v2.Platform{
				VCPU:                96,
				Architecture:        "EPYC 7R32",
				MemoryGB:            192,
				StorageDriveCount:   2,
				GPUCount:            0,
				GPUName:             "N/A",
				MemoryScope3:        244.1,
				StorageScope3:       200,
				GPUScope3:           0,
				CPUScope3:           0,
				TotalScope3:         1444.1,
				HardwareInformation: "Up to 3.3 GHz 2nd generation AMD EPYC 7002 Series Processor",
			},
		},
	}

	output, err := getInstanceData(testPath)
	assert.Nil(t, err)
	if !assert.True(t, cmp.Equal(output, expected)) {
		changes, err := diff.Diff(output, expected)
		if err != nil {
			fmt.Println("error with diff: ", err)
		}

		for _, chn := range changes {
			fmt.Printf("DIFF: actual: %+v expected: %+v\n", chn.From, chn.To)
		}
	}
}
