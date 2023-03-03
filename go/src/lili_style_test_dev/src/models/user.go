package models

type User struct {
	TimeStamp       string
	Mail      		string
	CorporateID     string
	Tel       		string
	Name      		string
	Class    		string

	University      string //大学(学生)
	UnderGraduate   string //学部(学生)
	Department      string //学科(学生)
	GraduationYear  string //卒業年度(学生)
	CurrentEmployer string //現在の勤め先(選考希望者)
	Answers
}

type UserSearchByCorporate struct {
	TimeStamp       string
	Mail      		string
	CorporateID     string
	Tel       		string
	Name      		string
	Class    		string

	University      string //大学(学生)
	UnderGraduate   string //学部(学生)
	Department      string //学科(学生)
	GraduationYear  string //卒業年度(学生)
	CurrentEmployer string //現在の勤め先(選考希望者)
	AltruismRate       int
	BusinessStanceRate int
	CarrierUpRate      int
	CommitRate         int
	MentalityRate      int
	Answers
	Corporate
}

type Answers struct {
	Altruism
	BusinessStance
	CarrierUp
	Commit
	Mentality
	Personality
}




