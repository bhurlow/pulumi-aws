// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.NetworkManager.Inputs
{

    public sealed class GetCoreNetworkPolicyDocumentAttachmentPolicyConditionInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// string value
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Valid values include: `equals`, `not-equals`, `contains`, `begins-with`.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// Valid values include: `account-id`, `any`, `tag-value`, `tag-exists`, `resource-id`, `region`, `attachment-type`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// string value
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GetCoreNetworkPolicyDocumentAttachmentPolicyConditionInputArgs()
        {
        }
    }
}
