package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	// max_user_connection
	if variableMap[dbConfigMaxUserConnection] != dbConfigMaxUserConnectionValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigMaxUserConnection,
			variableMap[dbConfigMaxUserConnection],
		})
		advice = "It is recommended that " + dbConfigMaxUserConnection + " be set to " + dbConfigMaxUserConnectionValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// log_bin
	if variableMap[dbConfigLogBin] != dbConfigLogBinValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigLogBin,
			variableMap[dbConfigLogBin],
		})
		advice = "It is recommended that " + dbConfigLogBin + " be set to " + dbConfigLogBinValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// binlog_format
	if variableMap[dbConfigBinlogFormat] != dbConfigBinlogFormatValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigBinlogFormat,
			variableMap[dbConfigBinlogFormat],
		})
		advice = "It is recommended that " + dbConfigBinlogFormat + " be set to " + dbConfigBinlogFormatValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// binlog_row_image
	if variableMap[dbConfigBinlogRowImage] != dbConfigBinlogRowImageValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigBinlogRowImage,
			variableMap[dbConfigBinlogRowImage],
		})
		advice = "It is recommended that " + dbConfigBinlogRowImage + " be set to " + dbConfigBinlogRowImageValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// sync_binlog
	if variableMap[dbConfigSyncBinlog] != dbConfigSyncBinlogValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSyncBinlog,
			variableMap[dbConfigSyncBinlog],
		})
		advice = "It is recommended that " + dbConfigSyncBinlog + " be set to " + dbConfigSyncBinlogValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// innodb_flush_log_at_trx_commit
	if variableMap[dbConfigInnodbFlushLogAtTrxCommit] != dbConfigInnodbFlushLogAtTrxCommitValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbFlushLogAtTrxCommit,
			variableMap[dbConfigInnodbFlushLogAtTrxCommit],
		})
		advice = "It is recommended that " + dbConfigInnodbFlushLogAtTrxCommit + " be set to " + dbConfigInnodbFlushLogAtTrxCommitValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// gtid_mode
	if variableMap[dbConfigGtidMode] != dbConfigGtidModeValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigGtidMode,
			variableMap[dbConfigGtidMode],
		})
		advice = "It is recommended that " + dbConfigGtidMode + " be set to " + dbConfigGtidModeValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// enforce_gtid_consistency
	if variableMap[dbConfigEnforceGtidConsistency] != dbConfigEnforceGtidConsistencyValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigEnforceGtidConsistency,
			variableMap[dbConfigEnforceGtidConsistency],
		})
		advice = "It is recommended that " + dbConfigEnforceGtidConsistency + " be set to " + dbConfigEnforceGtidConsistencyValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// slave-parallel-type
	if variableMap[dbConfigSlaveParallelType] != dbConfigSlaveParallelTypeValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSlaveParallelType,
			variableMap[dbConfigSlaveParallelType],
		})
		advice = "It is recommended that " + dbConfigSlaveParallelType + " be set to " + dbConfigSlaveParallelTypeValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// slave-parallel-workers
	if variableMap[dbConfigSlaveParallelWorkers] != dbConfigSlaveParallelWorkersValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSlaveParallelWorkers,
			variableMap[dbConfigSlaveParallelWorkers],
		})
		advice = "It is recommended that " + dbConfigSlaveParallelWorkers + " be set to " + dbConfigSlaveParallelWorkersValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// master_info_repository
	if variableMap[dbConfigMasterInfoRepository] != dbConfigMasterInfoRepositoryValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigMasterInfoRepository,
			variableMap[dbConfigMasterInfoRepository],
		})
		advice = "It is recommended that " + dbConfigMasterInfoRepository + " be set to " + dbConfigMasterInfoRepositoryValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// relay_log_info_repository
	if variableMap[dbConfigRelayLogInfoRepository] != dbConfigRelayLogInfoRepositoryValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigRelayLogInfoRepository,
			variableMap[dbConfigRelayLogInfoRepository],
		})
		advice = "It is recommended that " + dbConfigRelayLogInfoRepository + " be set to " + dbConfigRelayLogInfoRepositoryValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// report_host
	serverName := de.operationInfo.MySQLServer.GetServerName()
	if variableMap[dbConfigReportHost] != serverName {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigReportHost,
			variableMap[dbConfigReportHost],
		})
		advice = "It is recommended that " + dbConfigReportHost + " be set to " + serverName + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// report_port
	portNum := de.operationInfo.MySQLServer.GetPortNum()
	if variableMap[dbConfigReportPort] != strconv.Itoa(portNum) {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigReportPort,
			variableMap[dbConfigReportPort],
		})
		advice = "It is recommended that " + dbConfigReportPort + " be set to " + strconv.Itoa(portNum) + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}

	// innodb_flush_method
	if variableMap[dbConfigInnodbFlushMethod] != dbConfigInnodbFlushMethodValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbFlushMethod,
			variableMap[dbConfigInnodbFlushMethod],
		})
		advice = "It is recommended that " + dbConfigInnodbFlushMethod + " be set to " + dbConfigInnodbFlushMethodValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// innodb_monitor_enable
	if variableMap[dbConfigInnodbMonitorEnable] != dbConfigInnodbMonitorEnableValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbMonitorEnable,
			variableMap[dbConfigInnodbMonitorEnable],
		})
		advice = "It is recommended that " + dbConfigInnodbMonitorEnable + " be set to " + dbConfigInnodbMonitorEnableValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// innodb_print_all_deadlocks
	if variableMap[dbConfigInnodbPrintAllDeadlocks] != dbConfigInnodbPrintAllDeadlocksValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigInnodbPrintAllDeadlocks,
			variableMap[dbConfigInnodbPrintAllDeadlocks],
		})
		advice = "It is recommended that " + dbConfigInnodbPrintAllDeadlocks + " be set to " + dbConfigInnodbPrintAllDeadlocksValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// slow_query_log
	if variableMap[dbConfigSlowQueryLog] != dbConfigSlowQueryLogValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigSlowQueryLog,
			variableMap[dbConfigSlowQueryLog],
		})
		advice = "It is recommended that " + dbConfigSlowQueryLog + " be set to " + dbConfigSlowQueryLogValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// performance_schema
	if variableMap[dbConfigPerformanceSchema] != dbConfigPerformanceSchemaValid {
		dbConfigCount++
		dbConfigInvalid = append(dbConfigInvalid, GlobalVariables{
			dbConfigPerformanceSchema,
			variableMap[dbConfigPerformanceSchema],
		})
		advice = "It is recommended that " + dbConfigPerformanceSchema + " be set to " + dbConfigPerformanceSchemaValid + "." + "\n"
		dbConfigAdvice = append(dbConfigAdvice, advice)
	}
	// database config data
	jsonBytesTotal, err := json.Marshal(dbConfigInvalid)
	if err != nil {
		return nil
	}
	de.result.DBConfigData = string(jsonBytesTotal)
	// database config advice
	jsonBytesAdvice, err := json.Marshal(dbConfigAdvice)
	if err != nil {
		return nil
	}
	de.result.DBConfigAdvice = string(jsonBytesAdvice)

	// database config score deduction
	dbConfigScoreDeduction := float64(dbConfigCount) * dbConfigConfig.ScoreDeductionPerUnitHigh
	if dbConfigScoreDeduction > dbConfigConfig.MaxScoreDeductionHigh {
		dbConfigScoreDeduction = dbConfigConfig.MaxScoreDeductionHigh
	}
	de.result.DBConfigScore = int(defaultMaxScore - dbConfigScoreDeduction)
	if de.result.DBConfigScore < constant.ZeroInt {
		de.result.DBConfigScore = constant.ZeroInt
	}
}
