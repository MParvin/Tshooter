package cmd

import (
	"fmt"

	tools "github.com/mparvin/tshooter/tools"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read config file(s) and list contexts",
	Long: `It read kubernetes config file and then list all contexts, for example:
	$ tshooter config ~/.kube/config
	will read config file ~/.kube/config
	$ tshooter config -l
	or
	$ tshooter config --list
	will list all contexts`,
	Run: func(cmd *cobra.Command, args []string) {

		listFlag := cmd.Flags().Lookup("list")
		if listFlag != nil && listFlag.Changed {
			fmt.Println("Listing all contexts")
			config := tools.GetConfig(
			tools.ShowContexts(config)
			return
		}
		fmt.Println("Reading config file(s) and list contexts")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().String("list", "l", "List contexts")

	configCmd.Flags().BoolP("list", "l", true, "List contexts")
}
