// THIS IS A WORK IN PROGRESS
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

func main() {
	// generateEmbodied()
	//	generateUSE()
	generateGrid()
}

// type embodied struct {
// 	Type              string
// 	AdditionalMemory  float64
// 	AdditionalStorage float64
// 	AdditionalCPUS    float64
// 	AdditionalGPUS    float64
// 	Total             float64
// 	Architecture      string
// }

// func generateEmbodied() {
// 	file, err := os.Open("../ccf-coefficients/output/coefficients-aws-embodied.csv")
// 	if err != nil {
// 		log.Fatal("Error while reading the file", err)
// 	}
//
// 	defer file.Close()
//
// 	reader := csv.NewReader(file)
//
// 	records, err := reader.ReadAll()
//
// 	if err != nil {
// 		fmt.Println("Error reading records")
// 	}
//
// 	var data []embodied
// 	//NOTE: different providers have differnt mappings
// 	for _, record := range records {
// 		data = append(data, embodied{
// 			Type:              record[1],
// 			AdditionalMemory:  parseStringToFloat(record[2]),
// 			AdditionalStorage: parseStringToFloat(record[3]),
// 			AdditionalCPUS:    parseStringToFloat(record[4]),
// 			AdditionalGPUS:    parseStringToFloat(record[5]),
// 			Total:             parseStringToFloat(record[6]),
// 		})
// 	}
// 	yamlData, _ := yaml.Marshal(&data)
//
// 	fileName := "data/aws-embodied.yaml"
// 	err = ioutil.WriteFile(fileName, yamlData, 0644)
// 	if err != nil {
// 		panic("Unable to write data into the file")
// 	}
//
// }

func parseStringToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// type architecture struct {
// 	Architecture string
// 	MinWatts     float64
// 	MaxWatts     float64
// 	Chip         float64
// }
//
// func generateUSE() {
// 	file, err := os.Open("../ccf-coefficients/output/coefficients-azure-use.csv")
// 	if err != nil {
// 		log.Fatal("Error while reading the file", err)
// 	}
//
// 	defer file.Close()
//
// 	reader := csv.NewReader(file)
//
// 	records, err := reader.ReadAll()
//
// 	if err != nil {
// 		fmt.Println("Error reading records")
// 	}
//
// 	var data []architecture
// 	//NOTE: different providers have differnt mappings
// 	for _, record := range records {
// 		data = append(data, architecture{
// 			Architecture: record[1],
// 			MinWatts:     parseStringToFloat(record[2]),
// 			MaxWatts:     parseStringToFloat(record[3]),
// 			Chip:         parseStringToFloat(record[4]),
// 		})
// 	}
// 	yamlData, _ := yaml.Marshal(&data)
//
// 	fileName := "data/azure-use.yaml"
// 	err = ioutil.WriteFile(fileName, yamlData, 0644)
// 	if err != nil {
// 		panic("Unable to write data into the file")
// 	}
//
// }

type grid struct {
	Region string
	CO2e   float64
}

func generateGrid() {
	file, err := os.Open("../ccf-coefficients/data/grid-emissions-factors-gcp.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	var data []grid
	//NOTE: different providers have differnt mappings
	for _, record := range records {
		fmt.Println(record)
		data = append(data, grid{
			Region: record[0],
			CO2e:   parseStringToFloat(record[2]),
		})
	}
	yamlData, _ := yaml.Marshal(&data)

	fileName := "data/gcp-grid.yaml"
	err = os.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}

}
