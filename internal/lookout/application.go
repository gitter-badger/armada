package lookout

import (
	"database/sql"
	"fmt"
	"net"
	"strings"
	"sync"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/G-Research/armada/internal/armada/authorization"
	"github.com/G-Research/armada/internal/common/grpc"
	stanUtil "github.com/G-Research/armada/internal/common/stan-util"
	"github.com/G-Research/armada/internal/common/util"
	"github.com/G-Research/armada/internal/lookout/configuration"
	"github.com/G-Research/armada/internal/lookout/events"
	"github.com/G-Research/armada/internal/lookout/repository"
	"github.com/G-Research/armada/internal/lookout/server"
	"github.com/G-Research/armada/pkg/api/lookout"
)

func StartUp(config configuration.LookoutConfiguration) (func(), *sync.WaitGroup) {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	grpcServer := grpc.CreateGrpcServer([]authorization.AuthService{&authorization.AnonymousAuthService{}})

	db, err := sql.Open("postgres", config.PostgresConnectionString)
	if err != nil {
		panic(err)
	}

	jobStore := repository.NewSQLJobStore(db)
	jobRepository := repository.NewSQLJobRepository(db)

	conn, err := stanUtil.DurableConnect(
		config.Nats.ClusterID,
		"armada-server-"+util.NewULID(),
		strings.Join(config.Nats.Servers, ","),
	)

	if err != nil {
		panic(err)
	}
	eventProcessor := events.NewEventProcessor(conn, jobStore, config.Nats.Subject, config.Nats.QueueGroup)
	eventProcessor.Start()

	lookoutServer := server.NewLookoutServer(jobRepository)
	lookout.RegisterLookoutServer(grpcServer, lookoutServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		defer log.Println("Stopping server.")

		log.Printf("Grpc listening on %d", config.GrpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		wg.Done()
	}()

	stop := func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("failed to close nats connection: %v", err)
		}
		err = db.Close()
		if err != nil {
			log.Errorf("failed to close db connection: %v", err)
		}
		grpcServer.GracefulStop()
	}

	return stop, wg
}
