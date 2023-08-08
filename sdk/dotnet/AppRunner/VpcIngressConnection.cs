// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppRunner
{
    /// <summary>
    /// Manages an App Runner VPC Ingress Connection.
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
    ///     var example = new Aws.AppRunner.VpcIngressConnection("example", new()
    ///     {
    ///         ServiceArn = aws_apprunner_service.Example.Arn,
    ///         IngressVpcConfiguration = new Aws.AppRunner.Inputs.VpcIngressConnectionIngressVpcConfigurationArgs
    ///         {
    ///             VpcId = aws_default_vpc.Default.Id,
    ///             VpcEndpointId = aws_vpc_endpoint.Apprunner.Id,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "foo", "bar" },
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
    ///  to = aws_apprunner_vpc_ingress_connection.example
    /// 
    ///  id = "arn:aws:apprunner:us-west-2:837424938642:vpcingressconnection/example/b379f86381d74825832c2e82080342fa" } Using `pulumi import`, import App Runner VPC Ingress Connection using the `arn`. For exampleconsole % pulumi import aws_apprunner_vpc_ingress_connection.example "arn:aws:apprunner:us-west-2:837424938642:vpcingressconnection/example/b379f86381d74825832c2e82080342fa"
    /// </summary>
    [AwsResourceType("aws:apprunner/vpcIngressConnection:VpcIngressConnection")]
    public partial class VpcIngressConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the VPC Ingress Connection.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The domain name associated with the VPC Ingress Connection resource.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
        /// </summary>
        [Output("ingressVpcConfiguration")]
        public Output<Outputs.VpcIngressConnectionIngressVpcConfiguration> IngressVpcConfiguration { get; private set; } = null!;

        /// <summary>
        /// A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
        /// </summary>
        [Output("serviceArn")]
        public Output<string> ServiceArn { get; private set; } = null!;

        /// <summary>
        /// The current status of the VPC Ingress Connection.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a VpcIngressConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcIngressConnection(string name, VpcIngressConnectionArgs args, CustomResourceOptions? options = null)
            : base("aws:apprunner/vpcIngressConnection:VpcIngressConnection", name, args ?? new VpcIngressConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcIngressConnection(string name, Input<string> id, VpcIngressConnectionState? state = null, CustomResourceOptions? options = null)
            : base("aws:apprunner/vpcIngressConnection:VpcIngressConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcIngressConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcIngressConnection Get(string name, Input<string> id, VpcIngressConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcIngressConnection(name, id, state, options);
        }
    }

    public sealed class VpcIngressConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
        /// </summary>
        [Input("ingressVpcConfiguration", required: true)]
        public Input<Inputs.VpcIngressConnectionIngressVpcConfigurationArgs> IngressVpcConfiguration { get; set; } = null!;

        /// <summary>
        /// A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
        /// </summary>
        [Input("serviceArn", required: true)]
        public Input<string> ServiceArn { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VpcIngressConnectionArgs()
        {
        }
        public static new VpcIngressConnectionArgs Empty => new VpcIngressConnectionArgs();
    }

    public sealed class VpcIngressConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the VPC Ingress Connection.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The domain name associated with the VPC Ingress Connection resource.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
        /// </summary>
        [Input("ingressVpcConfiguration")]
        public Input<Inputs.VpcIngressConnectionIngressVpcConfigurationGetArgs>? IngressVpcConfiguration { get; set; }

        /// <summary>
        /// A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
        /// </summary>
        [Input("serviceArn")]
        public Input<string>? ServiceArn { get; set; }

        /// <summary>
        /// The current status of the VPC Ingress Connection.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public VpcIngressConnectionState()
        {
        }
        public static new VpcIngressConnectionState Empty => new VpcIngressConnectionState();
    }
}
