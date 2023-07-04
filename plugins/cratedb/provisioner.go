package cratedb

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

type CrateArgsProvisioner struct {
}



func (p CrateArgsProvisioner) Provision(ctx context.Context, in sdk.ProvisionInput, out *sdk.ProvisionOutput) {
	if value, ok := in.ItemFields[fieldname.Password]; ok {
		out.AddEnvVar("CRATEPW", value)
	}
	
	var user, host string
	if fieldValue, ok := in.ItemFields[fieldname.Username]; ok {
		user=fieldValue
	}
	if fieldValue, ok := in.ItemFields[fieldname.Host]; ok {
		host=fieldValue
	}

	out.CommandLine = []string{"--username", user, "--hosts", host, }

	
}

func (p CrateArgsProvisioner) Deprovision(ctx context.Context, in sdk.DeprovisionInput, out *sdk.DeprovisionOutput) {
	// Nothing to do here: credentials get wiped automatically when the process exits.
}

func (p CrateArgsProvisioner) Description() string {
	return "Provision CrateDB username, host as command-line arguments && Password as Env ."
}