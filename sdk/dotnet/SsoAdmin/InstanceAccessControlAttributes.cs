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
    /// Provides a Single Sign-On (SSO) ABAC Resource: https://docs.aws.amazon.com/singlesignon/latest/userguide/abac.html
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_ssoadmin_instance_access_control_attributes.example
    /// 
    ///  id = "arn:aws:sso:::instance/ssoins-0123456789abcdef" } Using `pulumi import`, import SSO Account Assignments using the `instance_arn`. For exampleconsole % pulumi import aws_ssoadmin_instance_access_control_attributes.example arn:aws:sso:::instance/ssoins-0123456789abcdef
    /// </summary>
    [AwsResourceType("aws:ssoadmin/instanceAccessControlAttributes:InstanceAccessControlAttributes")]
    public partial class InstanceAccessControlAttributes : global::Pulumi.CustomResource
    {
        /// <summary>
        /// See AccessControlAttribute for more details.
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<Outputs.InstanceAccessControlAttributesAttribute>> Attributes { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the SSO Instance.
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("statusReason")]
        public Output<string> StatusReason { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceAccessControlAttributes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceAccessControlAttributes(string name, InstanceAccessControlAttributesArgs args, CustomResourceOptions? options = null)
            : base("aws:ssoadmin/instanceAccessControlAttributes:InstanceAccessControlAttributes", name, args ?? new InstanceAccessControlAttributesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceAccessControlAttributes(string name, Input<string> id, InstanceAccessControlAttributesState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssoadmin/instanceAccessControlAttributes:InstanceAccessControlAttributes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceAccessControlAttributes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceAccessControlAttributes Get(string name, Input<string> id, InstanceAccessControlAttributesState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceAccessControlAttributes(name, id, state, options);
        }
    }

    public sealed class InstanceAccessControlAttributesArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes", required: true)]
        private InputList<Inputs.InstanceAccessControlAttributesAttributeArgs>? _attributes;

        /// <summary>
        /// See AccessControlAttribute for more details.
        /// </summary>
        public InputList<Inputs.InstanceAccessControlAttributesAttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.InstanceAccessControlAttributesAttributeArgs>());
            set => _attributes = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the SSO Instance.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        public InstanceAccessControlAttributesArgs()
        {
        }
        public static new InstanceAccessControlAttributesArgs Empty => new InstanceAccessControlAttributesArgs();
    }

    public sealed class InstanceAccessControlAttributesState : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<Inputs.InstanceAccessControlAttributesAttributeGetArgs>? _attributes;

        /// <summary>
        /// See AccessControlAttribute for more details.
        /// </summary>
        public InputList<Inputs.InstanceAccessControlAttributesAttributeGetArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.InstanceAccessControlAttributesAttributeGetArgs>());
            set => _attributes = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the SSO Instance.
        /// </summary>
        [Input("instanceArn")]
        public Input<string>? InstanceArn { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("statusReason")]
        public Input<string>? StatusReason { get; set; }

        public InstanceAccessControlAttributesState()
        {
        }
        public static new InstanceAccessControlAttributesState Empty => new InstanceAccessControlAttributesState();
    }
}
