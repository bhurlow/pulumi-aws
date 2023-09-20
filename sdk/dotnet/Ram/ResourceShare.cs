// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ram
{
    /// <summary>
    /// Manages a Resource Access Manager (RAM) Resource Share. To associate principals with the share, see the `aws.ram.PrincipalAssociation` resource. To associate resources with the share, see the `aws.ram.ResourceAssociation` resource.
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
    ///     var example = new Aws.Ram.ResourceShare("example", new()
    ///     {
    ///         AllowExternalPrincipals = true,
    ///         Tags = 
    ///         {
    ///             { "Environment", "Production" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// In TODO v1.5.0 and later, use an `import` block to import resource shares using the `arn` of the resource share. For exampleterraform import {
    /// 
    ///  to = aws_ram_resource_share.example
    /// 
    ///  id = "arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12" } Using `TODO import`, import resource shares using the `arn` of the resource share. For exampleconsole % TODO import aws_ram_resource_share.example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12
    /// </summary>
    [AwsResourceType("aws:ram/resourceShare:ResourceShare")]
    public partial class ResourceShare : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether principals outside your organization can be associated with a resource share.
        /// </summary>
        [Output("allowExternalPrincipals")]
        public Output<bool?> AllowExternalPrincipals { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource share.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the resource share.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the Amazon Resource Names (ARNs) of the RAM permission to associate with the resource share. If you do not specify an ARN for the permission, RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
        /// </summary>
        [Output("permissionArns")]
        public Output<ImmutableArray<string>> PermissionArns { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource share. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceShare(string name, ResourceShareArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ram/resourceShare:ResourceShare", name, args ?? new ResourceShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceShare(string name, Input<string> id, ResourceShareState? state = null, CustomResourceOptions? options = null)
            : base("aws:ram/resourceShare:ResourceShare", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceShare Get(string name, Input<string> id, ResourceShareState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceShare(name, id, state, options);
        }
    }

    public sealed class ResourceShareArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether principals outside your organization can be associated with a resource share.
        /// </summary>
        [Input("allowExternalPrincipals")]
        public Input<bool>? AllowExternalPrincipals { get; set; }

        /// <summary>
        /// The name of the resource share.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissionArns")]
        private InputList<string>? _permissionArns;

        /// <summary>
        /// Specifies the Amazon Resource Names (ARNs) of the RAM permission to associate with the resource share. If you do not specify an ARN for the permission, RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
        /// </summary>
        public InputList<string> PermissionArns
        {
            get => _permissionArns ?? (_permissionArns = new InputList<string>());
            set => _permissionArns = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource share. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ResourceShareArgs()
        {
        }
        public static new ResourceShareArgs Empty => new ResourceShareArgs();
    }

    public sealed class ResourceShareState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether principals outside your organization can be associated with a resource share.
        /// </summary>
        [Input("allowExternalPrincipals")]
        public Input<bool>? AllowExternalPrincipals { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource share.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the resource share.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissionArns")]
        private InputList<string>? _permissionArns;

        /// <summary>
        /// Specifies the Amazon Resource Names (ARNs) of the RAM permission to associate with the resource share. If you do not specify an ARN for the permission, RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
        /// </summary>
        public InputList<string> PermissionArns
        {
            get => _permissionArns ?? (_permissionArns = new InputList<string>());
            set => _permissionArns = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource share. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public ResourceShareState()
        {
        }
        public static new ResourceShareState Empty => new ResourceShareState();
    }
}
