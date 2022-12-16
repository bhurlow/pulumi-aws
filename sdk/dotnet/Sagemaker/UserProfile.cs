// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker
{
    /// <summary>
    /// Provides a SageMaker User Profile resource.
    /// 
    /// ## Example Usage
    /// ### Basic usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Sagemaker.UserProfile("example", new()
    ///     {
    ///         DomainId = aws_sagemaker_domain.Test.Id,
    ///         UserProfileName = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SageMaker User Profiles can be imported using the `arn`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:sagemaker/userProfile:UserProfile test_user_profile arn:aws:sagemaker:us-west-2:123456789012:user-profile/domain-id/profile-name
    /// ```
    /// </summary>
    [AwsResourceType("aws:sagemaker/userProfile:UserProfile")]
    public partial class UserProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The user profile Amazon Resource Name (ARN).
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated Domain.
        /// </summary>
        [Output("domainId")]
        public Output<string> DomainId { get; private set; } = null!;

        /// <summary>
        /// The ID of the user's profile in the Amazon Elastic File System (EFS) volume.
        /// </summary>
        [Output("homeEfsFileSystemUid")]
        public Output<string> HomeEfsFileSystemUid { get; private set; } = null!;

        /// <summary>
        /// A specifier for the type of value specified in `single_sign_on_user_value`. Currently, the only supported value is `UserName`. If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
        /// </summary>
        [Output("singleSignOnUserIdentifier")]
        public Output<string?> SingleSignOnUserIdentifier { get; private set; } = null!;

        /// <summary>
        /// The username of the associated AWS Single Sign-On User for this User Profile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
        /// </summary>
        [Output("singleSignOnUserValue")]
        public Output<string?> SingleSignOnUserValue { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The name for the User Profile.
        /// </summary>
        [Output("userProfileName")]
        public Output<string> UserProfileName { get; private set; } = null!;

        /// <summary>
        /// The user settings. See User Settings below.
        /// </summary>
        [Output("userSettings")]
        public Output<Outputs.UserProfileUserSettings?> UserSettings { get; private set; } = null!;


        /// <summary>
        /// Create a UserProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserProfile(string name, UserProfileArgs args, CustomResourceOptions? options = null)
            : base("aws:sagemaker/userProfile:UserProfile", name, args ?? new UserProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserProfile(string name, Input<string> id, UserProfileState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/userProfile:UserProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserProfile Get(string name, Input<string> id, UserProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new UserProfile(name, id, state, options);
        }
    }

    public sealed class UserProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the associated Domain.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// A specifier for the type of value specified in `single_sign_on_user_value`. Currently, the only supported value is `UserName`. If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
        /// </summary>
        [Input("singleSignOnUserIdentifier")]
        public Input<string>? SingleSignOnUserIdentifier { get; set; }

        /// <summary>
        /// The username of the associated AWS Single Sign-On User for this User Profile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
        /// </summary>
        [Input("singleSignOnUserValue")]
        public Input<string>? SingleSignOnUserValue { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name for the User Profile.
        /// </summary>
        [Input("userProfileName", required: true)]
        public Input<string> UserProfileName { get; set; } = null!;

        /// <summary>
        /// The user settings. See User Settings below.
        /// </summary>
        [Input("userSettings")]
        public Input<Inputs.UserProfileUserSettingsArgs>? UserSettings { get; set; }

        public UserProfileArgs()
        {
        }
        public static new UserProfileArgs Empty => new UserProfileArgs();
    }

    public sealed class UserProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user profile Amazon Resource Name (ARN).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ID of the associated Domain.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The ID of the user's profile in the Amazon Elastic File System (EFS) volume.
        /// </summary>
        [Input("homeEfsFileSystemUid")]
        public Input<string>? HomeEfsFileSystemUid { get; set; }

        /// <summary>
        /// A specifier for the type of value specified in `single_sign_on_user_value`. Currently, the only supported value is `UserName`. If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
        /// </summary>
        [Input("singleSignOnUserIdentifier")]
        public Input<string>? SingleSignOnUserIdentifier { get; set; }

        /// <summary>
        /// The username of the associated AWS Single Sign-On User for this User Profile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
        /// </summary>
        [Input("singleSignOnUserValue")]
        public Input<string>? SingleSignOnUserValue { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        /// <summary>
        /// The name for the User Profile.
        /// </summary>
        [Input("userProfileName")]
        public Input<string>? UserProfileName { get; set; }

        /// <summary>
        /// The user settings. See User Settings below.
        /// </summary>
        [Input("userSettings")]
        public Input<Inputs.UserProfileUserSettingsGetArgs>? UserSettings { get; set; }

        public UserProfileState()
        {
        }
        public static new UserProfileState Empty => new UserProfileState();
    }
}
