// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LightSail
{
    /// <summary>
    /// Creates a Lightsail load balancer Certificate resource.
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
    ///     var testLb = new Aws.LightSail.Lb("testLb", new()
    ///     {
    ///         HealthCheckPath = "/",
    ///         InstancePort = 80,
    ///         Tags = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     });
    /// 
    ///     var testLbCertificate = new Aws.LightSail.LbCertificate("testLbCertificate", new()
    ///     {
    ///         LbName = testLb.Id,
    ///         DomainName = "test.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_lightsail_lb_certificate.test
    /// 
    ///  id = "example-load-balancer,example-load-balancer-certificate" } Using `pulumi import`, import `aws_lightsail_lb_certificate` using the id attribute. For exampleconsole % pulumi import aws_lightsail_lb_certificate.test example-load-balancer,example-load-balancer-certificate
    /// </summary>
    [AwsResourceType("aws:lightsail/lbCertificate:LbCertificate")]
    public partial class LbCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the lightsail certificate.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the instance was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The domain name (e.g., example.com) for your SSL/TLS certificate.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        [Output("domainValidationRecords")]
        public Output<ImmutableArray<Outputs.LbCertificateDomainValidationRecord>> DomainValidationRecords { get; private set; } = null!;

        /// <summary>
        /// The load balancer name where you want to create the SSL/TLS certificate.
        /// </summary>
        [Output("lbName")]
        public Output<string> LbName { get; private set; } = null!;

        /// <summary>
        /// The SSL/TLS certificate name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
        /// </summary>
        [Output("subjectAlternativeNames")]
        public Output<ImmutableArray<string>> SubjectAlternativeNames { get; private set; } = null!;

        [Output("supportCode")]
        public Output<string> SupportCode { get; private set; } = null!;


        /// <summary>
        /// Create a LbCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbCertificate(string name, LbCertificateArgs args, CustomResourceOptions? options = null)
            : base("aws:lightsail/lbCertificate:LbCertificate", name, args ?? new LbCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LbCertificate(string name, Input<string> id, LbCertificateState? state = null, CustomResourceOptions? options = null)
            : base("aws:lightsail/lbCertificate:LbCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LbCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbCertificate Get(string name, Input<string> id, LbCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new LbCertificate(name, id, state, options);
        }
    }

    public sealed class LbCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name (e.g., example.com) for your SSL/TLS certificate.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The load balancer name where you want to create the SSL/TLS certificate.
        /// </summary>
        [Input("lbName", required: true)]
        public Input<string> LbName { get; set; } = null!;

        /// <summary>
        /// The SSL/TLS certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;

        /// <summary>
        /// Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
        /// </summary>
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        public LbCertificateArgs()
        {
        }
        public static new LbCertificateArgs Empty => new LbCertificateArgs();
    }

    public sealed class LbCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the lightsail certificate.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The timestamp when the instance was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The domain name (e.g., example.com) for your SSL/TLS certificate.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("domainValidationRecords")]
        private InputList<Inputs.LbCertificateDomainValidationRecordGetArgs>? _domainValidationRecords;
        public InputList<Inputs.LbCertificateDomainValidationRecordGetArgs> DomainValidationRecords
        {
            get => _domainValidationRecords ?? (_domainValidationRecords = new InputList<Inputs.LbCertificateDomainValidationRecordGetArgs>());
            set => _domainValidationRecords = value;
        }

        /// <summary>
        /// The load balancer name where you want to create the SSL/TLS certificate.
        /// </summary>
        [Input("lbName")]
        public Input<string>? LbName { get; set; }

        /// <summary>
        /// The SSL/TLS certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;

        /// <summary>
        /// Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
        /// </summary>
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        [Input("supportCode")]
        public Input<string>? SupportCode { get; set; }

        public LbCertificateState()
        {
        }
        public static new LbCertificateState Empty => new LbCertificateState();
    }
}
