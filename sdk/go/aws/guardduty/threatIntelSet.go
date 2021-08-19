// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a GuardDuty ThreatIntelSet.
//
// > **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage ThreatIntelSets. ThreatIntelSets that are uploaded by the primary account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-threat-intel-set.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/guardduty"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := guardduty.NewDetector(ctx, "primary", &guardduty.DetectorArgs{
// 			Enable: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myThreatIntelSetBucketObject, err := s3.NewBucketObject(ctx, "myThreatIntelSetBucketObject", &s3.BucketObjectArgs{
// 			Acl:     pulumi.String("public-read"),
// 			Content: pulumi.String("10.0.0.0/8\n"),
// 			Bucket:  bucket.ID(),
// 			Key:     pulumi.String("MyThreatIntelSet"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = guardduty.NewThreatIntelSet(ctx, "myThreatIntelSetThreatIntelSet", &guardduty.ThreatIntelSetArgs{
// 			Activate:   pulumi.Bool(true),
// 			DetectorId: primary.ID(),
// 			Format:     pulumi.String("TXT"),
// 			Location: pulumi.All(myThreatIntelSetBucketObject.Bucket, myThreatIntelSetBucketObject.Key).ApplyT(func(_args []interface{}) (string, error) {
// 				bucket := _args[0].(string)
// 				key := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v", "https://s3.amazonaws.com/", bucket, "/", key), nil
// 			}).(pulumi.StringOutput),
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
// GuardDuty ThreatIntelSet can be imported using the the primary GuardDuty detector ID and ThreatIntelSetID, e.g.
//
// ```sh
//  $ pulumi import aws:guardduty/threatIntelSet:ThreatIntelSet MyThreatIntelSet 00b00fd5aecc0ab60a708659477e9617:123456789012
// ```
type ThreatIntelSet struct {
	pulumi.CustomResourceState

	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate pulumi.BoolOutput `pulumi:"activate"`
	// Amazon Resource Name (ARN) of the GuardDuty ThreatIntelSet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringOutput `pulumi:"format"`
	// The URI of the file that contains the ThreatIntelSet.
	Location pulumi.StringOutput `pulumi:"location"`
	// The friendly name to identify the ThreatIntelSet.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewThreatIntelSet registers a new resource with the given unique name, arguments, and options.
func NewThreatIntelSet(ctx *pulumi.Context,
	name string, args *ThreatIntelSetArgs, opts ...pulumi.ResourceOption) (*ThreatIntelSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Activate == nil {
		return nil, errors.New("invalid value for required argument 'Activate'")
	}
	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource ThreatIntelSet
	err := ctx.RegisterResource("aws:guardduty/threatIntelSet:ThreatIntelSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThreatIntelSet gets an existing ThreatIntelSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThreatIntelSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelSetState, opts ...pulumi.ResourceOption) (*ThreatIntelSet, error) {
	var resource ThreatIntelSet
	err := ctx.ReadResource("aws:guardduty/threatIntelSet:ThreatIntelSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThreatIntelSet resources.
type threatIntelSetState struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate *bool `pulumi:"activate"`
	// Amazon Resource Name (ARN) of the GuardDuty ThreatIntelSet.
	Arn *string `pulumi:"arn"`
	// The detector ID of the GuardDuty.
	DetectorId *string `pulumi:"detectorId"`
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format *string `pulumi:"format"`
	// The URI of the file that contains the ThreatIntelSet.
	Location *string `pulumi:"location"`
	// The friendly name to identify the ThreatIntelSet.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ThreatIntelSetState struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the GuardDuty ThreatIntelSet.
	Arn pulumi.StringPtrInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringPtrInput
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringPtrInput
	// The URI of the file that contains the ThreatIntelSet.
	Location pulumi.StringPtrInput
	// The friendly name to identify the ThreatIntelSet.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ThreatIntelSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelSetState)(nil)).Elem()
}

type threatIntelSetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate bool `pulumi:"activate"`
	// The detector ID of the GuardDuty.
	DetectorId string `pulumi:"detectorId"`
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format string `pulumi:"format"`
	// The URI of the file that contains the ThreatIntelSet.
	Location string `pulumi:"location"`
	// The friendly name to identify the ThreatIntelSet.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ThreatIntelSet resource.
type ThreatIntelSetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate pulumi.BoolInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringInput
	// The format of the file that contains the ThreatIntelSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringInput
	// The URI of the file that contains the ThreatIntelSet.
	Location pulumi.StringInput
	// The friendly name to identify the ThreatIntelSet.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ThreatIntelSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelSetArgs)(nil)).Elem()
}

type ThreatIntelSetInput interface {
	pulumi.Input

	ToThreatIntelSetOutput() ThreatIntelSetOutput
	ToThreatIntelSetOutputWithContext(ctx context.Context) ThreatIntelSetOutput
}

func (*ThreatIntelSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelSet)(nil))
}

func (i *ThreatIntelSet) ToThreatIntelSetOutput() ThreatIntelSetOutput {
	return i.ToThreatIntelSetOutputWithContext(context.Background())
}

func (i *ThreatIntelSet) ToThreatIntelSetOutputWithContext(ctx context.Context) ThreatIntelSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelSetOutput)
}

func (i *ThreatIntelSet) ToThreatIntelSetPtrOutput() ThreatIntelSetPtrOutput {
	return i.ToThreatIntelSetPtrOutputWithContext(context.Background())
}

func (i *ThreatIntelSet) ToThreatIntelSetPtrOutputWithContext(ctx context.Context) ThreatIntelSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelSetPtrOutput)
}

type ThreatIntelSetPtrInput interface {
	pulumi.Input

	ToThreatIntelSetPtrOutput() ThreatIntelSetPtrOutput
	ToThreatIntelSetPtrOutputWithContext(ctx context.Context) ThreatIntelSetPtrOutput
}

type threatIntelSetPtrType ThreatIntelSetArgs

func (*threatIntelSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelSet)(nil))
}

func (i *threatIntelSetPtrType) ToThreatIntelSetPtrOutput() ThreatIntelSetPtrOutput {
	return i.ToThreatIntelSetPtrOutputWithContext(context.Background())
}

func (i *threatIntelSetPtrType) ToThreatIntelSetPtrOutputWithContext(ctx context.Context) ThreatIntelSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelSetPtrOutput)
}

// ThreatIntelSetArrayInput is an input type that accepts ThreatIntelSetArray and ThreatIntelSetArrayOutput values.
// You can construct a concrete instance of `ThreatIntelSetArrayInput` via:
//
//          ThreatIntelSetArray{ ThreatIntelSetArgs{...} }
type ThreatIntelSetArrayInput interface {
	pulumi.Input

	ToThreatIntelSetArrayOutput() ThreatIntelSetArrayOutput
	ToThreatIntelSetArrayOutputWithContext(context.Context) ThreatIntelSetArrayOutput
}

type ThreatIntelSetArray []ThreatIntelSetInput

func (ThreatIntelSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ThreatIntelSet)(nil)).Elem()
}

func (i ThreatIntelSetArray) ToThreatIntelSetArrayOutput() ThreatIntelSetArrayOutput {
	return i.ToThreatIntelSetArrayOutputWithContext(context.Background())
}

func (i ThreatIntelSetArray) ToThreatIntelSetArrayOutputWithContext(ctx context.Context) ThreatIntelSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelSetArrayOutput)
}

// ThreatIntelSetMapInput is an input type that accepts ThreatIntelSetMap and ThreatIntelSetMapOutput values.
// You can construct a concrete instance of `ThreatIntelSetMapInput` via:
//
//          ThreatIntelSetMap{ "key": ThreatIntelSetArgs{...} }
type ThreatIntelSetMapInput interface {
	pulumi.Input

	ToThreatIntelSetMapOutput() ThreatIntelSetMapOutput
	ToThreatIntelSetMapOutputWithContext(context.Context) ThreatIntelSetMapOutput
}

type ThreatIntelSetMap map[string]ThreatIntelSetInput

func (ThreatIntelSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ThreatIntelSet)(nil)).Elem()
}

func (i ThreatIntelSetMap) ToThreatIntelSetMapOutput() ThreatIntelSetMapOutput {
	return i.ToThreatIntelSetMapOutputWithContext(context.Background())
}

func (i ThreatIntelSetMap) ToThreatIntelSetMapOutputWithContext(ctx context.Context) ThreatIntelSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelSetMapOutput)
}

type ThreatIntelSetOutput struct{ *pulumi.OutputState }

func (ThreatIntelSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelSet)(nil))
}

func (o ThreatIntelSetOutput) ToThreatIntelSetOutput() ThreatIntelSetOutput {
	return o
}

func (o ThreatIntelSetOutput) ToThreatIntelSetOutputWithContext(ctx context.Context) ThreatIntelSetOutput {
	return o
}

func (o ThreatIntelSetOutput) ToThreatIntelSetPtrOutput() ThreatIntelSetPtrOutput {
	return o.ToThreatIntelSetPtrOutputWithContext(context.Background())
}

func (o ThreatIntelSetOutput) ToThreatIntelSetPtrOutputWithContext(ctx context.Context) ThreatIntelSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThreatIntelSet) *ThreatIntelSet {
		return &v
	}).(ThreatIntelSetPtrOutput)
}

type ThreatIntelSetPtrOutput struct{ *pulumi.OutputState }

func (ThreatIntelSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelSet)(nil))
}

func (o ThreatIntelSetPtrOutput) ToThreatIntelSetPtrOutput() ThreatIntelSetPtrOutput {
	return o
}

func (o ThreatIntelSetPtrOutput) ToThreatIntelSetPtrOutputWithContext(ctx context.Context) ThreatIntelSetPtrOutput {
	return o
}

func (o ThreatIntelSetPtrOutput) Elem() ThreatIntelSetOutput {
	return o.ApplyT(func(v *ThreatIntelSet) ThreatIntelSet {
		if v != nil {
			return *v
		}
		var ret ThreatIntelSet
		return ret
	}).(ThreatIntelSetOutput)
}

type ThreatIntelSetArrayOutput struct{ *pulumi.OutputState }

func (ThreatIntelSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelSet)(nil))
}

func (o ThreatIntelSetArrayOutput) ToThreatIntelSetArrayOutput() ThreatIntelSetArrayOutput {
	return o
}

func (o ThreatIntelSetArrayOutput) ToThreatIntelSetArrayOutputWithContext(ctx context.Context) ThreatIntelSetArrayOutput {
	return o
}

func (o ThreatIntelSetArrayOutput) Index(i pulumi.IntInput) ThreatIntelSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThreatIntelSet {
		return vs[0].([]ThreatIntelSet)[vs[1].(int)]
	}).(ThreatIntelSetOutput)
}

type ThreatIntelSetMapOutput struct{ *pulumi.OutputState }

func (ThreatIntelSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ThreatIntelSet)(nil))
}

func (o ThreatIntelSetMapOutput) ToThreatIntelSetMapOutput() ThreatIntelSetMapOutput {
	return o
}

func (o ThreatIntelSetMapOutput) ToThreatIntelSetMapOutputWithContext(ctx context.Context) ThreatIntelSetMapOutput {
	return o
}

func (o ThreatIntelSetMapOutput) MapIndex(k pulumi.StringInput) ThreatIntelSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ThreatIntelSet {
		return vs[0].(map[string]ThreatIntelSet)[vs[1].(string)]
	}).(ThreatIntelSetOutput)
}

func init() {
	pulumi.RegisterOutputType(ThreatIntelSetOutput{})
	pulumi.RegisterOutputType(ThreatIntelSetPtrOutput{})
	pulumi.RegisterOutputType(ThreatIntelSetArrayOutput{})
	pulumi.RegisterOutputType(ThreatIntelSetMapOutput{})
}
