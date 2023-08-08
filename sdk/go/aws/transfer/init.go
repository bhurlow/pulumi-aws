// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

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
	case "aws:transfer/access:Access":
		r = &Access{}
	case "aws:transfer/agreement:Agreement":
		r = &Agreement{}
	case "aws:transfer/certificate:Certificate":
		r = &Certificate{}
	case "aws:transfer/connector:Connector":
		r = &Connector{}
	case "aws:transfer/profile:Profile":
		r = &Profile{}
	case "aws:transfer/server:Server":
		r = &Server{}
	case "aws:transfer/sshKey:SshKey":
		r = &SshKey{}
	case "aws:transfer/tag:Tag":
		r = &Tag{}
	case "aws:transfer/user:User":
		r = &User{}
	case "aws:transfer/workflow:Workflow":
		r = &Workflow{}
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
		"transfer/access",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/agreement",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/connector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/profile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/server",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/sshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/tag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"transfer/workflow",
		&module{version},
	)
}
