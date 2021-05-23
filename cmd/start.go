package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
	"github.com/terra-project/amino-decoder/api"

	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"
)

var (
	// Port to run server
	Port int

	// The actual app config
	server *api.Server
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println(fmt.Sprintf("Listening on port :%v... [or whatever you remapped to for docker]", server.Port))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", server.Port), handlers.LoggingHandler(os.Stdout, server.Router())))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	serveCmd.PersistentFlags().IntVar(&Port, "port", 3000, "listen port")

	rootCmd.AddCommand(serveCmd)
}

func initConfig() {
	server = &api.Server{
		Port:    Port,
		Version: Version,
		Commit:  Commit,
		Branch:  Branch,
	}

	viper.AutomaticEnv() // read in environment variables that match
}
