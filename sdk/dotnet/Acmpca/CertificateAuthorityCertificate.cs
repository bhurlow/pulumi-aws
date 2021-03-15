// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Acmpca
{
    /// <summary>
    /// Associates a certificate with an AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority). An ACM PCA Certificate Authority is unable to issue certificates until it has a certificate associated with it. A root level ACM PCA Certificate Authority is able to self-sign its own root certificate.
    /// 
    /// ## Example Usage
    /// ### Self-Signed Root Certificate Authority Certificate
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleCertificateAuthority = new Aws.Acmpca.CertificateAuthority("exampleCertificateAuthority", new Aws.Acmpca.CertificateAuthorityArgs
    ///         {
    ///             Type = "ROOT",
    ///             CertificateAuthorityConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs
    ///             {
    ///                 KeyAlgorithm = "RSA_4096",
    ///                 SigningAlgorithm = "SHA512WITHRSA",
    ///                 Subject = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs
    ///                 {
    ///                     CommonName = "example.com",
    ///                 },
    ///             },
    ///         });
    ///         var current = Output.Create(Aws.GetPartition.InvokeAsync());
    ///         var exampleCertificate = new Aws.Acmpca.Certificate("exampleCertificate", new Aws.Acmpca.CertificateArgs
    ///         {
    ///             CertificateAuthorityArn = exampleCertificateAuthority.Arn,
    ///             CertificateSigningRequest = exampleCertificateAuthority.CertificateSigningRequest,
    ///             SigningAlgorithm = "SHA512WITHRSA",
    ///             TemplateArn = current.Apply(current =&gt; $"arn:{current.Partition}:acm-pca:::template/RootCACertificate/V1"),
    ///             Validity = new Aws.Acmpca.Inputs.CertificateValidityArgs
    ///             {
    ///                 Type = "YEARS",
    ///                 Value = "1",
    ///             },
    ///         });
    ///         var exampleCertificateAuthorityCertificate = new Aws.Acmpca.CertificateAuthorityCertificate("exampleCertificateAuthorityCertificate", new Aws.Acmpca.CertificateAuthorityCertificateArgs
    ///         {
    ///             CertificateAuthorityArn = exampleCertificateAuthority.Arn,
    ///             Certificate = exampleCertificate.Certificate,
    ///             CertificateChain = exampleCertificate.CertificateChain,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Certificate for Subordinate Certificate Authority
    /// 
    /// Note that the certificate for the subordinate certificate authority must be issued by the root certificate authority using a signing request from the subordinate certificate authority.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var subordinateCertificateAuthority = new Aws.Acmpca.CertificateAuthority("subordinateCertificateAuthority", new Aws.Acmpca.CertificateAuthorityArgs
    ///         {
    ///             Type = "SUBORDINATE",
    ///             CertificateAuthorityConfiguration = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs
    ///             {
    ///                 KeyAlgorithm = "RSA_2048",
    ///                 SigningAlgorithm = "SHA512WITHRSA",
    ///                 Subject = new Aws.Acmpca.Inputs.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs
    ///                 {
    ///                     CommonName = "sub.example.com",
    ///                 },
    ///             },
    ///         });
    ///         var rootCertificateAuthority = new Aws.Acmpca.CertificateAuthority("rootCertificateAuthority", new Aws.Acmpca.CertificateAuthorityArgs
    ///         {
    ///         });
    ///         // ...
    ///         var current = Output.Create(Aws.GetPartition.InvokeAsync());
    ///         var subordinateCertificate = new Aws.Acmpca.Certificate("subordinateCertificate", new Aws.Acmpca.CertificateArgs
    ///         {
    ///             CertificateAuthorityArn = rootCertificateAuthority.Arn,
    ///             CertificateSigningRequest = subordinateCertificateAuthority.CertificateSigningRequest,
    ///             SigningAlgorithm = "SHA512WITHRSA",
    ///             TemplateArn = current.Apply(current =&gt; $"arn:{current.Partition}:acm-pca:::template/SubordinateCACertificate_PathLen0/V1"),
    ///             Validity = new Aws.Acmpca.Inputs.CertificateValidityArgs
    ///             {
    ///                 Type = "YEARS",
    ///                 Value = "1",
    ///             },
    ///         });
    ///         var subordinateCertificateAuthorityCertificate = new Aws.Acmpca.CertificateAuthorityCertificate("subordinateCertificateAuthorityCertificate", new Aws.Acmpca.CertificateAuthorityCertificateArgs
    ///         {
    ///             CertificateAuthorityArn = subordinateCertificateAuthority.Arn,
    ///             Certificate = subordinateCertificate.Certificate,
    ///             CertificateChain = subordinateCertificate.CertificateChain,
    ///         });
    ///         var rootCertificateAuthorityCertificate = new Aws.Acmpca.CertificateAuthorityCertificate("rootCertificateAuthorityCertificate", new Aws.Acmpca.CertificateAuthorityCertificateArgs
    ///         {
    ///         });
    ///         // ...
    ///         var rootCertificate = new Aws.Acmpca.Certificate("rootCertificate", new Aws.Acmpca.CertificateArgs
    ///         {
    ///         });
    ///         // ...
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AwsResourceType("aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate")]
    public partial class CertificateAuthorityCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The PEM-encoded certificate for the Certificate Authority.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Certificate Authority.
        /// </summary>
        [Output("certificateAuthorityArn")]
        public Output<string> CertificateAuthorityArn { get; private set; } = null!;

        /// <summary>
        /// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateAuthorityCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateAuthorityCertificate(string name, CertificateAuthorityCertificateArgs args, CustomResourceOptions? options = null)
            : base("aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate", name, args ?? new CertificateAuthorityCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateAuthorityCertificate(string name, Input<string> id, CertificateAuthorityCertificateState? state = null, CustomResourceOptions? options = null)
            : base("aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CertificateAuthorityCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateAuthorityCertificate Get(string name, Input<string> id, CertificateAuthorityCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateAuthorityCertificate(name, id, state, options);
        }
    }

    public sealed class CertificateAuthorityCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PEM-encoded certificate for the Certificate Authority.
        /// </summary>
        [Input("certificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Certificate Authority.
        /// </summary>
        [Input("certificateAuthorityArn", required: true)]
        public Input<string> CertificateAuthorityArn { get; set; } = null!;

        /// <summary>
        /// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        public CertificateAuthorityCertificateArgs()
        {
        }
    }

    public sealed class CertificateAuthorityCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PEM-encoded certificate for the Certificate Authority.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the Certificate Authority.
        /// </summary>
        [Input("certificateAuthorityArn")]
        public Input<string>? CertificateAuthorityArn { get; set; }

        /// <summary>
        /// The PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        public CertificateAuthorityCertificateState()
        {
        }
    }
}
