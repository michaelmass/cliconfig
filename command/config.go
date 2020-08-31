package command

import (
	"log"

	"github.com/spf13/cobra"
)

// ShowFunc is used to show the config file
type ShowFunc func() error

// OpenFunc is used to open the config file
type OpenFunc func() error

// ResetFunc is used to reset the config file
type ResetFunc func() error

func newConfig(showFunc ShowFunc, openFunc OpenFunc, resetFunc ResetFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Config has utilities for the configuration file.",
		Long:  `Config has utilities for the configuration file.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newConfigShow(showFunc))
	cmd.AddCommand(newConfigOpen(openFunc))
	cmd.AddCommand(newConfigReset(resetFunc))

	return cmd
}

func newConfigShow(showFunc ShowFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Shows the config file content.",
		Long:  `Shows the config file content.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := showFunc()

			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return cmd
}

func newConfigOpen(openFunc OpenFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open the config file inside the default editor.",
		Long:  `Open the config file inside the default editor.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := openFunc()

			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return cmd
}

func newConfigReset(resetFunc ResetFunc) *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset update the configuration file with default values.",
		Long:  `Reset update the configuration file with default values.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := resetFunc()

			if err != nil {
				log.Fatal(err)
			}
		},
	}
}
