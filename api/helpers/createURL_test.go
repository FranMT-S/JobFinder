package helpers

import (
	"testing"

	"github.com/FranMT-S/JobFinder/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUrl(t *testing.T) {
	var tests = []struct {
		name     string
		body     models.JobRequest
		expected string
	}{
		{
			name: "test create url",
			body: models.JobRequest{
				Location:                 "Worldwide",
				Level:                    "Senior",
				Skills:                   []models.Skill{"golang", "Python"},
				Modality:                 "Remote",
				MinimumSalaryExpectation: 10000,
				MaximumSalaryExpectation: 100000,
				Category:                 models.Backend,
			},
			expected: "https://remoteok.com/remote-golang+python-jobs?location=Worldwide&min_salary=10000&order_by=date",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url, err := CreateUrl("https://remoteok.com/", test.body)
			if err != nil {
				t.Errorf("error creating job request: %v", err)
			} else {
				assert.Equal(t, test.expected, url)
			}
		})
	}
}
