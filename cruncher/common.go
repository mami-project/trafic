package cruncher

import (
	"fmt"
	"math"
	"time"
)

func formatFlowID(title string, cookie string, start int, sd int) string {
	// Use the auto-generated cookie if title has not been explicitly set
	if title == "" {
		title = cookie
	}

	flowID := fmt.Sprintf("%s_%d_%d", title, start, sd)

	return flowID
}

func formatTimestamp(flowStart int, sampleStart float64) string {
	tsec, tnsec := math.Modf(float64(flowStart) + sampleStart)

	return time.Unix(int64(tsec), int64(tnsec*(1e9))).String()
}
