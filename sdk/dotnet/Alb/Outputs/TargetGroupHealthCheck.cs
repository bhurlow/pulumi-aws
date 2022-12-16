// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Alb.Outputs
{

    [OutputType]
    public sealed class TargetGroupHealthCheck
    {
        /// <summary>
        /// Boolean to enable / disable `stickiness`. Default is `true`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Number of consecutive health check successes required before considering a target healthy. The range is 2-10. Defaults to 3.
        /// </summary>
        public readonly int? HealthyThreshold;
        /// <summary>
        /// Approximate amount of time, in seconds, between health checks of an individual target. The range is 5-300. For `lambda` target groups, it needs to be greater than the timeout of the underlying `lambda`. Defaults to 30.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// Response codes to use when checking for a healthy responses from a target. You can specify multiple values (for example, "200,202" for HTTP(s) or "0,12" for GRPC) or a range of values (for example, "200-299" or "0-99"). Required for HTTP/HTTPS/GRPC ALB. Only applies to Application Load Balancers (i.e., HTTP/HTTPS/GRPC) not Network Load Balancers (i.e., TCP).
        /// </summary>
        public readonly string? Matcher;
        /// <summary>
        /// Destination for the health check request. Required for HTTP/HTTPS ALB and HTTP NLB. Only applies to HTTP/HTTPS.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The port the load balancer uses when performing health checks on targets. Default is traffic-port.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// Protocol the load balancer uses when performing health checks on targets. Must be either `TCP`, `HTTP`, or `HTTPS`. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Amount of time, in seconds, during which no response from a target means a failed health check. The range is 2–120 seconds. For target groups with a protocol of HTTP, the default is 6 seconds. For target groups with a protocol of TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a protocol of GENEVE, the default is 5 seconds. If the target type is lambda, the default is 30 seconds.
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// Number of consecutive health check failures required before considering a target unhealthy. The range is 2-10. Defaults to 3.
        /// </summary>
        public readonly int? UnhealthyThreshold;

        [OutputConstructor]
        private TargetGroupHealthCheck(
            bool? enabled,

            int? healthyThreshold,

            int? interval,

            string? matcher,

            string? path,

            string? port,

            string? protocol,

            int? timeout,

            int? unhealthyThreshold)
        {
            Enabled = enabled;
            HealthyThreshold = healthyThreshold;
            Interval = interval;
            Matcher = matcher;
            Path = path;
            Port = port;
            Protocol = protocol;
            Timeout = timeout;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
