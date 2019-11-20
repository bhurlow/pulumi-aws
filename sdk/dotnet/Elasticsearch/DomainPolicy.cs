// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch
{
    /// <summary>
    /// Allows setting policy to an Elasticsearch domain while referencing domain attributes (e.g. ARN)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticsearch_domain_policy.html.markdown.
    /// </summary>
    public partial class DomainPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// IAM policy document specifying the access policies for the domain
        /// </summary>
        [Output("accessPolicies")]
        public Output<string> AccessPolicies { get; private set; } = null!;

        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;


        /// <summary>
        /// Create a DomainPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainPolicy(string name, DomainPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticsearch/domainPolicy:DomainPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DomainPolicy(string name, Input<string> id, DomainPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticsearch/domainPolicy:DomainPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainPolicy Get(string name, Input<string> id, DomainPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainPolicy(name, id, state, options);
        }
    }

    public sealed class DomainPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAM policy document specifying the access policies for the domain
        /// </summary>
        [Input("accessPolicies", required: true)]
        public Input<string> AccessPolicies { get; set; } = null!;

        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        public DomainPolicyArgs()
        {
        }
    }

    public sealed class DomainPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAM policy document specifying the access policies for the domain
        /// </summary>
        [Input("accessPolicies")]
        public Input<string>? AccessPolicies { get; set; }

        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        public DomainPolicyState()
        {
        }
    }
}
