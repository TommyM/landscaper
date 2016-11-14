package landscaper

import (
	"reflect"

	"gopkg.in/validator.v2"
)

// Component contains information about the release, configuration and secrets of a component
type Component struct {
	Name          string        `json:"name" validate:"nonzero,max=12"`
	Release       *Release      `json:"release" validate:"nonzero"`
	Configuration Configuration `json:"configuration"`
	Secrets       *Secrets      `json:"secrets"`
}

// NewComponent creates a Component and adds Name to the configuration
func NewComponent(name string, release *Release, cfg Configuration, secrets *Secrets) *Component {
	cmp := &Component{
		Name:          name,
		Release:       release,
		Configuration: cfg,
		Secrets:       secrets,
	}

	cmp.Configuration[metadataKey] = map[string]interface{}{
		releaseVersionKey: cmp.Release.Version,
		landscaperTagKey:  true,
	}

	return cmp
}

// Validate the component on required fields and correct values
func (c *Component) Validate() error {
	return validator.Validate(c)
}

// Equals checks if this component's values are equal to another
func (c *Component) Equals(other *Component) bool {
	return reflect.DeepEqual(c, other)
}
