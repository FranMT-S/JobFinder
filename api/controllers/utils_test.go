package controllers

import (
	"testing"

	"github.com/FranMT-S/JobFinder/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateJobRequest(t *testing.T) {
	var tests = []struct {
		name     string
		body     []byte
		expected models.JobRequest
	}{
		{
			name: "test create job request",
			body: []byte(`{
				"location": "United States", 
				"level": "Senior", 
				"skills": ["Go", "Python"], 
				"modalities": "Remote", 
				"minimumSalaryExpectation": 10000, 
				"maximumSalaryExpectation": 100000, 
				"position": "Software Engineer"
			}`),
			expected: models.JobRequest{
				Location:                 "United States",
				Level:                    "Senior",
				Skills:                   []models.Skill{"Go", "Python"},
				Modality:                 "Remote",
				MinimumSalaryExpectation: 10000,
				MaximumSalaryExpectation: 100000,
				Category:                 models.Backend,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jobRequest, err := CreateJobRequest(test.body)
			if err != nil {
				t.Errorf("error creating job request: %v", err)
			} else {
				assert.Equal(t, test.expected, jobRequest)
			}
		})
	}
}

func TestCreateJobRequestError(t *testing.T) {
	var tests = []struct {
		name string
		body []byte
	}{
		{
			name: "error because no data valid send",
			body: []byte(`{
				"locasdfasdation": "United States", 
				"a": "Senior", 
				"sdfg": ["Go", "Python"], 
				"sdfgasd": "Remote", 
				"sdfg": 10000, 
				"asdf": 100000, 
				"sdfg": "Software Engineer"
			}`),
		},
		{
			name: "error because body is incorrect",
			body: []byte(`{
				"location": "United States", 
				"sdfg":asdfasd
			}`),
		},
		{
			name: "error because data type is incorrect",
			body: []byte(`{
				"location": []
			}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := CreateJobRequest(test.body)
			assert.Error(t, err)
		})
	}
}
