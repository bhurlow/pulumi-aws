// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Inputs
{

    public sealed class ProviderAssumeRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The duration, between 15 minutes and 12 hours, of the role session. Valid time units are ns, us (or µs), ms, s, h, or m.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// A unique identifier that might be required when you assume a role in another account.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("policyArns")]
        private InputList<string>? _policyArns;

        /// <summary>
        /// Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
        /// </summary>
        public InputList<string> PolicyArns
        {
            get => _policyArns ?? (_policyArns = new InputList<string>());
            set => _policyArns = value;
        }

        /// <summary>
        /// Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// An identifier for the assumed role session.
        /// </summary>
        [Input("sessionName")]
        public Input<string>? SessionName { get; set; }

        /// <summary>
        /// Source identity specified by the principal assuming the role.
        /// </summary>
        [Input("sourceIdentity")]
        public Input<string>? SourceIdentity { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Assume role session tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("transitiveTagKeys")]
        private InputList<string>? _transitiveTagKeys;

        /// <summary>
        /// Assume role session tag keys to pass to any subsequent sessions.
        /// </summary>
        public InputList<string> TransitiveTagKeys
        {
            get => _transitiveTagKeys ?? (_transitiveTagKeys = new InputList<string>());
            set => _transitiveTagKeys = value;
        }

        public ProviderAssumeRoleArgs()
        {
        }
        public static new ProviderAssumeRoleArgs Empty => new ProviderAssumeRoleArgs();
    }
}
