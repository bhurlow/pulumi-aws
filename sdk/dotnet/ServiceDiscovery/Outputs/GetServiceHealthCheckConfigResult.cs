// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceDiscovery.Outputs
{

    [OutputType]
    public sealed class GetServiceHealthCheckConfigResult
    {
        /// <summary>
        /// The number of 30-second intervals that you want service discovery to wait before it changes the health status of a service instance.  Maximum value of 10.
        /// </summary>
        public readonly int FailureThreshold;
        /// <summary>
        /// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
        /// </summary>
        public readonly string ResourcePath;
        /// <summary>
        /// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy. Valid Values: HTTP, HTTPS, TCP
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServiceHealthCheckConfigResult(
            int failureThreshold,

            string resourcePath,

            string type)
        {
            FailureThreshold = failureThreshold;
            ResourcePath = resourcePath;
            Type = type;
        }
    }
}