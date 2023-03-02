/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gopher-cli",
	Short: "Gopher CLI in go",
	Long: `Gopher CLI application written in Go.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName = "trafalgar"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}
		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

        fmt.Println("Try to get '" + gopherName + "' Gopher...")

		// Get the data
        response, err := http.Get(URL)
        if err != nil {
            fmt.Println(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {
            // Create the file
            out, err := os.Create(gopherName + ".png")
            if err != nil {
                fmt.Println(err)
            }
            defer out.Close()

            // Writer the body to file
            _, err = io.Copy(out, response.Body)
            if err != nil {
                fmt.Println(err)
            }

            fmt.Println("Perfect! Just saved in " + out.Name() + "!")
        } else {
            fmt.Println("Error: " + gopherName + " not exists! :-(")
        }
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gopher-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


