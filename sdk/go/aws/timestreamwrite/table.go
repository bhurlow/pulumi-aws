// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timestreamwrite

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Timestream table resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := timestreamwrite.NewTable(ctx, "example", &timestreamwrite.TableArgs{
//				DatabaseName: pulumi.Any(aws_timestreamwrite_database.Example.Database_name),
//				TableName:    pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Full usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := timestreamwrite.NewTable(ctx, "example", &timestreamwrite.TableArgs{
//				DatabaseName: pulumi.Any(aws_timestreamwrite_database.Example.Database_name),
//				TableName:    pulumi.String("example"),
//				RetentionProperties: &timestreamwrite.TableRetentionPropertiesArgs{
//					MagneticStoreRetentionPeriodInDays: pulumi.Int(30),
//					MemoryStoreRetentionPeriodInHours:  pulumi.Int(8),
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("example-timestream-table"),
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
// ### Customer-defined Partition Key
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := timestreamwrite.NewTable(ctx, "example", &timestreamwrite.TableArgs{
//				DatabaseName: pulumi.Any(aws_timestreamwrite_database.Example.Database_name),
//				TableName:    pulumi.String("example"),
//				Schema: &timestreamwrite.TableSchemaArgs{
//					CompositePartitionKey: &timestreamwrite.TableSchemaCompositePartitionKeyArgs{
//						EnforcementInRecord: pulumi.String("REQUIRED"),
//						Name:                pulumi.String("attr1"),
//						Type:                pulumi.String("DIMENSION"),
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
// Using `pulumi import`, import Timestream tables using the `table_name` and `database_name` separate by a colon (`:`). For example:
//
// ```sh
//
//	$ pulumi import aws:timestreamwrite/table:Table example ExampleTable:ExampleDatabase
//
// ```
type Table struct {
	pulumi.CustomResourceState

	// The ARN that uniquely identifies this table.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the Timestream database.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties TableMagneticStoreWritePropertiesOutput `pulumi:"magneticStoreWriteProperties"`
	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
	RetentionProperties TableRetentionPropertiesOutput `pulumi:"retentionProperties"`
	// The schema of the table. See Schema below for more details.
	Schema TableSchemaOutput `pulumi:"schema"`
	// The name of the Timestream table.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Table
	err := ctx.RegisterResource("aws:timestreamwrite/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("aws:timestreamwrite/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// The ARN that uniquely identifies this table.
	Arn *string `pulumi:"arn"`
	// The name of the Timestream database.
	DatabaseName *string `pulumi:"databaseName"`
	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties *TableMagneticStoreWriteProperties `pulumi:"magneticStoreWriteProperties"`
	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
	RetentionProperties *TableRetentionProperties `pulumi:"retentionProperties"`
	// The schema of the table. See Schema below for more details.
	Schema *TableSchema `pulumi:"schema"`
	// The name of the Timestream table.
	TableName *string `pulumi:"tableName"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type TableState struct {
	// The ARN that uniquely identifies this table.
	Arn pulumi.StringPtrInput
	// The name of the Timestream database.
	DatabaseName pulumi.StringPtrInput
	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties TableMagneticStoreWritePropertiesPtrInput
	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
	RetentionProperties TableRetentionPropertiesPtrInput
	// The schema of the table. See Schema below for more details.
	Schema TableSchemaPtrInput
	// The name of the Timestream table.
	TableName pulumi.StringPtrInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// The name of the Timestream database.
	DatabaseName string `pulumi:"databaseName"`
	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties *TableMagneticStoreWriteProperties `pulumi:"magneticStoreWriteProperties"`
	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
	RetentionProperties *TableRetentionProperties `pulumi:"retentionProperties"`
	// The schema of the table. See Schema below for more details.
	Schema *TableSchema `pulumi:"schema"`
	// The name of the Timestream table.
	TableName string `pulumi:"tableName"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// The name of the Timestream database.
	DatabaseName pulumi.StringInput
	// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
	MagneticStoreWriteProperties TableMagneticStoreWritePropertiesPtrInput
	// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
	RetentionProperties TableRetentionPropertiesPtrInput
	// The schema of the table. See Schema below for more details.
	Schema TableSchemaPtrInput
	// The name of the Timestream table.
	TableName pulumi.StringInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

func (i *Table) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: i.ToTableOutputWithContext(ctx).OutputState,
	}
}

// TableArrayInput is an input type that accepts TableArray and TableArrayOutput values.
// You can construct a concrete instance of `TableArrayInput` via:
//
//	TableArray{ TableArgs{...} }
type TableArrayInput interface {
	pulumi.Input

	ToTableArrayOutput() TableArrayOutput
	ToTableArrayOutputWithContext(context.Context) TableArrayOutput
}

type TableArray []TableInput

func (TableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (i TableArray) ToTableArrayOutput() TableArrayOutput {
	return i.ToTableArrayOutputWithContext(context.Background())
}

func (i TableArray) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableArrayOutput)
}

func (i TableArray) ToOutput(ctx context.Context) pulumix.Output[[]*Table] {
	return pulumix.Output[[]*Table]{
		OutputState: i.ToTableArrayOutputWithContext(ctx).OutputState,
	}
}

// TableMapInput is an input type that accepts TableMap and TableMapOutput values.
// You can construct a concrete instance of `TableMapInput` via:
//
//	TableMap{ "key": TableArgs{...} }
type TableMapInput interface {
	pulumi.Input

	ToTableMapOutput() TableMapOutput
	ToTableMapOutputWithContext(context.Context) TableMapOutput
}

type TableMap map[string]TableInput

func (TableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (i TableMap) ToTableMapOutput() TableMapOutput {
	return i.ToTableMapOutputWithContext(context.Background())
}

func (i TableMap) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableMapOutput)
}

func (i TableMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Table] {
	return pulumix.Output[map[string]*Table]{
		OutputState: i.ToTableMapOutputWithContext(ctx).OutputState,
	}
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: o.OutputState,
	}
}

// The ARN that uniquely identifies this table.
func (o TableOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the Timestream database.
func (o TableOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
func (o TableOutput) MagneticStoreWriteProperties() TableMagneticStoreWritePropertiesOutput {
	return o.ApplyT(func(v *Table) TableMagneticStoreWritePropertiesOutput { return v.MagneticStoreWriteProperties }).(TableMagneticStoreWritePropertiesOutput)
}

// The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
func (o TableOutput) RetentionProperties() TableRetentionPropertiesOutput {
	return o.ApplyT(func(v *Table) TableRetentionPropertiesOutput { return v.RetentionProperties }).(TableRetentionPropertiesOutput)
}

// The schema of the table. See Schema below for more details.
func (o TableOutput) Schema() TableSchemaOutput {
	return o.ApplyT(func(v *Table) TableSchemaOutput { return v.Schema }).(TableSchemaOutput)
}

// The name of the Timestream table.
func (o TableOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Table) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o TableOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Table) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type TableArrayOutput struct{ *pulumi.OutputState }

func (TableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (o TableArrayOutput) ToTableArrayOutput() TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Table] {
	return pulumix.Output[[]*Table]{
		OutputState: o.OutputState,
	}
}

func (o TableArrayOutput) Index(i pulumi.IntInput) TableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Table {
		return vs[0].([]*Table)[vs[1].(int)]
	}).(TableOutput)
}

type TableMapOutput struct{ *pulumi.OutputState }

func (TableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (o TableMapOutput) ToTableMapOutput() TableMapOutput {
	return o
}

func (o TableMapOutput) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return o
}

func (o TableMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Table] {
	return pulumix.Output[map[string]*Table]{
		OutputState: o.OutputState,
	}
}

func (o TableMapOutput) MapIndex(k pulumi.StringInput) TableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Table {
		return vs[0].(map[string]*Table)[vs[1].(string)]
	}).(TableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableArrayInput)(nil)).Elem(), TableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableMapInput)(nil)).Elem(), TableMap{})
	pulumi.RegisterOutputType(TableOutput{})
	pulumi.RegisterOutputType(TableArrayOutput{})
	pulumi.RegisterOutputType(TableMapOutput{})
}
