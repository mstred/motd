package cmd

import (
	"fmt"
	"os"

	"github.com/mstred/motd/message"
	"github.com/spf13/cobra"
)

var name, greeting string
var preview, prompt bool
var debug bool = false

var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !prompt && (name == "" || greeting == "") {
			cmd.Usage()
			os.Exit(1)
		}

		if debug {
			fmt.Println("Name:", name)
			fmt.Println("Greeting:", greeting)
			fmt.Println("Prompt:", prompt)
			fmt.Println("Preview:", preview)

			os.Exit(0)
		}

		// conditionally read from stdin
		if prompt {
			message.Readtovar(&name, "Your Name: ")
			message.Readtovar(&greeting, "Your Greeting: ")
		}

		// generate message
		output := message.Greeting(name, greeting)

		// either preview the message or write to file
		if preview {
			fmt.Println(output)
		} else {
			// write content
			file, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)
			defer file.Close()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			_, err = file.Write([]byte(output))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name to use in the message")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting to use in the message")
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "Preview message instead of writing to /etc/motd")
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", false, "Prompt for name and greeting")

	if os.Getenv("DEBUG") != "" {
		debug = true
	}
}
