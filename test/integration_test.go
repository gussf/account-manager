package test

import (
	apihttp "account-manager/api/http"
	httpapi "account-manager/api/http"
	"account-manager/config"
	"account-manager/domain/service/account"
	"account-manager/domain/service/transaction"
	"account-manager/repository/postgresql"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	dbName     = "account-manager-test"
	dbUser     = "test"
	dbPassword = "test"
)

type IntegrationSuite struct {
	suite.Suite
	pgContainer *postgres.PostgresContainer
	store       *postgresql.Store
	server      *httptest.Server
}

func (s *IntegrationSuite) SetupSuite() {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx, "postgres:18.3",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	assert.NoError(s.T(), err)
	s.pgContainer = pgContainer

	host, err := pgContainer.Host(ctx)
	assert.NoError(s.T(), err)

	mappedPort, err := pgContainer.MappedPort(ctx, "5432")
	assert.NoError(s.T(), err)

	// for migration files
	err = os.Chdir("..")
	assert.NoError(s.T(), err)

	cfg := config.Config{
		HTTPServerRequestTimeout: 10 * time.Second,
		DatabaseHostname:         host,
		DatabasePort:             mappedPort.Int(),
		DatabaseName:             dbName,
		DatabaseUser:             dbUser,
		DatabasePassword:         dbPassword,
	}

	store, err := postgresql.NewStore(cfg)
	assert.NoError(s.T(), err)
	s.store = store

	accSvc := account.NewService(store)
	txSvc := transaction.NewService(store)
	srv := httpapi.NewServer(cfg, accSvc, txSvc)

	s.server = httptest.NewServer(srv.Handler())
}

func (s *IntegrationSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
	if s.pgContainer != nil {
		testcontainers.TerminateContainer(s.pgContainer)
	}
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) CreateRandomAccount() apihttp.CreateAccountResponse {
	req := httpapi.CreateAccountRequest{
		DocumentNumber: uuid.New().String(),
	}

	payload, err := json.Marshal(req)
	assert.NoError(s.T(), err)

	resp, err := http.Post(s.server.URL+"/accounts", "application/json", bytes.NewReader(payload))
	assert.NoError(s.T(), err)

	defer resp.Body.Close()
	assert.Equal(s.T(), http.StatusCreated, resp.StatusCode)

	var created apihttp.CreateAccountResponse
	err = json.NewDecoder(resp.Body).Decode(&created)
	assert.NoError(s.T(), err)

	return created
}
