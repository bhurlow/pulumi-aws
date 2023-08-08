// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

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
	case "aws:cognito/identityPool:IdentityPool":
		r = &IdentityPool{}
	case "aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag":
		r = &IdentityPoolProviderPrincipalTag{}
	case "aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment":
		r = &IdentityPoolRoleAttachment{}
	case "aws:cognito/identityProvider:IdentityProvider":
		r = &IdentityProvider{}
	case "aws:cognito/managedUserPoolClient:ManagedUserPoolClient":
		r = &ManagedUserPoolClient{}
	case "aws:cognito/resourceServer:ResourceServer":
		r = &ResourceServer{}
	case "aws:cognito/riskConfiguration:RiskConfiguration":
		r = &RiskConfiguration{}
	case "aws:cognito/user:User":
		r = &User{}
	case "aws:cognito/userGroup:UserGroup":
		r = &UserGroup{}
	case "aws:cognito/userInGroup:UserInGroup":
		r = &UserInGroup{}
	case "aws:cognito/userPool:UserPool":
		r = &UserPool{}
	case "aws:cognito/userPoolClient:UserPoolClient":
		r = &UserPoolClient{}
	case "aws:cognito/userPoolDomain:UserPoolDomain":
		r = &UserPoolDomain{}
	case "aws:cognito/userPoolUICustomization:UserPoolUICustomization":
		r = &UserPoolUICustomization{}
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
		"cognito/identityPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/identityPoolProviderPrincipalTag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/identityPoolRoleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/identityProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/managedUserPoolClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/resourceServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/riskConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userInGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPoolClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPoolDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPoolUICustomization",
		&module{version},
	)
}
