// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Outputs
{

    [OutputType]
    public sealed class EndpointConfigurationProductionVariant
    {
        /// <summary>
        /// The size of the Elastic Inference (EI) instance to use for the production variant.
        /// </summary>
        public readonly string? AcceleratorType;
        /// <summary>
        /// The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting. For more information about health check, see [How Your Container Should Respond to Health Check (Ping) Requests](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-algo-ping-requests). Valid values between `60` and `3600`.
        /// </summary>
        public readonly int? ContainerStartupHealthCheckTimeoutInSeconds;
        /// <summary>
        /// Specifies configuration for a core dump from the model container when the process crashes. Fields are documented below.
        /// </summary>
        public readonly Outputs.EndpointConfigurationProductionVariantCoreDumpConfig? CoreDumpConfig;
        /// <summary>
        /// You can use this parameter to turn on native Amazon Web Services Systems Manager (SSM) access for a production variant behind an endpoint. By default, SSM access is disabled for all production variants behind an endpoints.
        /// </summary>
        public readonly bool? EnableSsmAccess;
        /// <summary>
        /// Initial number of instances used for auto-scaling.
        /// </summary>
        public readonly int? InitialInstanceCount;
        /// <summary>
        /// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration. If unspecified, it defaults to `1.0`.
        /// </summary>
        public readonly double? InitialVariantWeight;
        /// <summary>
        /// The type of instance to start.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant. Valid values between `60` and `3600`.
        /// </summary>
        public readonly int? ModelDataDownloadTimeoutInSeconds;
        /// <summary>
        /// The name of the model to use.
        /// </summary>
        public readonly string ModelName;
        /// <summary>
        /// Sets how the endpoint routes incoming traffic. See routing_config below.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointConfigurationProductionVariantRoutingConfig> RoutingConfigs;
        /// <summary>
        /// Specifies configuration for how an endpoint performs asynchronous inference.
        /// </summary>
        public readonly Outputs.EndpointConfigurationProductionVariantServerlessConfig? ServerlessConfig;
        /// <summary>
        /// The name of the variant. If omitted, this provider will assign a random, unique name.
        /// </summary>
        public readonly string? VariantName;
        /// <summary>
        /// The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant. Valid values between `1` and `512`.
        /// </summary>
        public readonly int? VolumeSizeInGb;

        [OutputConstructor]
        private EndpointConfigurationProductionVariant(
            string? acceleratorType,

            int? containerStartupHealthCheckTimeoutInSeconds,

            Outputs.EndpointConfigurationProductionVariantCoreDumpConfig? coreDumpConfig,

            bool? enableSsmAccess,

            int? initialInstanceCount,

            double? initialVariantWeight,

            string? instanceType,

            int? modelDataDownloadTimeoutInSeconds,

            string modelName,

            ImmutableArray<Outputs.EndpointConfigurationProductionVariantRoutingConfig> routingConfigs,

            Outputs.EndpointConfigurationProductionVariantServerlessConfig? serverlessConfig,

            string? variantName,

            int? volumeSizeInGb)
        {
            AcceleratorType = acceleratorType;
            ContainerStartupHealthCheckTimeoutInSeconds = containerStartupHealthCheckTimeoutInSeconds;
            CoreDumpConfig = coreDumpConfig;
            EnableSsmAccess = enableSsmAccess;
            InitialInstanceCount = initialInstanceCount;
            InitialVariantWeight = initialVariantWeight;
            InstanceType = instanceType;
            ModelDataDownloadTimeoutInSeconds = modelDataDownloadTimeoutInSeconds;
            ModelName = modelName;
            RoutingConfigs = routingConfigs;
            ServerlessConfig = serverlessConfig;
            VariantName = variantName;
            VolumeSizeInGb = volumeSizeInGb;
        }
    }
}
