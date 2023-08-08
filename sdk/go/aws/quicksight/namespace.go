// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS QuickSight Namespace.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := quicksight.NewNamespace(ctx, "example", &quicksight.NamespaceArgs{
//				Namespace: pulumi.String("example"),
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
//	to = aws_quicksight_namespace.example
//
//	id = "123456789012,example" } Using `pulumi import`, import QuickSight Namespace using the AWS account ID and namespace separated by commas (`,`). For exampleconsole % pulumi import aws_quicksight_namespace.example 123456789012,example
type Namespace struct {
	pulumi.CustomResourceState

	// ARN of the Namespace.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// Namespace AWS Region.
	CapacityRegion pulumi.StringOutput `pulumi:"capacityRegion"`
	// Creation status of the namespace.
	CreationStatus pulumi.StringOutput `pulumi:"creationStatus"`
	// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
	IdentityStore pulumi.StringOutput `pulumi:"identityStore"`
	// Name of the namespace.
	//
	// The following arguments are optional:
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll  pulumi.StringMapOutput     `pulumi:"tagsAll"`
	Timeouts NamespaceTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Namespace
	err := ctx.RegisterResource("aws:quicksight/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("aws:quicksight/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// ARN of the Namespace.
	Arn *string `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// Namespace AWS Region.
	CapacityRegion *string `pulumi:"capacityRegion"`
	// Creation status of the namespace.
	CreationStatus *string `pulumi:"creationStatus"`
	// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
	IdentityStore *string `pulumi:"identityStore"`
	// Name of the namespace.
	//
	// The following arguments are optional:
	Namespace *string `pulumi:"namespace"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll  map[string]string  `pulumi:"tagsAll"`
	Timeouts *NamespaceTimeouts `pulumi:"timeouts"`
}

type NamespaceState struct {
	// ARN of the Namespace.
	Arn pulumi.StringPtrInput
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// Namespace AWS Region.
	CapacityRegion pulumi.StringPtrInput
	// Creation status of the namespace.
	CreationStatus pulumi.StringPtrInput
	// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
	IdentityStore pulumi.StringPtrInput
	// Name of the namespace.
	//
	// The following arguments are optional:
	Namespace pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll  pulumi.StringMapInput
	Timeouts NamespaceTimeoutsPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
	IdentityStore *string `pulumi:"identityStore"`
	// Name of the namespace.
	//
	// The following arguments are optional:
	Namespace string `pulumi:"namespace"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags     map[string]string  `pulumi:"tags"`
	Timeouts *NamespaceTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
	IdentityStore pulumi.StringPtrInput
	// Name of the namespace.
	//
	// The following arguments are optional:
	Namespace pulumi.StringInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags     pulumi.StringMapInput
	Timeouts NamespaceTimeoutsPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

// NamespaceArrayInput is an input type that accepts NamespaceArray and NamespaceArrayOutput values.
// You can construct a concrete instance of `NamespaceArrayInput` via:
//
//	NamespaceArray{ NamespaceArgs{...} }
type NamespaceArrayInput interface {
	pulumi.Input

	ToNamespaceArrayOutput() NamespaceArrayOutput
	ToNamespaceArrayOutputWithContext(context.Context) NamespaceArrayOutput
}

type NamespaceArray []NamespaceInput

func (NamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (i NamespaceArray) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return i.ToNamespaceArrayOutputWithContext(context.Background())
}

func (i NamespaceArray) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceArrayOutput)
}

// NamespaceMapInput is an input type that accepts NamespaceMap and NamespaceMapOutput values.
// You can construct a concrete instance of `NamespaceMapInput` via:
//
//	NamespaceMap{ "key": NamespaceArgs{...} }
type NamespaceMapInput interface {
	pulumi.Input

	ToNamespaceMapOutput() NamespaceMapOutput
	ToNamespaceMapOutputWithContext(context.Context) NamespaceMapOutput
}

type NamespaceMap map[string]NamespaceInput

func (NamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

// ARN of the Namespace.
func (o NamespaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// AWS account ID.
func (o NamespaceOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

// Namespace AWS Region.
func (o NamespaceOutput) CapacityRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.CapacityRegion }).(pulumi.StringOutput)
}

// Creation status of the namespace.
func (o NamespaceOutput) CreationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.CreationStatus }).(pulumi.StringOutput)
}

// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
func (o NamespaceOutput) IdentityStore() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.IdentityStore }).(pulumi.StringOutput)
}

// Name of the namespace.
//
// The following arguments are optional:
func (o NamespaceOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o NamespaceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

func (o NamespaceOutput) Timeouts() NamespaceTimeoutsPtrOutput {
	return o.ApplyT(func(v *Namespace) NamespaceTimeoutsPtrOutput { return v.Timeouts }).(NamespaceTimeoutsPtrOutput)
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].([]*Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].(map[string]*Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceInput)(nil)).Elem(), &Namespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceArrayInput)(nil)).Elem(), NamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceMapInput)(nil)).Elem(), NamespaceMap{})
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}
