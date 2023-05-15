// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsoAdmin
{
    /// <summary>
    /// Provides a Single Sign-On (SSO) Permission Set resource
    /// 
    /// &gt; **NOTE:** Updating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.
    /// 
    /// ## Import
    /// 
    /// SSO Permission Sets can be imported using the `arn` and `instance_arn` separated by a comma (`,`) e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:ssoadmin/permissionSet:PermissionSet example arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
    /// ```
    /// </summary>
    [AwsResourceType("aws:ssoadmin/permissionSet:PermissionSet")]
    public partial class PermissionSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Permission Set.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The date the Permission Set was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// The description of the Permission Set.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;

        /// <summary>
        /// The name of the Permission Set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The relay state URL used to redirect users within the application during the federation authentication process.
        /// </summary>
        [Output("relayState")]
        public Output<string?> RelayState { get; private set; } = null!;

        /// <summary>
        /// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
        /// </summary>
        [Output("sessionDuration")]
        public Output<string?> SessionDuration { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a PermissionSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PermissionSet(string name, PermissionSetArgs args, CustomResourceOptions? options = null)
            : base("aws:ssoadmin/permissionSet:PermissionSet", name, args ?? new PermissionSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PermissionSet(string name, Input<string> id, PermissionSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssoadmin/permissionSet:PermissionSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PermissionSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PermissionSet Get(string name, Input<string> id, PermissionSetState? state = null, CustomResourceOptions? options = null)
        {
            return new PermissionSet(name, id, state, options);
        }
    }

    public sealed class PermissionSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Permission Set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        /// <summary>
        /// The name of the Permission Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The relay state URL used to redirect users within the application during the federation authentication process.
        /// </summary>
        [Input("relayState")]
        public Input<string>? RelayState { get; set; }

        /// <summary>
        /// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
        /// </summary>
        [Input("sessionDuration")]
        public Input<string>? SessionDuration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PermissionSetArgs()
        {
        }
        public static new PermissionSetArgs Empty => new PermissionSetArgs();
    }

    public sealed class PermissionSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Permission Set.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The date the Permission Set was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// The description of the Permission Set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
        /// </summary>
        [Input("instanceArn")]
        public Input<string>? InstanceArn { get; set; }

        /// <summary>
        /// The name of the Permission Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The relay state URL used to redirect users within the application during the federation authentication process.
        /// </summary>
        [Input("relayState")]
        public Input<string>? RelayState { get; set; }

        /// <summary>
        /// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
        /// </summary>
        [Input("sessionDuration")]
        public Input<string>? SessionDuration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public PermissionSetState()
        {
        }
        public static new PermissionSetState Empty => new PermissionSetState();
    }
}
