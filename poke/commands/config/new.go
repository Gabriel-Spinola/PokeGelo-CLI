package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/cmd_config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	JSON string = "json"
	YAML string = "yaml"
	TOML string = "toml"
)

var TYPES = [3]string{JSON, YAML, TOML}

var (
	configPath     string = ""
	configFileName string = cmd_config.DefaultConfigFileName
	configFileType string = cmd_config.DefaultConfigFileType
)

func setFileType() error {
	for i := 0; i < len(TYPES); i++ {
		if configFileType == string(TYPES[i]) {
			viper.SetConfigType(configFileType)

			return nil
		}
	}

	return errors.New("Invalid config file type (extension): " + configFileType)
}

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		output_path := configPath + configFileName + "." + string(configFileType)
		err := os.WriteFile(output_path, nil, 0644)
		if err != nil {
			log.Fatalf("Failed to create config file. %s\n", err)

			return
		}

		err = setFileType()
		if err != nil {
			log.Fatalln(err)

			return
		}

		viper.AddConfigPath(output_path)
		fmt.Println("Config file have been succesfuly created at: " + output_path)
	},
}

func setNewFlags() {
	NewCmd.Flags().StringVarP(&configPath, "path", "p", "", "Path to config file")
	NewCmd.Flags().StringVarP(&configFileType, "file-type", "t", cmd_config.DefaultConfigFileType, "Config file type (extension)")
	NewCmd.Flags().StringVarP(&configFileName, "file-name", "n", cmd_config.DefaultConfigFileName, "Config file name")

	if err := NewCmd.MarkFlagRequired("path"); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	setNewFlags()
}
