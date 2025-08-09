package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/FranMT-S/JobFinder/constants"
	"github.com/FranMT-S/JobFinder/models"
	"github.com/FranMT-S/JobFinder/scraper"
	"github.com/go-chi/render"
)

type AnalizedJob struct {
	Job           models.Job           `json:"job"`
	MatchAnalizer models.MatchAnalizer `json:"matchAnalizer"`
}

func GetHost(w http.ResponseWriter, r *http.Request) {

	hosts := []struct {
		Id   models.HostScrapper `json:"id"`
		Name string              `json:"name"`
	}{
		{Id: models.RemoteOk, Name: "RemoteOk"},
		{Id: models.WorkRemotely, Name: "WorkRemotely"},
	}

	response := models.NewResponse("Hosts retrieved successfully", hosts)
	render.JSON(w, r, response)
}

func Scrap(w http.ResponseWriter, r *http.Request) {
	page := GetPage(r)
	maxJobs := GetMaxJobs(r)
	jobRequest, status, err := decodeJobRequest(r)

	if err != nil {
		WriteJSONError(w, *models.NewResponseError(
			status,
			err.Error(),
			err))
		return
	}

	chJob := make(chan models.Job)
	chWorkRemotely := make(chan models.Job)
	chError := make(chan error)

	analizedJobs := make([]AnalizedJob, 0, maxJobs)
	errorsTotal := make([]error, 0, len(jobRequest.Host))
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	processJobFunc := func(job models.Job) *models.Response {
		matchAnalizer := scraper.MatchAnalizer(job, *jobRequest)
		analizedJobs = append(analizedJobs, AnalizedJob{
			Job:           job,
			MatchAnalizer: matchAnalizer,
		})

		if len(analizedJobs) >= maxJobs {
			cancel()
			response := models.NewResponse(fmt.Sprintf("%d Jobs found", len(analizedJobs)), analizedJobs)
			return response
		}

		return nil
	}

	wg.Add(1)
	go initScraper(ctx, models.RemoteOk, *jobRequest, page, maxJobs, wg, chJob, chError, constants.MAX_JOBS_PARSE)

	wg.Add(1)
	go initScraper(ctx, models.WorkRemotely, *jobRequest, page, maxJobs, wg, chWorkRemotely, chError, constants.MAX_JOBS_PARSE)

	go func() {
		wg.Wait()
		close(chJob)
		close(chWorkRemotely)
		close(chError)
	}()

	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				WriteJSONError(w, *models.NewResponseError(
					http.StatusInternalServerError,
					"Request timeout",
					ctx.Err()))
				return
			}
			continue
		case job, ok := <-chJob:

			if !ok {
				chJob = nil
				response := models.NewResponse(fmt.Sprintf("%d Jobs found", len(analizedJobs)), analizedJobs)
				render.JSON(w, r, response)
				return
			}

			response := processJobFunc(job)
			if response != nil {
				render.JSON(w, r, response)
				return
			}

		case job, ok := <-chWorkRemotely:
			if !ok {
				chWorkRemotely = nil
				response := models.NewResponse(fmt.Sprintf("%d Jobs found", len(analizedJobs)), analizedJobs)
				render.JSON(w, r, response)
				return
			}

			response := processJobFunc(job)
			if response != nil {
				render.JSON(w, r, response)
				return
			}

		case err, ok := <-chError:
			if !ok {
				chError = nil
				continue
			}
			errorsTotal = append(errorsTotal, err)
			if len(errorsTotal) >= len(jobRequest.Host) {
				for _, err := range errorsTotal {
					log.Println("Error getting the jobs:", err)
				}

				WriteJSONError(w, *models.NewResponseError(
					http.StatusInternalServerError,
					"There was a error, try later",
					nil))

				return
			}
		}
	}
}

// ScrapRemoteOk is a function that scrapes the jobs from RemoteOk
// page is the page number to scrape
func ScrapRemoteOk(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}
	chJob := make(chan models.Job)
	chError := make(chan error)
	maxJobs := GetMaxJobs(r)
	jobs := make([]models.Job, 0, maxJobs)
	page := GetPage(r)

	processJobFunc := func(job models.Job) *models.Response {
		jobs = append(jobs, job)
		if len(jobs) >= maxJobs {
			cancel()
			response := models.NewResponse(fmt.Sprintf("%d Jobs found", len(jobs)), jobs)
			render.JSON(w, r, response)
			return response
		}

		return nil
	}

	jobRequest, status, err := decodeJobRequest(r)
	if err != nil {
		WriteJSONError(w, *models.NewResponseError(
			status,
			err.Error(),
			err))
		return
	}

	wg.Add(1)
	go initScraper(ctx, models.RemoteOk, *jobRequest, page, maxJobs, wg, chJob, chError, constants.MAX_JOBS_PARSE)

	go func() {
		wg.Wait()
		close(chJob)
		close(chError)
	}()

	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				WriteJSONError(w, *models.NewResponseError(
					http.StatusInternalServerError,
					"Request timeout",
					ctx.Err()))
				return
			}
			continue
		case job, ok := <-chJob:
			if !ok {
				jobResponse := models.NewJobResponse(jobs)
				response := models.NewResponse("Jobs found", jobResponse)

				render.JSON(w, r, response)
				return
			}
			response := processJobFunc(job)
			if response != nil {
				render.JSON(w, r, response)
				return
			}

		case err, ok := <-chError:
			if !ok {
				chError = nil
				continue
			}
			log.Println("Error getting the jobs:", err)
			WriteJSONError(w, *models.NewResponseError(
				http.StatusInternalServerError,
				err.Error(),
				err))
			return
		}
	}
}

// ScrapWorkRemotely is a function that scrapes the jobs from WorkRemotely
// page is the page number to scrape
func ScrapWorkRemotely(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}
	chJob := make(chan models.Job)
	chError := make(chan error)
	maxJobs := GetMaxJobs(r)
	jobs := make([]models.Job, 0, maxJobs)
	page := GetPage(r)

	jobRequest, status, err := decodeJobRequest(r)

	processJobFunc := func(job models.Job) *models.Response {
		jobs = append(jobs, job)
		if len(jobs) >= maxJobs {
			cancel()
			response := models.NewResponse(fmt.Sprintf("%d Jobs found", len(jobs)), jobs)
			render.JSON(w, r, response)
			return response
		}

		return nil
	}

	if err != nil {
		WriteJSONError(w, *models.NewResponseError(
			status,
			err.Error(),
			err))
		return
	}

	wg.Add(1)
	go initScraper(ctx, models.WorkRemotely, *jobRequest, page, maxJobs, wg, chJob, chError, constants.MAX_JOBS_PARSE)

	go func() {
		wg.Wait()
		close(chJob)
		close(chError)
	}()

	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				WriteJSONError(w, *models.NewResponseError(
					http.StatusInternalServerError,
					"Request timeout",
					ctx.Err()))
				return
			}
			continue
		case job, ok := <-chJob:
			if !ok {
				jobResponse := models.NewJobResponse(jobs)
				response := models.NewResponse("Jobs found", jobResponse)

				render.JSON(w, r, response)
				return
			}

			response := processJobFunc(job)
			if response != nil {
				render.JSON(w, r, response)
				return
			}
		case err, ok := <-chError:
			if !ok {
				chError = nil
				continue
			}
			log.Println("Error getting the jobs:", err)
			WriteJSONError(w, *models.NewResponseError(
				http.StatusInternalServerError,
				err.Error(),
				err))
			return
		}
	}
}

// initScraper is a function that initializes the scraper
// must be called in a go routine and add 1 to the wait group before calling
// when the function finish automatically calls wg.Done()
func initScraper(
	ctx context.Context,
	host models.HostScrapper,
	jobRequest models.JobRequest,
	page int,
	maxJobs int,
	wg *sync.WaitGroup,
	chJob chan<- models.Job,
	chError chan<- error,
	maxJobsParse int,
) {
	url, err := scraper.CreateURLToScrapper(host, jobRequest, page)
	if err != nil {
		chError <- err
		return
	}

	switch host {
	case models.RemoteOk:
		remoteOkScrapper := scraper.NewRemoteOkScraper()
		remoteOkScrapper.GetJobs(ctx, url, wg, chJob, chError, maxJobs, maxJobsParse)
	case models.WorkRemotely:
		workRemotelyScrapper := scraper.NewWorkRemotelyScraper()
		workRemotelyScrapper.GetJobs(ctx, url, wg, chJob, chError, maxJobs, maxJobsParse)
	default:
		wg.Done()
	}
}
