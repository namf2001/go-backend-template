package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/namf2001/go-backend-template/cmd/cli/generator"
	"github.com/namf2001/go-backend-template/cmd/cli/tui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Project name")
	createCmd.Flags().StringP("module", "m", "", "Go module path (e.g., github.com/user/project)")
	createCmd.Flags().StringP("db", "d", "", "Database driver: postgres, mysql, sqlite, none")
	createCmd.Flags().StringSliceP("feature", "f", nil, "Features to enable: jwt, oauth, swagger, prometheus, docker, cicd")
	createCmd.Flags().BoolP("git", "g", false, "Initialize git repository")
	createCmd.Flags().BoolP("all-features", "a", false, "Enable all features")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go backend project",
	Long:  "Scaffold a new production-ready Go backend project with clean architecture, database setup, authentication, and more.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tui.LogoStyle.Render(tui.Logo))
		fmt.Println(tui.SubtitleStyle.Render("Scaffold a production-ready Go backend project\n"))

		project := &generator.ProjectData{
			Features: make(map[string]bool),
		}

		// --- Project Name ---
		flagName, _ := cmd.Flags().GetString("name")
		if flagName != "" {
			project.ProjectName = flagName
		} else {
			var result string
			p := tea.NewProgram(tui.NewTextInput(
				"What is the name of your project?",
				"my-awesome-backend",
				&result,
			))
			if _, err := p.Run(); err != nil {
				log.Fatalf("Error: %v", err)
			}
			if result == "" {
				fmt.Println(tui.ErrorStyle.Render("Project name is required"))
				os.Exit(1)
			}
			project.ProjectName = result
		}

		// Validate: directory doesn't already exist
		if _, err := os.Stat(project.ProjectName); err == nil {
			fmt.Println(tui.ErrorStyle.Render(fmt.Sprintf("Directory '%s' already exists", project.ProjectName)))
			os.Exit(1)
		}

		// --- Module Path ---
		flagModule, _ := cmd.Flags().GetString("module")
		if flagModule != "" {
			project.ModulePath = flagModule
		} else {
			var result string
			defaultModule := fmt.Sprintf("github.com/%s/%s", getGitUsername(), project.ProjectName)
			p := tea.NewProgram(tui.NewTextInput(
				"What is the Go module path?",
				defaultModule,
				&result,
			))
			if _, err := p.Run(); err != nil {
				log.Fatalf("Error: %v", err)
			}
			if result == "" {
				result = defaultModule
			}
			project.ModulePath = result
		}

		// --- Database Driver ---
		flagDB, _ := cmd.Flags().GetString("db")
		if flagDB != "" {
			project.DBDriver = flagDB
		} else {
			var result string
			dbDescs := map[string]string{
				"postgres": "PostgreSQL - Production recommended",
				"mysql":    "MySQL / MariaDB",
				"sqlite":   "SQLite - Lightweight, file-based",
				"none":     "No database",
			}
			p := tea.NewProgram(tui.NewSelect(
				"Choose a database driver:",
				tui.DBDrivers,
				dbDescs,
				&result,
			))
			if _, err := p.Run(); err != nil {
				log.Fatalf("Error: %v", err)
			}
			if result == "" {
				result = "postgres"
			}
			project.DBDriver = result
		}

		fmt.Printf("  %s Database: %s\n", tui.CheckMark, tui.HighlightStyle.Render(project.DBDriver))

		// --- Features ---
		allFeatures, _ := cmd.Flags().GetBool("all-features")
		flagFeatures, _ := cmd.Flags().GetStringSlice("feature")
		featureFlagSet := cmd.Flags().Changed("feature") || cmd.Flags().Changed("all-features")

		if allFeatures {
			for _, f := range tui.FeatureList {
				project.Features[f] = true
			}
		} else if featureFlagSet {
			for _, f := range flagFeatures {
				f = strings.TrimSpace(strings.ToLower(f))
				if f != "" {
					project.Features[f] = true
				}
			}
		} else {
			// Interactive multi-select
			result := make(map[string]bool)
			// Pre-select recommended features
			result["jwt"] = true
			result["swagger"] = true
			result["docker"] = true

			p := tea.NewProgram(tui.NewMultiSelect(
				"Select features to include:",
				tui.FeatureList,
				tui.FeatureDescriptions,
				result,
			))
			if _, err := p.Run(); err != nil {
				log.Fatalf("Error: %v", err)
			}
			project.Features = result
		}

		// Print selected features
		for _, f := range tui.FeatureList {
			if project.Features[f] {
				fmt.Printf("  %s %s\n", tui.CheckMark, tui.FeatureDescriptions[f])
			}
		}
		fmt.Println()

		// --- Git ---
		flagGit, _ := cmd.Flags().GetBool("git")
		project.GitInit = flagGit

		// --- Generate ---
		fmt.Println(tui.SubtitleStyle.Render("Generating project...\n"))

		// Use spinner during generation
		spinnerModel := tui.NewSpinner("Creating your project...")
		p := tea.NewProgram(spinnerModel)

		wg := sync.WaitGroup{}
		wg.Add(1)
		var genErr error

		go func() {
			defer wg.Done()
			genErr = generator.Generate(project)
			p.Send(tui.DoneMsg{Err: genErr})
		}()

		if _, err := p.Run(); err != nil {
			log.Fatalf("Spinner error: %v", err)
		}
		wg.Wait()

		if genErr != nil {
			fmt.Println(tui.ErrorStyle.Render(fmt.Sprintf("\n✗ Failed to generate project: %v", genErr)))
			os.Exit(1)
		}

		// --- Success ---
		fmt.Println(tui.SuccessStyle.Render(fmt.Sprintf("\n✓ Project '%s' created successfully!\n", project.ProjectName)))
		fmt.Println(tui.TitleStyle.Render("Next steps:\n"))
		fmt.Printf("  %s cd %s\n", tui.BulletStyle, project.ProjectName)
		fmt.Printf("  %s cp .env.example .env.dev\n", tui.BulletStyle)

		if project.HasFeature("docker") {
			fmt.Printf("  %s docker-compose up -d\n", tui.BulletStyle)
		}

		if project.HasDB() {
			fmt.Printf("  %s make migrate-up\n", tui.BulletStyle)
		}

		fmt.Printf("  %s make dev\n", tui.BulletStyle)
		fmt.Println()

		// Print non-interactive equivalent
		nonInteractive := buildNonInteractiveCmd(project)
		fmt.Println(tui.TipStyle.Render("Tip: Repeat with this non-interactive command:"))
		fmt.Println(tui.DimStyle.Render(fmt.Sprintf("  %s\n", nonInteractive)))
	},
}

func getGitUsername() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "username"
	}
	// Try to read git config to get username
	configPath := filepath.Join(home, ".gitconfig")
	content, err := os.ReadFile(configPath)
	if err != nil {
		return "username"
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "name") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[1])
				name = strings.ReplaceAll(name, " ", "-")
				return strings.ToLower(name)
			}
		}
	}
	return "username"
}

func buildNonInteractiveCmd(data *generator.ProjectData) string {
	cmd := fmt.Sprintf("go-backend create --name %s --module %s --db %s", data.ProjectName, data.ModulePath, data.DBDriver)

	for _, f := range tui.FeatureList {
		if data.Features[f] {
			cmd += fmt.Sprintf(" --feature %s", f)
		}
	}

	if data.GitInit {
		cmd += " --git"
	}

	return cmd
}
