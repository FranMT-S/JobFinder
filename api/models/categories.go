package models

type Category int

const (
	NotCategory Category = -1
	Engineer    Category = iota
	Executive
	Developer
	Finance
	SysAdmin
	Backend
	FrontEnd
	FullStack
	Cloud
	Mobile
	DevOps
	DataScience
	Testing
	Architect
	Analyst
	Database
	Design
)

type CategoryModel struct {
	ID   Category `json:"id"`
	Name string   `json:"name"`
}

var CategoryMapRemoteOk = map[Category]string{
	Engineer:    "engineer",
	Executive:   "exec",
	Developer:   "dev",
	Finance:     "finance",
	SysAdmin:    "sys-admin",
	Backend:     "backend",
	FrontEnd:    "front-end",
	FullStack:   "full-stack",
	Cloud:       "cloud",
	Mobile:      "mobile",
	DevOps:      "devops",
	DataScience: "data-science",
	Testing:     "testing",
	Architect:   "architect",
	Analyst:     "analyst",
	Database:    "database",
	Design:      "design",
	NotCategory: "any",
}

var CategoryMapWeWorkRemotely = map[Category]int{
	Design:    1,
	FullStack: 2,
	SysAdmin:  6,
	DevOps:    6,
	FrontEnd:  17,
	Backend:   18,
}

func GetCategoryAllCategories() []CategoryModel {
	categories := []CategoryModel{
		{ID: NotCategory, Name: "NotCategory"},
		{ID: Engineer, Name: "Engineer"},
		{ID: Executive, Name: "Executive"},
		{ID: Developer, Name: "Developer"},
		{ID: Finance, Name: "Finance"},
		{ID: SysAdmin, Name: "SysAdmin"},
		{ID: Backend, Name: "Backend"},
		{ID: FrontEnd, Name: "FrontEnd"},
		{ID: FullStack, Name: "FullStack"},
		{ID: Cloud, Name: "Cloud"},
		{ID: Mobile, Name: "Mobile"},
		{ID: DevOps, Name: "DevOps"},
		{ID: DataScience, Name: "DataScience"},
		{ID: Testing, Name: "Testing"},
		{ID: Architect, Name: "Architect"},
		{ID: Analyst, Name: "Analyst"},
		{ID: Database, Name: "Database"},
		{ID: Design, Name: "Design"},
	}

	return categories
}
