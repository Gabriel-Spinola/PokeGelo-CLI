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

// TODO - Passing directories

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
		log.Fatal("Failed to unmarshal file", err)

		return FileData{}, err
	}

	return FileData{Data: data, bytes: byteValue}, nil
}

func handleFileRead(filePath string) bool {
	var result = new(lib.Request)

	fileData, err := readFileData(filePath)
	if err != nil {
		return false
	}

	result.Body = make(map[string]any)
	data, ok := fileData.Data.(map[string]any)
	if !ok {
		log.Fatal("Failed to read file data (Wrong format)")

		return false
	}

	result.Body = data
	bytes, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal body")

		return false
	}

	err = os.WriteFile("output"+lib.PATH_SEPARATOR+"builder_output_"+filepath.Base(filePath), bytes, 0644)
	if err != nil {
		log.Fatal("Failed to write json to file: ", err)

		return false
	}

	return true
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Build a template request json with the given body data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		if len(targetFilePaths) == 1 {
			handleFileRead(targetFilePaths[0])

			fmt.Println("TOOK: ", time.Since(start))
			return
		}

		isDoneChan := make(chan bool)

		go lib.ConcurrentFileProcessor(targetFilePaths, isDoneChan, handleFileRead)

		for result := range isDoneChan {
			fmt.Println(result)
		}

		<-isDoneChan

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
