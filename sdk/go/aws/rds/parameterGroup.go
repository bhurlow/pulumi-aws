// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an RDS DB parameter group resource. Documentation of the available parameters for various RDS engines can be found at:
//
// * [Aurora MySQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Reference.html)
// * [Aurora PostgreSQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraPostgreSQL.Reference.html)
// * [MariaDB Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Parameters.html)
// * [Oracle Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ModifyInstance.Oracle.html#USER_ModifyInstance.Oracle.sqlnet)
// * [PostgreSQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.html#Appendix.PostgreSQL.CommonDBATasks.Parameters)
//
// > **NOTE:** After applying your changes, you may encounter a perpetual diff in your preview
// output for a `parameter` whose `value` remains unchanged but whose `applyMethod` is changing
// (e.g., from `immediate` to `pending-reboot`, or `pending-reboot` to `immediate`). If only the
// apply method of a parameter is changing, the AWS API will not register this change. To change
// the `applyMethod` of a parameter, its value must also change.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.NewParameterGroup(ctx, "default", &rds.ParameterGroupArgs{
//				Family: pulumi.String("mysql5.6"),
//				Parameters: rds.ParameterGroupParameterArray{
//					&rds.ParameterGroupParameterArgs{
//						Name:  pulumi.String("character_set_server"),
//						Value: pulumi.String("utf8"),
//					},
//					&rds.ParameterGroupParameterArgs{
//						Name:  pulumi.String("character_set_client"),
//						Value: pulumi.String("utf8"),
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
// ### `createBeforeDestroy` Lifecycle Configuration
//
// The `createBeforeDestroy`
// lifecycle configuration is necessary for modifications that force re-creation of an existing,
// in-use parameter group. This includes common situations like changing the group `name` or
// bumping the `family` version during a major version upgrade. This configuration will prevent destruction
// of the deposed parameter group while still in use by the database during upgrade.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleParameterGroup, err := rds.NewParameterGroup(ctx, "exampleParameterGroup", &rds.ParameterGroupArgs{
//				Family: pulumi.String("postgres13"),
//				Parameters: rds.ParameterGroupParameterArray{
//					&rds.ParameterGroupParameterArgs{
//						Name:  pulumi.String("log_connections"),
//						Value: pulumi.String("1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewInstance(ctx, "exampleInstance", &rds.InstanceArgs{
//				ParameterGroupName: exampleParameterGroup.Name,
//				ApplyImmediately:   pulumi.Bool(true),
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
//	to = aws_db_parameter_group.rds_pg
//
//	id = "rds-pg" } Using `pulumi import`, import DB Parameter groups using the `name`. For exampleconsole % pulumi import aws_db_parameter_group.rds_pg rds-pg
type ParameterGroup struct {
	pulumi.CustomResourceState

	// The ARN of the db parameter group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the DB parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// The family of the DB parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the DB parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) after initial creation of the group.
	Parameters ParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewParameterGroup(ctx *pulumi.Context,
	name string, args *ParameterGroupArgs, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ParameterGroup
	err := ctx.RegisterResource("aws:rds/parameterGroup:ParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterGroup gets an existing ParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterGroupState, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	var resource ParameterGroup
	err := ctx.ReadResource("aws:rds/parameterGroup:ParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterGroup resources.
type parameterGroupState struct {
	// The ARN of the db parameter group.
	Arn *string `pulumi:"arn"`
	// The description of the DB parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the DB parameter group.
	Family *string `pulumi:"family"`
	// The name of the DB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) after initial creation of the group.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ParameterGroupState struct {
	// The ARN of the db parameter group.
	Arn pulumi.StringPtrInput
	// The description of the DB parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the DB parameter group.
	Family pulumi.StringPtrInput
	// The name of the DB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) after initial creation of the group.
	Parameters ParameterGroupParameterArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupState)(nil)).Elem()
}

type parameterGroupArgs struct {
	// The description of the DB parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the DB parameter group.
	Family string `pulumi:"family"`
	// The name of the DB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) after initial creation of the group.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	// The description of the DB parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the DB parameter group.
	Family pulumi.StringInput
	// The name of the DB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) after initial creation of the group.
	Parameters ParameterGroupParameterArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupArgs)(nil)).Elem()
}

type ParameterGroupInput interface {
	pulumi.Input

	ToParameterGroupOutput() ParameterGroupOutput
	ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput
}

func (*ParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterGroup)(nil)).Elem()
}

func (i *ParameterGroup) ToParameterGroupOutput() ParameterGroupOutput {
	return i.ToParameterGroupOutputWithContext(context.Background())
}

func (i *ParameterGroup) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupOutput)
}

// ParameterGroupArrayInput is an input type that accepts ParameterGroupArray and ParameterGroupArrayOutput values.
// You can construct a concrete instance of `ParameterGroupArrayInput` via:
//
//	ParameterGroupArray{ ParameterGroupArgs{...} }
type ParameterGroupArrayInput interface {
	pulumi.Input

	ToParameterGroupArrayOutput() ParameterGroupArrayOutput
	ToParameterGroupArrayOutputWithContext(context.Context) ParameterGroupArrayOutput
}

type ParameterGroupArray []ParameterGroupInput

func (ParameterGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterGroup)(nil)).Elem()
}

func (i ParameterGroupArray) ToParameterGroupArrayOutput() ParameterGroupArrayOutput {
	return i.ToParameterGroupArrayOutputWithContext(context.Background())
}

func (i ParameterGroupArray) ToParameterGroupArrayOutputWithContext(ctx context.Context) ParameterGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupArrayOutput)
}

// ParameterGroupMapInput is an input type that accepts ParameterGroupMap and ParameterGroupMapOutput values.
// You can construct a concrete instance of `ParameterGroupMapInput` via:
//
//	ParameterGroupMap{ "key": ParameterGroupArgs{...} }
type ParameterGroupMapInput interface {
	pulumi.Input

	ToParameterGroupMapOutput() ParameterGroupMapOutput
	ToParameterGroupMapOutputWithContext(context.Context) ParameterGroupMapOutput
}

type ParameterGroupMap map[string]ParameterGroupInput

func (ParameterGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterGroup)(nil)).Elem()
}

func (i ParameterGroupMap) ToParameterGroupMapOutput() ParameterGroupMapOutput {
	return i.ToParameterGroupMapOutputWithContext(context.Background())
}

func (i ParameterGroupMap) ToParameterGroupMapOutputWithContext(ctx context.Context) ParameterGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupMapOutput)
}

type ParameterGroupOutput struct{ *pulumi.OutputState }

func (ParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterGroup)(nil)).Elem()
}

func (o ParameterGroupOutput) ToParameterGroupOutput() ParameterGroupOutput {
	return o
}

func (o ParameterGroupOutput) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return o
}

// The ARN of the db parameter group.
func (o ParameterGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the DB parameter group. Defaults to "Managed by Pulumi".
func (o ParameterGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The family of the DB parameter group.
func (o ParameterGroupOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// The name of the DB parameter.
func (o ParameterGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o ParameterGroupOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// A list of DB parameters to apply. Note that parameters may differ from a family to an other. Full list of all parameters can be discovered via [`aws rds describe-db-parameters`](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) after initial creation of the group.
func (o ParameterGroupOutput) Parameters() ParameterGroupParameterArrayOutput {
	return o.ApplyT(func(v *ParameterGroup) ParameterGroupParameterArrayOutput { return v.Parameters }).(ParameterGroupParameterArrayOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ParameterGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ParameterGroupOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ParameterGroup) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type ParameterGroupArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterGroup)(nil)).Elem()
}

func (o ParameterGroupArrayOutput) ToParameterGroupArrayOutput() ParameterGroupArrayOutput {
	return o
}

func (o ParameterGroupArrayOutput) ToParameterGroupArrayOutputWithContext(ctx context.Context) ParameterGroupArrayOutput {
	return o
}

func (o ParameterGroupArrayOutput) Index(i pulumi.IntInput) ParameterGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ParameterGroup {
		return vs[0].([]*ParameterGroup)[vs[1].(int)]
	}).(ParameterGroupOutput)
}

type ParameterGroupMapOutput struct{ *pulumi.OutputState }

func (ParameterGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterGroup)(nil)).Elem()
}

func (o ParameterGroupMapOutput) ToParameterGroupMapOutput() ParameterGroupMapOutput {
	return o
}

func (o ParameterGroupMapOutput) ToParameterGroupMapOutputWithContext(ctx context.Context) ParameterGroupMapOutput {
	return o
}

func (o ParameterGroupMapOutput) MapIndex(k pulumi.StringInput) ParameterGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ParameterGroup {
		return vs[0].(map[string]*ParameterGroup)[vs[1].(string)]
	}).(ParameterGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupInput)(nil)).Elem(), &ParameterGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupArrayInput)(nil)).Elem(), ParameterGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupMapInput)(nil)).Elem(), ParameterGroupMap{})
	pulumi.RegisterOutputType(ParameterGroupOutput{})
	pulumi.RegisterOutputType(ParameterGroupArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupMapOutput{})
}
