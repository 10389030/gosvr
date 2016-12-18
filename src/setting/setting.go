package setting

import (
	"log"
	"os"
)

// sect global
const (
	HOST             = ":8080"
	MONGO_HOST       = "mongodb://localhost:27017"
	MONGO_ASSETS_DB  = "db_assets"
	MONGO_ASSETS_TBL = "tbl_assets"
)

// sect. Logging Basic
var _loggers = make(map[string]*log.Logger)

const LOG_INFO = log.Lmicroseconds | log.Lshortfile

var LOG_FILE = os.Stderr

func GetLogger(moduleName string) (l *log.Logger) {
	_log, ok := _loggers[moduleName]
	if !ok {
		_log = log.New(LOG_FILE, moduleName, LOG_INFO)
		_loggers[moduleName] = _log
	}

	return _log
}
