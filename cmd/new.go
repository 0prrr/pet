package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/knqyf263/pet/config"
	"github.com/knqyf263/pet/snippet"
	"github.com/spf13/cobra"
)

var is_markdown_mode bool

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new COMMAND",
	Short: "Create a new snippet",
	Long:  `Create a new snippet (default: $HOME/.config/pet/snippet.toml)`,
	RunE:  new,
}

func scan(message string) (string, error) {
	tempFile := "/tmp/pet.tmp"
	if runtime.GOOS == "windows" {
		tempDir := os.Getenv("TEMP")
		tempFile = filepath.Join(tempDir, "pet.tmp")
	}
	l, err := readline.NewEx(&readline.Config{
		Prompt:          message,
		HistoryFile:     tempFile,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		return "", err
	}
	defer l.Close()

    var cmds []string

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
            continue
		}
	if is_markdown_mode {
	    if line != "eof" && line != "EOF" {
		line += " \\"
		cmds = append(cmds, line)
		l.SetPrompt(color.YellowString("Command> "))
		continue
	    }
	} else {
	    cmds = append(cmds, line)
	}

        var finalCmd string
        if strings.HasPrefix(cmds[0], "#") {
            finalCmd = strings.Join(cmds, "\n\n")
        } else {
            finalCmd = strings.Join(cmds, " ")
        }
		return finalCmd, nil
	}
	return "", errors.New(color.RedString("canceled..."))
}

func scan_desc(message string) (string, error) {
	var desc string = ""
	tempFile := "/tmp/pet.tmp"
	if runtime.GOOS == "windows" {
		tempDir := os.Getenv("TEMP")
		tempFile = filepath.Join(tempDir, "pet.tmp")
	}
	l, err := readline.NewEx(&readline.Config{
		Prompt:          message,
		HistoryFile:     tempFile,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		return "", err
	}
	defer l.Close()
	for {
		desc, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(desc) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		desc = strings.TrimSpace(desc)

		if desc == "" {
			l.SetPrompt(color.GreenString("Description> "))
			continue
		}

		return desc, nil
	}
	return desc, nil
}

func scan_mode(message string) (bool, error) {
	tempFile := "/tmp/pet.tmp"
	if runtime.GOOS == "windows" {
		tempDir := os.Getenv("TEMP")
		tempFile = filepath.Join(tempDir, "pet.tmp")
	}
	l, err := readline.NewEx(&readline.Config{
		Prompt:          message,
		HistoryFile:     tempFile,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		return false, err
	}
	defer l.Close()
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
	            l.SetPrompt(color.MagentaString("Mode> "))
		    continue
		}
		if line == "m" {
			return true, nil
		} else if line == "n" {
			return false, nil
		} else {
			break
		}
	}
	return false, errors.New(color.RedString("Choose between 'm' and 'n'..."))
}

func new(cmd *cobra.Command, args []string) (err error) {
	var command string
	var description string
	var tags []string

	var snippets snippet.Snippets
	if err := snippets.Load(); err != nil {
		return err
	}

	if len(args) > 0 {
		command = strings.Join(args, " ")
		fmt.Fprintf(color.Output, "%s %s\n", color.YellowString("Command>"), command)
	} else {
		is_markdown_mode, err = scan_mode(color.MagentaString("Mode> "))
		if err != nil {
			return err
		}
		command, err = scan(color.YellowString("Command> "))
		if err != nil {
			return err
		}
	}
	description, err = scan_desc(color.GreenString("Description> "))
	if err != nil {
		return err
	}

	if config.Flag.Tag {
		var t string
		if t, err = scan(color.CyanString("Tag> ")); err != nil {
			return err
		}
		tags = strings.Fields(t)
	}

	for _, s := range snippets.Snippets {
		if s.Description == description {
			return fmt.Errorf("Snippet [%s] already exists", description)
		}
	}

	newSnippet := snippet.SnippetInfo{
		Description: description,
		Command:     command,
		Tag:         tags,
	}
	snippets.Snippets = append(snippets.Snippets, newSnippet)
	if err = snippets.Save(); err != nil {
		return err
	}

	return nil
}

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolVarP(&config.Flag.Tag, "tag", "t", false,
		`Display tag prompt (delimiter: space)`)
}
