package models

import "strings"

// Skill is a type that represents the skill of the job
type Skill string

var (
	React           Skill = "react"
	ReactNative     Skill = "react-native"
	Go              Skill = "go"
	Golang          Skill = "golang"
	Python          Skill = "python"
	JavaScript      Skill = "javascript"
	Java            Skill = "java"
	CSharp          Skill = "c#"
	PHP             Skill = "php"
	Ruby            Skill = "ruby"
	Swift           Skill = "swift"
	Kotlin          Skill = "kotlin"
	TypeScript      Skill = "typescript"
	Git             Skill = "git"
	Node            Skill = "node"
	Next            Skill = "next"
	Nextjs          Skill = "nextjs"
	Nestjs          Skill = "nestjs"
	Angular         Skill = "angular"
	Rust            Skill = "rust"
	Flutter         Skill = "flutter"
	Dart            Skill = "dart"
	CPlusPlus       Skill = "c++"
	C               Skill = "c"
	GraphQL         Skill = "graphql"
	Vue             Skill = "vue"
	Vuejs           Skill = "vue.js"
	MachineLearning Skill = "machine learning"
	AI              Skill = "ai"
	ML              Skill = "ml"
	Blockchain      Skill = "blockchain"
	Security        Skill = "security"
	Crypto          Skill = "crypto"
	Web3            Skill = "web3"
	PostgreSQL      Skill = "postgresql"
	Postgres        Skill = "postgres"
	SQL             Skill = "sql"
	Elasticsearch   Skill = "elasticsearch"
	Docker          Skill = "docker"
	Ecommerce       Skill = "ecommerce"
	Salesforce      Skill = "salesforce"
	Shopify         Skill = "shopify"
	WordPress       Skill = "wordpress"
	Laravel         Skill = "laravel"
	Payroll         Skill = "payroll"
	SEO             Skill = "seo"
	Apache          Skill = "apache"
	Mongo           Skill = "mongo"
	Serverless      Skill = "serverless"
	CSS             Skill = "css"
	HTML            Skill = "html"
	ObjectiveC      Skill = "objective-c"
	Scala           Skill = "scala"
	Web             Skill = "web"
	WebDev          Skill = "web-dev"
	NoSQL           Skill = "nosql"
	Android         Skill = "android"
	API             Skill = "api"
	SAAS            Skill = "saas"
)

func NewSkill(name string) Skill {
	return Skill(strings.TrimSpace(strings.ToLower(name)))
}

type SkillCategory struct {
	name     Skill
	category Category
}

func NewSkillCategory(name Skill, category Category) SkillCategory {
	return SkillCategory{
		name:     name,
		category: category,
	}
}

// SkillCategoryMap is a map of skills and their categories
// used to normalize the skills names, work with remote ok
var SkillCategoryMap = map[Skill]SkillCategory{
	JavaScript:      NewSkillCategory("javascript", FrontEnd),
	Golang:          NewSkillCategory("golang", Backend),
	React:           NewSkillCategory("react", FrontEnd),
	SAAS:            NewSkillCategory("saas", Backend),
	API:             NewSkillCategory("api", Backend),
	Ruby:            NewSkillCategory("ruby", Backend),
	Python:          NewSkillCategory("python", Backend),
	Node:            NewSkillCategory("node", Backend),
	Ecommerce:       NewSkillCategory("ecommerce", Backend),
	Java:            NewSkillCategory("java", Backend),
	Crypto:          NewSkillCategory("crypto", Backend),
	Git:             NewSkillCategory("git", Backend),
	Android:         NewSkillCategory("android", Backend),
	PHP:             NewSkillCategory("php", Backend),
	Serverless:      NewSkillCategory("serverless", Backend),
	CSS:             NewSkillCategory("css", FrontEnd),
	Angular:         NewSkillCategory("angular", FrontEnd),
	HTML:            NewSkillCategory("html", FrontEnd),
	Salesforce:      NewSkillCategory("salesforce", Backend),
	SQL:             NewSkillCategory("sql", Database),
	C:               NewSkillCategory("c", Backend),
	WebDev:          NewSkillCategory("web-dev", Backend),
	NoSQL:           NewSkillCategory("nosql", Backend),
	Postgres:        NewSkillCategory("postgres", Database),
	CPlusPlus:       NewSkillCategory("c-plus-plus", Backend),
	CSharp:          NewSkillCategory("c-sharp", Backend),
	SEO:             NewSkillCategory("seo", Backend),
	Apache:          NewSkillCategory("apache", Backend),
	ReactNative:     NewSkillCategory("react-native", Mobile),
	Mongo:           NewSkillCategory("mongo", Database),
	Shopify:         NewSkillCategory("shopify", Backend),
	WordPress:       NewSkillCategory("wordpress", Backend),
	Laravel:         NewSkillCategory("laravel", Backend),
	Elasticsearch:   NewSkillCategory("elasticsearch", Backend),
	Blockchain:      NewSkillCategory("blockchain", Backend),
	Web3:            NewSkillCategory("web3", Backend),
	Docker:          NewSkillCategory("docker", Backend),
	GraphQL:         NewSkillCategory("graphql", Backend),
	Payroll:         NewSkillCategory("payroll", Backend),
	MachineLearning: NewSkillCategory("machine-learning", DataScience),
	Scala:           NewSkillCategory("scala", Backend),
	Web:             NewSkillCategory("web", Backend),
	ObjectiveC:      NewSkillCategory("objective-c", Backend),
	Vue:             NewSkillCategory("vue", FrontEnd),
}

func GetSkill(skill Skill) Skill {
	if _, ok := SkillCategoryMap[skill]; !ok {
		return ""
	}

	return SkillCategoryMap[skill].name
}

func GetSkillCategory(skill Skill) Category {
	if _, ok := SkillCategoryMap[skill]; !ok {
		return NotCategory
	}

	return SkillCategoryMap[skill].category
}

func GetAllSkills() []Skill {
	return []Skill{
		React,
		ReactNative,
		Go,
		Golang,
		Python,
		JavaScript,
		Java,
		CSharp,
		PHP,
		Ruby,
		Swift,
		Kotlin,
		TypeScript,
		Node,
		Next,
		Nestjs,
		Nextjs,
		Angular,
		Rust,
		Flutter,
		Dart,
		CPlusPlus,
		C,
		GraphQL,
		Vue,
		Vuejs,
		MachineLearning,
		AI,
		ML,
		Blockchain,
		Security,
		Crypto,
		Web3,
		PostgreSQL,
		Postgres,
		SQL,
		Elasticsearch,
		Docker,
	}
}
