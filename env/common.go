package env

import (
	"os"
	"strconv"
	"time"
)

func envi(name string, defaultValue int) int {
	if i, e := strconv.Atoi(os.Getenv(name)); e != nil {
		return defaultValue
	} else {
		return i
	}
}

const DefaultServerURL = "https://www.murphysec.com"

var ServerURLOverride = os.Getenv("MPS_CLI_SERVER")
var APITokenOverride = os.Getenv("API_TOKEN")
var ScannerScan = false
var ScannerShouldEnableMavenBackupScan = false
var CommandTimeout time.Duration
var NoWait bool

func init() {
	ctm := os.Getenv("COMMAND_TIMEOUT")
	ct, e := strconv.Atoi(ctm)
	if e != nil || ctm == "" {
		CommandTimeout = time.Second * 25
	} else {
		CommandTimeout = time.Second * time.Duration(ct)
	}
}
