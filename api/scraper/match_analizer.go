package scraper

import (
	"strings"

	"github.com/FranMT-S/JobFinder/models"
)

var points = map[models.Level]map[models.Level]float64{
	models.Junior: {
		models.Junior:     100,
		models.Mid:        60,
		models.SemiSenior: 60,
		models.Senior:     20,
	},
	models.Mid: {
		models.Junior:     40,
		models.Mid:        100,
		models.SemiSenior: 100,
		models.Senior:     80,
	},
	models.SemiSenior: {
		models.Junior:     40,
		models.Mid:        100,
		models.SemiSenior: 100,
		models.Senior:     80,
	},
	models.Senior: {
		models.Junior:     20,
		models.Mid:        80,
		models.SemiSenior: 80,
		models.Senior:     100,
	},
}

// MatchAnalizer is a function that analyzes the match between a job and a job request
// salary represent 40% of the match
// skills represent 20% of the match
// modalities represent 20% of the match
// levels represent 10% of the match
// position represent 10% of the match
// jobRequest is the job request to be analyzed
// if any of the values of the job request are nil o len 0, the match is 100% of the percentage of that value
// returns a float64 that represents the match between the job and the job request
func MatchAnalizer(job models.Job, jobRequest models.JobRequest) models.MatchAnalizer {
	salaryPorcent := 0.4
	skillsPorcent := 0.2
	modalitiesPorcent := 0.2
	levelsPorcent := 0.1
	positionPorcent := 0.1
	salaryMessage := ""
	salaryPoints := 100.0

	if jobRequest.MinimumSalaryExpectation > 0 {
		salaryPoints, salaryMessage = analizeSalary(job, jobRequest)
	}

	salary := salaryPoints * salaryPorcent
	skillsPoints, skillMatches := analizeSkills(job, jobRequest)
	skillsPoints = skillsPoints * skillsPorcent
	modalities := analizeModalities(job, jobRequest) * modalitiesPorcent
	levels := analizeLevels(job, jobRequest) * levelsPorcent
	position := analizePosition(job, jobRequest) * positionPorcent

	return models.MatchAnalizer{
		PorcentSalary:     salary,
		PorcentSkills:     skillsPoints,
		PorcentModalities: modalities,
		PorcentLevels:     levels,
		PorcentPosition:   position,
		TotalPorcent:      salary + skillsPoints + modalities + levels + position,
		SkillMatches:      skillMatches,
		SalaryMessage:     salaryMessage,
	}
}

// analizeSalary is a function that analyzes the salary of the job
// if salary is less or equals to 0 it's considered that the salary is not specified return 100
// if the salary of the job is between the salary of the job request, returns 100
// return number between 0 and 100
func analizeSalary(job models.Job, jobRequest models.JobRequest) (float64, string) {
	minSalary := jobRequest.MinimumSalaryExpectation
	maxSalary := jobRequest.MaximumSalaryExpectation
	jobMinSalary := job.MinimumSalary
	jobMaxSalary := job.MaximumSalary

	if minSalary <= 0 && maxSalary <= 0 {
		return 0, ""
	}

	if minSalary >= maxSalary {
		return 0, "salary wrong range"
	}

	if jobMinSalary <= 0 {
		return 0, "salary not specified"
	}

	if jobMinSalary >= minSalary {

		if maxSalary >= jobMaxSalary {
			return 100, "Perfect, up to expected"
		}

		if maxSalary <= jobMaxSalary {
			return 100, "Perfect"
		}
	}

	if jobMinSalary < minSalary {
		rangeSalary := maxSalary - minSalary
		percentBelowMin := (minSalary - jobMinSalary) / rangeSalary * 100

		if maxSalary == -1 && jobMaxSalary == -1 {
			if percentBelowMin <= 10 {
				return 90, "Acceptable"
			} else if percentBelowMin <= 25 {
				return 75, "Low Start, Good Potential"
			} else {
				return 60, "Too Low Start"
			}
		}

		if maxSalary >= jobMaxSalary {
			// the max salary of the job is greater or equals to the max salary of the job request
			if percentBelowMin <= 10 {
				return 90, "Acceptable"
			} else if percentBelowMin <= 25 {
				return 75, "Low Start, Good Potential"
			} else {
				return 60, "Too Low Start"
			}
		}

		if maxSalary < jobMaxSalary {
			// the max salary of the job is less than the max salary of the job request
			missingMax := (maxSalary - jobMaxSalary) / rangeSalary * 100
			score := 100 - percentBelowMin - missingMax

			switch {
			case score >= 80:
				return score, "Good"
			case score >= 70:
				return score, "Fair"
			default:
				return score, "Below Expectations"
			}
		}
	}

	return 0, ""
}

// analizeSkills is a function that analyzes the skills of the job
// if the job has no skills, return 100
// if the job has skills, return 0
// return number between 0 and 100
func analizeSkills(job models.Job, jobRequest models.JobRequest) (float64, []string) {
	skillMatchesMap := make(map[string]struct{})

	for _, skill := range jobRequest.Skills {
		skill := strings.ToLower(string(skill))

		for _, jobSkill := range job.Skills {
			jobSkill := string(jobSkill)
			if skill == jobSkill {
				skillMatchesMap[string(skill)] = struct{}{}
			}
		}
	}

	skillsMatches := make([]string, len(skillMatchesMap))
	i := 0
	for skill := range skillMatchesMap {
		skillsMatches[i] = skill
		i++
	}

	if len(skillsMatches) == len(jobRequest.Skills) {
		return 100, skillsMatches
	} else if len(skillsMatches) >= len(jobRequest.Skills)/2 {
		return 50, skillsMatches
	} else if len(skillsMatches) > 0 {
		return 20, skillsMatches
	}

	return 0, skillsMatches
}

// analizeLevels is a function that analyzes the levels of the job
// if the job has no levels, return 100
// if the job has levels, return 0
// return number between 0 and 100
func analizeLevels(job models.Job, jobRequest models.JobRequest) float64 {
	if jobRequest.Level == "" || jobRequest.Level == models.AnyLevel {
		return 100
	}

	levelPoint := map[models.Level]int{
		models.Junior:     1,
		models.Mid:        2,
		models.SemiSenior: 2,
		models.Senior:     3,
	}

	maxLevel := models.Junior
	for _, level := range job.Level {
		v, ok := levelPoint[level]
		if ok {

			if level == jobRequest.Level {
				return 100
			}

			if v > levelPoint[maxLevel] {
				maxLevel = level
			}
		}
	}

	return points[jobRequest.Level][maxLevel]
}

// analizeModalities is a function that analyzes the modalities of the job
// if jobrequest no has modalities, return 100
// if a job match with the jobrequest modality, return 100
// return number between 0 and 100
func analizeModalities(job models.Job, jobRequest models.JobRequest) float64 {
	if len(jobRequest.Modality) == 0 {
		return 100
	}

	for _, modality := range job.Modalities {
		if modality == jobRequest.Modality {
			return 100
		}
	}

	return 0
}

// analizePosition is a function that analyzes the position of the job
// if jobrequest no has position, return 100
// if a job match with the jobrequest position, return 100
// return number between 0 and 100
func analizePosition(job models.Job, jobRequest models.JobRequest) float64 {
	if jobRequest.Category == -1 {
		return 100
	}

	for _, category := range job.Categories {
		if category == jobRequest.Category {
			return 100
		}
	}

	return 0
}
