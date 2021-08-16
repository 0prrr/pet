package cmd

import (
	"fmt"
	"strings"
    "os"
	"os/exec"

	"github.com/knqyf263/pet/config"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var delimiter string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search snippets",
	Long:  `Search snippets interactively (default filtering tool: peco)`,
	RunE:  search,
}

func search(cmd *cobra.Command, args []string) (err error) {
	flag := config.Flag

	var options []string
	if flag.Query != "" {
		options = append(options, fmt.Sprintf("--query %s", flag.Query))
	}
	commands, err := filter(options)
	if err != nil {
		return err
	}

    finalCommand := strings.Join(commands, flag.Delimiter)
    if strings.HasPrefix(finalCommand, "firefox") || strings.HasPrefix(finalCommand, "chrome"){
        var cmd *exec.Cmd
		cmd = exec.Command("sh", "-c", finalCommand)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
        cmd.Stdin = os.Stdin
        if err := cmd.Run(); err != nil {
            return err
        }
        return nil
    }
    fmt.Println(finalCommand)
	if terminal.IsTerminal(1) {
		fmt.Print("\n")
	}

	return nil
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&config.Flag.Color, "color", "", false,
		`Enable colorized output (only fzf)`)
	searchCmd.Flags().StringVarP(&config.Flag.Query, "query", "q", "",
		`Initial value for query`)
	searchCmd.Flags().StringVarP(&config.Flag.Delimiter, "delimiter", "d", "; ",
		`Use delim as the command delimiter character`)
}
