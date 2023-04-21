package tf

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

type Terraform interface {
	newTerraform() *tfexec.Terraform
	Setup() *tfexec.Terraform
	Init() error
	Plan() error
	Apply() error
	Destroy() error
	Show(planOrApply string) (string, error)
}

type TfConfig struct {
	Version    string
	WorkingDir string
}

type Project struct {
	Name      string
	Variables []string
	TfConfig  TfConfig
	tfBin     *tfexec.Terraform
}

func (p *Project) newTerraform() (*tfexec.Terraform, error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion(p.TfConfig.Version)),
	}

	tfExecutablePath, err := installer.Install(context.Background())

	log.Println("Installing Terraform")

	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	tf, err := tfexec.NewTerraform(p.TfConfig.WorkingDir, tfExecutablePath)

	if err != nil {
		return nil, err
	}

	// tf.SetLog("DEBUG")

	log.Println("Install Terraform successfully")

	return tf, nil
}

func (p *Project) Setup() error {
	log.Println("Setting up Terraform")

	tf, err := p.newTerraform()

	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	p.tfBin = tf

	log.Println("Setup Terraform successfully")

	return nil
}

func (p *Project) Init() error {
	log.Println("Initializing Terraform resources")

	err := p.tfBin.Init(context.Background(), tfexec.Upgrade(true))

	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	return nil
}

func (p *Project) Plan() error {
	log.Println("Planning Terraform changes")

	planFilePath := "./plan.out"
	outFile := tfexec.Out(planFilePath)

	_, err := p.tfBin.Plan(context.Background(), outFile)
	if err != nil {
		log.Fatalf("error running Plan: %s", err)
	}

	planFile, err := p.tfBin.ShowPlanFileRaw(context.Background(), planFilePath)
	if err != nil {
		log.Fatalf("error running ShowPlanFile: %s", err)
	}

	fmt.Printf("planFile: %v", planFile)

	return nil
}

func (p *Project) Apply() error {
	log.Println("Applying Terraform changes")

	return nil
}

func (p *Project) Destroy() error {
	log.Println("Destroying Terraform resources")

	return nil
}

func (p *Project) Show() (string, error) {
	log.Println("Showing state information")

	state, err := p.tfBin.Show(context.Background())

	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	fmt.Println(state)
	fmt.Println(state.TerraformVersion)
	fmt.Println(state.Values)
	fmt.Println(state.FormatVersion) // "0.1"

	return "", nil
}
