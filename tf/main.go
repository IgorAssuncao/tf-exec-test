package tf

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	// "os"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

type command string

// type commands struct {
// 	Plan    command
// 	Apply   command
// 	Destroy command
// 	Show    command
// }

// func newCommands() *commands {
// 	return &commands{
// 		Plan: *Project.plan,
// 	}
// }

const (
	Plan    command = "plan"
	Apply   command = "apply"
	Destroy command = "destroy"
	Show    command = "show"
)

type Terraform interface {
	newTerraform() (*tfexec.Terraform, error)
	setup() error
	init() error
	setupVariables(string) (string, error)
	validate() error
	plan(string) error
	apply() error
	destroy() error
	show() (string, error)
	Run(cmd command)
}

type TfConfig struct {
	Version    string
	WorkingDir string
}

type Project struct {
	Name      string
	Variables ProjectVars
	TfConfig  TfConfig
	tfBin     *tfexec.Terraform
}

var _ Terraform = &Project{}

type ProjectVars map[string]string

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
	logger := log.Default()
	tf.SetLogger(logger)

	log.Println("Install Terraform successfully")

	return tf, nil
}

func (p *Project) setup() error {
	log.Println("Setting up Terraform")

	tf, err := p.newTerraform()

	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	p.tfBin = tf

	log.Println("Setup Terraform successfully")

	return nil
}

func (p *Project) init() error {
	log.Println("Initializing Terraform")

	err := p.tfBin.Init(context.Background(), tfexec.Upgrade(true))

	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	return nil
}

func (p *Project) setupVariables(path string) (string, error) {
	log.Println("Setting up Project Variables")

	// varsFilePath := os.TempDir() + "/" + p.Name + "/variables.tfvars"

	tfexecDir := strings.TrimSuffix(path, "/terraform")

	varsFile, err := os.CreateTemp(tfexecDir, p.Name+".tfvars-")
	if err != nil {
		log.Fatalf("error running CreateTemp in setupVariables(): %s", err)
	}

	// defer varsFile.Close()

	for key, value := range p.Variables {
		varsFile.WriteString(key + " = " + "\"" + value + "\"\n")
	}

	log.Println("Set up Project Variables Successfull")

	return varsFile.Name(), nil
}

func (p *Project) validate() error {
	return nil
}

func (p *Project) plan(varsFile string) error {
	log.Println("Planning Terraform changes")

	planFilePath := "./plan.out"
	outFile := tfexec.Out(planFilePath)

	file := tfexec.VarFile(varsFile)

	_, err := p.tfBin.Plan(context.Background(), outFile, file)
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

func (p *Project) apply() error {
	log.Println("Applying Terraform changes")

	return nil
}

func (p *Project) destroy() error {
	log.Println("Destroying Terraform resources")

	return nil
}

func (p *Project) show() (string, error) {
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

func (p *Project) Run(cmd command) {
	p.setup()
	varsFile, err := p.setupVariables(p.tfBin.ExecPath())
	if err != nil {
		log.Fatalf("error in Run: %s", err)
	}
	p.init()

	switch cmd {
	case "plan":
		p.plan(varsFile)
	case "apply":
		p.plan(varsFile)
		p.apply()
	case "destroy":
		p.destroy()
	case "show":
		p.show()
	default:
		log.Println("Command not found.")
	}
}
