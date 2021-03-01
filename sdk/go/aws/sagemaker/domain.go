// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Sagemaker Domain resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewDomain(ctx, "exampleDomain", &sagemaker.DomainArgs{
// 			DomainName: pulumi.String("example"),
// 			AuthMode:   pulumi.String("IAM"),
// 			VpcId:      pulumi.Any(aws_vpc.Test.Id),
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Test.Id),
// 			},
// 			DefaultUserSettings: &sagemaker.DomainDefaultUserSettingsArgs{
// 				ExecutionRole: pulumi.Any(aws_iam_role.Test.Arn),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"sts:AssumeRole",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "Service",
// 							Identifiers: []string{
// 								"sagemaker.amazonaws.com",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			Path:             pulumi.String("/"),
// 			AssumeRolePolicy: pulumi.String(examplePolicyDocument.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using Custom Images
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testImage, err := sagemaker.NewImage(ctx, "testImage", &sagemaker.ImageArgs{
// 			ImageName: pulumi.String("example"),
// 			RoleArn:   pulumi.Any(aws_iam_role.Test.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testAppImageConfig, err := sagemaker.NewAppImageConfig(ctx, "testAppImageConfig", &sagemaker.AppImageConfigArgs{
// 			AppImageConfigName: pulumi.String("example"),
// 			KernelGatewayImageConfig: &sagemaker.AppImageConfigKernelGatewayImageConfigArgs{
// 				KernelSpec: &sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs{
// 					Name: pulumi.String("example"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testImageVersion, err := sagemaker.NewImageVersion(ctx, "testImageVersion", &sagemaker.ImageVersionArgs{
// 			ImageName: testImage.ID(),
// 			BaseImage: pulumi.String("base-image"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sagemaker.NewDomain(ctx, "testDomain", &sagemaker.DomainArgs{
// 			DomainName: pulumi.String("example"),
// 			AuthMode:   pulumi.String("IAM"),
// 			VpcId:      pulumi.Any(aws_vpc.Test.Id),
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Test.Id),
// 			},
// 			DefaultUserSettings: &sagemaker.DomainDefaultUserSettingsArgs{
// 				ExecutionRole: pulumi.Any(aws_iam_role.Test.Arn),
// 				KernelGatewayAppSettings: &sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsArgs{
// 					CustomImages: sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArray{
// 						&sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArgs{
// 							AppImageConfigName: testAppImageConfig.AppImageConfigName,
// 							ImageName:          testImageVersion.ImageName,
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Sagemaker Code Domains can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/domain:Domain test_domain d-8jgsjtilstu8
// ```
type Domain struct {
	pulumi.CustomResourceState

	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType pulumi.StringPtrOutput `pulumi:"appNetworkAccessType"`
	// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode pulumi.StringOutput `pulumi:"authMode"`
	// The default user settings. See Default User Settings below.
	DefaultUserSettings DomainDefaultUserSettingsOutput `pulumi:"defaultUserSettings"`
	// The domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId pulumi.StringOutput `pulumi:"homeEfsFileSystemId"`
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId pulumi.StringOutput `pulumi:"singleSignOnManagedApplicationInstanceId"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The domain's URL.
	Url pulumi.StringOutput `pulumi:"url"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthMode == nil {
		return nil, errors.New("invalid value for required argument 'AuthMode'")
	}
	if args.DefaultUserSettings == nil {
		return nil, errors.New("invalid value for required argument 'DefaultUserSettings'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Domain
	err := ctx.RegisterResource("aws:sagemaker/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("aws:sagemaker/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType *string `pulumi:"appNetworkAccessType"`
	// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
	Arn *string `pulumi:"arn"`
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode *string `pulumi:"authMode"`
	// The default user settings. See Default User Settings below.
	DefaultUserSettings *DomainDefaultUserSettings `pulumi:"defaultUserSettings"`
	// The domain name.
	DomainName *string `pulumi:"domainName"`
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId *string `pulumi:"homeEfsFileSystemId"`
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId *string `pulumi:"singleSignOnManagedApplicationInstanceId"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The domain's URL.
	Url *string `pulumi:"url"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId *string `pulumi:"vpcId"`
}

type DomainState struct {
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
	Arn pulumi.StringPtrInput
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode pulumi.StringPtrInput
	// The default user settings. See Default User Settings below.
	DefaultUserSettings DomainDefaultUserSettingsPtrInput
	// The domain name.
	DomainName pulumi.StringPtrInput
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId pulumi.StringPtrInput
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId pulumi.StringPtrInput
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId pulumi.StringPtrInput
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The domain's URL.
	Url pulumi.StringPtrInput
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType *string `pulumi:"appNetworkAccessType"`
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode string `pulumi:"authMode"`
	// The default user settings. See Default User Settings below.
	DefaultUserSettings DomainDefaultUserSettings `pulumi:"defaultUserSettings"`
	// The domain name.
	DomainName string `pulumi:"domainName"`
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType pulumi.StringPtrInput
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode pulumi.StringInput
	// The default user settings. See Default User Settings below.
	DefaultUserSettings DomainDefaultUserSettingsInput
	// The domain name.
	DomainName pulumi.StringInput
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId pulumi.StringPtrInput
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

func (i *Domain) ToDomainPtrOutput() DomainPtrOutput {
	return i.ToDomainPtrOutputWithContext(context.Background())
}

func (i *Domain) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPtrOutput)
}

type DomainPtrInput interface {
	pulumi.Input

	ToDomainPtrOutput() DomainPtrOutput
	ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput
}

type domainPtrType DomainArgs

func (*domainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil))
}

func (i *domainPtrType) ToDomainPtrOutput() DomainPtrOutput {
	return i.ToDomainPtrOutputWithContext(context.Background())
}

func (i *domainPtrType) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPtrOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//          DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Domain)(nil))
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//          DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Domain)(nil))
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o.ToDomainPtrOutputWithContext(context.Background())
}

func (o DomainOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o.ApplyT(func(v Domain) *Domain {
		return &v
	}).(DomainPtrOutput)
}

type DomainPtrOutput struct {
	*pulumi.OutputState
}

func (DomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil))
}

func (o DomainPtrOutput) ToDomainPtrOutput() DomainPtrOutput {
	return o
}

func (o DomainPtrOutput) ToDomainPtrOutputWithContext(ctx context.Context) DomainPtrOutput {
	return o
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Domain)(nil))
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Domain {
		return vs[0].([]Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Domain)(nil))
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Domain {
		return vs[0].(map[string]Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainPtrOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
