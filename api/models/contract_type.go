package models

type ContractType string

var (
	FullTime           ContractType = "full-time"
	PartTime           ContractType = "part-time"
	Contract           ContractType = "contract"
	Internship         ContractType = "internship"
	ContractNoSpecific ContractType = "no-specific"
)
