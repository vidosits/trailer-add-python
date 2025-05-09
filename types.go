package recommendation

type VersionSpecifier string

const (
	UnconstrainedVersion VersionSpecifier = "unconstrained"
)

type Version struct {
	Specifier  VersionSpecifier `json:"specifier,omitempty"`
	Constraint string           `json:"constraint"`
}

type Package struct {
	Name        string  `json:"name"`
	Version     Version `json:"version,omitempty"`
	Description string  `json:"description,omitempty"`
	Channel     string  `json:"channel"`
}

type Environment struct {
	Packages []Package `json:"packages"`
}

type ImageConfiguration struct {
	Environments map[string]Environment `json:"environments"`
}
