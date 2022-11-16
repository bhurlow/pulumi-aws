// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaLive
{
    /// <summary>
    /// Resource for managing an AWS MediaLive Input.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleInputSecurityGroup = new Aws.MediaLive.InputSecurityGroup("exampleInputSecurityGroup", new()
    ///     {
    ///         WhitelistRules = new[]
    ///         {
    ///             new Aws.MediaLive.Inputs.InputSecurityGroupWhitelistRuleArgs
    ///             {
    ///                 Cidr = "10.0.0.8/32",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "ENVIRONMENT", "prod" },
    ///         },
    ///     });
    /// 
    ///     var exampleInput = new Aws.MediaLive.Input("exampleInput", new()
    ///     {
    ///         InputSecurityGroups = new[]
    ///         {
    ///             exampleInputSecurityGroup.Id,
    ///         },
    ///         Type = "UDP_PUSH",
    ///         Tags = 
    ///         {
    ///             { "ENVIRONMENT", "prod" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// MediaLive Input can be imported using the `id`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:medialive/input:Input example 12345678
    /// ```
    /// </summary>
    [AwsResourceType("aws:medialive/input:Input")]
    public partial class Input : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the Input.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Channels attached to Input.
        /// </summary>
        [Output("attachedChannels")]
        public Output<ImmutableArray<string>> AttachedChannels { get; private set; } = null!;

        /// <summary>
        /// Destination settings for PUSH type inputs. See Destinations for more details.
        /// </summary>
        [Output("destinations")]
        public Output<ImmutableArray<Outputs.InputDestination>> Destinations { get; private set; } = null!;

        /// <summary>
        /// The input class.
        /// </summary>
        [Output("inputClass")]
        public Output<string> InputClass { get; private set; } = null!;

        /// <summary>
        /// Settings for the devices. See Input Devices for more details.
        /// </summary>
        [Output("inputDevices")]
        public Output<ImmutableArray<Outputs.InputInputDevice>> InputDevices { get; private set; } = null!;

        /// <summary>
        /// A list of IDs for all Inputs which are partners of this one.
        /// </summary>
        [Output("inputPartnerIds")]
        public Output<ImmutableArray<string>> InputPartnerIds { get; private set; } = null!;

        /// <summary>
        /// List of input security groups.
        /// </summary>
        [Output("inputSecurityGroups")]
        public Output<ImmutableArray<string>> InputSecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Source type of the input.
        /// </summary>
        [Output("inputSourceType")]
        public Output<string> InputSourceType { get; private set; } = null!;

        /// <summary>
        /// A list of the MediaConnect Flows. See Media Connect Flows for more details.
        /// </summary>
        [Output("mediaConnectFlows")]
        public Output<ImmutableArray<Outputs.InputMediaConnectFlow>> MediaConnectFlows { get; private set; } = null!;

        /// <summary>
        /// Name of the input.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the role this input assumes during and after creation.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The source URLs for a PULL-type input. See Sources for more details.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.InputSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the Input. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The different types of inputs that AWS Elemental MediaLive supports.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Settings for a private VPC Input. See VPC for more details.
        /// </summary>
        [Output("vpc")]
        public Output<Outputs.InputVpc?> Vpc { get; private set; } = null!;


        /// <summary>
        /// Create a Input resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Input(string name, InputArgs args, CustomResourceOptions? options = null)
            : base("aws:medialive/input:Input", name, args ?? new InputArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Input(string name, Input<string> id, InputState? state = null, CustomResourceOptions? options = null)
            : base("aws:medialive/input:Input", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Input resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Input Get(string name, Input<string> id, InputState? state = null, CustomResourceOptions? options = null)
        {
            return new Input(name, id, state, options);
        }
    }

    public sealed class InputArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<Inputs.InputDestinationArgs>? _destinations;

        /// <summary>
        /// Destination settings for PUSH type inputs. See Destinations for more details.
        /// </summary>
        public InputList<Inputs.InputDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.InputDestinationArgs>());
            set => _destinations = value;
        }

        [Input("inputDevices")]
        private InputList<Inputs.InputInputDeviceArgs>? _inputDevices;

        /// <summary>
        /// Settings for the devices. See Input Devices for more details.
        /// </summary>
        public InputList<Inputs.InputInputDeviceArgs> InputDevices
        {
            get => _inputDevices ?? (_inputDevices = new InputList<Inputs.InputInputDeviceArgs>());
            set => _inputDevices = value;
        }

        [Input("inputSecurityGroups")]
        private InputList<string>? _inputSecurityGroups;

        /// <summary>
        /// List of input security groups.
        /// </summary>
        public InputList<string> InputSecurityGroups
        {
            get => _inputSecurityGroups ?? (_inputSecurityGroups = new InputList<string>());
            set => _inputSecurityGroups = value;
        }

        [Input("mediaConnectFlows")]
        private InputList<Inputs.InputMediaConnectFlowArgs>? _mediaConnectFlows;

        /// <summary>
        /// A list of the MediaConnect Flows. See Media Connect Flows for more details.
        /// </summary>
        public InputList<Inputs.InputMediaConnectFlowArgs> MediaConnectFlows
        {
            get => _mediaConnectFlows ?? (_mediaConnectFlows = new InputList<Inputs.InputMediaConnectFlowArgs>());
            set => _mediaConnectFlows = value;
        }

        /// <summary>
        /// Name of the input.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the role this input assumes during and after creation.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("sources")]
        private InputList<Inputs.InputSourceArgs>? _sources;

        /// <summary>
        /// The source URLs for a PULL-type input. See Sources for more details.
        /// </summary>
        public InputList<Inputs.InputSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.InputSourceArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the Input. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The different types of inputs that AWS Elemental MediaLive supports.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Settings for a private VPC Input. See VPC for more details.
        /// </summary>
        [Input("vpc")]
        public Input<Inputs.InputVpcArgs>? Vpc { get; set; }

        public InputArgs()
        {
        }
        public static new InputArgs Empty => new InputArgs();
    }

    public sealed class InputState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the Input.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("attachedChannels")]
        private InputList<string>? _attachedChannels;

        /// <summary>
        /// Channels attached to Input.
        /// </summary>
        public InputList<string> AttachedChannels
        {
            get => _attachedChannels ?? (_attachedChannels = new InputList<string>());
            set => _attachedChannels = value;
        }

        [Input("destinations")]
        private InputList<Inputs.InputDestinationGetArgs>? _destinations;

        /// <summary>
        /// Destination settings for PUSH type inputs. See Destinations for more details.
        /// </summary>
        public InputList<Inputs.InputDestinationGetArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.InputDestinationGetArgs>());
            set => _destinations = value;
        }

        /// <summary>
        /// The input class.
        /// </summary>
        [Input("inputClass")]
        public Input<string>? InputClass { get; set; }

        [Input("inputDevices")]
        private InputList<Inputs.InputInputDeviceGetArgs>? _inputDevices;

        /// <summary>
        /// Settings for the devices. See Input Devices for more details.
        /// </summary>
        public InputList<Inputs.InputInputDeviceGetArgs> InputDevices
        {
            get => _inputDevices ?? (_inputDevices = new InputList<Inputs.InputInputDeviceGetArgs>());
            set => _inputDevices = value;
        }

        [Input("inputPartnerIds")]
        private InputList<string>? _inputPartnerIds;

        /// <summary>
        /// A list of IDs for all Inputs which are partners of this one.
        /// </summary>
        public InputList<string> InputPartnerIds
        {
            get => _inputPartnerIds ?? (_inputPartnerIds = new InputList<string>());
            set => _inputPartnerIds = value;
        }

        [Input("inputSecurityGroups")]
        private InputList<string>? _inputSecurityGroups;

        /// <summary>
        /// List of input security groups.
        /// </summary>
        public InputList<string> InputSecurityGroups
        {
            get => _inputSecurityGroups ?? (_inputSecurityGroups = new InputList<string>());
            set => _inputSecurityGroups = value;
        }

        /// <summary>
        /// Source type of the input.
        /// </summary>
        [Input("inputSourceType")]
        public Input<string>? InputSourceType { get; set; }

        [Input("mediaConnectFlows")]
        private InputList<Inputs.InputMediaConnectFlowGetArgs>? _mediaConnectFlows;

        /// <summary>
        /// A list of the MediaConnect Flows. See Media Connect Flows for more details.
        /// </summary>
        public InputList<Inputs.InputMediaConnectFlowGetArgs> MediaConnectFlows
        {
            get => _mediaConnectFlows ?? (_mediaConnectFlows = new InputList<Inputs.InputMediaConnectFlowGetArgs>());
            set => _mediaConnectFlows = value;
        }

        /// <summary>
        /// Name of the input.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the role this input assumes during and after creation.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("sources")]
        private InputList<Inputs.InputSourceGetArgs>? _sources;

        /// <summary>
        /// The source URLs for a PULL-type input. See Sources for more details.
        /// </summary>
        public InputList<Inputs.InputSourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.InputSourceGetArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the Input. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
        /// The different types of inputs that AWS Elemental MediaLive supports.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Settings for a private VPC Input. See VPC for more details.
        /// </summary>
        [Input("vpc")]
        public Input<Inputs.InputVpcGetArgs>? Vpc { get; set; }

        public InputState()
        {
        }
        public static new InputState Empty => new InputState();
    }
}
