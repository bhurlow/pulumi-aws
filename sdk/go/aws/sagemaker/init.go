// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:sagemaker/app:App":
		r, err = NewApp(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/appImageConfig:AppImageConfig":
		r, err = NewAppImageConfig(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/codeRepository:CodeRepository":
		r, err = NewCodeRepository(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/domain:Domain":
		r, err = NewDomain(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/endpoint:Endpoint":
		r, err = NewEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/endpointConfiguration:EndpointConfiguration":
		r, err = NewEndpointConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/featureGroup:FeatureGroup":
		r, err = NewFeatureGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/image:Image":
		r, err = NewImage(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/imageVersion:ImageVersion":
		r, err = NewImageVersion(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/model:Model":
		r, err = NewModel(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/modelPackageGroup:ModelPackageGroup":
		r, err = NewModelPackageGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/notebookInstance:NotebookInstance":
		r, err = NewNotebookInstance(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration":
		r, err = NewNotebookInstanceLifecycleConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "aws:sagemaker/userProfile:UserProfile":
		r, err = NewUserProfile(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/app",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/appImageConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/codeRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/endpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/endpointConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/featureGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/image",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/imageVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/model",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/modelPackageGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/notebookInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/notebookInstanceLifecycleConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"sagemaker/userProfile",
		&module{version},
	)
}
