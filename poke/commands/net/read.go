package net

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
	"github.com/spf13/cobra"
)

var targetFilePaths []string

type FileData struct {
	Data  interface{}
	bytes []byte
}

func readFileData(filePath string) (FileData, error) {
	var data interface{}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed ro read json", err)

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

func processFile(filePath string) {
	var result = new(lib.Request)

	fileData, err := readFileData(filePath)
	if err != nil {
		return
	}

	result.Body = make(map[string]any)
	data, ok := fileData.Data.(map[string]any)
	if !ok {
		log.Fatal("Failed to read file data (Wrong format)")

		return
	}

	result.Body = data
	bytes, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal body")

		return
	}

	err = os.WriteFile("output"+lib.PATH_SEPARATOR+"builder_output_"+filepath.Base(filePath), bytes, 0644)
	if err != nil {
		log.Fatal("Failed to write json to file: ", err)

		return
	}
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Build a template request json with the given body data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		for _, filePath := range targetFilePaths {
			processFile(filePath)
		}

		fmt.Println("TOOK: ", time.Since(start))
	},
}

func setReadFlags() {
	readCmd.Flags().StringSliceVarP(&targetFilePaths, "filepath", "f", []string{}, "The path to the files")

	if err := readCmd.MarkFlagRequired("filepath"); err != nil {
		fmt.Println(err)
	}
}

func init() {
	NetCmd.AddCommand(readCmd)

	setReadFlags()
}
