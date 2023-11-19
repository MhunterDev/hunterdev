package logsalot

import (
	"fmt"
	"log"
	"os"
)

func LogProcessToDebug(event string) {

	path := "/usr/mhdev/logs/Base-processes.log"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "Process Log:", log.Lmsgprefix)

	logger.SetOutput(file)
	logger.Println(event)
}

func DbErr(err error) {
	path := "/usr/mhdev/logs/pgrunner.log"

	file, _ := os.OpenFile(path, os.O_APPEND, 0644)
	defer file.Close()

	log.SetFlags(log.Ldate | log.Lshortfile)
	log.SetOutput(file)
	log.SetPrefix("<<< Database >>>   ")
	message := fmt.Sprintf("Error : %s ", err)
	log.Println(message)
	fmt.Println(err)
}

func ApiErr(err error) {
	path := "/usr/mhdev/logs/pgrunner.log"

	file, _ := os.OpenFile(path, os.O_APPEND, 0644)
	defer file.Close()

	log.SetFlags(log.Ldate | log.Lshortfile)
	log.SetOutput(file)
	log.SetPrefix("<<<   API    >>>   ")
	message := fmt.Sprintf("Error : %s ", err)
	log.Println(message)
	fmt.Println(err)
}

func LogInit(event string) {
	path := "/usr/mhdev/logs/Install.log"

	file, _ := os.OpenFile(path, os.O_APPEND, 0644)
	defer file.Close()

	log.SetFlags(log.Ldate | log.Lshortfile)
	log.SetOutput(file)
	log.SetPrefix("<<< Init File System >>>   ")
	message := fmt.Sprintf("Creating: %s ", event)

	log.Println(message)
	fmt.Printf("Creating: %s \n", event)
}
