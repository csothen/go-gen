package generator

func (ctrl *Controller) NewEndpoint(httpMethod, path, controllerMethod string) *Endpoint {
	endpoint := &Endpoint{httpMethod, path, controllerMethod}

	ctrl.endpoints = append(ctrl.endpoints, endpoint)

	return endpoint
}

func (ctrl *Controller) build() error {
	for _, endpoint := range ctrl.endpoints {
		if err := endpoint.build(); err != nil {
			return err
		}
	}

	// Build actual code for controller

	return nil
}
