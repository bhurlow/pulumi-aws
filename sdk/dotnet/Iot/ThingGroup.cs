// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot
{
    /// <summary>
    /// Manages an AWS IoT Thing Group.
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
    ///     var parent = new Aws.Iot.ThingGroup("parent");
    /// 
    ///     var example = new Aws.Iot.ThingGroup("example", new()
    ///     {
    ///         ParentGroupName = parent.Name,
    ///         Properties = new Aws.Iot.Inputs.ThingGroupPropertiesArgs
    ///         {
    ///             AttributePayload = new Aws.Iot.Inputs.ThingGroupPropertiesAttributePayloadArgs
    ///             {
    ///                 Attributes = 
    ///                 {
    ///                     { "One", "11111" },
    ///                     { "Two", "TwoTwo" },
    ///                 },
    ///             },
    ///             Description = "This is my thing group",
    ///         },
    ///         Tags = 
    ///         {
    ///             { "managed", "true" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IoT Things Groups can be imported using the name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:iot/thingGroup:ThingGroup example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:iot/thingGroup:ThingGroup")]
    public partial class ThingGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Thing Group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("metadatas")]
        public Output<ImmutableArray<Outputs.ThingGroupMetadata>> Metadatas { get; private set; } = null!;

        /// <summary>
        /// The name of the Thing Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the parent Thing Group.
        /// </summary>
        [Output("parentGroupName")]
        public Output<string?> ParentGroupName { get; private set; } = null!;

        /// <summary>
        /// The Thing Group properties. Defined below.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ThingGroupProperties?> Properties { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The current version of the Thing Group record in the registry.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ThingGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ThingGroup(string name, ThingGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:iot/thingGroup:ThingGroup", name, args ?? new ThingGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ThingGroup(string name, Input<string> id, ThingGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:iot/thingGroup:ThingGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ThingGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ThingGroup Get(string name, Input<string> id, ThingGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ThingGroup(name, id, state, options);
        }
    }

    public sealed class ThingGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Thing Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the parent Thing Group.
        /// </summary>
        [Input("parentGroupName")]
        public Input<string>? ParentGroupName { get; set; }

        /// <summary>
        /// The Thing Group properties. Defined below.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ThingGroupPropertiesArgs>? Properties { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ThingGroupArgs()
        {
        }
        public static new ThingGroupArgs Empty => new ThingGroupArgs();
    }

    public sealed class ThingGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Thing Group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.ThingGroupMetadataGetArgs>? _metadatas;
        public InputList<Inputs.ThingGroupMetadataGetArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.ThingGroupMetadataGetArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// The name of the Thing Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the parent Thing Group.
        /// </summary>
        [Input("parentGroupName")]
        public Input<string>? ParentGroupName { get; set; }

        /// <summary>
        /// The Thing Group properties. Defined below.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ThingGroupPropertiesGetArgs>? Properties { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The current version of the Thing Group record in the registry.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public ThingGroupState()
        {
        }
        public static new ThingGroupState Empty => new ThingGroupState();
    }
}
