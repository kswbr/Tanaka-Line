package main
import (
    "log"
    "net"
pb "./pb"
    "./service"
    "google.golang.org/grpc"
)
func main() {
    listenPort, err := net.Listen("tcp", ":19003")
    if err != nil {
        log.Fatalln(err)
    }
    server := grpc.NewServer()
    tanakaService := &service.MyTanakaService{}
    // 実行したい実処理をseverに登録する
    pb.RegisterTanakaServer(server, tanakaService)
    server.Serve(listenPort)
}

