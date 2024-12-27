package services

import (
	"encoding/json"
	"os"
	"repair-service/models"
)

// GetRepairs fetches all repairs from the JSON file
func GetRepairs() ([]models.Repair, error) {
	file, err := os.Open("services/repairs.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var repairs []models.Repair
	if err := json.NewDecoder(file).Decode(&repairs); err != nil {
		return nil, err
	}
	return repairs, nil
}

// GetRepairsByMachineID fetches repairs filtered by machine ID
func GetRepairsByMachineID(machineID string) ([]models.Repair, error) {
	repairs, err := GetRepairs()
	if err != nil {
		return nil, err
	}

	// Filter repairs based on the machineID
	var filteredRepairs []models.Repair
	for _, repair := range repairs {
		if repair.MachineID == machineID {
			filteredRepairs = append(filteredRepairs, repair)
		}
	}
	return filteredRepairs, nil
}
