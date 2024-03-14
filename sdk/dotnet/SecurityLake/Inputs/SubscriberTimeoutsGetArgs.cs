// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityLake.Inputs
{

    public sealed class SubscriberTimeoutsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        /// </summary>
        [Input("create")]
        public Input<string>? Create { get; set; }

        /// <summary>
        /// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
        /// </summary>
        [Input("delete")]
        public Input<string>? Delete { get; set; }

        /// <summary>
        /// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
        /// </summary>
        [Input("update")]
        public Input<string>? Update { get; set; }

        public SubscriberTimeoutsGetArgs()
        {
        }
        public static new SubscriberTimeoutsGetArgs Empty => new SubscriberTimeoutsGetArgs();
    }
}
