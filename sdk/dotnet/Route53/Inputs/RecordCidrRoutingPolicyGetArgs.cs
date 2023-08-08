// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53.Inputs
{

    public sealed class RecordCidrRoutingPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR collection ID. See the `aws.route53.CidrCollection` resource for more details.
        /// </summary>
        [Input("collectionId", required: true)]
        public Input<string> CollectionId { get; set; } = null!;

        /// <summary>
        /// The CIDR collection location name. See the `aws.route53.CidrLocation` resource for more details. A `location_name` with an asterisk `"*"` can be used to create a default CIDR record. `collection_id` is still required for default record.
        /// </summary>
        [Input("locationName", required: true)]
        public Input<string> LocationName { get; set; } = null!;

        public RecordCidrRoutingPolicyGetArgs()
        {
        }
        public static new RecordCidrRoutingPolicyGetArgs Empty => new RecordCidrRoutingPolicyGetArgs();
    }
}
