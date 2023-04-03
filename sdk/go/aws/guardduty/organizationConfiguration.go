// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages the GuardDuty Organization Configuration in the current AWS Region. The AWS account utilizing this resource must have been assigned as a delegated Organization administrator account, e.g., via the `guardduty.OrganizationAdminAccount` resource. More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).
//
// > **NOTE:** This is an advanced resource. The provider will automatically assume management of the GuardDuty Organization Configuration without import and perform no actions on removal from the resource configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/guardduty"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDetector, err := guardduty.NewDetector(ctx, "exampleDetector", &guardduty.DetectorArgs{
//				Enable: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = guardduty.NewOrganizationConfiguration(ctx, "exampleOrganizationConfiguration", &guardduty.OrganizationConfigurationArgs{
//				AutoEnable: pulumi.Bool(true),
//				DetectorId: exampleDetector.ID(),
//				Datasources: &guardduty.OrganizationConfigurationDatasourcesArgs{
//					S3Logs: &guardduty.OrganizationConfigurationDatasourcesS3LogsArgs{
//						AutoEnable: pulumi.Bool(true),
//					},
//					Kubernetes: &guardduty.OrganizationConfigurationDatasourcesKubernetesArgs{
//						AuditLogs: &guardduty.OrganizationConfigurationDatasourcesKubernetesAuditLogsArgs{
//							Enable: pulumi.Bool(true),
//						},
//					},
//					MalwareProtection: &guardduty.OrganizationConfigurationDatasourcesMalwareProtectionArgs{
//						ScanEc2InstanceWithFindings: &guardduty.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsArgs{
//							EbsVolumes: &guardduty.OrganizationConfigurationDatasourcesMalwareProtectionScanEc2InstanceWithFindingsEbsVolumesArgs{
//								AutoEnable: pulumi.Bool(true),
//							},
//						},
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
// GuardDuty Organization Configurations can be imported using the GuardDuty Detector ID, e.g.,
//
// ```sh
//
//	$ pulumi import aws:guardduty/organizationConfiguration:OrganizationConfiguration example 00b00fd5aecc0ab60a708659477e9617
//
// ```
type OrganizationConfiguration struct {
	pulumi.CustomResourceState

	// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
	AutoEnable pulumi.BoolOutput `pulumi:"autoEnable"`
	// Configuration for the collected datasources.
	Datasources OrganizationConfigurationDatasourcesOutput `pulumi:"datasources"`
	// The detector ID of the GuardDuty account.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
}

// NewOrganizationConfiguration registers a new resource with the given unique name, arguments, and options.
func NewOrganizationConfiguration(ctx *pulumi.Context,
	name string, args *OrganizationConfigurationArgs, opts ...pulumi.ResourceOption) (*OrganizationConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoEnable == nil {
		return nil, errors.New("invalid value for required argument 'AutoEnable'")
	}
	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	var resource OrganizationConfiguration
	err := ctx.RegisterResource("aws:guardduty/organizationConfiguration:OrganizationConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationConfiguration gets an existing OrganizationConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationConfigurationState, opts ...pulumi.ResourceOption) (*OrganizationConfiguration, error) {
	var resource OrganizationConfiguration
	err := ctx.ReadResource("aws:guardduty/organizationConfiguration:OrganizationConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationConfiguration resources.
type organizationConfigurationState struct {
	// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
	AutoEnable *bool `pulumi:"autoEnable"`
	// Configuration for the collected datasources.
	Datasources *OrganizationConfigurationDatasources `pulumi:"datasources"`
	// The detector ID of the GuardDuty account.
	DetectorId *string `pulumi:"detectorId"`
}

type OrganizationConfigurationState struct {
	// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
	AutoEnable pulumi.BoolPtrInput
	// Configuration for the collected datasources.
	Datasources OrganizationConfigurationDatasourcesPtrInput
	// The detector ID of the GuardDuty account.
	DetectorId pulumi.StringPtrInput
}

func (OrganizationConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationConfigurationState)(nil)).Elem()
}

type organizationConfigurationArgs struct {
	// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
	AutoEnable bool `pulumi:"autoEnable"`
	// Configuration for the collected datasources.
	Datasources *OrganizationConfigurationDatasources `pulumi:"datasources"`
	// The detector ID of the GuardDuty account.
	DetectorId string `pulumi:"detectorId"`
}

// The set of arguments for constructing a OrganizationConfiguration resource.
type OrganizationConfigurationArgs struct {
	// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
	AutoEnable pulumi.BoolInput
	// Configuration for the collected datasources.
	Datasources OrganizationConfigurationDatasourcesPtrInput
	// The detector ID of the GuardDuty account.
	DetectorId pulumi.StringInput
}

func (OrganizationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationConfigurationArgs)(nil)).Elem()
}

type OrganizationConfigurationInput interface {
	pulumi.Input

	ToOrganizationConfigurationOutput() OrganizationConfigurationOutput
	ToOrganizationConfigurationOutputWithContext(ctx context.Context) OrganizationConfigurationOutput
}

func (*OrganizationConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationConfiguration)(nil)).Elem()
}

func (i *OrganizationConfiguration) ToOrganizationConfigurationOutput() OrganizationConfigurationOutput {
	return i.ToOrganizationConfigurationOutputWithContext(context.Background())
}

func (i *OrganizationConfiguration) ToOrganizationConfigurationOutputWithContext(ctx context.Context) OrganizationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConfigurationOutput)
}

// OrganizationConfigurationArrayInput is an input type that accepts OrganizationConfigurationArray and OrganizationConfigurationArrayOutput values.
// You can construct a concrete instance of `OrganizationConfigurationArrayInput` via:
//
//	OrganizationConfigurationArray{ OrganizationConfigurationArgs{...} }
type OrganizationConfigurationArrayInput interface {
	pulumi.Input

	ToOrganizationConfigurationArrayOutput() OrganizationConfigurationArrayOutput
	ToOrganizationConfigurationArrayOutputWithContext(context.Context) OrganizationConfigurationArrayOutput
}

type OrganizationConfigurationArray []OrganizationConfigurationInput

func (OrganizationConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationConfiguration)(nil)).Elem()
}

func (i OrganizationConfigurationArray) ToOrganizationConfigurationArrayOutput() OrganizationConfigurationArrayOutput {
	return i.ToOrganizationConfigurationArrayOutputWithContext(context.Background())
}

func (i OrganizationConfigurationArray) ToOrganizationConfigurationArrayOutputWithContext(ctx context.Context) OrganizationConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConfigurationArrayOutput)
}

// OrganizationConfigurationMapInput is an input type that accepts OrganizationConfigurationMap and OrganizationConfigurationMapOutput values.
// You can construct a concrete instance of `OrganizationConfigurationMapInput` via:
//
//	OrganizationConfigurationMap{ "key": OrganizationConfigurationArgs{...} }
type OrganizationConfigurationMapInput interface {
	pulumi.Input

	ToOrganizationConfigurationMapOutput() OrganizationConfigurationMapOutput
	ToOrganizationConfigurationMapOutputWithContext(context.Context) OrganizationConfigurationMapOutput
}

type OrganizationConfigurationMap map[string]OrganizationConfigurationInput

func (OrganizationConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationConfiguration)(nil)).Elem()
}

func (i OrganizationConfigurationMap) ToOrganizationConfigurationMapOutput() OrganizationConfigurationMapOutput {
	return i.ToOrganizationConfigurationMapOutputWithContext(context.Background())
}

func (i OrganizationConfigurationMap) ToOrganizationConfigurationMapOutputWithContext(ctx context.Context) OrganizationConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConfigurationMapOutput)
}

type OrganizationConfigurationOutput struct{ *pulumi.OutputState }

func (OrganizationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationConfiguration)(nil)).Elem()
}

func (o OrganizationConfigurationOutput) ToOrganizationConfigurationOutput() OrganizationConfigurationOutput {
	return o
}

func (o OrganizationConfigurationOutput) ToOrganizationConfigurationOutputWithContext(ctx context.Context) OrganizationConfigurationOutput {
	return o
}

// When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
func (o OrganizationConfigurationOutput) AutoEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrganizationConfiguration) pulumi.BoolOutput { return v.AutoEnable }).(pulumi.BoolOutput)
}

// Configuration for the collected datasources.
func (o OrganizationConfigurationOutput) Datasources() OrganizationConfigurationDatasourcesOutput {
	return o.ApplyT(func(v *OrganizationConfiguration) OrganizationConfigurationDatasourcesOutput { return v.Datasources }).(OrganizationConfigurationDatasourcesOutput)
}

// The detector ID of the GuardDuty account.
func (o OrganizationConfigurationOutput) DetectorId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationConfiguration) pulumi.StringOutput { return v.DetectorId }).(pulumi.StringOutput)
}

type OrganizationConfigurationArrayOutput struct{ *pulumi.OutputState }

func (OrganizationConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationConfiguration)(nil)).Elem()
}

func (o OrganizationConfigurationArrayOutput) ToOrganizationConfigurationArrayOutput() OrganizationConfigurationArrayOutput {
	return o
}

func (o OrganizationConfigurationArrayOutput) ToOrganizationConfigurationArrayOutputWithContext(ctx context.Context) OrganizationConfigurationArrayOutput {
	return o
}

func (o OrganizationConfigurationArrayOutput) Index(i pulumi.IntInput) OrganizationConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationConfiguration {
		return vs[0].([]*OrganizationConfiguration)[vs[1].(int)]
	}).(OrganizationConfigurationOutput)
}

type OrganizationConfigurationMapOutput struct{ *pulumi.OutputState }

func (OrganizationConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationConfiguration)(nil)).Elem()
}

func (o OrganizationConfigurationMapOutput) ToOrganizationConfigurationMapOutput() OrganizationConfigurationMapOutput {
	return o
}

func (o OrganizationConfigurationMapOutput) ToOrganizationConfigurationMapOutputWithContext(ctx context.Context) OrganizationConfigurationMapOutput {
	return o
}

func (o OrganizationConfigurationMapOutput) MapIndex(k pulumi.StringInput) OrganizationConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationConfiguration {
		return vs[0].(map[string]*OrganizationConfiguration)[vs[1].(string)]
	}).(OrganizationConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationConfigurationInput)(nil)).Elem(), &OrganizationConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationConfigurationArrayInput)(nil)).Elem(), OrganizationConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationConfigurationMapInput)(nil)).Elem(), OrganizationConfigurationMap{})
	pulumi.RegisterOutputType(OrganizationConfigurationOutput{})
	pulumi.RegisterOutputType(OrganizationConfigurationArrayOutput{})
	pulumi.RegisterOutputType(OrganizationConfigurationMapOutput{})
}
