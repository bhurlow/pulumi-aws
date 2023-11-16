// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package shield

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ApplicationLayerAutomaticResponseTimeouts struct {
	Create *string `pulumi:"create"`
	Delete *string `pulumi:"delete"`
	Update *string `pulumi:"update"`
}

// ApplicationLayerAutomaticResponseTimeoutsInput is an input type that accepts ApplicationLayerAutomaticResponseTimeoutsArgs and ApplicationLayerAutomaticResponseTimeoutsOutput values.
// You can construct a concrete instance of `ApplicationLayerAutomaticResponseTimeoutsInput` via:
//
//	ApplicationLayerAutomaticResponseTimeoutsArgs{...}
type ApplicationLayerAutomaticResponseTimeoutsInput interface {
	pulumi.Input

	ToApplicationLayerAutomaticResponseTimeoutsOutput() ApplicationLayerAutomaticResponseTimeoutsOutput
	ToApplicationLayerAutomaticResponseTimeoutsOutputWithContext(context.Context) ApplicationLayerAutomaticResponseTimeoutsOutput
}

type ApplicationLayerAutomaticResponseTimeoutsArgs struct {
	Create pulumi.StringPtrInput `pulumi:"create"`
	Delete pulumi.StringPtrInput `pulumi:"delete"`
	Update pulumi.StringPtrInput `pulumi:"update"`
}

func (ApplicationLayerAutomaticResponseTimeoutsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLayerAutomaticResponseTimeouts)(nil)).Elem()
}

func (i ApplicationLayerAutomaticResponseTimeoutsArgs) ToApplicationLayerAutomaticResponseTimeoutsOutput() ApplicationLayerAutomaticResponseTimeoutsOutput {
	return i.ToApplicationLayerAutomaticResponseTimeoutsOutputWithContext(context.Background())
}

func (i ApplicationLayerAutomaticResponseTimeoutsArgs) ToApplicationLayerAutomaticResponseTimeoutsOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseTimeoutsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLayerAutomaticResponseTimeoutsOutput)
}

func (i ApplicationLayerAutomaticResponseTimeoutsArgs) ToApplicationLayerAutomaticResponseTimeoutsPtrOutput() ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return i.ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(context.Background())
}

func (i ApplicationLayerAutomaticResponseTimeoutsArgs) ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLayerAutomaticResponseTimeoutsOutput).ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(ctx)
}

// ApplicationLayerAutomaticResponseTimeoutsPtrInput is an input type that accepts ApplicationLayerAutomaticResponseTimeoutsArgs, ApplicationLayerAutomaticResponseTimeoutsPtr and ApplicationLayerAutomaticResponseTimeoutsPtrOutput values.
// You can construct a concrete instance of `ApplicationLayerAutomaticResponseTimeoutsPtrInput` via:
//
//	        ApplicationLayerAutomaticResponseTimeoutsArgs{...}
//
//	or:
//
//	        nil
type ApplicationLayerAutomaticResponseTimeoutsPtrInput interface {
	pulumi.Input

	ToApplicationLayerAutomaticResponseTimeoutsPtrOutput() ApplicationLayerAutomaticResponseTimeoutsPtrOutput
	ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(context.Context) ApplicationLayerAutomaticResponseTimeoutsPtrOutput
}

type applicationLayerAutomaticResponseTimeoutsPtrType ApplicationLayerAutomaticResponseTimeoutsArgs

func ApplicationLayerAutomaticResponseTimeoutsPtr(v *ApplicationLayerAutomaticResponseTimeoutsArgs) ApplicationLayerAutomaticResponseTimeoutsPtrInput {
	return (*applicationLayerAutomaticResponseTimeoutsPtrType)(v)
}

func (*applicationLayerAutomaticResponseTimeoutsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLayerAutomaticResponseTimeouts)(nil)).Elem()
}

func (i *applicationLayerAutomaticResponseTimeoutsPtrType) ToApplicationLayerAutomaticResponseTimeoutsPtrOutput() ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return i.ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(context.Background())
}

func (i *applicationLayerAutomaticResponseTimeoutsPtrType) ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLayerAutomaticResponseTimeoutsPtrOutput)
}

type ApplicationLayerAutomaticResponseTimeoutsOutput struct{ *pulumi.OutputState }

func (ApplicationLayerAutomaticResponseTimeoutsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLayerAutomaticResponseTimeouts)(nil)).Elem()
}

func (o ApplicationLayerAutomaticResponseTimeoutsOutput) ToApplicationLayerAutomaticResponseTimeoutsOutput() ApplicationLayerAutomaticResponseTimeoutsOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseTimeoutsOutput) ToApplicationLayerAutomaticResponseTimeoutsOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseTimeoutsOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseTimeoutsOutput) ToApplicationLayerAutomaticResponseTimeoutsPtrOutput() ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return o.ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(context.Background())
}

func (o ApplicationLayerAutomaticResponseTimeoutsOutput) ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationLayerAutomaticResponseTimeouts) *ApplicationLayerAutomaticResponseTimeouts {
		return &v
	}).(ApplicationLayerAutomaticResponseTimeoutsPtrOutput)
}

func (o ApplicationLayerAutomaticResponseTimeoutsOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationLayerAutomaticResponseTimeouts) *string { return v.Create }).(pulumi.StringPtrOutput)
}

func (o ApplicationLayerAutomaticResponseTimeoutsOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationLayerAutomaticResponseTimeouts) *string { return v.Delete }).(pulumi.StringPtrOutput)
}

func (o ApplicationLayerAutomaticResponseTimeoutsOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationLayerAutomaticResponseTimeouts) *string { return v.Update }).(pulumi.StringPtrOutput)
}

type ApplicationLayerAutomaticResponseTimeoutsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationLayerAutomaticResponseTimeoutsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLayerAutomaticResponseTimeouts)(nil)).Elem()
}

func (o ApplicationLayerAutomaticResponseTimeoutsPtrOutput) ToApplicationLayerAutomaticResponseTimeoutsPtrOutput() ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseTimeoutsPtrOutput) ToApplicationLayerAutomaticResponseTimeoutsPtrOutputWithContext(ctx context.Context) ApplicationLayerAutomaticResponseTimeoutsPtrOutput {
	return o
}

func (o ApplicationLayerAutomaticResponseTimeoutsPtrOutput) Elem() ApplicationLayerAutomaticResponseTimeoutsOutput {
	return o.ApplyT(func(v *ApplicationLayerAutomaticResponseTimeouts) ApplicationLayerAutomaticResponseTimeouts {
		if v != nil {
			return *v
		}
		var ret ApplicationLayerAutomaticResponseTimeouts
		return ret
	}).(ApplicationLayerAutomaticResponseTimeoutsOutput)
}

func (o ApplicationLayerAutomaticResponseTimeoutsPtrOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationLayerAutomaticResponseTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Create
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationLayerAutomaticResponseTimeoutsPtrOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationLayerAutomaticResponseTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationLayerAutomaticResponseTimeoutsPtrOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationLayerAutomaticResponseTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Update
	}).(pulumi.StringPtrOutput)
}

type DrtAccessLogBucketAssociationTimeouts struct {
	Create *string `pulumi:"create"`
	Delete *string `pulumi:"delete"`
	Read   *string `pulumi:"read"`
}

// DrtAccessLogBucketAssociationTimeoutsInput is an input type that accepts DrtAccessLogBucketAssociationTimeoutsArgs and DrtAccessLogBucketAssociationTimeoutsOutput values.
// You can construct a concrete instance of `DrtAccessLogBucketAssociationTimeoutsInput` via:
//
//	DrtAccessLogBucketAssociationTimeoutsArgs{...}
type DrtAccessLogBucketAssociationTimeoutsInput interface {
	pulumi.Input

	ToDrtAccessLogBucketAssociationTimeoutsOutput() DrtAccessLogBucketAssociationTimeoutsOutput
	ToDrtAccessLogBucketAssociationTimeoutsOutputWithContext(context.Context) DrtAccessLogBucketAssociationTimeoutsOutput
}

type DrtAccessLogBucketAssociationTimeoutsArgs struct {
	Create pulumi.StringPtrInput `pulumi:"create"`
	Delete pulumi.StringPtrInput `pulumi:"delete"`
	Read   pulumi.StringPtrInput `pulumi:"read"`
}

func (DrtAccessLogBucketAssociationTimeoutsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DrtAccessLogBucketAssociationTimeouts)(nil)).Elem()
}

func (i DrtAccessLogBucketAssociationTimeoutsArgs) ToDrtAccessLogBucketAssociationTimeoutsOutput() DrtAccessLogBucketAssociationTimeoutsOutput {
	return i.ToDrtAccessLogBucketAssociationTimeoutsOutputWithContext(context.Background())
}

func (i DrtAccessLogBucketAssociationTimeoutsArgs) ToDrtAccessLogBucketAssociationTimeoutsOutputWithContext(ctx context.Context) DrtAccessLogBucketAssociationTimeoutsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrtAccessLogBucketAssociationTimeoutsOutput)
}

func (i DrtAccessLogBucketAssociationTimeoutsArgs) ToDrtAccessLogBucketAssociationTimeoutsPtrOutput() DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return i.ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(context.Background())
}

func (i DrtAccessLogBucketAssociationTimeoutsArgs) ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrtAccessLogBucketAssociationTimeoutsOutput).ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(ctx)
}

// DrtAccessLogBucketAssociationTimeoutsPtrInput is an input type that accepts DrtAccessLogBucketAssociationTimeoutsArgs, DrtAccessLogBucketAssociationTimeoutsPtr and DrtAccessLogBucketAssociationTimeoutsPtrOutput values.
// You can construct a concrete instance of `DrtAccessLogBucketAssociationTimeoutsPtrInput` via:
//
//	        DrtAccessLogBucketAssociationTimeoutsArgs{...}
//
//	or:
//
//	        nil
type DrtAccessLogBucketAssociationTimeoutsPtrInput interface {
	pulumi.Input

	ToDrtAccessLogBucketAssociationTimeoutsPtrOutput() DrtAccessLogBucketAssociationTimeoutsPtrOutput
	ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(context.Context) DrtAccessLogBucketAssociationTimeoutsPtrOutput
}

type drtAccessLogBucketAssociationTimeoutsPtrType DrtAccessLogBucketAssociationTimeoutsArgs

func DrtAccessLogBucketAssociationTimeoutsPtr(v *DrtAccessLogBucketAssociationTimeoutsArgs) DrtAccessLogBucketAssociationTimeoutsPtrInput {
	return (*drtAccessLogBucketAssociationTimeoutsPtrType)(v)
}

func (*drtAccessLogBucketAssociationTimeoutsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DrtAccessLogBucketAssociationTimeouts)(nil)).Elem()
}

func (i *drtAccessLogBucketAssociationTimeoutsPtrType) ToDrtAccessLogBucketAssociationTimeoutsPtrOutput() DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return i.ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(context.Background())
}

func (i *drtAccessLogBucketAssociationTimeoutsPtrType) ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrtAccessLogBucketAssociationTimeoutsPtrOutput)
}

type DrtAccessLogBucketAssociationTimeoutsOutput struct{ *pulumi.OutputState }

func (DrtAccessLogBucketAssociationTimeoutsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DrtAccessLogBucketAssociationTimeouts)(nil)).Elem()
}

func (o DrtAccessLogBucketAssociationTimeoutsOutput) ToDrtAccessLogBucketAssociationTimeoutsOutput() DrtAccessLogBucketAssociationTimeoutsOutput {
	return o
}

func (o DrtAccessLogBucketAssociationTimeoutsOutput) ToDrtAccessLogBucketAssociationTimeoutsOutputWithContext(ctx context.Context) DrtAccessLogBucketAssociationTimeoutsOutput {
	return o
}

func (o DrtAccessLogBucketAssociationTimeoutsOutput) ToDrtAccessLogBucketAssociationTimeoutsPtrOutput() DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return o.ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(context.Background())
}

func (o DrtAccessLogBucketAssociationTimeoutsOutput) ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DrtAccessLogBucketAssociationTimeouts) *DrtAccessLogBucketAssociationTimeouts {
		return &v
	}).(DrtAccessLogBucketAssociationTimeoutsPtrOutput)
}

func (o DrtAccessLogBucketAssociationTimeoutsOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DrtAccessLogBucketAssociationTimeouts) *string { return v.Create }).(pulumi.StringPtrOutput)
}

func (o DrtAccessLogBucketAssociationTimeoutsOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DrtAccessLogBucketAssociationTimeouts) *string { return v.Delete }).(pulumi.StringPtrOutput)
}

func (o DrtAccessLogBucketAssociationTimeoutsOutput) Read() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DrtAccessLogBucketAssociationTimeouts) *string { return v.Read }).(pulumi.StringPtrOutput)
}

type DrtAccessLogBucketAssociationTimeoutsPtrOutput struct{ *pulumi.OutputState }

func (DrtAccessLogBucketAssociationTimeoutsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DrtAccessLogBucketAssociationTimeouts)(nil)).Elem()
}

func (o DrtAccessLogBucketAssociationTimeoutsPtrOutput) ToDrtAccessLogBucketAssociationTimeoutsPtrOutput() DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return o
}

func (o DrtAccessLogBucketAssociationTimeoutsPtrOutput) ToDrtAccessLogBucketAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessLogBucketAssociationTimeoutsPtrOutput {
	return o
}

func (o DrtAccessLogBucketAssociationTimeoutsPtrOutput) Elem() DrtAccessLogBucketAssociationTimeoutsOutput {
	return o.ApplyT(func(v *DrtAccessLogBucketAssociationTimeouts) DrtAccessLogBucketAssociationTimeouts {
		if v != nil {
			return *v
		}
		var ret DrtAccessLogBucketAssociationTimeouts
		return ret
	}).(DrtAccessLogBucketAssociationTimeoutsOutput)
}

func (o DrtAccessLogBucketAssociationTimeoutsPtrOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DrtAccessLogBucketAssociationTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Create
	}).(pulumi.StringPtrOutput)
}

func (o DrtAccessLogBucketAssociationTimeoutsPtrOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DrtAccessLogBucketAssociationTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(pulumi.StringPtrOutput)
}

func (o DrtAccessLogBucketAssociationTimeoutsPtrOutput) Read() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DrtAccessLogBucketAssociationTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Read
	}).(pulumi.StringPtrOutput)
}

type DrtAccessRoleArnAssociationTimeouts struct {
	Create *string `pulumi:"create"`
	Delete *string `pulumi:"delete"`
	Read   *string `pulumi:"read"`
}

// DrtAccessRoleArnAssociationTimeoutsInput is an input type that accepts DrtAccessRoleArnAssociationTimeoutsArgs and DrtAccessRoleArnAssociationTimeoutsOutput values.
// You can construct a concrete instance of `DrtAccessRoleArnAssociationTimeoutsInput` via:
//
//	DrtAccessRoleArnAssociationTimeoutsArgs{...}
type DrtAccessRoleArnAssociationTimeoutsInput interface {
	pulumi.Input

	ToDrtAccessRoleArnAssociationTimeoutsOutput() DrtAccessRoleArnAssociationTimeoutsOutput
	ToDrtAccessRoleArnAssociationTimeoutsOutputWithContext(context.Context) DrtAccessRoleArnAssociationTimeoutsOutput
}

type DrtAccessRoleArnAssociationTimeoutsArgs struct {
	Create pulumi.StringPtrInput `pulumi:"create"`
	Delete pulumi.StringPtrInput `pulumi:"delete"`
	Read   pulumi.StringPtrInput `pulumi:"read"`
}

func (DrtAccessRoleArnAssociationTimeoutsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DrtAccessRoleArnAssociationTimeouts)(nil)).Elem()
}

func (i DrtAccessRoleArnAssociationTimeoutsArgs) ToDrtAccessRoleArnAssociationTimeoutsOutput() DrtAccessRoleArnAssociationTimeoutsOutput {
	return i.ToDrtAccessRoleArnAssociationTimeoutsOutputWithContext(context.Background())
}

func (i DrtAccessRoleArnAssociationTimeoutsArgs) ToDrtAccessRoleArnAssociationTimeoutsOutputWithContext(ctx context.Context) DrtAccessRoleArnAssociationTimeoutsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrtAccessRoleArnAssociationTimeoutsOutput)
}

func (i DrtAccessRoleArnAssociationTimeoutsArgs) ToDrtAccessRoleArnAssociationTimeoutsPtrOutput() DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return i.ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(context.Background())
}

func (i DrtAccessRoleArnAssociationTimeoutsArgs) ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrtAccessRoleArnAssociationTimeoutsOutput).ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(ctx)
}

// DrtAccessRoleArnAssociationTimeoutsPtrInput is an input type that accepts DrtAccessRoleArnAssociationTimeoutsArgs, DrtAccessRoleArnAssociationTimeoutsPtr and DrtAccessRoleArnAssociationTimeoutsPtrOutput values.
// You can construct a concrete instance of `DrtAccessRoleArnAssociationTimeoutsPtrInput` via:
//
//	        DrtAccessRoleArnAssociationTimeoutsArgs{...}
//
//	or:
//
//	        nil
type DrtAccessRoleArnAssociationTimeoutsPtrInput interface {
	pulumi.Input

	ToDrtAccessRoleArnAssociationTimeoutsPtrOutput() DrtAccessRoleArnAssociationTimeoutsPtrOutput
	ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(context.Context) DrtAccessRoleArnAssociationTimeoutsPtrOutput
}

type drtAccessRoleArnAssociationTimeoutsPtrType DrtAccessRoleArnAssociationTimeoutsArgs

func DrtAccessRoleArnAssociationTimeoutsPtr(v *DrtAccessRoleArnAssociationTimeoutsArgs) DrtAccessRoleArnAssociationTimeoutsPtrInput {
	return (*drtAccessRoleArnAssociationTimeoutsPtrType)(v)
}

func (*drtAccessRoleArnAssociationTimeoutsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DrtAccessRoleArnAssociationTimeouts)(nil)).Elem()
}

func (i *drtAccessRoleArnAssociationTimeoutsPtrType) ToDrtAccessRoleArnAssociationTimeoutsPtrOutput() DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return i.ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(context.Background())
}

func (i *drtAccessRoleArnAssociationTimeoutsPtrType) ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrtAccessRoleArnAssociationTimeoutsPtrOutput)
}

type DrtAccessRoleArnAssociationTimeoutsOutput struct{ *pulumi.OutputState }

func (DrtAccessRoleArnAssociationTimeoutsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DrtAccessRoleArnAssociationTimeouts)(nil)).Elem()
}

func (o DrtAccessRoleArnAssociationTimeoutsOutput) ToDrtAccessRoleArnAssociationTimeoutsOutput() DrtAccessRoleArnAssociationTimeoutsOutput {
	return o
}

func (o DrtAccessRoleArnAssociationTimeoutsOutput) ToDrtAccessRoleArnAssociationTimeoutsOutputWithContext(ctx context.Context) DrtAccessRoleArnAssociationTimeoutsOutput {
	return o
}

func (o DrtAccessRoleArnAssociationTimeoutsOutput) ToDrtAccessRoleArnAssociationTimeoutsPtrOutput() DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return o.ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(context.Background())
}

func (o DrtAccessRoleArnAssociationTimeoutsOutput) ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DrtAccessRoleArnAssociationTimeouts) *DrtAccessRoleArnAssociationTimeouts {
		return &v
	}).(DrtAccessRoleArnAssociationTimeoutsPtrOutput)
}

func (o DrtAccessRoleArnAssociationTimeoutsOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DrtAccessRoleArnAssociationTimeouts) *string { return v.Create }).(pulumi.StringPtrOutput)
}

func (o DrtAccessRoleArnAssociationTimeoutsOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DrtAccessRoleArnAssociationTimeouts) *string { return v.Delete }).(pulumi.StringPtrOutput)
}

func (o DrtAccessRoleArnAssociationTimeoutsOutput) Read() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DrtAccessRoleArnAssociationTimeouts) *string { return v.Read }).(pulumi.StringPtrOutput)
}

type DrtAccessRoleArnAssociationTimeoutsPtrOutput struct{ *pulumi.OutputState }

func (DrtAccessRoleArnAssociationTimeoutsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DrtAccessRoleArnAssociationTimeouts)(nil)).Elem()
}

func (o DrtAccessRoleArnAssociationTimeoutsPtrOutput) ToDrtAccessRoleArnAssociationTimeoutsPtrOutput() DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return o
}

func (o DrtAccessRoleArnAssociationTimeoutsPtrOutput) ToDrtAccessRoleArnAssociationTimeoutsPtrOutputWithContext(ctx context.Context) DrtAccessRoleArnAssociationTimeoutsPtrOutput {
	return o
}

func (o DrtAccessRoleArnAssociationTimeoutsPtrOutput) Elem() DrtAccessRoleArnAssociationTimeoutsOutput {
	return o.ApplyT(func(v *DrtAccessRoleArnAssociationTimeouts) DrtAccessRoleArnAssociationTimeouts {
		if v != nil {
			return *v
		}
		var ret DrtAccessRoleArnAssociationTimeouts
		return ret
	}).(DrtAccessRoleArnAssociationTimeoutsOutput)
}

func (o DrtAccessRoleArnAssociationTimeoutsPtrOutput) Create() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DrtAccessRoleArnAssociationTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Create
	}).(pulumi.StringPtrOutput)
}

func (o DrtAccessRoleArnAssociationTimeoutsPtrOutput) Delete() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DrtAccessRoleArnAssociationTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(pulumi.StringPtrOutput)
}

func (o DrtAccessRoleArnAssociationTimeoutsPtrOutput) Read() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DrtAccessRoleArnAssociationTimeouts) *string {
		if v == nil {
			return nil
		}
		return v.Read
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLayerAutomaticResponseTimeoutsInput)(nil)).Elem(), ApplicationLayerAutomaticResponseTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationLayerAutomaticResponseTimeoutsPtrInput)(nil)).Elem(), ApplicationLayerAutomaticResponseTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrtAccessLogBucketAssociationTimeoutsInput)(nil)).Elem(), DrtAccessLogBucketAssociationTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrtAccessLogBucketAssociationTimeoutsPtrInput)(nil)).Elem(), DrtAccessLogBucketAssociationTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrtAccessRoleArnAssociationTimeoutsInput)(nil)).Elem(), DrtAccessRoleArnAssociationTimeoutsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrtAccessRoleArnAssociationTimeoutsPtrInput)(nil)).Elem(), DrtAccessRoleArnAssociationTimeoutsArgs{})
	pulumi.RegisterOutputType(ApplicationLayerAutomaticResponseTimeoutsOutput{})
	pulumi.RegisterOutputType(ApplicationLayerAutomaticResponseTimeoutsPtrOutput{})
	pulumi.RegisterOutputType(DrtAccessLogBucketAssociationTimeoutsOutput{})
	pulumi.RegisterOutputType(DrtAccessLogBucketAssociationTimeoutsPtrOutput{})
	pulumi.RegisterOutputType(DrtAccessRoleArnAssociationTimeoutsOutput{})
	pulumi.RegisterOutputType(DrtAccessRoleArnAssociationTimeoutsPtrOutput{})
}
