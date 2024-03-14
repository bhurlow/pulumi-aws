// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityLake.Inputs
{

    public sealed class SubscriberSourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Security Lake supports log and event collection for natively supported AWS services.
        /// </summary>
        [Input("awsLogSourceResource")]
        public Input<Inputs.SubscriberSourceAwsLogSourceResourceGetArgs>? AwsLogSourceResource { get; set; }

        /// <summary>
        /// Amazon Security Lake supports custom source types.
        /// </summary>
        [Input("customLogSourceResource")]
        public Input<Inputs.SubscriberSourceCustomLogSourceResourceGetArgs>? CustomLogSourceResource { get; set; }

        public SubscriberSourceGetArgs()
        {
        }
        public static new SubscriberSourceGetArgs Empty => new SubscriberSourceGetArgs();
    }
}
