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
	defaultBlueprintsColumns = []string{}
	BlueprintId              = ""
)

func init() {
	defObject := cli.BlueprintApiObject()
	ptrObj := global.To_struct_ptr(defObject)
	defaultBlueprintsColumns = global.PrintFields(ptrObj)
	RootCmd.AddCommand(blueprintCmd)

	listBlueprintCobraCmd.PersistentFlags().StringSliceVar(&listingColumnsFlag, "columns", []string{}, "Select the properties to display in the columns. Ex: --columns id,name")
	listBlueprintCobraCmd.PersistentFlags().BoolVar(&noHeadersFlag, "no-headers", false, "Do not display headers")
	listBlueprintCobraCmd.PersistentFlags().BoolVar(&reverseFlag, "reverse", false, "Use in conjunction with --sort to reverse sort")
	listBlueprintCobraCmd.PersistentFlags().StringSliceVar(&sortBy, "sort", []string{"Id"}, "Sort tables by column(s) name(s)")

	cobra.EnableCommandSorting = false

	getBlueprintCobraCmd.PersistentFlags().StringVar(&BlueprintId, "blueprint", "", "Blueprint Id")

	// Global parametes
	blueprintCmd.PersistentFlags().StringSliceVar(&listingFiltersFlag, "filter", []string{}, "Filter resources given key/values fields (case insensitive). Ex: --filter name='Projct 1'")
	blueprintCmd.AddCommand(listBlueprintCobraCmd)

	blueprintCmd.AddCommand(getBlueprintCobraCmd)
	blueprintCmd.AddCommand(blueprintSelectCobraCmd)
}

var blueprintCmd = &cobra.Command{
	Use:               "blueprint",
	Aliases:           []string{"prj"},
	Example:           "  acentera blueprint list --sort creation\n  acentera blueprint list --output csv\n  acentera blueprint list --filter state=active,type=xxxxx\n  \n  acentera blueprint select --filter name='My Blueprint 1'",
	PersistentPreRun:  applyHooks(initAwlessEnvHook, initLoggerHook, initCloudServicesHook, firstInstallDoneHook, initCliEnvHook, normalizeColumns),
	PersistentPostRun: applyHooks(verifyNewVersionHook, onVersionUpgrade, networkMonitorHook),
	Short:             "Blueprint resources: sorting, filtering via tag/properties, output formatting, etc...",
}
var blueprintSelectCobraCmd = &cobra.Command{
	Use:   "select",
	Short: "Select an blueprint",
	Run:   blueprintSelectResource,
}

var listBlueprintCobraCmd = &cobra.Command{
	Use:   "list",
	Short: "List blueprints",
	Run:   blueprintListResource,
}

var getBlueprintCobraCmd = &cobra.Command{
	Use:   "get",
	Short: "Get blueprint info",
	Run:   blueprintGetbyIdResource,
}

func blueprintGetbyIdResource(*cobra.Command, []string) {
	if BlueprintId == "" {
		logger.Error("Missing --blueprint parameter")
		return
	}
	blueprints, _ := cli.API().Blueprints().GetBlueprintsById(config.GetProjectId(), BlueprintId)
	// fmt.Println("Test", blueprints)

	columns := defaultBlueprintsColumns
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	g := graph.NewGraph()
	res1Tmp := outputs.New("blueprint", blueprints.Id).Object1(reflect.ValueOf(*blueprints).Interface()) // .Build()

	/*
		jsonData, err := json.Marshal(blueprints.Stages)
		if err != nil {
		}
	*/

	theStage := ""
	firstStage := ""
	for r, s := range blueprints.Releases {
		if s.Enabled != "" {
			if firstStage == "" {
				firstStage = r
			}
			if s.Release == "master" {
				theStage = r
			}
		}
	}

	res1 := res1Tmp.Prop("Name", fmt.Sprintf("%v", blueprints.Name)).Prop("Project", blueprints.Project).Prop("Status", blueprints.Status).Prop("ACenterAType", blueprints.ACenterAType).Prop("Release", theStage).Build()
	// .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
	g.AddResource(
		res1,
	)
	displayer, err := console.BuildOptions(
		console.WithRdfType("blueprint"),
		console.WithColumns(columns),
		console.WithFormat(listingFormat),
		console.WithReverseSort(reverseFlag),
		console.WithSortBy(sortBy...),
	).SetSource(g).Build()
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))
	return
}

func blueprintListResource(*cobra.Command, []string) {
	blueprints, _ := cli.API().Blueprints().GetBlueprints(config.GetProjectId())

	// fmt.Println("Received of ", blueprints)
	columns := defaultBlueprintsColumns
	if len(listingColumnsFlag) >= 1 {
		columns = listingColumnsFlag
	}
	g := graph.NewGraph()
	for id, r := range blueprints {
		// blueprintInfo := r.(map[string]interface{})[id]
		// fmt.Println("1fz -- WWS INFO:", r)
		// res1 := outputs.New("blueprint", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		res1Tmp := outputs.New("blueprint", id).Object(reflect.ValueOf(r).Interface()) // .Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		// res1 := res1Tmp.Prop("Name", fmt.Sprintf("%v", r.Name)).Prop("Status", fmt.Sprintf("%v", r.Status)).Build()
		res1 := res1Tmp.Prop("Name", fmt.Sprintf("%v", r.Name)).Prop("Project", fmt.Sprintf("%v", r.Project)).Build()
		// res1 := outputs.New("blueprint", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		// res1 := outputs.New("blueprint", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Prop("ZState", fmt.Sprintf("%v", "OK")).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()

		// .Prop("Status", fmt.Sprintf("%v", r.Status))

		// fmt.Println("RES IS :")
		// fmt.Println(res1)
		// res1 := outputs.New("blueprint", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}
	// fmt.Println("COLUMS ARE :")
	// fmt.Println(columns)
	// console.WithMaxWidth(console.GetTerminalWidth()),
	displayer, err := console.BuildOptions(
		console.WithRdfType("blueprint"),
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
			console.WithRdfType("blueprint"),
			console.WithColumns(listingColumnsFlag),
			console.WithFilters(listingFiltersFlag),
			console.WithTagFilters(listingTagFiltersFlag),
			console.WithTagKeyFilters(listingTagKeyFiltersFlag),
			console.WithTagValueFilters(listingTagValueFiltersFlag),

			console.WithFormat(listingFormat),
			console.WithidsOnly(listOnlyids),
			console.WithSortBy(sortBy...),

		).SetSource(blueprints).Build()
	*/
	exitOn(err)

	exitOn(displayer.Print(os.Stdout))

}

func blueprintSelectResource(*cobra.Command, []string) {
	blueprints, _ := cli.API().Blueprints().GetBlueprints(config.GetProjectId())

	columns := defaultBlueprintsColumns

	g := graph.NewGraph()
	for id, r := range blueprints {
		// blueprintInfo := r.(map[string]interface{})[id]
		// fmt.Println("1 -- WEBS INFO:", r)
		// res1 := outputs.New("blueprint", id).Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		res1 := outputs.New("blueprint", id).Object(reflect.ValueOf(r).Interface()).Build() // .Build() // .Prop("Id", id).Prop("Name", fmt.Sprintf("%v", r.Name)).Build()
		g.AddResource(
			res1,
		)
	}

	// Build Query
	buildQuery := console.BuildOptions(
		console.WithRdfType("blueprint"),
		console.WithColumns(columns),
		console.WithFilters(listingFiltersFlag),
	)
	a, _ := buildQuery.BuildQuery()
	r, _ := g.Find(a)
	nbItems := len(r)

	// Reset
	config.Set("user.blueprint.id", "")
	config.Set("user.blueprint.name", "")
	if nbItems > 1 {
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitWithError(displayer.Print(os.Stdout), "Filters returned more than one record. Please refine query.")
	} else if nbItems <= 0 {
		exitWithError(nil, "Could not find any blueprint that match filters. Please update your filter and try again.")
	} else if nbItems == 1 {
		// All Good
		config.Set("user.blueprint.id", r[0].Id())
		config.Set("user.blueprint.name", r[0].Name())
		displayer, err := buildQuery.SetSource(g).Build()
		exitOn(err)
		exitOn(displayer.Print(os.Stdout))
		logger.Infof("Successfully selected blueprint") // , r[0].Name())
	} else {
		exitWithError(nil, "Invalid command ...")
	}
}
