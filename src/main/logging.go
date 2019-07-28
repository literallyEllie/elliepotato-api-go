package main

import "log"

// Simpe logging methods and may be expanded on later.

func LogInfo(info interface{}) {
	log.Println("[INFO]", info)
}

func LogWarn(warning interface{}) {
	log.Println("[WARN]", warning)
}

func LogErr(error interface{}) {
	log.Println("[ERR]", error)
}
