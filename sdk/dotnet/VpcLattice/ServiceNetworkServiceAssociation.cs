// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.VpcLattice
{
    /// <summary>
    /// Resource for managing an AWS VPC Lattice Service Network Service Association.
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
    ///     var example = new Aws.VpcLattice.ServiceNetworkServiceAssociation("example", new()
    ///     {
    ///         ServiceIdentifier = aws_vpclattice_service.Example.Id,
    ///         ServiceNetworkIdentifier = aws_vpclattice_service_network.Example.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import VPC Lattice Service Network Service Association using the `id`. For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation example snsa-05e2474658a88f6ba
    /// ```
    /// </summary>
    [AwsResourceType("aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation")]
    public partial class ServiceNetworkServiceAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Association.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The account that created the association.
        /// </summary>
        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// The custom domain name of the service.
        /// </summary>
        [Output("customDomainName")]
        public Output<string> CustomDomainName { get; private set; } = null!;

        /// <summary>
        /// The DNS name of the service.
        /// </summary>
        [Output("dnsEntries")]
        public Output<ImmutableArray<Outputs.ServiceNetworkServiceAssociationDnsEntry>> DnsEntries { get; private set; } = null!;

        /// <summary>
        /// The ID or Amazon Resource Identifier (ARN) of the service.
        /// </summary>
        [Output("serviceIdentifier")]
        public Output<string> ServiceIdentifier { get; private set; } = null!;

        /// <summary>
        /// The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
        /// The following arguments are optional:
        /// </summary>
        [Output("serviceNetworkIdentifier")]
        public Output<string> ServiceNetworkIdentifier { get; private set; } = null!;

        /// <summary>
        /// The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceNetworkServiceAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceNetworkServiceAssociation(string name, ServiceNetworkServiceAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation", name, args ?? new ServiceNetworkServiceAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceNetworkServiceAssociation(string name, Input<string> id, ServiceNetworkServiceAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceNetworkServiceAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceNetworkServiceAssociation Get(string name, Input<string> id, ServiceNetworkServiceAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceNetworkServiceAssociation(name, id, state, options);
        }
    }

    public sealed class ServiceNetworkServiceAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID or Amazon Resource Identifier (ARN) of the service.
        /// </summary>
        [Input("serviceIdentifier", required: true)]
        public Input<string> ServiceIdentifier { get; set; } = null!;

        /// <summary>
        /// The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
        /// The following arguments are optional:
        /// </summary>
        [Input("serviceNetworkIdentifier", required: true)]
        public Input<string> ServiceNetworkIdentifier { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceNetworkServiceAssociationArgs()
        {
        }
        public static new ServiceNetworkServiceAssociationArgs Empty => new ServiceNetworkServiceAssociationArgs();
    }

    public sealed class ServiceNetworkServiceAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Association.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The account that created the association.
        /// </summary>
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// The custom domain name of the service.
        /// </summary>
        [Input("customDomainName")]
        public Input<string>? CustomDomainName { get; set; }

        [Input("dnsEntries")]
        private InputList<Inputs.ServiceNetworkServiceAssociationDnsEntryGetArgs>? _dnsEntries;

        /// <summary>
        /// The DNS name of the service.
        /// </summary>
        public InputList<Inputs.ServiceNetworkServiceAssociationDnsEntryGetArgs> DnsEntries
        {
            get => _dnsEntries ?? (_dnsEntries = new InputList<Inputs.ServiceNetworkServiceAssociationDnsEntryGetArgs>());
            set => _dnsEntries = value;
        }

        /// <summary>
        /// The ID or Amazon Resource Identifier (ARN) of the service.
        /// </summary>
        [Input("serviceIdentifier")]
        public Input<string>? ServiceIdentifier { get; set; }

        /// <summary>
        /// The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
        /// The following arguments are optional:
        /// </summary>
        [Input("serviceNetworkIdentifier")]
        public Input<string>? ServiceNetworkIdentifier { get; set; }

        /// <summary>
        /// The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public ServiceNetworkServiceAssociationState()
        {
        }
        public static new ServiceNetworkServiceAssociationState Empty => new ServiceNetworkServiceAssociationState();
    }
}
