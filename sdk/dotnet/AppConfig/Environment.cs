// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppConfig
{
    /// <summary>
    /// Provides an AppConfig Environment resource for an `aws.appconfig.Application` resource. One or more environments can be defined for an application.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleApplication = new Aws.AppConfig.Application("exampleApplication", new()
    ///     {
    ///         Description = "Example AppConfig Application",
    ///         Tags = 
    ///         {
    ///             { "Type", "AppConfig Application" },
    ///         },
    ///     });
    /// 
    ///     var exampleEnvironment = new Aws.AppConfig.Environment("exampleEnvironment", new()
    ///     {
    ///         Description = "Example AppConfig Environment",
    ///         ApplicationId = exampleApplication.Id,
    ///         Monitors = new[]
    ///         {
    ///             new Aws.AppConfig.Inputs.EnvironmentMonitorArgs
    ///             {
    ///                 AlarmArn = aws_cloudwatch_metric_alarm.Example.Arn,
    ///                 AlarmRoleArn = aws_iam_role.Example.Arn,
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Type", "AppConfig Environment" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AppConfig Environments can be imported by using the environment ID and application ID separated by a colon (`:`), e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:appconfig/environment:Environment example 71abcde:11xxxxx
    /// ```
    /// </summary>
    [AwsResourceType("aws:appconfig/environment:Environment")]
    public partial class Environment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AppConfig application ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// ARN of the AppConfig Environment.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the environment. Can be at most 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// AppConfig environment ID.
        /// </summary>
        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
        /// </summary>
        [Output("monitors")]
        public Output<ImmutableArray<Outputs.EnvironmentMonitor>> Monitors { get; private set; } = null!;

        /// <summary>
        /// Name for the environment. Must be between 1 and 64 characters in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// State of the environment. Possible values are `READY_FOR_DEPLOYMENT`, `DEPLOYING`, `ROLLING_BACK`
        /// or `ROLLED_BACK`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

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
        /// Create a Environment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Environment(string name, EnvironmentArgs args, CustomResourceOptions? options = null)
            : base("aws:appconfig/environment:Environment", name, args ?? new EnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Environment(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:appconfig/environment:Environment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Environment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Environment Get(string name, Input<string> id, EnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Environment(name, id, state, options);
        }
    }

    public sealed class EnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AppConfig application ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Description of the environment. Can be at most 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("monitors")]
        private InputList<Inputs.EnvironmentMonitorArgs>? _monitors;

        /// <summary>
        /// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
        /// </summary>
        public InputList<Inputs.EnvironmentMonitorArgs> Monitors
        {
            get => _monitors ?? (_monitors = new InputList<Inputs.EnvironmentMonitorArgs>());
            set => _monitors = value;
        }

        /// <summary>
        /// Name for the environment. Must be between 1 and 64 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public EnvironmentArgs()
        {
        }
        public static new EnvironmentArgs Empty => new EnvironmentArgs();
    }

    public sealed class EnvironmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AppConfig application ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// ARN of the AppConfig Environment.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the environment. Can be at most 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// AppConfig environment ID.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        [Input("monitors")]
        private InputList<Inputs.EnvironmentMonitorGetArgs>? _monitors;

        /// <summary>
        /// Set of Amazon CloudWatch alarms to monitor during the deployment process. Maximum of 5. See Monitor below for more details.
        /// </summary>
        public InputList<Inputs.EnvironmentMonitorGetArgs> Monitors
        {
            get => _monitors ?? (_monitors = new InputList<Inputs.EnvironmentMonitorGetArgs>());
            set => _monitors = value;
        }

        /// <summary>
        /// Name for the environment. Must be between 1 and 64 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// State of the environment. Possible values are `READY_FOR_DEPLOYMENT`, `DEPLOYING`, `ROLLING_BACK`
        /// or `ROLLED_BACK`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

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

        public EnvironmentState()
        {
        }
        public static new EnvironmentState Empty => new EnvironmentState();
    }
}
