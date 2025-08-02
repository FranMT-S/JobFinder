package models

import "time"

type Modality string

// Modality is a type that represents the modality of the job
var (
	Hybrid     Modality = "hybrid"
	Remote     Modality = "remote"
	Presential Modality = "presential"
	Any        Modality = "any"
	NoSpecific Modality = "no-specific"
)

// Level is a type that represents the level of the job
type Level string

var (
	Junior     Level = "junior"
	Mid        Level = "mid"
	Senior     Level = "senior"
	SemiSenior Level = "semi-senior"
	AnyLevel   Level = "any"
)

var PositionMap = map[string]Level{
	"frontend":  "frontend",
	"backend":   "backend",
	"fullstack": "fullstack",
	"mobile":    "mobile",
	"design":    "design",
	"product":   "product",
	"data":      "data",
	"qa":        "qa",
	"sales":     "sales",
	"marketing": "marketing",
	"business":  "business",
	"hr":        "hr",
	"it":        "it",
	"other":     "other",
}

// Job is a type that represents the fields that a job has
type Job struct {
	Host          HostScrapper `json:"host"`
	Web           string       `json:"web"`
	Url           string       `json:"url"`
	Position      string       `json:"position"`
	MinimumSalary float64      `json:"minimumSalary"`
	MaximumSalary float64      `json:"maximumSalary"`
	Company       string       `json:"company"`
	Level         []Level      `json:"level"`
	Skills        []Skill      `json:"skills"`
	Modalities    []Modality   `json:"modalities"`
	Location      []string     `json:"location"`
	Tags          []string     `json:"tags"`
	ContractType  ContractType `json:"contractType"`
	IsRecentJob   bool         `json:"isRecentJob"`
	CreatedAt     *time.Time   `json:"createdAt"`
	Categories    []Category   `json:"categories"`
}

// NewBlankJob initialize a job with default values and empty slices
func NewBlankJob() *Job {
	return &Job{
		Host:          0,
		Web:           "",
		Url:           "",
		Position:      "",
		MinimumSalary: 0,
		MaximumSalary: 0,
		Company:       "",
		Level:         make([]Level, 0),
		Skills:        make([]Skill, 0),
		Modalities:    make([]Modality, 0),
		Location:      make([]string, 0),
		Tags:          make([]string, 0),
		ContractType:  ContractType(""),
		IsRecentJob:   false,
		CreatedAt:     nil,
		Categories:    make([]Category, 0),
	}
}

// JobResponse is a type that represents the response of the job
type JobResponse struct {
	Jobs  []Job `json:"jobs"`
	Total int   `json:"total"`
}

func NewJobResponse(jobs []Job) *JobResponse {
	return &JobResponse{
		Jobs:  jobs,
		Total: len(jobs),
	}
}

// JobRequest is a type that represents the request of the job with the filters
type JobRequest struct {
	Host                     []HostScrapper `json:"host"`
	Location                 string         `json:"location"`
	Level                    Level          `json:"level"`
	Skills                   []Skill        `json:"skills"`
	Modality                 Modality       `json:"modality"`
	MinimumSalaryExpectation float64        `json:"minimumSalaryExpectation"`
	MaximumSalaryExpectation float64        `json:"maximumSalaryExpectation"`
	Category                 Category       `json:"category"`
}

// IsJobRequestNil is a function that checks if the all values of the job request are nil
func IsJobRequestNil(jobRequest JobRequest) bool {
	return jobRequest.Location == "" && jobRequest.Level == "" && len(jobRequest.Skills) == 0 && jobRequest.Modality == "" && jobRequest.MinimumSalaryExpectation == 0 && jobRequest.MaximumSalaryExpectation == 0 && jobRequest.Category == -1
}
