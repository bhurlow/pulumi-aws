// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoverycontrol

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
	case "aws:route53recoverycontrol/cluster:Cluster":
		r = &Cluster{}
	case "aws:route53recoverycontrol/controlPanel:ControlPanel":
		r = &ControlPanel{}
	case "aws:route53recoverycontrol/routingControl:RoutingControl":
		r = &RoutingControl{}
	case "aws:route53recoverycontrol/safetyRule:SafetyRule":
		r = &SafetyRule{}
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
		"route53recoverycontrol/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53recoverycontrol/controlPanel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53recoverycontrol/routingControl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53recoverycontrol/safetyRule",
		&module{version},
	)
}
