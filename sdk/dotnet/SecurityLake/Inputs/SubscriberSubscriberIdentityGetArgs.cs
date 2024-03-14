// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityLake.Inputs
{

    public sealed class SubscriberSubscriberIdentityGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Regions where Security Lake is automatically enabled.
        /// </summary>
        [Input("externalId", required: true)]
        public Input<string> ExternalId { get; set; } = null!;

        /// <summary>
        /// Provides encryption details of Amazon Security Lake object.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        public SubscriberSubscriberIdentityGetArgs()
        {
        }
        public static new SubscriberSubscriberIdentityGetArgs Empty => new SubscriberSubscriberIdentityGetArgs();
    }
}
