package generate

import (
	"github.com/VladimirRytov/advsrv/internal/configparser"
	"github.com/spf13/cobra"
)

var GenConf = &cobra.Command{Use: "generate", Short: "print config template to stdout", RunE: func(cmd *cobra.Command, args []string) error {
	return Generate()
}}

func Generate() error {
	ch := configparser.NewConfigparser(nil, nil)
	return ch.GenerateTemplate()
}
