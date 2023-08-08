// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift
{
    /// <summary>
    /// Creates an HSM client certificate that an Amazon Redshift cluster will use to connect to the client's HSM in order to store and retrieve the keys used to encrypt the cluster databases.
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
    ///     var example = new Aws.RedShift.HsmClientCertificate("example", new()
    ///     {
    ///         HsmClientCertificateIdentifier = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_redshift_hsm_client_certificate.test
    /// 
    ///  id = "example" } Using `pulumi import`, import Redshift HSM Client Certificates using `hsm_client_certificate_identifier`. For exampleconsole % pulumi import aws_redshift_hsm_client_certificate.test example
    /// </summary>
    [AwsResourceType("aws:redshift/hsmClientCertificate:HsmClientCertificate")]
    public partial class HsmClientCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Hsm Client Certificate.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The identifier of the HSM client certificate.
        /// </summary>
        [Output("hsmClientCertificateIdentifier")]
        public Output<string> HsmClientCertificateIdentifier { get; private set; } = null!;

        /// <summary>
        /// The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
        /// </summary>
        [Output("hsmClientCertificatePublicKey")]
        public Output<string> HsmClientCertificatePublicKey { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a HsmClientCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HsmClientCertificate(string name, HsmClientCertificateArgs args, CustomResourceOptions? options = null)
            : base("aws:redshift/hsmClientCertificate:HsmClientCertificate", name, args ?? new HsmClientCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HsmClientCertificate(string name, Input<string> id, HsmClientCertificateState? state = null, CustomResourceOptions? options = null)
            : base("aws:redshift/hsmClientCertificate:HsmClientCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HsmClientCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HsmClientCertificate Get(string name, Input<string> id, HsmClientCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new HsmClientCertificate(name, id, state, options);
        }
    }

    public sealed class HsmClientCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the HSM client certificate.
        /// </summary>
        [Input("hsmClientCertificateIdentifier", required: true)]
        public Input<string> HsmClientCertificateIdentifier { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public HsmClientCertificateArgs()
        {
        }
        public static new HsmClientCertificateArgs Empty => new HsmClientCertificateArgs();
    }

    public sealed class HsmClientCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Hsm Client Certificate.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The identifier of the HSM client certificate.
        /// </summary>
        [Input("hsmClientCertificateIdentifier")]
        public Input<string>? HsmClientCertificateIdentifier { get; set; }

        /// <summary>
        /// The public key that the Amazon Redshift cluster will use to connect to the HSM. You must register the public key in the HSM.
        /// </summary>
        [Input("hsmClientCertificatePublicKey")]
        public Input<string>? HsmClientCertificatePublicKey { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public HsmClientCertificateState()
        {
        }
        public static new HsmClientCertificateState Empty => new HsmClientCertificateState();
    }
}
