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
    /// Provides an AWS Cognito Identity Pool Roles Attachment.
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_cognito_identity_pool_roles_attachment.example
    /// 
    ///  id = "us-west-2:b64805ad-cb56-40ba-9ffc-f5d8207e6d42" } Using `pulumi import`, import Cognito Identity Pool Roles Attachment using the Identity Pool ID. For exampleconsole % pulumi import aws_cognito_identity_pool_roles_attachment.example us-west-2:b64805ad-cb56-40ba-9ffc-f5d8207e6d42
    /// </summary>
    [AwsResourceType("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment")]
    public partial class IdentityPoolRoleAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An identity pool ID in the format `REGION_GUID`.
        /// </summary>
        [Output("identityPoolId")]
        public Output<string> IdentityPoolId { get; private set; } = null!;

        /// <summary>
        /// A List of Role Mapping.
        /// </summary>
        [Output("roleMappings")]
        public Output<ImmutableArray<Outputs.IdentityPoolRoleAttachmentRoleMapping>> RoleMappings { get; private set; } = null!;

        /// <summary>
        /// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableDictionary<string, string>> Roles { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityPoolRoleAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityPoolRoleAttachment(string name, IdentityPoolRoleAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, args ?? new IdentityPoolRoleAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityPoolRoleAttachment(string name, Input<string> id, IdentityPoolRoleAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityPoolRoleAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityPoolRoleAttachment Get(string name, Input<string> id, IdentityPoolRoleAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityPoolRoleAttachment(name, id, state, options);
        }
    }

    public sealed class IdentityPoolRoleAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identity pool ID in the format `REGION_GUID`.
        /// </summary>
        [Input("identityPoolId", required: true)]
        public Input<string> IdentityPoolId { get; set; } = null!;

        [Input("roleMappings")]
        private InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingArgs>? _roleMappings;

        /// <summary>
        /// A List of Role Mapping.
        /// </summary>
        public InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingArgs> RoleMappings
        {
            get => _roleMappings ?? (_roleMappings = new InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingArgs>());
            set => _roleMappings = value;
        }

        [Input("roles", required: true)]
        private InputMap<string>? _roles;

        /// <summary>
        /// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
        /// </summary>
        public InputMap<string> Roles
        {
            get => _roles ?? (_roles = new InputMap<string>());
            set => _roles = value;
        }

        public IdentityPoolRoleAttachmentArgs()
        {
        }
        public static new IdentityPoolRoleAttachmentArgs Empty => new IdentityPoolRoleAttachmentArgs();
    }

    public sealed class IdentityPoolRoleAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identity pool ID in the format `REGION_GUID`.
        /// </summary>
        [Input("identityPoolId")]
        public Input<string>? IdentityPoolId { get; set; }

        [Input("roleMappings")]
        private InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingGetArgs>? _roleMappings;

        /// <summary>
        /// A List of Role Mapping.
        /// </summary>
        public InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingGetArgs> RoleMappings
        {
            get => _roleMappings ?? (_roleMappings = new InputList<Inputs.IdentityPoolRoleAttachmentRoleMappingGetArgs>());
            set => _roleMappings = value;
        }

        [Input("roles")]
        private InputMap<string>? _roles;

        /// <summary>
        /// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
        /// </summary>
        public InputMap<string> Roles
        {
            get => _roles ?? (_roles = new InputMap<string>());
            set => _roles = value;
        }

        public IdentityPoolRoleAttachmentState()
        {
        }
        public static new IdentityPoolRoleAttachmentState Empty => new IdentityPoolRoleAttachmentState();
    }
}
