package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/blazed/shorten/server"
	"github.com/blazed/shorten/storage"
	"github.com/spf13/cobra"
)

func Serve() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			if err := serve(cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
}

func serve(cmd *cobra.Command, args []string) error {
	s, err := storage.Open()
	if err != nil {
		return err
	}

	serverConfig := server.Config{
		AllowedOrigins: []string{"*"},
		Storage:        s,
	}

	srv, err := server.NewServer(serverConfig)
	if err != nil {
		return err
	}

	fmt.Println("Server starting on :3000")
	if err := http.ListenAndServe(":3000", srv); err != nil {
		fmt.Errorf("sever start failed: %s", err)
		return err
	}
	return nil
}
