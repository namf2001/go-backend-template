package generator

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const originalModule = "github.com/namf2001/go-backend-template"

// fileRenames maps embedded template names (without dot prefix) to their real output names
var fileRenames = map[string]string{
	"air.toml":    ".air.toml",
	"env.example": ".env.example",
	"gitignore":   ".gitignore",
	"go_mod":      "go.mod",
	"go_sum":      "go.sum",
}

// dirRenames maps embedded directory names to their real output names
var dirRenames = map[string]string{
	"github_workflows": ".github",
}

// Generate creates a new project from templates
func Generate(data *ProjectData) error {
	outputDir := data.ProjectName

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Walk through embedded templates
	if err := fs.WalkDir(TemplateFS, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from "templates/" prefix
		relPath := strings.TrimPrefix(path, "templates/")
		if relPath == "" {
			return nil
		}

		// Skip files based on feature flags
		if shouldSkip(relPath, data) {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		// Apply renaming for dotfiles/directories
		outputRelPath := applyRenames(relPath)
		targetPath := filepath.Join(outputDir, outputRelPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		// Read template file
		content, err := TemplateFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", path, err)
		}

		// Process .tmpl files through Go template engine
		if strings.HasSuffix(path, ".tmpl") {
			targetPath = strings.TrimSuffix(targetPath, ".tmpl")
			processed, err := processTemplate(relPath, string(content), data)
			if err != nil {
				return fmt.Errorf("failed to process template %s: %w", path, err)
			}
			content = []byte(processed)
		}

		// Replace module name in Go files and go.mod/go.sum
		if isGoFile(targetPath) || isModFile(targetPath) {
			content = []byte(strings.ReplaceAll(string(content), originalModule, data.ModulePath))
		}

		// Ensure parent directory exists
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory for %s: %w", targetPath, err)
		}

		return os.WriteFile(targetPath, content, 0644)
	}); err != nil {
		return fmt.Errorf("failed to generate project files: %w", err)
	}

	// Git init if requested
	if data.GitInit {
		if err := initGit(outputDir); err != nil {
			return fmt.Errorf("git init failed: %w", err)
		}
	}

	return nil
}

// applyRenames transforms embedded template paths back to real output paths
func applyRenames(relPath string) string {
	// Check directory renames
	for old, newName := range dirRenames {
		if strings.HasPrefix(relPath, old+"/") {
			relPath = newName + relPath[len(old):]
		} else if relPath == old {
			relPath = newName
		}
	}

	// Check file renames (only for the basename)
	dir := filepath.Dir(relPath)
	base := filepath.Base(relPath)
	if newName, ok := fileRenames[base]; ok {
		if dir == "." {
			return newName
		}
		return filepath.Join(dir, newName)
	}

	return relPath
}

// shouldSkip determines if a file/dir should be skipped based on features
func shouldSkip(relPath string, data *ProjectData) bool {
	// Skip OAuth-related files if OAuth not enabled
	if !data.HasFeature("oauth") {
		if strings.Contains(relPath, "pkg/oauth") ||
			strings.Contains(relPath, "handler/rest/v1/auth") ||
			strings.Contains(relPath, "controller/auth") {
			return true
		}
	}

	// Skip JWT-related files if JWT not enabled
	if !data.HasFeature("jwt") {
		if strings.Contains(relPath, "pkg/jwt") ||
			strings.Contains(relPath, "handler/middleware/auth.go") {
			return true
		}
	}

	// Skip Swagger files if Swagger not enabled
	if !data.HasFeature("swagger") {
		if strings.Contains(relPath, "docs/swagger") {
			return true
		}
	}

	// Skip Docker files if Docker not enabled
	if !data.HasFeature("docker") {
		if relPath == "Dockerfile" || strings.HasPrefix(relPath, "docker-compose") {
			return true
		}
	}

	// Skip CI/CD files if not enabled (renamed from .github to github_workflows)
	if !data.HasFeature("cicd") {
		if strings.HasPrefix(relPath, "github_workflows") {
			return true
		}
	}

	// Skip database-related files if no DB selected
	if !data.HasDB() {
		if strings.Contains(relPath, "migrations") ||
			strings.Contains(relPath, "repository") ||
			strings.Contains(relPath, "pkg/database") ||
			strings.Contains(relPath, "pkg/testdb") {
			return true
		}
	}

	return false
}

// processTemplate processes a Go template string with ProjectData
func processTemplate(name, content string, data *ProjectData) (string, error) {
	funcMap := template.FuncMap{
		"hasFeature": data.HasFeature,
		"hasDB":      data.HasDB,
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"title":      strings.Title, //nolint:staticcheck
	}

	tmpl, err := template.New(name).Funcs(funcMap).Parse(content)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func isGoFile(path string) bool {
	return strings.HasSuffix(path, ".go")
}

func isModFile(path string) bool {
	return strings.HasSuffix(path, "go.mod") || strings.HasSuffix(path, "go.sum")
}

func runCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func initGit(dir string) error {
	if err := runCommand(dir, "git", "init"); err != nil {
		return err
	}
	if err := runCommand(dir, "git", "add", "."); err != nil {
		return err
	}
	return runCommand(dir, "git", "commit", "-m", "Initial commit from go-backend scaffold")
}
