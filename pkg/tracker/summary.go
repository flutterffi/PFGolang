package tracker

import "fmt"

func MarkComplete(session *Session) {
	session.Completed = true
}

func Summary(session Session) string {
	status := "in progress"
	if session.Completed {
		status = "complete"
	}

	return fmt.Sprintf("%s: %d minutes (%s)", session.Title, session.MinutesSpent, status)
}
