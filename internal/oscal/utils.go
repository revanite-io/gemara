package oscalexporter

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	oscalValidation "github.com/defenseunicorns/go-oscal/src/pkg/validation"
	oscal "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

const (
	OSCALVersion    = "1.1.3"
	GemaraNamespace = "https://github.com/ossf/gemara/ns/oscal"
)

func NilIfEmpty[T any](slice *[]T) *[]T {
	if slice == nil || len(*slice) == 0 {
		return nil
	}
	return slice
}

func NormalizeControl(input string) string {
	re := regexp.MustCompile(`\((\d+)\)`)
	replacedString := re.ReplaceAllString(input, ".$1")
	finalString := strings.ToLower(replacedString)
	return finalString
}

func GetTimeWithFallback(timeStr string, fallback time.Time) time.Time {
	if parsedTime := GetTime(timeStr); parsedTime != nil {
		return *parsedTime
	}
	return fallback
}

func GetTime(timeStr string) *time.Time {
	if timeStr == "" {
		return nil
	}
	if parsedTime, err := time.Parse(time.RFC3339, timeStr); err == nil {
		return &parsedTime
	}
	return nil
}

func Validate(oscalModels oscal.OscalModels) error {
	validator, err := oscalValidation.NewValidatorDesiredVersion(oscalModels, OSCALVersion)
	if err != nil {
		return fmt.Errorf("failed to create validator: %w", err)
	}
	if err := validator.Validate(); err != nil {
		return fmt.Errorf("model failed validation: %w", err)
	}
	return nil
}
