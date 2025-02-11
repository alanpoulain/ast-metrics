package Analyzer

import (
	pb "github.com/halleck45/ast-metrics/src/NodeType"
)

type Comparator struct {
	ComparedBranch string
}

const (
	ADDED    = "added"
	DELETED  = "deleted"
	MODIFIED = "modified"
	UNCHANGED = "unchanged"
)

type ChangedFile struct {
	Path            string
	Comparaison     Comparaison
	Status          string
	LatestVersion   *pb.File
	PreviousVersion *pb.File
	IsNegligible    bool
}
type Comparaison struct {
	ComparedBranch                       string
	NbFiles                              int
	NbFunctions                          int
	NbClasses                            int
	NbClassesWithCode                    int
	NbMethods                            int
	Loc                                  int
	Cloc                                 int
	Lloc                                 int
	AverageMethodsPerClass               float64
	AverageLocPerMethod                  float64
	AverageLlocPerMethod                 float64
	AverageClocPerMethod                 float64
	AverageCyclomaticComplexityPerMethod float64
	AverageCyclomaticComplexityPerClass  float64
	MinCyclomaticComplexity              int
	MaxCyclomaticComplexity              int
	AverageHalsteadDifficulty            float64
	AverageHalsteadEffort                float64
	AverageHalsteadVolume                float64
	AverageHalsteadTime                  float64
	AverageHalsteadBugs                  float64
	SumHalsteadDifficulty                float64
	SumHalsteadEffort                    float64
	SumHalsteadVolume                    float64
	SumHalsteadTime                      float64
	SumHalsteadBugs                      float64
	AverageMI                            float64
	AverageMIwoc                         float64
	AverageMIcw                          float64
	AverageMIPerMethod                   float64
	AverageMIwocPerMethod                float64
	AverageMIcwPerMethod                 float64
	AverageAfferentCoupling              float64
	AverageEfferentCoupling              float64
	AverageInstability                   float64
	CommitCountForPeriod                 int
	CommittedFilesCountForPeriod         int // for example if one commit concerns 10 files, it will be 10
	BusFactor                            int
	Risk                                 float64
	ChangedFiles                         []ChangedFile
	NbNewFiles                           int
	NbDeletedFiles                       int
	NbModifiedFiles                      int
	NbModifiedFilesIncludingNegligible   int
}

func NewComparator(comparedBranch string) *Comparator {
	return &Comparator{
		ComparedBranch: comparedBranch,
	}
}

func (c *Comparator) Compare(first Aggregated, second Aggregated) Comparaison {
	comparaison := Comparaison{
		ComparedBranch: c.ComparedBranch,
	}

	// Compare the two projects
	comparaison.NbFiles = first.NbFiles - second.NbFiles
	comparaison.NbFunctions = first.NbFunctions - second.NbFunctions
	comparaison.NbClasses = first.NbClasses - second.NbClasses
	comparaison.NbClassesWithCode = first.NbClassesWithCode - second.NbClassesWithCode
	comparaison.NbMethods = first.NbMethods - second.NbMethods
	comparaison.Loc = first.Loc - second.Loc
	comparaison.Cloc = first.Cloc - second.Cloc
	comparaison.Lloc = first.Lloc - second.Lloc
	comparaison.AverageMethodsPerClass = first.AverageMethodsPerClass - second.AverageMethodsPerClass
	comparaison.AverageLocPerMethod = first.AverageLocPerMethod - second.AverageLocPerMethod
	comparaison.AverageLlocPerMethod = first.AverageLlocPerMethod - second.AverageLlocPerMethod
	comparaison.AverageClocPerMethod = first.AverageClocPerMethod - second.AverageClocPerMethod
	comparaison.AverageCyclomaticComplexityPerMethod = first.AverageCyclomaticComplexityPerMethod - second.AverageCyclomaticComplexityPerMethod
	comparaison.AverageCyclomaticComplexityPerClass = first.AverageCyclomaticComplexityPerClass - second.AverageCyclomaticComplexityPerClass
	comparaison.MinCyclomaticComplexity = first.MinCyclomaticComplexity - second.MinCyclomaticComplexity
	comparaison.MaxCyclomaticComplexity = first.MaxCyclomaticComplexity - second.MaxCyclomaticComplexity
	comparaison.AverageHalsteadDifficulty = first.AverageHalsteadDifficulty - second.AverageHalsteadDifficulty
	comparaison.AverageHalsteadEffort = first.AverageHalsteadEffort - second.AverageHalsteadEffort
	comparaison.AverageHalsteadVolume = first.AverageHalsteadVolume - second.AverageHalsteadVolume
	comparaison.AverageHalsteadTime = first.AverageHalsteadTime - second.AverageHalsteadTime
	comparaison.AverageHalsteadBugs = first.AverageHalsteadBugs - second.AverageHalsteadBugs
	comparaison.SumHalsteadDifficulty = first.SumHalsteadDifficulty - second.SumHalsteadDifficulty
	comparaison.SumHalsteadEffort = first.SumHalsteadEffort - second.SumHalsteadEffort
	comparaison.SumHalsteadVolume = first.SumHalsteadVolume - second.SumHalsteadVolume
	comparaison.SumHalsteadTime = first.SumHalsteadTime - second.SumHalsteadTime
	comparaison.SumHalsteadBugs = first.SumHalsteadBugs - second.SumHalsteadBugs
	comparaison.AverageMI = first.AverageMI - second.AverageMI
	comparaison.AverageMIwoc = first.AverageMIwoc - second.AverageMIwoc
	comparaison.AverageMIcw = first.AverageMIcw - second.AverageMIcw
	comparaison.AverageMIPerMethod = first.AverageMIPerMethod - second.AverageMIPerMethod
	comparaison.AverageMIwocPerMethod = first.AverageMIwocPerMethod - second.AverageMIwocPerMethod
	comparaison.AverageMIcwPerMethod = first.AverageMIcwPerMethod - second.AverageMIcwPerMethod
	comparaison.AverageAfferentCoupling = first.AverageAfferentCoupling - second.AverageAfferentCoupling
	comparaison.AverageEfferentCoupling = first.AverageEfferentCoupling - second.AverageEfferentCoupling
	comparaison.AverageInstability = first.AverageInstability - second.AverageInstability
	comparaison.CommitCountForPeriod = first.CommitCountForPeriod - second.CommitCountForPeriod
	comparaison.CommittedFilesCountForPeriod = first.CommittedFilesCountForPeriod - second.CommittedFilesCountForPeriod
	comparaison.BusFactor = first.BusFactor - second.BusFactor

	for _, file := range first.ConcernedFiles {

		change := ChangedFile{
			Path: file.Path,
			Comparaison: Comparaison{
				NbFunctions:                          0,
				NbClasses:                            0,
				NbClassesWithCode:                    0,
				NbMethods:                            0,
				Loc:                                  0,
				Cloc:                                 0,
				Lloc:                                 0,
				AverageMethodsPerClass:               0,
				AverageLocPerMethod:                  0,
				AverageLlocPerMethod:                 0,
				AverageClocPerMethod:                 0,
				AverageCyclomaticComplexityPerMethod: 0,
				AverageCyclomaticComplexityPerClass:  0,
				MinCyclomaticComplexity:              0,
				MaxCyclomaticComplexity:              0,
				AverageHalsteadDifficulty:            0,
				AverageHalsteadEffort:                0,
				AverageHalsteadVolume:                0,
				AverageHalsteadTime:                  0,
				AverageHalsteadBugs:                  0,
				SumHalsteadDifficulty:                0,
				SumHalsteadEffort:                    0,
				SumHalsteadVolume:                    0,
				SumHalsteadTime:                      0,
				SumHalsteadBugs:                      0,
				AverageMI:                            0,
				AverageMIwoc:                         0,
				AverageMIcw:                          0,
				AverageMIPerMethod:                   0,
				AverageMIwocPerMethod:                0,
				AverageMIcwPerMethod:                 0,
				AverageAfferentCoupling:              0,
				AverageEfferentCoupling:              0,
				AverageInstability:                   0,
				CommitCountForPeriod:                 0,
				CommittedFilesCountForPeriod:         0,
				BusFactor:                            0,
				Risk:                                 0,
			},
			Status:          ADDED,
			LatestVersion:   file,
			PreviousVersion: nil,
			IsNegligible:    false,
		}

		for _, file2 := range second.ConcernedFiles {

			if file.Path != file2.Path {
				continue
			}

			if file.Checksum == file2.Checksum {
				// already present, no change
				change.Status = UNCHANGED
			}

			// nb functions
			change.Comparaison.NbFunctions = 0
			before := 0
			after := 0
			if file.Stmts != nil && file.Stmts.StmtFunction != nil {
				before = len(file.Stmts.StmtFunction)
			}
			if file2.Stmts != nil && file2.Stmts.StmtFunction != nil {
				after = len(file2.Stmts.StmtFunction)
			}
			change.Comparaison.NbFunctions = before - after

			// nb classes
			change.Comparaison.NbClasses = 0
			before = 0
			after = 0
			if file.Stmts != nil && file.Stmts.StmtClass != nil {
				before = len(file.Stmts.StmtClass)
			}
			if file2.Stmts != nil && file2.Stmts.StmtClass != nil {
				after = len(file2.Stmts.StmtClass)
			}
			change.Comparaison.NbClasses = before - after

			// Loc, cloc...
			if file.LinesOfCode != nil || file2.LinesOfCode != nil {
				change.Comparaison.Loc = int(file.LinesOfCode.LinesOfCode) - int(file2.LinesOfCode.LinesOfCode)
				change.Comparaison.Cloc = int(file.LinesOfCode.CommentLinesOfCode) - int(file2.LinesOfCode.CommentLinesOfCode)
				change.Comparaison.Lloc = int(file.LinesOfCode.LogicalLinesOfCode) - int(file2.LinesOfCode.LogicalLinesOfCode)
			}

			if file.Stmts.Analyze != nil && file2.Stmts.Analyze != nil {

				// Cyclomatic complexity
				if file.Stmts.Analyze.Complexity != nil && file2.Stmts.Analyze.Complexity != nil {
					change.Comparaison.AverageCyclomaticComplexityPerMethod = float64(*file.Stmts.Analyze.Complexity.Cyclomatic) - float64(*file2.Stmts.Analyze.Complexity.Cyclomatic)
				}

				// Halstead
				if file.Stmts.Analyze.Volume != nil && file.Stmts.Analyze.Volume.HalsteadDifficulty != nil &&
					file2.Stmts.Analyze.Volume != nil && file2.Stmts.Analyze.Volume.HalsteadDifficulty != nil {
					change.Comparaison.AverageHalsteadDifficulty = float64(*file.Stmts.Analyze.Volume.HalsteadDifficulty) - float64(*file2.Stmts.Analyze.Volume.HalsteadDifficulty)
					change.Comparaison.AverageHalsteadEffort = float64(*file.Stmts.Analyze.Volume.HalsteadEffort) - float64(*file2.Stmts.Analyze.Volume.HalsteadEffort)
					change.Comparaison.AverageHalsteadVolume = float64(*file.Stmts.Analyze.Volume.HalsteadVolume) - float64(*file2.Stmts.Analyze.Volume.HalsteadVolume)
					change.Comparaison.AverageHalsteadTime = float64(*file.Stmts.Analyze.Volume.HalsteadTime) - float64(*file2.Stmts.Analyze.Volume.HalsteadTime)
				}

				// Maintainability index
				if file.Stmts.Analyze.Maintainability != nil && file2.Stmts.Analyze.Maintainability != nil && file.Stmts.Analyze.Maintainability.MaintainabilityIndex != nil && file2.Stmts.Analyze.Maintainability.MaintainabilityIndex != nil {
					change.Comparaison.AverageMI = float64(*file.Stmts.Analyze.Maintainability.MaintainabilityIndex) - float64(*file2.Stmts.Analyze.Maintainability.MaintainabilityIndex)
				}
				if file.Stmts.Analyze.Maintainability != nil && file2.Stmts.Analyze.Maintainability != nil && file.Stmts.Analyze.Maintainability.MaintainabilityIndexWithoutComments != nil && file2.Stmts.Analyze.Maintainability.MaintainabilityIndexWithoutComments != nil {
					change.Comparaison.AverageMIwoc = float64(*file.Stmts.Analyze.Maintainability.MaintainabilityIndexWithoutComments) - float64(*file2.Stmts.Analyze.Maintainability.MaintainabilityIndexWithoutComments)
				}

				// Coupling
				if file.Stmts.Analyze.Coupling != nil && file2.Stmts.Analyze.Coupling != nil {
					change.Comparaison.AverageAfferentCoupling = float64(file.Stmts.Analyze.Coupling.Afferent) - float64(file2.Stmts.Analyze.Coupling.Afferent)
					change.Comparaison.AverageEfferentCoupling = float64(file.Stmts.Analyze.Coupling.Efferent) - float64(file2.Stmts.Analyze.Coupling.Efferent)
					change.Comparaison.AverageInstability = float64(file.Stmts.Analyze.Coupling.Instability) - float64(file2.Stmts.Analyze.Coupling.Instability)
				}

				// Risk
				if file.Stmts.Analyze.Risk != nil && file2.Stmts.Analyze.Risk != nil {
					change.Comparaison.Risk = float64(file.Stmts.Analyze.Risk.Score) - float64(file2.Stmts.Analyze.Risk.Score)
					// check if not NaN
					if change.Comparaison.Risk != change.Comparaison.Risk {
						change.Comparaison.Risk = 0
					}
				}
			}

			// if changes concerns only white spaces, etc. we don't want to include it
			change.IsNegligible = change.Comparaison.NbFunctions == 0 &&
				change.Comparaison.NbClasses == 0 &&
				change.Comparaison.Loc == 0 &&
				change.Comparaison.Cloc == 0 &&
				change.Comparaison.Lloc == 0 &&
				change.Comparaison.AverageCyclomaticComplexityPerMethod == 0 &&
				change.Comparaison.AverageHalsteadDifficulty == 0 &&
				change.Comparaison.AverageHalsteadEffort == 0 &&
				change.Comparaison.AverageHalsteadVolume == 0 &&
				change.Comparaison.AverageHalsteadTime == 0 &&
				change.Comparaison.AverageMI == 0 &&
				change.Comparaison.AverageMIwoc == 0 &&
				change.Comparaison.AverageAfferentCoupling == 0 &&
				change.Comparaison.AverageEfferentCoupling == 0 &&
				change.Comparaison.AverageInstability == 0 &&
				change.Comparaison.Risk == 0

			change.Status = MODIFIED
			change.PreviousVersion = file2
			change.LatestVersion = file
			break
		}

		if change.PreviousVersion != nil {
			if change.PreviousVersion.Checksum == change.LatestVersion.Checksum {
				// include only changed files
				continue
			}
		}

		switch change.Status {
		case ADDED:
			comparaison.NbNewFiles++
		case DELETED:
			comparaison.NbDeletedFiles++
		case MODIFIED:
			if change.IsNegligible {
				comparaison.NbModifiedFilesIncludingNegligible++
			} else {
				comparaison.NbModifiedFiles++
			}
		}

		if change.IsNegligible {
			continue
		}

		if change.Status == UNCHANGED {
			continue
		}

		comparaison.ChangedFiles = append(comparaison.ChangedFiles, change)
	}

	return comparaison
}
