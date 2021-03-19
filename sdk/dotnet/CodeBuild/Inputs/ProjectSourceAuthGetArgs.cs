// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild.Inputs
{

    public sealed class ProjectSourceAuthGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource value that applies to the specified authorization type. Use the `aws.codebuild.SourceCredential` resource instead.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// Authorization type to use. The only valid value is `OAUTH`. This data type is deprecated and is no longer accurate or used. Use the `aws.codebuild.SourceCredential` resource instead.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProjectSourceAuthGetArgs()
        {
        }
    }
}
