// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ScraperDestination struct {
	// Configuration block for an Amazon Managed Prometheus workspace destination. See `amp`.
	Amp *ScraperDestinationAmp `pulumi:"amp"`
}

// ScraperDestinationInput is an input type that accepts ScraperDestinationArgs and ScraperDestinationOutput values.
// You can construct a concrete instance of `ScraperDestinationInput` via:
//
//	ScraperDestinationArgs{...}
type ScraperDestinationInput interface {
	pulumi.Input

	ToScraperDestinationOutput() ScraperDestinationOutput
	ToScraperDestinationOutputWithContext(context.Context) ScraperDestinationOutput
}

type ScraperDestinationArgs struct {
	// Configuration block for an Amazon Managed Prometheus workspace destination. See `amp`.
	Amp ScraperDestinationAmpPtrInput `pulumi:"amp"`
}

func (ScraperDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperDestination)(nil)).Elem()
}

func (i ScraperDestinationArgs) ToScraperDestinationOutput() ScraperDestinationOutput {
	return i.ToScraperDestinationOutputWithContext(context.Background())
}

func (i ScraperDestinationArgs) ToScraperDestinationOutputWithContext(ctx context.Context) ScraperDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperDestinationOutput)
}

func (i ScraperDestinationArgs) ToScraperDestinationPtrOutput() ScraperDestinationPtrOutput {
	return i.ToScraperDestinationPtrOutputWithContext(context.Background())
}

func (i ScraperDestinationArgs) ToScraperDestinationPtrOutputWithContext(ctx context.Context) ScraperDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperDestinationOutput).ToScraperDestinationPtrOutputWithContext(ctx)
}

// ScraperDestinationPtrInput is an input type that accepts ScraperDestinationArgs, ScraperDestinationPtr and ScraperDestinationPtrOutput values.
// You can construct a concrete instance of `ScraperDestinationPtrInput` via:
//
//	        ScraperDestinationArgs{...}
//
//	or:
//
//	        nil
type ScraperDestinationPtrInput interface {
	pulumi.Input

	ToScraperDestinationPtrOutput() ScraperDestinationPtrOutput
	ToScraperDestinationPtrOutputWithContext(context.Context) ScraperDestinationPtrOutput
}

type scraperDestinationPtrType ScraperDestinationArgs

func ScraperDestinationPtr(v *ScraperDestinationArgs) ScraperDestinationPtrInput {
	return (*scraperDestinationPtrType)(v)
}

func (*scraperDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperDestination)(nil)).Elem()
}

func (i *scraperDestinationPtrType) ToScraperDestinationPtrOutput() ScraperDestinationPtrOutput {
	return i.ToScraperDestinationPtrOutputWithContext(context.Background())
}

func (i *scraperDestinationPtrType) ToScraperDestinationPtrOutputWithContext(ctx context.Context) ScraperDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperDestinationPtrOutput)
}

type ScraperDestinationOutput struct{ *pulumi.OutputState }

func (ScraperDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperDestination)(nil)).Elem()
}

func (o ScraperDestinationOutput) ToScraperDestinationOutput() ScraperDestinationOutput {
	return o
}

func (o ScraperDestinationOutput) ToScraperDestinationOutputWithContext(ctx context.Context) ScraperDestinationOutput {
	return o
}

func (o ScraperDestinationOutput) ToScraperDestinationPtrOutput() ScraperDestinationPtrOutput {
	return o.ToScraperDestinationPtrOutputWithContext(context.Background())
}

func (o ScraperDestinationOutput) ToScraperDestinationPtrOutputWithContext(ctx context.Context) ScraperDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScraperDestination) *ScraperDestination {
		return &v
	}).(ScraperDestinationPtrOutput)
}

// Configuration block for an Amazon Managed Prometheus workspace destination. See `amp`.
func (o ScraperDestinationOutput) Amp() ScraperDestinationAmpPtrOutput {
	return o.ApplyT(func(v ScraperDestination) *ScraperDestinationAmp { return v.Amp }).(ScraperDestinationAmpPtrOutput)
}

type ScraperDestinationPtrOutput struct{ *pulumi.OutputState }

func (ScraperDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperDestination)(nil)).Elem()
}

func (o ScraperDestinationPtrOutput) ToScraperDestinationPtrOutput() ScraperDestinationPtrOutput {
	return o
}

func (o ScraperDestinationPtrOutput) ToScraperDestinationPtrOutputWithContext(ctx context.Context) ScraperDestinationPtrOutput {
	return o
}

func (o ScraperDestinationPtrOutput) Elem() ScraperDestinationOutput {
	return o.ApplyT(func(v *ScraperDestination) ScraperDestination {
		if v != nil {
			return *v
		}
		var ret ScraperDestination
		return ret
	}).(ScraperDestinationOutput)
}

// Configuration block for an Amazon Managed Prometheus workspace destination. See `amp`.
func (o ScraperDestinationPtrOutput) Amp() ScraperDestinationAmpPtrOutput {
	return o.ApplyT(func(v *ScraperDestination) *ScraperDestinationAmp {
		if v == nil {
			return nil
		}
		return v.Amp
	}).(ScraperDestinationAmpPtrOutput)
}

type ScraperDestinationAmp struct {
	// The Amazon Resource Name (ARN) of the prometheus workspace.
	WorkspaceArn string `pulumi:"workspaceArn"`
}

// ScraperDestinationAmpInput is an input type that accepts ScraperDestinationAmpArgs and ScraperDestinationAmpOutput values.
// You can construct a concrete instance of `ScraperDestinationAmpInput` via:
//
//	ScraperDestinationAmpArgs{...}
type ScraperDestinationAmpInput interface {
	pulumi.Input

	ToScraperDestinationAmpOutput() ScraperDestinationAmpOutput
	ToScraperDestinationAmpOutputWithContext(context.Context) ScraperDestinationAmpOutput
}

type ScraperDestinationAmpArgs struct {
	// The Amazon Resource Name (ARN) of the prometheus workspace.
	WorkspaceArn pulumi.StringInput `pulumi:"workspaceArn"`
}

func (ScraperDestinationAmpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperDestinationAmp)(nil)).Elem()
}

func (i ScraperDestinationAmpArgs) ToScraperDestinationAmpOutput() ScraperDestinationAmpOutput {
	return i.ToScraperDestinationAmpOutputWithContext(context.Background())
}

func (i ScraperDestinationAmpArgs) ToScraperDestinationAmpOutputWithContext(ctx context.Context) ScraperDestinationAmpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperDestinationAmpOutput)
}

func (i ScraperDestinationAmpArgs) ToScraperDestinationAmpPtrOutput() ScraperDestinationAmpPtrOutput {
	return i.ToScraperDestinationAmpPtrOutputWithContext(context.Background())
}

func (i ScraperDestinationAmpArgs) ToScraperDestinationAmpPtrOutputWithContext(ctx context.Context) ScraperDestinationAmpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperDestinationAmpOutput).ToScraperDestinationAmpPtrOutputWithContext(ctx)
}

// ScraperDestinationAmpPtrInput is an input type that accepts ScraperDestinationAmpArgs, ScraperDestinationAmpPtr and ScraperDestinationAmpPtrOutput values.
// You can construct a concrete instance of `ScraperDestinationAmpPtrInput` via:
//
//	        ScraperDestinationAmpArgs{...}
//
//	or:
//
//	        nil
type ScraperDestinationAmpPtrInput interface {
	pulumi.Input

	ToScraperDestinationAmpPtrOutput() ScraperDestinationAmpPtrOutput
	ToScraperDestinationAmpPtrOutputWithContext(context.Context) ScraperDestinationAmpPtrOutput
}

type scraperDestinationAmpPtrType ScraperDestinationAmpArgs

func ScraperDestinationAmpPtr(v *ScraperDestinationAmpArgs) ScraperDestinationAmpPtrInput {
	return (*scraperDestinationAmpPtrType)(v)
}

func (*scraperDestinationAmpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperDestinationAmp)(nil)).Elem()
}

func (i *scraperDestinationAmpPtrType) ToScraperDestinationAmpPtrOutput() ScraperDestinationAmpPtrOutput {
	return i.ToScraperDestinationAmpPtrOutputWithContext(context.Background())
}

func (i *scraperDestinationAmpPtrType) ToScraperDestinationAmpPtrOutputWithContext(ctx context.Context) ScraperDestinationAmpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperDestinationAmpPtrOutput)
}

type ScraperDestinationAmpOutput struct{ *pulumi.OutputState }

func (ScraperDestinationAmpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperDestinationAmp)(nil)).Elem()
}

func (o ScraperDestinationAmpOutput) ToScraperDestinationAmpOutput() ScraperDestinationAmpOutput {
	return o
}

func (o ScraperDestinationAmpOutput) ToScraperDestinationAmpOutputWithContext(ctx context.Context) ScraperDestinationAmpOutput {
	return o
}

func (o ScraperDestinationAmpOutput) ToScraperDestinationAmpPtrOutput() ScraperDestinationAmpPtrOutput {
	return o.ToScraperDestinationAmpPtrOutputWithContext(context.Background())
}

func (o ScraperDestinationAmpOutput) ToScraperDestinationAmpPtrOutputWithContext(ctx context.Context) ScraperDestinationAmpPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScraperDestinationAmp) *ScraperDestinationAmp {
		return &v
	}).(ScraperDestinationAmpPtrOutput)
}

// The Amazon Resource Name (ARN) of the prometheus workspace.
func (o ScraperDestinationAmpOutput) WorkspaceArn() pulumi.StringOutput {
	return o.ApplyT(func(v ScraperDestinationAmp) string { return v.WorkspaceArn }).(pulumi.StringOutput)
}

type ScraperDestinationAmpPtrOutput struct{ *pulumi.OutputState }

func (ScraperDestinationAmpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperDestinationAmp)(nil)).Elem()
}

func (o ScraperDestinationAmpPtrOutput) ToScraperDestinationAmpPtrOutput() ScraperDestinationAmpPtrOutput {
	return o
}

func (o ScraperDestinationAmpPtrOutput) ToScraperDestinationAmpPtrOutputWithContext(ctx context.Context) ScraperDestinationAmpPtrOutput {
	return o
}

func (o ScraperDestinationAmpPtrOutput) Elem() ScraperDestinationAmpOutput {
	return o.ApplyT(func(v *ScraperDestinationAmp) ScraperDestinationAmp {
		if v != nil {
			return *v
		}
		var ret ScraperDestinationAmp
		return ret
	}).(ScraperDestinationAmpOutput)
}

// The Amazon Resource Name (ARN) of the prometheus workspace.
func (o ScraperDestinationAmpPtrOutput) WorkspaceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScraperDestinationAmp) *string {
		if v == nil {
			return nil
		}
		return &v.WorkspaceArn
	}).(pulumi.StringPtrOutput)
}

type ScraperSource struct {
	// Configuration block for an EKS cluster source. See `eks`.
	Eks *ScraperSourceEks `pulumi:"eks"`
}

// ScraperSourceInput is an input type that accepts ScraperSourceArgs and ScraperSourceOutput values.
// You can construct a concrete instance of `ScraperSourceInput` via:
//
//	ScraperSourceArgs{...}
type ScraperSourceInput interface {
	pulumi.Input

	ToScraperSourceOutput() ScraperSourceOutput
	ToScraperSourceOutputWithContext(context.Context) ScraperSourceOutput
}

type ScraperSourceArgs struct {
	// Configuration block for an EKS cluster source. See `eks`.
	Eks ScraperSourceEksPtrInput `pulumi:"eks"`
}

func (ScraperSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperSource)(nil)).Elem()
}

func (i ScraperSourceArgs) ToScraperSourceOutput() ScraperSourceOutput {
	return i.ToScraperSourceOutputWithContext(context.Background())
}

func (i ScraperSourceArgs) ToScraperSourceOutputWithContext(ctx context.Context) ScraperSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperSourceOutput)
}

func (i ScraperSourceArgs) ToScraperSourcePtrOutput() ScraperSourcePtrOutput {
	return i.ToScraperSourcePtrOutputWithContext(context.Background())
}

func (i ScraperSourceArgs) ToScraperSourcePtrOutputWithContext(ctx context.Context) ScraperSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperSourceOutput).ToScraperSourcePtrOutputWithContext(ctx)
}

// ScraperSourcePtrInput is an input type that accepts ScraperSourceArgs, ScraperSourcePtr and ScraperSourcePtrOutput values.
// You can construct a concrete instance of `ScraperSourcePtrInput` via:
//
//	        ScraperSourceArgs{...}
//
//	or:
//
//	        nil
type ScraperSourcePtrInput interface {
	pulumi.Input

	ToScraperSourcePtrOutput() ScraperSourcePtrOutput
	ToScraperSourcePtrOutputWithContext(context.Context) ScraperSourcePtrOutput
}

type scraperSourcePtrType ScraperSourceArgs

func ScraperSourcePtr(v *ScraperSourceArgs) ScraperSourcePtrInput {
	return (*scraperSourcePtrType)(v)
}

func (*scraperSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperSource)(nil)).Elem()
}

func (i *scraperSourcePtrType) ToScraperSourcePtrOutput() ScraperSourcePtrOutput {
	return i.ToScraperSourcePtrOutputWithContext(context.Background())
}

func (i *scraperSourcePtrType) ToScraperSourcePtrOutputWithContext(ctx context.Context) ScraperSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperSourcePtrOutput)
}

type ScraperSourceOutput struct{ *pulumi.OutputState }

func (ScraperSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperSource)(nil)).Elem()
}

func (o ScraperSourceOutput) ToScraperSourceOutput() ScraperSourceOutput {
	return o
}

func (o ScraperSourceOutput) ToScraperSourceOutputWithContext(ctx context.Context) ScraperSourceOutput {
	return o
}

func (o ScraperSourceOutput) ToScraperSourcePtrOutput() ScraperSourcePtrOutput {
	return o.ToScraperSourcePtrOutputWithContext(context.Background())
}

func (o ScraperSourceOutput) ToScraperSourcePtrOutputWithContext(ctx context.Context) ScraperSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScraperSource) *ScraperSource {
		return &v
	}).(ScraperSourcePtrOutput)
}

// Configuration block for an EKS cluster source. See `eks`.
func (o ScraperSourceOutput) Eks() ScraperSourceEksPtrOutput {
	return o.ApplyT(func(v ScraperSource) *ScraperSourceEks { return v.Eks }).(ScraperSourceEksPtrOutput)
}

type ScraperSourcePtrOutput struct{ *pulumi.OutputState }

func (ScraperSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperSource)(nil)).Elem()
}

func (o ScraperSourcePtrOutput) ToScraperSourcePtrOutput() ScraperSourcePtrOutput {
	return o
}

func (o ScraperSourcePtrOutput) ToScraperSourcePtrOutputWithContext(ctx context.Context) ScraperSourcePtrOutput {
	return o
}

func (o ScraperSourcePtrOutput) Elem() ScraperSourceOutput {
	return o.ApplyT(func(v *ScraperSource) ScraperSource {
		if v != nil {
			return *v
		}
		var ret ScraperSource
		return ret
	}).(ScraperSourceOutput)
}

// Configuration block for an EKS cluster source. See `eks`.
func (o ScraperSourcePtrOutput) Eks() ScraperSourceEksPtrOutput {
	return o.ApplyT(func(v *ScraperSource) *ScraperSourceEks {
		if v == nil {
			return nil
		}
		return v.Eks
	}).(ScraperSourceEksPtrOutput)
}

type ScraperSourceEks struct {
	ClusterArn string `pulumi:"clusterArn"`
	// List of the security group IDs for the Amazon EKS cluster VPC configuration.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// List of subnet IDs. Must be in at least two different availability zones.
	SubnetIds []string `pulumi:"subnetIds"`
}

// ScraperSourceEksInput is an input type that accepts ScraperSourceEksArgs and ScraperSourceEksOutput values.
// You can construct a concrete instance of `ScraperSourceEksInput` via:
//
//	ScraperSourceEksArgs{...}
type ScraperSourceEksInput interface {
	pulumi.Input

	ToScraperSourceEksOutput() ScraperSourceEksOutput
	ToScraperSourceEksOutputWithContext(context.Context) ScraperSourceEksOutput
}

type ScraperSourceEksArgs struct {
	ClusterArn pulumi.StringInput `pulumi:"clusterArn"`
	// List of the security group IDs for the Amazon EKS cluster VPC configuration.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// List of subnet IDs. Must be in at least two different availability zones.
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
}

func (ScraperSourceEksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperSourceEks)(nil)).Elem()
}

func (i ScraperSourceEksArgs) ToScraperSourceEksOutput() ScraperSourceEksOutput {
	return i.ToScraperSourceEksOutputWithContext(context.Background())
}

func (i ScraperSourceEksArgs) ToScraperSourceEksOutputWithContext(ctx context.Context) ScraperSourceEksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperSourceEksOutput)
}

func (i ScraperSourceEksArgs) ToScraperSourceEksPtrOutput() ScraperSourceEksPtrOutput {
	return i.ToScraperSourceEksPtrOutputWithContext(context.Background())
}

func (i ScraperSourceEksArgs) ToScraperSourceEksPtrOutputWithContext(ctx context.Context) ScraperSourceEksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperSourceEksOutput).ToScraperSourceEksPtrOutputWithContext(ctx)
}

// ScraperSourceEksPtrInput is an input type that accepts ScraperSourceEksArgs, ScraperSourceEksPtr and ScraperSourceEksPtrOutput values.
// You can construct a concrete instance of `ScraperSourceEksPtrInput` via:
//
//	        ScraperSourceEksArgs{...}
//
//	or:
//
//	        nil
type ScraperSourceEksPtrInput interface {
	pulumi.Input

	ToScraperSourceEksPtrOutput() ScraperSourceEksPtrOutput
	ToScraperSourceEksPtrOutputWithContext(context.Context) ScraperSourceEksPtrOutput
}

type scraperSourceEksPtrType ScraperSourceEksArgs

func ScraperSourceEksPtr(v *ScraperSourceEksArgs) ScraperSourceEksPtrInput {
	return (*scraperSourceEksPtrType)(v)
}

func (*scraperSourceEksPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperSourceEks)(nil)).Elem()
}

func (i *scraperSourceEksPtrType) ToScraperSourceEksPtrOutput() ScraperSourceEksPtrOutput {
	return i.ToScraperSourceEksPtrOutputWithContext(context.Background())
}

func (i *scraperSourceEksPtrType) ToScraperSourceEksPtrOutputWithContext(ctx context.Context) ScraperSourceEksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperSourceEksPtrOutput)
}

type ScraperSourceEksOutput struct{ *pulumi.OutputState }

func (ScraperSourceEksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperSourceEks)(nil)).Elem()
}

func (o ScraperSourceEksOutput) ToScraperSourceEksOutput() ScraperSourceEksOutput {
	return o
}

func (o ScraperSourceEksOutput) ToScraperSourceEksOutputWithContext(ctx context.Context) ScraperSourceEksOutput {
	return o
}

func (o ScraperSourceEksOutput) ToScraperSourceEksPtrOutput() ScraperSourceEksPtrOutput {
	return o.ToScraperSourceEksPtrOutputWithContext(context.Background())
}

func (o ScraperSourceEksOutput) ToScraperSourceEksPtrOutputWithContext(ctx context.Context) ScraperSourceEksPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScraperSourceEks) *ScraperSourceEks {
		return &v
	}).(ScraperSourceEksPtrOutput)
}

func (o ScraperSourceEksOutput) ClusterArn() pulumi.StringOutput {
	return o.ApplyT(func(v ScraperSourceEks) string { return v.ClusterArn }).(pulumi.StringOutput)
}

// List of the security group IDs for the Amazon EKS cluster VPC configuration.
func (o ScraperSourceEksOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScraperSourceEks) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// List of subnet IDs. Must be in at least two different availability zones.
func (o ScraperSourceEksOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScraperSourceEks) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

type ScraperSourceEksPtrOutput struct{ *pulumi.OutputState }

func (ScraperSourceEksPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperSourceEks)(nil)).Elem()
}

func (o ScraperSourceEksPtrOutput) ToScraperSourceEksPtrOutput() ScraperSourceEksPtrOutput {
	return o
}

func (o ScraperSourceEksPtrOutput) ToScraperSourceEksPtrOutputWithContext(ctx context.Context) ScraperSourceEksPtrOutput {
	return o
}

func (o ScraperSourceEksPtrOutput) Elem() ScraperSourceEksOutput {
	return o.ApplyT(func(v *ScraperSourceEks) ScraperSourceEks {
		if v != nil {
			return *v
		}
		var ret ScraperSourceEks
		return ret
	}).(ScraperSourceEksOutput)
}

func (o ScraperSourceEksPtrOutput) ClusterArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScraperSourceEks) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterArn
	}).(pulumi.StringPtrOutput)
}

// List of the security group IDs for the Amazon EKS cluster VPC configuration.
func (o ScraperSourceEksPtrOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScraperSourceEks) []string {
		if v == nil {
			return nil
		}
		return v.SecurityGroupIds
	}).(pulumi.StringArrayOutput)
}

// List of subnet IDs. Must be in at least two different availability zones.
func (o ScraperSourceEksPtrOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScraperSourceEks) []string {
		if v == nil {
			return nil
		}
		return v.SubnetIds
	}).(pulumi.StringArrayOutput)
}

type ScraperTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	Create *string `pulumi:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	Delete *string `pulumi:"delete"`
}

// ScraperTimeoutsInput is an input type that accepts ScraperTimeoutsArgs and ScraperTimeoutsOutput values.
// You can construct a concrete instance of `ScraperTimeoutsInput` via:
//
//	ScraperTimeoutsArgs{...}
type ScraperTimeoutsInput interface {
	pulumi.Input

	ToScraperTimeoutsOutput() ScraperTimeoutsOutput
	ToScraperTimeoutsOutputWithContext(context.Context) ScraperTimeoutsOutput
}

type ScraperTimeoutsArgs struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	Create pulumi.StringPtrInput `pulumi:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	Delete pulumi.StringPtrInput `pulumi:"delete"`
}

func (ScraperTimeoutsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperTimeouts)(nil)).Elem()
}

func (i ScraperTimeoutsArgs) ToScraperTimeoutsOutput() ScraperTimeoutsOutput {
	return i.ToScraperTimeoutsOutputWithContext(context.Background())
}

func (i ScraperTimeoutsArgs) ToScraperTimeoutsOutputWithContext(ctx context.Context) ScraperTimeoutsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperTimeoutsOutput)
}

func (i ScraperTimeoutsArgs) ToScraperTimeoutsPtrOutput() ScraperTimeoutsPtrOutput {
	return i.ToScraperTimeoutsPtrOutputWithContext(context.Background())
}

func (i ScraperTimeoutsArgs) ToScraperTimeoutsPtrOutputWithContext(ctx context.Context) ScraperTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperTimeoutsOutput).ToScraperTimeoutsPtrOutputWithContext(ctx)
}

// ScraperTimeoutsPtrInput is an input type that accepts ScraperTimeoutsArgs, ScraperTimeoutsPtr and ScraperTimeoutsPtrOutput values.
// You can construct a concrete instance of `ScraperTimeoutsPtrInput` via:
//
//	        ScraperTimeoutsArgs{...}
//
//	or:
//
//	        nil
type ScraperTimeoutsPtrInput interface {
	pulumi.Input

	ToScraperTimeoutsPtrOutput() ScraperTimeoutsPtrOutput
	ToScraperTimeoutsPtrOutputWithContext(context.Context) ScraperTimeoutsPtrOutput
}

type scraperTimeoutsPtrType ScraperTimeoutsArgs

func ScraperTimeoutsPtr(v *ScraperTimeoutsArgs) ScraperTimeoutsPtrInput {
	return (*scraperTimeoutsPtrType)(v)
}

func (*scraperTimeoutsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperTimeouts)(nil)).Elem()
}

func (i *scraperTimeoutsPtrType) ToScraperTimeoutsPtrOutput() ScraperTimeoutsPtrOutput {
	return i.ToScraperTimeoutsPtrOutputWithContext(context.Background())
}

func (i *scraperTimeoutsPtrType) ToScraperTimeoutsPtrOutputWithContext(ctx context.Context) ScraperTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScraperTimeoutsPtrOutput)
}

type ScraperTimeoutsOutput struct{ *pulumi.OutputState }

func (ScraperTimeoutsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScraperTimeouts)(nil)).Elem()
}

func (o ScraperTimeoutsOutput) ToScraperTimeoutsOutput() ScraperTimeoutsOutput {
	return o
}

func (o ScraperTimeoutsOutput) ToScraperTimeoutsOutputWithContext(ctx context.Context) ScraperTimeoutsOutput {
	return o
}

func (o ScraperTimeoutsOutput) ToScraperTimeoutsPtrOutput() ScraperTimeoutsPtrOutput {
	return o.ToScraperTimeoutsPtrOutputWithContext(context.Background())
}

func (o ScraperTimeoutsOutput) ToScraperTimeoutsPtrOutputWithContext(ctx context.Context) ScraperTimeoutsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScraperTimeouts) *ScraperTimeouts {
		return &v
	}).(ScraperTimeoutsPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
func (o ScraperTimeoutsOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScraperTimeouts) *string { return v.Create }).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
func (o ScraperTimeoutsOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScraperTimeouts) *string { return v.Delete }).(pulumi.StringPtrOutput)
}

type ScraperTimeoutsPtrOutput struct{ *pulumi.OutputState }

func (ScraperTimeoutsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScraperTimeouts)(nil)).Elem()
}

func (o ScraperTimeoutsPtrOutput) ToScraperTimeoutsPtrOutput() ScraperTimeoutsPtrOutput {
	return o
}

func (o ScraperTimeoutsPtrOutput) ToScraperTimeoutsPtrOutputWithContext(ctx context.Context) ScraperTimeoutsPtrOutput {
	return o
}

func (o ScraperTimeoutsPtrOutput) Elem() ScraperTimeoutsOutput {
	return o.ApplyT(func(v *ScraperTimeouts) ScraperTimeouts {
		if v != nil {
			return *v
		}
		var ret ScraperTimeouts
		return ret
	}).(ScraperTimeoutsOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
func (o ScraperTimeoutsPtrOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScraperTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Create
	}).(pulumi.StringPtrOutput)
}

// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
func (o ScraperTimeoutsPtrOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScraperTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(pulumi.StringPtrOutput)
}

type WorkspaceLoggingConfiguration struct {
	// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
	LogGroupArn string `pulumi:"logGroupArn"`
}

// WorkspaceLoggingConfigurationInput is an input type that accepts WorkspaceLoggingConfigurationArgs and WorkspaceLoggingConfigurationOutput values.
// You can construct a concrete instance of `WorkspaceLoggingConfigurationInput` via:
//
//	WorkspaceLoggingConfigurationArgs{...}
type WorkspaceLoggingConfigurationInput interface {
	pulumi.Input

	ToWorkspaceLoggingConfigurationOutput() WorkspaceLoggingConfigurationOutput
	ToWorkspaceLoggingConfigurationOutputWithContext(context.Context) WorkspaceLoggingConfigurationOutput
}

type WorkspaceLoggingConfigurationArgs struct {
	// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
	LogGroupArn pulumi.StringInput `pulumi:"logGroupArn"`
}

func (WorkspaceLoggingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationOutput() WorkspaceLoggingConfigurationOutput {
	return i.ToWorkspaceLoggingConfigurationOutputWithContext(context.Background())
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceLoggingConfigurationOutput)
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return i.ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i WorkspaceLoggingConfigurationArgs) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceLoggingConfigurationOutput).ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx)
}

// WorkspaceLoggingConfigurationPtrInput is an input type that accepts WorkspaceLoggingConfigurationArgs, WorkspaceLoggingConfigurationPtr and WorkspaceLoggingConfigurationPtrOutput values.
// You can construct a concrete instance of `WorkspaceLoggingConfigurationPtrInput` via:
//
//	        WorkspaceLoggingConfigurationArgs{...}
//
//	or:
//
//	        nil
type WorkspaceLoggingConfigurationPtrInput interface {
	pulumi.Input

	ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput
	ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Context) WorkspaceLoggingConfigurationPtrOutput
}

type workspaceLoggingConfigurationPtrType WorkspaceLoggingConfigurationArgs

func WorkspaceLoggingConfigurationPtr(v *WorkspaceLoggingConfigurationArgs) WorkspaceLoggingConfigurationPtrInput {
	return (*workspaceLoggingConfigurationPtrType)(v)
}

func (*workspaceLoggingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (i *workspaceLoggingConfigurationPtrType) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return i.ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i *workspaceLoggingConfigurationPtrType) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceLoggingConfigurationPtrOutput)
}

type WorkspaceLoggingConfigurationOutput struct{ *pulumi.OutputState }

func (WorkspaceLoggingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationOutput() WorkspaceLoggingConfigurationOutput {
	return o
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationOutput {
	return o
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return o.ToWorkspaceLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (o WorkspaceLoggingConfigurationOutput) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceLoggingConfiguration) *WorkspaceLoggingConfiguration {
		return &v
	}).(WorkspaceLoggingConfigurationPtrOutput)
}

// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
func (o WorkspaceLoggingConfigurationOutput) LogGroupArn() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceLoggingConfiguration) string { return v.LogGroupArn }).(pulumi.StringOutput)
}

type WorkspaceLoggingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceLoggingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceLoggingConfiguration)(nil)).Elem()
}

func (o WorkspaceLoggingConfigurationPtrOutput) ToWorkspaceLoggingConfigurationPtrOutput() WorkspaceLoggingConfigurationPtrOutput {
	return o
}

func (o WorkspaceLoggingConfigurationPtrOutput) ToWorkspaceLoggingConfigurationPtrOutputWithContext(ctx context.Context) WorkspaceLoggingConfigurationPtrOutput {
	return o
}

func (o WorkspaceLoggingConfigurationPtrOutput) Elem() WorkspaceLoggingConfigurationOutput {
	return o.ApplyT(func(v *WorkspaceLoggingConfiguration) WorkspaceLoggingConfiguration {
		if v != nil {
			return *v
		}
		var ret WorkspaceLoggingConfiguration
		return ret
	}).(WorkspaceLoggingConfigurationOutput)
}

// The ARN of the CloudWatch log group to which the vended log data will be published. This log group must exist.
func (o WorkspaceLoggingConfigurationPtrOutput) LogGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceLoggingConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.LogGroupArn
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperDestinationInput)(nil)).Elem(), ScraperDestinationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperDestinationPtrInput)(nil)).Elem(), ScraperDestinationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperDestinationAmpInput)(nil)).Elem(), ScraperDestinationAmpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperDestinationAmpPtrInput)(nil)).Elem(), ScraperDestinationAmpArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperSourceInput)(nil)).Elem(), ScraperSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperSourcePtrInput)(nil)).Elem(), ScraperSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperSourceEksInput)(nil)).Elem(), ScraperSourceEksArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperSourceEksPtrInput)(nil)).Elem(), ScraperSourceEksArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperTimeoutsInput)(nil)).Elem(), ScraperTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScraperTimeoutsPtrInput)(nil)).Elem(), ScraperTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceLoggingConfigurationInput)(nil)).Elem(), WorkspaceLoggingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceLoggingConfigurationPtrInput)(nil)).Elem(), WorkspaceLoggingConfigurationArgs{})
	pulumi.RegisterOutputType(ScraperDestinationOutput{})
	pulumi.RegisterOutputType(ScraperDestinationPtrOutput{})
	pulumi.RegisterOutputType(ScraperDestinationAmpOutput{})
	pulumi.RegisterOutputType(ScraperDestinationAmpPtrOutput{})
	pulumi.RegisterOutputType(ScraperSourceOutput{})
	pulumi.RegisterOutputType(ScraperSourcePtrOutput{})
	pulumi.RegisterOutputType(ScraperSourceEksOutput{})
	pulumi.RegisterOutputType(ScraperSourceEksPtrOutput{})
	pulumi.RegisterOutputType(ScraperTimeoutsOutput{})
	pulumi.RegisterOutputType(ScraperTimeoutsPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceLoggingConfigurationOutput{})
	pulumi.RegisterOutputType(WorkspaceLoggingConfigurationPtrOutput{})
}
