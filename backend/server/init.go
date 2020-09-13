package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mashun4ek/study_group/app/config"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

type Env struct {
	*config.Env
}

// NewRouter - should I leave it in server/init.go
func NewRouter(e *config.Env) *mux.Router {
	routes := buildRoutes(&Env{e})
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/app").Subrouter()

	for _, route := range routes {
		s.HandleFunc(route.Path, route.Handler).Name(route.Description).Methods(route.Method)
	}
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			log.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			log.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			log.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			log.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			log.Println("Methods:", strings.Join(methods, ","))
		}
		log.Println()
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return r
}

// StartServer to start
func StartServer() {
	if configVars, err := loadConfigVariables(); err != nil {
		log.Println("Couldn't load configurations for db")
	} else if appConfig, err := config.NewEnv(configVars); err != nil {
		log.Println("Couldn't create new environment")
	} else {
		router := NewRouter(appConfig)
		corsHandler := cors.New(cors.Options{
			AllowCredentials: true,
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			ExposedHeaders:   []string{"Authorization", "Content-Disposition"},
			Debug:            true,
		})
		handler := corsHandler.Handler(router)
		log.Fatal(http.ListenAndServe(":8080", handler))
	}
}

func loadConfigVariables() (*config.ConfigurationVars, error) {
	app := cli.NewApp()
	app.Name = "REST API"
	app.Usage = ""

	var c config.ConfigurationVars

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "is-prod",
			Usage:       "Set true if a production deployment; else false.Default:false",
			EnvVar:      "IS_PROD",
			Destination: &c.IsProd,
		},
		cli.StringFlag{
			Name:        "db-user",
			Usage:       "Database Username",
			EnvVar:      "DB_USER",
			Value:       "",
			Destination: &c.DBUser,
		},
		cli.StringFlag{
			Name:        "db-pass",
			Usage:       "Database Username",
			EnvVar:      "DB_PASS",
			Value:       "",
			Destination: &c.DBPass,
		},
		cli.StringFlag{
			Name:        "db-host",
			Usage:       "Database Username",
			EnvVar:      "DB_HOST",
			Value:       "",
			Destination: &c.DBHost,
		},
		cli.IntFlag{
			Name:        "db-port",
			Usage:       "Database Username",
			EnvVar:      "DB_PORT",
			Value:       5432,
			Destination: &c.DBPort,
		},
		cli.StringFlag{
			Name:        "db-name",
			Usage:       "Database Username",
			EnvVar:      "DB_NAME",
			Value:       "",
			Destination: &c.DBName,
		},
	}

	if err := app.Run(os.Args); err != nil {
		return nil, errors.Wrap(err, "configuration variables failed to load from the host environment")
	}
	return &c, nil

}
