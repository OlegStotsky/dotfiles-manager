/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/asverdlov/dotfiles/pkg/registry"
	"github.com/spf13/cobra"
)

const (
	tickEmoji         = "\u2705" // https://emojipedia.org/check-mark-button/
	crossEmoji        = "\u274C" // https://emojipedia.org/cross-mark/
	questionMarkEmoji = "\u2753" // https://emojipedia.org/question-mark/
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available components",
	Run: func(cmd *cobra.Command, args []string) {
		registry := registry.Registry()

		s := "Components:\n"
		for _, c := range registry {
			state := questionMarkEmoji
			if c.IsInstalled != nil {
				isInstalled, err := c.IsInstalled()
				if err != nil {
					state = fmt.Sprintf("error: %s", err)
				} else if isInstalled {
					state = tickEmoji
				} else {
					state = crossEmoji
				}
			}
			s += fmt.Sprintf("- %s (%s)\n", c.Name, state)
		}

		fmt.Printf(s)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
