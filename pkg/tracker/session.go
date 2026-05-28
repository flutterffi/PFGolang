package tracker

type Session struct {
	Title        string
	MinutesSpent int
	Completed    bool
}

func NewSession(title string, minutes int) Session {
	return Session{
		Title:        title,
		MinutesSpent: minutes,
		Completed:    false,
	}
}
