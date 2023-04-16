package cmd

import (
	// "fmt"

	"os"
	"strings"

	// 	"strconv"
	// 	"strings"

	// 	"path/filepath"

	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// // createfiles --size 10m --count 5 --path data/
var (
	// Used for flags.
	filePath       string
	calculationCmd = &cobra.Command{
		Use:   "createfiles",
		Short: "A generator for Cobra based Applications build by Qays",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application. build by Qays Dwekat`,
		RunE: func(cmd *cobra.Command, args []string) error {
			paths, err := getAllFilesInDirectory(filePath)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			totalWords := 0
			for _, file := range paths {
				words, err2 := calculationWords(file)
				if err2 != nil {
					return fmt.Errorf("%v", err2)
				}
				totalWords += words
			}
			fmt.Println("Total number of words in the Directory: ", totalWords)
			return nil
		},
	}
)

// Execute executes the root command.
func ExecuteCalculation() error {
	return calculationCmd.Execute()
}

func init() {
	// Add Path flag to the command
	calculationCmd.Flags().StringVar(&filePath, "path", "/", "path of the generated files")
}
func getAllFilesInDirectory(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	filePaths := []string{}
	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			paths, err1 := getAllFilesInDirectory(path + file.Name())
			if err1 != nil {
				return nil, err1
			}
			filePaths = append(filePaths, paths...)
		} else {
			filePaths = append(filePaths, path+"/"+file.Name())
		}
	}
	return filePaths, nil
}

func calculationWords(file string) (int, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	value := string(data)
	numberOfwords := strings.Split(value, " ")

	return len(numberOfwords), nil
}
