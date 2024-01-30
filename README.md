# README

## Data Emissions

This intention of this repo is to collect emission data, coefficients, and machine information for GCP, AWS, and Azure datacenters. The hope is to use this data to calculate CO2e of services running on various cloud servies.

It is still a WIP, and will be updated as the progress progresses.

## v1

The v1 data can be generated via the `cmd/generator.go` functionality. The data itself is generated from the [CCF](https://github.com/cloud-carbon-footprint/ccf-coefficients/). The files are organized by cloud provider and by data information, specifically:

* use.yaml: The server architecture and wattage information
* grid.yaml: CO2e grid data by region
* embodied.yaml: Embodied data by machine type
* default.yaml: Default cloud provider values for missing use information, calculated by the averaging known data.

We found that having wattage information for the server in total, made calculating memory consumption challenging for VM instances, so we started work on a v2 for emission data.

## v2

The v2 emission were inspired by the teads dataset. It provides extensive data on machine information for AWS EC2 instances.

### AWS

https://docs.google.com/spreadsheets/d/1DqYgQnEDLQVQm5acMAhLgHLD8xXCG9BIrk-_Nv6jF3k/edit

### GCP

### Azure
