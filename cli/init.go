/*
Copyright 2018 ACenterA

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

package cli

import (
	// "fmt"
	"strings"
        "github.com/spf13/cobra"
	gentleman "gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/url"

//	awsservices "github.com/wallix/awless/aws/services"
	global "github.com/wallix/awless/global"
	// config "github.com/wallix/awless/config"
)

var (
	// Client makes HTTP requests and parses the responses.
	Client *gentleman.Client
	AnonClient *gentleman.Client
	GitClient *gentleman.Client
	APIPrefix string
)

func init() {
	Client = gentleman.New()
	AnonClient = gentleman.New()
	GitClient = gentleman.New()
}

func InitCliEnv(cmd *cobra.Command, args []string) error {
	/*
        if localGlobalFlag {
                return nil
        }
	*/
	// Define the server url (must be first)
	Client.Use(url.URL(global.API_ENDPOINT))
	AnonClient.Use(url.URL(global.API_ENDPOINT))
	tmpEndpoint := strings.Split(global.API_ENDPOINT,".")
        idx := strings.Index(tmpEndpoint[len(tmpEndpoint)-1],"/")
        endpointPrefix := tmpEndpoint[len(tmpEndpoint)-1][idx:]
	APIPrefix = endpointPrefix

	UserAgentMiddleware(Client)
	// UserAgentMiddleware(AnonClient)
	// PathMiddleware(Client)
	// PathMiddleware(AnonClient)
	AuthorizationMiddleware(Client)
	LogMiddleware(Client, false)
	// LogMiddleware(AnonClient, false)
	// LogMiddleware(AnonClient, false)

	// fmt.Println("User pwd :" + config.GetPasswordPlainText())
	/*
	Client = gentleman.New()
	UserAgentMiddleware()
	LogMiddleware()
	*/
	return nil
}

