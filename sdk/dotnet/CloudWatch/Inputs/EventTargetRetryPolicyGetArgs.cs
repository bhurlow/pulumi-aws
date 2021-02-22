// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Inputs
{

    public sealed class EventTargetRetryPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The age in seconds to continue to make retry attempts.
        /// </summary>
        [Input("maximumEventAgeInSeconds")]
        public Input<int>? MaximumEventAgeInSeconds { get; set; }

        /// <summary>
        /// maximum number of retry attempts to make before the request fails
        /// </summary>
        [Input("maximumRetryAttempts")]
        public Input<int>? MaximumRetryAttempts { get; set; }

        public EventTargetRetryPolicyGetArgs()
        {
        }
    }
}
