// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an RDS DB Reserved Instance.
//
// > **NOTE:** Once created, a reservation is valid for the `duration` of the provided `offeringId` and cannot be deleted. Performing a `destroy` will only remove the resource from state. For more information see [RDS Reserved Instances Documentation](https://aws.amazon.com/rds/reserved-instances/) and [PurchaseReservedDBInstancesOffering](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_PurchaseReservedDBInstancesOffering.html).
//
// > **NOTE:** Due to the expense of testing this resource, we provide it as best effort. If you find it useful, and have the ability to help test or notice issues, consider reaching out to us on GitHub.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := rds.GetReservedInstanceOffering(ctx, &rds.GetReservedInstanceOfferingArgs{
//				DbInstanceClass:    "db.t2.micro",
//				Duration:           31536000,
//				MultiAz:            false,
//				OfferingType:       "All Upfront",
//				ProductDescription: "mysql",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewReservedInstance(ctx, "example", &rds.ReservedInstanceArgs{
//				OfferingId:    *pulumi.String(test.OfferingId),
//				ReservationId: pulumi.String("optionalCustomReservationID"),
//				InstanceCount: pulumi.Int(3),
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
// RDS DB Instance Reservations can be imported using the `instance_id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:rds/reservedInstance:ReservedInstance reservation_instance CustomReservationID
//
// ```
type ReservedInstance struct {
	pulumi.CustomResourceState

	// ARN for the reserved DB instance.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Currency code for the reserved DB instance.
	CurrencyCode pulumi.StringOutput `pulumi:"currencyCode"`
	// DB instance class for the reserved DB instance.
	DbInstanceClass pulumi.StringOutput `pulumi:"dbInstanceClass"`
	// Duration of the reservation in seconds.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// Fixed price charged for this reserved DB instance.
	FixedPrice pulumi.Float64Output `pulumi:"fixedPrice"`
	// Number of instances to reserve. Default value is `1`.
	InstanceCount pulumi.IntPtrOutput `pulumi:"instanceCount"`
	// Unique identifier for the lease associated with the reserved DB instance. Amazon Web Services Support might request the lease ID for an issue related to a reserved DB instance.
	LeaseId pulumi.StringOutput `pulumi:"leaseId"`
	// Whether the reservation applies to Multi-AZ deployments.
	MultiAz pulumi.BoolOutput `pulumi:"multiAz"`
	// ID of the Reserved DB instance offering to purchase. To determine an `offeringId`, see the `rds.getReservedInstanceOffering` data source.
	OfferingId pulumi.StringOutput `pulumi:"offeringId"`
	// Offering type of this reserved DB instance.
	OfferingType pulumi.StringOutput `pulumi:"offeringType"`
	// Description of the reserved DB instance.
	ProductDescription pulumi.StringOutput `pulumi:"productDescription"`
	// Recurring price charged to run this reserved DB instance.
	RecurringCharges ReservedInstanceRecurringChargeArrayOutput `pulumi:"recurringCharges"`
	// Customer-specified identifier to track this reservation.
	ReservationId pulumi.StringPtrOutput `pulumi:"reservationId"`
	// Time the reservation started.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// State of the reserved DB instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Map of tags to assign to the DB reservation. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Hourly price charged for this reserved DB instance.
	UsagePrice pulumi.Float64Output `pulumi:"usagePrice"`
}

// NewReservedInstance registers a new resource with the given unique name, arguments, and options.
func NewReservedInstance(ctx *pulumi.Context,
	name string, args *ReservedInstanceArgs, opts ...pulumi.ResourceOption) (*ReservedInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OfferingId == nil {
		return nil, errors.New("invalid value for required argument 'OfferingId'")
	}
	var resource ReservedInstance
	err := ctx.RegisterResource("aws:rds/reservedInstance:ReservedInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservedInstance gets an existing ReservedInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservedInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservedInstanceState, opts ...pulumi.ResourceOption) (*ReservedInstance, error) {
	var resource ReservedInstance
	err := ctx.ReadResource("aws:rds/reservedInstance:ReservedInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReservedInstance resources.
type reservedInstanceState struct {
	// ARN for the reserved DB instance.
	Arn *string `pulumi:"arn"`
	// Currency code for the reserved DB instance.
	CurrencyCode *string `pulumi:"currencyCode"`
	// DB instance class for the reserved DB instance.
	DbInstanceClass *string `pulumi:"dbInstanceClass"`
	// Duration of the reservation in seconds.
	Duration *int `pulumi:"duration"`
	// Fixed price charged for this reserved DB instance.
	FixedPrice *float64 `pulumi:"fixedPrice"`
	// Number of instances to reserve. Default value is `1`.
	InstanceCount *int `pulumi:"instanceCount"`
	// Unique identifier for the lease associated with the reserved DB instance. Amazon Web Services Support might request the lease ID for an issue related to a reserved DB instance.
	LeaseId *string `pulumi:"leaseId"`
	// Whether the reservation applies to Multi-AZ deployments.
	MultiAz *bool `pulumi:"multiAz"`
	// ID of the Reserved DB instance offering to purchase. To determine an `offeringId`, see the `rds.getReservedInstanceOffering` data source.
	OfferingId *string `pulumi:"offeringId"`
	// Offering type of this reserved DB instance.
	OfferingType *string `pulumi:"offeringType"`
	// Description of the reserved DB instance.
	ProductDescription *string `pulumi:"productDescription"`
	// Recurring price charged to run this reserved DB instance.
	RecurringCharges []ReservedInstanceRecurringCharge `pulumi:"recurringCharges"`
	// Customer-specified identifier to track this reservation.
	ReservationId *string `pulumi:"reservationId"`
	// Time the reservation started.
	StartTime *string `pulumi:"startTime"`
	// State of the reserved DB instance.
	State *string `pulumi:"state"`
	// Map of tags to assign to the DB reservation. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Hourly price charged for this reserved DB instance.
	UsagePrice *float64 `pulumi:"usagePrice"`
}

type ReservedInstanceState struct {
	// ARN for the reserved DB instance.
	Arn pulumi.StringPtrInput
	// Currency code for the reserved DB instance.
	CurrencyCode pulumi.StringPtrInput
	// DB instance class for the reserved DB instance.
	DbInstanceClass pulumi.StringPtrInput
	// Duration of the reservation in seconds.
	Duration pulumi.IntPtrInput
	// Fixed price charged for this reserved DB instance.
	FixedPrice pulumi.Float64PtrInput
	// Number of instances to reserve. Default value is `1`.
	InstanceCount pulumi.IntPtrInput
	// Unique identifier for the lease associated with the reserved DB instance. Amazon Web Services Support might request the lease ID for an issue related to a reserved DB instance.
	LeaseId pulumi.StringPtrInput
	// Whether the reservation applies to Multi-AZ deployments.
	MultiAz pulumi.BoolPtrInput
	// ID of the Reserved DB instance offering to purchase. To determine an `offeringId`, see the `rds.getReservedInstanceOffering` data source.
	OfferingId pulumi.StringPtrInput
	// Offering type of this reserved DB instance.
	OfferingType pulumi.StringPtrInput
	// Description of the reserved DB instance.
	ProductDescription pulumi.StringPtrInput
	// Recurring price charged to run this reserved DB instance.
	RecurringCharges ReservedInstanceRecurringChargeArrayInput
	// Customer-specified identifier to track this reservation.
	ReservationId pulumi.StringPtrInput
	// Time the reservation started.
	StartTime pulumi.StringPtrInput
	// State of the reserved DB instance.
	State pulumi.StringPtrInput
	// Map of tags to assign to the DB reservation. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Hourly price charged for this reserved DB instance.
	UsagePrice pulumi.Float64PtrInput
}

func (ReservedInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservedInstanceState)(nil)).Elem()
}

type reservedInstanceArgs struct {
	// Number of instances to reserve. Default value is `1`.
	InstanceCount *int `pulumi:"instanceCount"`
	// ID of the Reserved DB instance offering to purchase. To determine an `offeringId`, see the `rds.getReservedInstanceOffering` data source.
	OfferingId string `pulumi:"offeringId"`
	// Customer-specified identifier to track this reservation.
	ReservationId *string `pulumi:"reservationId"`
	// Map of tags to assign to the DB reservation. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ReservedInstance resource.
type ReservedInstanceArgs struct {
	// Number of instances to reserve. Default value is `1`.
	InstanceCount pulumi.IntPtrInput
	// ID of the Reserved DB instance offering to purchase. To determine an `offeringId`, see the `rds.getReservedInstanceOffering` data source.
	OfferingId pulumi.StringInput
	// Customer-specified identifier to track this reservation.
	ReservationId pulumi.StringPtrInput
	// Map of tags to assign to the DB reservation. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ReservedInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservedInstanceArgs)(nil)).Elem()
}

type ReservedInstanceInput interface {
	pulumi.Input

	ToReservedInstanceOutput() ReservedInstanceOutput
	ToReservedInstanceOutputWithContext(ctx context.Context) ReservedInstanceOutput
}

func (*ReservedInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservedInstance)(nil)).Elem()
}

func (i *ReservedInstance) ToReservedInstanceOutput() ReservedInstanceOutput {
	return i.ToReservedInstanceOutputWithContext(context.Background())
}

func (i *ReservedInstance) ToReservedInstanceOutputWithContext(ctx context.Context) ReservedInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedInstanceOutput)
}

// ReservedInstanceArrayInput is an input type that accepts ReservedInstanceArray and ReservedInstanceArrayOutput values.
// You can construct a concrete instance of `ReservedInstanceArrayInput` via:
//
//	ReservedInstanceArray{ ReservedInstanceArgs{...} }
type ReservedInstanceArrayInput interface {
	pulumi.Input

	ToReservedInstanceArrayOutput() ReservedInstanceArrayOutput
	ToReservedInstanceArrayOutputWithContext(context.Context) ReservedInstanceArrayOutput
}

type ReservedInstanceArray []ReservedInstanceInput

func (ReservedInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReservedInstance)(nil)).Elem()
}

func (i ReservedInstanceArray) ToReservedInstanceArrayOutput() ReservedInstanceArrayOutput {
	return i.ToReservedInstanceArrayOutputWithContext(context.Background())
}

func (i ReservedInstanceArray) ToReservedInstanceArrayOutputWithContext(ctx context.Context) ReservedInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedInstanceArrayOutput)
}

// ReservedInstanceMapInput is an input type that accepts ReservedInstanceMap and ReservedInstanceMapOutput values.
// You can construct a concrete instance of `ReservedInstanceMapInput` via:
//
//	ReservedInstanceMap{ "key": ReservedInstanceArgs{...} }
type ReservedInstanceMapInput interface {
	pulumi.Input

	ToReservedInstanceMapOutput() ReservedInstanceMapOutput
	ToReservedInstanceMapOutputWithContext(context.Context) ReservedInstanceMapOutput
}

type ReservedInstanceMap map[string]ReservedInstanceInput

func (ReservedInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReservedInstance)(nil)).Elem()
}

func (i ReservedInstanceMap) ToReservedInstanceMapOutput() ReservedInstanceMapOutput {
	return i.ToReservedInstanceMapOutputWithContext(context.Background())
}

func (i ReservedInstanceMap) ToReservedInstanceMapOutputWithContext(ctx context.Context) ReservedInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedInstanceMapOutput)
}

type ReservedInstanceOutput struct{ *pulumi.OutputState }

func (ReservedInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservedInstance)(nil)).Elem()
}

func (o ReservedInstanceOutput) ToReservedInstanceOutput() ReservedInstanceOutput {
	return o
}

func (o ReservedInstanceOutput) ToReservedInstanceOutputWithContext(ctx context.Context) ReservedInstanceOutput {
	return o
}

// ARN for the reserved DB instance.
func (o ReservedInstanceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Currency code for the reserved DB instance.
func (o ReservedInstanceOutput) CurrencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.CurrencyCode }).(pulumi.StringOutput)
}

// DB instance class for the reserved DB instance.
func (o ReservedInstanceOutput) DbInstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.DbInstanceClass }).(pulumi.StringOutput)
}

// Duration of the reservation in seconds.
func (o ReservedInstanceOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.IntOutput { return v.Duration }).(pulumi.IntOutput)
}

// Fixed price charged for this reserved DB instance.
func (o ReservedInstanceOutput) FixedPrice() pulumi.Float64Output {
	return o.ApplyT(func(v *ReservedInstance) pulumi.Float64Output { return v.FixedPrice }).(pulumi.Float64Output)
}

// Number of instances to reserve. Default value is `1`.
func (o ReservedInstanceOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.IntPtrOutput { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

// Unique identifier for the lease associated with the reserved DB instance. Amazon Web Services Support might request the lease ID for an issue related to a reserved DB instance.
func (o ReservedInstanceOutput) LeaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.LeaseId }).(pulumi.StringOutput)
}

// Whether the reservation applies to Multi-AZ deployments.
func (o ReservedInstanceOutput) MultiAz() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.BoolOutput { return v.MultiAz }).(pulumi.BoolOutput)
}

// ID of the Reserved DB instance offering to purchase. To determine an `offeringId`, see the `rds.getReservedInstanceOffering` data source.
func (o ReservedInstanceOutput) OfferingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.OfferingId }).(pulumi.StringOutput)
}

// Offering type of this reserved DB instance.
func (o ReservedInstanceOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.OfferingType }).(pulumi.StringOutput)
}

// Description of the reserved DB instance.
func (o ReservedInstanceOutput) ProductDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.ProductDescription }).(pulumi.StringOutput)
}

// Recurring price charged to run this reserved DB instance.
func (o ReservedInstanceOutput) RecurringCharges() ReservedInstanceRecurringChargeArrayOutput {
	return o.ApplyT(func(v *ReservedInstance) ReservedInstanceRecurringChargeArrayOutput { return v.RecurringCharges }).(ReservedInstanceRecurringChargeArrayOutput)
}

// Customer-specified identifier to track this reservation.
func (o ReservedInstanceOutput) ReservationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringPtrOutput { return v.ReservationId }).(pulumi.StringPtrOutput)
}

// Time the reservation started.
func (o ReservedInstanceOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// State of the reserved DB instance.
func (o ReservedInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Map of tags to assign to the DB reservation. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ReservedInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ReservedInstanceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReservedInstance) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Hourly price charged for this reserved DB instance.
func (o ReservedInstanceOutput) UsagePrice() pulumi.Float64Output {
	return o.ApplyT(func(v *ReservedInstance) pulumi.Float64Output { return v.UsagePrice }).(pulumi.Float64Output)
}

type ReservedInstanceArrayOutput struct{ *pulumi.OutputState }

func (ReservedInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReservedInstance)(nil)).Elem()
}

func (o ReservedInstanceArrayOutput) ToReservedInstanceArrayOutput() ReservedInstanceArrayOutput {
	return o
}

func (o ReservedInstanceArrayOutput) ToReservedInstanceArrayOutputWithContext(ctx context.Context) ReservedInstanceArrayOutput {
	return o
}

func (o ReservedInstanceArrayOutput) Index(i pulumi.IntInput) ReservedInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReservedInstance {
		return vs[0].([]*ReservedInstance)[vs[1].(int)]
	}).(ReservedInstanceOutput)
}

type ReservedInstanceMapOutput struct{ *pulumi.OutputState }

func (ReservedInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReservedInstance)(nil)).Elem()
}

func (o ReservedInstanceMapOutput) ToReservedInstanceMapOutput() ReservedInstanceMapOutput {
	return o
}

func (o ReservedInstanceMapOutput) ToReservedInstanceMapOutputWithContext(ctx context.Context) ReservedInstanceMapOutput {
	return o
}

func (o ReservedInstanceMapOutput) MapIndex(k pulumi.StringInput) ReservedInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReservedInstance {
		return vs[0].(map[string]*ReservedInstance)[vs[1].(string)]
	}).(ReservedInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedInstanceInput)(nil)).Elem(), &ReservedInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedInstanceArrayInput)(nil)).Elem(), ReservedInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedInstanceMapInput)(nil)).Elem(), ReservedInstanceMap{})
	pulumi.RegisterOutputType(ReservedInstanceOutput{})
	pulumi.RegisterOutputType(ReservedInstanceArrayOutput{})
	pulumi.RegisterOutputType(ReservedInstanceMapOutput{})
}
