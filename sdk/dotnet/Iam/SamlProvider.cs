// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    /// <summary>
    /// Provides an IAM SAML provider.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Aws.Iam.SamlProvider("default", new()
    ///     {
    ///         SamlMetadataDocument = File.ReadAllText("saml-metadata.xml"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_iam_saml_provider.default
    /// 
    ///  id = "arn:aws:iam::123456789012:saml-provider/SAMLADFS" } Using `pulumi import`, import IAM SAML Providers using the `arn`. For exampleconsole % pulumi import aws_iam_saml_provider.default arn:aws:iam::123456789012:saml-provider/SAMLADFS
    /// </summary>
    [AwsResourceType("aws:iam/samlProvider:SamlProvider")]
    public partial class SamlProvider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN assigned by AWS for this provider.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the provider to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An XML document generated by an identity provider that supports SAML 2.0.
        /// </summary>
        [Output("samlMetadataDocument")]
        public Output<string> SamlMetadataDocument { get; private set; } = null!;

        /// <summary>
        /// Map of resource tags for the IAM SAML provider. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The expiration date and time for the SAML provider in RFC1123 format, e.g., `Mon, 02 Jan 2006 15:04:05 MST`.
        /// </summary>
        [Output("validUntil")]
        public Output<string> ValidUntil { get; private set; } = null!;


        /// <summary>
        /// Create a SamlProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SamlProvider(string name, SamlProviderArgs args, CustomResourceOptions? options = null)
            : base("aws:iam/samlProvider:SamlProvider", name, args ?? new SamlProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SamlProvider(string name, Input<string> id, SamlProviderState? state = null, CustomResourceOptions? options = null)
            : base("aws:iam/samlProvider:SamlProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SamlProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SamlProvider Get(string name, Input<string> id, SamlProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new SamlProvider(name, id, state, options);
        }
    }

    public sealed class SamlProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the provider to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An XML document generated by an identity provider that supports SAML 2.0.
        /// </summary>
        [Input("samlMetadataDocument", required: true)]
        public Input<string> SamlMetadataDocument { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of resource tags for the IAM SAML provider. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SamlProviderArgs()
        {
        }
        public static new SamlProviderArgs Empty => new SamlProviderArgs();
    }

    public sealed class SamlProviderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN assigned by AWS for this provider.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the provider to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An XML document generated by an identity provider that supports SAML 2.0.
        /// </summary>
        [Input("samlMetadataDocument")]
        public Input<string>? SamlMetadataDocument { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of resource tags for the IAM SAML provider. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// The expiration date and time for the SAML provider in RFC1123 format, e.g., `Mon, 02 Jan 2006 15:04:05 MST`.
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public SamlProviderState()
        {
        }
        public static new SamlProviderState Empty => new SamlProviderState();
    }
}
