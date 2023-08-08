// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Chime
{
    /// <summary>
    /// Resource for managing Amazon Chime SDK Voice Global Settings.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Chime.SdkvoiceGlobalSettings("example", new()
    ///     {
    ///         VoiceConnector = new Aws.Chime.Inputs.SdkvoiceGlobalSettingsVoiceConnectorArgs
    ///         {
    ///             CdrBucket = "example-bucket-name",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_chimesdkvoice_global_settings.example
    /// 
    ///  id = "123456789012" } Using `pulumi import`, import AWS Chime SDK Voice Global Settings using the `id` (AWS account ID). For exampleconsole % pulumi import aws_chimesdkvoice_global_settings.example 123456789012
    /// </summary>
    [AwsResourceType("aws:chime/sdkvoiceGlobalSettings:SdkvoiceGlobalSettings")]
    public partial class SdkvoiceGlobalSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Voice Connector settings. See voice_connector.
        /// </summary>
        [Output("voiceConnector")]
        public Output<Outputs.SdkvoiceGlobalSettingsVoiceConnector> VoiceConnector { get; private set; } = null!;


        /// <summary>
        /// Create a SdkvoiceGlobalSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SdkvoiceGlobalSettings(string name, SdkvoiceGlobalSettingsArgs args, CustomResourceOptions? options = null)
            : base("aws:chime/sdkvoiceGlobalSettings:SdkvoiceGlobalSettings", name, args ?? new SdkvoiceGlobalSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SdkvoiceGlobalSettings(string name, Input<string> id, SdkvoiceGlobalSettingsState? state = null, CustomResourceOptions? options = null)
            : base("aws:chime/sdkvoiceGlobalSettings:SdkvoiceGlobalSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SdkvoiceGlobalSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SdkvoiceGlobalSettings Get(string name, Input<string> id, SdkvoiceGlobalSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new SdkvoiceGlobalSettings(name, id, state, options);
        }
    }

    public sealed class SdkvoiceGlobalSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Voice Connector settings. See voice_connector.
        /// </summary>
        [Input("voiceConnector", required: true)]
        public Input<Inputs.SdkvoiceGlobalSettingsVoiceConnectorArgs> VoiceConnector { get; set; } = null!;

        public SdkvoiceGlobalSettingsArgs()
        {
        }
        public static new SdkvoiceGlobalSettingsArgs Empty => new SdkvoiceGlobalSettingsArgs();
    }

    public sealed class SdkvoiceGlobalSettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Voice Connector settings. See voice_connector.
        /// </summary>
        [Input("voiceConnector")]
        public Input<Inputs.SdkvoiceGlobalSettingsVoiceConnectorGetArgs>? VoiceConnector { get; set; }

        public SdkvoiceGlobalSettingsState()
        {
        }
        public static new SdkvoiceGlobalSettingsState Empty => new SdkvoiceGlobalSettingsState();
    }
}
