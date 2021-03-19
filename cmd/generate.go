package cmd

import (
	"fmt"

	"github.com/onigura/paich/parser"
	"github.com/spf13/cobra"
)

var templateType string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code based available template ",
	Long: `This command for generate code from template available by providing template name.
`,
	Run: func(cmd *cobra.Command, args []string) {
		g := parser.GenerateCmdArgToken{}
		g.Parse(&args, templateType)

		err := parser.Flat(g)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	generateCmd.Flags().StringVarP(&templateType, "template", "t", "", "Template type, available: gin-rest")
	generateCmd.SetUsageTemplate(`Usage: paich generate [label name] [field name]:[field type]:[field attribute]

Field descriptor format:

[label name]	name of code want to generate (e.g. json object name)
[field name]	name of field (e.g. json key field name)
[field type]	data type of field (e.g. SQL table type)
[field attr]	misc. for field.

This field will inflated with template.

Flags:

--template	Template type; available: gin-rest
	
`)
	rootCmd.AddCommand(generateCmd)
}
