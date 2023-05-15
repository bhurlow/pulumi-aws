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
    /// Provides an AppConfig Deployment resource for an `aws.appconfig.Application` resource.
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
    ///     var example = new Aws.AppConfig.Deployment("example", new()
    ///     {
    ///         ApplicationId = aws_appconfig_application.Example.Id,
    ///         ConfigurationProfileId = aws_appconfig_configuration_profile.Example.Configuration_profile_id,
    ///         ConfigurationVersion = aws_appconfig_hosted_configuration_version.Example.Version_number,
    ///         DeploymentStrategyId = aws_appconfig_deployment_strategy.Example.Id,
    ///         Description = "My example deployment",
    ///         EnvironmentId = aws_appconfig_environment.Example.Environment_id,
    ///         Tags = 
    ///         {
    ///             { "Type", "AppConfig Deployment" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AppConfig Deployments can be imported by using the application ID, environment ID, and deployment number separated by a slash (`/`), e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:appconfig/deployment:Deployment example 71abcde/11xxxxx/1
    /// ```
    /// </summary>
    [AwsResourceType("aws:appconfig/deployment:Deployment")]
    public partial class Deployment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// ARN of the AppConfig Deployment.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Configuration profile ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Output("configurationProfileId")]
        public Output<string> ConfigurationProfileId { get; private set; } = null!;

        /// <summary>
        /// Configuration version to deploy. Can be at most 1024 characters.
        /// </summary>
        [Output("configurationVersion")]
        public Output<string> ConfigurationVersion { get; private set; } = null!;

        /// <summary>
        /// Deployment number.
        /// </summary>
        [Output("deploymentNumber")]
        public Output<int> DeploymentNumber { get; private set; } = null!;

        /// <summary>
        /// Deployment strategy ID or name of a predefined deployment strategy. See [Predefined Deployment Strategies](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html#appconfig-creating-deployment-strategy-predefined) for more details.
        /// </summary>
        [Output("deploymentStrategyId")]
        public Output<string> DeploymentStrategyId { get; private set; } = null!;

        /// <summary>
        /// Description of the deployment. Can be at most 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Environment ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// State of the deployment.
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
        /// Create a Deployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Deployment(string name, DeploymentArgs args, CustomResourceOptions? options = null)
            : base("aws:appconfig/deployment:Deployment", name, args ?? new DeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Deployment(string name, Input<string> id, DeploymentState? state = null, CustomResourceOptions? options = null)
            : base("aws:appconfig/deployment:Deployment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Deployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Deployment Get(string name, Input<string> id, DeploymentState? state = null, CustomResourceOptions? options = null)
        {
            return new Deployment(name, id, state, options);
        }
    }

    public sealed class DeploymentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Configuration profile ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("configurationProfileId", required: true)]
        public Input<string> ConfigurationProfileId { get; set; } = null!;

        /// <summary>
        /// Configuration version to deploy. Can be at most 1024 characters.
        /// </summary>
        [Input("configurationVersion", required: true)]
        public Input<string> ConfigurationVersion { get; set; } = null!;

        /// <summary>
        /// Deployment strategy ID or name of a predefined deployment strategy. See [Predefined Deployment Strategies](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html#appconfig-creating-deployment-strategy-predefined) for more details.
        /// </summary>
        [Input("deploymentStrategyId", required: true)]
        public Input<string> DeploymentStrategyId { get; set; } = null!;

        /// <summary>
        /// Description of the deployment. Can be at most 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Environment ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

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

        public DeploymentArgs()
        {
        }
        public static new DeploymentArgs Empty => new DeploymentArgs();
    }

    public sealed class DeploymentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// ARN of the AppConfig Deployment.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Configuration profile ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("configurationProfileId")]
        public Input<string>? ConfigurationProfileId { get; set; }

        /// <summary>
        /// Configuration version to deploy. Can be at most 1024 characters.
        /// </summary>
        [Input("configurationVersion")]
        public Input<string>? ConfigurationVersion { get; set; }

        /// <summary>
        /// Deployment number.
        /// </summary>
        [Input("deploymentNumber")]
        public Input<int>? DeploymentNumber { get; set; }

        /// <summary>
        /// Deployment strategy ID or name of a predefined deployment strategy. See [Predefined Deployment Strategies](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html#appconfig-creating-deployment-strategy-predefined) for more details.
        /// </summary>
        [Input("deploymentStrategyId")]
        public Input<string>? DeploymentStrategyId { get; set; }

        /// <summary>
        /// Description of the deployment. Can be at most 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Environment ID. Must be between 4 and 7 characters in length.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// State of the deployment.
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

        public DeploymentState()
        {
        }
        public static new DeploymentState Empty => new DeploymentState();
    }
}
