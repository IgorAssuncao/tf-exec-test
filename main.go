package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func main() {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	workingDir := "./repositories"
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	fmt.Println(state.TerraformVersion)
	fmt.Println(state.Values)
	fmt.Println(state.FormatVersion) // "0.1"

	planFilePath := "./plan.out"
	outFile := tfexec.Out(planFilePath)

	_, err = tf.Plan(context.Background(), outFile)
	if err != nil {
		log.Fatalf("error running Plan: %s", err)
	}

	planFile, err := tf.ShowPlanFileRaw(context.Background(), planFilePath)
	if err != nil {
		log.Fatalf("error running ShowPlanFile: %s", err)
	}

	fmt.Printf("planFile: %v", planFile)
}