package Env

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// Загрузка ENV файла и чтение конфигов

var config map[string]string = make(map[string]string)

func GetConfig(key string) (string, bool) {
	value, exist := config[key]
	return value, exist
}

func readConfig(line string) {
	if line != "" && !strings.HasPrefix(line, "#") {
		lineArr := strings.Split(line, "=")
		config[lineArr[0]] = strings.Trim(lineArr[1], "\"")
	}
}

func LoadEnv() {
	file, err := os.Open("./.env")
	if err != nil {
		panic("Error read config")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")

		if err == io.EOF {
			readConfig(line)
			break
		} else if err == nil {
			readConfig(line)
		} else {
			panic("Error read string")
		}
	}

	checkEnv()
}

func checkEnv() {
	mandatoryKeys := []string{"SERVER", "ENCRYPT_KEY", "TOKEN_KEY"}

	for _, key := range mandatoryKeys {
		_, exist := GetConfig(key)
		if !exist {
			log.Fatalf("ENVIROMENT BROKEN! KEY: %s is absent", key)
		}
	}
}
