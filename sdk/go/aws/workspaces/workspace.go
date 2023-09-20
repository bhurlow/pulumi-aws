// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a workspace in [AWS Workspaces](https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces.html) Service
//
// > **NOTE:** AWS WorkSpaces service requires [`workspaces_DefaultRole`](https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role) IAM role to operate normally.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/workspaces"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			valueWindows10, err := workspaces.GetBundle(ctx, &workspaces.GetBundleArgs{
//				BundleId: pulumi.StringRef("wsb-bh8rsxt14"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = workspaces.NewWorkspace(ctx, "example", &workspaces.WorkspaceArgs{
//				DirectoryId:                 pulumi.Any(aws_workspaces_directory.Example.Id),
//				BundleId:                    *pulumi.String(valueWindows10.Id),
//				UserName:                    pulumi.String("john.doe"),
//				RootVolumeEncryptionEnabled: pulumi.Bool(true),
//				UserVolumeEncryptionEnabled: pulumi.Bool(true),
//				VolumeEncryptionKey:         pulumi.String("alias/aws/workspaces"),
//				WorkspaceProperties: &workspaces.WorkspaceWorkspacePropertiesArgs{
//					ComputeTypeName:                     pulumi.String("VALUE"),
//					UserVolumeSizeGib:                   pulumi.Int(10),
//					RootVolumeSizeGib:                   pulumi.Int(80),
//					RunningMode:                         pulumi.String("AUTO_STOP"),
//					RunningModeAutoStopTimeoutInMinutes: pulumi.Int(60),
//				},
//				Tags: pulumi.StringMap{
//					"Department": pulumi.String("IT"),
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
// Using `pulumi import`, import Workspaces using their ID. For example:
//
// ```sh
//
//	$ pulumi import aws:workspaces/workspace:Workspace example ws-9z9zmbkhv
//
// ```
type Workspace struct {
	pulumi.CustomResourceState

	// The ID of the bundle for the WorkSpace.
	BundleId pulumi.StringOutput `pulumi:"bundleId"`
	// The name of the WorkSpace, as seen by the operating system.
	ComputerName pulumi.StringOutput `pulumi:"computerName"`
	// The ID of the directory for the WorkSpace.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// The IP address of the WorkSpace.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled pulumi.BoolPtrOutput `pulumi:"rootVolumeEncryptionEnabled"`
	// The operational state of the WorkSpace.
	State pulumi.StringOutput `pulumi:"state"`
	// The tags for the WorkSpace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled pulumi.BoolPtrOutput `pulumi:"userVolumeEncryptionEnabled"`
	// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey pulumi.StringPtrOutput `pulumi:"volumeEncryptionKey"`
	// The WorkSpace properties.
	WorkspaceProperties WorkspaceWorkspacePropertiesOutput `pulumi:"workspaceProperties"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BundleId == nil {
		return nil, errors.New("invalid value for required argument 'BundleId'")
	}
	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("aws:workspaces/workspace:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("aws:workspaces/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// The ID of the bundle for the WorkSpace.
	BundleId *string `pulumi:"bundleId"`
	// The name of the WorkSpace, as seen by the operating system.
	ComputerName *string `pulumi:"computerName"`
	// The ID of the directory for the WorkSpace.
	DirectoryId *string `pulumi:"directoryId"`
	// The IP address of the WorkSpace.
	IpAddress *string `pulumi:"ipAddress"`
	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled *bool `pulumi:"rootVolumeEncryptionEnabled"`
	// The operational state of the WorkSpace.
	State *string `pulumi:"state"`
	// The tags for the WorkSpace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
	UserName *string `pulumi:"userName"`
	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled *bool `pulumi:"userVolumeEncryptionEnabled"`
	// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey *string `pulumi:"volumeEncryptionKey"`
	// The WorkSpace properties.
	WorkspaceProperties *WorkspaceWorkspaceProperties `pulumi:"workspaceProperties"`
}

type WorkspaceState struct {
	// The ID of the bundle for the WorkSpace.
	BundleId pulumi.StringPtrInput
	// The name of the WorkSpace, as seen by the operating system.
	ComputerName pulumi.StringPtrInput
	// The ID of the directory for the WorkSpace.
	DirectoryId pulumi.StringPtrInput
	// The IP address of the WorkSpace.
	IpAddress pulumi.StringPtrInput
	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled pulumi.BoolPtrInput
	// The operational state of the WorkSpace.
	State pulumi.StringPtrInput
	// The tags for the WorkSpace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
	UserName pulumi.StringPtrInput
	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled pulumi.BoolPtrInput
	// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey pulumi.StringPtrInput
	// The WorkSpace properties.
	WorkspaceProperties WorkspaceWorkspacePropertiesPtrInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The ID of the bundle for the WorkSpace.
	BundleId string `pulumi:"bundleId"`
	// The ID of the directory for the WorkSpace.
	DirectoryId string `pulumi:"directoryId"`
	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled *bool `pulumi:"rootVolumeEncryptionEnabled"`
	// The tags for the WorkSpace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
	UserName string `pulumi:"userName"`
	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled *bool `pulumi:"userVolumeEncryptionEnabled"`
	// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey *string `pulumi:"volumeEncryptionKey"`
	// The WorkSpace properties.
	WorkspaceProperties *WorkspaceWorkspaceProperties `pulumi:"workspaceProperties"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The ID of the bundle for the WorkSpace.
	BundleId pulumi.StringInput
	// The ID of the directory for the WorkSpace.
	DirectoryId pulumi.StringInput
	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled pulumi.BoolPtrInput
	// The tags for the WorkSpace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
	UserName pulumi.StringInput
	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled pulumi.BoolPtrInput
	// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey pulumi.StringPtrInput
	// The WorkSpace properties.
	WorkspaceProperties WorkspaceWorkspacePropertiesPtrInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

func (i *Workspace) ToOutput(ctx context.Context) pulumix.Output[*Workspace] {
	return pulumix.Output[*Workspace]{
		OutputState: i.ToWorkspaceOutputWithContext(ctx).OutputState,
	}
}

// WorkspaceArrayInput is an input type that accepts WorkspaceArray and WorkspaceArrayOutput values.
// You can construct a concrete instance of `WorkspaceArrayInput` via:
//
//	WorkspaceArray{ WorkspaceArgs{...} }
type WorkspaceArrayInput interface {
	pulumi.Input

	ToWorkspaceArrayOutput() WorkspaceArrayOutput
	ToWorkspaceArrayOutputWithContext(context.Context) WorkspaceArrayOutput
}

type WorkspaceArray []WorkspaceInput

func (WorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workspace)(nil)).Elem()
}

func (i WorkspaceArray) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return i.ToWorkspaceArrayOutputWithContext(context.Background())
}

func (i WorkspaceArray) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceArrayOutput)
}

func (i WorkspaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*Workspace] {
	return pulumix.Output[[]*Workspace]{
		OutputState: i.ToWorkspaceArrayOutputWithContext(ctx).OutputState,
	}
}

// WorkspaceMapInput is an input type that accepts WorkspaceMap and WorkspaceMapOutput values.
// You can construct a concrete instance of `WorkspaceMapInput` via:
//
//	WorkspaceMap{ "key": WorkspaceArgs{...} }
type WorkspaceMapInput interface {
	pulumi.Input

	ToWorkspaceMapOutput() WorkspaceMapOutput
	ToWorkspaceMapOutputWithContext(context.Context) WorkspaceMapOutput
}

type WorkspaceMap map[string]WorkspaceInput

func (WorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workspace)(nil)).Elem()
}

func (i WorkspaceMap) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return i.ToWorkspaceMapOutputWithContext(context.Background())
}

func (i WorkspaceMap) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceMapOutput)
}

func (i WorkspaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workspace] {
	return pulumix.Output[map[string]*Workspace]{
		OutputState: i.ToWorkspaceMapOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToOutput(ctx context.Context) pulumix.Output[*Workspace] {
	return pulumix.Output[*Workspace]{
		OutputState: o.OutputState,
	}
}

// The ID of the bundle for the WorkSpace.
func (o WorkspaceOutput) BundleId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.BundleId }).(pulumi.StringOutput)
}

// The name of the WorkSpace, as seen by the operating system.
func (o WorkspaceOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ComputerName }).(pulumi.StringOutput)
}

// The ID of the directory for the WorkSpace.
func (o WorkspaceOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// The IP address of the WorkSpace.
func (o WorkspaceOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Indicates whether the data stored on the root volume is encrypted.
func (o WorkspaceOutput) RootVolumeEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.RootVolumeEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

// The operational state of the WorkSpace.
func (o WorkspaceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The tags for the WorkSpace. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o WorkspaceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
func (o WorkspaceOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// Indicates whether the data stored on the user volume is encrypted.
func (o WorkspaceOutput) UserVolumeEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.UserVolumeEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
func (o WorkspaceOutput) VolumeEncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.VolumeEncryptionKey }).(pulumi.StringPtrOutput)
}

// The WorkSpace properties.
func (o WorkspaceOutput) WorkspaceProperties() WorkspaceWorkspacePropertiesOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceWorkspacePropertiesOutput { return v.WorkspaceProperties }).(WorkspaceWorkspacePropertiesOutput)
}

type WorkspaceArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workspace)(nil)).Elem()
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Workspace] {
	return pulumix.Output[[]*Workspace]{
		OutputState: o.OutputState,
	}
}

func (o WorkspaceArrayOutput) Index(i pulumi.IntInput) WorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workspace {
		return vs[0].([]*Workspace)[vs[1].(int)]
	}).(WorkspaceOutput)
}

type WorkspaceMapOutput struct{ *pulumi.OutputState }

func (WorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workspace)(nil)).Elem()
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workspace] {
	return pulumix.Output[map[string]*Workspace]{
		OutputState: o.OutputState,
	}
}

func (o WorkspaceMapOutput) MapIndex(k pulumi.StringInput) WorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workspace {
		return vs[0].(map[string]*Workspace)[vs[1].(string)]
	}).(WorkspaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceInput)(nil)).Elem(), &Workspace{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceArrayInput)(nil)).Elem(), WorkspaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceMapInput)(nil)).Elem(), WorkspaceMap{})
	pulumi.RegisterOutputType(WorkspaceOutput{})
	pulumi.RegisterOutputType(WorkspaceArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceMapOutput{})
}
