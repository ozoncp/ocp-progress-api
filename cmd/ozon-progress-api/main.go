package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/enescakir/emoji"
	"github.com/ozoncp/ocp-progress-api/core/api"
	desc "github.com/ozoncp/ocp-progress-api/pkg/ocp-progress-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort              = 8082
	mainConfigName string = "config/config.yaml"
)

func startGrpc(port int) error {
	address := ":" + fmt.Sprint(port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return fmt.Errorf("failed to start listening: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	api := api.NewOcpProgressApi()
	desc.RegisterOcpProgressApiServer(server, api)

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve server: %v", err)
	}
	return nil
}

func main() {

	fmt.Printf("Hello, my name is Dima Larin. I`ll work on progress-api %v \n", emoji.WavingHand.Tone(emoji.Dark))

	if err := startGrpc(grpcPort); err != nil {
		log.Fatal().Msgf("failed to start GRPC server: %v", err)
	}

	_ = openAndCloseFile(5)
}

func openAndCloseFile(count int) error {

	var f error
	for i := 0; i < count; i++ {

		f = func() error {
			file, err := os.Open(mainConfigName)
			if err != nil {
				fmt.Println(err)
				return err
			}

			defer func() {
				fmt.Println("File closed")
				file.Close()
			}()

			in := bufio.NewReader(file)
			str, err := in.ReadString('\n')

			if errors.Is(err, io.EOF) {
				buf := make([]byte, in.Buffered())
				read, err := io.ReadFull(in, buf)
				if err != nil {
					fmt.Println(err)
					return err
				}
				str += string(buf[:read])
			} else if err != nil {
				fmt.Println("Error of read string from file ", err)
				return err
			}

			fmt.Println(str)
			time.Sleep(2 * time.Second)

			return nil
		}()
		if f != nil {
			break
		}
	}
	return f
}
