// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsearch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type DomainEndpointOptions struct {
	// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
	EnforceHttps *bool `pulumi:"enforceHttps"`
	// The minimum required TLS version. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_DomainEndpointOptions.html) for valid values.
	TlsSecurityPolicy *string `pulumi:"tlsSecurityPolicy"`
}

// DomainEndpointOptionsInput is an input type that accepts DomainEndpointOptionsArgs and DomainEndpointOptionsOutput values.
// You can construct a concrete instance of `DomainEndpointOptionsInput` via:
//
//	DomainEndpointOptionsArgs{...}
type DomainEndpointOptionsInput interface {
	pulumi.Input

	ToDomainEndpointOptionsOutput() DomainEndpointOptionsOutput
	ToDomainEndpointOptionsOutputWithContext(context.Context) DomainEndpointOptionsOutput
}

type DomainEndpointOptionsArgs struct {
	// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
	EnforceHttps pulumi.BoolPtrInput `pulumi:"enforceHttps"`
	// The minimum required TLS version. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_DomainEndpointOptions.html) for valid values.
	TlsSecurityPolicy pulumi.StringPtrInput `pulumi:"tlsSecurityPolicy"`
}

func (DomainEndpointOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainEndpointOptions)(nil)).Elem()
}

func (i DomainEndpointOptionsArgs) ToDomainEndpointOptionsOutput() DomainEndpointOptionsOutput {
	return i.ToDomainEndpointOptionsOutputWithContext(context.Background())
}

func (i DomainEndpointOptionsArgs) ToDomainEndpointOptionsOutputWithContext(ctx context.Context) DomainEndpointOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainEndpointOptionsOutput)
}

func (i DomainEndpointOptionsArgs) ToDomainEndpointOptionsPtrOutput() DomainEndpointOptionsPtrOutput {
	return i.ToDomainEndpointOptionsPtrOutputWithContext(context.Background())
}

func (i DomainEndpointOptionsArgs) ToDomainEndpointOptionsPtrOutputWithContext(ctx context.Context) DomainEndpointOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainEndpointOptionsOutput).ToDomainEndpointOptionsPtrOutputWithContext(ctx)
}

// DomainEndpointOptionsPtrInput is an input type that accepts DomainEndpointOptionsArgs, DomainEndpointOptionsPtr and DomainEndpointOptionsPtrOutput values.
// You can construct a concrete instance of `DomainEndpointOptionsPtrInput` via:
//
//	        DomainEndpointOptionsArgs{...}
//
//	or:
//
//	        nil
type DomainEndpointOptionsPtrInput interface {
	pulumi.Input

	ToDomainEndpointOptionsPtrOutput() DomainEndpointOptionsPtrOutput
	ToDomainEndpointOptionsPtrOutputWithContext(context.Context) DomainEndpointOptionsPtrOutput
}

type domainEndpointOptionsPtrType DomainEndpointOptionsArgs

func DomainEndpointOptionsPtr(v *DomainEndpointOptionsArgs) DomainEndpointOptionsPtrInput {
	return (*domainEndpointOptionsPtrType)(v)
}

func (*domainEndpointOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainEndpointOptions)(nil)).Elem()
}

func (i *domainEndpointOptionsPtrType) ToDomainEndpointOptionsPtrOutput() DomainEndpointOptionsPtrOutput {
	return i.ToDomainEndpointOptionsPtrOutputWithContext(context.Background())
}

func (i *domainEndpointOptionsPtrType) ToDomainEndpointOptionsPtrOutputWithContext(ctx context.Context) DomainEndpointOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainEndpointOptionsPtrOutput)
}

type DomainEndpointOptionsOutput struct{ *pulumi.OutputState }

func (DomainEndpointOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainEndpointOptions)(nil)).Elem()
}

func (o DomainEndpointOptionsOutput) ToDomainEndpointOptionsOutput() DomainEndpointOptionsOutput {
	return o
}

func (o DomainEndpointOptionsOutput) ToDomainEndpointOptionsOutputWithContext(ctx context.Context) DomainEndpointOptionsOutput {
	return o
}

func (o DomainEndpointOptionsOutput) ToDomainEndpointOptionsPtrOutput() DomainEndpointOptionsPtrOutput {
	return o.ToDomainEndpointOptionsPtrOutputWithContext(context.Background())
}

func (o DomainEndpointOptionsOutput) ToDomainEndpointOptionsPtrOutputWithContext(ctx context.Context) DomainEndpointOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainEndpointOptions) *DomainEndpointOptions {
		return &v
	}).(DomainEndpointOptionsPtrOutput)
}

// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
func (o DomainEndpointOptionsOutput) EnforceHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DomainEndpointOptions) *bool { return v.EnforceHttps }).(pulumi.BoolPtrOutput)
}

// The minimum required TLS version. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_DomainEndpointOptions.html) for valid values.
func (o DomainEndpointOptionsOutput) TlsSecurityPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainEndpointOptions) *string { return v.TlsSecurityPolicy }).(pulumi.StringPtrOutput)
}

type DomainEndpointOptionsPtrOutput struct{ *pulumi.OutputState }

func (DomainEndpointOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainEndpointOptions)(nil)).Elem()
}

func (o DomainEndpointOptionsPtrOutput) ToDomainEndpointOptionsPtrOutput() DomainEndpointOptionsPtrOutput {
	return o
}

func (o DomainEndpointOptionsPtrOutput) ToDomainEndpointOptionsPtrOutputWithContext(ctx context.Context) DomainEndpointOptionsPtrOutput {
	return o
}

func (o DomainEndpointOptionsPtrOutput) Elem() DomainEndpointOptionsOutput {
	return o.ApplyT(func(v *DomainEndpointOptions) DomainEndpointOptions {
		if v != nil {
			return *v
		}
		var ret DomainEndpointOptions
		return ret
	}).(DomainEndpointOptionsOutput)
}

// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
func (o DomainEndpointOptionsPtrOutput) EnforceHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DomainEndpointOptions) *bool {
		if v == nil {
			return nil
		}
		return v.EnforceHttps
	}).(pulumi.BoolPtrOutput)
}

// The minimum required TLS version. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_DomainEndpointOptions.html) for valid values.
func (o DomainEndpointOptionsPtrOutput) TlsSecurityPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainEndpointOptions) *string {
		if v == nil {
			return nil
		}
		return v.TlsSecurityPolicy
	}).(pulumi.StringPtrOutput)
}

type DomainIndexField struct {
	// The analysis scheme you want to use for a `text` field. The analysis scheme specifies the language-specific text processing options that are used during indexing.
	AnalysisScheme *string `pulumi:"analysisScheme"`
	// The default value for the field. This value is used when no value is specified for the field in the document data.
	DefaultValue *string `pulumi:"defaultValue"`
	// You can get facet information by enabling this.
	Facet *bool `pulumi:"facet"`
	// You can highlight information.
	Highlight *bool `pulumi:"highlight"`
	// A unique name for the field. Field names must begin with a letter and be at least 3 and no more than 64 characters long. The allowed characters are: `a`-`z` (lower-case letters), `0`-`9`, and `_` (underscore). The name `score` is reserved and cannot be used as a field name.
	Name string `pulumi:"name"`
	// You can enable returning the value of all searchable fields.
	Return *bool `pulumi:"return"`
	// You can set whether this index should be searchable or not.
	Search *bool `pulumi:"search"`
	// You can enable the property to be sortable.
	Sort *bool `pulumi:"sort"`
	// A comma-separated list of source fields to map to the field. Specifying a source field copies data from one field to another, enabling you to use the same source data in different ways by configuring different options for the fields.
	SourceFields *string `pulumi:"sourceFields"`
	// The field type. Valid values: `date`, `date-array`, `double`, `double-array`, `int`, `int-array`, `literal`, `literal-array`, `text`, `text-array`.
	Type string `pulumi:"type"`
}

// DomainIndexFieldInput is an input type that accepts DomainIndexFieldArgs and DomainIndexFieldOutput values.
// You can construct a concrete instance of `DomainIndexFieldInput` via:
//
//	DomainIndexFieldArgs{...}
type DomainIndexFieldInput interface {
	pulumi.Input

	ToDomainIndexFieldOutput() DomainIndexFieldOutput
	ToDomainIndexFieldOutputWithContext(context.Context) DomainIndexFieldOutput
}

type DomainIndexFieldArgs struct {
	// The analysis scheme you want to use for a `text` field. The analysis scheme specifies the language-specific text processing options that are used during indexing.
	AnalysisScheme pulumi.StringPtrInput `pulumi:"analysisScheme"`
	// The default value for the field. This value is used when no value is specified for the field in the document data.
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	// You can get facet information by enabling this.
	Facet pulumi.BoolPtrInput `pulumi:"facet"`
	// You can highlight information.
	Highlight pulumi.BoolPtrInput `pulumi:"highlight"`
	// A unique name for the field. Field names must begin with a letter and be at least 3 and no more than 64 characters long. The allowed characters are: `a`-`z` (lower-case letters), `0`-`9`, and `_` (underscore). The name `score` is reserved and cannot be used as a field name.
	Name pulumi.StringInput `pulumi:"name"`
	// You can enable returning the value of all searchable fields.
	Return pulumi.BoolPtrInput `pulumi:"return"`
	// You can set whether this index should be searchable or not.
	Search pulumi.BoolPtrInput `pulumi:"search"`
	// You can enable the property to be sortable.
	Sort pulumi.BoolPtrInput `pulumi:"sort"`
	// A comma-separated list of source fields to map to the field. Specifying a source field copies data from one field to another, enabling you to use the same source data in different ways by configuring different options for the fields.
	SourceFields pulumi.StringPtrInput `pulumi:"sourceFields"`
	// The field type. Valid values: `date`, `date-array`, `double`, `double-array`, `int`, `int-array`, `literal`, `literal-array`, `text`, `text-array`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (DomainIndexFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainIndexField)(nil)).Elem()
}

func (i DomainIndexFieldArgs) ToDomainIndexFieldOutput() DomainIndexFieldOutput {
	return i.ToDomainIndexFieldOutputWithContext(context.Background())
}

func (i DomainIndexFieldArgs) ToDomainIndexFieldOutputWithContext(ctx context.Context) DomainIndexFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIndexFieldOutput)
}

// DomainIndexFieldArrayInput is an input type that accepts DomainIndexFieldArray and DomainIndexFieldArrayOutput values.
// You can construct a concrete instance of `DomainIndexFieldArrayInput` via:
//
//	DomainIndexFieldArray{ DomainIndexFieldArgs{...} }
type DomainIndexFieldArrayInput interface {
	pulumi.Input

	ToDomainIndexFieldArrayOutput() DomainIndexFieldArrayOutput
	ToDomainIndexFieldArrayOutputWithContext(context.Context) DomainIndexFieldArrayOutput
}

type DomainIndexFieldArray []DomainIndexFieldInput

func (DomainIndexFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainIndexField)(nil)).Elem()
}

func (i DomainIndexFieldArray) ToDomainIndexFieldArrayOutput() DomainIndexFieldArrayOutput {
	return i.ToDomainIndexFieldArrayOutputWithContext(context.Background())
}

func (i DomainIndexFieldArray) ToDomainIndexFieldArrayOutputWithContext(ctx context.Context) DomainIndexFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainIndexFieldArrayOutput)
}

type DomainIndexFieldOutput struct{ *pulumi.OutputState }

func (DomainIndexFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainIndexField)(nil)).Elem()
}

func (o DomainIndexFieldOutput) ToDomainIndexFieldOutput() DomainIndexFieldOutput {
	return o
}

func (o DomainIndexFieldOutput) ToDomainIndexFieldOutputWithContext(ctx context.Context) DomainIndexFieldOutput {
	return o
}

// The analysis scheme you want to use for a `text` field. The analysis scheme specifies the language-specific text processing options that are used during indexing.
func (o DomainIndexFieldOutput) AnalysisScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *string { return v.AnalysisScheme }).(pulumi.StringPtrOutput)
}

// The default value for the field. This value is used when no value is specified for the field in the document data.
func (o DomainIndexFieldOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

// You can get facet information by enabling this.
func (o DomainIndexFieldOutput) Facet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *bool { return v.Facet }).(pulumi.BoolPtrOutput)
}

// You can highlight information.
func (o DomainIndexFieldOutput) Highlight() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *bool { return v.Highlight }).(pulumi.BoolPtrOutput)
}

// A unique name for the field. Field names must begin with a letter and be at least 3 and no more than 64 characters long. The allowed characters are: `a`-`z` (lower-case letters), `0`-`9`, and `_` (underscore). The name `score` is reserved and cannot be used as a field name.
func (o DomainIndexFieldOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DomainIndexField) string { return v.Name }).(pulumi.StringOutput)
}

// You can enable returning the value of all searchable fields.
func (o DomainIndexFieldOutput) Return() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *bool { return v.Return }).(pulumi.BoolPtrOutput)
}

// You can set whether this index should be searchable or not.
func (o DomainIndexFieldOutput) Search() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *bool { return v.Search }).(pulumi.BoolPtrOutput)
}

// You can enable the property to be sortable.
func (o DomainIndexFieldOutput) Sort() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *bool { return v.Sort }).(pulumi.BoolPtrOutput)
}

// A comma-separated list of source fields to map to the field. Specifying a source field copies data from one field to another, enabling you to use the same source data in different ways by configuring different options for the fields.
func (o DomainIndexFieldOutput) SourceFields() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainIndexField) *string { return v.SourceFields }).(pulumi.StringPtrOutput)
}

// The field type. Valid values: `date`, `date-array`, `double`, `double-array`, `int`, `int-array`, `literal`, `literal-array`, `text`, `text-array`.
func (o DomainIndexFieldOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DomainIndexField) string { return v.Type }).(pulumi.StringOutput)
}

type DomainIndexFieldArrayOutput struct{ *pulumi.OutputState }

func (DomainIndexFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainIndexField)(nil)).Elem()
}

func (o DomainIndexFieldArrayOutput) ToDomainIndexFieldArrayOutput() DomainIndexFieldArrayOutput {
	return o
}

func (o DomainIndexFieldArrayOutput) ToDomainIndexFieldArrayOutputWithContext(ctx context.Context) DomainIndexFieldArrayOutput {
	return o
}

func (o DomainIndexFieldArrayOutput) Index(i pulumi.IntInput) DomainIndexFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainIndexField {
		return vs[0].([]DomainIndexField)[vs[1].(int)]
	}).(DomainIndexFieldOutput)
}

type DomainScalingParameters struct {
	// The instance type that you want to preconfigure for your domain. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_ScalingParameters.html) for valid values.
	DesiredInstanceType *string `pulumi:"desiredInstanceType"`
	// The number of partitions you want to preconfigure for your domain. Only valid when you select `search.2xlarge` as the instance type.
	DesiredPartitionCount *int `pulumi:"desiredPartitionCount"`
	// The number of replicas you want to preconfigure for each index partition.
	DesiredReplicationCount *int `pulumi:"desiredReplicationCount"`
}

// DomainScalingParametersInput is an input type that accepts DomainScalingParametersArgs and DomainScalingParametersOutput values.
// You can construct a concrete instance of `DomainScalingParametersInput` via:
//
//	DomainScalingParametersArgs{...}
type DomainScalingParametersInput interface {
	pulumi.Input

	ToDomainScalingParametersOutput() DomainScalingParametersOutput
	ToDomainScalingParametersOutputWithContext(context.Context) DomainScalingParametersOutput
}

type DomainScalingParametersArgs struct {
	// The instance type that you want to preconfigure for your domain. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_ScalingParameters.html) for valid values.
	DesiredInstanceType pulumi.StringPtrInput `pulumi:"desiredInstanceType"`
	// The number of partitions you want to preconfigure for your domain. Only valid when you select `search.2xlarge` as the instance type.
	DesiredPartitionCount pulumi.IntPtrInput `pulumi:"desiredPartitionCount"`
	// The number of replicas you want to preconfigure for each index partition.
	DesiredReplicationCount pulumi.IntPtrInput `pulumi:"desiredReplicationCount"`
}

func (DomainScalingParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainScalingParameters)(nil)).Elem()
}

func (i DomainScalingParametersArgs) ToDomainScalingParametersOutput() DomainScalingParametersOutput {
	return i.ToDomainScalingParametersOutputWithContext(context.Background())
}

func (i DomainScalingParametersArgs) ToDomainScalingParametersOutputWithContext(ctx context.Context) DomainScalingParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainScalingParametersOutput)
}

func (i DomainScalingParametersArgs) ToDomainScalingParametersPtrOutput() DomainScalingParametersPtrOutput {
	return i.ToDomainScalingParametersPtrOutputWithContext(context.Background())
}

func (i DomainScalingParametersArgs) ToDomainScalingParametersPtrOutputWithContext(ctx context.Context) DomainScalingParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainScalingParametersOutput).ToDomainScalingParametersPtrOutputWithContext(ctx)
}

// DomainScalingParametersPtrInput is an input type that accepts DomainScalingParametersArgs, DomainScalingParametersPtr and DomainScalingParametersPtrOutput values.
// You can construct a concrete instance of `DomainScalingParametersPtrInput` via:
//
//	        DomainScalingParametersArgs{...}
//
//	or:
//
//	        nil
type DomainScalingParametersPtrInput interface {
	pulumi.Input

	ToDomainScalingParametersPtrOutput() DomainScalingParametersPtrOutput
	ToDomainScalingParametersPtrOutputWithContext(context.Context) DomainScalingParametersPtrOutput
}

type domainScalingParametersPtrType DomainScalingParametersArgs

func DomainScalingParametersPtr(v *DomainScalingParametersArgs) DomainScalingParametersPtrInput {
	return (*domainScalingParametersPtrType)(v)
}

func (*domainScalingParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainScalingParameters)(nil)).Elem()
}

func (i *domainScalingParametersPtrType) ToDomainScalingParametersPtrOutput() DomainScalingParametersPtrOutput {
	return i.ToDomainScalingParametersPtrOutputWithContext(context.Background())
}

func (i *domainScalingParametersPtrType) ToDomainScalingParametersPtrOutputWithContext(ctx context.Context) DomainScalingParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainScalingParametersPtrOutput)
}

type DomainScalingParametersOutput struct{ *pulumi.OutputState }

func (DomainScalingParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainScalingParameters)(nil)).Elem()
}

func (o DomainScalingParametersOutput) ToDomainScalingParametersOutput() DomainScalingParametersOutput {
	return o
}

func (o DomainScalingParametersOutput) ToDomainScalingParametersOutputWithContext(ctx context.Context) DomainScalingParametersOutput {
	return o
}

func (o DomainScalingParametersOutput) ToDomainScalingParametersPtrOutput() DomainScalingParametersPtrOutput {
	return o.ToDomainScalingParametersPtrOutputWithContext(context.Background())
}

func (o DomainScalingParametersOutput) ToDomainScalingParametersPtrOutputWithContext(ctx context.Context) DomainScalingParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainScalingParameters) *DomainScalingParameters {
		return &v
	}).(DomainScalingParametersPtrOutput)
}

// The instance type that you want to preconfigure for your domain. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_ScalingParameters.html) for valid values.
func (o DomainScalingParametersOutput) DesiredInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainScalingParameters) *string { return v.DesiredInstanceType }).(pulumi.StringPtrOutput)
}

// The number of partitions you want to preconfigure for your domain. Only valid when you select `search.2xlarge` as the instance type.
func (o DomainScalingParametersOutput) DesiredPartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainScalingParameters) *int { return v.DesiredPartitionCount }).(pulumi.IntPtrOutput)
}

// The number of replicas you want to preconfigure for each index partition.
func (o DomainScalingParametersOutput) DesiredReplicationCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainScalingParameters) *int { return v.DesiredReplicationCount }).(pulumi.IntPtrOutput)
}

type DomainScalingParametersPtrOutput struct{ *pulumi.OutputState }

func (DomainScalingParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainScalingParameters)(nil)).Elem()
}

func (o DomainScalingParametersPtrOutput) ToDomainScalingParametersPtrOutput() DomainScalingParametersPtrOutput {
	return o
}

func (o DomainScalingParametersPtrOutput) ToDomainScalingParametersPtrOutputWithContext(ctx context.Context) DomainScalingParametersPtrOutput {
	return o
}

func (o DomainScalingParametersPtrOutput) Elem() DomainScalingParametersOutput {
	return o.ApplyT(func(v *DomainScalingParameters) DomainScalingParameters {
		if v != nil {
			return *v
		}
		var ret DomainScalingParameters
		return ret
	}).(DomainScalingParametersOutput)
}

// The instance type that you want to preconfigure for your domain. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_ScalingParameters.html) for valid values.
func (o DomainScalingParametersPtrOutput) DesiredInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainScalingParameters) *string {
		if v == nil {
			return nil
		}
		return v.DesiredInstanceType
	}).(pulumi.StringPtrOutput)
}

// The number of partitions you want to preconfigure for your domain. Only valid when you select `search.2xlarge` as the instance type.
func (o DomainScalingParametersPtrOutput) DesiredPartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainScalingParameters) *int {
		if v == nil {
			return nil
		}
		return v.DesiredPartitionCount
	}).(pulumi.IntPtrOutput)
}

// The number of replicas you want to preconfigure for each index partition.
func (o DomainScalingParametersPtrOutput) DesiredReplicationCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainScalingParameters) *int {
		if v == nil {
			return nil
		}
		return v.DesiredReplicationCount
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainEndpointOptionsInput)(nil)).Elem(), DomainEndpointOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainEndpointOptionsPtrInput)(nil)).Elem(), DomainEndpointOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIndexFieldInput)(nil)).Elem(), DomainIndexFieldArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainIndexFieldArrayInput)(nil)).Elem(), DomainIndexFieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainScalingParametersInput)(nil)).Elem(), DomainScalingParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainScalingParametersPtrInput)(nil)).Elem(), DomainScalingParametersArgs{})
	pulumi.RegisterOutputType(DomainEndpointOptionsOutput{})
	pulumi.RegisterOutputType(DomainEndpointOptionsPtrOutput{})
	pulumi.RegisterOutputType(DomainIndexFieldOutput{})
	pulumi.RegisterOutputType(DomainIndexFieldArrayOutput{})
	pulumi.RegisterOutputType(DomainScalingParametersOutput{})
	pulumi.RegisterOutputType(DomainScalingParametersPtrOutput{})
}
