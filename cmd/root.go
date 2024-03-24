package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	Data    bool
	rootCmd = &cobra.Command{
		Use:   "gurl",
		Short: "curl but better",
		Run:   run,
		Args:  cobra.ExactArgs(1),
	}
)

func run(cmd *cobra.Command, args []string) {
	url := args[0]
	method, _ := cmd.Flags().GetString("method")

	if url == "" {
		cmd.Help()
		os.Exit(1)
	}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error executing request:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("method", "X", "GET", "HTTP method to use")
}
