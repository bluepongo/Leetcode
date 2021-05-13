package leetcode

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// checkCPUUsage checks cpu usage
func (de *DefaultEngine) checkCPUUsage() error {
	// get data
	serverName := de.operationInfo.MySQLServer.GetServerName()
	query := fmt.Sprintf(`
		sum(avg by (node_name,mode) (clamp_max(((avg by (mode,node_name) ((
		clamp_max(rate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[20s]),1)) or
		(clamp_max(irate(node_cpu_seconds_total{node_name=~"%s",mode!="idle"}[5m]),1)) ))*100 or
		(avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[20s]) or
		avg_over_time(node_cpu_average{node_name=~"%s", mode!="total", mode!="idle"}[5m]))),100)))
	`, serverName, serverName, serverName, serverName)
	result, err := de.monitorPrometheusConn.Execute(query, de.operationInfo.StartTime, de.operationInfo.EndTime, de.operationInfo.Step)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	cpuUsageConfig := de.getItemConfig(defaultCPUUsageItemName)

	var (
		cpuUsage            float64
		cpuUsageHighSum     float64
		cpuUsageHighCount   int
		cpuUsageMediumSum   float64
		cpuUsageMediumCount int

		cpuUsageHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		cpuUsage, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case cpuUsage >= cpuUsageConfig.HighWatermark:
			cpuUsageHigh = append(cpuUsageHigh, rowData)

			cpuUsageHighSum += cpuUsage
			cpuUsageHighCount++
		case cpuUsage >= cpuUsageConfig.LowWatermark:
			cpuUsageMediumSum += cpuUsage
			cpuUsageMediumCount++
		}
	}

	// cpu usage data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.CPUUsageData = string(jsonBytesTotal)
	// cpu usage high
	jsonBytesHigh, err := json.Marshal(cpuUsageHigh)
	if err != nil {
		return nil
	}
	de.result.CPUUsageHigh = string(jsonBytesHigh)

	// cpu usage high score deduction
	cpuUsageScoreDeductionHigh := (cpuUsageHighSum/float64(cpuUsageHighCount) - cpuUsageConfig.HighWatermark) / cpuUsageConfig.Unit * cpuUsageConfig.ScoreDeductionPerUnitHigh
	if cpuUsageScoreDeductionHigh > cpuUsageConfig.MaxScoreDeductionHigh {
		cpuUsageScoreDeductionHigh = cpuUsageConfig.MaxScoreDeductionHigh
	}
	// cpu usage medium score deduction
	cpuUsageScoreDeductionMedium := (cpuUsageMediumSum/float64(cpuUsageMediumCount) - cpuUsageConfig.LowWatermark) / cpuUsageConfig.Unit * cpuUsageConfig.ScoreDeductionPerUnitMedium
	if cpuUsageScoreDeductionMedium > cpuUsageConfig.MaxScoreDeductionMedium {
		cpuUsageScoreDeductionMedium = cpuUsageConfig.MaxScoreDeductionMedium
	}
	// cpu usage score
	de.result.CPUUsageScore = int(defaultMaxScore - cpuUsageScoreDeductionHigh - cpuUsageScoreDeductionMedium)
	if de.result.CPUUsageScore < constant.ZeroInt {
		de.result.CPUUsageScore = constant.ZeroInt
	}

	return nil
}
