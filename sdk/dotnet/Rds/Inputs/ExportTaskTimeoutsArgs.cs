// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Inputs
{

    public sealed class ExportTaskTimeoutsArgs : global::Pulumi.ResourceArgs
    {
        [Input("create")]
        public Input<string>? Create { get; set; }

        [Input("delete")]
        public Input<string>? Delete { get; set; }

        public ExportTaskTimeoutsArgs()
        {
        }
        public static new ExportTaskTimeoutsArgs Empty => new ExportTaskTimeoutsArgs();
    }
}
