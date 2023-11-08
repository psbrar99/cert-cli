package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var Config struct {
	Cert   map[string]*cert.Cert `yaml:cert`
	CACert *cert.CACert          `yaml:caCert`
}
var cfgFile Config
var cfgFilePath string

var rootCmd = &cobra.Command{
	Use:   "x509 certs",
	Short: "certcli create certs",
	Long:  `A Fast and Flexible way to create CA and client certs for go applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("here here")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "-c", "config", "", "config file (default is tls.yaml)")
}

func initConfig() {
	if cfgFilePath == "" {
		// Use config file from the flag.
		cfgFilePath = "tls.yaml"
	}

	readBytes, err := os.ReadFile("cfgFilePath")
	if err != nil {
		fmt.Printf("Not able to read file %s", err)
	}

	err = yaml.Unmarshal(readBytes, &cfgFile)

}
