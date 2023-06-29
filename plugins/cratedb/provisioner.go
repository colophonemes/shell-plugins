package cratedb

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

var argsToProvision = []string{
	"--username", fieldname.Username,
	"--hosts", fieldname.Host,
}

type crateArgsProvisioner struct {
}

func crateProvisioner() sdk.Provisioner {
	return crateArgsProvisioner{}
}

func (p crateArgsProvisioner) Provision(ctx context.Context, in sdk.ProvisionInput, out *sdk.ProvisionOutput) {
	if value, ok := in.ItemFields[fieldname.Password]; ok {
		out.AddEnvVar("CRATEPW", value)
	}
	
	for i, arg := range argsToProvision {
		if i%2 == 0 {
			argName := arg
			fieldName := sdk.FieldName(argsToProvision[i+1])
			if fieldValue, ok := in.ItemFields[fieldName]; ok {
				out.AddArgs(argName, fieldValue)
			}
		}
	
}

func (p crateArgsProvisioner) Deprovision(ctx context.Context, in sdk.DeprovisionInput, out *sdk.DeprovisionOutput) {
	// Nothing to do here: credentials get wiped automatically when the process exits.
}

func (p crateArgsProvisioner) Description() string {
	return "Provision CrateDB username, host as command-line arguments && Password as Env ."
}
