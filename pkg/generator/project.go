package generator

import (
	"fmt"
	"os"
	"strings"
)

func NewProject(name, projectRoot string) (*Project, error) {
	return &Project{name, projectRoot + name, []*Application{}}, nil
}

// Build : Builds the project following the structure established
func (p *Project) Build() error {
	if err := os.MkdirAll(p.root, 0755); err != nil {
		return err
	}

	fmt.Println("Building project applications...")

	for _, app := range p.applications {
		if err := app.build(); err != nil {
			return err
		}
	}

	fmt.Println("Done!")

	return nil
}

func (p *Project) NewApplication(name string) *Application {
	app := &Application{
		name,
		[]*Controller{},
		[]*Service{},
		[]*Repository{},
		[]*Model{},
	}

	p.applications = append(p.applications, app)

	return app
}

// Describe : Returns the structure of the project
func (p *Project) Describe() error {
	fmt.Println(p.name)

	if len(p.applications) > 0 {
		depth := buildTreeUntilPath(GeneratorConfig.APPLICATION_LOCATION)
		for _, app := range p.applications {
			fmt.Printf("%s-%s\n", strings.Repeat("  |", depth), app.name)
			fmt.Printf("%s-main.go\n", strings.Repeat("  |", depth+1))
		}

		depth = buildTreeUntilPath(GeneratorConfig.CONTROLLER_LOCATION)
		for _, app := range p.applications {
			if len(app.controllers) > 0 {
				fmt.Printf("%s-%s\n", strings.Repeat("  |", depth), app.name)
				fmt.Printf("%s-controllers\n", strings.Repeat("  |", depth+1))
				for _, ctrl := range app.controllers {
					fmt.Printf("%s-%s.go\n", strings.Repeat("  |", depth+2), ctrl.name)
				}
			}
		}

		depth = buildTreeUntilPath(GeneratorConfig.SERVICE_LOCATION)
		for _, app := range p.applications {
			if len(app.services) > 0 {
				fmt.Printf("%s-%s\n", strings.Repeat("  |", depth), app.name)
				fmt.Printf("%s-services\n", strings.Repeat("  |", depth+1))
				for _, serv := range app.services {
					fmt.Printf("%s-%s.go\n", strings.Repeat("  |", depth+2), serv.name)
				}
			}
		}

		depth = buildTreeUntilPath(GeneratorConfig.REPOSITORY_LOCATION)
		for _, app := range p.applications {
			if len(app.repositories) > 0 {
				fmt.Printf("%s-%s\n", strings.Repeat("  |", depth), app.name)
				fmt.Printf("%s-repositories\n", strings.Repeat("  |", depth+1))
				for _, repo := range app.repositories {
					fmt.Printf("%s-%s.go\n", strings.Repeat("  |", depth+2), repo.name)
				}
			}
		}

		depth = buildTreeUntilPath(GeneratorConfig.MODEL_LOCATION)
		for _, app := range p.applications {
			if len(app.models) > 0 {
				fmt.Printf("%s-%s\n", strings.Repeat("  |", depth), app.name)
				fmt.Printf("%s-models\n", strings.Repeat("  |", depth+1))
				for _, model := range app.models {
					fmt.Printf("%s-%s.go\n", strings.Repeat("  |", depth+2), model.name)
				}
			}
		}
	}

	return nil
}

func buildTreeUntilPath(path string) int {
	folders := strings.Split(path, "/")
	depth := len(folders)

	for i, folder := range folders {
		if len(folder) > 0 {
			fmt.Printf("%s-%s\n", strings.Repeat("  |", i+1), folder)
		}
	}

	return depth
}
