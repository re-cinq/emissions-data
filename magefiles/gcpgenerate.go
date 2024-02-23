//go:build mage

package main

import (
	"encoding/csv"
	"os"
	"strconv"

	_ "github.com/magefile/mage/sh"
	"gopkg.in/yaml.v3"
)

type gcpCSV struct {
	vCPU      float64
	totalVCPU float64
}

type Embodied struct {
	InstanceType      string  `yaml:"type"`
	Additionalmemory  float64 `yaml:"additionalmemory"`
	Additionalstorage float64 `yaml:"additionalstorage"`
	Additionalcpus    float64 `yaml:"additionalcpus"`
	Additionalgpus    float64 `yaml:"additionalgpus"`
	Total             float64 `yaml:"total"`
	Architecture      string  `yaml:"architecture"`
	TotalVCPU         float64 `yaml:"totalVCPU"`
	VCPU              float64 `yaml:"vCPU"`
}

// generates the yaml files for GCP emissions data from csv
func GenerateGCP() error {
	datacsv, err := getCSVData("data/raw/gcp-instance-data.csv")
	if err != nil {
		return err
	}

	data, err := loadEmbodiedYaml("data/v1/gcp-embodied.yaml")
	if err != nil {
		return err
	}

	for i, d := range data {
		v, ok := datacsv[d.InstanceType]
		if !ok {
			continue
		}
		d.TotalVCPU = v.totalVCPU
		d.VCPU = v.vCPU
		data[i] = d
	}
	err = saveEmbodiedYaml(data, "data/v1/gcp-embodied.yaml")
	if err != nil {
		return err
	}

	return nil
}

func getCSVData(filepath string) (map[string]gcpCSV, error) {
	instances := make(map[string]gcpCSV)
	// open file
	f, err := os.Open(filepath)
	if err != nil {
		return instances, err
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return instances, err
	}

	for i, line := range data {
		if i == 0 {
			continue
		}
		id := line[0]
		ins := gcpCSV{}
		if i, err := strconv.ParseFloat(line[2], 64); err == nil {
			ins.vCPU = i
		}
		if i, err := strconv.ParseFloat(line[3], 64); err == nil {
			ins.totalVCPU = i
		}
		// First we get a "copy" of the entry
		if _, ok := instances[id]; !ok {

			// Then we reassign map entry
			instances[id] = ins

		}
	}

	return instances, nil
}

func loadEmbodiedYaml(filepath string) ([]Embodied, error) {
	var e []Embodied
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return e, err
	}
	err = yaml.Unmarshal(yamlFile, &e)
	return e, err
}

func saveEmbodiedYaml(e []Embodied, filepath string) error {
	d, err := yaml.Marshal(e)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, d, 0644)
	return err
}
