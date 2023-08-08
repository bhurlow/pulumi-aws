// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provisions and manages a Service Catalog provisioned product.
//
// A provisioned product is a resourced instance of a product. For example, provisioning a product based on a CloudFormation template launches a CloudFormation stack and its underlying resources.
//
// Like this resource, the `awsServicecatalogRecord` data source also provides information about a provisioned product. Although a Service Catalog record provides some overlapping information with this resource, a record is tied to a provisioned product event, such as provisioning, termination, and updating.
//
// > **Tip:** If you include conflicted keys as tags, AWS will report an error, "Parameter validation failed: Missing required parameter in Tags[N]:Value".
//
// > **Tip:** A "provisioning artifact" is also referred to as a "version." A "distributor" is also referred to as a "vendor."
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicecatalog.NewProvisionedProduct(ctx, "example", &servicecatalog.ProvisionedProductArgs{
//				ProductName:              pulumi.String("Example product"),
//				ProvisioningArtifactName: pulumi.String("Example version"),
//				ProvisioningParameters: servicecatalog.ProvisionedProductProvisioningParameterArray{
//					&servicecatalog.ProvisionedProductProvisioningParameterArgs{
//						Key:   pulumi.String("foo"),
//						Value: pulumi.String("bar"),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
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
// terraform import {
//
//	to = aws_servicecatalog_provisioned_product.example
//
//	id = "pp-dnigbtea24ste" } Using `pulumi import`, import `aws_servicecatalog_provisioned_product` using the provisioned product ID. For exampleconsole % pulumi import aws_servicecatalog_provisioned_product.example pp-dnigbtea24ste
type ProvisionedProduct struct {
	pulumi.CustomResourceState

	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	// ARN of the provisioned product.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Set of CloudWatch dashboards that were created when provisioning the product.
	CloudwatchDashboardNames pulumi.StringArrayOutput `pulumi:"cloudwatchDashboardNames"`
	// Time when the provisioned product was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
	IgnoreErrors pulumi.BoolPtrOutput `pulumi:"ignoreErrors"`
	// Record identifier of the last request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
	LastProvisioningRecordId pulumi.StringOutput `pulumi:"lastProvisioningRecordId"`
	// Record identifier of the last request performed on this provisioned product.
	LastRecordId pulumi.StringOutput `pulumi:"lastRecordId"`
	// Record identifier of the last successful request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
	LastSuccessfulProvisioningRecordId pulumi.StringOutput `pulumi:"lastSuccessfulProvisioningRecordId"`
	// ARN of the launch role associated with the provisioned product.
	LaunchRoleArn pulumi.StringOutput `pulumi:"launchRoleArn"`
	// User-friendly name of the provisioned product.
	//
	// The following arguments are optional:
	Name pulumi.StringOutput `pulumi:"name"`
	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
	NotificationArns pulumi.StringArrayOutput `pulumi:"notificationArns"`
	// The set of outputs for the product created.
	Outputs ProvisionedProductOutputTypeArrayOutput `pulumi:"outputs"`
	// Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
	PathId pulumi.StringOutput `pulumi:"pathId"`
	// Name of the path. You must provide `pathId` or `pathName`, but not both.
	PathName pulumi.StringPtrOutput `pulumi:"pathName"`
	// Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// Name of the product. You must provide `productId` or `productName`, but not both.
	ProductName pulumi.StringPtrOutput `pulumi:"productName"`
	// Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactId pulumi.StringOutput `pulumi:"provisioningArtifactId"`
	// Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactName pulumi.StringPtrOutput `pulumi:"provisioningArtifactName"`
	// Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
	ProvisioningParameters ProvisionedProductProvisioningParameterArrayOutput `pulumi:"provisioningParameters"`
	// _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
	RetainPhysicalResources pulumi.BoolPtrOutput `pulumi:"retainPhysicalResources"`
	// Configuration block with information about the provisioning preferences for a stack set. See details below.
	StackSetProvisioningPreferences ProvisionedProductStackSetProvisioningPreferencesPtrOutput `pulumi:"stackSetProvisioningPreferences"`
	// Current status of the provisioned product. See meanings below.
	Status pulumi.StringOutput `pulumi:"status"`
	// Current status message of the provisioned product.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Type of provisioned product. Valid values are `CFN_STACK` and `CFN_STACKSET`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProvisionedProduct registers a new resource with the given unique name, arguments, and options.
func NewProvisionedProduct(ctx *pulumi.Context,
	name string, args *ProvisionedProductArgs, opts ...pulumi.ResourceOption) (*ProvisionedProduct, error) {
	if args == nil {
		args = &ProvisionedProductArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProvisionedProduct
	err := ctx.RegisterResource("aws:servicecatalog/provisionedProduct:ProvisionedProduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProvisionedProduct gets an existing ProvisionedProduct resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvisionedProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisionedProductState, opts ...pulumi.ResourceOption) (*ProvisionedProduct, error) {
	var resource ProvisionedProduct
	err := ctx.ReadResource("aws:servicecatalog/provisionedProduct:ProvisionedProduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProvisionedProduct resources.
type provisionedProductState struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// ARN of the provisioned product.
	Arn *string `pulumi:"arn"`
	// Set of CloudWatch dashboards that were created when provisioning the product.
	CloudwatchDashboardNames []string `pulumi:"cloudwatchDashboardNames"`
	// Time when the provisioned product was created.
	CreatedTime *string `pulumi:"createdTime"`
	// _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
	IgnoreErrors *bool `pulumi:"ignoreErrors"`
	// Record identifier of the last request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
	LastProvisioningRecordId *string `pulumi:"lastProvisioningRecordId"`
	// Record identifier of the last request performed on this provisioned product.
	LastRecordId *string `pulumi:"lastRecordId"`
	// Record identifier of the last successful request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
	LastSuccessfulProvisioningRecordId *string `pulumi:"lastSuccessfulProvisioningRecordId"`
	// ARN of the launch role associated with the provisioned product.
	LaunchRoleArn *string `pulumi:"launchRoleArn"`
	// User-friendly name of the provisioned product.
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
	NotificationArns []string `pulumi:"notificationArns"`
	// The set of outputs for the product created.
	Outputs []ProvisionedProductOutputType `pulumi:"outputs"`
	// Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
	PathId *string `pulumi:"pathId"`
	// Name of the path. You must provide `pathId` or `pathName`, but not both.
	PathName *string `pulumi:"pathName"`
	// Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
	ProductId *string `pulumi:"productId"`
	// Name of the product. You must provide `productId` or `productName`, but not both.
	ProductName *string `pulumi:"productName"`
	// Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactId *string `pulumi:"provisioningArtifactId"`
	// Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactName *string `pulumi:"provisioningArtifactName"`
	// Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
	ProvisioningParameters []ProvisionedProductProvisioningParameter `pulumi:"provisioningParameters"`
	// _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
	RetainPhysicalResources *bool `pulumi:"retainPhysicalResources"`
	// Configuration block with information about the provisioning preferences for a stack set. See details below.
	StackSetProvisioningPreferences *ProvisionedProductStackSetProvisioningPreferences `pulumi:"stackSetProvisioningPreferences"`
	// Current status of the provisioned product. See meanings below.
	Status *string `pulumi:"status"`
	// Current status message of the provisioned product.
	StatusMessage *string `pulumi:"statusMessage"`
	// Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Type of provisioned product. Valid values are `CFN_STACK` and `CFN_STACKSET`.
	Type *string `pulumi:"type"`
}

type ProvisionedProductState struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrInput
	// ARN of the provisioned product.
	Arn pulumi.StringPtrInput
	// Set of CloudWatch dashboards that were created when provisioning the product.
	CloudwatchDashboardNames pulumi.StringArrayInput
	// Time when the provisioned product was created.
	CreatedTime pulumi.StringPtrInput
	// _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
	IgnoreErrors pulumi.BoolPtrInput
	// Record identifier of the last request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
	LastProvisioningRecordId pulumi.StringPtrInput
	// Record identifier of the last request performed on this provisioned product.
	LastRecordId pulumi.StringPtrInput
	// Record identifier of the last successful request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
	LastSuccessfulProvisioningRecordId pulumi.StringPtrInput
	// ARN of the launch role associated with the provisioned product.
	LaunchRoleArn pulumi.StringPtrInput
	// User-friendly name of the provisioned product.
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
	NotificationArns pulumi.StringArrayInput
	// The set of outputs for the product created.
	Outputs ProvisionedProductOutputTypeArrayInput
	// Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
	PathId pulumi.StringPtrInput
	// Name of the path. You must provide `pathId` or `pathName`, but not both.
	PathName pulumi.StringPtrInput
	// Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
	ProductId pulumi.StringPtrInput
	// Name of the product. You must provide `productId` or `productName`, but not both.
	ProductName pulumi.StringPtrInput
	// Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactId pulumi.StringPtrInput
	// Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactName pulumi.StringPtrInput
	// Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
	ProvisioningParameters ProvisionedProductProvisioningParameterArrayInput
	// _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
	RetainPhysicalResources pulumi.BoolPtrInput
	// Configuration block with information about the provisioning preferences for a stack set. See details below.
	StackSetProvisioningPreferences ProvisionedProductStackSetProvisioningPreferencesPtrInput
	// Current status of the provisioned product. See meanings below.
	Status pulumi.StringPtrInput
	// Current status message of the provisioned product.
	StatusMessage pulumi.StringPtrInput
	// Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Type of provisioned product. Valid values are `CFN_STACK` and `CFN_STACKSET`.
	Type pulumi.StringPtrInput
}

func (ProvisionedProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedProductState)(nil)).Elem()
}

type provisionedProductArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
	IgnoreErrors *bool `pulumi:"ignoreErrors"`
	// User-friendly name of the provisioned product.
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
	NotificationArns []string `pulumi:"notificationArns"`
	// Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
	PathId *string `pulumi:"pathId"`
	// Name of the path. You must provide `pathId` or `pathName`, but not both.
	PathName *string `pulumi:"pathName"`
	// Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
	ProductId *string `pulumi:"productId"`
	// Name of the product. You must provide `productId` or `productName`, but not both.
	ProductName *string `pulumi:"productName"`
	// Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactId *string `pulumi:"provisioningArtifactId"`
	// Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactName *string `pulumi:"provisioningArtifactName"`
	// Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
	ProvisioningParameters []ProvisionedProductProvisioningParameter `pulumi:"provisioningParameters"`
	// _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
	RetainPhysicalResources *bool `pulumi:"retainPhysicalResources"`
	// Configuration block with information about the provisioning preferences for a stack set. See details below.
	StackSetProvisioningPreferences *ProvisionedProductStackSetProvisioningPreferences `pulumi:"stackSetProvisioningPreferences"`
	// Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ProvisionedProduct resource.
type ProvisionedProductArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrInput
	// _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
	IgnoreErrors pulumi.BoolPtrInput
	// User-friendly name of the provisioned product.
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
	NotificationArns pulumi.StringArrayInput
	// Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
	PathId pulumi.StringPtrInput
	// Name of the path. You must provide `pathId` or `pathName`, but not both.
	PathName pulumi.StringPtrInput
	// Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
	ProductId pulumi.StringPtrInput
	// Name of the product. You must provide `productId` or `productName`, but not both.
	ProductName pulumi.StringPtrInput
	// Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactId pulumi.StringPtrInput
	// Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
	ProvisioningArtifactName pulumi.StringPtrInput
	// Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
	ProvisioningParameters ProvisionedProductProvisioningParameterArrayInput
	// _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
	RetainPhysicalResources pulumi.BoolPtrInput
	// Configuration block with information about the provisioning preferences for a stack set. See details below.
	StackSetProvisioningPreferences ProvisionedProductStackSetProvisioningPreferencesPtrInput
	// Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ProvisionedProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedProductArgs)(nil)).Elem()
}

type ProvisionedProductInput interface {
	pulumi.Input

	ToProvisionedProductOutput() ProvisionedProductOutput
	ToProvisionedProductOutputWithContext(ctx context.Context) ProvisionedProductOutput
}

func (*ProvisionedProduct) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedProduct)(nil)).Elem()
}

func (i *ProvisionedProduct) ToProvisionedProductOutput() ProvisionedProductOutput {
	return i.ToProvisionedProductOutputWithContext(context.Background())
}

func (i *ProvisionedProduct) ToProvisionedProductOutputWithContext(ctx context.Context) ProvisionedProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedProductOutput)
}

// ProvisionedProductArrayInput is an input type that accepts ProvisionedProductArray and ProvisionedProductArrayOutput values.
// You can construct a concrete instance of `ProvisionedProductArrayInput` via:
//
//	ProvisionedProductArray{ ProvisionedProductArgs{...} }
type ProvisionedProductArrayInput interface {
	pulumi.Input

	ToProvisionedProductArrayOutput() ProvisionedProductArrayOutput
	ToProvisionedProductArrayOutputWithContext(context.Context) ProvisionedProductArrayOutput
}

type ProvisionedProductArray []ProvisionedProductInput

func (ProvisionedProductArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProvisionedProduct)(nil)).Elem()
}

func (i ProvisionedProductArray) ToProvisionedProductArrayOutput() ProvisionedProductArrayOutput {
	return i.ToProvisionedProductArrayOutputWithContext(context.Background())
}

func (i ProvisionedProductArray) ToProvisionedProductArrayOutputWithContext(ctx context.Context) ProvisionedProductArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedProductArrayOutput)
}

// ProvisionedProductMapInput is an input type that accepts ProvisionedProductMap and ProvisionedProductMapOutput values.
// You can construct a concrete instance of `ProvisionedProductMapInput` via:
//
//	ProvisionedProductMap{ "key": ProvisionedProductArgs{...} }
type ProvisionedProductMapInput interface {
	pulumi.Input

	ToProvisionedProductMapOutput() ProvisionedProductMapOutput
	ToProvisionedProductMapOutputWithContext(context.Context) ProvisionedProductMapOutput
}

type ProvisionedProductMap map[string]ProvisionedProductInput

func (ProvisionedProductMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProvisionedProduct)(nil)).Elem()
}

func (i ProvisionedProductMap) ToProvisionedProductMapOutput() ProvisionedProductMapOutput {
	return i.ToProvisionedProductMapOutputWithContext(context.Background())
}

func (i ProvisionedProductMap) ToProvisionedProductMapOutputWithContext(ctx context.Context) ProvisionedProductMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionedProductMapOutput)
}

type ProvisionedProductOutput struct{ *pulumi.OutputState }

func (ProvisionedProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionedProduct)(nil)).Elem()
}

func (o ProvisionedProductOutput) ToProvisionedProductOutput() ProvisionedProductOutput {
	return o
}

func (o ProvisionedProductOutput) ToProvisionedProductOutputWithContext(ctx context.Context) ProvisionedProductOutput {
	return o
}

// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
func (o ProvisionedProductOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

// ARN of the provisioned product.
func (o ProvisionedProductOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Set of CloudWatch dashboards that were created when provisioning the product.
func (o ProvisionedProductOutput) CloudwatchDashboardNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringArrayOutput { return v.CloudwatchDashboardNames }).(pulumi.StringArrayOutput)
}

// Time when the provisioned product was created.
func (o ProvisionedProductOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
func (o ProvisionedProductOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.BoolPtrOutput { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

// Record identifier of the last request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
func (o ProvisionedProductOutput) LastProvisioningRecordId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.LastProvisioningRecordId }).(pulumi.StringOutput)
}

// Record identifier of the last request performed on this provisioned product.
func (o ProvisionedProductOutput) LastRecordId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.LastRecordId }).(pulumi.StringOutput)
}

// Record identifier of the last successful request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
func (o ProvisionedProductOutput) LastSuccessfulProvisioningRecordId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.LastSuccessfulProvisioningRecordId }).(pulumi.StringOutput)
}

// ARN of the launch role associated with the provisioned product.
func (o ProvisionedProductOutput) LaunchRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.LaunchRoleArn }).(pulumi.StringOutput)
}

// User-friendly name of the provisioned product.
//
// The following arguments are optional:
func (o ProvisionedProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
func (o ProvisionedProductOutput) NotificationArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringArrayOutput { return v.NotificationArns }).(pulumi.StringArrayOutput)
}

// The set of outputs for the product created.
func (o ProvisionedProductOutput) Outputs() ProvisionedProductOutputTypeArrayOutput {
	return o.ApplyT(func(v *ProvisionedProduct) ProvisionedProductOutputTypeArrayOutput { return v.Outputs }).(ProvisionedProductOutputTypeArrayOutput)
}

// Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
func (o ProvisionedProductOutput) PathId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.PathId }).(pulumi.StringOutput)
}

// Name of the path. You must provide `pathId` or `pathName`, but not both.
func (o ProvisionedProductOutput) PathName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringPtrOutput { return v.PathName }).(pulumi.StringPtrOutput)
}

// Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
func (o ProvisionedProductOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

// Name of the product. You must provide `productId` or `productName`, but not both.
func (o ProvisionedProductOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringPtrOutput { return v.ProductName }).(pulumi.StringPtrOutput)
}

// Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
func (o ProvisionedProductOutput) ProvisioningArtifactId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.ProvisioningArtifactId }).(pulumi.StringOutput)
}

// Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
func (o ProvisionedProductOutput) ProvisioningArtifactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringPtrOutput { return v.ProvisioningArtifactName }).(pulumi.StringPtrOutput)
}

// Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
func (o ProvisionedProductOutput) ProvisioningParameters() ProvisionedProductProvisioningParameterArrayOutput {
	return o.ApplyT(func(v *ProvisionedProduct) ProvisionedProductProvisioningParameterArrayOutput {
		return v.ProvisioningParameters
	}).(ProvisionedProductProvisioningParameterArrayOutput)
}

// _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
func (o ProvisionedProductOutput) RetainPhysicalResources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.BoolPtrOutput { return v.RetainPhysicalResources }).(pulumi.BoolPtrOutput)
}

// Configuration block with information about the provisioning preferences for a stack set. See details below.
func (o ProvisionedProductOutput) StackSetProvisioningPreferences() ProvisionedProductStackSetProvisioningPreferencesPtrOutput {
	return o.ApplyT(func(v *ProvisionedProduct) ProvisionedProductStackSetProvisioningPreferencesPtrOutput {
		return v.StackSetProvisioningPreferences
	}).(ProvisionedProductStackSetProvisioningPreferencesPtrOutput)
}

// Current status of the provisioned product. See meanings below.
func (o ProvisionedProductOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Current status message of the provisioned product.
func (o ProvisionedProductOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ProvisionedProductOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ProvisionedProductOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Type of provisioned product. Valid values are `CFN_STACK` and `CFN_STACKSET`.
func (o ProvisionedProductOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProvisionedProduct) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ProvisionedProductArrayOutput struct{ *pulumi.OutputState }

func (ProvisionedProductArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProvisionedProduct)(nil)).Elem()
}

func (o ProvisionedProductArrayOutput) ToProvisionedProductArrayOutput() ProvisionedProductArrayOutput {
	return o
}

func (o ProvisionedProductArrayOutput) ToProvisionedProductArrayOutputWithContext(ctx context.Context) ProvisionedProductArrayOutput {
	return o
}

func (o ProvisionedProductArrayOutput) Index(i pulumi.IntInput) ProvisionedProductOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProvisionedProduct {
		return vs[0].([]*ProvisionedProduct)[vs[1].(int)]
	}).(ProvisionedProductOutput)
}

type ProvisionedProductMapOutput struct{ *pulumi.OutputState }

func (ProvisionedProductMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProvisionedProduct)(nil)).Elem()
}

func (o ProvisionedProductMapOutput) ToProvisionedProductMapOutput() ProvisionedProductMapOutput {
	return o
}

func (o ProvisionedProductMapOutput) ToProvisionedProductMapOutputWithContext(ctx context.Context) ProvisionedProductMapOutput {
	return o
}

func (o ProvisionedProductMapOutput) MapIndex(k pulumi.StringInput) ProvisionedProductOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProvisionedProduct {
		return vs[0].(map[string]*ProvisionedProduct)[vs[1].(string)]
	}).(ProvisionedProductOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionedProductInput)(nil)).Elem(), &ProvisionedProduct{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionedProductArrayInput)(nil)).Elem(), ProvisionedProductArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionedProductMapInput)(nil)).Elem(), ProvisionedProductMap{})
	pulumi.RegisterOutputType(ProvisionedProductOutput{})
	pulumi.RegisterOutputType(ProvisionedProductArrayOutput{})
	pulumi.RegisterOutputType(ProvisionedProductMapOutput{})
}
