package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CreateMainFile(cmd *cobra.Command) {
	flags, _ := cmd.Flags().GetBool("echo")
	if flags {
		fmt.Println("Utilizastes la flag del framework echo")
	}
}
