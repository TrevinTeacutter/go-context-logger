package logs

import (
	"regexp"

	"go.opentelemetry.io/otel/attribute"
)

type Matcher interface {
	Matches(log Log) bool
}

var _ Matcher = (*AttributeMatcher)(nil)

type AttributeMatcher struct {
	Key      attribute.Key
	Operator string
	Value    attribute.Value
}

func (matcher *AttributeMatcher) Matches(log Log) bool {
	switch matcher.Operator {
	case "EQUALS":
		for _, kvPair := range log.Attributes {
			if kvPair.Key == matcher.Key &&
				kvPair.Value == matcher.Value {
				return true
			}
		}
	case "NOT_EQUALS":
		for _, kvPair := range log.Attributes {
			if kvPair.Key == matcher.Key &&
				kvPair.Value != matcher.Value {
				return true
			}
		}
	}

	return false
}

type AttributeRegexMatcher struct {
	Key      attribute.Key
	Operator string
	Value    regexp.Regexp
}

func (matcher *AttributeRegexMatcher) Matches(log Log) bool {
	switch matcher.Operator {
	case "MATCHES":
		for _, kvPair := range log.Attributes {
			if kvPair.Key == matcher.Key &&
				matcher.Value.MatchString(kvPair.Value.Emit()) {
				return true
			}
		}
	case "NOT_MATCHES":
		for _, kvPair := range log.Attributes {
			if kvPair.Key == matcher.Key &&
				!matcher.Value.MatchString(kvPair.Value.Emit()) {
				return true
			}
		}
	}

	return false
}
