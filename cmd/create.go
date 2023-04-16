package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"path/filepath"

	"github.com/spf13/cobra"
)

// createfiles --size 10m --count 5 --path data/
var (
	// Used for flags.
	size, count, path string
	supportedUnits    = []string{"kb", "mb", "gb", "tb"}
	createCmd         = &cobra.Command{
		Use:   "createfiles",
		Short: "A generator for Cobra based Applications build by Qays",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application. build by Qays Dwekat`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// convert count integer
			countInt, err := strconv.Atoi(count)
			if err != nil {
				return fmt.Errorf("invalid count: %v", err)
			}

			// get file size in byte
			sizeInt, err := getFileSize(size)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			// convert directory if not exist
			createDirectory(path)

			for i := 1; i <= countInt; i++ {
				file := createFile(path, i)
				writeFileContent(sizeInt, file)
				defer file.Close()
			}
			return nil
		},
	}

	supportedUnitsCmd = &cobra.Command{
		Use:   "units",
		Short: "Print the supported units of createfiles",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("The supported units: ", supportedUnits)
		},
	}
)

// Execute executes the root command.
func ExecuteCreation() error {
	return createCmd.Execute()
}

func init() {
	// Add Count flag to the command
	createCmd.Flags().StringVar(&count, "count", "1", "number the generated files")

	// Add Size flag to the command
	createCmd.Flags().StringVar(&size, "size", "1024", "size of the generated files")

	// Add Path flag to the command
	createCmd.Flags().StringVar(&path, "path", "/", "path of the generated files")

	createCmd.AddCommand(supportedUnitsCmd)
}

func writeFileContent(size int, file *os.File) error {
	tries := size / 1024
	for i := 1; i <= tries; i++ {
		_, err2 := file.WriteString("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus pellentesque augue mi, a aliquet orci tristique eu. Suspendisse luctus turpis pharetra quam facilisis, quis hendrerit turpis rhoncus. Nullam interdum laoreet aliquet. Pellentesque et quam quam. Curabitur metus mi, dapibus nec est sit amet, fringilla placerat ante. Vivamus facilisis justo mollis sapien fermentum, et porta elit fringilla. Curabitur vel ultricies turpis, vel sodales diam. Nam aliquet tellus vitae ligula pellentesque, at venenatis sem condimentum. Nulla facilisi. Mauris auctor feugiat velit, in sodales sem accumsan non. Fusce blandit vel dui quis feugiat. Maecenas congue imperdiet quam, a mattis libero mattis varius.Nam aliquet ligula sit amet placerat tincidunt. Praesent eu ultricies nibh, id efficitur enim. Suspendisse eu molestie tortor. Sed pulvinar ornare massa, vitae convallis elit elementum at. Mauris egestas enim vel volutpat mollis. Nullam convallis ipsum augue, non tristique sem ultricies auctor. Mauris quis arcu id dui.")
		if err2 != nil {
			return fmt.Errorf("fail to write on file: %v", err2)
		}
	}
	fmt.Println("write data completed")
	return nil
}

func createFile(path string, count int) *os.File {
	filename := "file_number_" + strconv.Itoa(count) + ".txt"
	pathOfTheFile := filepath.Join(path, filename)
	fmt.Println(pathOfTheFile)
	file, err := os.Create(pathOfTheFile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File created successfully")
	}
	return file
}

func createDirectory(path string) {

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dir created successfully")
	}
}

func getFileSize(size string) (int, error) {
	fmt.Println(supportedUnits)
	unit := strings.ToLower(size[len(size)-2:])
	if contains(supportedUnits, unit) {
		value := size[:len(size)-2]
		sizeInt, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}

		if unit == "kb" {
			return sizeInt * 1024, nil
		} else if unit == "mb" {
			return sizeInt * 1024 * 1024, nil
		} else if unit == "gb" {
			return sizeInt * 1024 * 1024 * 1024, nil
		} else if unit == "tb" {
			return sizeInt * 1024 * 1024 * 1024 * 1024, nil
		}

		return sizeInt, nil
	}

	return 0, fmt.Errorf("pleas use valid unit.\nsupported units %v", supportedUnits)
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
