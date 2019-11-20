// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an API Gateway Usage Plan.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_usage_plan.html.markdown.
    /// </summary>
    public partial class UsagePlan : Pulumi.CustomResource
    {
        /// <summary>
        /// The associated API stages of the usage plan.
        /// </summary>
        [Output("apiStages")]
        public Output<ImmutableArray<Outputs.UsagePlanApiStages>> ApiStages { get; private set; } = null!;

        /// <summary>
        /// The description of a usage plan.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the usage plan.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        /// </summary>
        [Output("productCode")]
        public Output<string?> ProductCode { get; private set; } = null!;

        /// <summary>
        /// The quota settings of the usage plan.
        /// </summary>
        [Output("quotaSettings")]
        public Output<Outputs.UsagePlanQuotaSettings?> QuotaSettings { get; private set; } = null!;

        /// <summary>
        /// The throttling limits of the usage plan.
        /// </summary>
        [Output("throttleSettings")]
        public Output<Outputs.UsagePlanThrottleSettings?> ThrottleSettings { get; private set; } = null!;


        /// <summary>
        /// Create a UsagePlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UsagePlan(string name, UsagePlanArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/usagePlan:UsagePlan", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private UsagePlan(string name, Input<string> id, UsagePlanState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/usagePlan:UsagePlan", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UsagePlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UsagePlan Get(string name, Input<string> id, UsagePlanState? state = null, CustomResourceOptions? options = null)
        {
            return new UsagePlan(name, id, state, options);
        }
    }

    public sealed class UsagePlanArgs : Pulumi.ResourceArgs
    {
        [Input("apiStages")]
        private InputList<Inputs.UsagePlanApiStagesArgs>? _apiStages;

        /// <summary>
        /// The associated API stages of the usage plan.
        /// </summary>
        public InputList<Inputs.UsagePlanApiStagesArgs> ApiStages
        {
            get => _apiStages ?? (_apiStages = new InputList<Inputs.UsagePlanApiStagesArgs>());
            set => _apiStages = value;
        }

        /// <summary>
        /// The description of a usage plan.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the usage plan.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// The quota settings of the usage plan.
        /// </summary>
        [Input("quotaSettings")]
        public Input<Inputs.UsagePlanQuotaSettingsArgs>? QuotaSettings { get; set; }

        /// <summary>
        /// The throttling limits of the usage plan.
        /// </summary>
        [Input("throttleSettings")]
        public Input<Inputs.UsagePlanThrottleSettingsArgs>? ThrottleSettings { get; set; }

        public UsagePlanArgs()
        {
        }
    }

    public sealed class UsagePlanState : Pulumi.ResourceArgs
    {
        [Input("apiStages")]
        private InputList<Inputs.UsagePlanApiStagesGetArgs>? _apiStages;

        /// <summary>
        /// The associated API stages of the usage plan.
        /// </summary>
        public InputList<Inputs.UsagePlanApiStagesGetArgs> ApiStages
        {
            get => _apiStages ?? (_apiStages = new InputList<Inputs.UsagePlanApiStagesGetArgs>());
            set => _apiStages = value;
        }

        /// <summary>
        /// The description of a usage plan.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the usage plan.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        /// </summary>
        [Input("productCode")]
        public Input<string>? ProductCode { get; set; }

        /// <summary>
        /// The quota settings of the usage plan.
        /// </summary>
        [Input("quotaSettings")]
        public Input<Inputs.UsagePlanQuotaSettingsGetArgs>? QuotaSettings { get; set; }

        /// <summary>
        /// The throttling limits of the usage plan.
        /// </summary>
        [Input("throttleSettings")]
        public Input<Inputs.UsagePlanThrottleSettingsGetArgs>? ThrottleSettings { get; set; }

        public UsagePlanState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class UsagePlanApiStagesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API Id of the associated API stage in a usage plan.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// API stage name of the associated API stage in a usage plan.
        /// </summary>
        [Input("stage", required: true)]
        public Input<string> Stage { get; set; } = null!;

        public UsagePlanApiStagesArgs()
        {
        }
    }

    public sealed class UsagePlanApiStagesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API Id of the associated API stage in a usage plan.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// API stage name of the associated API stage in a usage plan.
        /// </summary>
        [Input("stage", required: true)]
        public Input<string> Stage { get; set; } = null!;

        public UsagePlanApiStagesGetArgs()
        {
        }
    }

    public sealed class UsagePlanQuotaSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of requests that can be made in a given time period.
        /// </summary>
        [Input("limit", required: true)]
        public Input<int> Limit { get; set; } = null!;

        /// <summary>
        /// The number of requests subtracted from the given limit in the initial time period.
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        /// <summary>
        /// The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
        /// </summary>
        [Input("period", required: true)]
        public Input<string> Period { get; set; } = null!;

        public UsagePlanQuotaSettingsArgs()
        {
        }
    }

    public sealed class UsagePlanQuotaSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of requests that can be made in a given time period.
        /// </summary>
        [Input("limit", required: true)]
        public Input<int> Limit { get; set; } = null!;

        /// <summary>
        /// The number of requests subtracted from the given limit in the initial time period.
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        /// <summary>
        /// The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
        /// </summary>
        [Input("period", required: true)]
        public Input<string> Period { get; set; } = null!;

        public UsagePlanQuotaSettingsGetArgs()
        {
        }
    }

    public sealed class UsagePlanThrottleSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
        /// </summary>
        [Input("burstLimit")]
        public Input<int>? BurstLimit { get; set; }

        /// <summary>
        /// The API request steady-state rate limit.
        /// </summary>
        [Input("rateLimit")]
        public Input<double>? RateLimit { get; set; }

        public UsagePlanThrottleSettingsArgs()
        {
        }
    }

    public sealed class UsagePlanThrottleSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
        /// </summary>
        [Input("burstLimit")]
        public Input<int>? BurstLimit { get; set; }

        /// <summary>
        /// The API request steady-state rate limit.
        /// </summary>
        [Input("rateLimit")]
        public Input<double>? RateLimit { get; set; }

        public UsagePlanThrottleSettingsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class UsagePlanApiStages
    {
        /// <summary>
        /// API Id of the associated API stage in a usage plan.
        /// </summary>
        public readonly string ApiId;
        /// <summary>
        /// API stage name of the associated API stage in a usage plan.
        /// </summary>
        public readonly string Stage;

        [OutputConstructor]
        private UsagePlanApiStages(
            string apiId,
            string stage)
        {
            ApiId = apiId;
            Stage = stage;
        }
    }

    [OutputType]
    public sealed class UsagePlanQuotaSettings
    {
        /// <summary>
        /// The maximum number of requests that can be made in a given time period.
        /// </summary>
        public readonly int Limit;
        /// <summary>
        /// The number of requests subtracted from the given limit in the initial time period.
        /// </summary>
        public readonly int? Offset;
        /// <summary>
        /// The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
        /// </summary>
        public readonly string Period;

        [OutputConstructor]
        private UsagePlanQuotaSettings(
            int limit,
            int? offset,
            string period)
        {
            Limit = limit;
            Offset = offset;
            Period = period;
        }
    }

    [OutputType]
    public sealed class UsagePlanThrottleSettings
    {
        /// <summary>
        /// The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.
        /// </summary>
        public readonly int? BurstLimit;
        /// <summary>
        /// The API request steady-state rate limit.
        /// </summary>
        public readonly double? RateLimit;

        [OutputConstructor]
        private UsagePlanThrottleSettings(
            int? burstLimit,
            double? rateLimit)
        {
            BurstLimit = burstLimit;
            RateLimit = rateLimit;
        }
    }
    }
}
