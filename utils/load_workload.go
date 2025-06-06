package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"kube-scheduler/centralunit"
)

func LoadWorkloadsFromCSV(path string) ([]centralunit.Workload, error) {
    file, err := os.Open(path)
    if err != nil { return nil, err }
    defer file.Close()

    reader := csv.NewReader(file)
    rows, err := reader.ReadAll()
    if err != nil { return nil, err }

    var workloads []centralunit.Workload
    for _, row := range rows[1:] {
        cpu, _ := strconv.Atoi(row[1])
        ep, _ := strconv.ParseFloat(row[2], 64)
        workloads = append(workloads, centralunit.Workload{
            ID: row[0], CPURequirement: cpu, EnergyPriority: ep,
        })
    }
    return workloads, nil
}