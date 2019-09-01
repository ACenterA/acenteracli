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
	gentleman "gopkg.in/h2non/gentleman.v2"

//	awsservices "github.com/wallix/awless/aws/services"
//	"github.com/wallix/awless/global"
	config "github.com/wallix/awless/config"
)

var (
	// Client makes HTTP requests and parses the responses.
	Client *gentleman.Client
)

func init() {
	Client = gentleman.New()
	UserAgentMiddleware()
	LogMiddleware()
}

func InitCliEnv() error {
	/*
	Client = gentleman.New()
	UserAgentMiddleware()
	LogMiddleware()
	*/
}

