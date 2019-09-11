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
	"github.com/wallix/awless/cli"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/logger"
)

func init() {
	RootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from your account",
	// PersistentPreRun: applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook),
	PersistentPreRun: applyHooks(initAwlessEnvHookLogout, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook),
	Run: func(cmd *cobra.Command, args []string) {
		logout(cmd, args)
		// Do not check for updates ... config.CheckForUpdate()
		return
	},
}

func logout(*cobra.Command, []string) {
	token := config.GetToken()
	if token == "" {
		logger.Infof("Successfully logged out. No previous saved tokens.")
		config.Set("user.username", "")
		config.Set("_enc", "")
		return
	}
	errLogout := cli.API().Account().Logout()
	if errLogout != nil {
		logger.Error("Could not send logout for current token. We still remove the cached token from your host.")
		config.Set("_token", "")
		config.Set("user.username", "")
		config.Set("_enc", "")
	} else {
		config.Set("_token", "")
		config.Set("user.username", "")
		config.Set("_enc", "")
		logger.Infof("Successfully logged out.")
	}
	/*
		config.Set("_token", "")
		config.AskUserPassword(&username, &pass)
		if validateToken() {
			logger.Infof("Successfully logged out.")
		} else {
			logger.Infof("Could not logout to server. Please validate your username and password")
		}
	*/
}
