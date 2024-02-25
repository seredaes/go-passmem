package Logger

import (
	"fmt"
	"os"
	"seredaes/go-passmem/App/Env"
	"time"
)

var (
	Info    = Teal
	Warning = Yellow
	Danger  = Red
	Server  = Green
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func Log(message string, typeColor string) {

	debug, _ := Env.GetConfig("DEBUG")

	if debug == "enabled" {
		logConsole(message, typeColor)
	} else {
		logFile(message)
	}

}

func logFile(errorMessage string) {
	dateObj := time.Now()
	dateString := dateObj.Format("2006-01-02 15:04:05")
	logFileName := dateObj.Format("2006-01-02")

	logFile, _ := os.OpenFile("./Logs/"+logFileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	result := dateString + " :: " + errorMessage + "\n-------------------------------------------------\n"
	logFile.Write([]byte(result))

	defer logFile.Close()
}

func logConsole(errorMessage string, msgType string) {
	dateObj := time.Now()
	dateString := dateObj.Format("2006-01-02 15:04:05")
	errorMsg := "\n" + dateString + " :: " + errorMessage + "\n-------------------------------------------------"

	switch msgType {
	case "warning":
		fmt.Println(Warning(errorMsg))
	case "danger":
		fmt.Println(Danger(errorMsg))
	case "info":
		fmt.Println(Info(errorMsg))
	case "server":
		fmt.Println(Server(errorMsg))
	default:
		fmt.Println(errorMsg)
	}

}
