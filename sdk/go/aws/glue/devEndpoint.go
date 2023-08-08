// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue Development Endpoint resource.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
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
//									"glue.amazonaws.com",
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
//				AssumeRolePolicy: *pulumi.String(examplePolicyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = glue.NewDevEndpoint(ctx, "exampleDevEndpoint", &glue.DevEndpointArgs{
//				RoleArn: exampleRole.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewRolePolicyAttachment(ctx, "example-AWSGlueServiceRole", &iam.RolePolicyAttachmentArgs{
//				PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"),
//				Role:      exampleRole.Name,
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
//	to = aws_glue_dev_endpoint.example
//
//	id = "foo" } Using `pulumi import`, import a Glue Development Endpoint using the `name`. For exampleconsole % pulumi import aws_glue_dev_endpoint.example foo
type DevEndpoint struct {
	pulumi.CustomResourceState

	// A map of arguments used to configure the endpoint.
	Arguments pulumi.StringMapOutput `pulumi:"arguments"`
	// The ARN of the endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The AWS availability zone where this endpoint is located.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
	ExtraJarsS3Path pulumi.StringPtrOutput `pulumi:"extraJarsS3Path"`
	// Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
	ExtraPythonLibsS3Path pulumi.StringPtrOutput `pulumi:"extraPythonLibsS3Path"`
	// The reason for a current failure in this endpoint.
	FailureReason pulumi.StringOutput `pulumi:"failureReason"`
	// Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
	GlueVersion pulumi.StringPtrOutput `pulumi:"glueVersion"`
	// The name of this endpoint. It must be unique in your account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
	NumberOfNodes pulumi.IntPtrOutput `pulumi:"numberOfNodes"`
	// The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
	NumberOfWorkers pulumi.IntPtrOutput `pulumi:"numberOfWorkers"`
	// A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
	PrivateAddress pulumi.StringOutput `pulumi:"privateAddress"`
	// The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
	PublicAddress pulumi.StringOutput `pulumi:"publicAddress"`
	// The public key to be used by this endpoint for authentication.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// A list of public keys to be used by this endpoint for authentication.
	PublicKeys pulumi.StringArrayOutput `pulumi:"publicKeys"`
	// The IAM role for this endpoint.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The name of the Security Configuration structure to be used with this endpoint.
	SecurityConfiguration pulumi.StringPtrOutput `pulumi:"securityConfiguration"`
	// Security group IDs for the security groups to be used by this endpoint.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The current status of this endpoint.
	Status pulumi.StringOutput `pulumi:"status"`
	// The subnet ID for the new endpoint to use.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// he ID of the VPC used by this endpoint.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrOutput `pulumi:"workerType"`
	// The YARN endpoint address used by this endpoint.
	YarnEndpointAddress pulumi.StringOutput `pulumi:"yarnEndpointAddress"`
	// The Apache Zeppelin port for the remote Apache Spark interpreter.
	ZeppelinRemoteSparkInterpreterPort pulumi.IntOutput `pulumi:"zeppelinRemoteSparkInterpreterPort"`
}

// NewDevEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDevEndpoint(ctx *pulumi.Context,
	name string, args *DevEndpointArgs, opts ...pulumi.ResourceOption) (*DevEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DevEndpoint
	err := ctx.RegisterResource("aws:glue/devEndpoint:DevEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevEndpoint gets an existing DevEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevEndpointState, opts ...pulumi.ResourceOption) (*DevEndpoint, error) {
	var resource DevEndpoint
	err := ctx.ReadResource("aws:glue/devEndpoint:DevEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevEndpoint resources.
type devEndpointState struct {
	// A map of arguments used to configure the endpoint.
	Arguments map[string]string `pulumi:"arguments"`
	// The ARN of the endpoint.
	Arn *string `pulumi:"arn"`
	// The AWS availability zone where this endpoint is located.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
	ExtraJarsS3Path *string `pulumi:"extraJarsS3Path"`
	// Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
	ExtraPythonLibsS3Path *string `pulumi:"extraPythonLibsS3Path"`
	// The reason for a current failure in this endpoint.
	FailureReason *string `pulumi:"failureReason"`
	// Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
	GlueVersion *string `pulumi:"glueVersion"`
	// The name of this endpoint. It must be unique in your account.
	Name *string `pulumi:"name"`
	// The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
	NumberOfNodes *int `pulumi:"numberOfNodes"`
	// The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
	PrivateAddress *string `pulumi:"privateAddress"`
	// The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
	PublicAddress *string `pulumi:"publicAddress"`
	// The public key to be used by this endpoint for authentication.
	PublicKey *string `pulumi:"publicKey"`
	// A list of public keys to be used by this endpoint for authentication.
	PublicKeys []string `pulumi:"publicKeys"`
	// The IAM role for this endpoint.
	RoleArn *string `pulumi:"roleArn"`
	// The name of the Security Configuration structure to be used with this endpoint.
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// Security group IDs for the security groups to be used by this endpoint.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The current status of this endpoint.
	Status *string `pulumi:"status"`
	// The subnet ID for the new endpoint to use.
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// he ID of the VPC used by this endpoint.
	VpcId *string `pulumi:"vpcId"`
	// The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType *string `pulumi:"workerType"`
	// The YARN endpoint address used by this endpoint.
	YarnEndpointAddress *string `pulumi:"yarnEndpointAddress"`
	// The Apache Zeppelin port for the remote Apache Spark interpreter.
	ZeppelinRemoteSparkInterpreterPort *int `pulumi:"zeppelinRemoteSparkInterpreterPort"`
}

type DevEndpointState struct {
	// A map of arguments used to configure the endpoint.
	Arguments pulumi.StringMapInput
	// The ARN of the endpoint.
	Arn pulumi.StringPtrInput
	// The AWS availability zone where this endpoint is located.
	AvailabilityZone pulumi.StringPtrInput
	// Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
	ExtraJarsS3Path pulumi.StringPtrInput
	// Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
	ExtraPythonLibsS3Path pulumi.StringPtrInput
	// The reason for a current failure in this endpoint.
	FailureReason pulumi.StringPtrInput
	// Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
	GlueVersion pulumi.StringPtrInput
	// The name of this endpoint. It must be unique in your account.
	Name pulumi.StringPtrInput
	// The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
	NumberOfNodes pulumi.IntPtrInput
	// The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
	NumberOfWorkers pulumi.IntPtrInput
	// A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
	PrivateAddress pulumi.StringPtrInput
	// The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
	PublicAddress pulumi.StringPtrInput
	// The public key to be used by this endpoint for authentication.
	PublicKey pulumi.StringPtrInput
	// A list of public keys to be used by this endpoint for authentication.
	PublicKeys pulumi.StringArrayInput
	// The IAM role for this endpoint.
	RoleArn pulumi.StringPtrInput
	// The name of the Security Configuration structure to be used with this endpoint.
	SecurityConfiguration pulumi.StringPtrInput
	// Security group IDs for the security groups to be used by this endpoint.
	SecurityGroupIds pulumi.StringArrayInput
	// The current status of this endpoint.
	Status pulumi.StringPtrInput
	// The subnet ID for the new endpoint to use.
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// he ID of the VPC used by this endpoint.
	VpcId pulumi.StringPtrInput
	// The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrInput
	// The YARN endpoint address used by this endpoint.
	YarnEndpointAddress pulumi.StringPtrInput
	// The Apache Zeppelin port for the remote Apache Spark interpreter.
	ZeppelinRemoteSparkInterpreterPort pulumi.IntPtrInput
}

func (DevEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*devEndpointState)(nil)).Elem()
}

type devEndpointArgs struct {
	// A map of arguments used to configure the endpoint.
	Arguments map[string]string `pulumi:"arguments"`
	// Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
	ExtraJarsS3Path *string `pulumi:"extraJarsS3Path"`
	// Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
	ExtraPythonLibsS3Path *string `pulumi:"extraPythonLibsS3Path"`
	// Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
	GlueVersion *string `pulumi:"glueVersion"`
	// The name of this endpoint. It must be unique in your account.
	Name *string `pulumi:"name"`
	// The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
	NumberOfNodes *int `pulumi:"numberOfNodes"`
	// The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// The public key to be used by this endpoint for authentication.
	PublicKey *string `pulumi:"publicKey"`
	// A list of public keys to be used by this endpoint for authentication.
	PublicKeys []string `pulumi:"publicKeys"`
	// The IAM role for this endpoint.
	RoleArn string `pulumi:"roleArn"`
	// The name of the Security Configuration structure to be used with this endpoint.
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// Security group IDs for the security groups to be used by this endpoint.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The subnet ID for the new endpoint to use.
	SubnetId *string `pulumi:"subnetId"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType *string `pulumi:"workerType"`
}

// The set of arguments for constructing a DevEndpoint resource.
type DevEndpointArgs struct {
	// A map of arguments used to configure the endpoint.
	Arguments pulumi.StringMapInput
	// Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
	ExtraJarsS3Path pulumi.StringPtrInput
	// Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
	ExtraPythonLibsS3Path pulumi.StringPtrInput
	// Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
	GlueVersion pulumi.StringPtrInput
	// The name of this endpoint. It must be unique in your account.
	Name pulumi.StringPtrInput
	// The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
	NumberOfNodes pulumi.IntPtrInput
	// The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
	NumberOfWorkers pulumi.IntPtrInput
	// The public key to be used by this endpoint for authentication.
	PublicKey pulumi.StringPtrInput
	// A list of public keys to be used by this endpoint for authentication.
	PublicKeys pulumi.StringArrayInput
	// The IAM role for this endpoint.
	RoleArn pulumi.StringInput
	// The name of the Security Configuration structure to be used with this endpoint.
	SecurityConfiguration pulumi.StringPtrInput
	// Security group IDs for the security groups to be used by this endpoint.
	SecurityGroupIds pulumi.StringArrayInput
	// The subnet ID for the new endpoint to use.
	SubnetId pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrInput
}

func (DevEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devEndpointArgs)(nil)).Elem()
}

type DevEndpointInput interface {
	pulumi.Input

	ToDevEndpointOutput() DevEndpointOutput
	ToDevEndpointOutputWithContext(ctx context.Context) DevEndpointOutput
}

func (*DevEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**DevEndpoint)(nil)).Elem()
}

func (i *DevEndpoint) ToDevEndpointOutput() DevEndpointOutput {
	return i.ToDevEndpointOutputWithContext(context.Background())
}

func (i *DevEndpoint) ToDevEndpointOutputWithContext(ctx context.Context) DevEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevEndpointOutput)
}

// DevEndpointArrayInput is an input type that accepts DevEndpointArray and DevEndpointArrayOutput values.
// You can construct a concrete instance of `DevEndpointArrayInput` via:
//
//	DevEndpointArray{ DevEndpointArgs{...} }
type DevEndpointArrayInput interface {
	pulumi.Input

	ToDevEndpointArrayOutput() DevEndpointArrayOutput
	ToDevEndpointArrayOutputWithContext(context.Context) DevEndpointArrayOutput
}

type DevEndpointArray []DevEndpointInput

func (DevEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DevEndpoint)(nil)).Elem()
}

func (i DevEndpointArray) ToDevEndpointArrayOutput() DevEndpointArrayOutput {
	return i.ToDevEndpointArrayOutputWithContext(context.Background())
}

func (i DevEndpointArray) ToDevEndpointArrayOutputWithContext(ctx context.Context) DevEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevEndpointArrayOutput)
}

// DevEndpointMapInput is an input type that accepts DevEndpointMap and DevEndpointMapOutput values.
// You can construct a concrete instance of `DevEndpointMapInput` via:
//
//	DevEndpointMap{ "key": DevEndpointArgs{...} }
type DevEndpointMapInput interface {
	pulumi.Input

	ToDevEndpointMapOutput() DevEndpointMapOutput
	ToDevEndpointMapOutputWithContext(context.Context) DevEndpointMapOutput
}

type DevEndpointMap map[string]DevEndpointInput

func (DevEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DevEndpoint)(nil)).Elem()
}

func (i DevEndpointMap) ToDevEndpointMapOutput() DevEndpointMapOutput {
	return i.ToDevEndpointMapOutputWithContext(context.Background())
}

func (i DevEndpointMap) ToDevEndpointMapOutputWithContext(ctx context.Context) DevEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevEndpointMapOutput)
}

type DevEndpointOutput struct{ *pulumi.OutputState }

func (DevEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevEndpoint)(nil)).Elem()
}

func (o DevEndpointOutput) ToDevEndpointOutput() DevEndpointOutput {
	return o
}

func (o DevEndpointOutput) ToDevEndpointOutputWithContext(ctx context.Context) DevEndpointOutput {
	return o
}

// A map of arguments used to configure the endpoint.
func (o DevEndpointOutput) Arguments() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringMapOutput { return v.Arguments }).(pulumi.StringMapOutput)
}

// The ARN of the endpoint.
func (o DevEndpointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The AWS availability zone where this endpoint is located.
func (o DevEndpointOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
func (o DevEndpointOutput) ExtraJarsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.ExtraJarsS3Path }).(pulumi.StringPtrOutput)
}

// Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
func (o DevEndpointOutput) ExtraPythonLibsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.ExtraPythonLibsS3Path }).(pulumi.StringPtrOutput)
}

// The reason for a current failure in this endpoint.
func (o DevEndpointOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

// Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
func (o DevEndpointOutput) GlueVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.GlueVersion }).(pulumi.StringPtrOutput)
}

// The name of this endpoint. It must be unique in your account.
func (o DevEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
func (o DevEndpointOutput) NumberOfNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.IntPtrOutput { return v.NumberOfNodes }).(pulumi.IntPtrOutput)
}

// The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
func (o DevEndpointOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.IntPtrOutput { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

// A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
func (o DevEndpointOutput) PrivateAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.PrivateAddress }).(pulumi.StringOutput)
}

// The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
func (o DevEndpointOutput) PublicAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.PublicAddress }).(pulumi.StringOutput)
}

// The public key to be used by this endpoint for authentication.
func (o DevEndpointOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// A list of public keys to be used by this endpoint for authentication.
func (o DevEndpointOutput) PublicKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringArrayOutput { return v.PublicKeys }).(pulumi.StringArrayOutput)
}

// The IAM role for this endpoint.
func (o DevEndpointOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The name of the Security Configuration structure to be used with this endpoint.
func (o DevEndpointOutput) SecurityConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.SecurityConfiguration }).(pulumi.StringPtrOutput)
}

// Security group IDs for the security groups to be used by this endpoint.
func (o DevEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The current status of this endpoint.
func (o DevEndpointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The subnet ID for the new endpoint to use.
func (o DevEndpointOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DevEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o DevEndpointOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// he ID of the VPC used by this endpoint.
func (o DevEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
func (o DevEndpointOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.WorkerType }).(pulumi.StringPtrOutput)
}

// The YARN endpoint address used by this endpoint.
func (o DevEndpointOutput) YarnEndpointAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.YarnEndpointAddress }).(pulumi.StringOutput)
}

// The Apache Zeppelin port for the remote Apache Spark interpreter.
func (o DevEndpointOutput) ZeppelinRemoteSparkInterpreterPort() pulumi.IntOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.IntOutput { return v.ZeppelinRemoteSparkInterpreterPort }).(pulumi.IntOutput)
}

type DevEndpointArrayOutput struct{ *pulumi.OutputState }

func (DevEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DevEndpoint)(nil)).Elem()
}

func (o DevEndpointArrayOutput) ToDevEndpointArrayOutput() DevEndpointArrayOutput {
	return o
}

func (o DevEndpointArrayOutput) ToDevEndpointArrayOutputWithContext(ctx context.Context) DevEndpointArrayOutput {
	return o
}

func (o DevEndpointArrayOutput) Index(i pulumi.IntInput) DevEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DevEndpoint {
		return vs[0].([]*DevEndpoint)[vs[1].(int)]
	}).(DevEndpointOutput)
}

type DevEndpointMapOutput struct{ *pulumi.OutputState }

func (DevEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DevEndpoint)(nil)).Elem()
}

func (o DevEndpointMapOutput) ToDevEndpointMapOutput() DevEndpointMapOutput {
	return o
}

func (o DevEndpointMapOutput) ToDevEndpointMapOutputWithContext(ctx context.Context) DevEndpointMapOutput {
	return o
}

func (o DevEndpointMapOutput) MapIndex(k pulumi.StringInput) DevEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DevEndpoint {
		return vs[0].(map[string]*DevEndpoint)[vs[1].(string)]
	}).(DevEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevEndpointInput)(nil)).Elem(), &DevEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevEndpointArrayInput)(nil)).Elem(), DevEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevEndpointMapInput)(nil)).Elem(), DevEndpointMap{})
	pulumi.RegisterOutputType(DevEndpointOutput{})
	pulumi.RegisterOutputType(DevEndpointArrayOutput{})
	pulumi.RegisterOutputType(DevEndpointMapOutput{})
}
