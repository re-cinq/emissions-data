package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	v2 "github.com/re-cinq/emissions-data/pkg/types/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	path := "data/v2/input-AWS-EC2-Dataset.csv"
	instances, err := getInstanceData(path)
	if err != nil {
		fmt.Println("error getting instances: ", err)
		return
	}

	yamlData, _ := yaml.Marshal(&instances)
	fileName := "data/v2/aws-instances.yaml"
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		fmt.Println("Unable to write data into the YAML file: ", err)
		return
	}
}

// getInstanceData parses the csv data to the Instance and Platform structs
// It requires the data to have the following column ordering

// s 0: Instance.Kind
//   1: -------------
// i 2: Instance.VCPU
// i 3: Platform.VCPU
// s 4: Platform.Architecture
// f 5: Instance.MemoryGB
// f 6: Platform.MemoryGB
// s 7: Instance.StorageInfoAndSizeGB
// s 8: Instance.StorageType
// i 9: Platform.StorageDriveCount
// i 10: Platform.GPUCount
// s 11: Platform.GPUName
// i 12: Instance.GPUCount
// f 13: Instance.GPUMemoryGB
// w 14: Instance.PkgWatt 0
// w 15: Instance.PkgWatt 10
// w 16: Instance.PkgWatt 50
// w 17: Instance.PkgWatt 100
// w 18: Instance.RAMWatt 0
// w 19: Instance.RAMWatt 10
// w 20: Instance.RAMWatt 50
// w 21: Instance.RAMWatt 100
// w 22: Instance.GPUWatt 0
// w 23: Instance.GPUWatt 10
// w 24: Instance.GPUWatt 50
// w 25: Instance.GPUWatt 100
// w 26: Instance.DeltaFullMachine
// w 27: Instance.TotalWatt 0
// w 28: Instance.TotalWatt 10
// w 29: Instance.TotalWatt 50
// w 30: Instance.TotalWatt 100
// f 31: Platform.MemoryScope3
// f 32: Platform.StorageScope3
// f 33: Platform.GPUScope3
// f 34: Platform.CPUScope3
// f 35: Platform.TotalScope3
// f 36: Instance.EmbodiedHourlyGCO2e
// s 37: Platform.HardwareInformation

func getInstanceData(path string) ([]v2.Instance, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error while getting the file: %+v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading in csv: %+v", err)
	}

	var instances []v2.Instance
	for _, record := range records[1:] {

		ivCPU, err := parseInt(record[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing instance vCPU: %+v", err)
		}

		pvCPU, err := parseInt(record[3])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform vCPU: %+v", err)
		}

		iMem, err := parseFloat(record[5])
		if err != nil {
			return nil, fmt.Errorf("error parsing instance memory: %+v", err)
		}

		pMem, err := parseFloat(record[6])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform memory: %+v", err)
		}

		pStorCt, err := parseInt(record[9])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform storage drive count: %+v", err)
		}

		pGPUCt, err := parseInt(record[10])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform GPU count: %+v", err)
		}

		iGPUCt, err := parseInt(record[12])
		if err != nil {
			return nil, fmt.Errorf("error parsing instance GPU count: %+v", err)
		}

		iGPUMem, err := parseFloat(record[13])
		if err != nil {
			return nil, fmt.Errorf("error parsing instance GPU Memory: %+v", err)
		}

		iPkgWatt, err := parseWattage(record[14:18]...)
		if err != nil {
			return nil, fmt.Errorf("error parsing instance PkgWattage: %+v", err)
		}

		iRAMWatt, err := parseWattage(record[18:22]...)
		if err != nil {
			return nil, fmt.Errorf("error parsing instance RAMWattage: %+v", err)
		}

		iGPUWatt, err := parseWattage(record[22:26]...)
		if err != nil {
			return nil, fmt.Errorf("error parsing instance GPUWattage: %+v", err)
		}

		iTotalWatt, err := parseWattage(record[27:31]...)
		if err != nil {
			return nil, fmt.Errorf("error parsing instance Total Wattage: %+v", err)
		}

		iDelta, err := parseFloat(record[26])
		if err != nil {
			return nil, fmt.Errorf("error parsing instance Delta Full Machine: %+v", err)
		}

		pMemS3, err := parseFloat(record[31])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform Scope 3 Memory Emissions: %+v", err)
		}

		pStorS3, err := parseFloat(record[32])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform Scope 3 Storage Emissions: %+v", err)
		}

		pGPUS3, err := parseFloat(record[33])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform Scope 3 GPU Emissions: %+v", err)
		}

		pCPUS3, err := parseFloat(record[34])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform Scope 3 CPU Emissions: %+v", err)
		}

		pTotalS3, err := parseFloat(record[35])
		if err != nil {
			return nil, fmt.Errorf("error parsing platform Scope 3 Total Emissions: %+v", err)
		}

		iEmbodiedHr, err := parseFloat(record[36])
		if err != nil {
			return nil, fmt.Errorf("error parsing instance hourly embodied emissions: %+v", err)
		}

		i := v2.Instance{
			Kind:                 record[0],
			VCPU:                 ivCPU,
			MemoryGB:             iMem,
			GPUMemoryGB:          iGPUMem,
			GPUCount:             iGPUCt,
			StorageInfoAndSizeGB: record[7],
			StorageType:          record[8],
			PkgWatt:              iPkgWatt,
			RAMWatt:              iRAMWatt,
			GPUWatt:              iGPUWatt,
			TotalWatt:            iTotalWatt,
			DeltaFullMachine:     iDelta,
			EmbodiedHourlyGCO2e:  iEmbodiedHr,
		}
		p := v2.Platform{
			Architecture:        record[4],
			HardwareInformation: record[37],
			VCPU:                pvCPU,
			GPUName:             record[11],
			MemoryGB:            pMem,
			StorageDriveCount:   pStorCt,
			GPUCount:            pGPUCt,
			MemoryScope3:        pMemS3,
			StorageScope3:       pStorS3,
			GPUScope3:           pGPUS3,
			CPUScope3:           pCPUS3,
			TotalScope3:         pTotalS3,
		}

		i.Platform = p
		instances = append(instances, i)
	}

	return instances, nil
}

// parseInt converts the string to an Int, and
// checks the conversion matching the original
func parseInt(s string) (int, error) {
	if s == "N/A" {
		s = "0"
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return -1, fmt.Errorf("error parsing %s: %+v", s, err)
	}

	if strconv.Itoa(i) != s {
		return -1, fmt.Errorf("error validating: string(%s) does not equal int(%d)", s, i)
	}

	return i, nil
}

// parseFloat covnerts the input string to a
// float64 and validates the output
func parseFloat(s string) (float64, error) {
	if s == "N/A" {
		s = "0"
	}

	s = strings.ReplaceAll(s, ",", ".")

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1, fmt.Errorf("error parsing %s: %+v", s, err)
	}

	// TODO: Need to fix this. If the string is a decimal `0.29` the %0f is 0,
	// so then is a false negative
	//if fmt.Sprintf("%.0f", i) != s {
	//	fmt.Println("XXXXXX: ", i, s)
	//	return -1, fmt.Errorf("error validating: string(%s) does not equal float(%.0f)", s, i)
	//}

	return i, nil
}

func parseWattage(vals ...string) ([]v2.Wattage, error) {

	wattages := []v2.Wattage{}
	for i, val := range vals {
		v, err := parseFloat(val)
		if err != nil {
			return nil, fmt.Errorf("error parsing Wattage: %+v", err)
		}

		var p int
		switch i {
		case 0:
			p = 0
		case 1:
			p = 10
		case 2:
			p = 50
		case 3:
			p = 100
		}

		wattages = append(wattages, v2.Wattage{
			Percentage: p,
			Wattage:    v,
		})
	}

	return wattages, nil

}
