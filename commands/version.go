/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/logger"
)
var (
	performUpgrade bool
)

func init() {
	RootCmd.AddCommand(versionCmd)

	// versionExtraCmd.Flags().BoolVar(&performUpgrade, "upgrade", false, "Update the Client to the latest version")

	versionCmd.AddCommand(versionCheckCmd)
	versionCmd.AddCommand(versionUpgradCmd)
	versionCmd.AddCommand(versionExtraCmd)
}

var versionCmd = &cobra.Command{
        Hidden: true,
        Use:    "version",
        Short:  "Show / Upgrade client version",
}

var versionExtraCmd = &cobra.Command{
	Use:   "show",
	Short: "Show client version",

        Run: func(cmd *cobra.Command, args []string) {
                if performUpgrade {
			config.ConfirmAndSelfUpdate()
			return
		}
		printVersion(cmd, args)
		// Do not check for updates ... config.CheckForUpdate()
		return
	},
}
var versionUpgradCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade client version",

        Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprint(os.Stderr, config.AWLESS_ASCII_LOGO)
		config.ConfirmAndSelfUpdate()
		return
	},
}
var versionCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check cliet Upgrade",

        Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprint(os.Stderr, config.AWLESS_ASCII_LOGO)
		os.Exit(config.CheckForUpdateExists())
		return
	},
}

func printVersion(*cobra.Command, []string) {
	fmt.Fprint(os.Stderr, config.AWLESS_ASCII_LOGO)
        logger.Infof("Current version %s", config.Version)
	// fmt.Println(config.CurrentBuildInfo)
}
