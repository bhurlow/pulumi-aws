// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SageMaker Domain resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"sagemaker.amazonaws.com",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
//				Path:             pulumi.String("/"),
//				AssumeRolePolicy: *pulumi.String(examplePolicyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sagemaker.NewDomain(ctx, "exampleDomain", &sagemaker.DomainArgs{
//				DomainName: pulumi.String("example"),
//				AuthMode:   pulumi.String("IAM"),
//				VpcId:      pulumi.Any(aws_vpc.Example.Id),
//				SubnetIds: pulumi.StringArray{
//					aws_subnet.Example.Id,
//				},
//				DefaultUserSettings: &sagemaker.DomainDefaultUserSettingsArgs{
//					ExecutionRole: exampleRole.Arn,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Using Custom Images
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleImage, err := sagemaker.NewImage(ctx, "exampleImage", &sagemaker.ImageArgs{
//				ImageName: pulumi.String("example"),
//				RoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
//			})
//			if err != nil {
//				return err
//			}
//			exampleAppImageConfig, err := sagemaker.NewAppImageConfig(ctx, "exampleAppImageConfig", &sagemaker.AppImageConfigArgs{
//				AppImageConfigName: pulumi.String("example"),
//				KernelGatewayImageConfig: &sagemaker.AppImageConfigKernelGatewayImageConfigArgs{
//					KernelSpec: &sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs{
//						Name: pulumi.String("example"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleImageVersion, err := sagemaker.NewImageVersion(ctx, "exampleImageVersion", &sagemaker.ImageVersionArgs{
//				ImageName: exampleImage.ID(),
//				BaseImage: pulumi.String("base-image"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sagemaker.NewDomain(ctx, "exampleDomain", &sagemaker.DomainArgs{
//				DomainName: pulumi.String("example"),
//				AuthMode:   pulumi.String("IAM"),
//				VpcId:      pulumi.Any(aws_vpc.Example.Id),
//				SubnetIds: pulumi.StringArray{
//					aws_subnet.Example.Id,
//				},
//				DefaultUserSettings: &sagemaker.DomainDefaultUserSettingsArgs{
//					ExecutionRole: pulumi.Any(aws_iam_role.Example.Arn),
//					KernelGatewayAppSettings: &sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsArgs{
//						CustomImages: sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArray{
//							&sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArgs{
//								AppImageConfigName: exampleAppImageConfig.AppImageConfigName,
//								ImageName:          exampleImageVersion.ImageName,
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// SageMaker Domains can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:sagemaker/domain:Domain test_domain d-8jgsjtilstu8
//
// ```
type Domain struct {
	pulumi.CustomResourceState

	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType pulumi.StringPtrOutput `pulumi:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
	AppSecurityGroupManagement pulumi.StringPtrOutput `pulumi:"appSecurityGroupManagement"`
	// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode pulumi.StringOutput `pulumi:"authMode"`
	// The default space settings. See Default Space Settings below.
	DefaultSpaceSettings DomainDefaultSpaceSettingsPtrOutput `pulumi:"defaultSpaceSettings"`
	// The default user settings. See Default User Settings below.* `domainName` - (Required) The domain name.
	DefaultUserSettings DomainDefaultUserSettingsOutput `pulumi:"defaultUserSettings"`
	DomainName          pulumi.StringOutput             `pulumi:"domainName"`
	// The domain's settings.
	DomainSettings DomainDomainSettingsPtrOutput `pulumi:"domainSettings"`
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId pulumi.StringOutput `pulumi:"homeEfsFileSystemId"`
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
	RetentionPolicy DomainRetentionPolicyPtrOutput `pulumi:"retentionPolicy"`
	// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
	SecurityGroupIdForDomainBoundary pulumi.StringOutput `pulumi:"securityGroupIdForDomainBoundary"`
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId pulumi.StringOutput `pulumi:"singleSignOnManagedApplicationInstanceId"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
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
	// The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
	AppSecurityGroupManagement *string `pulumi:"appSecurityGroupManagement"`
	// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
	Arn *string `pulumi:"arn"`
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode *string `pulumi:"authMode"`
	// The default space settings. See Default Space Settings below.
	DefaultSpaceSettings *DomainDefaultSpaceSettings `pulumi:"defaultSpaceSettings"`
	// The default user settings. See Default User Settings below.* `domainName` - (Required) The domain name.
	DefaultUserSettings *DomainDefaultUserSettings `pulumi:"defaultUserSettings"`
	DomainName          *string                    `pulumi:"domainName"`
	// The domain's settings.
	DomainSettings *DomainDomainSettings `pulumi:"domainSettings"`
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId *string `pulumi:"homeEfsFileSystemId"`
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
	RetentionPolicy *DomainRetentionPolicy `pulumi:"retentionPolicy"`
	// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
	SecurityGroupIdForDomainBoundary *string `pulumi:"securityGroupIdForDomainBoundary"`
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId *string `pulumi:"singleSignOnManagedApplicationInstanceId"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The domain's URL.
	Url *string `pulumi:"url"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId *string `pulumi:"vpcId"`
}

type DomainState struct {
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType pulumi.StringPtrInput
	// The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
	AppSecurityGroupManagement pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
	Arn pulumi.StringPtrInput
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode pulumi.StringPtrInput
	// The default space settings. See Default Space Settings below.
	DefaultSpaceSettings DomainDefaultSpaceSettingsPtrInput
	// The default user settings. See Default User Settings below.* `domainName` - (Required) The domain name.
	DefaultUserSettings DomainDefaultUserSettingsPtrInput
	DomainName          pulumi.StringPtrInput
	// The domain's settings.
	DomainSettings DomainDomainSettingsPtrInput
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId pulumi.StringPtrInput
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId pulumi.StringPtrInput
	// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
	RetentionPolicy DomainRetentionPolicyPtrInput
	// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
	SecurityGroupIdForDomainBoundary pulumi.StringPtrInput
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId pulumi.StringPtrInput
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
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
	// The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
	AppSecurityGroupManagement *string `pulumi:"appSecurityGroupManagement"`
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode string `pulumi:"authMode"`
	// The default space settings. See Default Space Settings below.
	DefaultSpaceSettings *DomainDefaultSpaceSettings `pulumi:"defaultSpaceSettings"`
	// The default user settings. See Default User Settings below.* `domainName` - (Required) The domain name.
	DefaultUserSettings DomainDefaultUserSettings `pulumi:"defaultUserSettings"`
	DomainName          string                    `pulumi:"domainName"`
	// The domain's settings.
	DomainSettings *DomainDomainSettings `pulumi:"domainSettings"`
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
	RetentionPolicy *DomainRetentionPolicy `pulumi:"retentionPolicy"`
	// The VPC subnets that Studio uses for communication.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
	AppNetworkAccessType pulumi.StringPtrInput
	// The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
	AppSecurityGroupManagement pulumi.StringPtrInput
	// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
	AuthMode pulumi.StringInput
	// The default space settings. See Default Space Settings below.
	DefaultSpaceSettings DomainDefaultSpaceSettingsPtrInput
	// The default user settings. See Default User Settings below.* `domainName` - (Required) The domain name.
	DefaultUserSettings DomainDefaultUserSettingsInput
	DomainName          pulumi.StringInput
	// The domain's settings.
	DomainSettings DomainDomainSettingsPtrInput
	// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
	KmsKeyId pulumi.StringPtrInput
	// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
	RetentionPolicy DomainRetentionPolicyPtrInput
	// The VPC subnets that Studio uses for communication.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//	DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
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
//	DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`. Valid values are `PublicInternetOnly` and `VpcOnly`.
func (o DomainOutput) AppNetworkAccessType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.AppNetworkAccessType }).(pulumi.StringPtrOutput)
}

// The entity that creates and manages the required security groups for inter-app communication in `VPCOnly` mode. Valid values are `Service` and `Customer`.
func (o DomainOutput) AppSecurityGroupManagement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.AppSecurityGroupManagement }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) assigned by AWS to this Domain.
func (o DomainOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The mode of authentication that members use to access the domain. Valid values are `IAM` and `SSO`.
func (o DomainOutput) AuthMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.AuthMode }).(pulumi.StringOutput)
}

// The default space settings. See Default Space Settings below.
func (o DomainOutput) DefaultSpaceSettings() DomainDefaultSpaceSettingsPtrOutput {
	return o.ApplyT(func(v *Domain) DomainDefaultSpaceSettingsPtrOutput { return v.DefaultSpaceSettings }).(DomainDefaultSpaceSettingsPtrOutput)
}

// The default user settings. See Default User Settings below.* `domainName` - (Required) The domain name.
func (o DomainOutput) DefaultUserSettings() DomainDefaultUserSettingsOutput {
	return o.ApplyT(func(v *Domain) DomainDefaultUserSettingsOutput { return v.DefaultUserSettings }).(DomainDefaultUserSettingsOutput)
}

func (o DomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The domain's settings.
func (o DomainOutput) DomainSettings() DomainDomainSettingsPtrOutput {
	return o.ApplyT(func(v *Domain) DomainDomainSettingsPtrOutput { return v.DomainSettings }).(DomainDomainSettingsPtrOutput)
}

// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
func (o DomainOutput) HomeEfsFileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.HomeEfsFileSystemId }).(pulumi.StringOutput)
}

// The AWS KMS customer managed CMK used to encrypt the EFS volume attached to the domain.
func (o DomainOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The retention policy for this domain, which specifies whether resources will be retained after the Domain is deleted. By default, all resources are retained. See Retention Policy below.
func (o DomainOutput) RetentionPolicy() DomainRetentionPolicyPtrOutput {
	return o.ApplyT(func(v *Domain) DomainRetentionPolicyPtrOutput { return v.RetentionPolicy }).(DomainRetentionPolicyPtrOutput)
}

// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
func (o DomainOutput) SecurityGroupIdForDomainBoundary() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.SecurityGroupIdForDomainBoundary }).(pulumi.StringOutput)
}

// The SSO managed application instance ID.
func (o DomainOutput) SingleSignOnManagedApplicationInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.SingleSignOnManagedApplicationInstanceId }).(pulumi.StringOutput)
}

// The VPC subnets that Studio uses for communication.
func (o DomainOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o DomainOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The domain's URL.
func (o DomainOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
func (o DomainOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].([]*Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].(map[string]*Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainArrayInput)(nil)).Elem(), DomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMapInput)(nil)).Elem(), DomainMap{})
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
