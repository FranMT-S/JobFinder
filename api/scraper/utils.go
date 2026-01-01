package scraper

import (
	"regexp"
	"strings"

	"github.com/FranMT-S/JobFinder/models"
)

// SkillsRegex is a regex that matches the most common skills
var SkillsRegex = regexp.MustCompile(`golang|react|angular|vue|javascript|python|java|c#|php|ruby|swift|kotlin|typescript|next.js|nestjs|node.js|node|next|nextjs|go|rust`)
var CategoriesRegex = regexp.MustCompile(`fullstack|frontend|backend|devops`)

// analyzeLevel is a function that analyzes the level of the job
// s is the string that contains the level
// returns a pointer to a level
func AnalyzeLevel(s string) *models.Level {

	if s == "" {
		return nil
	}

	if strings.Contains(s, "senior") || strings.Contains(s, "principal") || s == "sr" || strings.Contains(s, "head") {
		return &models.Senior
	}

	if strings.Contains(s, "junior") || strings.Contains(s, "intern") {
		return &models.Junior
	}

	if strings.Contains(s, "mid-level") || s == "mid" {
		return &models.Mid
	}

	if strings.Contains(s, "semi-senior") || strings.Contains(s, "semi-lead") || strings.Contains(s, "semi-principal") || strings.Contains(s, "semi-mid") || s == "ssr" || s == "ssr-senior" {
		return &models.SemiSenior
	}

	return nil
}

// getModality is a function that gets the modality from the location
// location is a slice of strings that contains the location
// if the location contains "remote" or "worldwide" the job is remote
// if the location is empty, the job is no specific
// returns a slice of modalities
func GetModalities(location []string) (modalities []models.Modality) {

	isRemote := false
	for _, l := range location {
		l = strings.TrimSpace(strings.ToLower(l))
		if strings.Contains(l, "remote") || strings.Contains(l, "worldwide") || strings.Contains(l, "any") {
			isRemote = true
			break
		}
	}

	if isRemote {
		modalities = append(modalities, models.Remote)
	}

	// if the job is not remote, check if it has a location
	if !isRemote && len(location) > 0 {
		modalities = append(modalities, models.Presential)
	}

	// if the job is not remote and has no location, return no specific
	if !isRemote && len(location) == 0 {
		modalities = append(modalities, models.NoSpecific)
	}

	return modalities
}

// GetLevels is a function that gets the levels from the title job and the skills
// titleJob is the title of the job that will be analyzed to get the level
// skills is a slice of skills that will be analyzed to get the level
// returns a slice of levels
func GetLevels(titleJob string, skills []models.Skill) []models.Level {

	levelsMap := map[models.Level]bool{}

	levels := make([]models.Level, 0)

	// analyze the skills
	for _, skill := range skills {
		level := AnalyzeLevel(string(skill))
		if level != nil {
			levelsMap[*level] = true
		}
	}

	// analyze the title job
	titleJob = strings.ToLower(titleJob)
	level := AnalyzeLevel(titleJob)
	if level != nil {
		levelsMap[*level] = true
	}

	// if no level was found, return any level
	if len(levelsMap) == 0 {
		levelsMap[models.AnyLevel] = true
	}

	for level := range levelsMap {
		levels = append(levels, level)
	}

	return levels
}

// RetryFunc is a function that retries a function a number of times
func RetryFunc(f func() error, attempts int) (err error) {
	for i := 0; i < attempts; i++ {
		if err := f(); err != nil {
			continue
		}
		return nil
	}

	return err
}

func FindCategory(text string) models.Category {
	text = strings.TrimSpace(strings.ToLower(text))

	if text == "" {
		return models.NotCategory
	}

	switch {
	case
		strings.Contains(text, "frontend") ||
			strings.Contains(text, "front-end") ||
			strings.Contains(text, "front end"):
		return models.FrontEnd
	case
		strings.Contains(text, "backend") ||
			strings.Contains(text, "back-end") ||
			strings.Contains(text, "back end"):
		return models.Backend
	case
		strings.Contains(text, "fullstack") ||
			strings.Contains(text, "full-stack") ||
			strings.Contains(text, "full stack"):
		return models.FullStack
	case
		strings.Contains(text, "data-science") ||
			strings.Contains(text, "data science"):
		return models.DataScience
	case
		strings.Contains(text, "testing") ||
			strings.Contains(text, "test") ||
			strings.Contains(text, "tester"):
		return models.Testing
	case
		strings.Contains(text, "architect") ||
			strings.Contains(text, "arch"):
		return models.Architect
	case
		strings.Contains(text, "analyst") ||
			strings.Contains(text, "anal"):
		return models.Analyst
	case strings.Contains(text, "database") ||
		strings.Contains(text, "db"):
		return models.Database
	default:
		return models.NotCategory
	}
}
