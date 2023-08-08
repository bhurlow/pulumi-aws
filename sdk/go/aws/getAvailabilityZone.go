// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getAvailabilityZone` provides details about a specific availability zone (AZ)
// in the current region.
//
// This can be used both to validate an availability zone given in a variable
// and to split the AZ name into its component parts of an AWS region and an
// AZ identifier letter. The latter may be useful e.g., for implementing a
// consistent subnet numbering scheme across several regions by mapping both
// the region and the subnet letter to network numbers.
//
// This is different from the `getAvailabilityZones` (plural) data source,
// which provides a list of the available zones.
func GetAvailabilityZone(ctx *pulumi.Context, args *GetAvailabilityZoneArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAvailabilityZoneResult
	err := ctx.Invoke("aws:index/getAvailabilityZone:getAvailabilityZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZone.
type GetAvailabilityZoneArgs struct {
	// Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
	AllAvailabilityZones *bool `pulumi:"allAvailabilityZones"`
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetAvailabilityZoneFilter `pulumi:"filters"`
	// Full name of the availability zone to select.
	Name *string `pulumi:"name"`
	// Specific availability zone state to require. May be any of `"available"`, `"information"` or `"impaired"`.
	State *string `pulumi:"state"`
	// Zone ID of the availability zone to select.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getAvailabilityZone.
type GetAvailabilityZoneResult struct {
	AllAvailabilityZones *bool                       `pulumi:"allAvailabilityZones"`
	Filters              []GetAvailabilityZoneFilter `pulumi:"filters"`
	// For Availability Zones, this is the same value as the Region name. For Local Zones, the name of the associated group, for example `us-west-2-lax-1`.
	GroupName string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Part of the AZ name that appears after the region name, uniquely identifying the AZ within its region.
	// For Availability Zones this is usually a single letter, for example `a` for the `us-west-2a` zone.
	// For Local and Wavelength Zones this is a longer string, for example `wl1-sfo-wlz-1` for the `us-west-2-wl1-sfo-wlz-1` zone.
	NameSuffix string `pulumi:"nameSuffix"`
	// The name of the location from which the address is advertised.
	NetworkBorderGroup string `pulumi:"networkBorderGroup"`
	// For Availability Zones, this always has the value of `opt-in-not-required`. For Local Zones, this is the opt in status. The possible values are `opted-in` and `not-opted-in`.
	OptInStatus string `pulumi:"optInStatus"`
	// ID of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.
	ParentZoneId string `pulumi:"parentZoneId"`
	// Name of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.
	ParentZoneName string `pulumi:"parentZoneName"`
	// Region where the selected availability zone resides. This is always the region selected on the provider, since this data source searches only within that region.
	Region string `pulumi:"region"`
	State  string `pulumi:"state"`
	ZoneId string `pulumi:"zoneId"`
	// Type of zone. Values are `availability-zone`, `local-zone`, and `wavelength-zone`.
	ZoneType string `pulumi:"zoneType"`
}

func GetAvailabilityZoneOutput(ctx *pulumi.Context, args GetAvailabilityZoneOutputArgs, opts ...pulumi.InvokeOption) GetAvailabilityZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAvailabilityZoneResult, error) {
			args := v.(GetAvailabilityZoneArgs)
			r, err := GetAvailabilityZone(ctx, &args, opts...)
			var s GetAvailabilityZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAvailabilityZoneResultOutput)
}

// A collection of arguments for invoking getAvailabilityZone.
type GetAvailabilityZoneOutputArgs struct {
	// Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
	AllAvailabilityZones pulumi.BoolPtrInput `pulumi:"allAvailabilityZones"`
	// Configuration block(s) for filtering. Detailed below.
	Filters GetAvailabilityZoneFilterArrayInput `pulumi:"filters"`
	// Full name of the availability zone to select.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specific availability zone state to require. May be any of `"available"`, `"information"` or `"impaired"`.
	State pulumi.StringPtrInput `pulumi:"state"`
	// Zone ID of the availability zone to select.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetAvailabilityZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZoneArgs)(nil)).Elem()
}

// A collection of values returned by getAvailabilityZone.
type GetAvailabilityZoneResultOutput struct{ *pulumi.OutputState }

func (GetAvailabilityZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZoneResult)(nil)).Elem()
}

func (o GetAvailabilityZoneResultOutput) ToGetAvailabilityZoneResultOutput() GetAvailabilityZoneResultOutput {
	return o
}

func (o GetAvailabilityZoneResultOutput) ToGetAvailabilityZoneResultOutputWithContext(ctx context.Context) GetAvailabilityZoneResultOutput {
	return o
}

func (o GetAvailabilityZoneResultOutput) AllAvailabilityZones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) *bool { return v.AllAvailabilityZones }).(pulumi.BoolPtrOutput)
}

func (o GetAvailabilityZoneResultOutput) Filters() GetAvailabilityZoneFilterArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) []GetAvailabilityZoneFilter { return v.Filters }).(GetAvailabilityZoneFilterArrayOutput)
}

// For Availability Zones, this is the same value as the Region name. For Local Zones, the name of the associated group, for example `us-west-2-lax-1`.
func (o GetAvailabilityZoneResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAvailabilityZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAvailabilityZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// Part of the AZ name that appears after the region name, uniquely identifying the AZ within its region.
// For Availability Zones this is usually a single letter, for example `a` for the `us-west-2a` zone.
// For Local and Wavelength Zones this is a longer string, for example `wl1-sfo-wlz-1` for the `us-west-2-wl1-sfo-wlz-1` zone.
func (o GetAvailabilityZoneResultOutput) NameSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.NameSuffix }).(pulumi.StringOutput)
}

// The name of the location from which the address is advertised.
func (o GetAvailabilityZoneResultOutput) NetworkBorderGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.NetworkBorderGroup }).(pulumi.StringOutput)
}

// For Availability Zones, this always has the value of `opt-in-not-required`. For Local Zones, this is the opt in status. The possible values are `opted-in` and `not-opted-in`.
func (o GetAvailabilityZoneResultOutput) OptInStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.OptInStatus }).(pulumi.StringOutput)
}

// ID of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.
func (o GetAvailabilityZoneResultOutput) ParentZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.ParentZoneId }).(pulumi.StringOutput)
}

// Name of the zone that handles some of the Local Zone or Wavelength Zone control plane operations, such as API calls.
func (o GetAvailabilityZoneResultOutput) ParentZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.ParentZoneName }).(pulumi.StringOutput)
}

// Region where the selected availability zone resides. This is always the region selected on the provider, since this data source searches only within that region.
func (o GetAvailabilityZoneResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetAvailabilityZoneResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetAvailabilityZoneResultOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.ZoneId }).(pulumi.StringOutput)
}

// Type of zone. Values are `availability-zone`, `local-zone`, and `wavelength-zone`.
func (o GetAvailabilityZoneResultOutput) ZoneType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZoneResult) string { return v.ZoneType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAvailabilityZoneResultOutput{})
}
