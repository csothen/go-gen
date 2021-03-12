package generator

func (app *Application) NewController(name string) *Controller {
	ctrl := &Controller{name, []*Endpoint{}}

	app.controllers = append(app.controllers, ctrl)

	return ctrl
}

func (app *Application) NewService(name string) *Service {
	service := &Service{name}

	app.services = append(app.services, service)

	return service
}

func (app *Application) NewRepository(name string) *Repository {
	repo := &Repository{name}

	app.repositories = append(app.repositories, repo)

	return repo
}

func (app *Application) NewModel(name string, structure Structure) *Model {
	model := &Model{name, structure}

	app.models = append(app.models, model)

	return model
}
