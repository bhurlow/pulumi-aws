// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS CloudFront Continuous Deployment Policy.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			staging, err := cloudfront.NewDistribution(ctx, "staging", &cloudfront.DistributionArgs{
//				Enabled: pulumi.Bool(true),
//				Staging: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := cloudfront.NewContinuousDeploymentPolicy(ctx, "example", &cloudfront.ContinuousDeploymentPolicyArgs{
//				Enabled: pulumi.Bool(true),
//				StagingDistributionDnsNames: &cloudfront.ContinuousDeploymentPolicyStagingDistributionDnsNamesArgs{
//					Items: pulumi.StringArray{
//						staging.DomainName,
//					},
//					Quantity: pulumi.Int(1),
//				},
//				TrafficConfig: &cloudfront.ContinuousDeploymentPolicyTrafficConfigArgs{
//					Type: pulumi.String("SingleWeight"),
//					SingleWeightConfig: &cloudfront.ContinuousDeploymentPolicyTrafficConfigSingleWeightConfigArgs{
//						Weight: pulumi.Float64(0.01),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudfront.NewDistribution(ctx, "production", &cloudfront.DistributionArgs{
//				Enabled:                      pulumi.Bool(true),
//				ContinuousDeploymentPolicyId: example.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Single Weight Config with Session Stickiness
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfront.NewContinuousDeploymentPolicy(ctx, "example", &cloudfront.ContinuousDeploymentPolicyArgs{
//				Enabled: pulumi.Bool(true),
//				StagingDistributionDnsNames: &cloudfront.ContinuousDeploymentPolicyStagingDistributionDnsNamesArgs{
//					Items: pulumi.StringArray{
//						aws_cloudfront_distribution.Staging.Domain_name,
//					},
//					Quantity: pulumi.Int(1),
//				},
//				TrafficConfig: &cloudfront.ContinuousDeploymentPolicyTrafficConfigArgs{
//					Type: pulumi.String("SingleWeight"),
//					SingleWeightConfig: &cloudfront.ContinuousDeploymentPolicyTrafficConfigSingleWeightConfigArgs{
//						Weight: pulumi.Float64(0.01),
//						SessionStickinessConfig: &cloudfront.ContinuousDeploymentPolicyTrafficConfigSingleWeightConfigSessionStickinessConfigArgs{
//							IdleTtl:    pulumi.Int(300),
//							MaximumTtl: pulumi.Int(600),
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
// ### Single Header Config
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfront.NewContinuousDeploymentPolicy(ctx, "example", &cloudfront.ContinuousDeploymentPolicyArgs{
//				Enabled: pulumi.Bool(true),
//				StagingDistributionDnsNames: &cloudfront.ContinuousDeploymentPolicyStagingDistributionDnsNamesArgs{
//					Items: pulumi.StringArray{
//						aws_cloudfront_distribution.Staging.Domain_name,
//					},
//					Quantity: pulumi.Int(1),
//				},
//				TrafficConfig: &cloudfront.ContinuousDeploymentPolicyTrafficConfigArgs{
//					Type: pulumi.String("SingleHeader"),
//					SingleHeaderConfig: &cloudfront.ContinuousDeploymentPolicyTrafficConfigSingleHeaderConfigArgs{
//						Header: pulumi.String("aws-cf-cd-example"),
//						Value:  pulumi.String("example"),
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
// Using `pulumi import`, import CloudFront Continuous Deployment Policy using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:cloudfront/continuousDeploymentPolicy:ContinuousDeploymentPolicy example abcd-1234
//
// ```
type ContinuousDeploymentPolicy struct {
	pulumi.CustomResourceState

	// Whether this continuous deployment policy is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Current version of the continuous distribution policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Date and time the continuous deployment policy was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// CloudFront domain name of the staging distribution. See `stagingDistributionDnsNames`.
	StagingDistributionDnsNames ContinuousDeploymentPolicyStagingDistributionDnsNamesPtrOutput `pulumi:"stagingDistributionDnsNames"`
	// Parameters for routing production traffic from primary to staging distributions. See `trafficConfig`.
	TrafficConfig ContinuousDeploymentPolicyTrafficConfigPtrOutput `pulumi:"trafficConfig"`
}

// NewContinuousDeploymentPolicy registers a new resource with the given unique name, arguments, and options.
func NewContinuousDeploymentPolicy(ctx *pulumi.Context,
	name string, args *ContinuousDeploymentPolicyArgs, opts ...pulumi.ResourceOption) (*ContinuousDeploymentPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContinuousDeploymentPolicy
	err := ctx.RegisterResource("aws:cloudfront/continuousDeploymentPolicy:ContinuousDeploymentPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContinuousDeploymentPolicy gets an existing ContinuousDeploymentPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContinuousDeploymentPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContinuousDeploymentPolicyState, opts ...pulumi.ResourceOption) (*ContinuousDeploymentPolicy, error) {
	var resource ContinuousDeploymentPolicy
	err := ctx.ReadResource("aws:cloudfront/continuousDeploymentPolicy:ContinuousDeploymentPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContinuousDeploymentPolicy resources.
type continuousDeploymentPolicyState struct {
	// Whether this continuous deployment policy is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Current version of the continuous distribution policy.
	Etag *string `pulumi:"etag"`
	// Date and time the continuous deployment policy was last modified.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// CloudFront domain name of the staging distribution. See `stagingDistributionDnsNames`.
	StagingDistributionDnsNames *ContinuousDeploymentPolicyStagingDistributionDnsNames `pulumi:"stagingDistributionDnsNames"`
	// Parameters for routing production traffic from primary to staging distributions. See `trafficConfig`.
	TrafficConfig *ContinuousDeploymentPolicyTrafficConfig `pulumi:"trafficConfig"`
}

type ContinuousDeploymentPolicyState struct {
	// Whether this continuous deployment policy is enabled.
	Enabled pulumi.BoolPtrInput
	// Current version of the continuous distribution policy.
	Etag pulumi.StringPtrInput
	// Date and time the continuous deployment policy was last modified.
	LastModifiedTime pulumi.StringPtrInput
	// CloudFront domain name of the staging distribution. See `stagingDistributionDnsNames`.
	StagingDistributionDnsNames ContinuousDeploymentPolicyStagingDistributionDnsNamesPtrInput
	// Parameters for routing production traffic from primary to staging distributions. See `trafficConfig`.
	TrafficConfig ContinuousDeploymentPolicyTrafficConfigPtrInput
}

func (ContinuousDeploymentPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*continuousDeploymentPolicyState)(nil)).Elem()
}

type continuousDeploymentPolicyArgs struct {
	// Whether this continuous deployment policy is enabled.
	Enabled bool `pulumi:"enabled"`
	// CloudFront domain name of the staging distribution. See `stagingDistributionDnsNames`.
	StagingDistributionDnsNames *ContinuousDeploymentPolicyStagingDistributionDnsNames `pulumi:"stagingDistributionDnsNames"`
	// Parameters for routing production traffic from primary to staging distributions. See `trafficConfig`.
	TrafficConfig *ContinuousDeploymentPolicyTrafficConfig `pulumi:"trafficConfig"`
}

// The set of arguments for constructing a ContinuousDeploymentPolicy resource.
type ContinuousDeploymentPolicyArgs struct {
	// Whether this continuous deployment policy is enabled.
	Enabled pulumi.BoolInput
	// CloudFront domain name of the staging distribution. See `stagingDistributionDnsNames`.
	StagingDistributionDnsNames ContinuousDeploymentPolicyStagingDistributionDnsNamesPtrInput
	// Parameters for routing production traffic from primary to staging distributions. See `trafficConfig`.
	TrafficConfig ContinuousDeploymentPolicyTrafficConfigPtrInput
}

func (ContinuousDeploymentPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*continuousDeploymentPolicyArgs)(nil)).Elem()
}

type ContinuousDeploymentPolicyInput interface {
	pulumi.Input

	ToContinuousDeploymentPolicyOutput() ContinuousDeploymentPolicyOutput
	ToContinuousDeploymentPolicyOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyOutput
}

func (*ContinuousDeploymentPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ContinuousDeploymentPolicy)(nil)).Elem()
}

func (i *ContinuousDeploymentPolicy) ToContinuousDeploymentPolicyOutput() ContinuousDeploymentPolicyOutput {
	return i.ToContinuousDeploymentPolicyOutputWithContext(context.Background())
}

func (i *ContinuousDeploymentPolicy) ToContinuousDeploymentPolicyOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContinuousDeploymentPolicyOutput)
}

// ContinuousDeploymentPolicyArrayInput is an input type that accepts ContinuousDeploymentPolicyArray and ContinuousDeploymentPolicyArrayOutput values.
// You can construct a concrete instance of `ContinuousDeploymentPolicyArrayInput` via:
//
//	ContinuousDeploymentPolicyArray{ ContinuousDeploymentPolicyArgs{...} }
type ContinuousDeploymentPolicyArrayInput interface {
	pulumi.Input

	ToContinuousDeploymentPolicyArrayOutput() ContinuousDeploymentPolicyArrayOutput
	ToContinuousDeploymentPolicyArrayOutputWithContext(context.Context) ContinuousDeploymentPolicyArrayOutput
}

type ContinuousDeploymentPolicyArray []ContinuousDeploymentPolicyInput

func (ContinuousDeploymentPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContinuousDeploymentPolicy)(nil)).Elem()
}

func (i ContinuousDeploymentPolicyArray) ToContinuousDeploymentPolicyArrayOutput() ContinuousDeploymentPolicyArrayOutput {
	return i.ToContinuousDeploymentPolicyArrayOutputWithContext(context.Background())
}

func (i ContinuousDeploymentPolicyArray) ToContinuousDeploymentPolicyArrayOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContinuousDeploymentPolicyArrayOutput)
}

// ContinuousDeploymentPolicyMapInput is an input type that accepts ContinuousDeploymentPolicyMap and ContinuousDeploymentPolicyMapOutput values.
// You can construct a concrete instance of `ContinuousDeploymentPolicyMapInput` via:
//
//	ContinuousDeploymentPolicyMap{ "key": ContinuousDeploymentPolicyArgs{...} }
type ContinuousDeploymentPolicyMapInput interface {
	pulumi.Input

	ToContinuousDeploymentPolicyMapOutput() ContinuousDeploymentPolicyMapOutput
	ToContinuousDeploymentPolicyMapOutputWithContext(context.Context) ContinuousDeploymentPolicyMapOutput
}

type ContinuousDeploymentPolicyMap map[string]ContinuousDeploymentPolicyInput

func (ContinuousDeploymentPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContinuousDeploymentPolicy)(nil)).Elem()
}

func (i ContinuousDeploymentPolicyMap) ToContinuousDeploymentPolicyMapOutput() ContinuousDeploymentPolicyMapOutput {
	return i.ToContinuousDeploymentPolicyMapOutputWithContext(context.Background())
}

func (i ContinuousDeploymentPolicyMap) ToContinuousDeploymentPolicyMapOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContinuousDeploymentPolicyMapOutput)
}

type ContinuousDeploymentPolicyOutput struct{ *pulumi.OutputState }

func (ContinuousDeploymentPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContinuousDeploymentPolicy)(nil)).Elem()
}

func (o ContinuousDeploymentPolicyOutput) ToContinuousDeploymentPolicyOutput() ContinuousDeploymentPolicyOutput {
	return o
}

func (o ContinuousDeploymentPolicyOutput) ToContinuousDeploymentPolicyOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyOutput {
	return o
}

// Whether this continuous deployment policy is enabled.
func (o ContinuousDeploymentPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ContinuousDeploymentPolicy) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Current version of the continuous distribution policy.
func (o ContinuousDeploymentPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ContinuousDeploymentPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Date and time the continuous deployment policy was last modified.
func (o ContinuousDeploymentPolicyOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ContinuousDeploymentPolicy) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// CloudFront domain name of the staging distribution. See `stagingDistributionDnsNames`.
func (o ContinuousDeploymentPolicyOutput) StagingDistributionDnsNames() ContinuousDeploymentPolicyStagingDistributionDnsNamesPtrOutput {
	return o.ApplyT(func(v *ContinuousDeploymentPolicy) ContinuousDeploymentPolicyStagingDistributionDnsNamesPtrOutput {
		return v.StagingDistributionDnsNames
	}).(ContinuousDeploymentPolicyStagingDistributionDnsNamesPtrOutput)
}

// Parameters for routing production traffic from primary to staging distributions. See `trafficConfig`.
func (o ContinuousDeploymentPolicyOutput) TrafficConfig() ContinuousDeploymentPolicyTrafficConfigPtrOutput {
	return o.ApplyT(func(v *ContinuousDeploymentPolicy) ContinuousDeploymentPolicyTrafficConfigPtrOutput {
		return v.TrafficConfig
	}).(ContinuousDeploymentPolicyTrafficConfigPtrOutput)
}

type ContinuousDeploymentPolicyArrayOutput struct{ *pulumi.OutputState }

func (ContinuousDeploymentPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContinuousDeploymentPolicy)(nil)).Elem()
}

func (o ContinuousDeploymentPolicyArrayOutput) ToContinuousDeploymentPolicyArrayOutput() ContinuousDeploymentPolicyArrayOutput {
	return o
}

func (o ContinuousDeploymentPolicyArrayOutput) ToContinuousDeploymentPolicyArrayOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyArrayOutput {
	return o
}

func (o ContinuousDeploymentPolicyArrayOutput) Index(i pulumi.IntInput) ContinuousDeploymentPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContinuousDeploymentPolicy {
		return vs[0].([]*ContinuousDeploymentPolicy)[vs[1].(int)]
	}).(ContinuousDeploymentPolicyOutput)
}

type ContinuousDeploymentPolicyMapOutput struct{ *pulumi.OutputState }

func (ContinuousDeploymentPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContinuousDeploymentPolicy)(nil)).Elem()
}

func (o ContinuousDeploymentPolicyMapOutput) ToContinuousDeploymentPolicyMapOutput() ContinuousDeploymentPolicyMapOutput {
	return o
}

func (o ContinuousDeploymentPolicyMapOutput) ToContinuousDeploymentPolicyMapOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyMapOutput {
	return o
}

func (o ContinuousDeploymentPolicyMapOutput) MapIndex(k pulumi.StringInput) ContinuousDeploymentPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContinuousDeploymentPolicy {
		return vs[0].(map[string]*ContinuousDeploymentPolicy)[vs[1].(string)]
	}).(ContinuousDeploymentPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContinuousDeploymentPolicyInput)(nil)).Elem(), &ContinuousDeploymentPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContinuousDeploymentPolicyArrayInput)(nil)).Elem(), ContinuousDeploymentPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContinuousDeploymentPolicyMapInput)(nil)).Elem(), ContinuousDeploymentPolicyMap{})
	pulumi.RegisterOutputType(ContinuousDeploymentPolicyOutput{})
	pulumi.RegisterOutputType(ContinuousDeploymentPolicyArrayOutput{})
	pulumi.RegisterOutputType(ContinuousDeploymentPolicyMapOutput{})
}
