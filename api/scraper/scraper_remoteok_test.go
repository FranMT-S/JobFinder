package scraper

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/FranMT-S/JobFinder/models"
	"github.com/stretchr/testify/assert"
)

func TestScrapper(t *testing.T) {

	var createdAt = time.Date(2025, 6, 16, 22, 15, 2, 0, time.UTC)
	var createdAt2 = time.Date(2025, 5, 31, 0, 0, 8, 0, time.UTC)
	var testcases = []struct {
		expected    []models.Job
		errResponse error
		statusCode  int
		name        string
	}{
		{
			name: "remoteOk test get jobs",
			expected: []models.Job{
				{
					Position:      "Senior Platform Engineer open across ANZ",
					Level:         []models.Level{models.Senior},
					MinimumSalary: 65000,
					MaximumSalary: 125000,
					Skills: []models.Skill{
						"golang",
						"senior",
					},
					Modalities: []models.Modality{
						models.Remote,
					},
					Company:   "Canva",
					Location:  []string{"Probably worldwide"},
					Url:       "/remote-jobs/remote-senior-platform-engineer-open-across-anz-canva-1093391",
					CreatedAt: &createdAt,
				},
				{
					Position:      "Software Engineer II",
					Level:         []models.Level{models.AnyLevel},
					MinimumSalary: 70000,
					MaximumSalary: 110000,
					Skills: []models.Skill{
						"software",
						"system",
						"front-end",
						"security",
						"full-stack",
						"developer",
						"web",
						"scrum",
						"devops",
						"javascript",
						"financial",
						"typescript",
						"mongo",
						"api",
						"management",
						"health",
						"engineering",
					},
					Modalities: []models.Modality{
						models.Presential,
					},
					Company:   "Everbridge",
					Location:  []string{"United States"},
					Url:       "/remote-jobs/remote-software-engineer-ii-everbridge-1093283",
					CreatedAt: &createdAt2,
				},
			},
		},
	}

	server, host, err := MockServer("../testdata/remoteOkSampleData.html")
	if err != nil {
		t.Fatalf("failed reading the html file: %v", err)
	}

	defer DisabledMockServer(server)

	ctx := context.Background()
	wg := &sync.WaitGroup{}
	chJob := make(chan models.Job)
	chError := make(chan error)
	scraper := NewRemoteOkScraper()
	data := make([]models.Job, 0)

	wg.Add(1)
	go func() {
		for job := range chJob {
			data = append(data, job)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		scraper.GetJobs(ctx, server.URL, wg, chJob, chError, 2, 2)
		close(chJob)
		close(chError)
	}()

	wg.Wait()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, len(tt.expected), len(data), "the number of jobs is not the same")
			assert.NoError(t, err, "the error is not nil")
			for i, jobExpected := range tt.expected {
				absoluteURL := fmt.Sprintf("%s://%s%s", "http", host, jobExpected.Url)
				assert.Equal(t, jobExpected.Position, data[i].Position, "the position is not the same")
				assert.Equal(t, absoluteURL, data[i].Url, "the url is not the same")
				assert.Equal(t, jobExpected.Company, data[i].Company, "the company is not the same")
				assert.Equal(t, jobExpected.Level, data[i].Level, "the level is not the same")
				assert.Equal(t, jobExpected.MinimumSalary, data[i].MinimumSalary, "the minimum salary is not the same")
				assert.Equal(t, jobExpected.MaximumSalary, data[i].MaximumSalary, "the maximum salary is not the same")
				assert.Equal(t, jobExpected.Skills, data[i].Skills, "the skills are not the same")
				assert.Equal(t, jobExpected.Modalities, data[i].Modalities, "the modalities are not the same")
				assert.Equal(t, jobExpected.Location, data[i].Location, "the location is not the same")
				assert.Equal(t, jobExpected.CreatedAt, data[i].CreatedAt, "the created at is not the same")

			}
		})
	}
}
