// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Acm
{
    /// <summary>
    /// The ACM certificate resource allows requesting and management of certificates
    /// from the Amazon Certificate Manager.
    /// 
    /// It deals with requesting certificates and managing their attributes and life-cycle.
    /// This resource does not deal with validation of a certificate but can provide inputs
    /// for other resources implementing the validation. It does not wait for a certificate to be issued.
    /// Use a `aws.acm.CertificateValidation` resource for this.
    /// 
    /// Most commonly, this resource is used to together with `aws.route53.Record` and
    /// `aws.acm.CertificateValidation` to request a DNS validated certificate,
    /// deploy the required validation records and wait for validation to complete.
    /// 
    /// Domain validation through E-Mail is also supported but should be avoided as it requires a manual step outside
    /// of this provider.
    /// 
    /// It's recommended to specify `create_before_destroy = true` in a [lifecycle][1] block to replace a certificate
    /// which is currently in use (eg, by `aws.lb.Listener`).
    /// 
    /// ## options Configuration Block
    /// 
    /// Supported nested arguments for the `options` configuration block:
    /// 
    /// * `certificate_transparency_logging_preference` - (Optional) Specifies whether certificate details should be added to a certificate transparency log. Valid values are `ENABLED` or `DISABLED`. See https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency for more details.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acm_certificate.html.markdown.
    /// </summary>
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the certificate
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// ARN of an ACMPCA
        /// </summary>
        [Output("certificateAuthorityArn")]
        public Output<string?> CertificateAuthorityArn { get; private set; } = null!;

        /// <summary>
        /// The certificate's PEM-formatted public key
        /// </summary>
        [Output("certificateBody")]
        public Output<string?> CertificateBody { get; private set; } = null!;

        /// <summary>
        /// The certificate's PEM-formatted chain
        /// * Creating a private CA issued certificate
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// A domain name for which the certificate should be issued
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// A list of attributes to feed into other resources to complete certificate validation. Can have more than one element, e.g. if SANs are defined. Only set if `DNS`-validation was used.
        /// </summary>
        [Output("domainValidationOptions")]
        public Output<ImmutableArray<Outputs.CertificateDomainValidationOptions>> DomainValidationOptions { get; private set; } = null!;

        [Output("options")]
        public Output<Outputs.CertificateOptions?> Options { get; private set; } = null!;

        /// <summary>
        /// The certificate's PEM-formatted private key
        /// </summary>
        [Output("privateKey")]
        public Output<string?> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// A list of domains that should be SANs in the issued certificate
        /// </summary>
        [Output("subjectAlternativeNames")]
        public Output<ImmutableArray<string>> SubjectAlternativeNames { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A list of addresses that received a validation E-Mail. Only set if `EMAIL`-validation was used.
        /// </summary>
        [Output("validationEmails")]
        public Output<ImmutableArray<string>> ValidationEmails { get; private set; } = null!;

        /// <summary>
        /// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into state managed by this provider.
        /// * Importing an existing certificate
        /// </summary>
        [Output("validationMethod")]
        public Output<string> ValidationMethod { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:acm/certificate:Certificate", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("aws:acm/certificate:Certificate", name, state, MakeResourceOptions(options, id))
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
        /// ARN of an ACMPCA
        /// </summary>
        [Input("certificateAuthorityArn")]
        public Input<string>? CertificateAuthorityArn { get; set; }

        /// <summary>
        /// The certificate's PEM-formatted public key
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// The certificate's PEM-formatted chain
        /// * Creating a private CA issued certificate
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// A domain name for which the certificate should be issued
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("options")]
        public Input<Inputs.CertificateOptionsArgs>? Options { get; set; }

        /// <summary>
        /// The certificate's PEM-formatted private key
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;

        /// <summary>
        /// A list of domains that should be SANs in the issued certificate
        /// </summary>
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into state managed by this provider.
        /// * Importing an existing certificate
        /// </summary>
        [Input("validationMethod")]
        public Input<string>? ValidationMethod { get; set; }

        public CertificateArgs()
        {
        }
    }

    public sealed class CertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the certificate
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// ARN of an ACMPCA
        /// </summary>
        [Input("certificateAuthorityArn")]
        public Input<string>? CertificateAuthorityArn { get; set; }

        /// <summary>
        /// The certificate's PEM-formatted public key
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// The certificate's PEM-formatted chain
        /// * Creating a private CA issued certificate
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// A domain name for which the certificate should be issued
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("domainValidationOptions")]
        private InputList<Inputs.CertificateDomainValidationOptionsGetArgs>? _domainValidationOptions;

        /// <summary>
        /// A list of attributes to feed into other resources to complete certificate validation. Can have more than one element, e.g. if SANs are defined. Only set if `DNS`-validation was used.
        /// </summary>
        public InputList<Inputs.CertificateDomainValidationOptionsGetArgs> DomainValidationOptions
        {
            get => _domainValidationOptions ?? (_domainValidationOptions = new InputList<Inputs.CertificateDomainValidationOptionsGetArgs>());
            set => _domainValidationOptions = value;
        }

        [Input("options")]
        public Input<Inputs.CertificateOptionsGetArgs>? Options { get; set; }

        /// <summary>
        /// The certificate's PEM-formatted private key
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;

        /// <summary>
        /// A list of domains that should be SANs in the issued certificate
        /// </summary>
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("validationEmails")]
        private InputList<string>? _validationEmails;

        /// <summary>
        /// A list of addresses that received a validation E-Mail. Only set if `EMAIL`-validation was used.
        /// </summary>
        public InputList<string> ValidationEmails
        {
            get => _validationEmails ?? (_validationEmails = new InputList<string>());
            set => _validationEmails = value;
        }

        /// <summary>
        /// Which method to use for validation. `DNS` or `EMAIL` are valid, `NONE` can be used for certificates that were imported into ACM and then into state managed by this provider.
        /// * Importing an existing certificate
        /// </summary>
        [Input("validationMethod")]
        public Input<string>? ValidationMethod { get; set; }

        public CertificateState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class CertificateDomainValidationOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A domain name for which the certificate should be issued
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The name of the DNS record to create to validate the certificate
        /// </summary>
        [Input("resourceRecordName")]
        public Input<string>? ResourceRecordName { get; set; }

        /// <summary>
        /// The type of DNS record to create
        /// </summary>
        [Input("resourceRecordType")]
        public Input<string>? ResourceRecordType { get; set; }

        /// <summary>
        /// The value the DNS record needs to have
        /// </summary>
        [Input("resourceRecordValue")]
        public Input<string>? ResourceRecordValue { get; set; }

        public CertificateDomainValidationOptionsGetArgs()
        {
        }
    }

    public sealed class CertificateOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("certificateTransparencyLoggingPreference")]
        public Input<string>? CertificateTransparencyLoggingPreference { get; set; }

        public CertificateOptionsArgs()
        {
        }
    }

    public sealed class CertificateOptionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("certificateTransparencyLoggingPreference")]
        public Input<string>? CertificateTransparencyLoggingPreference { get; set; }

        public CertificateOptionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class CertificateDomainValidationOptions
    {
        /// <summary>
        /// A domain name for which the certificate should be issued
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The name of the DNS record to create to validate the certificate
        /// </summary>
        public readonly string ResourceRecordName;
        /// <summary>
        /// The type of DNS record to create
        /// </summary>
        public readonly string ResourceRecordType;
        /// <summary>
        /// The value the DNS record needs to have
        /// </summary>
        public readonly string ResourceRecordValue;

        [OutputConstructor]
        private CertificateDomainValidationOptions(
            string domainName,
            string resourceRecordName,
            string resourceRecordType,
            string resourceRecordValue)
        {
            DomainName = domainName;
            ResourceRecordName = resourceRecordName;
            ResourceRecordType = resourceRecordType;
            ResourceRecordValue = resourceRecordValue;
        }
    }

    [OutputType]
    public sealed class CertificateOptions
    {
        public readonly string? CertificateTransparencyLoggingPreference;

        [OutputConstructor]
        private CertificateOptions(string? certificateTransparencyLoggingPreference)
        {
            CertificateTransparencyLoggingPreference = certificateTransparencyLoggingPreference;
        }
    }
    }
}
