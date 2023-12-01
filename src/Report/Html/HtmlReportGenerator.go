package Report

import (
	"fmt"
	"os"
	"sort"

	log "github.com/sirupsen/logrus"

	"github.com/flosch/pongo2/v5"
	"github.com/halleck45/ast-metrics/src/Analyzer"
	"github.com/halleck45/ast-metrics/src/Cli"
	"github.com/halleck45/ast-metrics/src/Engine"
	pb "github.com/halleck45/ast-metrics/src/NodeType"
)

type ReportGenerator struct {
	// The path where the report will be generated
	ReportPath string
}

func NewReportGenerator(reportPath string) *ReportGenerator {
	return &ReportGenerator{
		ReportPath: reportPath,
	}
}

func (v *ReportGenerator) Generate(files []*pb.File, projectAggregated Analyzer.ProjectAggregated) error {

	// Ensure destination folder exists
	err := v.EnsureFolder(v.ReportPath)
	if err != nil {
		return err
	}

	// Define loader in order to retrieve templates in the Report/Html/templates folder
	loader := pongo2.MustNewLocalFileSystemLoader("src/Report/Html/templates")
	pongo2.DefaultSet = pongo2.NewSet("src/Report/Html/templates", loader)

	// Custom filters
	v.RegisterFilters()

	// Overview
	v.GenerateLanguagePage("index.html", "All", projectAggregated.Combined, files, projectAggregated)
	// by language overview
	for language, currentView := range projectAggregated.ByProgrammingLanguage {
		v.GenerateLanguagePage("index.html", language, currentView, files, projectAggregated)
	}

	return nil
}

func (v *ReportGenerator) GenerateLanguagePage(template string, language string, currentView Analyzer.Aggregated, files []*pb.File, projectAggregated Analyzer.ProjectAggregated) error {

	// Compile the index.html template
	tpl, err := pongo2.DefaultSet.FromFile("index.html")
	if err != nil {
		log.Error(err)
	}
	// Render it, passing projectAggregated and files as context
	out, err := tpl.Execute(pongo2.Context{"currentLanguage": language, "currentView": currentView, "projectAggregated": projectAggregated, "files": files})
	if err != nil {
		log.Error(err)
	}

	// Write the result to the file
	pageSuffix := ""
	if language != "All" {
		pageSuffix = fmt.Sprintf("_%s", language)
	}
	file, err := os.Create(fmt.Sprintf("%s/index%s.html", v.ReportPath, pageSuffix))
	if err != nil {
		log.Error(err)
	}
	defer file.Close()
	file.WriteString(out)

	return nil
}

func (v *ReportGenerator) EnsureFolder(path string) error {
	// check if the folder exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create it
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *ReportGenerator) RegisterFilters() {

	pongo2.RegisterFilter("sortMaintainabilityIndex", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the list to sort
		// create new empty list
		list := make([]*pb.StmtClass, 0)

		// append to the list when file contians at lease one class
		for _, file := range in.Interface().([]*pb.File) {
			if len(file.Stmts.StmtClass) == 0 {
				continue
			}

			classes := Engine.GetClassesInFile(file)

			for _, class := range classes {
				if class.Stmts.Analyze.Maintainability == nil {
					continue
				}

				if *class.Stmts.Analyze.Maintainability.MaintainabilityIndex < 1 {
					continue
				}

				if *class.Stmts.Analyze.Maintainability.MaintainabilityIndex > 65 {
					continue
				}

				list = append(list, class)
			}
		}

		// sort the list, manually
		sort.Slice(list, func(i, j int) bool {
			if list[i].Stmts.Analyze.Maintainability == nil {
				return false
			}
			if list[j].Stmts.Analyze.Maintainability == nil {
				return true
			}

			// get first class in file
			class1 := list[i]
			class2 := list[j]

			return *class1.Stmts.Analyze.Maintainability.MaintainabilityIndex < *class2.Stmts.Analyze.Maintainability.MaintainabilityIndex
		})

		// keep only the first 10
		if len(list) > 10 {
			list = list[:10]
		}

		return pongo2.AsValue(list), nil
	})

	pongo2.RegisterFilter("sortRisk", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the list to sort
		// create new empty list
		list := make([]*pb.File, 0)

		// append to the list when file contains at lease one class
		for _, file := range in.Interface().([]*pb.File) {
			if file.Stmts.StmtClass == nil {
				continue
			}

			list = append(list, file)
		}

		// sort the list
		sort.Slice(list, func(i, j int) bool {

			if list[i].Stmts.Analyze.Risk == nil {
				return false
			}

			if list[i].Stmts.StmtClass == nil {
				return true
			}

			if list[j].Stmts.StmtClass == nil {
				return true
			}

			class1 := list[i].Stmts.StmtClass[0]
			class2 := list[j].Stmts.StmtClass[0]

			if class1.Stmts.Analyze.Risk == nil {
				return false
			}
			if class2.Stmts.Analyze.Risk == nil {
				return true
			}

			return class1.Stmts.Analyze.Risk.Score > class2.Stmts.Analyze.Risk.Score
		})

		// keep only the first 10
		if len(list) > 10 {
			list = list[:10]
		}

		return pongo2.AsValue(list), nil
	})

	// filter to format number. Ex: 1234 -> 1 K
	pongo2.RegisterFilter("stringifyNumber", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the number to format
		number := in.Integer()

		// format it
		if number > 1000000 {
			return pongo2.AsValue(fmt.Sprintf("%.1f M", float64(number)/1000000)), nil
		} else if number > 1000 {
			return pongo2.AsValue(fmt.Sprintf("%.1f K", float64(number)/1000)), nil
		}

		return pongo2.AsValue(number), nil
	})

	// filter that Return new Cli.NewComponentBarchartCyclomaticByMethodRepartition(aggregated, files)
	pongo2.RegisterFilter("barchartCyclomaticByMethodRepartition", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the aggregated and files
		aggregated := in.Interface().(Analyzer.Aggregated)
		files := aggregated.ConcernedFiles

		// create the component
		comp := Cli.NewComponentBarchartCyclomaticByMethodRepartition(aggregated, files)
		return pongo2.AsSafeValue(comp.RenderHTML()), nil
	})

	// filter barchartCyclomaticByMethodRepartition
	pongo2.RegisterFilter("barchartCyclomaticByMethodRepartition", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the aggregated and files
		aggregated := in.Interface().(Analyzer.Aggregated)
		files := aggregated.ConcernedFiles

		// create the component
		comp := Cli.NewComponentBarchartCyclomaticByMethodRepartition(aggregated, files)
		return pongo2.AsSafeValue(comp.RenderHTML()), nil
	})

	// filter barchartMaintainabilityIndexRepartition
	pongo2.RegisterFilter("barchartMaintainabilityIndexRepartition", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the aggregated and files
		aggregated := in.Interface().(Analyzer.Aggregated)
		files := aggregated.ConcernedFiles

		// create the component
		comp := Cli.NewComponentBarchartMaintainabilityIndexRepartition(aggregated, files)
		return pongo2.AsSafeValue(comp.RenderHTML()), nil
	})

	// filter barchartLocPerMethodRepartition
	pongo2.RegisterFilter("barchartLocPerMethodRepartition", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the aggregated and files
		aggregated := in.Interface().(Analyzer.Aggregated)
		files := aggregated.ConcernedFiles

		// create the component
		comp := Cli.NewComponentBarchartLocByMethodRepartition(aggregated, files)
		return pongo2.AsSafeValue(comp.RenderHTML()), nil
	})

	// filter lineChartGitActivity
	pongo2.RegisterFilter("lineChartGitActivity", func(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
		// get the aggregated and files
		aggregated := in.Interface().(Analyzer.Aggregated)
		files := aggregated.ConcernedFiles

		// create the component
		comp := Cli.NewComponentLineChartGitActivity(aggregated, files)
		return pongo2.AsSafeValue(comp.RenderHTML()), nil
	})
}
