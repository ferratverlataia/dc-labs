package main

import (

	"context"
	"crypto/rand"
	"flag"
	"fmt"
	pb "github.com/ferratverlataia/dc-labs/challenges/third-partial/proto"

	"google.golang.org/grpc"
	"log"
	"math/big"
	"nanomsg.org/go/mangos/v2/protocol/respondent"
	"net"
	"os"
	"strconv"
	"time"

	// register transports
	_ "nanomsg.org/go/mangos/v2/transport/all"
)
import mangos "nanomsg.org/go/mangos/v2"
var (
	defaultRPCPort = 50051
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

var (
	controllerAddress = ""
	workerName        = ""
	tags              = ""
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("RPC: Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func init() {
	flag.StringVar(&controllerAddress, "controller", "tcp://localhost:40899", "Controller address")
	flag.StringVar(&workerName, "worker-name", "hard-worker", "Worker Name")
	flag.StringVar(&tags, "tags", "gpu,superCPU,largeMemory", "Comma-separated worker tags")
}

// joinCluster is meant to join the controller message-passing server
func joinCluster() {
	var sock mangos.Socket
	var err error
	var msg []byte
	var name="1"
	if sock, err = respondent.NewSocket(); err != nil {
		die("can't get new sub socket: %s", err.Error())
	}
	 err=sock.Dial(controllerAddress)
	if err!=nil {

		die("cant dial on respondent socket %s",err.Error())

	}
	log.Printf("Connecting to controller on: %s", controllerAddress)

	for {
		log.Printf("sending information petition")
		if msg, err = sock.Recv(); err != nil {
			die("Cannot recv: %s", err.Error())
		}
		fmt.Printf("CLIENT(%s): RECEIVED \"%s\" SURVEY REQUEST\n",
			name, string(msg))
		port := getAvailablePort()
		fmt.Printf("CLIENT(%s): SENDING DATE SURVEY RESPONSE\n", name)
		t := time.Now().Format(time.ANSIC)
		usage, err := rand.Int(rand.Reader, big.NewInt(100))

		if err != nil {
			panic(err)
		}
		conn, err:= net.Dial("udp","8.8.8:80")
		IP:=""
		if(err!=nil){
			IP="not found"

		}else{
			IP=conn.LocalAddr().(*net.UDPAddr).String()
		}

		workerMetadata := workerName + "*" + tags + "*" + IP + "*" + strconv.Itoa(port) + "*" + t + "*" + usage.String()
		fmt.Printf("Worker meta data: %s",workerMetadata)
		if err = sock.Send([]byte(workerMetadata)); err != nil {
			die("Cannot send: %s", err.Error())
		}
	}


}

func getAvailablePort() int {
	port := defaultRPCPort
	for {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			port = port + 1
			continue
		}
		ln.Close()
		break
	}
	return port
}

func main() {
	flag.Parse()

	// Subscribe to Controller
	go joinCluster()

	// Setup Worker RPC Server
	rpcPort := getAvailablePort()
	log.Printf("Starting RPC Service on localhost:%v", rpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
