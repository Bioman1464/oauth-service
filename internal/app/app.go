package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	httpserver "auth-service/internal/adapters/http-server"
	"auth-service/internal/providers/repo"
	"auth-service/internal/providers/service"

	_ "github.com/lib/pq"
)

type Config struct {
	PublicURL string
	DBCfg     DatabaseConfig
}

type DatabaseConfig struct {
	URL     string
	Driver  string
	User    string
	Pass    string
	Host    string
	Port    string
	Name    string
	SSLMode string
}

type Application struct {
	cfg          *Config
	log          *log.Logger
	postgresqlDB *sql.DB
	httpServer   *httpserver.Server
}

var app Application

func Start(ctx context.Context) {
	app = Application{}

	fmt.Println("Parsing config")
	cfg, err := parseConfig()
	if err != nil {
		log.WithError(err).Fatal("Config parse failed")
	}

	app.cfg = cfg
	ctx = context.WithValue(ctx, "app_url", app.cfg.PublicURL)

	fmt.Println("Initializing logger")
	logger, err := app.logger()
	if err != nil {
		log.WithError(err).Fatal("Logger initialization failed")
	}

	app.log = logger

	defer func() {
		if r := recover(); r != nil {
			logger.
				WithField("RecoverMessage", r).
				Fatal("Application Crashed")
		}
	}()

	app.log.Info("Initializing db connection")
	db, err := app.database()
	if err != nil {
		app.log.WithError(err).Fatal("Database initialization failed")
	}

	app.postgresqlDB = db

	repoProvider := app.repositoryProvider()

	serviceProvider := app.serviceProvider(repoProvider)

	httpServer, err := app.initHTTPServer(ctx, serviceProvider)
	if err != nil {
		app.log.WithError(err).Fatal("HTTP Server initialization failed")
	}

	app.log.Info("Starting HTTP server")
	err = httpServer.Start()
	if err != nil {
		app.log.WithError(err).Fatal("HTTP Server start failed")
	}

	//TODO:: graceful shutdown
}

func Stop(ctx context.Context) {
}

func parseConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := Config{
		PublicURL: os.Getenv("PUBLIC_URL"),
		DBCfg: DatabaseConfig{
			URL:     os.Getenv("DATABASE_URL"),
			Driver:  os.Getenv("DATABASE_DRIVER"),
			Host:    os.Getenv("DATABASE_HOST"),
			Port:    os.Getenv("DATABASE_PORT"),
			User:    os.Getenv("DATABASE_USER"),
			Pass:    os.Getenv("DATABASE_PASS"),
			Name:    os.Getenv("DATABASE_NAME"),
			SSLMode: os.Getenv("DATABASE_SSL_MODE"),
		},
	}

	return &cfg, nil
}

func (a *Application) logger() (*log.Logger, error) {
	return log.New(), nil
}

func (a *Application) database() (*sql.DB, error) {
	dbCfg := a.cfg.DBCfg

	a.log.Println(dbCfg)

	dbUrl := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Pass, dbCfg.Name, dbCfg.SSLMode,
	)

	db, err := sql.Open(a.cfg.DBCfg.Driver, dbUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (a *Application) repositoryProvider() repo.RepositoryProvider {
	return repo.NewPostgresqlRepositoryProvider(a.postgresqlDB)
}

func (a *Application) serviceProvider(provider repo.RepositoryProvider) service.Provider {
	return service.NewServiceProvider(provider)
}

func (a *Application) initHTTPServer(ctx context.Context, provider service.Provider) (*httpserver.Server, error) {
	return httpserver.NewServer(ctx, a.cfg.PublicURL, provider, a.log)
}
