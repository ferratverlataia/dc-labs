package controller

import (
	"fmt"
	"log"
	mangos "nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/surveyor"
	"os"
	"time"
	// register transports
	_ "nanomsg.org/go/mangos/v2/transport/all"
)

var controllerAddress = "tcp://localhost:40899"

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func Start() {
	var sock mangos.Socket
	var err error
	var msg []byte

	if sock, err = surveyor.NewSocket(); err != nil {
		die("can't get new pub socket: %s", err)
	}
	if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on pub socket: %s", err.Error())
	}
	err= sock.SetOption(mangos.OptionSurveyTime,time.Second/2)
	if err!=nil{

		die("SetOption(): %s",err.Error())

	}
	for {
	time.Sleep(time.Second)
		d := date()
		log.Printf("Controller: Publishing Date %s\n", d)
		if err = sock.Send([]byte("DATE")); err != nil {
			die("Failed publishing: %s", err.Error())
		}
		time.Sleep(time.Second * 3)

			fmt.Printf("receiving data")
			if msg,err =sock.Recv(); err != nil{
				fmt.Printf("Cannot recv: %s", err.Error())
				continue
			}
	fmt.Printf("Server Succesfully received response *%s*",string(msg) )

	}
}
