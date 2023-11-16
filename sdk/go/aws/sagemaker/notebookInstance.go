// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SageMaker Notebook Instance resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sagemaker.NewNotebookInstance(ctx, "ni", &sagemaker.NotebookInstanceArgs{
//				RoleArn:      pulumi.Any(aws_iam_role.Role.Arn),
//				InstanceType: pulumi.String("ml.t2.medium"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("foo"),
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
// ### Code repository usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := sagemaker.NewCodeRepository(ctx, "example", &sagemaker.CodeRepositoryArgs{
//				CodeRepositoryName: pulumi.String("my-notebook-instance-code-repo"),
//				GitConfig: &sagemaker.CodeRepositoryGitConfigArgs{
//					RepositoryUrl: pulumi.String("https://github.com/github/docs.git"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sagemaker.NewNotebookInstance(ctx, "ni", &sagemaker.NotebookInstanceArgs{
//				RoleArn:               pulumi.Any(aws_iam_role.Role.Arn),
//				InstanceType:          pulumi.String("ml.t2.medium"),
//				DefaultCodeRepository: example.CodeRepositoryName,
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("foo"),
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
// Using `pulumi import`, import SageMaker Notebook Instances using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:sagemaker/notebookInstance:NotebookInstance test_notebook_instance my-notebook-instance
//
// ```
type NotebookInstance struct {
	pulumi.CustomResourceState

	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See [Elastic Inference Accelerator](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) for more details. Valid values: `ml.eia1.medium`, `ml.eia1.large`, `ml.eia1.xlarge`, `ml.eia2.medium`, `ml.eia2.large`, `ml.eia2.xlarge`.
	AcceleratorTypes pulumi.StringArrayOutput `pulumi:"acceleratorTypes"`
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories pulumi.StringArrayOutput `pulumi:"additionalCodeRepositories"`
	// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository pulumi.StringPtrOutput `pulumi:"defaultCodeRepository"`
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess pulumi.StringPtrOutput `pulumi:"directInternetAccess"`
	// Information on the IMDS configuration of the notebook instance. Conflicts with `instanceMetadataServiceConfiguration`. see details below.
	InstanceMetadataServiceConfiguration NotebookInstanceInstanceMetadataServiceConfigurationPtrOutput `pulumi:"instanceMetadataServiceConfiguration"`
	// The name of ML compute instance type.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName pulumi.StringPtrOutput `pulumi:"lifecycleConfigName"`
	// The name of the notebook instance (must be unique).
	Name pulumi.StringOutput `pulumi:"name"`
	// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnetId`.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1`, `notebook-al2-v1`, or  `notebook-al2-v2`, depending on which version of Amazon Linux you require.
	PlatformIdentifier pulumi.StringOutput `pulumi:"platformIdentifier"`
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess pulumi.StringPtrOutput `pulumi:"rootAccess"`
	// The associated security groups.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// The VPC subnet ID.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
	Url pulumi.StringOutput `pulumi:"url"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize pulumi.IntPtrOutput `pulumi:"volumeSize"`
}

// NewNotebookInstance registers a new resource with the given unique name, arguments, and options.
func NewNotebookInstance(ctx *pulumi.Context,
	name string, args *NotebookInstanceArgs, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NotebookInstance
	err := ctx.RegisterResource("aws:sagemaker/notebookInstance:NotebookInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookInstance gets an existing NotebookInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookInstanceState, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	var resource NotebookInstance
	err := ctx.ReadResource("aws:sagemaker/notebookInstance:NotebookInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookInstance resources.
type notebookInstanceState struct {
	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See [Elastic Inference Accelerator](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) for more details. Valid values: `ml.eia1.medium`, `ml.eia1.large`, `ml.eia1.xlarge`, `ml.eia2.medium`, `ml.eia2.large`, `ml.eia2.xlarge`.
	AcceleratorTypes []string `pulumi:"acceleratorTypes"`
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories []string `pulumi:"additionalCodeRepositories"`
	// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
	Arn *string `pulumi:"arn"`
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository *string `pulumi:"defaultCodeRepository"`
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess *string `pulumi:"directInternetAccess"`
	// Information on the IMDS configuration of the notebook instance. Conflicts with `instanceMetadataServiceConfiguration`. see details below.
	InstanceMetadataServiceConfiguration *NotebookInstanceInstanceMetadataServiceConfiguration `pulumi:"instanceMetadataServiceConfiguration"`
	// The name of ML compute instance type.
	InstanceType *string `pulumi:"instanceType"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName *string `pulumi:"lifecycleConfigName"`
	// The name of the notebook instance (must be unique).
	Name *string `pulumi:"name"`
	// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnetId`.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1`, `notebook-al2-v1`, or  `notebook-al2-v2`, depending on which version of Amazon Linux you require.
	PlatformIdentifier *string `pulumi:"platformIdentifier"`
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess *string `pulumi:"rootAccess"`
	// The associated security groups.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The VPC subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
	Url *string `pulumi:"url"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize *int `pulumi:"volumeSize"`
}

type NotebookInstanceState struct {
	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See [Elastic Inference Accelerator](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) for more details. Valid values: `ml.eia1.medium`, `ml.eia1.large`, `ml.eia1.xlarge`, `ml.eia2.medium`, `ml.eia2.large`, `ml.eia2.xlarge`.
	AcceleratorTypes pulumi.StringArrayInput
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
	Arn pulumi.StringPtrInput
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository pulumi.StringPtrInput
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess pulumi.StringPtrInput
	// Information on the IMDS configuration of the notebook instance. Conflicts with `instanceMetadataServiceConfiguration`. see details below.
	InstanceMetadataServiceConfiguration NotebookInstanceInstanceMetadataServiceConfigurationPtrInput
	// The name of ML compute instance type.
	InstanceType pulumi.StringPtrInput
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId pulumi.StringPtrInput
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName pulumi.StringPtrInput
	// The name of the notebook instance (must be unique).
	Name pulumi.StringPtrInput
	// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnetId`.
	NetworkInterfaceId pulumi.StringPtrInput
	// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1`, `notebook-al2-v1`, or  `notebook-al2-v2`, depending on which version of Amazon Linux you require.
	PlatformIdentifier pulumi.StringPtrInput
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn pulumi.StringPtrInput
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess pulumi.StringPtrInput
	// The associated security groups.
	SecurityGroups pulumi.StringArrayInput
	// The VPC subnet ID.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
	Url pulumi.StringPtrInput
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize pulumi.IntPtrInput
}

func (NotebookInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceState)(nil)).Elem()
}

type notebookInstanceArgs struct {
	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See [Elastic Inference Accelerator](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) for more details. Valid values: `ml.eia1.medium`, `ml.eia1.large`, `ml.eia1.xlarge`, `ml.eia2.medium`, `ml.eia2.large`, `ml.eia2.xlarge`.
	AcceleratorTypes []string `pulumi:"acceleratorTypes"`
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories []string `pulumi:"additionalCodeRepositories"`
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository *string `pulumi:"defaultCodeRepository"`
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess *string `pulumi:"directInternetAccess"`
	// Information on the IMDS configuration of the notebook instance. Conflicts with `instanceMetadataServiceConfiguration`. see details below.
	InstanceMetadataServiceConfiguration *NotebookInstanceInstanceMetadataServiceConfiguration `pulumi:"instanceMetadataServiceConfiguration"`
	// The name of ML compute instance type.
	InstanceType string `pulumi:"instanceType"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName *string `pulumi:"lifecycleConfigName"`
	// The name of the notebook instance (must be unique).
	Name *string `pulumi:"name"`
	// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1`, `notebook-al2-v1`, or  `notebook-al2-v2`, depending on which version of Amazon Linux you require.
	PlatformIdentifier *string `pulumi:"platformIdentifier"`
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess *string `pulumi:"rootAccess"`
	// The associated security groups.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The VPC subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize *int `pulumi:"volumeSize"`
}

// The set of arguments for constructing a NotebookInstance resource.
type NotebookInstanceArgs struct {
	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See [Elastic Inference Accelerator](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) for more details. Valid values: `ml.eia1.medium`, `ml.eia1.large`, `ml.eia1.xlarge`, `ml.eia2.medium`, `ml.eia2.large`, `ml.eia2.xlarge`.
	AcceleratorTypes pulumi.StringArrayInput
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	AdditionalCodeRepositories pulumi.StringArrayInput
	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository pulumi.StringPtrInput
	// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess pulumi.StringPtrInput
	// Information on the IMDS configuration of the notebook instance. Conflicts with `instanceMetadataServiceConfiguration`. see details below.
	InstanceMetadataServiceConfiguration NotebookInstanceInstanceMetadataServiceConfigurationPtrInput
	// The name of ML compute instance type.
	InstanceType pulumi.StringInput
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId pulumi.StringPtrInput
	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName pulumi.StringPtrInput
	// The name of the notebook instance (must be unique).
	Name pulumi.StringPtrInput
	// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1`, `notebook-al2-v1`, or  `notebook-al2-v2`, depending on which version of Amazon Linux you require.
	PlatformIdentifier pulumi.StringPtrInput
	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn pulumi.StringInput
	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess pulumi.StringPtrInput
	// The associated security groups.
	SecurityGroups pulumi.StringArrayInput
	// The VPC subnet ID.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize pulumi.IntPtrInput
}

func (NotebookInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceArgs)(nil)).Elem()
}

type NotebookInstanceInput interface {
	pulumi.Input

	ToNotebookInstanceOutput() NotebookInstanceOutput
	ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput
}

func (*NotebookInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookInstance)(nil)).Elem()
}

func (i *NotebookInstance) ToNotebookInstanceOutput() NotebookInstanceOutput {
	return i.ToNotebookInstanceOutputWithContext(context.Background())
}

func (i *NotebookInstance) ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceOutput)
}

// NotebookInstanceArrayInput is an input type that accepts NotebookInstanceArray and NotebookInstanceArrayOutput values.
// You can construct a concrete instance of `NotebookInstanceArrayInput` via:
//
//	NotebookInstanceArray{ NotebookInstanceArgs{...} }
type NotebookInstanceArrayInput interface {
	pulumi.Input

	ToNotebookInstanceArrayOutput() NotebookInstanceArrayOutput
	ToNotebookInstanceArrayOutputWithContext(context.Context) NotebookInstanceArrayOutput
}

type NotebookInstanceArray []NotebookInstanceInput

func (NotebookInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotebookInstance)(nil)).Elem()
}

func (i NotebookInstanceArray) ToNotebookInstanceArrayOutput() NotebookInstanceArrayOutput {
	return i.ToNotebookInstanceArrayOutputWithContext(context.Background())
}

func (i NotebookInstanceArray) ToNotebookInstanceArrayOutputWithContext(ctx context.Context) NotebookInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceArrayOutput)
}

// NotebookInstanceMapInput is an input type that accepts NotebookInstanceMap and NotebookInstanceMapOutput values.
// You can construct a concrete instance of `NotebookInstanceMapInput` via:
//
//	NotebookInstanceMap{ "key": NotebookInstanceArgs{...} }
type NotebookInstanceMapInput interface {
	pulumi.Input

	ToNotebookInstanceMapOutput() NotebookInstanceMapOutput
	ToNotebookInstanceMapOutputWithContext(context.Context) NotebookInstanceMapOutput
}

type NotebookInstanceMap map[string]NotebookInstanceInput

func (NotebookInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotebookInstance)(nil)).Elem()
}

func (i NotebookInstanceMap) ToNotebookInstanceMapOutput() NotebookInstanceMapOutput {
	return i.ToNotebookInstanceMapOutputWithContext(context.Background())
}

func (i NotebookInstanceMap) ToNotebookInstanceMapOutputWithContext(ctx context.Context) NotebookInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceMapOutput)
}

type NotebookInstanceOutput struct{ *pulumi.OutputState }

func (NotebookInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookInstance)(nil)).Elem()
}

func (o NotebookInstanceOutput) ToNotebookInstanceOutput() NotebookInstanceOutput {
	return o
}

func (o NotebookInstanceOutput) ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput {
	return o
}

// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See [Elastic Inference Accelerator](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) for more details. Valid values: `ml.eia1.medium`, `ml.eia1.large`, `ml.eia1.xlarge`, `ml.eia2.medium`, `ml.eia2.large`, `ml.eia2.xlarge`.
func (o NotebookInstanceOutput) AcceleratorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringArrayOutput { return v.AcceleratorTypes }).(pulumi.StringArrayOutput)
}

// An array of up to three Git repositories to associate with the notebook instance.
// These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
func (o NotebookInstanceOutput) AdditionalCodeRepositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringArrayOutput { return v.AdditionalCodeRepositories }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) assigned by AWS to this notebook instance.
func (o NotebookInstanceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
func (o NotebookInstanceOutput) DefaultCodeRepository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringPtrOutput { return v.DefaultCodeRepository }).(pulumi.StringPtrOutput)
}

// Set to `Disabled` to disable internet access to notebook. Requires `securityGroups` and `subnetId` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
func (o NotebookInstanceOutput) DirectInternetAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringPtrOutput { return v.DirectInternetAccess }).(pulumi.StringPtrOutput)
}

// Information on the IMDS configuration of the notebook instance. Conflicts with `instanceMetadataServiceConfiguration`. see details below.
func (o NotebookInstanceOutput) InstanceMetadataServiceConfiguration() NotebookInstanceInstanceMetadataServiceConfigurationPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) NotebookInstanceInstanceMetadataServiceConfigurationPtrOutput {
		return v.InstanceMetadataServiceConfiguration
	}).(NotebookInstanceInstanceMetadataServiceConfigurationPtrOutput)
}

// The name of ML compute instance type.
func (o NotebookInstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
func (o NotebookInstanceOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The name of a lifecycle configuration to associate with the notebook instance.
func (o NotebookInstanceOutput) LifecycleConfigName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringPtrOutput { return v.LifecycleConfigName }).(pulumi.StringPtrOutput)
}

// The name of the notebook instance (must be unique).
func (o NotebookInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network interface ID that Amazon SageMaker created at the time of creating the instance. Only available when setting `subnetId`.
func (o NotebookInstanceOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1`, `notebook-al2-v1`, or  `notebook-al2-v2`, depending on which version of Amazon Linux you require.
func (o NotebookInstanceOutput) PlatformIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringOutput { return v.PlatformIdentifier }).(pulumi.StringOutput)
}

// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
func (o NotebookInstanceOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
func (o NotebookInstanceOutput) RootAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringPtrOutput { return v.RootAccess }).(pulumi.StringPtrOutput)
}

// The associated security groups.
func (o NotebookInstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// The VPC subnet ID.
func (o NotebookInstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NotebookInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o NotebookInstanceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The URL that you use to connect to the Jupyter notebook that is running in your notebook instance.
func (o NotebookInstanceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
func (o NotebookInstanceOutput) VolumeSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NotebookInstance) pulumi.IntPtrOutput { return v.VolumeSize }).(pulumi.IntPtrOutput)
}

type NotebookInstanceArrayOutput struct{ *pulumi.OutputState }

func (NotebookInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotebookInstance)(nil)).Elem()
}

func (o NotebookInstanceArrayOutput) ToNotebookInstanceArrayOutput() NotebookInstanceArrayOutput {
	return o
}

func (o NotebookInstanceArrayOutput) ToNotebookInstanceArrayOutputWithContext(ctx context.Context) NotebookInstanceArrayOutput {
	return o
}

func (o NotebookInstanceArrayOutput) Index(i pulumi.IntInput) NotebookInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotebookInstance {
		return vs[0].([]*NotebookInstance)[vs[1].(int)]
	}).(NotebookInstanceOutput)
}

type NotebookInstanceMapOutput struct{ *pulumi.OutputState }

func (NotebookInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotebookInstance)(nil)).Elem()
}

func (o NotebookInstanceMapOutput) ToNotebookInstanceMapOutput() NotebookInstanceMapOutput {
	return o
}

func (o NotebookInstanceMapOutput) ToNotebookInstanceMapOutputWithContext(ctx context.Context) NotebookInstanceMapOutput {
	return o
}

func (o NotebookInstanceMapOutput) MapIndex(k pulumi.StringInput) NotebookInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotebookInstance {
		return vs[0].(map[string]*NotebookInstance)[vs[1].(string)]
	}).(NotebookInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotebookInstanceInput)(nil)).Elem(), &NotebookInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotebookInstanceArrayInput)(nil)).Elem(), NotebookInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotebookInstanceMapInput)(nil)).Elem(), NotebookInstanceMap{})
	pulumi.RegisterOutputType(NotebookInstanceOutput{})
	pulumi.RegisterOutputType(NotebookInstanceArrayOutput{})
	pulumi.RegisterOutputType(NotebookInstanceMapOutput{})
}
