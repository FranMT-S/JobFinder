package scraper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/FranMT-S/JobFinder/models"
)

func CreateURLToScrapper(host models.HostScrapper, jobRequest models.JobRequest, page int) (string, error) {
	switch host {
	case models.RemoteOk:
		return createURLRemoteOk(jobRequest, page), nil
	case models.WorkRemotely:
		return createURLWorkRemotely(jobRequest), nil
	default:
		return "", fmt.Errorf("host not found")
	}
}

func createURLRemoteOk(jobRequest models.JobRequest, page int) string {
	offset := ((page) * 50) - 1

	tagsList := make([]string, len(jobRequest.Skills))
	for i, skill := range jobRequest.Skills {
		tagsList[i] = string(skill)
	}

	urlraw, err := url.Parse("https://remoteok.com/")
	if err != nil {
		return ""
	}
	q := urlraw.Query()

	category := models.CategoryMapRemoteOk[jobRequest.Category]
	if category != "" && category != "any" {
		tagsList = append(tagsList, category)
	}

	switch jobRequest.Level {
	case models.Junior:
		tagsList = append(tagsList, "junior")
	case models.Senior:
		tagsList = append(tagsList, "senior")
	}

	tags := strings.Join(tagsList, ",")

	q.Set("tags", tags)
	q.Set("order_by", "date")
	q.Set("action", "get_jobs")
	q.Set("offset", fmt.Sprintf("%d", offset))

	urlraw.RawQuery = q.Encode()

	return urlraw.String()
}

func createURLWorkRemotely(jobRequest models.JobRequest) string {
	categoryId := models.CategoryMapWeWorkRemotely[jobRequest.Category]
	categoryParam := ""

	if categoryId > 0 {
		categoryParam = fmt.Sprintf("&categories[]=%d", categoryId)
	}

	termsList := []string{}

	for _, skill := range jobRequest.Skills {
		termsList = append(termsList, string(skill))
	}

	switch jobRequest.Level {
	case models.Junior:
		termsList = append(termsList, "junior")
	case models.Senior:
		termsList = append(termsList, "senior")
	case models.Mid:
		termsList = append(termsList, "mid")
	case models.SemiSenior:
		termsList = append(termsList, "mid")
	}

	urlraw, err := url.Parse("https://weworkremotely.com/remote-jobs/search")
	if err != nil {
		return ""
	}
	q := urlraw.Query()
	for _, salary := range CreateSalaryRangeWeWorkremotely(jobRequest) {
		if salary != "" {
			q.Set("salary_range[]", salary)
		}
	}
	q.Set("term", strings.Join(termsList, " "))
	q.Set("categories_chosen", categoryParam)
	q.Set("countries_chosen", "")
	q.Set("skills_chosen", "")
	q.Set("commit", "Apply filters")

	urlraw.RawQuery = q.Encode()

	return urlraw.String()
}

// CreateSalaryRangeWeWorkremotely is a function that return
// the range salary used in weworkremotely in the api
func CreateSalaryRangeWeWorkremotely(jobRequest models.JobRequest) []string {
	minLevel, maxLevel := 0, 4
	minExpected, maxExpected := jobRequest.MinimumSalaryExpectation, jobRequest.MaximumSalaryExpectation

	// not specified
	if minExpected <= 0 && maxExpected <= 0 {
		return []string{}
	}

	// specified but not match return the min salary
	if minExpected < 10000 && maxExpected < 10000 {
		minLevel = 0
		maxLevel = 0
	}

	salaryRangeLevel := make([]string, 5)
	levels := []struct {
		min, max float64
	}{
		{10000, 25000},
		{25000, 49000},
		{49000, 75000},
		{75000, 100000},
	}

	// get the min and max range of salary
	for i, level := range levels {
		if minExpected >= level.min && minExpected <= level.max {
			minLevel = i
		}

		if maxExpected > level.min && maxExpected <= level.max {
			maxLevel = i
		}
	}

	// if the min salary is above 100000, return the max salary
	if minExpected >= 100000 {
		minLevel = 4
		maxLevel = 4
	}

	if maxExpected > 100000 {
		maxLevel = 4
	}

	for i := minLevel; i <= maxLevel; i++ {
		switch i {
		case 0:
			salaryRangeLevel[i] = WORK_REMOTELY_RANGE_SALARY_1
		case 1:
			salaryRangeLevel[i] = WORK_REMOTELY_RANGE_SALARY_2
		case 2:
			salaryRangeLevel[i] = WORK_REMOTELY_RANGE_SALARY_3
		case 3:
			salaryRangeLevel[i] = WORK_REMOTELY_RANGE_SALARY_4
		case 4:
			salaryRangeLevel[i] = WORK_REMOTELY_RANGE_SALARY_5
		}
	}

	// get only the range that not is empty string
	salaryRangeSlice := make([]string, 0)
	for _, salary := range salaryRangeLevel {
		if salary != "" {
			salaryRangeSlice = append(salaryRangeSlice, salary)
		}
	}

	return salaryRangeSlice
}
