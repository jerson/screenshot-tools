package utils

import (
	"path/filepath"
	"strings"
)

// EvidenceSuggestion ...
type EvidenceSuggestion struct {
	Model       string
	Name        string
	Description string
}

// GetEvidenceSuggestion ...
func GetEvidenceSuggestion(ref string) (EvidenceSuggestion, error) {

	suggestion := EvidenceSuggestion{}
	path, err := filepath.Abs(ref)
	if err != nil {
		return suggestion, err
	}
	currentDir := filepath.Base(path)
	currentDirNames := strings.Split(currentDir, "_")
	if len(currentDirNames) > 1 {
		suggestion.Model = currentDirNames[0]
		suggestion.Name = strings.Join(currentDirNames[1:], "_")
	} else {
		suggestion.Name = currentDir
	}
	return suggestion, nil
}
