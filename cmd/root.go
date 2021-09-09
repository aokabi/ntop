package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getNginxStub() {
	req, err := client.Get(fmt.Sprintf("http://%s:%d/nginx_status",host, port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer req.Body.Close()
	// read request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print current time
	fmt.Println(time.Now().Truncate(time.Second))
	fmt.Println(string(body))
}

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			for {
				time.Sleep(1 * time.Second)
				getNginxStub()
			}
		},
	}
	client = &http.Client{}
	port int
	host string
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().IntVar(&port, "port", 80, "port to request nginx")
	rootCmd.PersistentFlags().StringVar(&host, "host", "127.0.0.1", "host to request nginx")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("license", "apache")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
