package generator

type Project struct {
	name         string
	root         string
	applications []*Application
}

type Application struct {
	name         string
	controllers  []*Controller
	services     []*Service
	repositories []*Repository
	models       []*Model
}

type Endpoint struct {
	httpMethod string
	route      string
	handler    string
}

type Controller struct {
	name      string
	endpoints []*Endpoint
}

type Service struct {
	name string
}

type Repository struct {
	name string
}

type Model struct {
	name      string
	structure Structure
}

type Structure []Attribute

type Attribute struct {
	name    string
	atrType string
}
