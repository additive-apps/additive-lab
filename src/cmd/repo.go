package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"additive-apps/additive-lab/src/print"
	"additive-apps/additive-lab/src/util"
	"additive-apps/additive-lab/src/util/constants"
	"additive-apps/additive-lab/src/util/dir"
	"additive-apps/additive-lab/src/util/dirpath"
)

var repositories = map[string]string{
	"account":              "git@github.com:additive-apps/account.git",
	"content":              "git@github.com:additive-apps/content.git",
	"crm":                  "git@github.com:additive-apps/crm.git",
	"journal-widget":       "git@github.com:additive-apps/journal-widget.git",
	"laravel-apps":         "git@github.com:additive-apps/laravel-apps.git",
	"mailer":               "git@github.com:additive-apps/mailer.git",
	"marketing-automation": "git@github.com:additive-apps/marketing-automation.git",
	"multimedia":           "git@github.com:additive-apps/multimedia.git",
	"newsletter":           "git@github.com:additive-apps/newsletter.git",
	"newsletter-widget":    "git@github.com:additive-apps/newsletter-widget.git",
	"popup-widget":         "git@github.com:additive-apps/popup-widget.git",
	"ui":                   "git@github.com:additive-apps/ui.git",
	"voucher":              "git@github.com:additive-apps/voucher.git",
	"voucher-widget":       "git@github.com:additive-apps/voucher-widget.git",
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Intracts with ADDITIVE repositories",
	Run: func(cmd *cobra.Command, args []string) {
		// Shows help only if no sub-commands were passed.
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones ADDITIVE repositories",
	Long:  "Clones a given repository into $HOME/additive-lab/repositories",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		repoUrl := repositories[repoName]

		if len(repoUrl) == 0 {
			print.Fatal("No repository named %s could be found.", repoName)
		}

		repoPath := dirpath.Join(
			util.HomePath(),
			constants.DIR_ADDITIVE_LAB,
			constants.DIR_REPOSITORIES,
			repoName,
		)

		dir.NotExists(repoPath, func() {
			cmd := exec.Command("git", "clone", repoUrl, repoPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()

			if err != nil {
				print.Fatal(err)
			}
		})
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories that can be cloned",
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "REPOSITORY\tURL")

		for name, url := range repositories {
			fmt.Fprintf(w, "%s\t%s\n", name, url)
		}

		w.Flush()
	},
}

func init() {
	repoCmd.AddCommand(cloneCmd)
	repoCmd.AddCommand(listCmd)
	rootCmd.AddCommand(repoCmd)
}
