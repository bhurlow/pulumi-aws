// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53.Inputs
{

    public sealed class RecordGeoproximityRoutingPolicyCoordinateArgs : global::Pulumi.ResourceArgs
    {
        [Input("latitude", required: true)]
        public Input<string> Latitude { get; set; } = null!;

        [Input("longitude", required: true)]
        public Input<string> Longitude { get; set; } = null!;

        public RecordGeoproximityRoutingPolicyCoordinateArgs()
        {
        }
        public static new RecordGeoproximityRoutingPolicyCoordinateArgs Empty => new RecordGeoproximityRoutingPolicyCoordinateArgs();
    }
}
