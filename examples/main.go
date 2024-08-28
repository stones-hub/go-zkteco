package main

import (
	gozk "github.com/stones-hub/go-zkteco"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	zkSocket := gozk.NewZK("192.168.205.55", 4370, 0, gozk.DefaultTimezone)

	zkSocket.Connect()

	defer zkSocket.Disconnect()

	/*
		atts, err := zkSocket.GetAttendances()
		if err != nil {
			fmt.Printf("err : %v\n", err)
		}

		for _, att := range atts {
			fmt.Printf("%v ==> %d\n", att.AttendedAt, att.UserID)
		}
	*/

	zkSocket.GetUsersByK3()
}

func gracefulQuit(f func()) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan

		log.Println("Stopping...")
		f()

		time.Sleep(time.Second * 1)
		os.Exit(1)
	}()

	for {
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}
