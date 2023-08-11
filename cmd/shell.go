// cmd/shell.go

package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	// ... potentially other imports ...
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a REPL shell for procedural language",
	Run:   shellHandler,
}

func shellHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Enter your commands (end with ;; to execute):")
	scanner := bufio.NewScanner(os.Stdin)
	var batch string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == ";;" {
			// Call to your language executor here
			// executeQuery(batch)
			batch = ""
			fmt.Println("Enter your commands (end with ;; to execute):")
		} else {
			batch += line + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
