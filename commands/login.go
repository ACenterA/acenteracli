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
	"github.com/spf13/cobra"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/logger"
)

func init() {
	RootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:              "login",
	Short:            "Login to your account",
	PersistentPreRun: applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook),
	Run: func(cmd *cobra.Command, args []string) {
		login(cmd, args)
		// Do not check for updates ... config.CheckForUpdate()
		return
	},
}

func login(*cobra.Command, []string) {
	var username string
	var pass string
	config.Set("_token", "")
	config.AskUserPassword(&username, &pass)
	if validateToken() {
		logger.Infof("Successfully logged in.")
	} else {
		logger.Infof("Could not login to server. Please validate your username and password")
	}
}
