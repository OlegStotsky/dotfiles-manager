/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/asverdlov/dotfiles/pkg/registry"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a component",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		componentName := args[0]
		registry := registry.Registry()

		component, found := registry[componentName]
		if !found {
			return fmt.Errorf("No such component: %s", componentName)
		}

		if component.Install == nil {
			return fmt.Errorf("Install is not configured for component: %s", componentName)
		}

		return component.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	//installCmd.PersistentFlags().String("component", "", "Component to install")
}
