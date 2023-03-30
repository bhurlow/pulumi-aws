// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Inputs
{

    public sealed class VirtualNodeSpecLoggingAccessLogFileGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The specified format for the logs.
        /// </summary>
        [Input("format")]
        public Input<Inputs.VirtualNodeSpecLoggingAccessLogFileFormatGetArgs>? Format { get; set; }

        /// <summary>
        /// File path to write access logs to. You can use `/dev/stdout` to send access logs to standard out. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public VirtualNodeSpecLoggingAccessLogFileGetArgs()
        {
        }
        public static new VirtualNodeSpecLoggingAccessLogFileGetArgs Empty => new VirtualNodeSpecLoggingAccessLogFileGetArgs();
    }
}
