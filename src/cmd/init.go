package cmd

import (
	"additive-apps/additive-lab/src/print"
	"additive-apps/additive-lab/src/util"
	"additive-apps/additive-lab/src/util/dir"
	"additive-apps/additive-lab/src/util/ssh"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generates SSH key and the additive-lab directories",
	Run: func(cmd *cobra.Command, args []string) {
		ssh.Generate()
		ssh.GenerateConfigFile()
		createLabDirs()
	},
}

func createLabDirs() {
	labPath := util.HomePath() + "/additive-lab/"

	dir.NotExists(labPath, func() {
		reposPath := labPath + "repositories/"
		err := os.MkdirAll(reposPath, 0750)

		if err != nil && !os.IsExist(err) {
			print.Fatal(err)
		}

		print.Create(labPath)
		print.Create(reposPath)
	})
}

func init() {
	rootCmd.AddCommand(initCmd)
}
