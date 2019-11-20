// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides a CloudWatch Dashboard resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_dashboard.html.markdown.
    /// </summary>
    public partial class Dashboard : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the dashboard.
        /// </summary>
        [Output("dashboardArn")]
        public Output<string> DashboardArn { get; private set; } = null!;

        /// <summary>
        /// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
        /// </summary>
        [Output("dashboardBody")]
        public Output<string> DashboardBody { get; private set; } = null!;

        /// <summary>
        /// The name of the dashboard.
        /// </summary>
        [Output("dashboardName")]
        public Output<string> DashboardName { get; private set; } = null!;


        /// <summary>
        /// Create a Dashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dashboard(string name, DashboardArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/dashboard:Dashboard", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Dashboard(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/dashboard:Dashboard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dashboard Get(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
        {
            return new Dashboard(name, id, state, options);
        }
    }

    public sealed class DashboardArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
        /// </summary>
        [Input("dashboardBody", required: true)]
        public Input<string> DashboardBody { get; set; } = null!;

        /// <summary>
        /// The name of the dashboard.
        /// </summary>
        [Input("dashboardName", required: true)]
        public Input<string> DashboardName { get; set; } = null!;

        public DashboardArgs()
        {
        }
    }

    public sealed class DashboardState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the dashboard.
        /// </summary>
        [Input("dashboardArn")]
        public Input<string>? DashboardArn { get; set; }

        /// <summary>
        /// The detailed information about the dashboard, including what widgets are included and their location on the dashboard. You can read more about the body structure in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
        /// </summary>
        [Input("dashboardBody")]
        public Input<string>? DashboardBody { get; set; }

        /// <summary>
        /// The name of the dashboard.
        /// </summary>
        [Input("dashboardName")]
        public Input<string>? DashboardName { get; set; }

        public DashboardState()
        {
        }
    }
}
