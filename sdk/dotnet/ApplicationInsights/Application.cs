// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApplicationInsights
{
    /// <summary>
    /// Provides a ApplicationInsights Application resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleGroup = new Aws.ResourceGroups.Group("exampleGroup", new()
    ///     {
    ///         ResourceQuery = new Aws.ResourceGroups.Inputs.GroupResourceQueryArgs
    ///         {
    ///             Query = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["ResourceTypeFilters"] = new[]
    ///                 {
    ///                     "AWS::EC2::Instance",
    ///                 },
    ///                 ["TagFilters"] = new[]
    ///                 {
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["Key"] = "Stage",
    ///                         ["Values"] = new[]
    ///                         {
    ///                             "Test",
    ///                         },
    ///                     },
    ///                 },
    ///             }),
    ///         },
    ///     });
    /// 
    ///     var exampleApplication = new Aws.ApplicationInsights.Application("exampleApplication", new()
    ///     {
    ///         ResourceGroupName = exampleGroup.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ApplicationInsights Applications can be imported using the `resource_group_name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:applicationinsights/application:Application some some-application
    /// ```
    /// </summary>
    [AwsResourceType("aws:applicationinsights/application:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the Application.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
        /// </summary>
        [Output("autoConfigEnabled")]
        public Output<bool?> AutoConfigEnabled { get; private set; } = null!;

        /// <summary>
        /// Configures all of the resources in the resource group by applying the recommended configurations.
        /// </summary>
        [Output("autoCreate")]
        public Output<bool?> AutoCreate { get; private set; } = null!;

        /// <summary>
        /// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
        /// </summary>
        [Output("cweMonitorEnabled")]
        public Output<bool?> CweMonitorEnabled { get; private set; } = null!;

        /// <summary>
        /// Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED`.
        /// </summary>
        [Output("groupingType")]
        public Output<string?> GroupingType { get; private set; } = null!;

        /// <summary>
        /// When set to `true`, creates opsItems for any problems detected on an application.
        /// </summary>
        [Output("opsCenterEnabled")]
        public Output<bool?> OpsCenterEnabled { get; private set; } = null!;

        /// <summary>
        /// SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
        /// </summary>
        [Output("opsItemSnsTopicArn")]
        public Output<string?> OpsItemSnsTopicArn { get; private set; } = null!;

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("aws:applicationinsights/application:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
            : base("aws:applicationinsights/application:Application", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Application(name, id, state, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
        /// </summary>
        [Input("autoConfigEnabled")]
        public Input<bool>? AutoConfigEnabled { get; set; }

        /// <summary>
        /// Configures all of the resources in the resource group by applying the recommended configurations.
        /// </summary>
        [Input("autoCreate")]
        public Input<bool>? AutoCreate { get; set; }

        /// <summary>
        /// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
        /// </summary>
        [Input("cweMonitorEnabled")]
        public Input<bool>? CweMonitorEnabled { get; set; }

        /// <summary>
        /// Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED`.
        /// </summary>
        [Input("groupingType")]
        public Input<string>? GroupingType { get; set; }

        /// <summary>
        /// When set to `true`, creates opsItems for any problems detected on an application.
        /// </summary>
        [Input("opsCenterEnabled")]
        public Input<bool>? OpsCenterEnabled { get; set; }

        /// <summary>
        /// SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
        /// </summary>
        [Input("opsItemSnsTopicArn")]
        public Input<string>? OpsItemSnsTopicArn { get; set; }

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }

    public sealed class ApplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the Application.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Indicates whether Application Insights automatically configures unmonitored resources in the resource group.
        /// </summary>
        [Input("autoConfigEnabled")]
        public Input<bool>? AutoConfigEnabled { get; set; }

        /// <summary>
        /// Configures all of the resources in the resource group by applying the recommended configurations.
        /// </summary>
        [Input("autoCreate")]
        public Input<bool>? AutoCreate { get; set; }

        /// <summary>
        /// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as instance terminated, failed deployment, and others.
        /// </summary>
        [Input("cweMonitorEnabled")]
        public Input<bool>? CweMonitorEnabled { get; set; }

        /// <summary>
        /// Application Insights can create applications based on a resource group or on an account. To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED`.
        /// </summary>
        [Input("groupingType")]
        public Input<string>? GroupingType { get; set; }

        /// <summary>
        /// When set to `true`, creates opsItems for any problems detected on an application.
        /// </summary>
        [Input("opsCenterEnabled")]
        public Input<bool>? OpsCenterEnabled { get; set; }

        /// <summary>
        /// SNS topic provided to Application Insights that is associated to the created opsItem. Allows you to receive notifications for updates to the opsItem.
        /// </summary>
        [Input("opsItemSnsTopicArn")]
        public Input<string>? OpsItemSnsTopicArn { get; set; }

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public ApplicationState()
        {
        }
        public static new ApplicationState Empty => new ApplicationState();
    }
}
