package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FranMT-S/JobFinder/constants"
	"github.com/FranMT-S/JobFinder/models"
)

func CreateJobRequest(body []byte) (models.JobRequest, error) {
	var jobRequest models.JobRequest
	err := json.Unmarshal(body, &jobRequest)
	if err != nil {
		return models.JobRequest{}, fmt.Errorf("invalid body:%v", err)
	}

	if models.IsJobRequestNil(jobRequest) {
		return models.JobRequest{}, fmt.Errorf("invalid body")
	}

	return jobRequest, nil
}

// WriteJSONError is a function that writes a JSON error to the response
func WriteJSONError(w http.ResponseWriter, err models.ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)

	json.NewEncoder(w).Encode(err)
}

// GetPage return the page number from the request, if the page is not specified or invalidad, it returns 1
func GetPage(r *http.Request) int {
	page := r.URL.Query().Get(constants.PAGE)

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}

	if pageNumber < 1 {
		pageNumber = 1
	}

	return pageNumber
}

// GetMaxJobs return the max number of jobs to scrape from the request, if the max is not specified or invalidad, it returns 50
// if the max is greater than 100, it returns 100
func GetMaxJobs(r *http.Request) int {
	maxJobs := r.URL.Query().Get(constants.MAX)
	maxJobsNumber, err := strconv.Atoi(maxJobs)
	if err != nil {
		maxJobsNumber = constants.MIN_JOBS
	}

	if maxJobsNumber < 1 {
		maxJobsNumber = constants.MIN_JOBS
	}

	if maxJobsNumber > constants.MAX_JOBS {
		maxJobsNumber = constants.MAX_JOBS
	}

	return maxJobsNumber
}

// decodeJobRequest is a function that get a JobRequest from a request body
func decodeJobRequest(r *http.Request) (*models.JobRequest, int, error) {
	var jobRequest models.JobRequest

	err := json.NewDecoder(r.Body).Decode(&jobRequest)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("invalid body:%v", err)
	}

	if models.IsJobRequestNil(jobRequest) {
		return nil, http.StatusBadRequest, fmt.Errorf("the body is empty")
	}

	if jobRequest.MinimumSalaryExpectation > jobRequest.MaximumSalaryExpectation &&
		(jobRequest.MinimumSalaryExpectation > 0 && jobRequest.MaximumSalaryExpectation > 0) {
		return nil, http.StatusBadRequest, fmt.Errorf("minimum salary expectation must be less than maximum salary expectation")
	}

	if jobRequest.Host == nil {
		jobRequest.Host = []models.HostScrapper{models.RemoteOk, models.WorkRemotely}
	}

	return &jobRequest, http.StatusOK, nil
}
