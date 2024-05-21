package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/VladimirRytov/advsrv/internal/configparser"
	"github.com/VladimirRytov/advsrv/internal/datatransferobjects"
	"github.com/VladimirRytov/advsrv/internal/front/cli/generate"
	"github.com/VladimirRytov/advsrv/internal/front/cli/start"
	"github.com/VladimirRytov/advsrv/internal/front/rpcworker"
	"github.com/spf13/cobra"
)

type CommandLine struct {
	root *cobra.Command
}

const (
	stopCmd        = "stop"
	reloadCmd      = "reload"
	restartCmd     = "restart"
	reloadCacheCmd = "reloadcache"
	cleanCacheCmd  = "cleancache"
	versionCmd     = "version"
)

var (
	stop = &cobra.Command{Use: stopCmd,
		Short: "stop server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rpcworker.Send(&datatransferobjects.Message{Code: rpcworker.Stop})
		}}

	reload = &cobra.Command{Use: reloadCmd, Short: "reload server",
		Long: "close all existing connections and creating new ones with params in the configuration file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rpcworker.Send(&datatransferobjects.Message{Code: rpcworker.Reload})
		}}
	restart = &cobra.Command{Use: restartCmd, Short: "restart server",
		Long: "close all existing connections and creating new ones with params in the configuration file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := rpcworker.Send(&datatransferobjects.Message{Code: rpcworker.Stop})
			if err != nil {
				return err
			}
			return start.RunMain()
		}}

	reloadCache = &cobra.Command{Use: reloadCacheCmd,
		Short: "reload miniatures cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rpcworker.Send(&datatransferobjects.Message{Code: rpcworker.ReloadCache})
		}}
	cleanCache = &cobra.Command{Use: cleanCacheCmd,
		Short: "remove not existing files from miniatures cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			return rpcworker.Send(&datatransferobjects.Message{Code: rpcworker.CleanCache})
		}}

	version = &cobra.Command{Use: versionCmd, Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(start.Version)
	}}
)

func NewCommandLine() *CommandLine {
	c := new(CommandLine)
	c.root = &cobra.Command{Use: "advsrv"}
	start.Start.Flags().StringVarP(&start.ConfigFilePath, "file", "f", c.defaultConfigPath(), "path to the config file")
	reload.Flags().StringVarP(&start.ConfigFilePath, "file", "f", c.defaultConfigPath(), "path to the config file")
	restart.Flags().StringVarP(&start.ConfigFilePath, "file", "f", c.defaultConfigPath(), "path to the config file")

	return c
}

func (c *CommandLine) Init() error {
	c.root.AddCommand(start.Start, stop, reload, restart, reloadCache, cleanCache, generate.GenConf, version)
	return c.root.Execute()
}

func (c *CommandLine) defaultConfigPath() string {
	switch runtime.GOOS {
	case "linux":
		return filepath.Join("/etc/advsrv", configparser.ConfigFileName+"."+configparser.ConfigFileFormat)
	default:
		exe, err := os.Executable()
		if err != nil {
			exe, _ = os.Getwd()
		}
		return filepath.Join(filepath.Dir(exe), configparser.ConfigFileName+"."+configparser.ConfigFileFormat)
	}
}
