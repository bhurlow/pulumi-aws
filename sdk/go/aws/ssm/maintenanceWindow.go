// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an SSM Maintenance Window resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewMaintenanceWindow(ctx, "production", &ssm.MaintenanceWindowArgs{
//				Cutoff:   pulumi.Int(1),
//				Duration: pulumi.Int(3),
//				Schedule: pulumi.String("cron(0 16 ? * TUE *)"),
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
// # Using `pulumi import`, import SSM
//
// Maintenance Windows using the maintenance window `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ssm/maintenanceWindow:MaintenanceWindow imported-window mw-0123456789
//
// ```
type MaintenanceWindow struct {
	pulumi.CustomResourceState

	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets pulumi.BoolPtrOutput `pulumi:"allowUnassociatedTargets"`
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff pulumi.IntOutput `pulumi:"cutoff"`
	// A description for the maintenance window.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The duration of the Maintenance Window in hours.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// The name of the maintenance window.
	Name pulumi.StringOutput `pulumi:"name"`
	// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset pulumi.IntPtrOutput `pulumi:"scheduleOffset"`
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone pulumi.StringPtrOutput `pulumi:"scheduleTimezone"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cutoff == nil {
		return nil, errors.New("invalid value for required argument 'Cutoff'")
	}
	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MaintenanceWindow
	err := ctx.RegisterResource("aws:ssm/maintenanceWindow:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("aws:ssm/maintenanceWindow:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets *bool `pulumi:"allowUnassociatedTargets"`
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff *int `pulumi:"cutoff"`
	// A description for the maintenance window.
	Description *string `pulumi:"description"`
	// The duration of the Maintenance Window in hours.
	Duration *int `pulumi:"duration"`
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate *string `pulumi:"endDate"`
	// The name of the maintenance window.
	Name *string `pulumi:"name"`
	// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
	Schedule *string `pulumi:"schedule"`
	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset *int `pulumi:"scheduleOffset"`
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone *string `pulumi:"scheduleTimezone"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate *string `pulumi:"startDate"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type MaintenanceWindowState struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets pulumi.BoolPtrInput
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff pulumi.IntPtrInput
	// A description for the maintenance window.
	Description pulumi.StringPtrInput
	// The duration of the Maintenance Window in hours.
	Duration pulumi.IntPtrInput
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate pulumi.StringPtrInput
	// The name of the maintenance window.
	Name pulumi.StringPtrInput
	// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
	Schedule pulumi.StringPtrInput
	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset pulumi.IntPtrInput
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone pulumi.StringPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets *bool `pulumi:"allowUnassociatedTargets"`
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff int `pulumi:"cutoff"`
	// A description for the maintenance window.
	Description *string `pulumi:"description"`
	// The duration of the Maintenance Window in hours.
	Duration int `pulumi:"duration"`
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate *string `pulumi:"endDate"`
	// The name of the maintenance window.
	Name *string `pulumi:"name"`
	// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
	Schedule string `pulumi:"schedule"`
	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset *int `pulumi:"scheduleOffset"`
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone *string `pulumi:"scheduleTimezone"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate *string `pulumi:"startDate"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets pulumi.BoolPtrInput
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff pulumi.IntInput
	// A description for the maintenance window.
	Description pulumi.StringPtrInput
	// The duration of the Maintenance Window in hours.
	Duration pulumi.IntInput
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate pulumi.StringPtrInput
	// The name of the maintenance window.
	Name pulumi.StringPtrInput
	// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
	Schedule pulumi.StringInput
	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset pulumi.IntPtrInput
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone pulumi.StringPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}

type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput
}

func (*MaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

func (i *MaintenanceWindow) ToOutput(ctx context.Context) pulumix.Output[*MaintenanceWindow] {
	return pulumix.Output[*MaintenanceWindow]{
		OutputState: i.ToMaintenanceWindowOutputWithContext(ctx).OutputState,
	}
}

// MaintenanceWindowArrayInput is an input type that accepts MaintenanceWindowArray and MaintenanceWindowArrayOutput values.
// You can construct a concrete instance of `MaintenanceWindowArrayInput` via:
//
//	MaintenanceWindowArray{ MaintenanceWindowArgs{...} }
type MaintenanceWindowArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput
	ToMaintenanceWindowArrayOutputWithContext(context.Context) MaintenanceWindowArrayOutput
}

type MaintenanceWindowArray []MaintenanceWindowInput

func (MaintenanceWindowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return i.ToMaintenanceWindowArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowArrayOutput)
}

func (i MaintenanceWindowArray) ToOutput(ctx context.Context) pulumix.Output[[]*MaintenanceWindow] {
	return pulumix.Output[[]*MaintenanceWindow]{
		OutputState: i.ToMaintenanceWindowArrayOutputWithContext(ctx).OutputState,
	}
}

// MaintenanceWindowMapInput is an input type that accepts MaintenanceWindowMap and MaintenanceWindowMapOutput values.
// You can construct a concrete instance of `MaintenanceWindowMapInput` via:
//
//	MaintenanceWindowMap{ "key": MaintenanceWindowArgs{...} }
type MaintenanceWindowMapInput interface {
	pulumi.Input

	ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput
	ToMaintenanceWindowMapOutputWithContext(context.Context) MaintenanceWindowMapOutput
}

type MaintenanceWindowMap map[string]MaintenanceWindowInput

func (MaintenanceWindowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return i.ToMaintenanceWindowMapOutputWithContext(context.Background())
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowMapOutput)
}

func (i MaintenanceWindowMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MaintenanceWindow] {
	return pulumix.Output[map[string]*MaintenanceWindow]{
		OutputState: i.ToMaintenanceWindowMapOutputWithContext(ctx).OutputState,
	}
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToOutput(ctx context.Context) pulumix.Output[*MaintenanceWindow] {
	return pulumix.Output[*MaintenanceWindow]{
		OutputState: o.OutputState,
	}
}

// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
func (o MaintenanceWindowOutput) AllowUnassociatedTargets() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.BoolPtrOutput { return v.AllowUnassociatedTargets }).(pulumi.BoolPtrOutput)
}

// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
func (o MaintenanceWindowOutput) Cutoff() pulumi.IntOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntOutput { return v.Cutoff }).(pulumi.IntOutput)
}

// A description for the maintenance window.
func (o MaintenanceWindowOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The duration of the Maintenance Window in hours.
func (o MaintenanceWindowOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntOutput { return v.Duration }).(pulumi.IntOutput)
}

// Whether the maintenance window is enabled. Default: `true`.
func (o MaintenanceWindowOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
func (o MaintenanceWindowOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

// The name of the maintenance window.
func (o MaintenanceWindowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
func (o MaintenanceWindowOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
func (o MaintenanceWindowOutput) ScheduleOffset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntPtrOutput { return v.ScheduleOffset }).(pulumi.IntPtrOutput)
}

// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
func (o MaintenanceWindowOutput) ScheduleTimezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.ScheduleTimezone }).(pulumi.StringPtrOutput)
}

// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
func (o MaintenanceWindowOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o MaintenanceWindowOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o MaintenanceWindowOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type MaintenanceWindowArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MaintenanceWindow] {
	return pulumix.Output[[]*MaintenanceWindow]{
		OutputState: o.OutputState,
	}
}

func (o MaintenanceWindowArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].([]*MaintenanceWindow)[vs[1].(int)]
	}).(MaintenanceWindowOutput)
}

type MaintenanceWindowMapOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MaintenanceWindow] {
	return pulumix.Output[map[string]*MaintenanceWindow]{
		OutputState: o.OutputState,
	}
}

func (o MaintenanceWindowMapOutput) MapIndex(k pulumi.StringInput) MaintenanceWindowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].(map[string]*MaintenanceWindow)[vs[1].(string)]
	}).(MaintenanceWindowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowInput)(nil)).Elem(), &MaintenanceWindow{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowArrayInput)(nil)).Elem(), MaintenanceWindowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowMapInput)(nil)).Elem(), MaintenanceWindowMap{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowMapOutput{})
}
