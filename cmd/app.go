package cmd

import (
	"context"
	"fmt"
	"os"
	
	"github.com/gobuffalo/genny"
	"github.com/spf13/cobra"
	
	"github.com/CosmicCompass/starport/pkg/gomodulepath"
	"github.com/CosmicCompass/starport/templates/app"
)

var appCmd = &cobra.Command{
	Use:   "app [github.com/org/repo]",
	Short: "Generates an empty application",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := gomodulepath.Parse(args[0])
		if err != nil {
			return err
		}
		denom, _ := cmd.Flags().GetString("denom")
		account_prefix, _ := cmd.Flags().GetString("account_prefix")
		
		g, _ := app.New(&app.Options{
			ModulePath:       path.RawPath,
			AppName:          path.Package,
			BinaryNamePrefix: path.Root,
			Denom:            denom,
			AccountPrefix:    account_prefix,
		})
		run := genny.WetRunner(context.Background())
		run.With(g)
		pwd, _ := os.Getwd()
		run.Root = pwd + "/" + path.Root
		run.Run()
		message := `
⭐️ Successfully created a Cosmos app '%[1]v'.
👉 Get started with the following commands:

 %% cd %[1]v
 %% starport serve

NOTE: add -v flag for advanced use.
`
		fmt.Printf(message, path.Root)
		return nil
	},
}
