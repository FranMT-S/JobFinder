package scraper

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/FranMT-S/JobFinder/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type WorkRemotelyScraper struct{}

func NewWorkRemotelyScraper() *WorkRemotelyScraper {
	return &WorkRemotelyScraper{}
}

// func (s WorkRemotelyScraper) GetJobs(urlServer string) ([]models.Job, error) {
// 	c := SetupBasicCollector()
// 	var err error
// 	jobs := make([]models.Job, 0)

// 	c.OnRequest(func(r *colly.Request) {
// 		r.Headers.Set("Content-Type", "text/html; charset=UTF-8")
// 	})

// 	c.OnResponse(func(r *colly.Response) {
// 		fmt.Println("Visiting", r.Request.URL.String())

// 	})

// 	c.OnError(func(r *colly.Response, err error) {
// 		fmt.Println("Error visiting", r.Request.URL.String(), err)
// 	})

// 	c.OnHTML("section.jobs a div.new-listing", func(e *colly.HTMLElement) {

// 		if len(jobs) >= 4 {
// 			return
// 		}

// 		e.DOM.Each(s.findJobs(func(job models.Job, errf error) {
// 			if errf != nil {
// 				log.Println("Error parsing job:", errf)
// 				return
// 			}

// 			jobs = append(jobs, job)
// 		}))
// 	})

// 	err = RetryFunc(func() error {
// 		return c.Visit(urlServer)
// 	}, 3)

// 	if err != nil {
// 		return nil, fmt.Errorf("error visiting url: %v", err)
// 	}

// 	return jobs, nil
// }

// ScrapperJob is an interface that represents a job scrapper
// urlServer is the url of the server that will be scraped
// wg is a WaitGroup that will be used to wait for the scraper to finish
// chJob is a channel that will be used to send the jobs to the main function
// chError is a channel that will be used to send the errors to the main function
// maxJobs is the maximum number of jobs that will be scraped
// maxJobsParse is the maximum number of tags of job that must be try to parsed
func (s WorkRemotelyScraper) GetJobs(ctx context.Context,
	url string,
	wg *sync.WaitGroup,
	chJob chan<- models.Job,
	chError chan<- error,
	maxJobs int,
	maxJobsParse int,
) {
	c := SetupBasicCollector()
	var err error
	jobsSended := 0
	jobsParsed := 0

	defer wg.Done()
	isFinished := false

	go func() {
		<-ctx.Done()
		isFinished = true
	}()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "text/html; charset=UTF-8")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visiting", r.Request.URL.String())

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error visiting", r.Request.URL.String(), err)
	})

	c.OnHTML("section.jobs a div.new-listing", func(e *colly.HTMLElement) {
		if isFinished {
			e.Request.Abort()
			return
		}

		e.DOM.Each(s.findJobs(func(job models.Job, errf error) {
			if isFinished {
				e.Request.Abort()
				return
			}

			jobsParsed++
			if jobsParsed >= maxJobsParse {
				isFinished = true
				return
			}

			if errf != nil {
				log.Println("Error parsing job:", errf)
				return
			}

			absoluteURL := fmt.Sprintf("%s://%s%s", e.Request.URL.Scheme, e.Request.URL.Host, job.Url)
			job.Url = absoluteURL

			chJob <- job
			jobsSended++
			if jobsSended >= maxJobs {
				isFinished = true
				return
			}
		}))
	})

	err = RetryFunc(func() error {
		return c.Visit(url)
	}, 3)

	if err != nil {
		chError <- fmt.Errorf("error visiting url: %v", err)
	}

}

// FindJobs is a function that finds the jobs from the page
// callback is a function that is called when a job is found
// returns a function that can be used in the OnHTML method of the collector of colly
func (scraper WorkRemotelyScraper) findJobs(callback func(models.Job, error)) func(int, *goquery.Selection) {
	var (
		position      string
		minimunSalary float64
		maximumSalary float64
		skills        []models.Skill = make([]models.Skill, 0)
		company       string
		location      []string
		url           string
		createdAt     *time.Time
		levels        []models.Level    = make([]models.Level, 0)
		modalities    []models.Modality = make([]models.Modality, 0)
		isRecentJob   bool
		contractType  models.ContractType = models.ContractNoSpecific
		tags          []string
	)

	return func(i int, e *goquery.Selection) {
		position = scraper.getPosition(e)
		createdAt = scraper.getCreatedAt(e)
		company = scraper.getCompanyAndNewJob(e)
		location = scraper.getLocation(e)
		url = e.Parent().AttrOr("href", "")
		tags, contractType, minimunSalary, maximumSalary, modalities = scraper.parseCategories(e)
		var err error

		if position == "" || url == "" || company == "" {
			err = errors.New("job details are not valid")
		}

		skills = append(skills, scraper.findSkills(position)...)

		isRecentJob = createdAt != nil && createdAt.Day() == time.Now().Day()

		job := models.NewBlankJob()
		job.Host = models.WorkRemotely
		job.Web = "workremotely"
		job.Position = position
		job.Level = levels
		job.MinimumSalary = minimunSalary
		job.MaximumSalary = maximumSalary
		job.Skills = skills
		job.Modalities = modalities
		job.Company = company
		job.Location = location
		job.Url = url
		job.CreatedAt = createdAt
		job.IsRecentJob = isRecentJob
		job.ContractType = contractType
		job.Tags = tags
		job.Categories = make([]models.Category, 0)

		callback(*job, err)
	}
}

func (scraper WorkRemotelyScraper) findSkills(position string) []models.Skill {
	skills := make([]models.Skill, 0)
	skillsMatches := SkillsRegex.FindStringSubmatch(position)
	for _, skill := range skillsMatches {
		skill = strings.TrimSpace(skill)
		if skill != "" {
			skills = append(skills, models.NewSkill(skill))
		}
	}

	return skills
}

func (scraper WorkRemotelyScraper) getPosition(e *goquery.Selection) string {
	return strings.TrimSpace(e.Find("h4.new-listing__header__title").First().Text())
}

func (scraper WorkRemotelyScraper) getCompanyAndNewJob(e *goquery.Selection) string {
	return strings.TrimSpace(e.Find(".new-listing__company-name").First().Text())
}

func (scraper WorkRemotelyScraper) getCreatedAt(e *goquery.Selection) *time.Time {
	dateText := strings.ToLower(e.Find(".new-listing__header__icons__date").First().Text())
	return scraper.parseDateTime(dateText)
}

func (scraper WorkRemotelyScraper) getLocation(e *goquery.Selection) []string {
	return []string{strings.TrimSpace(e.Find(".new-listing__company-headquarters").First().Text())}
}

func (scraper WorkRemotelyScraper) parseCategories(e *goquery.Selection) (tags []string, contractType models.ContractType, minSalary float64, maxSalary float64, modalities []models.Modality) {
	minSalary = -1
	maxSalary = -1
	contractType = models.ContractNoSpecific
	modalities = make([]models.Modality, 0)
	tags = make([]string, 0)

	e.Find(".new-listing__categories__category").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		text = strings.ToLower(text)

		if strings.Contains(text, "featured") || strings.Contains(text, "top 100") {
			return
		}

		if contractType == models.ContractNoSpecific {
			contractType = scraper.getContract(text)
		} else if strings.Contains(text, "anywhere in the world") {
			modalities = append(modalities, models.Remote)
		} else if strings.Contains(text, "$") && (minSalary == -1 || maxSalary == -1) {
			minSalary, maxSalary = scraper.ParseSalary(text)
		} else {
			tags = append(tags, text)
		}
	})

	return tags, contractType, minSalary, maxSalary, modalities
}

func (scraper WorkRemotelyScraper) ParseSalary(text string) (minSalary float64, maxSalary float64) {
	minSalary = -1
	maxSalary = -1
	var err error

	cleanRegex := regexp.MustCompile(`[^0-9. ]`).ReplaceAllString(text, "")

	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	if strings.Contains(text, "or more usd") {
		cleanText := strings.TrimSpace(cleanRegex)

		minSalary, err = strconv.ParseFloat(cleanText, 64)
		if err != nil {
			minSalary = -1
		}
		return minSalary, -1
	}

	cleanRegex = strings.TrimSpace(cleanRegex)
	parts := strings.SplitN(cleanRegex, " ", 2)

	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])

	minSalary, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		minSalary = -1
	}

	maxSalary, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		maxSalary = -1
	}

	return minSalary, maxSalary
}

func (scraper WorkRemotelyScraper) getContract(s string) models.ContractType {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "" {
		return models.ContractNoSpecific
	}

	if strings.Contains(s, "contract") {
		return models.Contract
	}

	if strings.Contains(s, "full-time") || strings.Contains(s, "fulltime") {
		return models.FullTime
	}

	if strings.Contains(s, "part-time") || strings.Contains(s, "partime") {
		return models.PartTime
	}

	return models.ContractNoSpecific
}

func (scraper WorkRemotelyScraper) parseDateTime(dateText string) *time.Time {
	var createdAt *time.Time

	if strings.Contains(dateText, "new") {
		now := time.Now()
		createdAt = &now
		return createdAt
	}

	dateText = strings.TrimSpace(dateText)
	day, err := strconv.Atoi(strings.ReplaceAll(dateText, "d", ""))
	if err != nil {
		log.Println("Error parsing day:", err)
		createdAt = nil
		return createdAt
	}

	date := time.Now().AddDate(0, 0, -day)
	createdAt = &date

	return createdAt
}
