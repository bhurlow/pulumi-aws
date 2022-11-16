// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RolesAnywhere
{
    /// <summary>
    /// Resource for managing a Roles Anywhere Profile.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testRole = new Aws.Iam.Role("testRole", new()
    ///     {
    ///         Path = "/",
    ///         AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2012-10-17",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "sts:AssumeRole",
    ///                         "sts:TagSession",
    ///                         "sts:SetSourceIdentity",
    ///                     },
    ///                     ["Principal"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["Service"] = "rolesanywhere.amazonaws.com",
    ///                     },
    ///                     ["Effect"] = "Allow",
    ///                     ["Sid"] = "",
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    ///     var testProfile = new Aws.RolesAnywhere.Profile("testProfile", new()
    ///     {
    ///         RoleArns = new[]
    ///         {
    ///             testRole.Arn,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_rolesanywhere_profile` can be imported using its `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:rolesanywhere/profile:Profile example db138a85-8925-4f9f-a409-08231233cacf
    /// ```
    /// </summary>
    [AwsResourceType("aws:rolesanywhere/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Profile
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The number of seconds the vended session credentials are valid for. Defaults to 3600.
        /// </summary>
        [Output("durationSeconds")]
        public Output<int> DurationSeconds { get; private set; } = null!;

        /// <summary>
        /// Whether or not the Profile is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// A list of managed policy ARNs that apply to the vended session credentials.
        /// </summary>
        [Output("managedPolicyArns")]
        public Output<ImmutableArray<string>> ManagedPolicyArns { get; private set; } = null!;

        /// <summary>
        /// The name of the Profile.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        /// </summary>
        [Output("requireInstanceProperties")]
        public Output<bool?> RequireInstanceProperties { get; private set; } = null!;

        /// <summary>
        /// A list of IAM roles that this profile can assume
        /// </summary>
        [Output("roleArns")]
        public Output<ImmutableArray<string>> RoleArns { get; private set; } = null!;

        /// <summary>
        /// A session policy that applies to the trust boundary of the vended session credentials.
        /// </summary>
        [Output("sessionPolicy")]
        public Output<string?> SessionPolicy { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs args, CustomResourceOptions? options = null)
            : base("aws:rolesanywhere/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("aws:rolesanywhere/profile:Profile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds the vended session credentials are valid for. Defaults to 3600.
        /// </summary>
        [Input("durationSeconds")]
        public Input<int>? DurationSeconds { get; set; }

        /// <summary>
        /// Whether or not the Profile is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("managedPolicyArns")]
        private InputList<string>? _managedPolicyArns;

        /// <summary>
        /// A list of managed policy ARNs that apply to the vended session credentials.
        /// </summary>
        public InputList<string> ManagedPolicyArns
        {
            get => _managedPolicyArns ?? (_managedPolicyArns = new InputList<string>());
            set => _managedPolicyArns = value;
        }

        /// <summary>
        /// The name of the Profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        /// </summary>
        [Input("requireInstanceProperties")]
        public Input<bool>? RequireInstanceProperties { get; set; }

        [Input("roleArns", required: true)]
        private InputList<string>? _roleArns;

        /// <summary>
        /// A list of IAM roles that this profile can assume
        /// </summary>
        public InputList<string> RoleArns
        {
            get => _roleArns ?? (_roleArns = new InputList<string>());
            set => _roleArns = value;
        }

        /// <summary>
        /// A session policy that applies to the trust boundary of the vended session credentials.
        /// </summary>
        [Input("sessionPolicy")]
        public Input<string>? SessionPolicy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Profile
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The number of seconds the vended session credentials are valid for. Defaults to 3600.
        /// </summary>
        [Input("durationSeconds")]
        public Input<int>? DurationSeconds { get; set; }

        /// <summary>
        /// Whether or not the Profile is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("managedPolicyArns")]
        private InputList<string>? _managedPolicyArns;

        /// <summary>
        /// A list of managed policy ARNs that apply to the vended session credentials.
        /// </summary>
        public InputList<string> ManagedPolicyArns
        {
            get => _managedPolicyArns ?? (_managedPolicyArns = new InputList<string>());
            set => _managedPolicyArns = value;
        }

        /// <summary>
        /// The name of the Profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
        /// </summary>
        [Input("requireInstanceProperties")]
        public Input<bool>? RequireInstanceProperties { get; set; }

        [Input("roleArns")]
        private InputList<string>? _roleArns;

        /// <summary>
        /// A list of IAM roles that this profile can assume
        /// </summary>
        public InputList<string> RoleArns
        {
            get => _roleArns ?? (_roleArns = new InputList<string>());
            set => _roleArns = value;
        }

        /// <summary>
        /// A session policy that applies to the trust boundary of the vended session credentials.
        /// </summary>
        [Input("sessionPolicy")]
        public Input<string>? SessionPolicy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}
