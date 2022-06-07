package logs

type Route struct {
	Matchers []Matcher
	Children []Route
	Receiver Receiver
	Continue bool
}

func (r *Route) Route(log Log) []Receiver {
	for _, matcher := range r.Matchers {
		if !matcher.Matches(log) {
			return nil
		}
	}

	receivers := []Receiver{r.Receiver}

	for _, child := range r.Children {
		receivers = append(receivers, child.Route(log)...)
	}

	return receivers
}
