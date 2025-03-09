package vnstat

import (
	"os/exec"
)

type UsageData struct {
	InUsage    int64
	OutUsage   int64
	TotalUsage int64
}

// GetUsage executes the vnStat command and returns sample usage data.
func GetUsage() (UsageData, error) {
	var data UsageData
	// Run the vnstat command with JSON output (assuming vnstat is installed)
	_, err := exec.Command("vnstat", "--json").Output()
	if err != nil {
		// Return an error if the command fails
		return data, err
	}
	// TODO: Parse the JSON output properly.
	// For now, return some placeholder values.
	data.InUsage = 1000  // Example placeholder
	data.OutUsage = 1500 // Example placeholder
	data.TotalUsage = data.InUsage + data.OutUsage
	return data, nil
}
