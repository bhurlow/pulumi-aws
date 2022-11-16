// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AppStream stack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/appstream"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appstream.NewStack(ctx, "example", &appstream.StackArgs{
//				ApplicationSettings: &appstream.StackApplicationSettingsArgs{
//					Enabled:       pulumi.Bool(true),
//					SettingsGroup: pulumi.String("SettingsGroup"),
//				},
//				Description: pulumi.String("stack description"),
//				DisplayName: pulumi.String("stack display name"),
//				FeedbackUrl: pulumi.String("http://your-domain/feedback"),
//				RedirectUrl: pulumi.String("http://your-domain/redirect"),
//				StorageConnectors: appstream.StackStorageConnectorArray{
//					&appstream.StackStorageConnectorArgs{
//						ConnectorType: pulumi.String("HOMEFOLDERS"),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"TagName": pulumi.String("TagValue"),
//				},
//				UserSettings: appstream.StackUserSettingArray{
//					&appstream.StackUserSettingArgs{
//						Action:     pulumi.String("CLIPBOARD_COPY_FROM_LOCAL_DEVICE"),
//						Permission: pulumi.String("ENABLED"),
//					},
//					&appstream.StackUserSettingArgs{
//						Action:     pulumi.String("CLIPBOARD_COPY_TO_LOCAL_DEVICE"),
//						Permission: pulumi.String("ENABLED"),
//					},
//					&appstream.StackUserSettingArgs{
//						Action:     pulumi.String("FILE_UPLOAD"),
//						Permission: pulumi.String("ENABLED"),
//					},
//					&appstream.StackUserSettingArgs{
//						Action:     pulumi.String("FILE_DOWNLOAD"),
//						Permission: pulumi.String("ENABLED"),
//					},
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
// `aws_appstream_stack` can be imported using the id, e.g.,
//
// ```sh
//
//	$ pulumi import aws:appstream/stack:Stack example stackID
//
// ```
type Stack struct {
	pulumi.CustomResourceState

	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See `accessEndpoints` below.
	AccessEndpoints StackAccessEndpointArrayOutput `pulumi:"accessEndpoints"`
	// Settings for application settings persistence.
	// See `applicationSettings` below.
	ApplicationSettings StackApplicationSettingsOutput `pulumi:"applicationSettings"`
	// ARN of the appstream stack.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Date and time, in UTC and extended RFC 3339 format, when the stack was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Description for the AppStream stack.
	Description pulumi.StringOutput `pulumi:"description"`
	// Stack name to display.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	EmbedHostDomains pulumi.StringArrayOutput `pulumi:"embedHostDomains"`
	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackUrl pulumi.StringOutput `pulumi:"feedbackUrl"`
	// Unique name for the AppStream stack.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL that users are redirected to after their streaming session ends.
	RedirectUrl pulumi.StringOutput `pulumi:"redirectUrl"`
	// Configuration block for the storage connectors to enable.
	// See `storageConnectors` below.
	StorageConnectors StackStorageConnectorArrayOutput `pulumi:"storageConnectors"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
	// See `userSettings` below.
	UserSettings StackUserSettingArrayOutput `pulumi:"userSettings"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil {
		args = &StackArgs{}
	}

	var resource Stack
	err := ctx.RegisterResource("aws:appstream/stack:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("aws:appstream/stack:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See `accessEndpoints` below.
	AccessEndpoints []StackAccessEndpoint `pulumi:"accessEndpoints"`
	// Settings for application settings persistence.
	// See `applicationSettings` below.
	ApplicationSettings *StackApplicationSettings `pulumi:"applicationSettings"`
	// ARN of the appstream stack.
	Arn *string `pulumi:"arn"`
	// Date and time, in UTC and extended RFC 3339 format, when the stack was created.
	CreatedTime *string `pulumi:"createdTime"`
	// Description for the AppStream stack.
	Description *string `pulumi:"description"`
	// Stack name to display.
	DisplayName *string `pulumi:"displayName"`
	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	EmbedHostDomains []string `pulumi:"embedHostDomains"`
	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackUrl *string `pulumi:"feedbackUrl"`
	// Unique name for the AppStream stack.
	Name *string `pulumi:"name"`
	// URL that users are redirected to after their streaming session ends.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// Configuration block for the storage connectors to enable.
	// See `storageConnectors` below.
	StorageConnectors []StackStorageConnector `pulumi:"storageConnectors"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
	// See `userSettings` below.
	UserSettings []StackUserSetting `pulumi:"userSettings"`
}

type StackState struct {
	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See `accessEndpoints` below.
	AccessEndpoints StackAccessEndpointArrayInput
	// Settings for application settings persistence.
	// See `applicationSettings` below.
	ApplicationSettings StackApplicationSettingsPtrInput
	// ARN of the appstream stack.
	Arn pulumi.StringPtrInput
	// Date and time, in UTC and extended RFC 3339 format, when the stack was created.
	CreatedTime pulumi.StringPtrInput
	// Description for the AppStream stack.
	Description pulumi.StringPtrInput
	// Stack name to display.
	DisplayName pulumi.StringPtrInput
	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	EmbedHostDomains pulumi.StringArrayInput
	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackUrl pulumi.StringPtrInput
	// Unique name for the AppStream stack.
	Name pulumi.StringPtrInput
	// URL that users are redirected to after their streaming session ends.
	RedirectUrl pulumi.StringPtrInput
	// Configuration block for the storage connectors to enable.
	// See `storageConnectors` below.
	StorageConnectors StackStorageConnectorArrayInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
	// See `userSettings` below.
	UserSettings StackUserSettingArrayInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See `accessEndpoints` below.
	AccessEndpoints []StackAccessEndpoint `pulumi:"accessEndpoints"`
	// Settings for application settings persistence.
	// See `applicationSettings` below.
	ApplicationSettings *StackApplicationSettings `pulumi:"applicationSettings"`
	// Description for the AppStream stack.
	Description *string `pulumi:"description"`
	// Stack name to display.
	DisplayName *string `pulumi:"displayName"`
	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	EmbedHostDomains []string `pulumi:"embedHostDomains"`
	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackUrl *string `pulumi:"feedbackUrl"`
	// Unique name for the AppStream stack.
	Name *string `pulumi:"name"`
	// URL that users are redirected to after their streaming session ends.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// Configuration block for the storage connectors to enable.
	// See `storageConnectors` below.
	StorageConnectors []StackStorageConnector `pulumi:"storageConnectors"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
	// See `userSettings` below.
	UserSettings []StackUserSetting `pulumi:"userSettings"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	// See `accessEndpoints` below.
	AccessEndpoints StackAccessEndpointArrayInput
	// Settings for application settings persistence.
	// See `applicationSettings` below.
	ApplicationSettings StackApplicationSettingsPtrInput
	// Description for the AppStream stack.
	Description pulumi.StringPtrInput
	// Stack name to display.
	DisplayName pulumi.StringPtrInput
	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	EmbedHostDomains pulumi.StringArrayInput
	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackUrl pulumi.StringPtrInput
	// Unique name for the AppStream stack.
	Name pulumi.StringPtrInput
	// URL that users are redirected to after their streaming session ends.
	RedirectUrl pulumi.StringPtrInput
	// Configuration block for the storage connectors to enable.
	// See `storageConnectors` below.
	StorageConnectors StackStorageConnectorArrayInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
	// See `userSettings` below.
	UserSettings StackUserSettingArrayInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}

type StackInput interface {
	pulumi.Input

	ToStackOutput() StackOutput
	ToStackOutputWithContext(ctx context.Context) StackOutput
}

func (*Stack) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil)).Elem()
}

func (i *Stack) ToStackOutput() StackOutput {
	return i.ToStackOutputWithContext(context.Background())
}

func (i *Stack) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackOutput)
}

// StackArrayInput is an input type that accepts StackArray and StackArrayOutput values.
// You can construct a concrete instance of `StackArrayInput` via:
//
//	StackArray{ StackArgs{...} }
type StackArrayInput interface {
	pulumi.Input

	ToStackArrayOutput() StackArrayOutput
	ToStackArrayOutputWithContext(context.Context) StackArrayOutput
}

type StackArray []StackInput

func (StackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stack)(nil)).Elem()
}

func (i StackArray) ToStackArrayOutput() StackArrayOutput {
	return i.ToStackArrayOutputWithContext(context.Background())
}

func (i StackArray) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackArrayOutput)
}

// StackMapInput is an input type that accepts StackMap and StackMapOutput values.
// You can construct a concrete instance of `StackMapInput` via:
//
//	StackMap{ "key": StackArgs{...} }
type StackMapInput interface {
	pulumi.Input

	ToStackMapOutput() StackMapOutput
	ToStackMapOutputWithContext(context.Context) StackMapOutput
}

type StackMap map[string]StackInput

func (StackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stack)(nil)).Elem()
}

func (i StackMap) ToStackMapOutput() StackMapOutput {
	return i.ToStackMapOutputWithContext(context.Background())
}

func (i StackMap) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackMapOutput)
}

type StackOutput struct{ *pulumi.OutputState }

func (StackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil)).Elem()
}

func (o StackOutput) ToStackOutput() StackOutput {
	return o
}

func (o StackOutput) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return o
}

// Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
// See `accessEndpoints` below.
func (o StackOutput) AccessEndpoints() StackAccessEndpointArrayOutput {
	return o.ApplyT(func(v *Stack) StackAccessEndpointArrayOutput { return v.AccessEndpoints }).(StackAccessEndpointArrayOutput)
}

// Settings for application settings persistence.
// See `applicationSettings` below.
func (o StackOutput) ApplicationSettings() StackApplicationSettingsOutput {
	return o.ApplyT(func(v *Stack) StackApplicationSettingsOutput { return v.ApplicationSettings }).(StackApplicationSettingsOutput)
}

// ARN of the appstream stack.
func (o StackOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Date and time, in UTC and extended RFC 3339 format, when the stack was created.
func (o StackOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Description for the AppStream stack.
func (o StackOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Stack name to display.
func (o StackOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
func (o StackOutput) EmbedHostDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringArrayOutput { return v.EmbedHostDomains }).(pulumi.StringArrayOutput)
}

// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
func (o StackOutput) FeedbackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.FeedbackUrl }).(pulumi.StringOutput)
}

// Unique name for the AppStream stack.
func (o StackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// URL that users are redirected to after their streaming session ends.
func (o StackOutput) RedirectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringOutput { return v.RedirectUrl }).(pulumi.StringOutput)
}

// Configuration block for the storage connectors to enable.
// See `storageConnectors` below.
func (o StackOutput) StorageConnectors() StackStorageConnectorArrayOutput {
	return o.ApplyT(func(v *Stack) StackStorageConnectorArrayOutput { return v.StorageConnectors }).(StackStorageConnectorArrayOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o StackOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StackOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Stack) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Configuration block for the actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.
// See `userSettings` below.
func (o StackOutput) UserSettings() StackUserSettingArrayOutput {
	return o.ApplyT(func(v *Stack) StackUserSettingArrayOutput { return v.UserSettings }).(StackUserSettingArrayOutput)
}

type StackArrayOutput struct{ *pulumi.OutputState }

func (StackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stack)(nil)).Elem()
}

func (o StackArrayOutput) ToStackArrayOutput() StackArrayOutput {
	return o
}

func (o StackArrayOutput) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return o
}

func (o StackArrayOutput) Index(i pulumi.IntInput) StackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Stack {
		return vs[0].([]*Stack)[vs[1].(int)]
	}).(StackOutput)
}

type StackMapOutput struct{ *pulumi.OutputState }

func (StackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stack)(nil)).Elem()
}

func (o StackMapOutput) ToStackMapOutput() StackMapOutput {
	return o
}

func (o StackMapOutput) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return o
}

func (o StackMapOutput) MapIndex(k pulumi.StringInput) StackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Stack {
		return vs[0].(map[string]*Stack)[vs[1].(string)]
	}).(StackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackInput)(nil)).Elem(), &Stack{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackArrayInput)(nil)).Elem(), StackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackMapInput)(nil)).Elem(), StackMap{})
	pulumi.RegisterOutputType(StackOutput{})
	pulumi.RegisterOutputType(StackArrayOutput{})
	pulumi.RegisterOutputType(StackMapOutput{})
}
