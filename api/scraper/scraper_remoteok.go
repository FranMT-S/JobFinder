package scraper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/FranMT-S/JobFinder/config"
	"github.com/FranMT-S/JobFinder/helpers"
	"github.com/FranMT-S/JobFinder/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// remoteOkDetails is a struct that contains the information of a job from RemoteOk
type remoteOkDetails struct {
	positionName string
	company      string
	location     []string
	isRecentJob  bool
	minSalary    float64
	maxSalary    float64
	url          string
}

type RemoteOkScraper struct{}

func NewRemoteOkScraper() *RemoteOkScraper {
	return &RemoteOkScraper{}
}

// ScrapperJob is an interface that represents a job scrapper
// urlServer is the url of the server that will be scraped
// wg is a WaitGroup that will be used to wait for the scraper to finish
// chJob is a channel that will be used to send the jobs to the main function
// chError is a channel that will be used to send the errors to the main function
// maxJobs is the maximum number of jobs that will be scraped
// maxJobsParse is the maximum number of tags of job that must be try to parsed
func (s RemoteOkScraper) GetJobs(ctx context.Context,
	url string,
	wg *sync.WaitGroup,
	chJob chan<- models.Job,
	chError chan<- error,
	maxJobs int,
	maxJobsParse int,
) {

	c := SetupBasicCollector()
	isFinished := false

	counter := 0

	defer wg.Done()
	jobsSended, jobsParsed := 0, 0
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error visiting", r.Request.URL.String(), err)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "text/html; charset=UTF-8")
		fmt.Println("Visiting", r.URL.String())
	})

	go func() {
		<-ctx.Done()
		isFinished = true
	}()

	c.OnResponse(func(r *colly.Response) {
		// add table tags to the html to fix the parsing because the response is just "tr" tags
		html := string(r.Body)
		html = "<table>" + html + "</table>"
		doc, errint := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
		if errint != nil {
			chError <- fmt.Errorf("error parsing html: %v", errint)
			return
		}

		if config.ENVIRONMENT == "development" {
			helpers.SaveHTMLResponse(html, "remoteok.html")
		}

		doc.Find("tr.job").Each(s.findJobs(func(job models.Job, err error) {
			if isFinished {
				r.Request.Abort()

				return
			}

			jobsParsed++
			if jobsParsed > maxJobsParse {
				isFinished = true
				return
			}

			absoluteURL := fmt.Sprintf("%s://%s%s", r.Request.URL.Scheme, r.Request.URL.Host, job.Url)
			job.Url = absoluteURL

			counter++
			if counter >= 3 {
				<-time.After(1 * time.Second)
				counter = 0
			}

			chJob <- job
			jobsSended++
			if jobsSended >= maxJobs {
				isFinished = true
				return
			}
		}))
	})

	errVisit := c.Visit(url)
	if errVisit != nil {
		chError <- fmt.Errorf("error visiting url: %v", errVisit)
	}

}

// FindJobs is a function that finds the jobs from the page
// callback is a function that is called when a job is found
// returns a function that can be used in the OnHTML method of the collector of colly
func (scraper RemoteOkScraper) findJobs(callback func(models.Job, error)) func(int, *goquery.Selection) {
	var (
		position      string
		minimunSalary float64
		maximumSalary float64
		skills        []models.Skill = make([]models.Skill, 0)
		company       string
		location      []string
		url           string
		createdAt     *time.Time
		tags          []string
		categories    []models.Category
	)

	return func(i int, e *goquery.Selection) {
		jobDetails := scraper.parseJobDetails(e)
		position = jobDetails.positionName
		minimunSalary = jobDetails.minSalary
		maximumSalary = jobDetails.maxSalary
		company = jobDetails.company
		location = jobDetails.location
		url = jobDetails.url
		var err error

		if position == "" || url == "" || company == "" {
			err = errors.New("job details are not valid")
		}

		skills = scraper.parseSkills(e)
		createdAt, _ = scraper.parseDateTime(e)

		levels := GetLevels(position, skills)
		modalities := GetModalities(location)
		categories = scraper.GetCategories(position, skills)

		job := models.NewBlankJob()
		job.Host = models.RemoteOk
		job.Web = "remoteok"
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
		job.IsRecentJob = jobDetails.isRecentJob
		job.ContractType = models.ContractNoSpecific
		job.Tags = tags
		job.Categories = categories

		callback(*job, err)
	}
}

// parseJobDetails is a function that parses the job details from the page
// e is the element the container that contains the job details
// returns a remoteOkDetails struct
func (scraper RemoteOkScraper) parseJobDetails(e *goquery.Selection) remoteOkDetails {
	var (
		positionName string
		company      string
		location     []string
		isRecentJob  bool
		minSalary    float64
		maxSalary    float64
		url          string
	)

	companyContainer := e.Find(JOB_DETAILS_TAG).First()

	if companyContainer == nil {
		return remoteOkDetails{}
	}

	positionName, link := scraper.getLinkAndPositionName(companyContainer)
	company, isRecentJob = scraper.getCompanyAndNewJob(companyContainer)
	location, minSalary, maxSalary = scraper.getLocationAndSalary(companyContainer)
	url = strings.TrimSpace(link)

	return remoteOkDetails{
		positionName: positionName,
		company:      company,
		location:     location,
		isRecentJob:  isRecentJob,
		minSalary:    minSalary,
		maxSalary:    maxSalary,
		url:          url,
	}
}

// parseSkills is a function that parses the skills from the page
// e is the element the container that contains the skills tags
// returns a slice of skills
func (scraper RemoteOkScraper) parseSkills(e *goquery.Selection) []models.Skill {
	skills := make([]models.Skill, 0)
	e.Find(SKILLS_TAG).Each(func(_ int, s *goquery.Selection) {
		skills = append(skills, models.NewSkill(s.Text()))
	})

	return skills
}

// parseDateTime is a function that parses the date and time from the page
// e is the element the container that contains the date and time tag
// returns a pointer to a time.Time and an error
func (scraper RemoteOkScraper) parseDateTime(e *goquery.Selection) (*time.Time, error) {
	timeTag := e.Find(DATE_TIME_TAG).First()

	timeText := timeTag.AttrOr("datetime", "")

	dateTime, err := time.ParseInLocation(time.RFC3339, timeText, time.UTC)
	if err != nil {
		return nil, err
	}
	return &dateTime, nil
}

// getLinkAndPositionName is a function that gets the link and position name from the page
// e is the element the tag element that contains the link and position name
// returns the position name and the link
func (scraper RemoteOkScraper) getLinkAndPositionName(e *goquery.Selection) (positionName string, link string) {
	container := e.Find("a[itemprop='url']").First()
	link, exist := container.Attr("href")
	if exist {
		link = strings.TrimSpace(link)
	}
	positionName = strings.TrimSpace(container.Find("h2").First().Text())

	return positionName, link
}

// getCompanyAndNewJob is a function that gets the company and if the job is recent from the page
// e is the element the tag element that contains the company and the new job tag
// returns the company and if the job is recent
func (scraper RemoteOkScraper) getCompanyAndNewJob(e *goquery.Selection) (company string, isRecentJob bool) {
	container := e.Find("span[itemprop='hiringOrganization']").First()
	company = strings.TrimSpace(container.Find("h3").First().Text())
	isRecentJob = strings.Contains(container.Find("img[alt='This job has just been']").First().AttrOr("alt", ""), "has just been")
	return company, isRecentJob
}

// getLocationAndSalary is a function that gets the location and salary from the page
// e is the element the tag element that contains the location and salary
// returns the location, the minimum salary and the maximum salary
func (scraper RemoteOkScraper) getLocationAndSalary(e *goquery.Selection) (locations []string, minSalary float64, maxSalary float64) {
	minSalary, maxSalary = -1, -1

	e.Find(".location").Each(func(_ int, s *goquery.Selection) {
		locationText := helpers.CleanEmojiAndTrimString(s.Text())
		isSalaryTag := regexp.MustCompile(`\$\d+`).MatchString(locationText)

		if isSalaryTag {
			minSalary, maxSalary = scraper.getSalaryFromLocation(locationText)
		}

		if !isSalaryTag {
			locations = append(locations, locationText)
		}

	})

	return locations, minSalary, maxSalary
}

// getSalaryFromLocation is a function that gets the salary from the location text
// locationText is the text of the location tag
// returns the minimum salary and the maximum salary
func (scraper RemoteOkScraper) getSalaryFromLocation(locationText string) (minSalary float64, maxSalary float64) {
	minSalary, maxSalary = -1, -1

	parts := strings.Split(locationText, " ")

	parts[0] = helpers.CleanOnlyNumbersFromString(parts[0])
	parts[2] = helpers.CleanOnlyNumbersFromString(parts[2])
	min, err := strconv.ParseFloat(parts[0], 64)

	if err != nil {
		minSalary = -1
	}

	// remote ok uses k as a multiplier for the salary
	minSalary = min * 1000
	max, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		maxSalary = -1
	}
	maxSalary = max * 1000

	return minSalary, maxSalary
}

func (scraper RemoteOkScraper) GetCategories(position string, skills []models.Skill) []models.Category {
	mapCategories := map[models.Category]struct{}{}

	category := FindCategory(position)
	if category != models.NotCategory {
		mapCategories[category] = struct{}{}
	}

	for _, skill := range skills {
		category := models.GetSkillCategory(skill)
		if category != models.NotCategory {
			mapCategories[category] = struct{}{}
		}
	}

	categories := make([]models.Category, 0, len(mapCategories))
	for category := range mapCategories {
		categories = append(categories, category)
	}

	return categories
}
