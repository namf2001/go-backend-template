package generator

// ProjectData holds all configuration for project generation
type ProjectData struct {
	ProjectName string          // Short name, e.g., "my-app"
	ModulePath  string          // Full Go module path, e.g., "github.com/user/my-app"
	DBDriver    string          // "postgres", "mysql", "sqlite", "none"
	Features    map[string]bool // Feature toggles: "jwt", "oauth", "swagger", "prometheus", "docker", "cicd"
	GitInit     bool            // Whether to initialize git
}

// HasFeature checks if a feature is enabled
func (p *ProjectData) HasFeature(name string) bool {
	return p.Features[name]
}

// HasDB returns true if a real database driver is selected
func (p *ProjectData) HasDB() bool {
	return p.DBDriver != "" && p.DBDriver != "none"
}
