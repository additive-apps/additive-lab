package cmd

import (
	"additive-apps/additive-lab/src/print"
	"additive-apps/additive-lab/src/util/file"
	"additive-apps/additive-lab/src/util/ssh"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Generate and configure SSH key",
	Run: func(cmd *cobra.Command, args []string) {
		// Shows help only if no sub-commands were passed.
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates ssh key",
	Run: func(cmd *cobra.Command, args []string) {
		isConfigEnabled, err := cmd.Flags().GetBool("config")

		if err != nil {
			print.Fatal("Error getting flag value:", err)
			return
		}

		if isConfigEnabled {
			ssh.GenerateConfigFile()
		}

		ssh.Generate()
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Creates ssh config",
	Run: func(cmd *cobra.Command, args []string) {
		ssh.GenerateConfigFile()
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show public ssh key content",
	Run: func(cmd *cobra.Command, args []string) {
		sshPaths := ssh.NewSSHPaths()
		content := file.Content(sshPaths.Config)
		print.Normal(content)
	},
}

func init() {
	createCmd.Flags().Bool("config", true, "create config file")

	sshCmd.AddCommand(createCmd)
	sshCmd.AddCommand(configCmd)
	sshCmd.AddCommand(showCmd)

	rootCmd.AddCommand(sshCmd)
}
