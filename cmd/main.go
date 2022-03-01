package cmd

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	log "github.com/MalukiMuthusi/logger"
	"github.com/gorilla/mux"
	"github.com/riviatechs/mt940_server/graph/generated"
	resolver "github.com/riviatechs/mt940_server/graph/resolvers"
	"github.com/riviatechs/mt940_server/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "golf",
	Short: "golf service",
	Long: `
	API for golf website
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// check for configurations that must be set
		checkMustBeSet()

		Start()

	},
}

func checkMustBeSet() {
	mustBeSet(util.ProjectID)

	mustBeSet(util.DbUser)

	mustBeSet(util.DbName)

	mustBeSet(util.DbPwd)

	mustBeSet(util.DbHost)

	mustBeSet(util.DbPort)

	mustBeSet(util.DbConnectionName)
}

func mustBeSet(env string) {
	if viper.GetString(env) == "" {
		log.Fatalf(fmt.Sprintf("%s is not set", env))
	}
}

func StartServer() {

	rootCmd.Execute()
}

func initConfig() {
	viper.SetEnvPrefix(util.ProgName)
	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize(initConfig)

	// debug flag
	rootCmd.PersistentFlags().Bool(util.Debug, false, "Enable debug messages for tracing and troubleshooting")
	bind(util.Debug)

	// project_id flag
	rootCmd.PersistentFlags().String(util.ProjectID, "", "Google Cloud Project ID")
	bind(util.ProjectID)

	// wait flag
	rootCmd.PersistentFlags().Int64(util.Wait, int64(time.Second*15), "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	bind(util.Wait)

	// port flag
	rootCmd.PersistentFlags().String(util.Port, "", "port the server to listen at")
	bind(util.Port)

	// db user name
	rootCmd.PersistentFlags().String(util.DbUser, "", "Database user name")
	bind(util.DbUser)

	// db name
	rootCmd.PersistentFlags().String(util.DbName, "", "Database Name")
	bind(util.DbName)

	// db password
	rootCmd.PersistentFlags().String(util.DbPwd, "", "Database password")
	bind(util.DbPwd)

	// db host
	rootCmd.PersistentFlags().String(util.DbHost, "", "Database host location")
	bind(util.DbHost)

	// db Port
	rootCmd.PersistentFlags().String(util.DbPort, "", "Database Port")
	bind(util.DbPort)

	// db connection name
	rootCmd.PersistentFlags().String(util.DbConnectionName, "", "Database Connection name")

	// db hosted on cloud?
	rootCmd.PersistentFlags().Bool(util.DbHostedCloud, false, "Database hosted on cloud")

	bind(util.DbPort)

	// Set up the structured logging library
	log.Setup(viper.GetBool(util.Debug))

}

// bind viper flags to the cobra command
func bind(flag string) {
	viper.BindPFlag(flag, rootCmd.Flags().Lookup(flag))
}

func Start() {
	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	graphSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	r := mux.NewRouter()
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", graphSrv)

	port := os.Getenv("PORT")
	if port == "" {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("port not set, setting it to 8080")

		port = "8080"
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {

		log.Info(fmt.Sprintf("Starting server %s", srv.Addr))

		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed tp start server", zap.Error(err))
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	log.Info("shutting down")
	os.Exit(0)
}
