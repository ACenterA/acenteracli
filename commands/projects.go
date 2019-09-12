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
	"github.com/wallix/awless/cli"
	"github.com/wallix/awless/config"
	"github.com/wallix/awless/console"
	"github.com/wallix/awless/graph"
	"github.com/wallix/awless/graph/outputs"
	"github.com/wallix/awless/logger"
)

var (
	defaultProjectsColumns = []string{"Id", "Name"}
	dbType                 = "mysql"
	dbName                 = ""
)

func init() {
	RootCmd.AddCommand(projectCmd)

	listProjectCobraCmd.PersistentFlags().StringSliceVar(&listingColumnsFlag, "columns", []string{}, "Select the properties to display in the columns. Ex: --columns id,name")
	listProjectCobraCmd.PersistentFlags().BoolVar(&noHeadersFlag, "no-headers", false, "Do not display headers")
	listProjectCobraCmd.PersistentFlags().BoolVar(&reverseFlag, "reverse", false, "Use in conjunction with --sort to reverse sort")
	listProjectCobraCmd.PersistentFlags().StringSliceVar(&sortBy, "sort", []string{"Id"}, "Sort tables by column(s) name(s)")

	cobra.EnableCommandSorting = false

	// Global parametes
	projectCmd.PersistentFlags().StringSliceVar(&listingFiltersFlag, "filter", []string{}, "Filter resources given key/values fields (case insensitive). Ex: --filter name='Projct 1'")

	createProjectDbServeCobraCmd.PersistentFlags().StringVar(&dbName, "name", "mysql", "Name for the Database resources Ex: --name='MySQL01'")
	// createProjectDbServeCobraCmd.PersistentFlags().StringSliceVar(&dbType, "type", string, "Type of the Database resources Ex: --type='mysql'")

	projectCmd.AddCommand(listProjectCobraCmd)
	projectCmd.AddCommand(projectSelectCobraCmd)
	projectCmd.AddCommand(createProjectDbServeCobraCmd)
}

var projectCmd = &cobra.Command{
	Use:               "project",
	Aliases:           []string{"prj"},
	Example:           "  acentera project list --sort creation\n  acentera project list --output csv\n  acentera project list --filter state=active,type=xxxxx\n  \n  acentera project select --filter name='My Project 1'",
	PersistentPreRun:  applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook, normalizeColumns),
	PersistentPostRun: applyHooks(verifyNewVersionHook, onVersionUpgrade, networkMonitorHook),
	Short:             "Project resources: sorting, filtering via tag/properties, output formatting, etc...",
}
var projectSelectCobraCmd = &cobra.Command{
	Use:   "select",
	Short: "Select an project",
	Run:   projectSelectResource,
}

var listProjectCobraCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run:   projectListResource,
}

var createProjectDbServeCobraCmd = &cobra.Command{
	Use:   "create-database-server",
	Short: "Create a database server for the project",
	Run:   projectCreateDatabase,
}

var listProjectDbServeCobraCmd = &cobra.Command{
	Use:   "list-database-server",
	Short: "List database servers for the project",
	Run:   projectListDatabase,
}

func projectListResource(*cobra.Command, []string) {
	projects, _ := cli.API().Projects().GetProjects()

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
			console.WithRdfType("project"),
			console.WithColumns(listingColumnsFlag),
			console.WithFilters(listingFiltersFlag),
			console.WithTagFilters(listingTagFiltersFlag),
			console.WithTagKeyFilters(listingTagKeyFiltersFlag),
			console.WithTagValueFilters(listingTagValueFiltersFlag),

			console.WithFormat(listingFormat),
			console.WithidsOnly(listOnlyids),
			console.WithSortBy(sortBy...),

		).SetSource(projects).Build()
	*/
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))

}

func projectSelectResource(*cobra.Command, []string) {
	projects, _ := cli.API().Projects().GetProjects()

	columns := defaultProjectsColumns

	g := graph.NewGraph()
	for id, r := range projects {
		// projectInfo := r.(map[string]interface{})[id]
		// fmt.Println("1 -- PRJ INFO:", r)
		res1 := outputs.New("project", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}

	// Build Query
	buildQuery := console.BuildOptions(
		console.WithRdfType("project"),
		console.WithColumns(columns),
		console.WithFilters(listingFiltersFlag),
	)
	a, _ := buildQuery.BuildQuery()
	r, _ := g.Find(a)
	nbItems := len(r)

	// Reset
	config.Set("user.project.id", "")
	config.Set("user.project.name", "")
	if nbItems > 1 {
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitWithError(displayer.Print(os.Stdout), "Filters returned more than one record. Please refine query.")
	} else if nbItems <= 0 {
		exitWithError(nil, "Could not find any project that match filters. Please update your filter and try again.")
	} else if nbItems == 1 {
		// All Good
		config.Set("user.project.id", r[0].Id())
		config.Set("user.project.name", r[0].Name())
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitOn(displayer.Print(os.Stdout))
		logger.Infof("Successfully selected project") // , r[0].Name())
	} else {
		exitWithError(nil, "Invalid command ...")
	}
}

func projectCreateDatabase(*cobra.Command, []string) {
	fmt.Println("CEATE DB HERE")
	projects := cli.API().Projects()
	fmt.Println("CEATE DB HERE 1")
	a, ez := projects.CreatDatabase(config.GetProjectId(), dbName, dbType)
	fmt.Println("CEATE DB HERE 2")
	logger.Infof("Test crete projects", a, ez) // , r[0].Name())
	fmt.Println("CEATE DB HERE 3")
	/*} else {
		exitWithError(nil, "Invalid command ...")
	}*/
}

func projectListDatabase(*cobra.Command, []string) {
	fmt.Println("LIST DB HERE")
	/*
		projects := cli.API().Projects()
		fmt.Println("CEATE DB HERE 1")

			a, ez := projects.ListDatabase(config.GetProjectId())
			fmt.Println("CEATE DB HERE 2")
			logger.Infof("Test crete projects", a, ez) // , r[0].Name())
	*/
	fmt.Println("CEATE DB HERE 3")
	/*} else {
		exitWithError(nil, "Invalid command ...")
	}*/
}
