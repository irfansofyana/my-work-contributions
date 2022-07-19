package model

// Work struct defines a work that a person does
type Work struct {
	What        string
	RelatedLink string
	Type        TypeOfWork
}

// TypeOfWork is custom type to hold value of type of work
type TypeOfWork int

const (
	GithubPR TypeOfWork = iota
	GithubIssue
	JiraTicket
)

// String is a function to create common behavior - give the type a String function
func (t TypeOfWork) String() string {
	return [...]string{
		"Github Pull Request",
		"Github Issue",
		"JIRA Ticket",
	}[t]
}

// NewWork is function to create new instance of Work
func NewWork(what, relatedLink string, typeOfWork TypeOfWork) Work {
	return Work{what, relatedLink, typeOfWork}
}
