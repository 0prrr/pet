package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"regexp"

	"github.com/fatih/color"
	"github.com/knqyf263/pet/dialog"
	"github.com/knqyf263/pet/config"
	"github.com/knqyf263/pet/snippet"
)

func editFile(command, file string) error {
	command += " " + file
	return run(command, os.Stdin, os.Stdout)
}

func run(command string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}

func filter(options []string) (commands []string, imageUrls []string, err error) {
	var snippets snippet.Snippets
	if err := snippets.Load(); err != nil {
		return commands, nil,fmt.Errorf("Load snippet failed: %v", err)
	}

	snippetTexts := map[string]snippet.SnippetInfo{}
	var text string
	var command string
	for _, s := range snippets.Snippets {
		command = s.Command
		if strings.ContainsAny(command, "\n") {
			command = strings.Replace(command, "\n", "\\n", -1)
		}
		t := fmt.Sprintf("[%s]: %s", s.Description, command)

		tags := ""
		for _, tag := range s.Tag {
			tags += fmt.Sprintf(" #%s", tag)
		}
		t += tags

		snippetTexts[t] = s
		if config.Flag.Color {
			t = fmt.Sprintf("[%s]: %s%s",
				color.RedString(s.Description), command, color.BlueString(tags))
		}
		text += t + "\n"
	}

	var buf bytes.Buffer
	selectCmd := fmt.Sprintf("%s %s",
		config.Conf.General.SelectCmd, strings.Join(options, " "))
	err = run(selectCmd, strings.NewReader(text), &buf)
	if err != nil {
		return nil, nil, nil
	}

	lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

	if !strings.Contains(lines[0], "#") {
	    params := dialog.SearchForParams(lines)
	    if params != nil {
		snippetInfo := snippetTexts[lines[0]]
		dialog.CurrentCommand = snippetInfo.Command
		dialog.GenerateParamsLayout(params, dialog.CurrentCommand)
		res := []string{dialog.FinalCommand}
		return res, nil, nil
	    }
	}

	var urls []string
	for _, line := range lines {
		snippetInfo := snippetTexts[line]
        cmd := snippetInfo.Command
        if strings.Contains(cmd, "\\") {
            cmd = strings.Replace(cmd, "\\", "", -1)
	    if strings.Contains(cmd, "#") {
		if strings.Contains(cmd, "|") || strings.Contains(cmd, "```") {
                    cmd = strings.Replace(cmd, "\n\n", "\n", -1)
                }
	    }

	    if strings.Contains(cmd, "img::") {
		r, _ := regexp.Compile("img([ -~]+)")
		results := r.FindAllString(cmd, -1)
		for _, url := range results {
		    urls = append(urls, strings.Replace(url, "img::", "", -1))
		}
	    }
            commands = append(commands, fmt.Sprint(cmd))
        } else {
            commands = append(commands, fmt.Sprint(snippetInfo.Command))
        }
	}
	return commands, urls, nil
}
