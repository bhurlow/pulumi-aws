// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

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
	case "aws:servicecatalog/budgetResourceAssociation:BudgetResourceAssociation":
		r = &BudgetResourceAssociation{}
	case "aws:servicecatalog/constraint:Constraint":
		r = &Constraint{}
	case "aws:servicecatalog/organizationsAccess:OrganizationsAccess":
		r = &OrganizationsAccess{}
	case "aws:servicecatalog/portfolio:Portfolio":
		r = &Portfolio{}
	case "aws:servicecatalog/portfolioShare:PortfolioShare":
		r = &PortfolioShare{}
	case "aws:servicecatalog/principalPortfolioAssociation:PrincipalPortfolioAssociation":
		r = &PrincipalPortfolioAssociation{}
	case "aws:servicecatalog/product:Product":
		r = &Product{}
	case "aws:servicecatalog/productPortfolioAssociation:ProductPortfolioAssociation":
		r = &ProductPortfolioAssociation{}
	case "aws:servicecatalog/provisionedProduct:ProvisionedProduct":
		r = &ProvisionedProduct{}
	case "aws:servicecatalog/provisioningArtifact:ProvisioningArtifact":
		r = &ProvisioningArtifact{}
	case "aws:servicecatalog/serviceAction:ServiceAction":
		r = &ServiceAction{}
	case "aws:servicecatalog/tagOption:TagOption":
		r = &TagOption{}
	case "aws:servicecatalog/tagOptionResourceAssociation:TagOptionResourceAssociation":
		r = &TagOptionResourceAssociation{}
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
		"servicecatalog/budgetResourceAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/constraint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/organizationsAccess",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/portfolio",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/portfolioShare",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/principalPortfolioAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/product",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/productPortfolioAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/provisionedProduct",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/provisioningArtifact",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/serviceAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/tagOption",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicecatalog/tagOptionResourceAssociation",
		&module{version},
	)
}
