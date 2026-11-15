
package incident

import (
    "errors"
    "time"
)

// Severity represents the seriousness of an incident.
type Severity string

const (
	SeverityLow      Severity = "LOW"
	SeverityMedium   Severity = "MEDIUM"
	SeverityHigh     Severity = "HIGH"
	SeverityCritical Severity = "CRITICAL"
)

// Incident represents an infrastructure failure.
type Incident struct {
	ID        string
	Resource  string
	Namespace string
	Reason    string
	Message   string
	Severity  Severity
	DetectedAt time.Time
}
// Validate checks whether an incident contains valid data.
func (i Incident) Validate() error {
	if i.ID == "" {
		return errors.New("incident ID is required")
	}

	if i.Resource == "" {
		return errors.New("resource is required")
	}

	if i.Reason == "" {
		return errors.New("incident reason is required")
	}

	switch i.Severity {
	case SeverityLow, SeverityMedium, SeverityHigh, SeverityCritical:
	default:
		return errors.New("invalid incident severity")
	}

	if i.DetectedAt.IsZero() {
		return errors.New("detection time is required")
	}

	return nil
}
