package chart

// Chart describes a Chart's quality
type Chart struct {
	Name   string
	Status int `json:"status"`
}

// IsComplete allows for a quick check to see if a Chart is valid
func (chart *Chart) IsComplete() bool {
	if chart.Name == "" {
		return false
	}
	return true
}
