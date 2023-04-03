// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a static IP address attachment - relationship between a Lightsail static IP & Lightsail instance.
//
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lightsail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testStaticIp, err := lightsail.NewStaticIp(ctx, "testStaticIp", nil)
//			if err != nil {
//				return err
//			}
//			testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
//				AvailabilityZone: pulumi.String("us-east-1b"),
//				BlueprintId:      pulumi.String("string"),
//				BundleId:         pulumi.String("string"),
//				KeyPairName:      pulumi.String("some_key_name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = lightsail.NewStaticIpAttachment(ctx, "testStaticIpAttachment", &lightsail.StaticIpAttachmentArgs{
//				StaticIpName: testStaticIp.ID(),
//				InstanceName: testInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type StaticIpAttachment struct {
	pulumi.CustomResourceState

	// The name of the Lightsail instance to attach the IP to
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The allocated static IP address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The name of the allocated static IP
	StaticIpName pulumi.StringOutput `pulumi:"staticIpName"`
}

// NewStaticIpAttachment registers a new resource with the given unique name, arguments, and options.
func NewStaticIpAttachment(ctx *pulumi.Context,
	name string, args *StaticIpAttachmentArgs, opts ...pulumi.ResourceOption) (*StaticIpAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.StaticIpName == nil {
		return nil, errors.New("invalid value for required argument 'StaticIpName'")
	}
	var resource StaticIpAttachment
	err := ctx.RegisterResource("aws:lightsail/staticIpAttachment:StaticIpAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticIpAttachment gets an existing StaticIpAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticIpAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticIpAttachmentState, opts ...pulumi.ResourceOption) (*StaticIpAttachment, error) {
	var resource StaticIpAttachment
	err := ctx.ReadResource("aws:lightsail/staticIpAttachment:StaticIpAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticIpAttachment resources.
type staticIpAttachmentState struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName *string `pulumi:"instanceName"`
	// The allocated static IP address
	IpAddress *string `pulumi:"ipAddress"`
	// The name of the allocated static IP
	StaticIpName *string `pulumi:"staticIpName"`
}

type StaticIpAttachmentState struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName pulumi.StringPtrInput
	// The allocated static IP address
	IpAddress pulumi.StringPtrInput
	// The name of the allocated static IP
	StaticIpName pulumi.StringPtrInput
}

func (StaticIpAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpAttachmentState)(nil)).Elem()
}

type staticIpAttachmentArgs struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName string `pulumi:"instanceName"`
	// The name of the allocated static IP
	StaticIpName string `pulumi:"staticIpName"`
}

// The set of arguments for constructing a StaticIpAttachment resource.
type StaticIpAttachmentArgs struct {
	// The name of the Lightsail instance to attach the IP to
	InstanceName pulumi.StringInput
	// The name of the allocated static IP
	StaticIpName pulumi.StringInput
}

func (StaticIpAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpAttachmentArgs)(nil)).Elem()
}

type StaticIpAttachmentInput interface {
	pulumi.Input

	ToStaticIpAttachmentOutput() StaticIpAttachmentOutput
	ToStaticIpAttachmentOutputWithContext(ctx context.Context) StaticIpAttachmentOutput
}

func (*StaticIpAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticIpAttachment)(nil)).Elem()
}

func (i *StaticIpAttachment) ToStaticIpAttachmentOutput() StaticIpAttachmentOutput {
	return i.ToStaticIpAttachmentOutputWithContext(context.Background())
}

func (i *StaticIpAttachment) ToStaticIpAttachmentOutputWithContext(ctx context.Context) StaticIpAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpAttachmentOutput)
}

// StaticIpAttachmentArrayInput is an input type that accepts StaticIpAttachmentArray and StaticIpAttachmentArrayOutput values.
// You can construct a concrete instance of `StaticIpAttachmentArrayInput` via:
//
//	StaticIpAttachmentArray{ StaticIpAttachmentArgs{...} }
type StaticIpAttachmentArrayInput interface {
	pulumi.Input

	ToStaticIpAttachmentArrayOutput() StaticIpAttachmentArrayOutput
	ToStaticIpAttachmentArrayOutputWithContext(context.Context) StaticIpAttachmentArrayOutput
}

type StaticIpAttachmentArray []StaticIpAttachmentInput

func (StaticIpAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticIpAttachment)(nil)).Elem()
}

func (i StaticIpAttachmentArray) ToStaticIpAttachmentArrayOutput() StaticIpAttachmentArrayOutput {
	return i.ToStaticIpAttachmentArrayOutputWithContext(context.Background())
}

func (i StaticIpAttachmentArray) ToStaticIpAttachmentArrayOutputWithContext(ctx context.Context) StaticIpAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpAttachmentArrayOutput)
}

// StaticIpAttachmentMapInput is an input type that accepts StaticIpAttachmentMap and StaticIpAttachmentMapOutput values.
// You can construct a concrete instance of `StaticIpAttachmentMapInput` via:
//
//	StaticIpAttachmentMap{ "key": StaticIpAttachmentArgs{...} }
type StaticIpAttachmentMapInput interface {
	pulumi.Input

	ToStaticIpAttachmentMapOutput() StaticIpAttachmentMapOutput
	ToStaticIpAttachmentMapOutputWithContext(context.Context) StaticIpAttachmentMapOutput
}

type StaticIpAttachmentMap map[string]StaticIpAttachmentInput

func (StaticIpAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticIpAttachment)(nil)).Elem()
}

func (i StaticIpAttachmentMap) ToStaticIpAttachmentMapOutput() StaticIpAttachmentMapOutput {
	return i.ToStaticIpAttachmentMapOutputWithContext(context.Background())
}

func (i StaticIpAttachmentMap) ToStaticIpAttachmentMapOutputWithContext(ctx context.Context) StaticIpAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpAttachmentMapOutput)
}

type StaticIpAttachmentOutput struct{ *pulumi.OutputState }

func (StaticIpAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticIpAttachment)(nil)).Elem()
}

func (o StaticIpAttachmentOutput) ToStaticIpAttachmentOutput() StaticIpAttachmentOutput {
	return o
}

func (o StaticIpAttachmentOutput) ToStaticIpAttachmentOutputWithContext(ctx context.Context) StaticIpAttachmentOutput {
	return o
}

// The name of the Lightsail instance to attach the IP to
func (o StaticIpAttachmentOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticIpAttachment) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The allocated static IP address
func (o StaticIpAttachmentOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticIpAttachment) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The name of the allocated static IP
func (o StaticIpAttachmentOutput) StaticIpName() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticIpAttachment) pulumi.StringOutput { return v.StaticIpName }).(pulumi.StringOutput)
}

type StaticIpAttachmentArrayOutput struct{ *pulumi.OutputState }

func (StaticIpAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticIpAttachment)(nil)).Elem()
}

func (o StaticIpAttachmentArrayOutput) ToStaticIpAttachmentArrayOutput() StaticIpAttachmentArrayOutput {
	return o
}

func (o StaticIpAttachmentArrayOutput) ToStaticIpAttachmentArrayOutputWithContext(ctx context.Context) StaticIpAttachmentArrayOutput {
	return o
}

func (o StaticIpAttachmentArrayOutput) Index(i pulumi.IntInput) StaticIpAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StaticIpAttachment {
		return vs[0].([]*StaticIpAttachment)[vs[1].(int)]
	}).(StaticIpAttachmentOutput)
}

type StaticIpAttachmentMapOutput struct{ *pulumi.OutputState }

func (StaticIpAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticIpAttachment)(nil)).Elem()
}

func (o StaticIpAttachmentMapOutput) ToStaticIpAttachmentMapOutput() StaticIpAttachmentMapOutput {
	return o
}

func (o StaticIpAttachmentMapOutput) ToStaticIpAttachmentMapOutputWithContext(ctx context.Context) StaticIpAttachmentMapOutput {
	return o
}

func (o StaticIpAttachmentMapOutput) MapIndex(k pulumi.StringInput) StaticIpAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StaticIpAttachment {
		return vs[0].(map[string]*StaticIpAttachment)[vs[1].(string)]
	}).(StaticIpAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticIpAttachmentInput)(nil)).Elem(), &StaticIpAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticIpAttachmentArrayInput)(nil)).Elem(), StaticIpAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticIpAttachmentMapInput)(nil)).Elem(), StaticIpAttachmentMap{})
	pulumi.RegisterOutputType(StaticIpAttachmentOutput{})
	pulumi.RegisterOutputType(StaticIpAttachmentArrayOutput{})
	pulumi.RegisterOutputType(StaticIpAttachmentMapOutput{})
}
