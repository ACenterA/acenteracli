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
	"github.com/wallix/awless/console"
	"github.com/wallix/awless/graph"
	"github.com/wallix/awless/graph/outputs"
	"github.com/wallix/awless/logger"
)

var ()

func init() {
	RootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(listProjectCobraCmd)

	projectCmd.PersistentFlags().StringSliceVar(&listingColumnsFlag, "columns", []string{}, "Select the properties to display in the columns. Ex: --columns id,name")

	cobra.EnableCommandSorting = false
}

var projectCmd = &cobra.Command{
	Use:               "project",
	Aliases:           []string{"prj"},
	Example:           "  acentera project list --sort creation\n  acentera project list --format csv\n  acentera list project --filter state=active,type=xxxxx\n",
	PersistentPreRun:  applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook, normalizeColumns),
	PersistentPostRun: applyHooks(verifyNewVersionHook, onVersionUpgrade, networkMonitorHook),
	Short:             "Project resources: sorting, filtering via tag/properties, output formatting, etc...",
}

var listProjectCobraCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run:   projctListResource,
}

func projctListResource(*cobra.Command, []string) {
	projects, _ := cli.API().Projects().GetProjects()

	logger.Infof("%s", projects)

	columns := []string{"Name", "id"}
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	fmt.Println(columns)
	g := graph.NewGraph()
	for id, r := range projects {
		// projectInfo := r.(map[string]interface{})[id]
		// fmt.Println("1 -- PRJ INFO:", r)
		res1 := outputs.New("project", id).Prop("id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}
	// r, err := g.GetResource("sub_1", fmt.Sprintf("%v", z))
	// fmt.Println(r)
	// exitOn(err)
	displayer, err := console.BuildOptions(
		console.WithMaxWidth(console.GetTerminalWidth()),
		console.WithRdfType("project"),
		console.WithColumns(columns),
		console.WithColumns(listingColumnsFlag),
		console.WithFormat(listingFormat),
		console.WithReverseSort(reverseFlag),
		console.WithMaxWidth(console.GetTerminalWidth()),
		console.WithNoHeaders(noHeadersFlag),
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

/*
func printResources(g cloud.GraphAPI, resType string) {
	displayer, err := console.BuildOptions(
		console.WithRdfType(resType),
		console.WithColumns(listingColumnsFlag),
		console.WithFilters(listingFiltersFlag),
		console.WithTagFilters(listingTagFiltersFlag),
		console.WithTagKeyFilters(listingTagKeyFiltersFlag),
		console.WithTagValueFilters(listingTagValueFiltersFlag),
		console.WithMaxWidth(console.GetTerminalWidth()),
		console.WithFormat(listingFormat),
		console.WithidsOnly(listOnlyids),
		console.WithSortBy(sortBy...),
		console.WithReverseSort(reverseFlag),
		console.WithNoHeaders(noHeadersFlag),
	).SetSource(g).Build()
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))
}
*/
