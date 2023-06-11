package main

import (
	"log"
	"os"
	"tf-exec/tf"
)

func main() {
	project := &tf.Project{
		Name: "tf-exec-test",
		TfConfig: tf.TfConfig{
			Version:    "1.4.6",
			WorkingDir: "./terraform/repositories/",
		},
	}

	gh_token, ok := os.LookupEnv("GH_TOKEN")

	if !ok {
		log.Fatalln("GH_TOKEN not found in env")
	}

	variables := tf.ProjectVars{
		"gh_token": gh_token,
	}

	project.Variables = variables

	project.Run(tf.Plan)
	// project.Run(tf.Show)
	project.Run(tf.Destroy)
}
