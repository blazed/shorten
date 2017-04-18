package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
		Storage: s,
	}

	srv, err := server.NewServer(serverConfig)
	if err != nil {
		return err
	}

	h := &http.Server{Addr: ":3000", Handler: srv}

	go func() {
		fmt.Println("Server starting on :3000")
		if err := h.ListenAndServe(); err != nil {
			fmt.Printf("server start failed: %s", err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c

	fmt.Println("Shutting down server...")

	h.Shutdown(context.Background())
	s.Close()

	fmt.Println("Server stopped")

	return nil
}
