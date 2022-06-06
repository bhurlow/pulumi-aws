// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty.Inputs
{

    public sealed class DetectorDatasourcesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures [Kubernetes protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html).
        /// See Kubernetes and Kubernetes Audit Logs below for more details.
        /// </summary>
        [Input("kubernetes")]
        public Input<Inputs.DetectorDatasourcesKubernetesGetArgs>? Kubernetes { get; set; }

        /// <summary>
        /// Configures [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
        /// See S3 Logs below for more details.
        /// </summary>
        [Input("s3Logs")]
        public Input<Inputs.DetectorDatasourcesS3LogsGetArgs>? S3Logs { get; set; }

        public DetectorDatasourcesGetArgs()
        {
        }
    }
}
