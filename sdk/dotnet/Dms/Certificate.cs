// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms
{
    /// <summary>
    /// Provides a DMS (Data Migration Service) certificate resource. DMS certificates can be created, deleted, and imported.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a new certificate
    ///         var test = new Aws.Dms.Certificate("test", new Aws.Dms.CertificateArgs
    ///         {
    ///             CertificateId = "test-dms-certificate-tf",
    ///             CertificatePem = "...",
    ///             Tags = 
    ///             {
    ///                 { "Name", "test" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Certificates can be imported using the `certificate_arn`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:dms/certificate:Certificate test arn:aws:dms:us-west-2:123456789:cert:xxxxxxxxxx
    /// ```
    /// </summary>
    [AwsResourceType("aws:dms/certificate:Certificate")]
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the certificate.
        /// </summary>
        [Output("certificateArn")]
        public Output<string> CertificateArn { get; private set; } = null!;

        /// <summary>
        /// The certificate identifier.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
        /// </summary>
        [Output("certificatePem")]
        public Output<string?> CertificatePem { get; private set; } = null!;

        /// <summary>
        /// The contents of the Oracle Wallet certificate for use with SSL, provided as a base64-encoded String. Either `certificate_pem` or `certificate_wallet` must be set.
        /// </summary>
        [Output("certificateWallet")]
        public Output<string?> CertificateWallet { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("aws:dms/certificate:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("aws:dms/certificate:Certificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate identifier.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        /// <summary>
        /// The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
        /// </summary>
        [Input("certificatePem")]
        public Input<string>? CertificatePem { get; set; }

        /// <summary>
        /// The contents of the Oracle Wallet certificate for use with SSL, provided as a base64-encoded String. Either `certificate_pem` or `certificate_wallet` must be set.
        /// </summary>
        [Input("certificateWallet")]
        public Input<string>? CertificateWallet { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CertificateArgs()
        {
        }
    }

    public sealed class CertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the certificate.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// The certificate identifier.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
        /// </summary>
        [Input("certificatePem")]
        public Input<string>? CertificatePem { get; set; }

        /// <summary>
        /// The contents of the Oracle Wallet certificate for use with SSL, provided as a base64-encoded String. Either `certificate_pem` or `certificate_wallet` must be set.
        /// </summary>
        [Input("certificateWallet")]
        public Input<string>? CertificateWallet { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CertificateState()
        {
        }
    }
}
