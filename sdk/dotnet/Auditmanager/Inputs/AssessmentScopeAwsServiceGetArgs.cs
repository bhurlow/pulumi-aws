// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Auditmanager.Inputs
{

    public sealed class AssessmentScopeAwsServiceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Amazon Web Service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public AssessmentScopeAwsServiceGetArgs()
        {
        }
        public static new AssessmentScopeAwsServiceGetArgs Empty => new AssessmentScopeAwsServiceGetArgs();
    }
}
