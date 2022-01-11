package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func loggingSettings(logfile string) {
	file, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, file)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)

}

func main() {
	loggingSettings("test.log")
	_, err := os.Open("hogehoge")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	log.Println("logging!")
	log.Printf("%T, %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!!")

	// log.Fatalが実行された時点でコードの処理は止まる
	fmt.Println("ok?")
}
