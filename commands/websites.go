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

	"github.com/spf13/cobra"
	"github.com/wallix/awless/cli"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/console"
	"github.com/wallix/awless/global"
	"github.com/wallix/awless/graph"
	"github.com/wallix/awless/graph/outputs"
	"github.com/wallix/awless/logger"
)

var (
	defaultWebsitesColumns = []string{}
	WebsiteId              = ""
	DBNameTmp              = ""
	DBPrefixTmp            = ""
	UploadDirTmp           = ""
)

func init() {
	defObject := cli.WebsiteApiObject()
	ptrObj := global.To_struct_ptr(defObject)
	defaultWebsitesColumns = global.PrintFields(ptrObj)
	RootCmd.AddCommand(websiteCmd)

	listWebsiteCobraCmd.PersistentFlags().StringSliceVar(&listingColumnsFlag, "columns", []string{}, "Select the properties to display in the columns. Ex: --columns id,name")
	listWebsiteCobraCmd.PersistentFlags().BoolVar(&noHeadersFlag, "no-headers", false, "Do not display headers")
	listWebsiteCobraCmd.PersistentFlags().BoolVar(&reverseFlag, "reverse", false, "Use in conjunction with --sort to reverse sort")
	listWebsiteCobraCmd.PersistentFlags().StringSliceVar(&sortBy, "sort", []string{"Id"}, "Sort tables by column(s) name(s)")

	cobra.EnableCommandSorting = false

	getWebsiteCobraCmd.PersistentFlags().StringVar(&WebsiteId, "website", "", "Website Id")

	// Global parametes
	websiteCmd.PersistentFlags().StringSliceVar(&listingFiltersFlag, "filter", []string{}, "Filter resources given key/values fields (case insensitive). Ex: --filter name='Projct 1'")
	websiteCmd.AddCommand(listWebsiteCobraCmd)

	websiteCmd.AddCommand(getWebsiteCobraCmd)
	websiteCmd.AddCommand(websiteSelectCobraCmd)

	createSimpleCobraCmd.PersistentFlags().StringVar(&BluePrintId, "blueprintid", "", "Blueprint Id")
	createSimpleCobraCmd.PersistentFlags().StringVar(&GitRepoName, "name", "", "Website Short Name")
	createSimpleCobraCmd.PersistentFlags().StringVar(&DBNameTmp, "dbname", "", "DB Short Name ie: prod_team_shortname")
	createSimpleCobraCmd.PersistentFlags().StringVar(&DBPrefixTmp, "dbprefix", "wp_", "Enter the DB Prefix")
	createSimpleCobraCmd.PersistentFlags().StringVar(&UploadDirTmp, "uploaddir", "", "Enter an Upload dir ie: prod_XXXX")
	createSimpleCobraCmd.PersistentFlags().StringVar(&DatbaseServerId, "database", "", "Database ServerId to create the website databaes. See `database list-servers`")
	// createSimpleCobraCmd.PersistentFlags().StringVar(&gitTeamName, "team", "", "Git Team name")
	createSimpleCobraCmd.PersistentFlags().StringVar(&gitDisplayName, "description", "", "Git Short description")
	websiteCmd.AddCommand(createSimpleCobraCmd)
}

var websiteCmd = &cobra.Command{
	Use:               "website",
	Aliases:           []string{"prj"},
	Example:           "  acentera website list --sort creation\n  acentera website list --output csv\n  acentera website list --filter state=active,type=xxxxx\n  \n  acentera website select --filter name='My Website 1'",
	PersistentPreRun:  applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook, normalizeColumns),
	PersistentPostRun: applyHooks(verifyNewVersionHook, onVersionUpgrade, networkMonitorHook),
	Short:             "Website resources: sorting, filtering via tag/properties, output formatting, etc...",
}
var websiteSelectCobraCmd = &cobra.Command{
	Use:   "select",
	Short: "Select an website",
	Run:   websiteSelectResource,
}

var listWebsiteCobraCmd = &cobra.Command{
	Use:   "list",
	Short: "List websites",
	Run:   websiteListResource,
}

var createSimpleCobraCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Website without Git repository using Blueprints",
	Run:   gitWebsiteWithBlueprintAndWithoutGitResource,
}

var getWebsiteCobraCmd = &cobra.Command{
	Use:   "get",
	Short: "Get website info",
	Run:   websiteGetbyIdResource,
}

func websiteGetbyIdResource(*cobra.Command, []string) {
	if WebsiteId == "" {
		logger.Error("Missing --website parameter")
		return
	}
	websites, _ := cli.API().Websites().GetWebsitesById(config.GetProjectId(), WebsiteId)
	// fmt.Println("Test", websites)

	columns := defaultWebsitesColumns
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	g := graph.NewGraph()
	res1Tmp := outputs.New("website", websites.Id).Object1(reflect.ValueOf(*websites).Interface()) // .Build()

	/*
		jsonData, err := json.Marshal(websites.Stages)
		if err != nil {
		}
	*/

	theStage := ""
	firstStage := ""
	for r, s := range websites.Stages {
		if s.Enabled != "" {
			if firstStage == "" {
				firstStage = r
			}
			if s.Stage == "master" {
				theStage = r
			}
		}
	}

	res1 := res1Tmp.Prop("Name", fmt.Sprintf("%v", websites.Name)).Prop("Project", websites.Project).Prop("Status", websites.Status).Prop("Code", websites.Code).Prop("ACenterAType", websites.ACenterAType).Prop("Stage", theStage).Build()
	// .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
	g.AddResource(
		res1,
	)
	displayer, err := console.BuildOptions(
		console.WithRdfType("website"),
		console.WithColumns(columns),
		console.WithFormat(listingFormat),
		console.WithReverseSort(reverseFlag),
		console.WithSortBy(sortBy...),
	).SetSource(g).Build()
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))
	return
}

func websiteListResource(*cobra.Command, []string) {
	websites, _ := cli.API().Websites().GetWebsites(config.GetProjectId())

	// fmt.Println("Received of ", websites)
	columns := defaultWebsitesColumns
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	g := graph.NewGraph()
	for id, r := range websites {
		// websiteInfo := r.(map[string]interface{})[id]
		// fmt.Println("1fz -- WWS INFO:", r)
		// res1 := outputs.New("website", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		res1Tmp := outputs.New("website", id).Object(reflect.ValueOf(r).Interface()) // .Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		// res1 := res1Tmp.Prop("Name", fmt.Sprintf("%v", r.Name)).Prop("Status", fmt.Sprintf("%v", r.Status)).Build()
		res1 := res1Tmp.Prop("Name", fmt.Sprintf("%v", r.Name)).Prop("Project", fmt.Sprintf("%v", r.Project)).Build()
		// res1 := outputs.New("website", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		// res1 := outputs.New("website", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Prop("ZState", fmt.Sprintf("%v", "OK")).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()

		// .Prop("Status", fmt.Sprintf("%v", r.Status))

		// fmt.Println("RES IS :")
		// fmt.Println(res1)
		// res1 := outputs.New("website", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}
	// fmt.Println("COLUMS ARE :")
	// fmt.Println(columns)
	// console.WithMaxWidth(console.GetTerminalWidth()),
	displayer, err := console.BuildOptions(
		console.WithRdfType("website"),
		console.WithColumns(columns),
		console.WithFormat(listingFormat),
		console.WithReverseSort(reverseFlag),
		console.WithSortBy(sortBy...),
		// console.WithNoHeaders(false),
	).SetSource(g).Build()
	/*
		console.WithFilters(listingFiltersFlag),
			console.WithMaxWidth(console.GetTerminalWidth()),
			console.WithNoHeaders(noHeadersFlag),
			console.WithIDsOnly(listOnlyIDs),
	*/

	/*
		fmt.Println("FORMT IS SET TO ", listingFormat)
		displayer, er := console.BuildOptions(
			console.WithMaxWidth(console.GetTerminalWidth()),
			console.WithRdfType("ff"),
			console.WithColumns(columns),
			console.WithFormat(listingFormat),
			// console.WithNoHeaders(false),
		).SetSource(g).Build()

		var w bytes.Buffer
		displayer, err := console.BuildOptions(
			console.WithRdfType("website"),
			console.WithColumns(listingColumnsFlag),
			console.WithFilters(listingFiltersFlag),
			console.WithTagFilters(listingTagFiltersFlag),
			console.WithTagKeyFilters(listingTagKeyFiltersFlag),
			console.WithTagValueFilters(listingTagValueFiltersFlag),

			console.WithFormat(listingFormat),
			console.WithidsOnly(listOnlyids),
			console.WithSortBy(sortBy...),

		).SetSource(websites).Build()
	*/
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))

}

func websiteSelectResource(*cobra.Command, []string) {
	websites, _ := cli.API().Websites().GetWebsites(config.GetProjectId())

	columns := defaultWebsitesColumns

	g := graph.NewGraph()
	for id, r := range websites {
		// websiteInfo := r.(map[string]interface{})[id]
		// fmt.Println("1 -- WEBS INFO:", r)
		// res1 := outputs.New("website", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		res1 := outputs.New("website", id).Object(reflect.ValueOf(r).Interface()).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}

	// Build Query
	buildQuery := console.BuildOptions(
		console.WithRdfType("website"),
		console.WithColumns(columns),
		console.WithFilters(listingFiltersFlag),
	)
	a, _ := buildQuery.BuildQuery()
	r, _ := g.Find(a)
	nbItems := len(r)

	// Reset
	config.Set("user.website.id", "")
	config.Set("user.website.name", "")
	if nbItems > 1 {
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitWithError(displayer.Print(os.Stdout), "Filters returned more than one record. Please refine query.")
	} else if nbItems <= 0 {
		exitWithError(nil, "Could not find any website that match filters. Please update your filter and try again.")
	} else if nbItems == 1 {
		// All Good
		config.Set("user.website.id", r[0].Id())
		config.Set("user.website.name", r[0].Name())
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitOn(displayer.Print(os.Stdout))
		logger.Infof("Successfully selected website") // , r[0].Name())
	} else {
		exitWithError(nil, "Invalid command ...")
	}
}
func gitWebsiteWithBlueprintAndWithoutGitResource(*cobra.Command, []string) {
	//RepositoryOptions
	if GitRepoName == "" {
		logger.Error("Missing --name parameter")
		return
	}
	if DBNameTmp == "" {
		logger.Error("Missing --dbname parameter")
		return
	}
	if gitDisplayName == "" {
		logger.Error("Missing --description parameter")
		return
	}
	if DatbaseServerId == "" {
		logger.Error("Missing --database parameter")
		d, _ := cli.API().Databases().GetDatabases(config.GetProjectId())
		DatbaseServerId = d[0].Id
	}

	proj := config.GetProjectId()
	fmt.Println(fmt.Sprintf("Project: %s, Will create website [%s] with Description %s - BlueprintId: %s on Database %s", proj, GitRepoName, gitDisplayName, BluePrintId, DatbaseServerId))
	cli.API().Websites().CreateSiteWithBlueprintAndDbWithoutGit(GitRepoName, gitDisplayName, BluePrintId, DatbaseServerId, DBNameTmp, DBPrefixTmp, UploadDirTmp)
}
