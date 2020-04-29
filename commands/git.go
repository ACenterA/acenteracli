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
	"reflect"
	"strings"

	bitbucket "github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
	"github.com/wallix/awless/cli"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/console"
	"github.com/wallix/awless/graph"
	"github.com/wallix/awless/graph/outputs"
	"github.com/wallix/awless/logger"
)

var (
	GitProvider    string
	BitBucket      *bitbucket.Client
	GitRepoName    string
	gitTeamName    string
	gitDisplayName string
	GitOwner       string
	BluePrintId    string
	//GitHub github
)

func init() {
	RootCmd.AddCommand(gitCmd)

	/*
		listProjectCobraCmd.PersistentFlags().StringSliceVar(&listingColumnsFlag, "columns", []string{}, "Select the properties to display in the columns. Ex: --columns id,name")
		listProjectCobraCmd.PersistentFlags().BoolVar(&noHeadersFlag, "no-headers", false, "Do not display headers")
		listProjectCobraCmd.PersistentFlags().BoolVar(&reverseFlag, "reverse", false, "Use in conjunction with --sort to reverse sort")
		listProjectCobraCmd.PersistentFlags().StringSliceVar(&sortBy, "sort", []string{"Id"}, "Sort tables by column(s) name(s)")
	*/
	cobra.EnableCommandSorting = false

	// Global parametes
	// projectCmd.PersistentFlags().StringSliceVar(&listingFiltersFlag, "filter", []string{}, "Filter resources given key/values fields (case insensitive). Ex: --filter name='Projct 1'")

	// createProjectDbServeCobraCmd.PersistentFlags().StringVar(&dbName, "name", "mysql", "Name for the Database resources Ex: --name='MySQL01'")
	// createProjectDbServeCobraCmd.PersistentFlags().StringSliceVar(&dbType, "type", string, "Type of the Database resources Ex: --type='mysql'")

	gitCmd.PersistentFlags().StringVar(&GitProvider, "provider", "bitbucket", "Provider to use github/gitlab/bitbucket (default: bitbucket)")

	createBitbucketCobraCmd.PersistentFlags().StringVar(&BluePrintId, "blueprintid", "", "Blueprint Id")

	createBitbucketCobraCmd.PersistentFlags().StringVar(&GitRepoName, "name", "", "Repository name")
	createBitbucketCobraCmd.PersistentFlags().StringVar(&gitTeamName, "team", "", "Git Team name")
	createBitbucketCobraCmd.PersistentFlags().StringVar(&gitDisplayName, "description", "", "Git Short description")

	createBitbucketBlueprintCobraCmd.PersistentFlags().StringVar(&GitRepoName, "name", "", "Repository name")
	createBitbucketBlueprintCobraCmd.PersistentFlags().StringVar(&gitTeamName, "team", "", "Git Team name")
	createBitbucketBlueprintCobraCmd.PersistentFlags().StringVar(&gitDisplayName, "description", "", "Git Short description")

	listBitbucketCobraCmd.PersistentFlags().StringVar(&gitTeamName, "team", "", "Git Team name")

	gitCmd.AddCommand(listBitbucketCobraCmd)
	gitCmd.AddCommand(logoutGitCobraCmd)
	gitCmd.AddCommand(createBitbucketCobraCmd)
	gitCmd.AddCommand(createBitbucketBlueprintCobraCmd)
	// projectCmd.AddCommand(createProjectDbServeCobraCmd)
}

var gitCmd = &cobra.Command{
	Use:               "git",
	Aliases:           []string{"git"},
	Example:           "  acentera git list --provider github\n",
	PersistentPreRun:  applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook, normalizeColumns, initGitHook),
	PersistentPostRun: applyHooks(verifyNewVersionHook, onVersionUpgrade, networkMonitorHook),
	Short:             "Git resources: create,list and manage repositories from Github/Bitbucket/Gitlab ...",
}

var listBitbucketCobraCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories",
	Run:   gitBitBucketListResource,
}

var createBitbucketCobraCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Website and repository with blueprint",
	Run:   gitCreateBitBucketResource,
}

var createBitbucketBlueprintCobraCmd = &cobra.Command{
	Use:   "create-blueprint",
	Short: "Create Blueprint and repository",
	Run:   gitCreateBlueprintBitBucketResource,
}

var logoutGitCobraCmd = &cobra.Command{
	Use:              "logout",
	Short:            "Remove stored authentication info from cache",
	PersistentPreRun: applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook, normalizeColumns),
	Run:              gitLogoutResource,
}

func gitLogoutResource(*cobra.Command, []string) {
	config.ResetGitCreds(strings.ToLower(GitProvider))
}

func gitBitBucketListResource(*cobra.Command, []string) {

	// fmt.Println("OK GOT PROVIDER of :", GitProvider)
	/*projects, _ := cli.API().Projects().GetProjects()

	columns := defaultProjectsColumns
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	g := graph.NewGraph()
	for id, r := range projects {
		// projectInfo := r.(map[string]interface{})[id]
		// fmt.Println("1 -- PRJ INFO:", r)
		res1 := outputs.New("project", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}
	displayer, err := console.BuildOptions(
		console.WithMaxWidth(console.GetTerminalWidth()),
		console.WithRdfType("project"),
		console.WithColumns(columns),
		console.WithFormat(listingFormat),
		console.WithReverseSort(reverseFlag),
		console.WithFilters(listingFiltersFlag),
		console.WithMaxWidth(console.GetTerminalWidth()),
		console.WithNoHeaders(noHeadersFlag),
		console.WithIDsOnly(listOnlyIDs),
		console.WithSortBy(sortBy...),
		// console.WithNoHeaders(false),
	).SetSource(g).Build()
	*/

	c := BitBucket
	//fmt.Println("DID WE GOT USER?")
	//fmt.Println(c.User)
	f, errz := c.User.Profile()
	if errz == nil {

		accountId := f.(map[string]interface{})["account_id"].(string) // or team ?

		teamNameOrOwner := gitTeamName
		if teamNameOrOwner == "" {
			teamNameOrOwner = accountId
		}
		opt := &bitbucket.RepositoriesOptions{
			Owner: teamNameOrOwner, // f.(map[string]interface{})["account_id"].(string), // or team ?
		}

		r, err := c.Repositories.ListForAccount(opt)
		if err != nil {
			panic(err)
			return
		}
		// fmt.Println(r)

		/* ?? pages ...
				Page    int32
		        Pagelen int32
		        Size    int32
				Items   []Repository
		*/

		g := graph.NewGraph()
		for _, r := range r.Items {
			res1 := outputs.New("repositories", fmt.Sprintf("%s", r.Slug)).Prop("Id", fmt.Sprintf("%v", r.Slug)).Object1(reflect.ValueOf(r).Interface()).Prop("Full_name", fmt.Sprintf("%v", r.Full_name)).Build()
			g.AddResource(
				res1,
			)
		}

		columns := []string{"Id", "Slug", "Full_name"}
		displayer, err := console.BuildOptions(
			console.WithRdfType("repositories"),
			console.WithColumns(columns),
			console.WithFormat(listingFormat),
			console.WithReverseSort(reverseFlag),
			console.WithSortBy(sortBy...),
			// console.WithNoHeaders(false),
		).SetSource(g).Build()
		exitOn(err)

		exitOn(displayer.Print(os.Stdout))

	}
}

func gitCreateBitBucketResource(*cobra.Command, []string) {
	//RepositoryOptions
	if GitRepoName == "" {
		logger.Error("Missing --name parameter")
		return
	}

	/*
		Owner    string `json:"owner"`
		RepoSlug string `json:"repo_slug"`
		Scm      string `json:"scm"`
		//      Name        string `json:"name"`
		IsPrivate   string `json:"is_private"`
		Description string `json:"description"`
		ForkPolicy  string `json:"fork_policy"`
		Language    string `json:"language"`
		HasIssues   string `json:"has_issues"`
		HasWiki     string `json:"has_wiki"`
		Project     string `json:"project"`
	*/

	c := BitBucket
	teamNameOrOwner := gitTeamName
	if teamNameOrOwner == "" {
		teamNameOrOwner = GitOwner
	}
	displayNameOrName := gitDisplayName
	if displayNameOrName == "" {
		displayNameOrName = GitRepoName
	}

	opt := &bitbucket.RepositoryOptions{
		Owner:       teamNameOrOwner,
		RepoSlug:    GitRepoName,
		IsPrivate:   "true",
		Description: fmt.Sprintf("[ACenterA] - %s", displayNameOrName),
		Scm:         "git",
	}
	res, err := c.Repositories.Repository.Create(opt)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// fmt.Println(res.Owner)
	// fmt.Println("Create of ", GitRepoName)
	// fmt.Println(res.Links)
	// fmt.Println(res.Links["avatar"].(map[string]interface{})["href"])
	data := strings.Split(res.Links["avatar"].(map[string]interface{})["href"].(string), "/")
	uuidInfoTmp := data[len(data)-1]
	uuidInfoTmpArr := strings.Split(uuidInfoTmp, "?")
	uuidInfo := uuidInfoTmpArr[0][3 : len(uuidInfoTmpArr[0])-3]
	// GOT UUID :  uuidInfo

	cloneHttpsArr := res.Links["clone"].([]interface{})
	sshRepo := ""
	httpRepo := ""
	for _, t := range cloneHttpsArr {
		//		fmt.Println("A:", a)
		//		fmt.Println("T:", t)
		t1 := t.(map[string]interface{})
		hrefTmp := t1["href"].(string)[0 : len(t1["href"].(string))-4]
		isSsh := false
		if t1["name"] == "https" {
			isSsh = false
		} else if t1["name"] == "ssh" {
			isSsh = true
		}
		//		fmt.Println("Its an repo of :", hrefTmp)
		if isSsh {
			sshRepo = hrefTmp
		} else {
			httpRepo = hrefTmp
		}
	}
	// repository = https://api.bitbucket.org/2.0/repositories/
	/*
		fmt.Println("SSH REPO:", sshRepo)
		fmt.Println("HTTP REPO:", httpRepo)
		fmt.Println("Repository Name Path ..:", fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v", res.Full_name))
		fmt.Println("GOT UUID", uuidInfo)
	*/

	if BluePrintId == "" {
		proj, _ := cli.API().Websites().CreateSite(GitRepoName, uuidInfo, fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v", res.Full_name), sshRepo, httpRepo, displayNameOrName)
		fmt.Println(proj)
	} else {
		// fmt.Println("CREATE SITE USING BLUEPRINTID:", BluePrintId)
		proj, _ := cli.API().Websites().CreateSiteWithBlueprint(GitRepoName, uuidInfo, fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v", res.Full_name), sshRepo, httpRepo, displayNameOrName, BluePrintId)
		fmt.Println(proj)
	}
	/*
			Its an repo of : git@bitbucket.org:zizani-dev/f1
		SSH REPO: git@bitbucket.org:zizani-dev/f1
		HTTP REPO: https://zizacom@bitbucket.org/zizani-dev/f1
		Repository Name Path ..: https://api.bitbucket.org/2.0/repositories/zizani-dev/f1


		{
		    "projectId":"10744400d19811e992d8910affe1eb9e",

		    "ssh_repository":"bitbucket.org:zizacom/wwwmetslatvauwebcom.git",
		    "http_repository":"https://bitbucket.org/zizacom/wwwmetslatvauwebcom",
		    "repository":"https://api.bitbucket.org/2.0/repositories/zizacom/wwwmetslatvauwebcom",

		    "repodetails":{
		           "uuid":"c3cafb61-dd33-4b83-a298-ed4377108180",
		           "name":"wwwmetslatvauwebcom"
		    },

		    "refresh":"ygwTju7T8bzDbVzxDL", // ?????????

		    "branch":"master",
		    "title":"www.metslatvauweb.com",
		    "type":"git",
		    "acentera_type":"docker-simple",

		    "auth":"Ubm_dEMswVNdhUXKKgM9zIFeTmdEmtnBZ6HpfjtyvdWywskM_VZjVp_e_A5qAAG-fChNK_bcV-8PxwDzNO8=",
		    "token":"Ubm_dEMswVNdhUXKKgM9zIFeTmdEmtnBZ6HpfjtyvdWywskM_VZjVp_e_A5qAAG-fChNK_bcV-8PxwDzNO8=",
		}

		'
	*/
	// cloneHttps := strings.Split(cloneHttpsArr, ".")[0]
	// fmt.Println("GOT HTTPS :", cloneHttps)

	/*
		// map[
		  // branches:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/refs/branches]
		  //
		  clone:[
		  	map[href:https://zizacom@bitbucket.org/zizani-dev/ff1.git name:https]
		  	map[href:git@bitbucket.org:zizani-dev/ff1.git name:ssh]
		  ]

		  // commits:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/commits] downloads:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/downloads] forks:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/forks] hooks:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/hooks] html:map[href:https://bitbucket.org/zizani-dev/ff1] pullrequests:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/pullrequests] self:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1] source:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/src] tags:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/refs/tags] watchers:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/watchers]]
	*/

}

func gitCreateBlueprintBitBucketResource(*cobra.Command, []string) {
	//RepositoryOptions
	if GitRepoName == "" {
		logger.Error("Missing --name parameter")
		return
	}

	/*
		Owner    string `json:"owner"`
		RepoSlug string `json:"repo_slug"`
		Scm      string `json:"scm"`
		//      Name        string `json:"name"`
		IsPrivate   string `json:"is_private"`
		Description string `json:"description"`
		ForkPolicy  string `json:"fork_policy"`
		Language    string `json:"language"`
		HasIssues   string `json:"has_issues"`
		HasWiki     string `json:"has_wiki"`
		Project     string `json:"project"`
	*/

	c := BitBucket
	teamNameOrOwner := gitTeamName
	if teamNameOrOwner == "" {
		teamNameOrOwner = GitOwner
	}
	displayNameOrName := gitDisplayName
	if displayNameOrName == "" {
		displayNameOrName = GitRepoName
	}

	opt := &bitbucket.RepositoryOptions{
		Owner:       teamNameOrOwner,
		RepoSlug:    GitRepoName,
		IsPrivate:   "true",
		Description: fmt.Sprintf("[ACenterA] - %s", displayNameOrName),
		Scm:         "git",
	}
	res, err := c.Repositories.Repository.Create(opt)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// fmt.Println(res.Owner)
	// fmt.Println("Create of ", GitRepoName)
	// fmt.Println(res.Links)
	// fmt.Println(res.Links["avatar"].(map[string]interface{})["href"])
	data := strings.Split(res.Links["avatar"].(map[string]interface{})["href"].(string), "/")
	uuidInfoTmp := data[len(data)-1]
	uuidInfoTmpArr := strings.Split(uuidInfoTmp, "?")
	uuidInfo := uuidInfoTmpArr[0][3 : len(uuidInfoTmpArr[0])-3]
	// GOT UUID :  uuidInfo

	cloneHttpsArr := res.Links["clone"].([]interface{})
	sshRepo := ""
	httpRepo := ""
	for _, t := range cloneHttpsArr {
		//		fmt.Println("A:", a)
		//		fmt.Println("T:", t)
		t1 := t.(map[string]interface{})
		hrefTmp := t1["href"].(string)[0 : len(t1["href"].(string))-4]
		isSsh := false
		if t1["name"] == "https" {
			isSsh = false
		} else if t1["name"] == "ssh" {
			isSsh = true
		}
		//		fmt.Println("Its an repo of :", hrefTmp)
		if isSsh {
			sshRepo = hrefTmp
		} else {
			httpRepo = hrefTmp
		}
	}
	// repository = https://api.bitbucket.org/2.0/repositories/
	/*
		fmt.Println("SSH REPO:", sshRepo)
		fmt.Println("HTTP REPO:", httpRepo)
		fmt.Println("Repository Name Path ..:", fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v", res.Full_name))
		fmt.Println("GOT UUID", uuidInfo)
	*/

	proj, _ := cli.API().Blueprints().CreateBlueprint(GitRepoName, uuidInfo, fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v", res.Full_name), sshRepo, httpRepo, displayNameOrName)
	fmt.Println(proj)
	/*
			Its an repo of : git@bitbucket.org:zizani-dev/f1
		SSH REPO: git@bitbucket.org:zizani-dev/f1
		HTTP REPO: https://zizacom@bitbucket.org/zizani-dev/f1
		Repository Name Path ..: https://api.bitbucket.org/2.0/repositories/zizani-dev/f1


		{
		    "projectId":"10744400d19811e992d8910affe1eb9e",

		    "ssh_repository":"bitbucket.org:zizacom/wwwmetslatvauwebcom.git",
		    "http_repository":"https://bitbucket.org/zizacom/wwwmetslatvauwebcom",
		    "repository":"https://api.bitbucket.org/2.0/repositories/zizacom/wwwmetslatvauwebcom",

		    "repodetails":{
		           "uuid":"c3cafb61-dd33-4b83-a298-ed4377108180",
		           "name":"wwwmetslatvauwebcom"
		    },

		    "refresh":"ygwTju7T8bzDbVzxDL", // ?????????

		    "branch":"master",
		    "title":"www.metslatvauweb.com",
		    "type":"git",
		    "acentera_type":"docker-simple",

		    "auth":"Ubm_dEMswVNdhUXKKgM9zIFeTmdEmtnBZ6HpfjtyvdWywskM_VZjVp_e_A5qAAG-fChNK_bcV-8PxwDzNO8=",
		    "token":"Ubm_dEMswVNdhUXKKgM9zIFeTmdEmtnBZ6HpfjtyvdWywskM_VZjVp_e_A5qAAG-fChNK_bcV-8PxwDzNO8=",
		}

		'
	*/
	// cloneHttps := strings.Split(cloneHttpsArr, ".")[0]
	// fmt.Println("GOT HTTPS :", cloneHttps)

	/*
		// map[
		  // branches:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/refs/branches]
		  //
		  clone:[
		  	map[href:https://zizacom@bitbucket.org/zizani-dev/ff1.git name:https]
		  	map[href:git@bitbucket.org:zizani-dev/ff1.git name:ssh]
		  ]

		  // commits:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/commits] downloads:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/downloads] forks:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/forks] hooks:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/hooks] html:map[href:https://bitbucket.org/zizani-dev/ff1] pullrequests:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/pullrequests] self:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1] source:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/src] tags:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/refs/tags] watchers:map[href:https://api.bitbucket.org/2.0/repositories/zizani-dev/ff1/watchers]]
	*/

}
