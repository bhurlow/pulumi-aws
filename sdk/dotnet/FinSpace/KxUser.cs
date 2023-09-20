// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.FinSpace
{
    /// <summary>
    /// Resource for managing an AWS FinSpace Kx User.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleKey = new Aws.Kms.Key("exampleKey", new()
    ///     {
    ///         Description = "Example KMS Key",
    ///         DeletionWindowInDays = 7,
    ///     });
    /// 
    ///     var exampleKxEnvironment = new Aws.FinSpace.KxEnvironment("exampleKxEnvironment", new()
    ///     {
    ///         KmsKeyId = exampleKey.Arn,
    ///     });
    /// 
    ///     var exampleRole = new Aws.Iam.Role("exampleRole", new()
    ///     {
    ///         AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Version"] = "2012-10-17",
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Action"] = "sts:AssumeRole",
    ///                     ["Effect"] = "Allow",
    ///                     ["Sid"] = "",
    ///                     ["Principal"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["Service"] = "ec2.amazonaws.com",
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    ///     var exampleKxUser = new Aws.FinSpace.KxUser("exampleKxUser", new()
    ///     {
    ///         EnvironmentId = exampleKxEnvironment.Id,
    ///         IamRole = exampleRole.Arn,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Using `pulumi import`, import an AWS FinSpace Kx User using the `id` (environment ID and user name, comma-delimited). For example:
    /// 
    /// ```sh
    ///  $ pulumi import aws:finspace/kxUser:KxUser example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-user
    /// ```
    /// </summary>
    [AwsResourceType("aws:finspace/kxUser:KxUser")]
    public partial class KxUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) identifier of the KX user.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the KX environment.
        /// </summary>
        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// IAM role ARN to be associated with the user.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("iamRole")]
        public Output<string> IamRole { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Create a KxUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KxUser(string name, KxUserArgs args, CustomResourceOptions? options = null)
            : base("aws:finspace/kxUser:KxUser", name, args ?? new KxUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KxUser(string name, Input<string> id, KxUserState? state = null, CustomResourceOptions? options = null)
            : base("aws:finspace/kxUser:KxUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KxUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KxUser Get(string name, Input<string> id, KxUserState? state = null, CustomResourceOptions? options = null)
        {
            return new KxUser(name, id, state, options);
        }
    }

    public sealed class KxUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for the KX environment.
        /// </summary>
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// IAM role ARN to be associated with the user.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("iamRole", required: true)]
        public Input<string> IamRole { get; set; } = null!;

        /// <summary>
        /// A unique identifier for the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public KxUserArgs()
        {
        }
        public static new KxUserArgs Empty => new KxUserArgs();
    }

    public sealed class KxUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) identifier of the KX user.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Unique identifier for the KX environment.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// IAM role ARN to be associated with the user.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// A unique identifier for the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public KxUserState()
        {
        }
        public static new KxUserState Empty => new KxUserState();
    }
}
