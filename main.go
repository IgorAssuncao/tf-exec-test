package main

import (
	"tf-exec/tf"
)

func main() {
	project := &tf.Project{
		Name:      "tf-exec-test",
		Variables: []string{},
		TfConfig: tf.TfConfig{
			Version:    "1.3.2",
			WorkingDir: "./terraform/repositories/",
		},
	}

	project.Setup()
	project.Init()
	project.Plan()
	project.Show()
}
