// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker
{
    /// <summary>
    /// Provides a SageMaker endpoint configuration resource.
    /// 
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ec = new Aws.Sagemaker.EndpointConfiguration("ec", new()
    ///     {
    ///         ProductionVariants = new[]
    ///         {
    ///             new Aws.Sagemaker.Inputs.EndpointConfigurationProductionVariantArgs
    ///             {
    ///                 VariantName = "variant-1",
    ///                 ModelName = aws_sagemaker_model.M.Name,
    ///                 InitialInstanceCount = 1,
    ///                 InstanceType = "ml.t2.medium",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Name", "foo" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import endpoint configurations using the `name`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:sagemaker/endpointConfiguration:EndpointConfiguration test_endpoint_config endpoint-config-foo
    /// ```
    /// </summary>
    [AwsResourceType("aws:sagemaker/endpointConfiguration:EndpointConfiguration")]
    public partial class EndpointConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies configuration for how an endpoint performs asynchronous inference.
        /// </summary>
        [Output("asyncInferenceConfig")]
        public Output<Outputs.EndpointConfigurationAsyncInferenceConfig?> AsyncInferenceConfig { get; private set; } = null!;

        /// <summary>
        /// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
        /// </summary>
        [Output("dataCaptureConfig")]
        public Output<Outputs.EndpointConfigurationDataCaptureConfig?> DataCaptureConfig { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
        /// </summary>
        [Output("productionVariants")]
        public Output<ImmutableArray<Outputs.EndpointConfigurationProductionVariant>> ProductionVariants { get; private set; } = null!;

        /// <summary>
        /// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants. If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
        /// </summary>
        [Output("shadowProductionVariants")]
        public Output<ImmutableArray<Outputs.EndpointConfigurationShadowProductionVariant>> ShadowProductionVariants { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointConfiguration(string name, EndpointConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:sagemaker/endpointConfiguration:EndpointConfiguration", name, args ?? new EndpointConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointConfiguration(string name, Input<string> id, EndpointConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/endpointConfiguration:EndpointConfiguration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "tagsAll",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EndpointConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointConfiguration Get(string name, Input<string> id, EndpointConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointConfiguration(name, id, state, options);
        }
    }

    public sealed class EndpointConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies configuration for how an endpoint performs asynchronous inference.
        /// </summary>
        [Input("asyncInferenceConfig")]
        public Input<Inputs.EndpointConfigurationAsyncInferenceConfigArgs>? AsyncInferenceConfig { get; set; }

        /// <summary>
        /// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
        /// </summary>
        [Input("dataCaptureConfig")]
        public Input<Inputs.EndpointConfigurationDataCaptureConfigArgs>? DataCaptureConfig { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("productionVariants", required: true)]
        private InputList<Inputs.EndpointConfigurationProductionVariantArgs>? _productionVariants;

        /// <summary>
        /// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
        /// </summary>
        public InputList<Inputs.EndpointConfigurationProductionVariantArgs> ProductionVariants
        {
            get => _productionVariants ?? (_productionVariants = new InputList<Inputs.EndpointConfigurationProductionVariantArgs>());
            set => _productionVariants = value;
        }

        [Input("shadowProductionVariants")]
        private InputList<Inputs.EndpointConfigurationShadowProductionVariantArgs>? _shadowProductionVariants;

        /// <summary>
        /// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants. If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
        /// </summary>
        public InputList<Inputs.EndpointConfigurationShadowProductionVariantArgs> ShadowProductionVariants
        {
            get => _shadowProductionVariants ?? (_shadowProductionVariants = new InputList<Inputs.EndpointConfigurationShadowProductionVariantArgs>());
            set => _shadowProductionVariants = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EndpointConfigurationArgs()
        {
        }
        public static new EndpointConfigurationArgs Empty => new EndpointConfigurationArgs();
    }

    public sealed class EndpointConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies configuration for how an endpoint performs asynchronous inference.
        /// </summary>
        [Input("asyncInferenceConfig")]
        public Input<Inputs.EndpointConfigurationAsyncInferenceConfigGetArgs>? AsyncInferenceConfig { get; set; }

        /// <summary>
        /// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
        /// </summary>
        [Input("dataCaptureConfig")]
        public Input<Inputs.EndpointConfigurationDataCaptureConfigGetArgs>? DataCaptureConfig { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("productionVariants")]
        private InputList<Inputs.EndpointConfigurationProductionVariantGetArgs>? _productionVariants;

        /// <summary>
        /// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
        /// </summary>
        public InputList<Inputs.EndpointConfigurationProductionVariantGetArgs> ProductionVariants
        {
            get => _productionVariants ?? (_productionVariants = new InputList<Inputs.EndpointConfigurationProductionVariantGetArgs>());
            set => _productionVariants = value;
        }

        [Input("shadowProductionVariants")]
        private InputList<Inputs.EndpointConfigurationShadowProductionVariantGetArgs>? _shadowProductionVariants;

        /// <summary>
        /// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants. If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
        /// </summary>
        public InputList<Inputs.EndpointConfigurationShadowProductionVariantGetArgs> ShadowProductionVariants
        {
            get => _shadowProductionVariants ?? (_shadowProductionVariants = new InputList<Inputs.EndpointConfigurationShadowProductionVariantGetArgs>());
            set => _shadowProductionVariants = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        [Obsolete(@"Please use `tags` instead.")]
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _tagsAll = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        public EndpointConfigurationState()
        {
        }
        public static new EndpointConfigurationState Empty => new EndpointConfigurationState();
    }
}
