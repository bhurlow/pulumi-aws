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
    /// Creates and manages an AWS IoT Thing Type.
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
    ///     var foo = new Aws.Iot.ThingType("foo");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_iot_thing_type.example
    /// 
    ///  id = "example" } Using `pulumi import`, import IOT Thing Types using the name. For exampleconsole % pulumi import aws_iot_thing_type.example example
    /// </summary>
    [AwsResourceType("aws:iot/thingType:ThingType")]
    public partial class ThingType : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the created AWS IoT Thing Type.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Whether the thing type is deprecated. If true, no new things could be associated with this type.
        /// </summary>
        [Output("deprecated")]
        public Output<bool?> Deprecated { get; private set; } = null!;

        /// <summary>
        /// The name of the thing type.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// , Configuration block that can contain the following properties of the thing type:
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ThingTypeProperties?> Properties { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ThingType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ThingType(string name, ThingTypeArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:iot/thingType:ThingType", name, args ?? new ThingTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ThingType(string name, Input<string> id, ThingTypeState? state = null, CustomResourceOptions? options = null)
            : base("aws:iot/thingType:ThingType", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ThingType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ThingType Get(string name, Input<string> id, ThingTypeState? state = null, CustomResourceOptions? options = null)
        {
            return new ThingType(name, id, state, options);
        }
    }

    public sealed class ThingTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the thing type is deprecated. If true, no new things could be associated with this type.
        /// </summary>
        [Input("deprecated")]
        public Input<bool>? Deprecated { get; set; }

        /// <summary>
        /// The name of the thing type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// , Configuration block that can contain the following properties of the thing type:
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ThingTypePropertiesArgs>? Properties { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ThingTypeArgs()
        {
        }
        public static new ThingTypeArgs Empty => new ThingTypeArgs();
    }

    public sealed class ThingTypeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the created AWS IoT Thing Type.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Whether the thing type is deprecated. If true, no new things could be associated with this type.
        /// </summary>
        [Input("deprecated")]
        public Input<bool>? Deprecated { get; set; }

        /// <summary>
        /// The name of the thing type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// , Configuration block that can contain the following properties of the thing type:
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ThingTypePropertiesGetArgs>? Properties { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
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

        public ThingTypeState()
        {
        }
        public static new ThingTypeState Empty => new ThingTypeState();
    }
}
