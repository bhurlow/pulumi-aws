// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    /// <summary>
    /// Manages an Amazon API Gateway Version 2 domain name.
    /// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
    /// 
    /// &gt; **Note:** This resource establishes ownership of and the TLS settings for
    /// a particular domain name. An API stage can be associated with the domain name using the `aws.apigatewayv2.ApiMapping` resource.
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.ApiGatewayV2.DomainName("example", new()
    ///     {
    ///         Domain = "ws-api.example.com",
    ///         DomainNameConfiguration = new Aws.ApiGatewayV2.Inputs.DomainNameDomainNameConfigurationArgs
    ///         {
    ///             CertificateArn = aws_acm_certificate.Example.Arn,
    ///             EndpointType = "REGIONAL",
    ///             SecurityPolicy = "TLS_1_2",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Associated Route 53 Resource Record
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleDomainName = new Aws.ApiGatewayV2.DomainName("exampleDomainName", new()
    ///     {
    ///         Domain = "http-api.example.com",
    ///         DomainNameConfiguration = new Aws.ApiGatewayV2.Inputs.DomainNameDomainNameConfigurationArgs
    ///         {
    ///             CertificateArn = aws_acm_certificate.Example.Arn,
    ///             EndpointType = "REGIONAL",
    ///             SecurityPolicy = "TLS_1_2",
    ///         },
    ///     });
    /// 
    ///     var exampleRecord = new Aws.Route53.Record("exampleRecord", new()
    ///     {
    ///         Name = exampleDomainName.Domain,
    ///         Type = "A",
    ///         ZoneId = aws_route53_zone.Example.Zone_id,
    ///         Aliases = new[]
    ///         {
    ///             new Aws.Route53.Inputs.RecordAliasArgs
    ///             {
    ///                 Name = exampleDomainName.DomainNameConfiguration.Apply(domainNameConfiguration =&gt; domainNameConfiguration.TargetDomainName),
    ///                 ZoneId = exampleDomainName.DomainNameConfiguration.Apply(domainNameConfiguration =&gt; domainNameConfiguration.HostedZoneId),
    ///                 EvaluateTargetHealth = false,
    ///             },
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
    ///  to = aws_apigatewayv2_domain_name.example
    /// 
    ///  id = "ws-api.example.com" } Using `pulumi import`, import `aws_apigatewayv2_domain_name` using the domain name. For exampleconsole % pulumi import aws_apigatewayv2_domain_name.example ws-api.example.com
    /// </summary>
    [AwsResourceType("aws:apigatewayv2/domainName:DomainName")]
    public partial class DomainName : global::Pulumi.CustomResource
    {
        /// <summary>
        /// [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
        /// </summary>
        [Output("apiMappingSelectionExpression")]
        public Output<string> ApiMappingSelectionExpression { get; private set; } = null!;

        /// <summary>
        /// ARN of the domain name.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Domain name. Must be between 1 and 512 characters in length.
        /// </summary>
        [Output("domainName")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Domain name configuration. See below.
        /// </summary>
        [Output("domainNameConfiguration")]
        public Output<Outputs.DomainNameDomainNameConfiguration> DomainNameConfiguration { get; private set; } = null!;

        /// <summary>
        /// Mutual TLS authentication configuration for the domain name.
        /// </summary>
        [Output("mutualTlsAuthentication")]
        public Output<Outputs.DomainNameMutualTlsAuthentication?> MutualTlsAuthentication { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the domain name. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a DomainName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainName(string name, DomainNameArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/domainName:DomainName", name, args ?? new DomainNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainName(string name, Input<string> id, DomainNameState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/domainName:DomainName", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainName Get(string name, Input<string> id, DomainNameState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainName(name, id, state, options);
        }
    }

    public sealed class DomainNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain name. Must be between 1 and 512 characters in length.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Domain name configuration. See below.
        /// </summary>
        [Input("domainNameConfiguration", required: true)]
        public Input<Inputs.DomainNameDomainNameConfigurationArgs> DomainNameConfiguration { get; set; } = null!;

        /// <summary>
        /// Mutual TLS authentication configuration for the domain name.
        /// </summary>
        [Input("mutualTlsAuthentication")]
        public Input<Inputs.DomainNameMutualTlsAuthenticationArgs>? MutualTlsAuthentication { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the domain name. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DomainNameArgs()
        {
        }
        public static new DomainNameArgs Empty => new DomainNameArgs();
    }

    public sealed class DomainNameState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
        /// </summary>
        [Input("apiMappingSelectionExpression")]
        public Input<string>? ApiMappingSelectionExpression { get; set; }

        /// <summary>
        /// ARN of the domain name.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Domain name. Must be between 1 and 512 characters in length.
        /// </summary>
        [Input("domainName")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Domain name configuration. See below.
        /// </summary>
        [Input("domainNameConfiguration")]
        public Input<Inputs.DomainNameDomainNameConfigurationGetArgs>? DomainNameConfiguration { get; set; }

        /// <summary>
        /// Mutual TLS authentication configuration for the domain name.
        /// </summary>
        [Input("mutualTlsAuthentication")]
        public Input<Inputs.DomainNameMutualTlsAuthenticationGetArgs>? MutualTlsAuthentication { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the domain name. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public DomainNameState()
        {
        }
        public static new DomainNameState Empty => new DomainNameState();
    }
}
