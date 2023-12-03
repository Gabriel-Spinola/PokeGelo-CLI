package net

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
	"github.com/spf13/cobra"
)

var targetFilePath string

type FileData struct {
	Data  interface{}
	bytes []byte
}

func readFileData() (FileData, error) {
	var data interface{}

	jsonFile, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal("Failed ro read json")

		return FileData{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if len(byteValue) <= 0 {
		log.Fatal("The json file can't be epty")

		return FileData{}, err
	}

	if err := json.Unmarshal(byteValue, &data); err != nil {
		return FileData{}, err
	}

	return FileData{Data: data, bytes: byteValue}, nil
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Build a template request json with the given body data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var result = new(lib.Request)

		fileData, err := readFileData()
		if err != nil {
			return
		}

		result.Body = make(map[string]interface{})
		result.Body["data"] = fileData.Data

		bytes, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatal("Failed to marshal body")

			return
		}

		err = os.WriteFile("output/builder_output.json", bytes, 0644)
		if err != nil {
			log.Fatal("Failed to write json to file")

			return
		}
	},
}

func setReadFlags() {
	readCmd.Flags().StringVarP(&targetFilePath, "filepath", "f", "", "The path to the file")

	if err := readCmd.MarkFlagRequired("filepath"); err != nil {
		fmt.Println(err)
	}
}

func init() {
	NetCmd.AddCommand(readCmd)

	setReadFlags()
}
