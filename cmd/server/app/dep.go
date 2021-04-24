package app

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Dependency of routers and their controllers
type Dependency struct {
	ProjectConfig
}

var dep Dependency

// SetDep returns the initialized dependency of Rubik app
func SetDep(d Dependency) error {
	b, err := ioutil.ReadFile(d.ProjectConfig.PodServerConfig.Path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &d.ProjectConfig.PodDetails)
	if err != nil {
		return err
	}

	dep = d
	fmt.Printf("%+v \n", d)

	return nil
}

// GetDep returns the initialized dependency of Rubik app
func GetDep() Dependency {
	return dep
}
