package models

type MatchAnalizer struct {
	PorcentSalary     float64  `json:"porcentSalary"`
	PorcentSkills     float64  `json:"porcentSkills"`
	PorcentModalities float64  `json:"porcentModalities"`
	PorcentLevels     float64  `json:"porcentLevels"`
	PorcentPosition   float64  `json:"porcentPosition"`
	TotalPorcent      float64  `json:"totalPorcent"`
	SkillMatches      []string `json:"skillMatches"`
	SalaryMessage     string   `json:"salaryMessage"`
}
