package commands

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry/bosh-bootloader/storage"
)

const (
	UsageHeader = `
Usage:
  bbl [GLOBAL OPTIONS] %s [OPTIONS]

Global Options:
  --help      [-h]       Prints usage. Use "bbl [command] --help" for more information about a command
  --state-dir            Directory containing the bbl state
  --debug                Prints debugging output
  --version   [-v]       Prints version
%s
`
	CommandUsage = `
[%s command options]
  %s`
)

const GlobalUsage = `
Basic Commands: A good place to start
  up                      Deploys BOSH director on an IAAS. Updates existing director
  print-env               All environment variables needed for targeting BOSH. Use with: eval "$(bbl print-env)"
  create-lbs              Creates recommended load balancer(s) for CF, Concourse

Maintenance Lifecycle Commands:
  destroy                 Tears down BOSH director infrastructure. Cleans up state directory
  update-lbs              Updates load balancer(s)
  delete-lbs              Deletes attached load balancer(s)
  rotate                  Rotates SSH key for the jumpbox user
  plan                    Populates a state directory with the latest config without applying it

Environmental Detail Commands: Useful for automation and gaining access
  bosh-deployment-vars    Prints required variables for BOSH deployment
  jumpbox-deployment-vars Prints required variables for jumpbox deployment
  cloud-config            Prints suggested cloud configuration for BOSH environment
  jumpbox-address         Prints BOSH jumpbox address
  director-address        Prints BOSH director address
  director-username       Prints BOSH director username
  director-password       Prints BOSH director password
  director-ca-cert        Prints BOSH director CA certificate
  env-id                  Prints environment ID
  ssh-key                 Prints jumpbox SSH private key
  director-ssh-key        Prints director SSH private key
  lbs                     Prints load balancer(s) and DNS records

Troubleshooting Commands:
  help                    Prints usage
  version                 Prints version
  latest-error            Prints the output from the latest call to terraform`

type Usage struct {
	logger logger
}

func NewUsage(logger logger) Usage {
	return Usage{
		logger: logger,
	}
}

func (u Usage) CheckFastFails(subcommandFlags []string, state storage.State) error {
	return nil
}

func (u Usage) Execute(subcommandFlags []string, state storage.State) error {
	u.Print()
	return nil
}

func (u Usage) Print() {
	content := fmt.Sprintf(UsageHeader, "COMMAND", GlobalUsage)
	u.logger.Println(strings.TrimLeft(content, "\n"))
}

func (u Usage) PrintCommandUsage(command, message string) {
	commandUsage := fmt.Sprintf(CommandUsage, command, message)
	content := fmt.Sprintf(UsageHeader, command, commandUsage)
	u.logger.Println(strings.TrimLeft(content, "\n"))
}
