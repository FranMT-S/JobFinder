package models

type MatchAnalizer struct {
	PorcentSalary     float64
	PorcentSkills     float64
	PorcentModalities float64
	PorcentLevels     float64
	PorcentPosition   float64
	TotalPorcent      float64
	SkillMatches      []string
	SalaryMessage     string
}
