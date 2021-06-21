package main

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/enescakir/emoji"
	"github.com/ozoncp/ocp-progress-api/core/api"
	"github.com/ozoncp/ocp-progress-api/core/repo"
	"github.com/ozoncp/ocp-progress-api/internal/producer"
	desc "github.com/ozoncp/ocp-progress-api/pkg/ocp-progress-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort              = 8082
	mainConfigName string = "config/config.yaml"
	kafkaBroker    string = `env:"KAFKA_BROKER" envDefault:"127.0.0.1:9094"`
)

func getProgressRepo() *repo.Repo {
	const dbName = "ozon"
	const address = "postgres://postgres:postgres@localhost:5432/" + dbName + "?sslmode=disable"

	db, err := sql.Open("pgx", address)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open postgres")
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping postgres")
	}

	log.Debug().Msgf("Connected to DB %v", dbName)

	progressRepo := repo.New(db)

	return &progressRepo
}

func startGrpc(port int) error {
	address := ":" + fmt.Sprint(port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return fmt.Errorf("failed to start listening: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	ctx := context.Background()

	LogProducer, err := producer.New(ctx, []string{kafkaBroker}, "events", 128)
	if err != nil {
		return fmt.Errorf("failed to start LogProducer: %v", err)
	}

	api := api.NewOcpProgressApi(*getProgressRepo(), LogProducer)
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
