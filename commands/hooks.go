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
	"path/filepath"
	"strings"

	// "strings"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
	awsservices "github.com/wallix/awless/aws/services"
	"github.com/wallix/awless/cli"

	// "github.com/wallix/awless/cloud"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/database"
	"github.com/wallix/awless/global"
	"github.com/wallix/awless/logger"
	"github.com/wallix/awless/sync"
)

func applyHooks(funcs ...func(*cobra.Command, []string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		for _, fn := range funcs {
			if err := fn(cmd, args); err != nil {
				exitOn(err)
			}
		}
	}
}

func initAwlessEnvHookLogout(cmd *cobra.Command, args []string) error {
	if err := config.InitAwlessEnv(false); err != nil {
		return fmt.Errorf("cannot init awless environment: %s", err)
	}

	return applyRegionAndProfilePrecedence()
}
func initAwlessEnvHook(cmd *cobra.Command, args []string) error {
	if err := config.InitAwlessEnv(true); err != nil {
		return fmt.Errorf("cannot init awless environment: %s", err)
	}

	return applyRegionAndProfilePrecedence()
}

func toUpper(item string) string {
	return strings.ToUpper(item)
}
func toLower(item string) string {
	return strings.ToLower(item)
}

func Normalize(item string) string {
	return strings.Title(strings.ToLower(item))
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func normalizeColumns(cmd *cobra.Command, args []string) error {
	listingColumnsFlag = Map(listingColumnsFlag, Normalize)
	return nil
}

func initGitHook(cmd *cobra.Command, args []string) error {
	// fmt.Println("Provider is set to :" + GitProvider)

	tokenKey := fmt.Sprintf("_%s.token", strings.Trim(strings.ToLower(GitProvider), ""))
	// usernameKey := fmt.Sprintf("%s.user.username", strings.Trim(strings.ToLower(GitProvider), ""))
	// passKey := fmt.Sprintf("_%s.user.password", strings.Trim(strings.ToLower(GitProvider), ""))

	t, _ := config.Get(tokenKey)
	if t != nil && t != "" {
		logger.Verbosef("[%s] using token", GitProvider)
		//Ok no token
		return nil
	}
	// fmt.Println("No Token")
	var username string
	var pass string
	// resp := false
	// fmt.Println("LOWER TEST?")

	if strings.ToLower(GitProvider) == "bitbucket" {
		// fmt.Println("LOWER TEST? - 1")
		_, err := AskBitBucketUserPassword(&username, &pass)
		return err
	} else if strings.ToLower(GitProvider) == "github" {
		// AskGithubUserPassword(&username, &pass)
	}

	/*
		if res {
			// fmt.Println("Keys are:", tokenKey, usernameKey, passKey)
		}
	*/

	return nil
}

func initCliEnvHook(cmd *cobra.Command, args []string) error {
	cli.InitCliEnv()
	return nil
}

var profileOverridenThrough, regionOverridenThrough string

func applyRegionAndProfilePrecedence() error {
	if awsProfileGlobalFlag != "" {
		if err := config.SetVolatile(config.ProfileConfigKey, awsProfileGlobalFlag); err != nil {
			return err
		}
		profileOverridenThrough = "command flag"
	} else if envProfile := os.Getenv("AWS_DEFAULT_PROFILE"); envProfile != "" {
		if err := config.SetVolatile(config.ProfileConfigKey, envProfile); err != nil {
			return err
		}
		profileOverridenThrough = "AWS_DEFAULT_PROFILE variable"
	} else if envProfile := os.Getenv("AWS_PROFILE"); envProfile != "" {
		if err := config.SetVolatile(config.ProfileConfigKey, envProfile); err != nil {
			return err
		}
		profileOverridenThrough = "AWS_PROFILE variable"
	}

	profile := config.GetAWSProfile()

	if region, embedded, err := hasEmbeddedRegionInSharedConfigForProfile(profile); err == nil {
		if embedded {
			if e := config.SetVolatile(config.RegionConfigKey, region); e != nil {
				return e
			}
			regionOverridenThrough = fmt.Sprintf("profile '%s' (see AWS config files $HOME/.aws/{credentials,config})", profile)
		} else {
			regionOverridenThrough = ""
		}
	} else {
		return err
	}

	if awsRegionGlobalFlag != "" {
		if err := config.SetVolatile(config.RegionConfigKey, awsRegionGlobalFlag); err != nil {
			return err
		}
		regionOverridenThrough = "command flag"
	} else if envRegion := os.Getenv("AWS_DEFAULT_REGION"); envRegion != "" {
		if err := config.SetVolatile(config.RegionConfigKey, envRegion); err != nil {
			return err
		}
		regionOverridenThrough = "AWS_DEFAULT_REGION variable"
	}

	return nil
}

func notifyOnRegionOrProfilePrecedenceHook(*cobra.Command, []string) error {
	applyRegionAndProfilePrecedence()

	if m := profileOverridenThrough; len(m) > 0 {
		logger.Infof("profile precedence: '%s' loaded through %s", config.GetAWSProfile(), m)
	}
	if m := regionOverridenThrough; len(m) > 0 {
		logger.Infof("region precedence: '%s' loaded through %s", config.GetAWSRegion(), m)
	}

	return nil
}

func initCloudServicesHook(cmd *cobra.Command, args []string) error {
	if localGlobalFlag {
		return nil
	}

	profile, region := config.GetAWSProfile(), config.GetAWSRegion()

	// logger.Verbosef("awless %s - loading AWS session with profile '%s' and region '%s'", config.Version, profile, region)

	if err := awsservices.Init(profile, region, config.GetConfigWithPrefix("aws."), logger.DefaultLogger, config.SetProfileCallback, networkMonitorFlag); err != nil {
		return err
	}

	/*
		if config.TriggerSyncOnConfigUpdate && !strings.HasPrefix(cmd.Name(), "sync") {
			var services []cloud.Service
			for _, s := range cloud.ServiceRegistry {
				services = append(services, s)
			}
			if !noSyncGlobalFlag {
				logger.Infof("Syncing new region '%s'... (disable with --no-sync global flag)", region)
				sync.NewSyncer(logger.DefaultLogger).Sync(services...)
			}
		}
	*/

	return nil
}

func includeHookIf(cond *bool, hook func(*cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, args []string) error {
		if *cond {
			return hook(c, args)
		}
		return nil
	}
}

func initSyncerHook(cmd *cobra.Command, args []string) error {
	if noSyncGlobalFlag {
		sync.DefaultSyncer = sync.NoOpSyncer()
	} else {
		sync.DefaultSyncer = sync.NoOpSyncer()
		// sync.DefaultSyncer = sync.NewSyncer(logger.DefaultLogger)
	}
	return nil
}

func initLoggerHook(cmd *cobra.Command, args []string) error {
	var flag int
	if verboseGlobalFlag {
		flag = logger.VerboseF
		global.Verbose = "verbose"
	}
	if extraVerboseGlobalFlag {
		flag = flag | logger.ExtraVerboseF
		global.Verbose = "extra"
	}

	logger.DefaultLogger.SetVerbose(flag)
	if silentGlobalFlag {
		logger.DefaultLogger = logger.DiscardLogger
	}
	return nil
}

func onVersionUpgrade(cmd *cobra.Command, args []string) error {
	var lastVersion string
	if derr := database.Execute(func(db *database.DB) (err error) {
		lastVersion, err = db.GetStringValue("current.version")
		return
	}); derr != nil {
		fmt.Printf("cannot verify stored version in db: %s\n", derr)
	}

	if config.IsSemverUpgrade(lastVersion, config.Version) {
		if err := database.Execute(func(db *database.DB) error {
			return db.SetStringValue("current.version", config.Version)
		}); err != nil {
			fmt.Printf("cannot store upgraded version in db: %s\n", err)
		}
		migrationActionsAndExtraMessages(config.Version)
		logger.Infof("You have just upgraded acentera from %s to %s", lastVersion, config.Version)
		logger.Infof("Check out %s latest features at https://github.com/ACenterA/acenteracli/blob/master/CHANGELOG.md", config.Version)
	}

	return nil
}

func verifyNewVersionHook(cmd *cobra.Command, args []string) error {
	if localGlobalFlag {
		return nil
	}
	cliBranch := os.Getenv("ACENTERA_CLI_BRANCH")
	if cliBranch == "" {
		cliBranch = "master"
	}
	config.VerifyNewVersionAvailable(fmt.Sprintf("https://raw.githubusercontent.com/ACenterA/acenteracli/%v/VERSION", cliBranch), os.Stderr)
	return nil
}

func networkMonitorHook(cmd *cobra.Command, args []string) error {
	return nil
}

func firstInstallDoneHook(cmd *cobra.Command, args []string) error {
	if config.TriggerSyncOnConfigUpdate {
		fmt.Fprintln(os.Stderr, "\nAll done. Enjoy!")
		fmt.Fprintln(os.Stderr, "You can review and configure awless with `awless config`")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintf(os.Stderr, "Now running: `%s`\n", cmd.CommandPath())
	}
	return nil
}

func migrationActionsAndExtraMessages(current string) {
	switch current {
	case "v0.1.7":
		config.Set("instance.distro", "amazonlinux")
		logger.Info("In v0.1.7, the default template config value 'instance.image' has been deprecated in favor of 'instance.distro'")

		ami, _ := config.Get("instance.image")
		if isNotAwlessFormerDefaultAMI(fmt.Sprint(ami)) {
			logger.Warningf("\tYou had a customized value of '%s' for the now deprecated 'instance.image'", fmt.Sprint(ami))
			logger.Warning("\tThis value will not be taken into account anymore as default when running templates")
		} else {
			logger.Info("\tMigrated correctly the deprecated 'instance.image' to 'instance.distro'")
		}
		config.Unset("instance.image")
		logger.Info("\tYou can always check your config values with 'awless config'")
	case "v0.1.9":
		logger.Info("In v0.1.9, the local data file model has been moved to support multi-account transparently")
		oldData := filepath.Join(os.Getenv("__AWLESS_HOME"), "aws", "rdf")
		if err := os.RemoveAll(oldData); err == nil {
			logger.Info("-> Stale data have been removed. The local model (ex: used for completion) will progressively be synced again through your usage of awless.")
			logger.Info("-> You can also manually run `awless sync`")
		}
	}
}

func hasEmbeddedRegionInSharedConfigForProfile(profile string) (string, bool, error) {
	s, err := session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
		Profile:                 profile,
	})
	if err != nil {
		return "", false, fmt.Errorf("cannot check profile '%s' has embedded region in shared config file: %s", profile, err)
	}
	region := *s.Config.Region
	return region, len(region) > 0, nil
}

func isNotAwlessFormerDefaultAMI(s string) bool {
	amis := []string{"ami-c58c1dd3", "ami-4191b524", "ami-7a85a01a", "ami-4836a428", "ami-0bd66a6f", "ami-d3c0c4b5", "ami-b6daced2", "ami-b968bad6", "ami-fc5ae39f", "ami-762a2315", "ami-923d12f5", "ami-9d15c7f3", "ami-52c7b43d", "ami-2bccae47"}
	for _, e := range amis {
		if e == s {
			return false
		}
	}
	return true
}

func AskBitBucketUserPassword(username *string, pass *string) (bool, error) {

	prompted := false

	token := os.Getenv("ACENTERA_BITBUCKET_TOKEN")
	c := bitbucket.NewOAuthbearerToken(token)
	r, er := c.User.Profile()
	if er == nil {
		BitBucket = c
		fmt.Println("Successfully loggedin using Bitbucket Token")
		config.Set("bitbucket.user.account_id", r.(map[string]interface{})["account_id"].(string))
		ac, _ := config.Get("bitbucket.user.account_id")
		config.Set("_bitbucket.token", token)
		GitOwner = ac.(string)
		return prompted, er
	}

	if *username == "" {
		*username = config.GetGitUsername("bitbucket")
	}
	if *username == "" {
		*username = os.Getenv("ACENTERA_BITBUCKET_USERNAME")
		if *username == "" {
			if !prompted {
				fmt.Printf("\nPlease enter your bitbucket credentials.\n")
			}
			prompted = true
			config.PromptUntilNonEmpty("\nUsername: ", username)
			config.Set("bitbucket.user.username", *username)
			ac, _ := config.Get("bitbucket.user.account_id")
			GitOwner = ac.(string)
		} else {
			config.Set("bitbucket.user.username", *username)
		}
	}

	/*if oldUsername != *username {
		ResetUserSettings()
	}
	*/

	*pass = config.GetGitPassword("bitbucket")
	if *pass != "" {
		c := bitbucket.NewBasicAuth(*username, *pass)
		r, er := c.User.Profile()
		if er == nil {
			config.Set("_bitbucket.token", string(""))
			config.Set("bitbucket.user.username", *username)
			config.Set("_bitbucket.user.password", *pass)
			config.Set("bitbucket.user.account_id", r.(map[string]interface{})["account_id"].(string))
			BitBucket = c

			ac, _ := config.Get("bitbucket.user.account_id")
			GitOwner = ac.(string)
			return prompted, er
		}
		*pass = ""
	}
	if *pass == "" {
		*pass = os.Getenv("ACENTERA_BITBUCKET_PASSWORD")
		if *pass != "" {
			c := bitbucket.NewBasicAuth(*username, *pass)
			r, e := c.User.Profile()
			if e == nil {
				config.Set("_bitbucket.token", string(""))
				config.Set("bitbucket.user.username", *username)
				config.Set("_bitbucket.user.password", *pass)
				config.Set("bitbucket.user.account_id", r.(map[string]interface{})["account_id"].(string))
				BitBucket = c
				ac, _ := config.Get("bitbucket.user.account_id")
				GitOwner = ac.(string)
				return prompted, e
			}
		}
		*pass = ""

		if *pass == "" {
			if !prompted {
				fmt.Printf("\nPlease enter your bitbucket credentials.\n")
			}
			prompted = true
			config.PromptUntilNonEmptySecure("Password: ", pass)
		}
	}
	if prompted {
		config.Set("_bitbucket.token", string(""))
		c := bitbucket.NewBasicAuth(*username, *pass)
		r, ez := c.User.Profile()
		if ez == nil {
			config.Set("_bitbucket.user.password", *pass)
			config.Set("bitbucket.user.account_id", r.(map[string]interface{})["account_id"].(string))

			ac, _ := config.Get("bitbucket.user.account_id")
			GitOwner = ac.(string)

			BitBucket = c
			return prompted, ez
		} else {
			logger.Error("Could not login to bitbucket. Please try again.")
			fmt.Println(ez)
			return prompted, ez
		}
	}
	return prompted, nil
}

func AskGithubUserPassword(username *string, pass *string) bool {

	prompted := false
	oldUsername := config.GetUsername()
	if *username == "" {
		*username = os.Getenv("ACENTERA_GITHUB_USERNAME")
		if *username == "" {
			if !prompted {
				fmt.Printf("\nPlease enter your Github credentials.\n")
			}
			prompted = true
			config.PromptUntilNonEmpty("\nUsername: ", username)
			config.Set("github.user.username", *username)
		}
	}

	if oldUsername != *username {
		config.ResetUserSettings()
	}

	if *pass == "" {
		*pass = os.Getenv("ACENTERA_GITHUB_PASSWORD")
		if *pass == "" {
			if !prompted {
				fmt.Printf("\nPlease enter your github credentials.\n")
			}
			prompted = true
			config.PromptUntilNonEmptySecure("Password: ", pass)
			// enc := encrypt([]byte(*pass), global.ENC_PWD)
			// config.Set("_github.user.password", string(enc))
		}
	}
	if prompted {
		config.Set("_github.user.password", string(""))
	}
	return prompted
}
