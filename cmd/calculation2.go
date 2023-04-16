package cmd

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/spf13/cobra"
)

// createfiles --size 10m --count 5 --path data/
var (
	// Used for flags.
	filePath2        string
	perfile, perline bool
	calculation2Cmd  = &cobra.Command{
		Use:   "createfiles",
		Short: "A generator for Cobra based Applications build by Qays",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application. build by Qays Dwekat`,
		RunE: func(cmd *cobra.Command, args []string) error {
			paths, err := getAllFilesInDirectory(filePath2)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			var totalWords int64

			var wg sync.WaitGroup

			for _, file := range paths {
				wg.Add(2)

				go func() error {
					words, err2 := calculationWords(file)
					if err2 != nil {
						return fmt.Errorf("%v", err2)
					}

					atomic.AddInt64(&totalWords, int64(words))

					wg.Done()
					return nil
				}()
			}

			wg.Wait()

			fmt.Println("Total number of words in the Directory: ", totalWords)
			return nil
		},
	}
)

// Execute executes the root command.
func ExecuteCalculation2() error {
	return calculation2Cmd.Execute()
}

func init() {
	// Add Path flag to the command
	calculation2Cmd.Flags().StringVar(&filePath2, "path", "/", "path of the generated files")
	calculation2Cmd.Flags().BoolVar(&perfile, "perfile", false, "path of the generated files")
	calculation2Cmd.Flags().BoolVar(&perline, "perline", false, "path of the generated files")
}
