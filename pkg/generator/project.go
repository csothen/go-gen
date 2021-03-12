package generator

func NewProject(projectRoot string) (*Project, error) {
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
