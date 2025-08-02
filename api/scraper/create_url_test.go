package scraper

import (
	"fmt"
	"testing"

	"github.com/FranMT-S/JobFinder/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateSalaryRangeWeWorkremotely(t *testing.T) {

	ttc := []struct {
		name       string
		jobRequest models.JobRequest
		expected   []string
	}{
		{
			name: "not must be containts salary",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 0,
				MaximumSalaryExpectation: -1,
			},
			expected: []string{},
		},
		{
			name: "min beetwen 1 and 10000",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 5000,
				MaximumSalaryExpectation: -1,
			},
			expected: []string{WORK_REMOTELY_RANGE_SALARY_1},
		},
		{
			name: "salary level 1",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 10000,
				MaximumSalaryExpectation: 25000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_1,
			},
		},
		{
			name: "salary level 2",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 25000,
				MaximumSalaryExpectation: 49000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_2,
			},
		},
		{
			name: "salary level 3",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 49000,
				MaximumSalaryExpectation: 75000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_3,
			},
		},
		{
			name: "salary level 4",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 75000,
				MaximumSalaryExpectation: 100000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_4,
			},
		},
		{
			name: "salary level 5",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 100000,
				MaximumSalaryExpectation: 120000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_5,
			},
		},
		{
			name: "salary level 1,2,3,4,5",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 0,
				MaximumSalaryExpectation: 120000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_1,
				WORK_REMOTELY_RANGE_SALARY_2,
				WORK_REMOTELY_RANGE_SALARY_3,
				WORK_REMOTELY_RANGE_SALARY_4,
				WORK_REMOTELY_RANGE_SALARY_5,
			},
		},
		{
			name: "salary level 2,3,4",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 25000,
				MaximumSalaryExpectation: 100000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_2,
				WORK_REMOTELY_RANGE_SALARY_3,
				WORK_REMOTELY_RANGE_SALARY_4,
			},
		},
		{
			name: "salary level 2,3",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 25000,
				MaximumSalaryExpectation: 75000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_2,
				WORK_REMOTELY_RANGE_SALARY_3,
			},
		},
		{
			name: "salary above or equals 49000",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: 49000,
				MaximumSalaryExpectation: -1,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_3,
				WORK_REMOTELY_RANGE_SALARY_4,
				WORK_REMOTELY_RANGE_SALARY_5,
			},
		},
		{
			name: "salary below 75000",
			jobRequest: models.JobRequest{
				MinimumSalaryExpectation: -1,
				MaximumSalaryExpectation: 75000,
			},
			expected: []string{
				WORK_REMOTELY_RANGE_SALARY_1,
				WORK_REMOTELY_RANGE_SALARY_2,
				WORK_REMOTELY_RANGE_SALARY_3,
			},
		},
	}

	for _, test := range ttc {

		salaryRange := CreateSalaryRangeWeWorkremotely(test.jobRequest)
		fmt.Println(salaryRange)
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, salaryRange)
		})
	}

}
