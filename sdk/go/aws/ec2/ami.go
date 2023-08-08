// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AMI resource allows the creation and management of a completely-custom
// *Amazon Machine Image* (AMI).
//
// If you just want to duplicate an existing AMI, possibly copying it to another
// region, it's better to use `ec2.AmiCopy` instead.
//
// If you just want to share an existing AMI with another AWS account,
// it's better to use `ec2.AmiLaunchPermission` instead.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewAmi(ctx, "example", &ec2.AmiArgs{
//				EbsBlockDevices: ec2.AmiEbsBlockDeviceArray{
//					&ec2.AmiEbsBlockDeviceArgs{
//						DeviceName: pulumi.String("/dev/xvda"),
//						SnapshotId: pulumi.String("snap-xxxxxxxx"),
//						VolumeSize: pulumi.Int(8),
//					},
//				},
//				ImdsSupport:        pulumi.String("v2.0"),
//				RootDeviceName:     pulumi.String("/dev/xvda"),
//				VirtualizationType: pulumi.String("hvm"),
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
// terraform import {
//
//	to = aws_ami.example
//
//	id = "ami-12345678" } Using `pulumi import`, import `aws_ami` using the ID of the AMI. For exampleconsole % pulumi import aws_ami.example ami-12345678
type Ami struct {
	pulumi.CustomResourceState

	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// ARN of the AMI.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
	BootMode pulumi.StringPtrOutput `pulumi:"bootMode"`
	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime pulumi.StringPtrOutput `pulumi:"deprecationTime"`
	// Longer, human-readable description for the AMI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiEbsBlockDeviceArrayOutput `pulumi:"ebsBlockDevices"`
	// Whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrOutput `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiEphemeralBlockDeviceArrayOutput `pulumi:"ephemeralBlockDevices"`
	// Hypervisor type of the image.
	Hypervisor pulumi.StringOutput `pulumi:"hypervisor"`
	// Path to an S3 object containing an image manifest, e.g., created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringOutput `pulumi:"imageLocation"`
	// AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.
	ImageOwnerAlias pulumi.StringOutput `pulumi:"imageOwnerAlias"`
	// Type of image.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to `v2.0`. For more information, see [Configure instance metadata options for new instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration).
	ImdsSupport pulumi.StringPtrOutput `pulumi:"imdsSupport"`
	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId           pulumi.StringPtrOutput `pulumi:"kernelId"`
	ManageEbsSnapshots pulumi.BoolOutput      `pulumi:"manageEbsSnapshots"`
	// Region-unique name for the AMI.
	Name pulumi.StringOutput `pulumi:"name"`
	// AWS account ID of the image owner.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// This value is set to windows for Windows AMIs; otherwise, it is blank.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// Platform details associated with the billing code of the AMI.
	PlatformDetails pulumi.StringOutput `pulumi:"platformDetails"`
	// Whether the image has public launch permissions.
	Public pulumi.BoolOutput `pulumi:"public"`
	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrOutput `pulumi:"ramdiskId"`
	// Name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrOutput `pulumi:"rootDeviceName"`
	// Snapshot ID for the root volume (for EBS-backed AMIs)
	RootSnapshotId pulumi.StringOutput `pulumi:"rootSnapshotId"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrOutput `pulumi:"sriovNetSupport"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// If the image is configured for NitroTPM support, the value is `v2.0`. For more information, see [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html) in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport pulumi.StringPtrOutput `pulumi:"tpmSupport"`
	// Operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
	UsageOperation pulumi.StringOutput `pulumi:"usageOperation"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrOutput `pulumi:"virtualizationType"`
}

// NewAmi registers a new resource with the given unique name, arguments, and options.
func NewAmi(ctx *pulumi.Context,
	name string, args *AmiArgs, opts ...pulumi.ResourceOption) (*Ami, error) {
	if args == nil {
		args = &AmiArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ami
	err := ctx.RegisterResource("aws:ec2/ami:Ami", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAmi gets an existing Ami resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAmi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AmiState, opts ...pulumi.ResourceOption) (*Ami, error) {
	var resource Ami
	err := ctx.ReadResource("aws:ec2/ami:Ami", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ami resources.
type amiState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture *string `pulumi:"architecture"`
	// ARN of the AMI.
	Arn *string `pulumi:"arn"`
	// Boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
	BootMode *string `pulumi:"bootMode"`
	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime *string `pulumi:"deprecationTime"`
	// Longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []AmiEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport *bool `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// Hypervisor type of the image.
	Hypervisor *string `pulumi:"hypervisor"`
	// Path to an S3 object containing an image manifest, e.g., created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation *string `pulumi:"imageLocation"`
	// AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.
	ImageOwnerAlias *string `pulumi:"imageOwnerAlias"`
	// Type of image.
	ImageType *string `pulumi:"imageType"`
	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to `v2.0`. For more information, see [Configure instance metadata options for new instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration).
	ImdsSupport *string `pulumi:"imdsSupport"`
	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId           *string `pulumi:"kernelId"`
	ManageEbsSnapshots *bool   `pulumi:"manageEbsSnapshots"`
	// Region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// AWS account ID of the image owner.
	OwnerId *string `pulumi:"ownerId"`
	// This value is set to windows for Windows AMIs; otherwise, it is blank.
	Platform *string `pulumi:"platform"`
	// Platform details associated with the billing code of the AMI.
	PlatformDetails *string `pulumi:"platformDetails"`
	// Whether the image has public launch permissions.
	Public *bool `pulumi:"public"`
	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId *string `pulumi:"ramdiskId"`
	// Name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName *string `pulumi:"rootDeviceName"`
	// Snapshot ID for the root volume (for EBS-backed AMIs)
	RootSnapshotId *string `pulumi:"rootSnapshotId"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `pulumi:"sriovNetSupport"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// If the image is configured for NitroTPM support, the value is `v2.0`. For more information, see [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html) in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport *string `pulumi:"tpmSupport"`
	// Operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
	UsageOperation *string `pulumi:"usageOperation"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

type AmiState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrInput
	// ARN of the AMI.
	Arn pulumi.StringPtrInput
	// Boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
	BootMode pulumi.StringPtrInput
	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime pulumi.StringPtrInput
	// Longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiEbsBlockDeviceArrayInput
	// Whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiEphemeralBlockDeviceArrayInput
	// Hypervisor type of the image.
	Hypervisor pulumi.StringPtrInput
	// Path to an S3 object containing an image manifest, e.g., created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringPtrInput
	// AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.
	ImageOwnerAlias pulumi.StringPtrInput
	// Type of image.
	ImageType pulumi.StringPtrInput
	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to `v2.0`. For more information, see [Configure instance metadata options for new instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration).
	ImdsSupport pulumi.StringPtrInput
	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId           pulumi.StringPtrInput
	ManageEbsSnapshots pulumi.BoolPtrInput
	// Region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// AWS account ID of the image owner.
	OwnerId pulumi.StringPtrInput
	// This value is set to windows for Windows AMIs; otherwise, it is blank.
	Platform pulumi.StringPtrInput
	// Platform details associated with the billing code of the AMI.
	PlatformDetails pulumi.StringPtrInput
	// Whether the image has public launch permissions.
	Public pulumi.BoolPtrInput
	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrInput
	// Name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrInput
	// Snapshot ID for the root volume (for EBS-backed AMIs)
	RootSnapshotId pulumi.StringPtrInput
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// If the image is configured for NitroTPM support, the value is `v2.0`. For more information, see [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html) in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport pulumi.StringPtrInput
	// Operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
	UsageOperation pulumi.StringPtrInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrInput
}

func (AmiState) ElementType() reflect.Type {
	return reflect.TypeOf((*amiState)(nil)).Elem()
}

type amiArgs struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture *string `pulumi:"architecture"`
	// Boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
	BootMode *string `pulumi:"bootMode"`
	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime *string `pulumi:"deprecationTime"`
	// Longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []AmiEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport *bool `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g., created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation *string `pulumi:"imageLocation"`
	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to `v2.0`. For more information, see [Configure instance metadata options for new instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration).
	ImdsSupport *string `pulumi:"imdsSupport"`
	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId *string `pulumi:"kernelId"`
	// Region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId *string `pulumi:"ramdiskId"`
	// Name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName *string `pulumi:"rootDeviceName"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `pulumi:"sriovNetSupport"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// If the image is configured for NitroTPM support, the value is `v2.0`. For more information, see [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html) in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport *string `pulumi:"tpmSupport"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

// The set of arguments for constructing a Ami resource.
type AmiArgs struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrInput
	// Boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
	BootMode pulumi.StringPtrInput
	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime pulumi.StringPtrInput
	// Longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiEbsBlockDeviceArrayInput
	// Whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiEphemeralBlockDeviceArrayInput
	// Path to an S3 object containing an image manifest, e.g., created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringPtrInput
	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to `v2.0`. For more information, see [Configure instance metadata options for new instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration).
	ImdsSupport pulumi.StringPtrInput
	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringPtrInput
	// Region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrInput
	// Name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrInput
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// If the image is configured for NitroTPM support, the value is `v2.0`. For more information, see [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html) in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport pulumi.StringPtrInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrInput
}

func (AmiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*amiArgs)(nil)).Elem()
}

type AmiInput interface {
	pulumi.Input

	ToAmiOutput() AmiOutput
	ToAmiOutputWithContext(ctx context.Context) AmiOutput
}

func (*Ami) ElementType() reflect.Type {
	return reflect.TypeOf((**Ami)(nil)).Elem()
}

func (i *Ami) ToAmiOutput() AmiOutput {
	return i.ToAmiOutputWithContext(context.Background())
}

func (i *Ami) ToAmiOutputWithContext(ctx context.Context) AmiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiOutput)
}

// AmiArrayInput is an input type that accepts AmiArray and AmiArrayOutput values.
// You can construct a concrete instance of `AmiArrayInput` via:
//
//	AmiArray{ AmiArgs{...} }
type AmiArrayInput interface {
	pulumi.Input

	ToAmiArrayOutput() AmiArrayOutput
	ToAmiArrayOutputWithContext(context.Context) AmiArrayOutput
}

type AmiArray []AmiInput

func (AmiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ami)(nil)).Elem()
}

func (i AmiArray) ToAmiArrayOutput() AmiArrayOutput {
	return i.ToAmiArrayOutputWithContext(context.Background())
}

func (i AmiArray) ToAmiArrayOutputWithContext(ctx context.Context) AmiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiArrayOutput)
}

// AmiMapInput is an input type that accepts AmiMap and AmiMapOutput values.
// You can construct a concrete instance of `AmiMapInput` via:
//
//	AmiMap{ "key": AmiArgs{...} }
type AmiMapInput interface {
	pulumi.Input

	ToAmiMapOutput() AmiMapOutput
	ToAmiMapOutputWithContext(context.Context) AmiMapOutput
}

type AmiMap map[string]AmiInput

func (AmiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ami)(nil)).Elem()
}

func (i AmiMap) ToAmiMapOutput() AmiMapOutput {
	return i.ToAmiMapOutputWithContext(context.Background())
}

func (i AmiMap) ToAmiMapOutputWithContext(ctx context.Context) AmiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiMapOutput)
}

type AmiOutput struct{ *pulumi.OutputState }

func (AmiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ami)(nil)).Elem()
}

func (o AmiOutput) ToAmiOutput() AmiOutput {
	return o
}

func (o AmiOutput) ToAmiOutputWithContext(ctx context.Context) AmiOutput {
	return o
}

// Machine architecture for created instances. Defaults to "x8664".
func (o AmiOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.Architecture }).(pulumi.StringPtrOutput)
}

// ARN of the AMI.
func (o AmiOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
func (o AmiOutput) BootMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.BootMode }).(pulumi.StringPtrOutput)
}

// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
func (o AmiOutput) DeprecationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.DeprecationTime }).(pulumi.StringPtrOutput)
}

// Longer, human-readable description for the AMI.
func (o AmiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Nested block describing an EBS block device that should be
// attached to created instances. The structure of this block is described below.
func (o AmiOutput) EbsBlockDevices() AmiEbsBlockDeviceArrayOutput {
	return o.ApplyT(func(v *Ami) AmiEbsBlockDeviceArrayOutput { return v.EbsBlockDevices }).(AmiEbsBlockDeviceArrayOutput)
}

// Whether enhanced networking with ENA is enabled. Defaults to `false`.
func (o AmiOutput) EnaSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.BoolPtrOutput { return v.EnaSupport }).(pulumi.BoolPtrOutput)
}

// Nested block describing an ephemeral block device that
// should be attached to created instances. The structure of this block is described below.
func (o AmiOutput) EphemeralBlockDevices() AmiEphemeralBlockDeviceArrayOutput {
	return o.ApplyT(func(v *Ami) AmiEphemeralBlockDeviceArrayOutput { return v.EphemeralBlockDevices }).(AmiEphemeralBlockDeviceArrayOutput)
}

// Hypervisor type of the image.
func (o AmiOutput) Hypervisor() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.Hypervisor }).(pulumi.StringOutput)
}

// Path to an S3 object containing an image manifest, e.g., created
// by the `ec2-upload-bundle` command in the EC2 command line tools.
func (o AmiOutput) ImageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.ImageLocation }).(pulumi.StringOutput)
}

// AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.
func (o AmiOutput) ImageOwnerAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.ImageOwnerAlias }).(pulumi.StringOutput)
}

// Type of image.
func (o AmiOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.ImageType }).(pulumi.StringOutput)
}

// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to `v2.0`. For more information, see [Configure instance metadata options for new instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration).
func (o AmiOutput) ImdsSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.ImdsSupport }).(pulumi.StringPtrOutput)
}

// ID of the kernel image (AKI) that will be used as the paravirtual
// kernel in created instances.
func (o AmiOutput) KernelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.KernelId }).(pulumi.StringPtrOutput)
}

func (o AmiOutput) ManageEbsSnapshots() pulumi.BoolOutput {
	return o.ApplyT(func(v *Ami) pulumi.BoolOutput { return v.ManageEbsSnapshots }).(pulumi.BoolOutput)
}

// Region-unique name for the AMI.
func (o AmiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// AWS account ID of the image owner.
func (o AmiOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// This value is set to windows for Windows AMIs; otherwise, it is blank.
func (o AmiOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

// Platform details associated with the billing code of the AMI.
func (o AmiOutput) PlatformDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.PlatformDetails }).(pulumi.StringOutput)
}

// Whether the image has public launch permissions.
func (o AmiOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v *Ami) pulumi.BoolOutput { return v.Public }).(pulumi.BoolOutput)
}

// ID of an initrd image (ARI) that will be used when booting the
// created instances.
func (o AmiOutput) RamdiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.RamdiskId }).(pulumi.StringPtrOutput)
}

// Name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
func (o AmiOutput) RootDeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.RootDeviceName }).(pulumi.StringPtrOutput)
}

// Snapshot ID for the root volume (for EBS-backed AMIs)
func (o AmiOutput) RootSnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.RootSnapshotId }).(pulumi.StringOutput)
}

// When set to "simple" (the default), enables enhanced networking
// for created instances. No other value is supported at this time.
func (o AmiOutput) SriovNetSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.SriovNetSupport }).(pulumi.StringPtrOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o AmiOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o AmiOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// If the image is configured for NitroTPM support, the value is `v2.0`. For more information, see [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html) in the Amazon Elastic Compute Cloud User Guide.
func (o AmiOutput) TpmSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.TpmSupport }).(pulumi.StringPtrOutput)
}

// Operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
func (o AmiOutput) UsageOperation() pulumi.StringOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringOutput { return v.UsageOperation }).(pulumi.StringOutput)
}

// Keyword to choose what virtualization mode created instances
// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
// changes the set of further arguments that are required, as described below.
func (o AmiOutput) VirtualizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ami) pulumi.StringPtrOutput { return v.VirtualizationType }).(pulumi.StringPtrOutput)
}

type AmiArrayOutput struct{ *pulumi.OutputState }

func (AmiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ami)(nil)).Elem()
}

func (o AmiArrayOutput) ToAmiArrayOutput() AmiArrayOutput {
	return o
}

func (o AmiArrayOutput) ToAmiArrayOutputWithContext(ctx context.Context) AmiArrayOutput {
	return o
}

func (o AmiArrayOutput) Index(i pulumi.IntInput) AmiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ami {
		return vs[0].([]*Ami)[vs[1].(int)]
	}).(AmiOutput)
}

type AmiMapOutput struct{ *pulumi.OutputState }

func (AmiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ami)(nil)).Elem()
}

func (o AmiMapOutput) ToAmiMapOutput() AmiMapOutput {
	return o
}

func (o AmiMapOutput) ToAmiMapOutputWithContext(ctx context.Context) AmiMapOutput {
	return o
}

func (o AmiMapOutput) MapIndex(k pulumi.StringInput) AmiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ami {
		return vs[0].(map[string]*Ami)[vs[1].(string)]
	}).(AmiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AmiInput)(nil)).Elem(), &Ami{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmiArrayInput)(nil)).Elem(), AmiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmiMapInput)(nil)).Elem(), AmiMap{})
	pulumi.RegisterOutputType(AmiOutput{})
	pulumi.RegisterOutputType(AmiArrayOutput{})
	pulumi.RegisterOutputType(AmiMapOutput{})
}
