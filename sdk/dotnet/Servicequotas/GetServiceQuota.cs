// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceQuotas
{
    public static partial class Invokes
    {
        /// <summary>
        /// Retrieve information about a Service Quota.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/servicequotas_service_quota.html.markdown.
        /// </summary>
        public static Task<GetServiceQuotaResult> GetServiceQuota(GetServiceQuotaArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceQuotaResult>("aws:servicequotas/getServiceQuota:getServiceQuota", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetServiceQuotaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Quota code within the service. When configured, the data source directly looks up the service quota. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        /// </summary>
        [Input("quotaCode")]
        public Input<string>? QuotaCode { get; set; }

        /// <summary>
        /// Quota name within the service. When configured, the data source searches through all service quotas to find the matching quota name. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
        /// </summary>
        [Input("quotaName")]
        public Input<string>? QuotaName { get; set; }

        /// <summary>
        /// Service code for the quota. Available values can be found with the [`aws.servicequotas.getService` data source](https://www.terraform.io/docs/providers/aws/d/servicequotas_service.html) or [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
        /// </summary>
        [Input("serviceCode", required: true)]
        public Input<string> ServiceCode { get; set; } = null!;

        public GetServiceQuotaArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetServiceQuotaResult
    {
        /// <summary>
        /// Whether the service quota is adjustable.
        /// </summary>
        public readonly bool Adjustable;
        /// <summary>
        /// Amazon Resource Name (ARN) of the service quota.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Default value of the service quota.
        /// </summary>
        public readonly double DefaultValue;
        /// <summary>
        /// Whether the service quota is global for the AWS account.
        /// </summary>
        public readonly bool GlobalQuota;
        public readonly string QuotaCode;
        public readonly string QuotaName;
        public readonly string ServiceCode;
        /// <summary>
        /// Name of the service.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Current value of the service quota.
        /// </summary>
        public readonly double Value;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetServiceQuotaResult(
            bool adjustable,
            string arn,
            double defaultValue,
            bool globalQuota,
            string quotaCode,
            string quotaName,
            string serviceCode,
            string serviceName,
            double value,
            string id)
        {
            Adjustable = adjustable;
            Arn = arn;
            DefaultValue = defaultValue;
            GlobalQuota = globalQuota;
            QuotaCode = quotaCode;
            QuotaName = quotaName;
            ServiceCode = serviceCode;
            ServiceName = serviceName;
            Value = value;
            Id = id;
        }
    }
}
