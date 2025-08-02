package helpers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/FranMT-S/JobFinder/models"
)

func CreateUrl(baseurl string, jobRequest models.JobRequest) (string, error) {
	// If the job request has skills, add them to the url
	if len(jobRequest.Skills) > 0 {
		var skills []string
		// Convert the skills to strings
		for _, skill := range jobRequest.Skills {
			skill := strings.ToLower(string(skill))
			skills = append(skills, skill)
		}
		// Add the skills to the url
		baseurl = fmt.Sprintf("%sremote-%s-jobs", baseurl, strings.Join(skills, "+"))
	}

	u, err := url.Parse(baseurl)
	if err != nil {
		return "", err
	}

	query := u.Query()

	addQuery(&query, "min_salary", strconv.FormatFloat(jobRequest.MinimumSalaryExpectation, 'f', -1, 64))
	addQuery(&query, "location", jobRequest.Location)
	addQuery(&query, "order_by", "date")

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func addQuery(query *url.Values, name string, value string) {
	if value == "" || name == "" || query == nil {
		return
	}

	query.Add(name, value)
}
