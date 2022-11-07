package registry

import (
	"context"

	"github.com/hashicorp/hcl/v2"
	sdkpacker "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer/hcl2template"
	"github.com/hashicorp/packer/internal/registry/hcp"
	"github.com/hashicorp/packer/packer"
)

// Registry is an entity capable to orchestrate a Packer build and upload metadata to HCP
type Registry interface {
	PopulateIteration(context.Context) error
	BuildStart(context.Context, string) error
	BuildDone(ctx context.Context, buildName string, artifacts []sdkpacker.Artifact, buildErr error) ([]sdkpacker.Artifact, error)
}

// GetRegistry instanciates the appropriate handler for the configuration type given as parameter.
//
// If no HCP-related data is present, it will be a NoopHandler.
func GetRegistry(cfg packer.Handler) (Registry, hcl.Diagnostics) {
	if !hcp.IsHCPEnabled(cfg) {
		return newNullHandler(), nil
	}

	switch config := cfg.(type) {
	case *hcl2template.PackerConfig:
		return hcp.NewHCLRegistry(config)
	case *packer.Core:
		return hcp.NewJSONRegistry(config)
	}

	return nil, hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Unknown Config type",
			Detail: "The config type %s does not match a Packer-known template type. " +
				"This is a Packer error and should be brought up to the Packer " +
				"team via a Github Issue.",
		},
	}
}
