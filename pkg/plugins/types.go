package plugins

// Descriptor The static part of a plugin configuration (prod).
type Descriptor struct {
	// ModuleName (required)
	ModuleName string `description:"plugin's module name." json:"moduleName,omitempty" toml:"moduleName,omitempty" yaml:"moduleName,omitempty" export:"true"`

	// Version (required)
	Version string `description:"plugin's version." json:"version,omitempty" toml:"version,omitempty" yaml:"version,omitempty" export:"true"`

    // CleanArchives
	CleanArchives bool `description:"Whether to clear the history packet?" json:"cleanArchives,omitempty" toml:"cleanArchives,omitempty" yaml:"cleanArchives,omitempty" export:"true"`

	// OnLine
	Online bool `description:"online" json:"online,omitempty" toml:"online,omitempty" yaml:"online,omitempty" export:"true"`
}

// DevPlugin The static part of a plugin configuration (only for dev).
type DevPlugin struct {
	// GoPath plugin's GOPATH. (required)
	GoPath string `description:"plugin's GOPATH." json:"goPath,omitempty" toml:"goPath,omitempty" yaml:"goPath,omitempty" export:"true"`

	// ModuleName (required)
	ModuleName string `description:"plugin's module name."  json:"moduleName,omitempty" toml:"moduleName,omitempty" yaml:"moduleName,omitempty" export:"true"`
}

// Manifest The plugin manifest.
type Manifest struct {
	DisplayName   string                 `yaml:"displayName"`
	Type          string                 `yaml:"type"`
	Import        string                 `yaml:"import"`
	BasePkg       string                 `yaml:"basePkg"`
	Compatibility string                 `yaml:"compatibility"`
	Summary       string                 `yaml:"summary"`
	TestData      map[string]interface{} `yaml:"testData"`
}
