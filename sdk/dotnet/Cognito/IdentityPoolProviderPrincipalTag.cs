// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito
{
    /// <summary>
    /// Provides an AWS Cognito Identity Principal Mapping.
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_cognito_identity_pool_provider_principal_tag.example
    /// 
    ///  id = "us-west-2_abc123:CorpAD" } Using `pulumi import`, import Cognito Identity Pool Roles Attachment using the Identity Pool ID and provider name. For exampleconsole % pulumi import aws_cognito_identity_pool_provider_principal_tag.example us-west-2_abc123:CorpAD
    /// </summary>
    [AwsResourceType("aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag")]
    public partial class IdentityPoolProviderPrincipalTag : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An identity pool ID.
        /// </summary>
        [Output("identityPoolId")]
        public Output<string> IdentityPoolId { get; private set; } = null!;

        /// <summary>
        /// The name of the identity provider.
        /// </summary>
        [Output("identityProviderName")]
        public Output<string> IdentityProviderName { get; private set; } = null!;

        /// <summary>
        /// String to string map of variables.
        /// </summary>
        [Output("principalTags")]
        public Output<ImmutableDictionary<string, string>?> PrincipalTags { get; private set; } = null!;

        /// <summary>
        /// use default (username and clientID) attribute mappings.
        /// </summary>
        [Output("useDefaults")]
        public Output<bool?> UseDefaults { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityPoolProviderPrincipalTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityPoolProviderPrincipalTag(string name, IdentityPoolProviderPrincipalTagArgs args, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag", name, args ?? new IdentityPoolProviderPrincipalTagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityPoolProviderPrincipalTag(string name, Input<string> id, IdentityPoolProviderPrincipalTagState? state = null, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityPoolProviderPrincipalTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityPoolProviderPrincipalTag Get(string name, Input<string> id, IdentityPoolProviderPrincipalTagState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityPoolProviderPrincipalTag(name, id, state, options);
        }
    }

    public sealed class IdentityPoolProviderPrincipalTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identity pool ID.
        /// </summary>
        [Input("identityPoolId", required: true)]
        public Input<string> IdentityPoolId { get; set; } = null!;

        /// <summary>
        /// The name of the identity provider.
        /// </summary>
        [Input("identityProviderName", required: true)]
        public Input<string> IdentityProviderName { get; set; } = null!;

        [Input("principalTags")]
        private InputMap<string>? _principalTags;

        /// <summary>
        /// String to string map of variables.
        /// </summary>
        public InputMap<string> PrincipalTags
        {
            get => _principalTags ?? (_principalTags = new InputMap<string>());
            set => _principalTags = value;
        }

        /// <summary>
        /// use default (username and clientID) attribute mappings.
        /// </summary>
        [Input("useDefaults")]
        public Input<bool>? UseDefaults { get; set; }

        public IdentityPoolProviderPrincipalTagArgs()
        {
        }
        public static new IdentityPoolProviderPrincipalTagArgs Empty => new IdentityPoolProviderPrincipalTagArgs();
    }

    public sealed class IdentityPoolProviderPrincipalTagState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identity pool ID.
        /// </summary>
        [Input("identityPoolId")]
        public Input<string>? IdentityPoolId { get; set; }

        /// <summary>
        /// The name of the identity provider.
        /// </summary>
        [Input("identityProviderName")]
        public Input<string>? IdentityProviderName { get; set; }

        [Input("principalTags")]
        private InputMap<string>? _principalTags;

        /// <summary>
        /// String to string map of variables.
        /// </summary>
        public InputMap<string> PrincipalTags
        {
            get => _principalTags ?? (_principalTags = new InputMap<string>());
            set => _principalTags = value;
        }

        /// <summary>
        /// use default (username and clientID) attribute mappings.
        /// </summary>
        [Input("useDefaults")]
        public Input<bool>? UseDefaults { get; set; }

        public IdentityPoolProviderPrincipalTagState()
        {
        }
        public static new IdentityPoolProviderPrincipalTagState Empty => new IdentityPoolProviderPrincipalTagState();
    }
}
