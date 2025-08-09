package scraper

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"html"

	"github.com/FranMT-S/JobFinder/helpers"
	"github.com/FranMT-S/JobFinder/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type WorkRemotelyScraper struct{}

func NewWorkRemotelyScraper() *WorkRemotelyScraper {
	return &WorkRemotelyScraper{}
}

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

	jobPageCollector := c.Clone()

	c.OnHTML("section.jobs a div.new-listing", func(e *colly.HTMLElement) {
		if isFinished {
			e.Request.Abort()
			return
		}

		e.DOM.Each(func(i int, s *goquery.Selection) {
			urlJob := s.Parent().AttrOr("href", "")
			urlJob = fmt.Sprintf("%s://%s%s", e.Request.URL.Scheme, e.Request.URL.Host, urlJob)
			fmt.Println("Visiting", urlJob)
			jobPageCollector.Visit(urlJob)
			jobsParsed++
			if jobsParsed >= maxJobsParse {
				isFinished = true
				return
			}
		})

	})

	jobPageCollector.OnHTML(".lis-container", func(e *colly.HTMLElement) {

		job := models.NewBlankJob()
		job.Host = models.WorkRemotely
		job.Web = "workremotely"

		// parse url
		absoluteURL := e.Request.URL.String()
		job.Url = absoluteURL

		// collect data from details card
		e.ForEach("li.lis-container__job__sidebar__job-about__list__item", func(i int, ch *colly.HTMLElement) {
			text := ch.DOM.Text()
			text = strings.ToLower(text)
			dataText := ch.ChildText("span")
			switch {
			case strings.Contains(text, "posted on"):
				// parse created at
				createdAt := s.parseDateTime(dataText)
				job.CreatedAt = createdAt
			case strings.Contains(text, "job type"):
				// parse contract type
				job.ContractType = s.getContract(dataText)
			case strings.Contains(text, "salary"):
				// parse salary
				minSalary, maxSalary := s.ParseSalary(dataText)
				job.MinimumSalary = minSalary
				job.MaximumSalary = maxSalary
			case strings.Contains(text, "category"):
				// parse categories
				category := dataText
				category = strings.TrimSpace(category)
				category = strings.ToLower(category)
				switch category {
				case "design":
					job.Categories = append(job.Categories, models.Design)
				case "full stack":
					job.Categories = append(job.Categories, models.FullStack)
				case "sysadmin":
					job.Categories = append(job.Categories, models.SysAdmin)
				case "devops and sysadmin":
					job.Categories = append(job.Categories, models.DevOps)
				case "front end":
					job.Categories = append(job.Categories, models.FrontEnd)
				case "backend":
					job.Categories = append(job.Categories, models.Backend)
				}
			case strings.Contains(text, "skills"):
				// parse skills
				ch.ForEach(".boxes a", func(i int, itag *colly.HTMLElement) {
					tag := itag.DOM.Text()
					tag = strings.TrimSpace(tag)
					tag = strings.ToLower(tag)
					job.Skills = append(job.Skills, models.NewSkill(tag))
				})
			case strings.Contains(text, "country"):
				// parse locations
				ch.ForEach(".boxes a", func(i int, icountry *colly.HTMLElement) {
					country := icountry.DOM.Text()
					country = strings.TrimSpace(country)
					if helpers.StartsWithRegionalCharacters(country) {
						splitString := strings.SplitN(country, " ", 2)
						if len(splitString) == 2 {
							country = splitString[1]
							country = strings.TrimSpace(country)
						}
					}

					job.Location = append(job.Location, country)
				})
			}
		})

		// parse job name
		position := e.ChildText(".lis-container__header__hero__company-info__title")
		position = strings.TrimSpace(position)

		// parse description
		description, err := e.DOM.Find(".lis-container__job__content__description").Html()
		if err != nil {
			job.Description = ""
		} else {
			job.Description = description
		}

		// parse name company
		company := e.ChildText(".lis-container__job__sidebar__companyDetails__info__title")
		company = strings.TrimSpace(company)

		level := GetLevels(position, job.Skills)

		job.Company = company
		job.Position = position
		job.IsRecentJob = s.isRecentJob(job)
		job.Level = level

		job.Url = html.EscapeString(job.Url)

		chJob <- *job
		jobsSended++
		if jobsSended >= maxJobs {
			isFinished = true
			return
		}
	})

	err = RetryFunc(func() error {
		return c.Visit(url)
	}, 3)

	if err != nil {
		chError <- fmt.Errorf("error visiting url: %v", err)
	}
}

func (scraper WorkRemotelyScraper) isRecentJob(job *models.Job) bool {
	day := time.Now().Day()
	month := time.Now().Month()
	year := time.Now().Year()
	return job.CreatedAt != nil && job.CreatedAt.Day() == day && job.CreatedAt.Month() == month && job.CreatedAt.Year() == year
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
	dateText = strings.TrimSpace(dateText)
	dateText = strings.ToLower(dateText)

	splitDate := strings.Split(dateText, " ")

	if len(splitDate) < 2 {
		return nil
	}

	timeStr, interval := splitDate[0], splitDate[1]
	timeInt, err := strconv.Atoi(timeStr)
	if err != nil {
		return nil
	}

	var createdAt *time.Time = nil

	now := time.Now()
	switch interval {
	case "days", "day":
		now = now.AddDate(0, 0, -timeInt)
		createdAt = &now
	case "hours", "hour":
		now = now.Add(time.Hour * time.Duration(timeInt))
		createdAt = &now
	}

	return createdAt

}
