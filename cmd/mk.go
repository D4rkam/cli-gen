package cmd

import (
	"github.com/spf13/cobra"
)

// mkCmd represents the mk command
var mkCmd = &cobra.Command{
	Use:   "mk",
	Short: "Generador de paquetes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		CreateMainFile(cmd)
	},
}

func init() {
	rootCmd.AddCommand(mkCmd)

	//Frameworks
	mkCmd.Flags().BoolP("echo", "e", false, "Creacion estandar para el framework echo")
	mkCmd.Flags().BoolP("fiber", "f", false, "Creacion estandar para el framework fiber")
	mkCmd.Flags().BoolP("gin", "g", false, "Creacion estandar para el framework gin")

	//Template
	mkCmd.Flags().BoolP("template", "t", false, "Creacion basada en template")

	//Package
	mkCmd.Flags().BoolP("controller", "c", false, "Creacion de paquete controller")
	mkCmd.Flags().BoolP("service", "s", false, "Creacion de paquete service")
	mkCmd.Flags().BoolP("repository", "r", false, "Creacion de paquete repository")
}
