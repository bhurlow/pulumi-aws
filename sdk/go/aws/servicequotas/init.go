// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

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
	case "aws:servicequotas/serviceQuota:ServiceQuota":
		r = &ServiceQuota{}
	case "aws:servicequotas/template:Template":
		r = &Template{}
	case "aws:servicequotas/templateAssociation:TemplateAssociation":
		r = &TemplateAssociation{}
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
		"servicequotas/serviceQuota",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicequotas/template",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicequotas/templateAssociation",
		&module{version},
	)
}
