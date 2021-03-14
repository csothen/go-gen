package generator

import "fmt"

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

func (app *Application) build() error {

	fmt.Println(fmt.Sprintf("Building application %s controllers...", app.name))

	for _, ctrl := range app.controllers {
		if err := ctrl.build(); err != nil {
			return err
		}
	}

	fmt.Println(fmt.Sprintf("Building application %s services...", app.name))

	for _, serv := range app.services {
		if err := serv.build(); err != nil {
			return err
		}
	}

	fmt.Println(fmt.Sprintf("Building application %s repositories...", app.name))

	for _, repo := range app.repositories {
		if err := repo.build(); err != nil {
			return err
		}
	}

	fmt.Println(fmt.Sprintf("Building application %s models...", app.name))

	for _, model := range app.models {
		if err := model.build(); err != nil {
			return err
		}
	}

	// Build actual code for application

	return nil
}
