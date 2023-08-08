// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

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
	case "aws:quicksight/accountSubscription:AccountSubscription":
		r = &AccountSubscription{}
	case "aws:quicksight/analysis:Analysis":
		r = &Analysis{}
	case "aws:quicksight/dashboard:Dashboard":
		r = &Dashboard{}
	case "aws:quicksight/dataSet:DataSet":
		r = &DataSet{}
	case "aws:quicksight/dataSource:DataSource":
		r = &DataSource{}
	case "aws:quicksight/folder:Folder":
		r = &Folder{}
	case "aws:quicksight/folderMembership:FolderMembership":
		r = &FolderMembership{}
	case "aws:quicksight/group:Group":
		r = &Group{}
	case "aws:quicksight/groupMembership:GroupMembership":
		r = &GroupMembership{}
	case "aws:quicksight/iamPolicyAssignment:IamPolicyAssignment":
		r = &IamPolicyAssignment{}
	case "aws:quicksight/ingestion:Ingestion":
		r = &Ingestion{}
	case "aws:quicksight/namespace:Namespace":
		r = &Namespace{}
	case "aws:quicksight/refreshSchedule:RefreshSchedule":
		r = &RefreshSchedule{}
	case "aws:quicksight/template:Template":
		r = &Template{}
	case "aws:quicksight/templateAlias:TemplateAlias":
		r = &TemplateAlias{}
	case "aws:quicksight/theme:Theme":
		r = &Theme{}
	case "aws:quicksight/user:User":
		r = &User{}
	case "aws:quicksight/vpcConnection:VpcConnection":
		r = &VpcConnection{}
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
		"quicksight/accountSubscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/analysis",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/dashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/dataSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/dataSource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/folder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/folderMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/groupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/iamPolicyAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/ingestion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/namespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/refreshSchedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/template",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/templateAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/theme",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"quicksight/vpcConnection",
		&module{version},
	)
}
