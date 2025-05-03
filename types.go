package recommendation

type VersionSpecifier string

const (
	ExactVersion         VersionSpecifier = "exactly"
	MinimumVersion       VersionSpecifier = "minimum"
	MaximumVersion       VersionSpecifier = "maximum"
	ExcludeVersion       VersionSpecifier = "exclude"
	CompatibleVersion    VersionSpecifier = "compatible"
	UnconstrainedVersion VersionSpecifier = "unconstrained"
	CustomVersion        VersionSpecifier = "custom"
)

type Version struct {
	Specifier  VersionSpecifier `json:"specifier,omitempty"`
	Constraint string           `json:"constraint"`
}

type Package[V Version | string] struct {
	Name        string `json:"name"`
	Version     V      `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
	Channel     string `json:"channel"`
}

type Environment struct {
	Channels []string           `json:"channels"`
	Packages []Package[Version] `json:"packages"`
}

type startupConfiguration struct {
	Environment string   `json:"environment"`
	Command     []string `json:"command"`
}

type startup struct {
	DefaultEnvironment    string                 `json:"defaultEnvironment"`
	StartupConfigurations []startupConfiguration `json:"startupConfigurations"`
}
type UserConfiguration struct {
	Name  string `json:"name"`
	Group string `json:"group"`
	UID   int64  `json:"uid"`
	GID   int64  `json:"gid"`
}

type File struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	InternalFileName string `json:"internalFileName"`
	URL              string `json:"url"`
}

type ImageConfiguration struct {
	Environments map[string]Environment `json:"environments"`
	FilePaths    map[string]string      `json:"filePaths"`
	User         *UserConfiguration     `json:"user"`
	BaseImage    string                 `json:"baseImage"`
	Startup      startup                `json:"startup"`
}

type Image struct {
	Dockerfile     string             `json:"dockerfile"`
	Status         string             `json:"status"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	ID             string             `json:"id"`
	FileToken      string             `json:"fileToken"`
	Repository     string             `json:"repository"`
	Digest         string             `json:"digest"`
	Tag            string             `json:"tag"`
	Configuration  ImageConfiguration `json:"configuration"`
	Files          []File             `json:"files"`
	PushAfterBuild bool               `json:"pushAfterBuild"`
	External       bool               `json:"external"`
	ForcePullBase  bool               `json:"forcePullBase"`
}
