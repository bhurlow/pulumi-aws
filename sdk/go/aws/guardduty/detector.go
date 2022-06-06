// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a GuardDuty detector.
//
// > **NOTE:** Deleting this resource is equivalent to "disabling" GuardDuty for an AWS region, which removes all existing findings. You can set the `enable` attribute to `false` to instead "suspend" monitoring and feedback reporting while keeping existing data. See the [Suspending or Disabling Amazon GuardDuty documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_suspend-disable.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/guardduty"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := guardduty.NewDetector(ctx, "myDetector", &guardduty.DetectorArgs{
// 			Datasources: &guardduty.DetectorDatasourcesArgs{
// 				Kubernetes: &guardduty.DetectorDatasourcesKubernetesArgs{
// 					AuditLogs: &guardduty.DetectorDatasourcesKubernetesAuditLogsArgs{
// 						Enable: pulumi.Bool(false),
// 					},
// 				},
// 				S3Logs: &guardduty.DetectorDatasourcesS3LogsArgs{
// 					Enable: pulumi.Bool(true),
// 				},
// 			},
// 			Enable: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// GuardDuty detectors can be imported using the detector ID, e.g.,
//
// ```sh
//  $ pulumi import aws:guardduty/detector:Detector MyDetector 00b00fd5aecc0ab60a708659477e9617
// ```
type Detector struct {
	pulumi.CustomResourceState

	// The AWS account ID of the GuardDuty detector
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Amazon Resource Name (ARN) of the GuardDuty detector
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
	Datasources DetectorDatasourcesOutput `pulumi:"datasources"`
	// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
	// Defaults to `true`.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
	FindingPublishingFrequency pulumi.StringOutput `pulumi:"findingPublishingFrequency"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDetector registers a new resource with the given unique name, arguments, and options.
func NewDetector(ctx *pulumi.Context,
	name string, args *DetectorArgs, opts ...pulumi.ResourceOption) (*Detector, error) {
	if args == nil {
		args = &DetectorArgs{}
	}

	var resource Detector
	err := ctx.RegisterResource("aws:guardduty/detector:Detector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDetector gets an existing Detector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDetector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DetectorState, opts ...pulumi.ResourceOption) (*Detector, error) {
	var resource Detector
	err := ctx.ReadResource("aws:guardduty/detector:Detector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Detector resources.
type detectorState struct {
	// The AWS account ID of the GuardDuty detector
	AccountId *string `pulumi:"accountId"`
	// Amazon Resource Name (ARN) of the GuardDuty detector
	Arn *string `pulumi:"arn"`
	// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
	Datasources *DetectorDatasources `pulumi:"datasources"`
	// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
	// Defaults to `true`.
	Enable *bool `pulumi:"enable"`
	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
	FindingPublishingFrequency *string `pulumi:"findingPublishingFrequency"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DetectorState struct {
	// The AWS account ID of the GuardDuty detector
	AccountId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the GuardDuty detector
	Arn pulumi.StringPtrInput
	// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
	Datasources DetectorDatasourcesPtrInput
	// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
	// Defaults to `true`.
	Enable pulumi.BoolPtrInput
	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
	FindingPublishingFrequency pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (DetectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*detectorState)(nil)).Elem()
}

type detectorArgs struct {
	// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
	Datasources *DetectorDatasources `pulumi:"datasources"`
	// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
	// Defaults to `true`.
	Enable *bool `pulumi:"enable"`
	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
	FindingPublishingFrequency *string `pulumi:"findingPublishingFrequency"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Detector resource.
type DetectorArgs struct {
	// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
	Datasources DetectorDatasourcesPtrInput
	// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
	// Defaults to `true`.
	Enable pulumi.BoolPtrInput
	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
	FindingPublishingFrequency pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (DetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*detectorArgs)(nil)).Elem()
}

type DetectorInput interface {
	pulumi.Input

	ToDetectorOutput() DetectorOutput
	ToDetectorOutputWithContext(ctx context.Context) DetectorOutput
}

func (*Detector) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (i *Detector) ToDetectorOutput() DetectorOutput {
	return i.ToDetectorOutputWithContext(context.Background())
}

func (i *Detector) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput)
}

// DetectorArrayInput is an input type that accepts DetectorArray and DetectorArrayOutput values.
// You can construct a concrete instance of `DetectorArrayInput` via:
//
//          DetectorArray{ DetectorArgs{...} }
type DetectorArrayInput interface {
	pulumi.Input

	ToDetectorArrayOutput() DetectorArrayOutput
	ToDetectorArrayOutputWithContext(context.Context) DetectorArrayOutput
}

type DetectorArray []DetectorInput

func (DetectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Detector)(nil)).Elem()
}

func (i DetectorArray) ToDetectorArrayOutput() DetectorArrayOutput {
	return i.ToDetectorArrayOutputWithContext(context.Background())
}

func (i DetectorArray) ToDetectorArrayOutputWithContext(ctx context.Context) DetectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorArrayOutput)
}

// DetectorMapInput is an input type that accepts DetectorMap and DetectorMapOutput values.
// You can construct a concrete instance of `DetectorMapInput` via:
//
//          DetectorMap{ "key": DetectorArgs{...} }
type DetectorMapInput interface {
	pulumi.Input

	ToDetectorMapOutput() DetectorMapOutput
	ToDetectorMapOutputWithContext(context.Context) DetectorMapOutput
}

type DetectorMap map[string]DetectorInput

func (DetectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Detector)(nil)).Elem()
}

func (i DetectorMap) ToDetectorMapOutput() DetectorMapOutput {
	return i.ToDetectorMapOutputWithContext(context.Background())
}

func (i DetectorMap) ToDetectorMapOutputWithContext(ctx context.Context) DetectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorMapOutput)
}

type DetectorOutput struct{ *pulumi.OutputState }

func (DetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (o DetectorOutput) ToDetectorOutput() DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return o
}

// The AWS account ID of the GuardDuty detector
func (o DetectorOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Detector) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the GuardDuty detector
func (o DetectorOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Detector) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Describes which data sources will be enabled for the detector. See Data Sources below for more details.
func (o DetectorOutput) Datasources() DetectorDatasourcesOutput {
	return o.ApplyT(func(v *Detector) DetectorDatasourcesOutput { return v.Datasources }).(DetectorDatasourcesOutput)
}

// If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
// Defaults to `true`.
func (o DetectorOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Detector) pulumi.BoolPtrOutput { return v.Enable }).(pulumi.BoolPtrOutput)
}

// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
func (o DetectorOutput) FindingPublishingFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Detector) pulumi.StringOutput { return v.FindingPublishingFrequency }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DetectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Detector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider .
func (o DetectorOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Detector) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type DetectorArrayOutput struct{ *pulumi.OutputState }

func (DetectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Detector)(nil)).Elem()
}

func (o DetectorArrayOutput) ToDetectorArrayOutput() DetectorArrayOutput {
	return o
}

func (o DetectorArrayOutput) ToDetectorArrayOutputWithContext(ctx context.Context) DetectorArrayOutput {
	return o
}

func (o DetectorArrayOutput) Index(i pulumi.IntInput) DetectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Detector {
		return vs[0].([]*Detector)[vs[1].(int)]
	}).(DetectorOutput)
}

type DetectorMapOutput struct{ *pulumi.OutputState }

func (DetectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Detector)(nil)).Elem()
}

func (o DetectorMapOutput) ToDetectorMapOutput() DetectorMapOutput {
	return o
}

func (o DetectorMapOutput) ToDetectorMapOutputWithContext(ctx context.Context) DetectorMapOutput {
	return o
}

func (o DetectorMapOutput) MapIndex(k pulumi.StringInput) DetectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Detector {
		return vs[0].(map[string]*Detector)[vs[1].(string)]
	}).(DetectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorInput)(nil)).Elem(), &Detector{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorArrayInput)(nil)).Elem(), DetectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorMapInput)(nil)).Elem(), DetectorMap{})
	pulumi.RegisterOutputType(DetectorOutput{})
	pulumi.RegisterOutputType(DetectorArrayOutput{})
	pulumi.RegisterOutputType(DetectorMapOutput{})
}
