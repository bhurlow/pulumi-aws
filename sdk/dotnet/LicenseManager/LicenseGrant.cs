// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LicenseManager
{
    /// <summary>
    /// Provides a License Manager grant. This allows for sharing licenses with other AWS accounts.
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_licensemanager_grant.test
    /// 
    ///  id = "arn:aws:license-manager::123456789011:grant:g-01d313393d9e443d8664cc054db1e089" } Using `pulumi import`, import `aws_licensemanager_grant` using the grant arn. For exampleconsole % pulumi import aws_licensemanager_grant.test arn:aws:license-manager::123456789011:grant:g-01d313393d9e443d8664cc054db1e089
    /// </summary>
    [AwsResourceType("aws:licensemanager/licenseGrant:LicenseGrant")]
    public partial class LicenseGrant : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of the allowed operations for the grant. This is a subset of the allowed operations on the license.
        /// </summary>
        [Output("allowedOperations")]
        public Output<ImmutableArray<string>> AllowedOperations { get; private set; } = null!;

        /// <summary>
        /// The grant ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The home region for the license.
        /// </summary>
        [Output("homeRegion")]
        public Output<string> HomeRegion { get; private set; } = null!;

        /// <summary>
        /// The ARN of the license to grant.
        /// </summary>
        [Output("licenseArn")]
        public Output<string> LicenseArn { get; private set; } = null!;

        /// <summary>
        /// The Name of the grant.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent ARN.
        /// </summary>
        [Output("parentArn")]
        public Output<string> ParentArn { get; private set; } = null!;

        /// <summary>
        /// The target account for the grant in the form of the ARN for an account principal of the root user.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The grant status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The grant version.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a LicenseGrant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LicenseGrant(string name, LicenseGrantArgs args, CustomResourceOptions? options = null)
            : base("aws:licensemanager/licenseGrant:LicenseGrant", name, args ?? new LicenseGrantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LicenseGrant(string name, Input<string> id, LicenseGrantState? state = null, CustomResourceOptions? options = null)
            : base("aws:licensemanager/licenseGrant:LicenseGrant", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LicenseGrant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LicenseGrant Get(string name, Input<string> id, LicenseGrantState? state = null, CustomResourceOptions? options = null)
        {
            return new LicenseGrant(name, id, state, options);
        }
    }

    public sealed class LicenseGrantArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedOperations", required: true)]
        private InputList<string>? _allowedOperations;

        /// <summary>
        /// A list of the allowed operations for the grant. This is a subset of the allowed operations on the license.
        /// </summary>
        public InputList<string> AllowedOperations
        {
            get => _allowedOperations ?? (_allowedOperations = new InputList<string>());
            set => _allowedOperations = value;
        }

        /// <summary>
        /// The ARN of the license to grant.
        /// </summary>
        [Input("licenseArn", required: true)]
        public Input<string> LicenseArn { get; set; } = null!;

        /// <summary>
        /// The Name of the grant.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The target account for the grant in the form of the ARN for an account principal of the root user.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        public LicenseGrantArgs()
        {
        }
        public static new LicenseGrantArgs Empty => new LicenseGrantArgs();
    }

    public sealed class LicenseGrantState : global::Pulumi.ResourceArgs
    {
        [Input("allowedOperations")]
        private InputList<string>? _allowedOperations;

        /// <summary>
        /// A list of the allowed operations for the grant. This is a subset of the allowed operations on the license.
        /// </summary>
        public InputList<string> AllowedOperations
        {
            get => _allowedOperations ?? (_allowedOperations = new InputList<string>());
            set => _allowedOperations = value;
        }

        /// <summary>
        /// The grant ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The home region for the license.
        /// </summary>
        [Input("homeRegion")]
        public Input<string>? HomeRegion { get; set; }

        /// <summary>
        /// The ARN of the license to grant.
        /// </summary>
        [Input("licenseArn")]
        public Input<string>? LicenseArn { get; set; }

        /// <summary>
        /// The Name of the grant.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent ARN.
        /// </summary>
        [Input("parentArn")]
        public Input<string>? ParentArn { get; set; }

        /// <summary>
        /// The target account for the grant in the form of the ARN for an account principal of the root user.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// The grant status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The grant version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public LicenseGrantState()
        {
        }
        public static new LicenseGrantState Empty => new LicenseGrantState();
    }
}
