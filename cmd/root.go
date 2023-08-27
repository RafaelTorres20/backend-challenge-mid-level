package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cvtm",
	Short: "Api de consulta de ativos",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
