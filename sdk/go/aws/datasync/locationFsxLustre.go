// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS DataSync FSx Lustre Location.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datasync"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datasync.NewLocationFsxLustre(ctx, "example", &datasync.LocationFsxLustreArgs{
//				FsxFilesystemArn: pulumi.Any(aws_fsx_lustre_file_system.Example.Arn),
//				SecurityGroupArns: pulumi.StringArray{
//					aws_security_group.Example.Arn,
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
// terraform import {
//
//	to = aws_datasync_location_fsx_lustre_file_system.example
//
//	id = "arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:476956259333:file-system/fs-08e04cd442c1bb94a" } Using `pulumi import`, import `aws_datasync_location_fsx_lustre_file_system` using the `DataSync-ARN#FSx-Lustre-ARN`. For exampleconsole % pulumi import aws_datasync_location_fsx_lustre_file_system.example arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:476956259333:file-system/fs-08e04cd442c1bb94a
type LocationFsxLustre struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The time that the FSx for Lustre location was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn pulumi.StringOutput `pulumi:"fsxFilesystemArn"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
	SecurityGroupArns pulumi.StringArrayOutput `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The URL of the FSx for Lustre location that was described.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewLocationFsxLustre registers a new resource with the given unique name, arguments, and options.
func NewLocationFsxLustre(ctx *pulumi.Context,
	name string, args *LocationFsxLustreArgs, opts ...pulumi.ResourceOption) (*LocationFsxLustre, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FsxFilesystemArn == nil {
		return nil, errors.New("invalid value for required argument 'FsxFilesystemArn'")
	}
	if args.SecurityGroupArns == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupArns'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocationFsxLustre
	err := ctx.RegisterResource("aws:datasync/locationFsxLustre:LocationFsxLustre", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationFsxLustre gets an existing LocationFsxLustre resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationFsxLustre(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationFsxLustreState, opts ...pulumi.ResourceOption) (*LocationFsxLustre, error) {
	var resource LocationFsxLustre
	err := ctx.ReadResource("aws:datasync/locationFsxLustre:LocationFsxLustre", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationFsxLustre resources.
type locationFsxLustreState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// The time that the FSx for Lustre location was created.
	CreationTime *string `pulumi:"creationTime"`
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn *string `pulumi:"fsxFilesystemArn"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The URL of the FSx for Lustre location that was described.
	Uri *string `pulumi:"uri"`
}

type LocationFsxLustreState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// The time that the FSx for Lustre location was created.
	CreationTime pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn pulumi.StringPtrInput
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
	SecurityGroupArns pulumi.StringArrayInput
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The URL of the FSx for Lustre location that was described.
	Uri pulumi.StringPtrInput
}

func (LocationFsxLustreState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFsxLustreState)(nil)).Elem()
}

type locationFsxLustreArgs struct {
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn string `pulumi:"fsxFilesystemArn"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LocationFsxLustre resource.
type LocationFsxLustreArgs struct {
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn pulumi.StringInput
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
	SecurityGroupArns pulumi.StringArrayInput
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (LocationFsxLustreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFsxLustreArgs)(nil)).Elem()
}

type LocationFsxLustreInput interface {
	pulumi.Input

	ToLocationFsxLustreOutput() LocationFsxLustreOutput
	ToLocationFsxLustreOutputWithContext(ctx context.Context) LocationFsxLustreOutput
}

func (*LocationFsxLustre) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationFsxLustre)(nil)).Elem()
}

func (i *LocationFsxLustre) ToLocationFsxLustreOutput() LocationFsxLustreOutput {
	return i.ToLocationFsxLustreOutputWithContext(context.Background())
}

func (i *LocationFsxLustre) ToLocationFsxLustreOutputWithContext(ctx context.Context) LocationFsxLustreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxLustreOutput)
}

// LocationFsxLustreArrayInput is an input type that accepts LocationFsxLustreArray and LocationFsxLustreArrayOutput values.
// You can construct a concrete instance of `LocationFsxLustreArrayInput` via:
//
//	LocationFsxLustreArray{ LocationFsxLustreArgs{...} }
type LocationFsxLustreArrayInput interface {
	pulumi.Input

	ToLocationFsxLustreArrayOutput() LocationFsxLustreArrayOutput
	ToLocationFsxLustreArrayOutputWithContext(context.Context) LocationFsxLustreArrayOutput
}

type LocationFsxLustreArray []LocationFsxLustreInput

func (LocationFsxLustreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocationFsxLustre)(nil)).Elem()
}

func (i LocationFsxLustreArray) ToLocationFsxLustreArrayOutput() LocationFsxLustreArrayOutput {
	return i.ToLocationFsxLustreArrayOutputWithContext(context.Background())
}

func (i LocationFsxLustreArray) ToLocationFsxLustreArrayOutputWithContext(ctx context.Context) LocationFsxLustreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxLustreArrayOutput)
}

// LocationFsxLustreMapInput is an input type that accepts LocationFsxLustreMap and LocationFsxLustreMapOutput values.
// You can construct a concrete instance of `LocationFsxLustreMapInput` via:
//
//	LocationFsxLustreMap{ "key": LocationFsxLustreArgs{...} }
type LocationFsxLustreMapInput interface {
	pulumi.Input

	ToLocationFsxLustreMapOutput() LocationFsxLustreMapOutput
	ToLocationFsxLustreMapOutputWithContext(context.Context) LocationFsxLustreMapOutput
}

type LocationFsxLustreMap map[string]LocationFsxLustreInput

func (LocationFsxLustreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocationFsxLustre)(nil)).Elem()
}

func (i LocationFsxLustreMap) ToLocationFsxLustreMapOutput() LocationFsxLustreMapOutput {
	return i.ToLocationFsxLustreMapOutputWithContext(context.Background())
}

func (i LocationFsxLustreMap) ToLocationFsxLustreMapOutputWithContext(ctx context.Context) LocationFsxLustreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxLustreMapOutput)
}

type LocationFsxLustreOutput struct{ *pulumi.OutputState }

func (LocationFsxLustreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationFsxLustre)(nil)).Elem()
}

func (o LocationFsxLustreOutput) ToLocationFsxLustreOutput() LocationFsxLustreOutput {
	return o
}

func (o LocationFsxLustreOutput) ToLocationFsxLustreOutputWithContext(ctx context.Context) LocationFsxLustreOutput {
	return o
}

// Amazon Resource Name (ARN) of the DataSync Location.
func (o LocationFsxLustreOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The time that the FSx for Lustre location was created.
func (o LocationFsxLustreOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
func (o LocationFsxLustreOutput) FsxFilesystemArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringOutput { return v.FsxFilesystemArn }).(pulumi.StringOutput)
}

// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
func (o LocationFsxLustreOutput) SecurityGroupArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringArrayOutput { return v.SecurityGroupArns }).(pulumi.StringArrayOutput)
}

// Subdirectory to perform actions as source or destination.
func (o LocationFsxLustreOutput) Subdirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringOutput { return v.Subdirectory }).(pulumi.StringOutput)
}

// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o LocationFsxLustreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o LocationFsxLustreOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The URL of the FSx for Lustre location that was described.
func (o LocationFsxLustreOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationFsxLustre) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

type LocationFsxLustreArrayOutput struct{ *pulumi.OutputState }

func (LocationFsxLustreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocationFsxLustre)(nil)).Elem()
}

func (o LocationFsxLustreArrayOutput) ToLocationFsxLustreArrayOutput() LocationFsxLustreArrayOutput {
	return o
}

func (o LocationFsxLustreArrayOutput) ToLocationFsxLustreArrayOutputWithContext(ctx context.Context) LocationFsxLustreArrayOutput {
	return o
}

func (o LocationFsxLustreArrayOutput) Index(i pulumi.IntInput) LocationFsxLustreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocationFsxLustre {
		return vs[0].([]*LocationFsxLustre)[vs[1].(int)]
	}).(LocationFsxLustreOutput)
}

type LocationFsxLustreMapOutput struct{ *pulumi.OutputState }

func (LocationFsxLustreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocationFsxLustre)(nil)).Elem()
}

func (o LocationFsxLustreMapOutput) ToLocationFsxLustreMapOutput() LocationFsxLustreMapOutput {
	return o
}

func (o LocationFsxLustreMapOutput) ToLocationFsxLustreMapOutputWithContext(ctx context.Context) LocationFsxLustreMapOutput {
	return o
}

func (o LocationFsxLustreMapOutput) MapIndex(k pulumi.StringInput) LocationFsxLustreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocationFsxLustre {
		return vs[0].(map[string]*LocationFsxLustre)[vs[1].(string)]
	}).(LocationFsxLustreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFsxLustreInput)(nil)).Elem(), &LocationFsxLustre{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFsxLustreArrayInput)(nil)).Elem(), LocationFsxLustreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFsxLustreMapInput)(nil)).Elem(), LocationFsxLustreMap{})
	pulumi.RegisterOutputType(LocationFsxLustreOutput{})
	pulumi.RegisterOutputType(LocationFsxLustreArrayOutput{})
	pulumi.RegisterOutputType(LocationFsxLustreMapOutput{})
}
