package main

import (
	"fmt"

	"github.com/csothen/go-gen/pkg/generator"
)

func main() {

	p, _ := generator.NewProject("myProject", "./")

	a1 := p.NewApplication("app1")
	a2 := p.NewApplication("app2")

	a1.NewController("app1_user_controller")
	a1.NewController("app1_cars_controller")

	a2.NewController("app2_house_controller")
	a2.NewController("app2_logs_controller")

	a1.NewService("app1_service")
	a2.NewService("app2_service")

	//p.Describe()
	err := p.Build()
	if err != nil {
		fmt.Println(err)
	}
}
