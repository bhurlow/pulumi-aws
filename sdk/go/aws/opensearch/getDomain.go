// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an OpenSearch Domain
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := opensearch.LookupDomain(ctx, &opensearch.LookupDomainArgs{
//				DomainName: "my-domain-name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("aws:opensearch/getDomain:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomain.
type LookupDomainArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// Off Peak update options
	OffPeakWindowOptions *GetDomainOffPeakWindowOptions `pulumi:"offPeakWindowOptions"`
	// Tags assigned to the domain.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDomain.
type LookupDomainResult struct {
	// Policy document attached to the domain.
	AccessPolicies string `pulumi:"accessPolicies"`
	// Key-value string pairs to specify advanced configuration options.
	AdvancedOptions map[string]string `pulumi:"advancedOptions"`
	// Status of the OpenSearch domain's advanced security options. The block consists of the following attributes:
	AdvancedSecurityOptions []GetDomainAdvancedSecurityOption `pulumi:"advancedSecurityOptions"`
	// ARN of the domain.
	Arn string `pulumi:"arn"`
	// Configuration of the Auto-Tune options of the domain.
	AutoTuneOptions []GetDomainAutoTuneOption `pulumi:"autoTuneOptions"`
	// Cluster configuration of the domain.
	ClusterConfigs []GetDomainClusterConfig `pulumi:"clusterConfigs"`
	// Domain Amazon Cognito Authentication options for Dashboard.
	CognitoOptions []GetDomainCognitoOption `pulumi:"cognitoOptions"`
	// Status of the creation of the domain.
	Created bool `pulumi:"created"`
	// Domain-specific endpoint used to access the [Dashboard application](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/dashboards.html).
	DashboardEndpoint string `pulumi:"dashboardEndpoint"`
	// Status of the deletion of the domain.
	Deleted bool `pulumi:"deleted"`
	// Unique identifier for the domain.
	DomainId   string `pulumi:"domainId"`
	DomainName string `pulumi:"domainName"`
	// EBS Options for the instances in the domain.
	EbsOptions []GetDomainEbsOption `pulumi:"ebsOptions"`
	// Domain encryption at rest related options.
	EncryptionAtRests []GetDomainEncryptionAtRest `pulumi:"encryptionAtRests"`
	// Domain-specific endpoint used to submit index, search, and data upload requests.
	Endpoint string `pulumi:"endpoint"`
	// OpenSearch version for the domain.
	EngineVersion string `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (**Deprecated**) Domain-specific endpoint for kibana without https scheme. Use the `dashboardEndpoint` attribute instead.
	//
	// Deprecated: use 'dashboard_endpoint' attribute instead
	KibanaEndpoint string `pulumi:"kibanaEndpoint"`
	// Domain log publishing related options.
	LogPublishingOptions []GetDomainLogPublishingOption `pulumi:"logPublishingOptions"`
	// Domain in transit encryption related options.
	NodeToNodeEncryptions []GetDomainNodeToNodeEncryption `pulumi:"nodeToNodeEncryptions"`
	// Off Peak update options
	OffPeakWindowOptions *GetDomainOffPeakWindowOptions `pulumi:"offPeakWindowOptions"`
	// Status of a configuration change in the domain.
	Processing bool `pulumi:"processing"`
	// Domain snapshot related options.
	SnapshotOptions []GetDomainSnapshotOption `pulumi:"snapshotOptions"`
	// Software update options for the domain
	SoftwareUpdateOptions []GetDomainSoftwareUpdateOption `pulumi:"softwareUpdateOptions"`
	// Tags assigned to the domain.
	Tags map[string]string `pulumi:"tags"`
	// VPC Options for private OpenSearch domains.
	VpcOptions []GetDomainVpcOption `pulumi:"vpcOptions"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

// A collection of arguments for invoking getDomain.
type LookupDomainOutputArgs struct {
	// Name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Off Peak update options
	OffPeakWindowOptions GetDomainOffPeakWindowOptionsPtrInput `pulumi:"offPeakWindowOptions"`
	// Tags assigned to the domain.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

// A collection of values returned by getDomain.
type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// Policy document attached to the domain.
func (o LookupDomainResultOutput) AccessPolicies() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.AccessPolicies }).(pulumi.StringOutput)
}

// Key-value string pairs to specify advanced configuration options.
func (o LookupDomainResultOutput) AdvancedOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.AdvancedOptions }).(pulumi.StringMapOutput)
}

// Status of the OpenSearch domain's advanced security options. The block consists of the following attributes:
func (o LookupDomainResultOutput) AdvancedSecurityOptions() GetDomainAdvancedSecurityOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainAdvancedSecurityOption { return v.AdvancedSecurityOptions }).(GetDomainAdvancedSecurityOptionArrayOutput)
}

// ARN of the domain.
func (o LookupDomainResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Configuration of the Auto-Tune options of the domain.
func (o LookupDomainResultOutput) AutoTuneOptions() GetDomainAutoTuneOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainAutoTuneOption { return v.AutoTuneOptions }).(GetDomainAutoTuneOptionArrayOutput)
}

// Cluster configuration of the domain.
func (o LookupDomainResultOutput) ClusterConfigs() GetDomainClusterConfigArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainClusterConfig { return v.ClusterConfigs }).(GetDomainClusterConfigArrayOutput)
}

// Domain Amazon Cognito Authentication options for Dashboard.
func (o LookupDomainResultOutput) CognitoOptions() GetDomainCognitoOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainCognitoOption { return v.CognitoOptions }).(GetDomainCognitoOptionArrayOutput)
}

// Status of the creation of the domain.
func (o LookupDomainResultOutput) Created() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainResult) bool { return v.Created }).(pulumi.BoolOutput)
}

// Domain-specific endpoint used to access the [Dashboard application](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/dashboards.html).
func (o LookupDomainResultOutput) DashboardEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DashboardEndpoint }).(pulumi.StringOutput)
}

// Status of the deletion of the domain.
func (o LookupDomainResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

// Unique identifier for the domain.
func (o LookupDomainResultOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DomainId }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// EBS Options for the instances in the domain.
func (o LookupDomainResultOutput) EbsOptions() GetDomainEbsOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainEbsOption { return v.EbsOptions }).(GetDomainEbsOptionArrayOutput)
}

// Domain encryption at rest related options.
func (o LookupDomainResultOutput) EncryptionAtRests() GetDomainEncryptionAtRestArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainEncryptionAtRest { return v.EncryptionAtRests }).(GetDomainEncryptionAtRestArrayOutput)
}

// Domain-specific endpoint used to submit index, search, and data upload requests.
func (o LookupDomainResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// OpenSearch version for the domain.
func (o LookupDomainResultOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.EngineVersion }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// (**Deprecated**) Domain-specific endpoint for kibana without https scheme. Use the `dashboardEndpoint` attribute instead.
//
// Deprecated: use 'dashboard_endpoint' attribute instead
func (o LookupDomainResultOutput) KibanaEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.KibanaEndpoint }).(pulumi.StringOutput)
}

// Domain log publishing related options.
func (o LookupDomainResultOutput) LogPublishingOptions() GetDomainLogPublishingOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainLogPublishingOption { return v.LogPublishingOptions }).(GetDomainLogPublishingOptionArrayOutput)
}

// Domain in transit encryption related options.
func (o LookupDomainResultOutput) NodeToNodeEncryptions() GetDomainNodeToNodeEncryptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainNodeToNodeEncryption { return v.NodeToNodeEncryptions }).(GetDomainNodeToNodeEncryptionArrayOutput)
}

// Off Peak update options
func (o LookupDomainResultOutput) OffPeakWindowOptions() GetDomainOffPeakWindowOptionsPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *GetDomainOffPeakWindowOptions { return v.OffPeakWindowOptions }).(GetDomainOffPeakWindowOptionsPtrOutput)
}

// Status of a configuration change in the domain.
func (o LookupDomainResultOutput) Processing() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainResult) bool { return v.Processing }).(pulumi.BoolOutput)
}

// Domain snapshot related options.
func (o LookupDomainResultOutput) SnapshotOptions() GetDomainSnapshotOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainSnapshotOption { return v.SnapshotOptions }).(GetDomainSnapshotOptionArrayOutput)
}

// Software update options for the domain
func (o LookupDomainResultOutput) SoftwareUpdateOptions() GetDomainSoftwareUpdateOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainSoftwareUpdateOption { return v.SoftwareUpdateOptions }).(GetDomainSoftwareUpdateOptionArrayOutput)
}

// Tags assigned to the domain.
func (o LookupDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// VPC Options for private OpenSearch domains.
func (o LookupDomainResultOutput) VpcOptions() GetDomainVpcOptionArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []GetDomainVpcOption { return v.VpcOptions }).(GetDomainVpcOptionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
