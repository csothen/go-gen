package generator

import "os"

func NewProject(projectRoot string) (*Project, error) {
	err := os.MkdirAll(projectRoot, 0755)
	if err != nil {
		return nil, err
	}

	return &Project{projectRoot, []*Application{}}, nil
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
	return nil
}

// Create : Creates the project following the structure established
func (p *Project) Create() error {
	return nil
}
