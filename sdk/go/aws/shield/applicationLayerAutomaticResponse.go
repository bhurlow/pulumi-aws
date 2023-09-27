// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package shield

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource for managing an AWS Shield Application Layer Automatic Response for automatic DDoS mitigation.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			currentPartition, err := aws.GetPartition(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = shield.NewApplicationLayerAutomaticResponse(ctx, "example", &shield.ApplicationLayerAutomaticResponseArgs{
//				Action:      pulumi.String("COUNT"),
//				ResourceArn: pulumi.String(fmt.Sprintf("arn:%v:cloudfront:%v:distribution/%v", currentPartition.Partition, currentCallerIdentity.AccountId, _var.Distribution_id)),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ApplicationLayerAutomaticResponse struct {
	pulumi.CustomResourceState

	// One of `COUNT` or `BLOCK`
	Action pulumi.StringOutput `pulumi:"action"`
	// ARN of the resource to protect (Cloudfront Distributions and ALBs only at this time).
	ResourceArn pulumi.StringOutput                                `pulumi:"resourceArn"`
	Timeouts    ApplicationLayerAutomaticResponseTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewApplicationLayerAutomaticResponse registers a new resource with the given unique name, arguments, and options.
func NewApplicationLayerAutomaticResponse(ctx *pulumi.Context,
	name string, args *ApplicationLayerAutomaticResponseArgs, opts ...pulumi.ResourceOption) (*ApplicationLayerAutomaticResponse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationLayerAutomaticResponse
	err := ctx.RegisterResource("aws:shield/applicationLayerAutomaticResponse:ApplicationLayerAutomaticResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationLayerAutomaticResponse gets an existing ApplicationLayerAutomaticResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationLayerAutomaticResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationLayerAutomaticResponseState, opts ...pulumi.ResourceOption) (*ApplicationLayerAutomaticResponse, error) {
	var resource ApplicationLayerAutomaticResponse
	err := ctx.ReadResource("aws:shield/applicationLayerAutomaticResponse:ApplicationLayerAutomaticResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationLayerAutomaticResponse resources.
type applicationLayerAutomaticResponseState struct {
	// One of `COUNT` or `BLOCK`
	Action *string `pulumi:"action"`
	// ARN of the resource to protect (Cloudfront Distributions and ALBs only at this time).
	ResourceArn *string                                    `pulumi:"resourceArn"`
	Timeouts    *ApplicationLayerAutomaticResponseTimeouts `pulumi:"timeouts"`
}

type ApplicationLayerAutomaticResponseState struct {
	// One of `COUNT` or `BLOCK`
	Action pulumi.StringPtrInput
	// ARN of the resource to protect (Cloudfront Distributions and ALBs only at this time).
	ResourceArn pulumi.StringPtrInput
	Timeouts    ApplicationLayerAutomaticResponseTimeoutsPtrInput
}

func (ApplicationLayerAutomaticResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationLayerAutomaticResponseState)(nil)).Elem()
}

type applicationLayerAutomaticResponseArgs struct {
	// One of `COUNT` or `BLOCK`
	Action string `pulumi:"action"`
	// ARN of the resource to protect (Cloudfront Distributions and ALBs only at this time).
	ResourceArn string                                     `pulumi:"resourceArn"`
	Timeouts    *ApplicationLayerAutomaticResponseTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApplicationLayerAutomaticResponse resource.
type ApplicationLayerAutomaticResponseArgs struct {
	// One of `COUNT` or `BLOCK`
	Action pulumi.StringInput
	// ARN of the resource to protect (Cloudfront Distributions and ALBs only at this time).
	ResourceArn pulumi.StringInput
	Timeouts    ApplicationLayerAutomaticResponseTimeoutsPtrInput
}

func (ApplicationLayerAutomaticResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationLayerAutomaticResponseArgs)(nil)).Elem()
}

type ApplicationLayerAutomaticResponseInput interface {
	pulumi.Input

	ToApplicationLayerAutomaticResponseOutput() ApplicationLayerAutomaticResponseOutput
	ToApplicationLayerAutomaticResponseOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseOutput
}

func (*ApplicationLayerAutomaticResponse) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLayerAutomaticResponse)(nil)).Elem()
}

func (i *ApplicationLayerAutomaticResponse) ToApplicationLayerAutomaticResponseOutput() ApplicationLayerAutomaticResponseOutput {
	return i.ToApplicationLayerAutomaticResponseOutputWithContext(context.Background())
}

func (i *ApplicationLayerAutomaticResponse) ToApplicationLayerAutomaticResponseOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLayerAutomaticResponseOutput)
}

func (i *ApplicationLayerAutomaticResponse) ToOutput(ctx context.Context) pulumix.Output[*ApplicationLayerAutomaticResponse] {
	return pulumix.Output[*ApplicationLayerAutomaticResponse]{
		OutputState: i.ToApplicationLayerAutomaticResponseOutputWithContext(ctx).OutputState,
	}
}

// ApplicationLayerAutomaticResponseArrayInput is an input type that accepts ApplicationLayerAutomaticResponseArray and ApplicationLayerAutomaticResponseArrayOutput values.
// You can construct a concrete instance of `ApplicationLayerAutomaticResponseArrayInput` via:
//
//	ApplicationLayerAutomaticResponseArray{ ApplicationLayerAutomaticResponseArgs{...} }
type ApplicationLayerAutomaticResponseArrayInput interface {
	pulumi.Input

	ToApplicationLayerAutomaticResponseArrayOutput() ApplicationLayerAutomaticResponseArrayOutput
	ToApplicationLayerAutomaticResponseArrayOutputWithContext(context.Context) ApplicationLayerAutomaticResponseArrayOutput
}

type ApplicationLayerAutomaticResponseArray []ApplicationLayerAutomaticResponseInput

func (ApplicationLayerAutomaticResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationLayerAutomaticResponse)(nil)).Elem()
}

func (i ApplicationLayerAutomaticResponseArray) ToApplicationLayerAutomaticResponseArrayOutput() ApplicationLayerAutomaticResponseArrayOutput {
	return i.ToApplicationLayerAutomaticResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationLayerAutomaticResponseArray) ToApplicationLayerAutomaticResponseArrayOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLayerAutomaticResponseArrayOutput)
}

func (i ApplicationLayerAutomaticResponseArray) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationLayerAutomaticResponse] {
	return pulumix.Output[[]*ApplicationLayerAutomaticResponse]{
		OutputState: i.ToApplicationLayerAutomaticResponseArrayOutputWithContext(ctx).OutputState,
	}
}

// ApplicationLayerAutomaticResponseMapInput is an input type that accepts ApplicationLayerAutomaticResponseMap and ApplicationLayerAutomaticResponseMapOutput values.
// You can construct a concrete instance of `ApplicationLayerAutomaticResponseMapInput` via:
//
//	ApplicationLayerAutomaticResponseMap{ "key": ApplicationLayerAutomaticResponseArgs{...} }
type ApplicationLayerAutomaticResponseMapInput interface {
	pulumi.Input

	ToApplicationLayerAutomaticResponseMapOutput() ApplicationLayerAutomaticResponseMapOutput
	ToApplicationLayerAutomaticResponseMapOutputWithContext(context.Context) ApplicationLayerAutomaticResponseMapOutput
}

type ApplicationLayerAutomaticResponseMap map[string]ApplicationLayerAutomaticResponseInput

func (ApplicationLayerAutomaticResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationLayerAutomaticResponse)(nil)).Elem()
}

func (i ApplicationLayerAutomaticResponseMap) ToApplicationLayerAutomaticResponseMapOutput() ApplicationLayerAutomaticResponseMapOutput {
	return i.ToApplicationLayerAutomaticResponseMapOutputWithContext(context.Background())
}

func (i ApplicationLayerAutomaticResponseMap) ToApplicationLayerAutomaticResponseMapOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLayerAutomaticResponseMapOutput)
}

func (i ApplicationLayerAutomaticResponseMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationLayerAutomaticResponse] {
	return pulumix.Output[map[string]*ApplicationLayerAutomaticResponse]{
		OutputState: i.ToApplicationLayerAutomaticResponseMapOutputWithContext(ctx).OutputState,
	}
}

type ApplicationLayerAutomaticResponseOutput struct{ *pulumi.OutputState }

func (ApplicationLayerAutomaticResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLayerAutomaticResponse)(nil)).Elem()
}

func (o ApplicationLayerAutomaticResponseOutput) ToApplicationLayerAutomaticResponseOutput() ApplicationLayerAutomaticResponseOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseOutput) ToApplicationLayerAutomaticResponseOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseOutput) ToOutput(ctx context.Context) pulumix.Output[*ApplicationLayerAutomaticResponse] {
	return pulumix.Output[*ApplicationLayerAutomaticResponse]{
		OutputState: o.OutputState,
	}
}

// One of `COUNT` or `BLOCK`
func (o ApplicationLayerAutomaticResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationLayerAutomaticResponse) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// ARN of the resource to protect (Cloudfront Distributions and ALBs only at this time).
func (o ApplicationLayerAutomaticResponseOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationLayerAutomaticResponse) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

func (o ApplicationLayerAutomaticResponseOutput) Timeouts() ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApplicationLayerAutomaticResponse) ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
		return v.Timeouts
	}).(ApplicationLayerAutomaticResponseTimeoutsPtrOutput)
}

type ApplicationLayerAutomaticResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationLayerAutomaticResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationLayerAutomaticResponse)(nil)).Elem()
}

func (o ApplicationLayerAutomaticResponseArrayOutput) ToApplicationLayerAutomaticResponseArrayOutput() ApplicationLayerAutomaticResponseArrayOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseArrayOutput) ToApplicationLayerAutomaticResponseArrayOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseArrayOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ApplicationLayerAutomaticResponse] {
	return pulumix.Output[[]*ApplicationLayerAutomaticResponse]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationLayerAutomaticResponseArrayOutput) Index(i pulumi.IntInput) ApplicationLayerAutomaticResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationLayerAutomaticResponse {
		return vs[0].([]*ApplicationLayerAutomaticResponse)[vs[1].(int)]
	}).(ApplicationLayerAutomaticResponseOutput)
}

type ApplicationLayerAutomaticResponseMapOutput struct{ *pulumi.OutputState }

func (ApplicationLayerAutomaticResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationLayerAutomaticResponse)(nil)).Elem()
}

func (o ApplicationLayerAutomaticResponseMapOutput) ToApplicationLayerAutomaticResponseMapOutput() ApplicationLayerAutomaticResponseMapOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseMapOutput) ToApplicationLayerAutomaticResponseMapOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseMapOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApplicationLayerAutomaticResponse] {
	return pulumix.Output[map[string]*ApplicationLayerAutomaticResponse]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationLayerAutomaticResponseMapOutput) MapIndex(k pulumi.StringInput) ApplicationLayerAutomaticResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationLayerAutomaticResponse {
		return vs[0].(map[string]*ApplicationLayerAutomaticResponse)[vs[1].(string)]
	}).(ApplicationLayerAutomaticResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLayerAutomaticResponseInput)(nil)).Elem(), &ApplicationLayerAutomaticResponse{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLayerAutomaticResponseArrayInput)(nil)).Elem(), ApplicationLayerAutomaticResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLayerAutomaticResponseMapInput)(nil)).Elem(), ApplicationLayerAutomaticResponseMap{})
	pulumi.RegisterOutputType(ApplicationLayerAutomaticResponseOutput{})
	pulumi.RegisterOutputType(ApplicationLayerAutomaticResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationLayerAutomaticResponseMapOutput{})
}
