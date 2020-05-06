package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	mangos "nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/surveyor"
	"os"
	"strings"
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

	type employee struct {
		worker ,tags,IP,port,usage string

	}

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
	for{

		fmt.Printf("receiving data \n")
		if msg,err =sock.Recv(); err != nil{
			fmt.Printf("Cannot recv: %s\n", err.Error())
			break
		}

		fmt.Printf("Server Succesfully received response:  %s \n",string(msg) )
		 rpcmachine := strings.Split(string(msg),"*")

			 file,_:=json.MarshalIndent( employee{worker:rpcmachine[0],tags:rpcmachine[1],IP:rpcmachine[2],port:rpcmachine[3],usage:rpcmachine[4]},"","")

		 _=ioutil.WriteFile("test.json",file,0644)
	}


	}
}
