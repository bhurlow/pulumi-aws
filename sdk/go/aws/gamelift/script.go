// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an GameLift Script resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/gamelift"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gamelift.NewScript(ctx, "example", &gamelift.ScriptArgs{
//				StorageLocation: &gamelift.ScriptStorageLocationArgs{
//					Bucket:  pulumi.Any(aws_s3_bucket.Example.Id),
//					Key:     pulumi.Any(aws_s3_object.Example.Key),
//					RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
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
// GameLift Scripts can be imported using the ID, e.g.,
//
// ```sh
//
//	$ pulumi import aws:gamelift/script:Script example <script-id>
//
// ```
type Script struct {
	pulumi.CustomResourceState

	// GameLift Script ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the script
	Name pulumi.StringOutput `pulumi:"name"`
	// Information indicating where your game script files are stored. See below.
	StorageLocation ScriptStorageLocationOutput `pulumi:"storageLocation"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Version that is associated with this script.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	ZipFile pulumi.StringPtrOutput `pulumi:"zipFile"`
}

// NewScript registers a new resource with the given unique name, arguments, and options.
func NewScript(ctx *pulumi.Context,
	name string, args *ScriptArgs, opts ...pulumi.ResourceOption) (*Script, error) {
	if args == nil {
		args = &ScriptArgs{}
	}

	var resource Script
	err := ctx.RegisterResource("aws:gamelift/script:Script", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScript gets an existing Script resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptState, opts ...pulumi.ResourceOption) (*Script, error) {
	var resource Script
	err := ctx.ReadResource("aws:gamelift/script:Script", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Script resources.
type scriptState struct {
	// GameLift Script ARN.
	Arn *string `pulumi:"arn"`
	// Name of the script
	Name *string `pulumi:"name"`
	// Information indicating where your game script files are stored. See below.
	StorageLocation *ScriptStorageLocation `pulumi:"storageLocation"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Version that is associated with this script.
	Version *string `pulumi:"version"`
	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	ZipFile *string `pulumi:"zipFile"`
}

type ScriptState struct {
	// GameLift Script ARN.
	Arn pulumi.StringPtrInput
	// Name of the script
	Name pulumi.StringPtrInput
	// Information indicating where your game script files are stored. See below.
	StorageLocation ScriptStorageLocationPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Version that is associated with this script.
	Version pulumi.StringPtrInput
	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	ZipFile pulumi.StringPtrInput
}

func (ScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptState)(nil)).Elem()
}

type scriptArgs struct {
	// Name of the script
	Name *string `pulumi:"name"`
	// Information indicating where your game script files are stored. See below.
	StorageLocation *ScriptStorageLocation `pulumi:"storageLocation"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Version that is associated with this script.
	Version *string `pulumi:"version"`
	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	ZipFile *string `pulumi:"zipFile"`
}

// The set of arguments for constructing a Script resource.
type ScriptArgs struct {
	// Name of the script
	Name pulumi.StringPtrInput
	// Information indicating where your game script files are stored. See below.
	StorageLocation ScriptStorageLocationPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Version that is associated with this script.
	Version pulumi.StringPtrInput
	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	ZipFile pulumi.StringPtrInput
}

func (ScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptArgs)(nil)).Elem()
}

type ScriptInput interface {
	pulumi.Input

	ToScriptOutput() ScriptOutput
	ToScriptOutputWithContext(ctx context.Context) ScriptOutput
}

func (*Script) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (i *Script) ToScriptOutput() ScriptOutput {
	return i.ToScriptOutputWithContext(context.Background())
}

func (i *Script) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptOutput)
}

// ScriptArrayInput is an input type that accepts ScriptArray and ScriptArrayOutput values.
// You can construct a concrete instance of `ScriptArrayInput` via:
//
//	ScriptArray{ ScriptArgs{...} }
type ScriptArrayInput interface {
	pulumi.Input

	ToScriptArrayOutput() ScriptArrayOutput
	ToScriptArrayOutputWithContext(context.Context) ScriptArrayOutput
}

type ScriptArray []ScriptInput

func (ScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Script)(nil)).Elem()
}

func (i ScriptArray) ToScriptArrayOutput() ScriptArrayOutput {
	return i.ToScriptArrayOutputWithContext(context.Background())
}

func (i ScriptArray) ToScriptArrayOutputWithContext(ctx context.Context) ScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptArrayOutput)
}

// ScriptMapInput is an input type that accepts ScriptMap and ScriptMapOutput values.
// You can construct a concrete instance of `ScriptMapInput` via:
//
//	ScriptMap{ "key": ScriptArgs{...} }
type ScriptMapInput interface {
	pulumi.Input

	ToScriptMapOutput() ScriptMapOutput
	ToScriptMapOutputWithContext(context.Context) ScriptMapOutput
}

type ScriptMap map[string]ScriptInput

func (ScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Script)(nil)).Elem()
}

func (i ScriptMap) ToScriptMapOutput() ScriptMapOutput {
	return i.ToScriptMapOutputWithContext(context.Background())
}

func (i ScriptMap) ToScriptMapOutputWithContext(ctx context.Context) ScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptMapOutput)
}

type ScriptOutput struct{ *pulumi.OutputState }

func (ScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (o ScriptOutput) ToScriptOutput() ScriptOutput {
	return o
}

func (o ScriptOutput) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return o
}

// GameLift Script ARN.
func (o ScriptOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of the script
func (o ScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Information indicating where your game script files are stored. See below.
func (o ScriptOutput) StorageLocation() ScriptStorageLocationOutput {
	return o.ApplyT(func(v *Script) ScriptStorageLocationOutput { return v.StorageLocation }).(ScriptStorageLocationOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ScriptOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Script) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ScriptOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Script) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Version that is associated with this script.
func (o ScriptOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
func (o ScriptOutput) ZipFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.StringPtrOutput { return v.ZipFile }).(pulumi.StringPtrOutput)
}

type ScriptArrayOutput struct{ *pulumi.OutputState }

func (ScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Script)(nil)).Elem()
}

func (o ScriptArrayOutput) ToScriptArrayOutput() ScriptArrayOutput {
	return o
}

func (o ScriptArrayOutput) ToScriptArrayOutputWithContext(ctx context.Context) ScriptArrayOutput {
	return o
}

func (o ScriptArrayOutput) Index(i pulumi.IntInput) ScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Script {
		return vs[0].([]*Script)[vs[1].(int)]
	}).(ScriptOutput)
}

type ScriptMapOutput struct{ *pulumi.OutputState }

func (ScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Script)(nil)).Elem()
}

func (o ScriptMapOutput) ToScriptMapOutput() ScriptMapOutput {
	return o
}

func (o ScriptMapOutput) ToScriptMapOutputWithContext(ctx context.Context) ScriptMapOutput {
	return o
}

func (o ScriptMapOutput) MapIndex(k pulumi.StringInput) ScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Script {
		return vs[0].(map[string]*Script)[vs[1].(string)]
	}).(ScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptInput)(nil)).Elem(), &Script{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptArrayInput)(nil)).Elem(), ScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptMapInput)(nil)).Elem(), ScriptMap{})
	pulumi.RegisterOutputType(ScriptOutput{})
	pulumi.RegisterOutputType(ScriptArrayOutput{})
	pulumi.RegisterOutputType(ScriptMapOutput{})
}
