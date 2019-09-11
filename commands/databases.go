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
	defaultDatabasesColumns = []string{}
	DatbaseName             = ""
	DatbaesStage            = ""
	DatbaseServerId         = ""
	WSId                    = ""
	DatabaseProject         = ""
)

func init() {
	defObject := cli.DatabaseApiObject()
	ptrObj := global.To_struct_ptr(defObject)
	defaultDatabasesColumns = global.PrintFields(ptrObj)
	RootCmd.AddCommand(databaseCmd)

	// listDatabaseCobraCmd.PersistentFlags().StringSliceVar(&listingColumnsFlag, "columns", []string{}, "Select the properties to display in the columns. Ex: --columns id,name")
	// listDatabaseCobraCmd.PersistentFlags().BoolVar(&noHeadersFlag, "no-headers", false, "Do not display headers")
	// listDatabaseCobraCmd.PersistentFlags().BoolVar(&reverseFlag, "reverse", false, "Use in conjunction with --sort to reverse sort")
	// listDatabaseCobraCmd.PersistentFlags().StringSliceVar(&sortBy, "sort", []string{"Id"}, "Sort tables by column(s) name(s)")

	databaseCreateWStageCobraCmd.PersistentFlags().StringVar(&DatbaseName, "name", "", "Database name")
	databaseCreateWStageCobraCmd.PersistentFlags().StringVar(&DatbaesStage, "stage", "", "Stage Id")
	databaseCreateWStageCobraCmd.PersistentFlags().StringVar(&DatabaseProject, "project", "", "Project Id")
	databaseCreateWStageCobraCmd.PersistentFlags().StringVar(&DatbaseServerId, "database", "", "Database Server Id")
	databaseCreateWStageCobraCmd.PersistentFlags().StringVar(&WSId, "website", "", "Website Id")

	cobra.EnableCommandSorting = false

	// Global parametes
	// databaseCmd.PersistentFlags().StringSliceVar(&listingFiltersFlag, "filter", []string{}, "Filter resources given key/values fields (case insensitive). Ex: --filter name='Projct 1'")
	listDatabaseDBCobraCmd.PersistentFlags().StringSliceVar(&listingFiltersFlag, "filter", []string{}, "Filter resources given key/values fields (case insensitive). Ex: --filter name='db_name'")
	listDatabaseDBCobraCmd.PersistentFlags().StringVar(&DatbaseServerId, "database", "", "Database Server Id")
	databaseCmd.AddCommand(listDatabaseDBCobraCmd)

	databaseCmd.AddCommand(listDatabaseCobraCmd)
	databaseCmd.AddCommand(databaseCreateWStageCobraCmd)
}

var databaseCmd = &cobra.Command{
	Use:               "database",
	Aliases:           []string{"prj"},
	Example:           "  acentera database list --sort creation\n  acentera database list --output csv\n  acentera database list --filter state=active,type=xxxxx\n  \n  acentera database select --filter name='My Database 1'",
	PersistentPreRun:  applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook, normalizeColumns),
	PersistentPostRun: applyHooks(verifyNewVersionHook, onVersionUpgrade, networkMonitorHook),
	Short:             "Database resources: sorting, filtering via tag/properties, output formatting, etc...",
}

var databaseCreateWStageCobraCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an database",
	Run:   dbCreateWebsiteStage,
}

func dbCreateWebsiteStage(*cobra.Command, []string) {
	//RepositoryOptions
	if DatbaseName == "" {
		logger.Error("Missing --name parameter")
		return
	}
	if DatbaesStage == "" {
		logger.Error("Missing --stage parameter")
		return
	}
	if DatabaseProject == "" {
		logger.Error("Missing --project parameter")
		return
	}
	if DatbaseServerId == "" {
		logger.Error("Missing --database parameter")
		return
	}

	if WSId == "" {
		logger.Error("Missing --website parameter")
		return
	}

	fmt.Println("Got WS:", WSId, " name:", DatbaseName, " Stage", DatbaesStage, "Project:", DatabaseProject, "db severid:", DatbaseServerId)
	databases, _ := cli.API().Databases().CreateDatabase(DatabaseProject, WSId, DatbaesStage, DatbaseName, DatbaseServerId)
	fmt.Println("create of db?", databases)
	return
}

var listDatabaseCobraCmd = &cobra.Command{
	Use:   "list-servers",
	Short: "List databases servers",
	Run:   databaseListResource,
}

func databaseListResource(*cobra.Command, []string) {
	d, _ := cli.API().Databases().GetDatabases(config.GetProjectId())

	columns := defaultDatabasesColumns
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	g := graph.NewGraph()
	for _, r := range d {
		// databaseInfo := r.(map[string]interface{})[id]
		// fmt.Println("1fz -- WWS INFO:", r)
		// res1 := outputs.New("database", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()

		res1 := outputs.New("database", r.Id).Object(reflect.ValueOf(r).Interface()).Prop("User", fmt.Sprintf("%v", "******")).Prop("Pass", fmt.Sprintf("%v", "******")).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()

		// res1 := outputs.New("database", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		// res1 := outputs.New("database", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Prop("ZState", fmt.Sprintf("%v", "OK")).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()

		// .Prop("Status", fmt.Sprintf("%v", r.Status))

		// fmt.Println("RES IS :")
		// fmt.Println(res1)
		// res1 := outputs.New("database", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}
	// fmt.Println("COLUMS ARE :")
	// fmt.Println(columns)
	// console.WithMaxWidth(console.GetTerminalWidth()),
	displayer, err := console.BuildOptions(
		console.WithRdfType("database"),
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
			console.WithRdfType("database"),
			console.WithColumns(listingColumnsFlag),
			console.WithFilters(listingFiltersFlag),
			console.WithTagFilters(listingTagFiltersFlag),
			console.WithTagKeyFilters(listingTagKeyFiltersFlag),
			console.WithTagValueFilters(listingTagValueFiltersFlag),

			console.WithFormat(listingFormat),
			console.WithidsOnly(listOnlyids),
			console.WithSortBy(sortBy...),

		).SetSource(databases).Build()
	*/
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))

}

var listDatabaseDBCobraCmd = &cobra.Command{
	Use:   "list",
	Short: "List databases",
	Run:   databaseListDBsResource,
}

func databaseListDBsResource(*cobra.Command, []string) {
	d, _ := cli.API().Databases().ListDatabases(DatbaseServerId)

	columns := defaultDatabasesColumns
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	g := graph.NewGraph()
	for _, r := range d {
		res1 := outputs.New("database", r.Id).Object(reflect.ValueOf(r).Interface()).Prop("Name", r.Name).Prop("User", fmt.Sprintf("%v", r.User)).Prop("Database", fmt.Sprintf("%v", r.Database)).Prop("Pass", fmt.Sprintf("%v", r.Pass)).Build()
		g.AddResource(
			res1,
		)
	}

	displayer, err := console.BuildOptions(
		console.WithRdfType("database"),
		console.WithColumns(columns),
		console.WithFormat(listingFormat),
		console.WithFilters(listingFiltersFlag),
		console.WithReverseSort(reverseFlag),
		console.WithSortBy(sortBy...),
		// console.WithNoHeaders(false),
	).SetSource(g).Build()

	exitOn(err)

	exitOn(displayer.Print(os.Stdout))

}

/*
func databaseSelectResource(*cobra.Command, []string) {
	databases, _ := cli.API().Databases().GetDatabases(config.GetProjectId())

	columns := defaultDatabasesColumns

	g := graph.NewGraph()
	for id, r := range databases {
		// databaseInfo := r.(map[string]interface{})[id]
		// fmt.Println("1 -- WEBS INFO:", r)
		// res1 := outputs.New("database", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		res1 := outputs.New("database", id).Object(reflect.ValueOf(r).Interface()).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}

	// Build Query
	buildQuery := console.BuildOptions(
		console.WithRdfType("database"),
		console.WithColumns(columns),
		console.WithFilters(listingFiltersFlag),
	)
	a, _ := buildQuery.BuildQuery()
	r, _ := g.Find(a)
	nbItems := len(r)

	// Reset
	config.Set("user.database.id", "")
	config.Set("user.database.name", "")
	if nbItems > 1 {
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitWithError(displayer.Print(os.Stdout), "Filters returned more than one record. Please refine query.")
	} else if nbItems <= 0 {
		exitWithError(nil, "Could not find any database that match filters. Please update your filter and try again.")
	} else if nbItems == 1 {
		// All Good
		config.Set("user.database.id", r[0].Id())
		config.Set("user.database.name", r[0].Name())
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitOn(displayer.Print(os.Stdout))
		logger.Infof("Successfully selected database") // , r[0].Name())
	} else {
		exitWithError(nil, "Invalid command ...")
	}
}
*/
