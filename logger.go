package main

import "log"

func init() {
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
}

func Info(str string) {
	log.Println(str)
}

func Warn(str string) {
	log.Fatalln(str)
}

func Error(str string) {
	log.Panicln(str)
}
