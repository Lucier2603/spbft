package pbft

import (
l4g "code.google.com/p/log4go"
"os"
)

// default logger. usually used in dev mode when no config found
var defaultLogger l4g.Logger = l4g.NewDefaultLogger(l4g.DEBUG)

func init() {
	// overwrite its format output to print file name and code line number
	defaultLogger.AddFilter("stdout", l4g.DEBUG, l4g.NewFormatLogWriter(os.Stdout, l4g.FORMAT_DEFAULT))
}

var loggers = make(map[string]l4g.Logger)

// construct logger and save to logger map
func InitializeLogger(name, logPath string) {
	if name == "" {
		name = "default"
	}

	println(name + "=" + logPath)
	logger := make(l4g.Logger)
	logger.LoadConfiguration(logPath)
	loggers[name] = logger
}

// get the configured logger instance. if not exists(usually in dev mode), return our default debug console logger
func DefaultLogger() l4g.Logger {
	logger, ok := loggers["default"]

	if !ok {
		logger = defaultLogger
	}

	return logger
}

func HeartbeatLogger() l4g.Logger {
	return Logger("heartbeat")
}

func StorageLogger() l4g.Logger {
	return Logger("storage")
}

func ConsensusLogger() l4g.Logger {
	return Logger("consensus")
}

func ContractLogger() l4g.Logger {
	return Logger("contract")
}

// try get logger, return default logger if does not exists
func Logger(name string) l4g.Logger {

	logger, ok := loggers[name]

	// try to return the default logger
	if name == "default" || name == "" || !ok {
		logger = DefaultLogger()
	}

	return logger
}

