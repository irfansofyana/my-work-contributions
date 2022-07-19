package report

import (
	"context"
	"fmt"

	"github.com/irfansofyana/workcontributions/internal/domain/repository"
	"google.golang.org/appengine/log"
)

// Print is function to print the report into stdout
func Print(repo repository.WorkRepository, param *repository.WorkRepositoryParam) {
	listOfWork, err := repo.GetListOfWork(param)
	if err != nil {
		log.Errorf(context.Background(), "Print: error print report to stdout", err)
		return
	}

	for i, work := range listOfWork {
		fmt.Printf("%d. What: %s\n", i+1, work.What)
		fmt.Printf("\t Type: %s\n", work.Type.String())
		fmt.Printf("\t Related Link: %s\n", work.RelatedLink)
	}
}
