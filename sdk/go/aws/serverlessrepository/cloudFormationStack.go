// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serverlessrepository

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Deploys an Application CloudFormation Stack from the Serverless Application Repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/serverlessrepository"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			currentPartition, err := aws.GetPartition(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			currentRegion, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = serverlessrepository.NewCloudFormationStack(ctx, "postgres-rotator", &serverlessrepository.CloudFormationStackArgs{
//				ApplicationId: pulumi.String("arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser"),
//				Capabilities: pulumi.StringArray{
//					pulumi.String("CAPABILITY_IAM"),
//					pulumi.String("CAPABILITY_RESOURCE_POLICY"),
//				},
//				Parameters: pulumi.StringMap{
//					"endpoint":     pulumi.String(fmt.Sprintf("secretsmanager.%v.%v", currentRegion.Name, currentPartition.DnsSuffix)),
//					"functionName": pulumi.String("func-postgres-rotator"),
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
// Using `pulumi import`, import Serverless Application Repository Stack using the CloudFormation Stack name (with or without the `serverlessrepo-` prefix) or the CloudFormation Stack ID. For example:
//
// ```sh
//
//	$ pulumi import aws:serverlessrepository/cloudFormationStack:CloudFormationStack example serverlessrepo-postgres-rotator
//
// ```
type CloudFormationStack struct {
	pulumi.CustomResourceState

	// The ARN of the application from the Serverless Application Repository.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of outputs from the stack.
	Outputs pulumi.StringMapOutput `pulumi:"outputs"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion pulumi.StringOutput `pulumi:"semanticVersion"`
	// A list of tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewCloudFormationStack registers a new resource with the given unique name, arguments, and options.
func NewCloudFormationStack(ctx *pulumi.Context,
	name string, args *CloudFormationStackArgs, opts ...pulumi.ResourceOption) (*CloudFormationStack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.Capabilities == nil {
		return nil, errors.New("invalid value for required argument 'Capabilities'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudFormationStack
	err := ctx.RegisterResource("aws:serverlessrepository/cloudFormationStack:CloudFormationStack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudFormationStack gets an existing CloudFormationStack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudFormationStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudFormationStackState, opts ...pulumi.ResourceOption) (*CloudFormationStack, error) {
	var resource CloudFormationStack
	err := ctx.ReadResource("aws:serverlessrepository/cloudFormationStack:CloudFormationStack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudFormationStack resources.
type cloudFormationStackState struct {
	// The ARN of the application from the Serverless Application Repository.
	ApplicationId *string `pulumi:"applicationId"`
	// A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities []string `pulumi:"capabilities"`
	// The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
	Name *string `pulumi:"name"`
	// A map of outputs from the stack.
	Outputs map[string]string `pulumi:"outputs"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `pulumi:"parameters"`
	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion *string `pulumi:"semanticVersion"`
	// A list of tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type CloudFormationStackState struct {
	// The ARN of the application from the Serverless Application Repository.
	ApplicationId pulumi.StringPtrInput
	// A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput
	// The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
	Name pulumi.StringPtrInput
	// A map of outputs from the stack.
	Outputs pulumi.StringMapInput
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapInput
	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion pulumi.StringPtrInput
	// A list of tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (CloudFormationStackState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudFormationStackState)(nil)).Elem()
}

type cloudFormationStackArgs struct {
	// The ARN of the application from the Serverless Application Repository.
	ApplicationId string `pulumi:"applicationId"`
	// A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities []string `pulumi:"capabilities"`
	// The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
	Name *string `pulumi:"name"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `pulumi:"parameters"`
	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion *string `pulumi:"semanticVersion"`
	// A list of tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CloudFormationStack resource.
type CloudFormationStackArgs struct {
	// The ARN of the application from the Serverless Application Repository.
	ApplicationId pulumi.StringInput
	// A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput
	// The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
	Name pulumi.StringPtrInput
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapInput
	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion pulumi.StringPtrInput
	// A list of tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (CloudFormationStackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudFormationStackArgs)(nil)).Elem()
}

type CloudFormationStackInput interface {
	pulumi.Input

	ToCloudFormationStackOutput() CloudFormationStackOutput
	ToCloudFormationStackOutputWithContext(ctx context.Context) CloudFormationStackOutput
}

func (*CloudFormationStack) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudFormationStack)(nil)).Elem()
}

func (i *CloudFormationStack) ToCloudFormationStackOutput() CloudFormationStackOutput {
	return i.ToCloudFormationStackOutputWithContext(context.Background())
}

func (i *CloudFormationStack) ToCloudFormationStackOutputWithContext(ctx context.Context) CloudFormationStackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationStackOutput)
}

func (i *CloudFormationStack) ToOutput(ctx context.Context) pulumix.Output[*CloudFormationStack] {
	return pulumix.Output[*CloudFormationStack]{
		OutputState: i.ToCloudFormationStackOutputWithContext(ctx).OutputState,
	}
}

// CloudFormationStackArrayInput is an input type that accepts CloudFormationStackArray and CloudFormationStackArrayOutput values.
// You can construct a concrete instance of `CloudFormationStackArrayInput` via:
//
//	CloudFormationStackArray{ CloudFormationStackArgs{...} }
type CloudFormationStackArrayInput interface {
	pulumi.Input

	ToCloudFormationStackArrayOutput() CloudFormationStackArrayOutput
	ToCloudFormationStackArrayOutputWithContext(context.Context) CloudFormationStackArrayOutput
}

type CloudFormationStackArray []CloudFormationStackInput

func (CloudFormationStackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudFormationStack)(nil)).Elem()
}

func (i CloudFormationStackArray) ToCloudFormationStackArrayOutput() CloudFormationStackArrayOutput {
	return i.ToCloudFormationStackArrayOutputWithContext(context.Background())
}

func (i CloudFormationStackArray) ToCloudFormationStackArrayOutputWithContext(ctx context.Context) CloudFormationStackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationStackArrayOutput)
}

func (i CloudFormationStackArray) ToOutput(ctx context.Context) pulumix.Output[[]*CloudFormationStack] {
	return pulumix.Output[[]*CloudFormationStack]{
		OutputState: i.ToCloudFormationStackArrayOutputWithContext(ctx).OutputState,
	}
}

// CloudFormationStackMapInput is an input type that accepts CloudFormationStackMap and CloudFormationStackMapOutput values.
// You can construct a concrete instance of `CloudFormationStackMapInput` via:
//
//	CloudFormationStackMap{ "key": CloudFormationStackArgs{...} }
type CloudFormationStackMapInput interface {
	pulumi.Input

	ToCloudFormationStackMapOutput() CloudFormationStackMapOutput
	ToCloudFormationStackMapOutputWithContext(context.Context) CloudFormationStackMapOutput
}

type CloudFormationStackMap map[string]CloudFormationStackInput

func (CloudFormationStackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudFormationStack)(nil)).Elem()
}

func (i CloudFormationStackMap) ToCloudFormationStackMapOutput() CloudFormationStackMapOutput {
	return i.ToCloudFormationStackMapOutputWithContext(context.Background())
}

func (i CloudFormationStackMap) ToCloudFormationStackMapOutputWithContext(ctx context.Context) CloudFormationStackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationStackMapOutput)
}

func (i CloudFormationStackMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CloudFormationStack] {
	return pulumix.Output[map[string]*CloudFormationStack]{
		OutputState: i.ToCloudFormationStackMapOutputWithContext(ctx).OutputState,
	}
}

type CloudFormationStackOutput struct{ *pulumi.OutputState }

func (CloudFormationStackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudFormationStack)(nil)).Elem()
}

func (o CloudFormationStackOutput) ToCloudFormationStackOutput() CloudFormationStackOutput {
	return o
}

func (o CloudFormationStackOutput) ToCloudFormationStackOutputWithContext(ctx context.Context) CloudFormationStackOutput {
	return o
}

func (o CloudFormationStackOutput) ToOutput(ctx context.Context) pulumix.Output[*CloudFormationStack] {
	return pulumix.Output[*CloudFormationStack]{
		OutputState: o.OutputState,
	}
}

// The ARN of the application from the Serverless Application Repository.
func (o CloudFormationStackOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
func (o CloudFormationStackOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringArrayOutput { return v.Capabilities }).(pulumi.StringArrayOutput)
}

// The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
func (o CloudFormationStackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A map of outputs from the stack.
func (o CloudFormationStackOutput) Outputs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringMapOutput { return v.Outputs }).(pulumi.StringMapOutput)
}

// A map of Parameter structures that specify input parameters for the stack.
func (o CloudFormationStackOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// The version of the application to deploy. If not supplied, deploys the latest version.
func (o CloudFormationStackOutput) SemanticVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringOutput { return v.SemanticVersion }).(pulumi.StringOutput)
}

// A list of tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CloudFormationStackOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o CloudFormationStackOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudFormationStack) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type CloudFormationStackArrayOutput struct{ *pulumi.OutputState }

func (CloudFormationStackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudFormationStack)(nil)).Elem()
}

func (o CloudFormationStackArrayOutput) ToCloudFormationStackArrayOutput() CloudFormationStackArrayOutput {
	return o
}

func (o CloudFormationStackArrayOutput) ToCloudFormationStackArrayOutputWithContext(ctx context.Context) CloudFormationStackArrayOutput {
	return o
}

func (o CloudFormationStackArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CloudFormationStack] {
	return pulumix.Output[[]*CloudFormationStack]{
		OutputState: o.OutputState,
	}
}

func (o CloudFormationStackArrayOutput) Index(i pulumi.IntInput) CloudFormationStackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudFormationStack {
		return vs[0].([]*CloudFormationStack)[vs[1].(int)]
	}).(CloudFormationStackOutput)
}

type CloudFormationStackMapOutput struct{ *pulumi.OutputState }

func (CloudFormationStackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudFormationStack)(nil)).Elem()
}

func (o CloudFormationStackMapOutput) ToCloudFormationStackMapOutput() CloudFormationStackMapOutput {
	return o
}

func (o CloudFormationStackMapOutput) ToCloudFormationStackMapOutputWithContext(ctx context.Context) CloudFormationStackMapOutput {
	return o
}

func (o CloudFormationStackMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CloudFormationStack] {
	return pulumix.Output[map[string]*CloudFormationStack]{
		OutputState: o.OutputState,
	}
}

func (o CloudFormationStackMapOutput) MapIndex(k pulumi.StringInput) CloudFormationStackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudFormationStack {
		return vs[0].(map[string]*CloudFormationStack)[vs[1].(string)]
	}).(CloudFormationStackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudFormationStackInput)(nil)).Elem(), &CloudFormationStack{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudFormationStackArrayInput)(nil)).Elem(), CloudFormationStackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudFormationStackMapInput)(nil)).Elem(), CloudFormationStackMap{})
	pulumi.RegisterOutputType(CloudFormationStackOutput{})
	pulumi.RegisterOutputType(CloudFormationStackArrayOutput{})
	pulumi.RegisterOutputType(CloudFormationStackMapOutput{})
}
