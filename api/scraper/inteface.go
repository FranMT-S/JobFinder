package scraper

import (
	"context"
	"sync"

	"github.com/FranMT-S/JobFinder/models"
)

// ScrapperJob is an interface that represents a job scrapper
// urlServer is the url of the server that will be scraped
// wg is a WaitGroup that will be used to wait for the scraper to finish
// chJob is a channel that will be used to send the jobs to the main function
// chError is a channel that will be used to send the errors to the main function
// maxJobs is the maximum number of jobs that will be scraped
// maxJobsParse is the maximum number of tags of job that must be try to parsed
type ScrapperJob interface {
	GetJobs(ctx context.Context,
		url string,
		wg *sync.WaitGroup,
		chJob chan<- models.Job,
		chError chan<- error,
		maxJobs int,
		maxJobsParse int,
	)
}
