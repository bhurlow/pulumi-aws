// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
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
	case "aws:licensemanager/association:Association":
		r = &Association{}
	case "aws:licensemanager/licenseConfiguration:LicenseConfiguration":
		r = &LicenseConfiguration{}
	case "aws:licensemanager/licenseGrant:LicenseGrant":
		r = &LicenseGrant{}
	case "aws:licensemanager/licenseGrantAccepter:LicenseGrantAccepter":
		r = &LicenseGrantAccepter{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws",
		"licensemanager/association",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"licensemanager/licenseConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"licensemanager/licenseGrant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"licensemanager/licenseGrantAccepter",
		&module{version},
	)
}
