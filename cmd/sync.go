package cmd

import (
	"github.com/knqyf263/pet/config"
	petSync "github.com/knqyf263/pet/sync"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
}

func sync(cmd *cobra.Command, args []string) (err error) {
	return petSync.AutoSync(config.Conf.General.SnippetFile)
}

func init() {
	RootCmd.AddCommand(syncCmd)
}
