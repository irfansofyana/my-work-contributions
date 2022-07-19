package model

import (
	"reflect"
	"testing"
)

func TestTypeOfWorkString(t *testing.T) {
	testCases := []struct {
		input TypeOfWork
		want  string
	}{
		{input: GithubPR, want: "Github Pull Request"},
		{input: GithubIssue, want: "Github Issue"},
		{input: JiraTicket, want: "JIRA Ticket"},
	}

	for i, tc := range testCases {
		got := tc.input.String()
		if got != tc.want {
			t.Errorf("TestTypeOfWorkString: test %d failed: expected: %v, got: %v", i, tc.want, got)
		}
	}
}

func TestNewWork(t *testing.T) {
	testCases := []struct {
		what        string
		relatedLink string
		typeWork    TypeOfWork
		want        Work
	}{
		{
			what:        "feat: implement authentication",
			relatedLink: "https://github.com/someproject/pull/1",
			typeWork:    GithubPR,
			want: Work{
				"feat: implement authentication",
				"https://github.com/someproject/pull/1",
				GithubPR,
			},
		},
	}

	for i, tc := range testCases {
		got := NewWork(tc.what, tc.relatedLink, tc.typeWork)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TestNewWork: test %d failed: expected: %v, got: %v", i, tc.want, got)
		}
	}
}
