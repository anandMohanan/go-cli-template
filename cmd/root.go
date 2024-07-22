package cmd

import (
	"strings"

	"github.com/anandMohanan/go-cli-template/constants"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var cobraInit = &cobra.Command{
	Use:   strings.ToLower(constants.AppName),
	Short: constants.ShortDescription,
	Long:  constants.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		// Add your root command logic here
		// Example usage of help
		cmd.Help()
	},
}

func init() {
	cobraInit.Flags().BoolP("version", "v", false, constants.AppName+" version")
}

func Execute() {

	cc.Init(&cc.Config{
		RootCmd:         cobraInit,
		Headings:        cc.Yellow + cc.Bold + cc.Underline,
		Commands:        cc.Green + cc.Bold,
		Example:         cc.Cyan + cc.Bold,
		Flags:           cc.Red + cc.Italic,
		FlagsDataType:   cc.HiWhite + cc.Italic,
		NoExtraNewlines: false,
		NoBottomNewline: false,
	})

	 cobraInit.Execute()
}
