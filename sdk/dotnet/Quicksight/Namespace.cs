// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Quicksight
{
    /// <summary>
    /// Resource for managing an AWS QuickSight Namespace.
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
    ///     var example = new Aws.Quicksight.Namespace("example", new()
    ///     {
    ///         NameSpace = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_quicksight_namespace.example
    /// 
    ///  id = "123456789012,example" } Using `pulumi import`, import QuickSight Namespace using the AWS account ID and namespace separated by commas (`,`). For exampleconsole % pulumi import aws_quicksight_namespace.example 123456789012,example
    /// </summary>
    [AwsResourceType("aws:quicksight/namespace:Namespace")]
    public partial class Namespace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the Namespace.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        /// <summary>
        /// Namespace AWS Region.
        /// </summary>
        [Output("capacityRegion")]
        public Output<string> CapacityRegion { get; private set; } = null!;

        /// <summary>
        /// Creation status of the namespace.
        /// </summary>
        [Output("creationStatus")]
        public Output<string> CreationStatus { get; private set; } = null!;

        /// <summary>
        /// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
        /// </summary>
        [Output("identityStore")]
        public Output<string> IdentityStore { get; private set; } = null!;

        /// <summary>
        /// Name of the namespace.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("namespace")]
        public Output<string> NameSpace { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        [Output("timeouts")]
        public Output<Outputs.NamespaceTimeouts?> Timeouts { get; private set; } = null!;


        /// <summary>
        /// Create a Namespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Namespace(string name, NamespaceArgs args, CustomResourceOptions? options = null)
            : base("aws:quicksight/namespace:Namespace", name, args ?? new NamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Namespace(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:quicksight/namespace:Namespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Namespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Namespace Get(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Namespace(name, id, state, options);
        }
    }

    public sealed class NamespaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
        /// </summary>
        [Input("identityStore")]
        public Input<string>? IdentityStore { get; set; }

        /// <summary>
        /// Name of the namespace.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> NameSpace { get; set; } = null!;

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

        [Input("timeouts")]
        public Input<Inputs.NamespaceTimeoutsArgs>? Timeouts { get; set; }

        public NamespaceArgs()
        {
        }
        public static new NamespaceArgs Empty => new NamespaceArgs();
    }

    public sealed class NamespaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the Namespace.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// AWS account ID.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        /// <summary>
        /// Namespace AWS Region.
        /// </summary>
        [Input("capacityRegion")]
        public Input<string>? CapacityRegion { get; set; }

        /// <summary>
        /// Creation status of the namespace.
        /// </summary>
        [Input("creationStatus")]
        public Input<string>? CreationStatus { get; set; }

        /// <summary>
        /// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
        /// </summary>
        [Input("identityStore")]
        public Input<string>? IdentityStore { get; set; }

        /// <summary>
        /// Name of the namespace.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("namespace")]
        public Input<string>? NameSpace { get; set; }

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
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("timeouts")]
        public Input<Inputs.NamespaceTimeoutsGetArgs>? Timeouts { get; set; }

        public NamespaceState()
        {
        }
        public static new NamespaceState Empty => new NamespaceState();
    }
}
