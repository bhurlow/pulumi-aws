// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodePipeline
{
    /// <summary>
    /// Provides a CodeDeploy CustomActionType
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.CodePipeline.CustomActionType("example", new()
    ///     {
    ///         Category = "Build",
    ///         InputArtifactDetails = new Aws.CodePipeline.Inputs.CustomActionTypeInputArtifactDetailsArgs
    ///         {
    ///             MaximumCount = 1,
    ///             MinimumCount = 0,
    ///         },
    ///         OutputArtifactDetails = new Aws.CodePipeline.Inputs.CustomActionTypeOutputArtifactDetailsArgs
    ///         {
    ///             MaximumCount = 1,
    ///             MinimumCount = 0,
    ///         },
    ///         ProviderName = "example",
    ///         Version = "1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CodeDeploy CustomActionType can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:codepipeline/customActionType:CustomActionType example Build:terraform:1
    /// ```
    /// </summary>
    [AwsResourceType("aws:codepipeline/customActionType:CustomActionType")]
    public partial class CustomActionType : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The action ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The category of the custom action. Valid values: `Source`, `Build`, `Deploy`, `Test`, `Invoke`, `Approval`
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// The configuration properties for the custom action. Max 10 items.
        /// </summary>
        [Output("configurationProperties")]
        public Output<ImmutableArray<Outputs.CustomActionTypeConfigurationProperty>> ConfigurationProperties { get; private set; } = null!;

        /// <summary>
        /// The details of the input artifact for the action.
        /// </summary>
        [Output("inputArtifactDetails")]
        public Output<Outputs.CustomActionTypeInputArtifactDetails> InputArtifactDetails { get; private set; } = null!;

        /// <summary>
        /// The details of the output artifact of the action.
        /// </summary>
        [Output("outputArtifactDetails")]
        public Output<Outputs.CustomActionTypeOutputArtifactDetails> OutputArtifactDetails { get; private set; } = null!;

        /// <summary>
        /// The creator of the action being called.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// The provider of the service used in the custom action
        /// </summary>
        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;

        /// <summary>
        /// The settings for an action type.
        /// </summary>
        [Output("settings")]
        public Output<Outputs.CustomActionTypeSettings?> Settings { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The version identifier of the custom action.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CustomActionType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomActionType(string name, CustomActionTypeArgs args, CustomResourceOptions? options = null)
            : base("aws:codepipeline/customActionType:CustomActionType", name, args ?? new CustomActionTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomActionType(string name, Input<string> id, CustomActionTypeState? state = null, CustomResourceOptions? options = null)
            : base("aws:codepipeline/customActionType:CustomActionType", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomActionType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomActionType Get(string name, Input<string> id, CustomActionTypeState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomActionType(name, id, state, options);
        }
    }

    public sealed class CustomActionTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The category of the custom action. Valid values: `Source`, `Build`, `Deploy`, `Test`, `Invoke`, `Approval`
        /// </summary>
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        [Input("configurationProperties")]
        private InputList<Inputs.CustomActionTypeConfigurationPropertyArgs>? _configurationProperties;

        /// <summary>
        /// The configuration properties for the custom action. Max 10 items.
        /// </summary>
        public InputList<Inputs.CustomActionTypeConfigurationPropertyArgs> ConfigurationProperties
        {
            get => _configurationProperties ?? (_configurationProperties = new InputList<Inputs.CustomActionTypeConfigurationPropertyArgs>());
            set => _configurationProperties = value;
        }

        /// <summary>
        /// The details of the input artifact for the action.
        /// </summary>
        [Input("inputArtifactDetails", required: true)]
        public Input<Inputs.CustomActionTypeInputArtifactDetailsArgs> InputArtifactDetails { get; set; } = null!;

        /// <summary>
        /// The details of the output artifact of the action.
        /// </summary>
        [Input("outputArtifactDetails", required: true)]
        public Input<Inputs.CustomActionTypeOutputArtifactDetailsArgs> OutputArtifactDetails { get; set; } = null!;

        /// <summary>
        /// The provider of the service used in the custom action
        /// </summary>
        [Input("providerName", required: true)]
        public Input<string> ProviderName { get; set; } = null!;

        /// <summary>
        /// The settings for an action type.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.CustomActionTypeSettingsArgs>? Settings { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The version identifier of the custom action.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public CustomActionTypeArgs()
        {
        }
        public static new CustomActionTypeArgs Empty => new CustomActionTypeArgs();
    }

    public sealed class CustomActionTypeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The category of the custom action. Valid values: `Source`, `Build`, `Deploy`, `Test`, `Invoke`, `Approval`
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("configurationProperties")]
        private InputList<Inputs.CustomActionTypeConfigurationPropertyGetArgs>? _configurationProperties;

        /// <summary>
        /// The configuration properties for the custom action. Max 10 items.
        /// </summary>
        public InputList<Inputs.CustomActionTypeConfigurationPropertyGetArgs> ConfigurationProperties
        {
            get => _configurationProperties ?? (_configurationProperties = new InputList<Inputs.CustomActionTypeConfigurationPropertyGetArgs>());
            set => _configurationProperties = value;
        }

        /// <summary>
        /// The details of the input artifact for the action.
        /// </summary>
        [Input("inputArtifactDetails")]
        public Input<Inputs.CustomActionTypeInputArtifactDetailsGetArgs>? InputArtifactDetails { get; set; }

        /// <summary>
        /// The details of the output artifact of the action.
        /// </summary>
        [Input("outputArtifactDetails")]
        public Input<Inputs.CustomActionTypeOutputArtifactDetailsGetArgs>? OutputArtifactDetails { get; set; }

        /// <summary>
        /// The creator of the action being called.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The provider of the service used in the custom action
        /// </summary>
        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        /// <summary>
        /// The settings for an action type.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.CustomActionTypeSettingsGetArgs>? Settings { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The version identifier of the custom action.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public CustomActionTypeState()
        {
        }
        public static new CustomActionTypeState Empty => new CustomActionTypeState();
    }
}
