// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:ram/principalAssociation:PrincipalAssociation":
		r = &PrincipalAssociation{}
	case "aws:ram/resourceAssociation:ResourceAssociation":
		r = &ResourceAssociation{}
	case "aws:ram/resourceShare:ResourceShare":
		r = &ResourceShare{}
	case "aws:ram/resourceShareAccepter:ResourceShareAccepter":
		r = &ResourceShareAccepter{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws",
		"ram/principalAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ram/resourceAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ram/resourceShare",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"ram/resourceShareAccepter",
		&module{version},
	)
}
