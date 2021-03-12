package generator

func (ctrl *Controller) NewEndpoint(httpMethod, path, controllerMethod string) *Endpoint {
	endpoint := &Endpoint{httpMethod, path, controllerMethod}

	ctrl.endpoints = append(ctrl.endpoints, endpoint)

	return endpoint
}
