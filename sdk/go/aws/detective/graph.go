// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package detective

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a resource to manage an [AWS Detective Graph](https://docs.aws.amazon.com/detective/latest/APIReference/API_CreateGraph.html). As an AWS account may own only one Detective graph per region, provisioning multiple Detective graphs requires a separate provider configuration for each graph.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/detective"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := detective.NewGraph(ctx, "example", &detective.GraphArgs{
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("example-detective-graph"),
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
// Using `pulumi import`, import `aws_detective_graph` using the ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:detective/graph:Graph example arn:aws:detective:us-east-1:123456789101:graph:231684d34gh74g4bae1dbc7bd807d02d
//
// ```
type Graph struct {
	pulumi.CustomResourceState

	// Date and time, in UTC and extended RFC 3339 format, when the Amazon Detective Graph was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// ARN of the Detective Graph.
	GraphArn pulumi.StringOutput `pulumi:"graphArn"`
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewGraph registers a new resource with the given unique name, arguments, and options.
func NewGraph(ctx *pulumi.Context,
	name string, args *GraphArgs, opts ...pulumi.ResourceOption) (*Graph, error) {
	if args == nil {
		args = &GraphArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Graph
	err := ctx.RegisterResource("aws:detective/graph:Graph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraph gets an existing Graph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphState, opts ...pulumi.ResourceOption) (*Graph, error) {
	var resource Graph
	err := ctx.ReadResource("aws:detective/graph:Graph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Graph resources.
type graphState struct {
	// Date and time, in UTC and extended RFC 3339 format, when the Amazon Detective Graph was created.
	CreatedTime *string `pulumi:"createdTime"`
	// ARN of the Detective Graph.
	GraphArn *string `pulumi:"graphArn"`
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type GraphState struct {
	// Date and time, in UTC and extended RFC 3339 format, when the Amazon Detective Graph was created.
	CreatedTime pulumi.StringPtrInput
	// ARN of the Detective Graph.
	GraphArn pulumi.StringPtrInput
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (GraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphState)(nil)).Elem()
}

type graphArgs struct {
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Graph resource.
type GraphArgs struct {
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (GraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphArgs)(nil)).Elem()
}

type GraphInput interface {
	pulumi.Input

	ToGraphOutput() GraphOutput
	ToGraphOutputWithContext(ctx context.Context) GraphOutput
}

func (*Graph) ElementType() reflect.Type {
	return reflect.TypeOf((**Graph)(nil)).Elem()
}

func (i *Graph) ToGraphOutput() GraphOutput {
	return i.ToGraphOutputWithContext(context.Background())
}

func (i *Graph) ToGraphOutputWithContext(ctx context.Context) GraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphOutput)
}

func (i *Graph) ToOutput(ctx context.Context) pulumix.Output[*Graph] {
	return pulumix.Output[*Graph]{
		OutputState: i.ToGraphOutputWithContext(ctx).OutputState,
	}
}

// GraphArrayInput is an input type that accepts GraphArray and GraphArrayOutput values.
// You can construct a concrete instance of `GraphArrayInput` via:
//
//	GraphArray{ GraphArgs{...} }
type GraphArrayInput interface {
	pulumi.Input

	ToGraphArrayOutput() GraphArrayOutput
	ToGraphArrayOutputWithContext(context.Context) GraphArrayOutput
}

type GraphArray []GraphInput

func (GraphArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Graph)(nil)).Elem()
}

func (i GraphArray) ToGraphArrayOutput() GraphArrayOutput {
	return i.ToGraphArrayOutputWithContext(context.Background())
}

func (i GraphArray) ToGraphArrayOutputWithContext(ctx context.Context) GraphArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphArrayOutput)
}

func (i GraphArray) ToOutput(ctx context.Context) pulumix.Output[[]*Graph] {
	return pulumix.Output[[]*Graph]{
		OutputState: i.ToGraphArrayOutputWithContext(ctx).OutputState,
	}
}

// GraphMapInput is an input type that accepts GraphMap and GraphMapOutput values.
// You can construct a concrete instance of `GraphMapInput` via:
//
//	GraphMap{ "key": GraphArgs{...} }
type GraphMapInput interface {
	pulumi.Input

	ToGraphMapOutput() GraphMapOutput
	ToGraphMapOutputWithContext(context.Context) GraphMapOutput
}

type GraphMap map[string]GraphInput

func (GraphMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Graph)(nil)).Elem()
}

func (i GraphMap) ToGraphMapOutput() GraphMapOutput {
	return i.ToGraphMapOutputWithContext(context.Background())
}

func (i GraphMap) ToGraphMapOutputWithContext(ctx context.Context) GraphMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphMapOutput)
}

func (i GraphMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Graph] {
	return pulumix.Output[map[string]*Graph]{
		OutputState: i.ToGraphMapOutputWithContext(ctx).OutputState,
	}
}

type GraphOutput struct{ *pulumi.OutputState }

func (GraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Graph)(nil)).Elem()
}

func (o GraphOutput) ToGraphOutput() GraphOutput {
	return o
}

func (o GraphOutput) ToGraphOutputWithContext(ctx context.Context) GraphOutput {
	return o
}

func (o GraphOutput) ToOutput(ctx context.Context) pulumix.Output[*Graph] {
	return pulumix.Output[*Graph]{
		OutputState: o.OutputState,
	}
}

// Date and time, in UTC and extended RFC 3339 format, when the Amazon Detective Graph was created.
func (o GraphOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Graph) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// ARN of the Detective Graph.
func (o GraphOutput) GraphArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Graph) pulumi.StringOutput { return v.GraphArn }).(pulumi.StringOutput)
}

// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o GraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Graph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o GraphOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Graph) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type GraphArrayOutput struct{ *pulumi.OutputState }

func (GraphArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Graph)(nil)).Elem()
}

func (o GraphArrayOutput) ToGraphArrayOutput() GraphArrayOutput {
	return o
}

func (o GraphArrayOutput) ToGraphArrayOutputWithContext(ctx context.Context) GraphArrayOutput {
	return o
}

func (o GraphArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Graph] {
	return pulumix.Output[[]*Graph]{
		OutputState: o.OutputState,
	}
}

func (o GraphArrayOutput) Index(i pulumi.IntInput) GraphOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Graph {
		return vs[0].([]*Graph)[vs[1].(int)]
	}).(GraphOutput)
}

type GraphMapOutput struct{ *pulumi.OutputState }

func (GraphMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Graph)(nil)).Elem()
}

func (o GraphMapOutput) ToGraphMapOutput() GraphMapOutput {
	return o
}

func (o GraphMapOutput) ToGraphMapOutputWithContext(ctx context.Context) GraphMapOutput {
	return o
}

func (o GraphMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Graph] {
	return pulumix.Output[map[string]*Graph]{
		OutputState: o.OutputState,
	}
}

func (o GraphMapOutput) MapIndex(k pulumi.StringInput) GraphOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Graph {
		return vs[0].(map[string]*Graph)[vs[1].(string)]
	}).(GraphOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GraphInput)(nil)).Elem(), &Graph{})
	pulumi.RegisterInputType(reflect.TypeOf((*GraphArrayInput)(nil)).Elem(), GraphArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GraphMapInput)(nil)).Elem(), GraphMap{})
	pulumi.RegisterOutputType(GraphOutput{})
	pulumi.RegisterOutputType(GraphArrayOutput{})
	pulumi.RegisterOutputType(GraphMapOutput{})
}
