package tui

import "github.com/charmbracelet/lipgloss"

const Logo = `
  ____         ____             _                  _ 
 / ___| ___   | __ )  __ _  ___| | _____ _ __   __| |
| |  _ / _ \  |  _ \ / _` + "`" + ` |/ __| |/ / _ \ '_ \ / _` + "`" + ` |
| |_| | (_) | | |_) | (_| | (__|   <  __/ | | | (_| |
 \____|\___/  |____/ \__,_|\___|_|\_\___|_| |_|\__,_|

`

var (
	LogoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	TitleStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Bold(true).PaddingLeft(1)
	SubtitleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).PaddingLeft(1)
	SuccessStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).PaddingLeft(1)
	ErrorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4672")).PaddingLeft(1)
	TipStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Italic(true).PaddingLeft(1)
	DimStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).PaddingLeft(1)
	HighlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	CheckMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).SetString("✓")
	BulletStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).SetString("•")
)

// DBDrivers is the list of supported database drivers
var DBDrivers = []string{"postgres", "mysql", "sqlite", "none"}

// FeatureList is the list of supported advanced features
var FeatureList = []string{"jwt", "oauth", "swagger", "prometheus", "docker", "cicd"}

// FeatureDescriptions describes each feature
var FeatureDescriptions = map[string]string{
	"jwt":        "JWT Authentication",
	"oauth":      "Google OAuth2",
	"swagger":    "Swagger API Docs",
	"prometheus": "Prometheus Metrics",
	"docker":     "Docker + Docker Compose",
	"cicd":       "GitHub Actions CI/CD",
}
