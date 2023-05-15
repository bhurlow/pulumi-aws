// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS Location Geofence Collection.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/location"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := location.NewGeofenceCollection(ctx, "example", &location.GeofenceCollectionArgs{
//				CollectionName: pulumi.String("example"),
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
// Location Geofence Collection can be imported using the `collection_name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:location/geofenceCollection:GeofenceCollection example example
//
// ```
type GeofenceCollection struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS.
	CollectionArn pulumi.StringOutput `pulumi:"collectionArn"`
	// The name of the geofence collection.
	CollectionName pulumi.StringOutput `pulumi:"collectionName"`
	// The timestamp for when the geofence collection resource was created in ISO 8601 format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The optional description for the geofence collection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Key-value tags for the geofence collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The timestamp for when the geofence collection resource was last updated in ISO 8601 format.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGeofenceCollection registers a new resource with the given unique name, arguments, and options.
func NewGeofenceCollection(ctx *pulumi.Context,
	name string, args *GeofenceCollectionArgs, opts ...pulumi.ResourceOption) (*GeofenceCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectionName == nil {
		return nil, errors.New("invalid value for required argument 'CollectionName'")
	}
	var resource GeofenceCollection
	err := ctx.RegisterResource("aws:location/geofenceCollection:GeofenceCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeofenceCollection gets an existing GeofenceCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeofenceCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeofenceCollectionState, opts ...pulumi.ResourceOption) (*GeofenceCollection, error) {
	var resource GeofenceCollection
	err := ctx.ReadResource("aws:location/geofenceCollection:GeofenceCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeofenceCollection resources.
type geofenceCollectionState struct {
	// The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS.
	CollectionArn *string `pulumi:"collectionArn"`
	// The name of the geofence collection.
	CollectionName *string `pulumi:"collectionName"`
	// The timestamp for when the geofence collection resource was created in ISO 8601 format.
	CreateTime *string `pulumi:"createTime"`
	// The optional description for the geofence collection.
	Description *string `pulumi:"description"`
	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Key-value tags for the geofence collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The timestamp for when the geofence collection resource was last updated in ISO 8601 format.
	UpdateTime *string `pulumi:"updateTime"`
}

type GeofenceCollectionState struct {
	// The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS.
	CollectionArn pulumi.StringPtrInput
	// The name of the geofence collection.
	CollectionName pulumi.StringPtrInput
	// The timestamp for when the geofence collection resource was created in ISO 8601 format.
	CreateTime pulumi.StringPtrInput
	// The optional description for the geofence collection.
	Description pulumi.StringPtrInput
	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KmsKeyId pulumi.StringPtrInput
	// Key-value tags for the geofence collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// The timestamp for when the geofence collection resource was last updated in ISO 8601 format.
	UpdateTime pulumi.StringPtrInput
}

func (GeofenceCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*geofenceCollectionState)(nil)).Elem()
}

type geofenceCollectionArgs struct {
	// The name of the geofence collection.
	CollectionName string `pulumi:"collectionName"`
	// The optional description for the geofence collection.
	Description *string `pulumi:"description"`
	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Key-value tags for the geofence collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GeofenceCollection resource.
type GeofenceCollectionArgs struct {
	// The name of the geofence collection.
	CollectionName pulumi.StringInput
	// The optional description for the geofence collection.
	Description pulumi.StringPtrInput
	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KmsKeyId pulumi.StringPtrInput
	// Key-value tags for the geofence collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (GeofenceCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geofenceCollectionArgs)(nil)).Elem()
}

type GeofenceCollectionInput interface {
	pulumi.Input

	ToGeofenceCollectionOutput() GeofenceCollectionOutput
	ToGeofenceCollectionOutputWithContext(ctx context.Context) GeofenceCollectionOutput
}

func (*GeofenceCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**GeofenceCollection)(nil)).Elem()
}

func (i *GeofenceCollection) ToGeofenceCollectionOutput() GeofenceCollectionOutput {
	return i.ToGeofenceCollectionOutputWithContext(context.Background())
}

func (i *GeofenceCollection) ToGeofenceCollectionOutputWithContext(ctx context.Context) GeofenceCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeofenceCollectionOutput)
}

// GeofenceCollectionArrayInput is an input type that accepts GeofenceCollectionArray and GeofenceCollectionArrayOutput values.
// You can construct a concrete instance of `GeofenceCollectionArrayInput` via:
//
//	GeofenceCollectionArray{ GeofenceCollectionArgs{...} }
type GeofenceCollectionArrayInput interface {
	pulumi.Input

	ToGeofenceCollectionArrayOutput() GeofenceCollectionArrayOutput
	ToGeofenceCollectionArrayOutputWithContext(context.Context) GeofenceCollectionArrayOutput
}

type GeofenceCollectionArray []GeofenceCollectionInput

func (GeofenceCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GeofenceCollection)(nil)).Elem()
}

func (i GeofenceCollectionArray) ToGeofenceCollectionArrayOutput() GeofenceCollectionArrayOutput {
	return i.ToGeofenceCollectionArrayOutputWithContext(context.Background())
}

func (i GeofenceCollectionArray) ToGeofenceCollectionArrayOutputWithContext(ctx context.Context) GeofenceCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeofenceCollectionArrayOutput)
}

// GeofenceCollectionMapInput is an input type that accepts GeofenceCollectionMap and GeofenceCollectionMapOutput values.
// You can construct a concrete instance of `GeofenceCollectionMapInput` via:
//
//	GeofenceCollectionMap{ "key": GeofenceCollectionArgs{...} }
type GeofenceCollectionMapInput interface {
	pulumi.Input

	ToGeofenceCollectionMapOutput() GeofenceCollectionMapOutput
	ToGeofenceCollectionMapOutputWithContext(context.Context) GeofenceCollectionMapOutput
}

type GeofenceCollectionMap map[string]GeofenceCollectionInput

func (GeofenceCollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GeofenceCollection)(nil)).Elem()
}

func (i GeofenceCollectionMap) ToGeofenceCollectionMapOutput() GeofenceCollectionMapOutput {
	return i.ToGeofenceCollectionMapOutputWithContext(context.Background())
}

func (i GeofenceCollectionMap) ToGeofenceCollectionMapOutputWithContext(ctx context.Context) GeofenceCollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeofenceCollectionMapOutput)
}

type GeofenceCollectionOutput struct{ *pulumi.OutputState }

func (GeofenceCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeofenceCollection)(nil)).Elem()
}

func (o GeofenceCollectionOutput) ToGeofenceCollectionOutput() GeofenceCollectionOutput {
	return o
}

func (o GeofenceCollectionOutput) ToGeofenceCollectionOutputWithContext(ctx context.Context) GeofenceCollectionOutput {
	return o
}

// The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS.
func (o GeofenceCollectionOutput) CollectionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringOutput { return v.CollectionArn }).(pulumi.StringOutput)
}

// The name of the geofence collection.
func (o GeofenceCollectionOutput) CollectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringOutput { return v.CollectionName }).(pulumi.StringOutput)
}

// The timestamp for when the geofence collection resource was created in ISO 8601 format.
func (o GeofenceCollectionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The optional description for the geofence collection.
func (o GeofenceCollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
func (o GeofenceCollectionOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Key-value tags for the geofence collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o GeofenceCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GeofenceCollectionOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The timestamp for when the geofence collection resource was last updated in ISO 8601 format.
func (o GeofenceCollectionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GeofenceCollection) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type GeofenceCollectionArrayOutput struct{ *pulumi.OutputState }

func (GeofenceCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GeofenceCollection)(nil)).Elem()
}

func (o GeofenceCollectionArrayOutput) ToGeofenceCollectionArrayOutput() GeofenceCollectionArrayOutput {
	return o
}

func (o GeofenceCollectionArrayOutput) ToGeofenceCollectionArrayOutputWithContext(ctx context.Context) GeofenceCollectionArrayOutput {
	return o
}

func (o GeofenceCollectionArrayOutput) Index(i pulumi.IntInput) GeofenceCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GeofenceCollection {
		return vs[0].([]*GeofenceCollection)[vs[1].(int)]
	}).(GeofenceCollectionOutput)
}

type GeofenceCollectionMapOutput struct{ *pulumi.OutputState }

func (GeofenceCollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GeofenceCollection)(nil)).Elem()
}

func (o GeofenceCollectionMapOutput) ToGeofenceCollectionMapOutput() GeofenceCollectionMapOutput {
	return o
}

func (o GeofenceCollectionMapOutput) ToGeofenceCollectionMapOutputWithContext(ctx context.Context) GeofenceCollectionMapOutput {
	return o
}

func (o GeofenceCollectionMapOutput) MapIndex(k pulumi.StringInput) GeofenceCollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GeofenceCollection {
		return vs[0].(map[string]*GeofenceCollection)[vs[1].(string)]
	}).(GeofenceCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GeofenceCollectionInput)(nil)).Elem(), &GeofenceCollection{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeofenceCollectionArrayInput)(nil)).Elem(), GeofenceCollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeofenceCollectionMapInput)(nil)).Elem(), GeofenceCollectionMap{})
	pulumi.RegisterOutputType(GeofenceCollectionOutput{})
	pulumi.RegisterOutputType(GeofenceCollectionArrayOutput{})
	pulumi.RegisterOutputType(GeofenceCollectionMapOutput{})
}
