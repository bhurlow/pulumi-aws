// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityLake.Outputs
{

    [OutputType]
    public sealed class SubscriberSourceCustomLogSourceResource
    {
        /// <summary>
        /// The attributes of a third-party custom source.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubscriberSourceCustomLogSourceResourceAttribute> Attributes;
        public readonly ImmutableArray<Outputs.SubscriberSourceCustomLogSourceResourceProvider> Providers;
        /// <summary>
        /// The name for a third-party custom source. This must be a Regionally unique value.
        /// </summary>
        public readonly string? SourceName;
        /// <summary>
        /// The version for a third-party custom source. This must be a Regionally unique value.
        /// </summary>
        public readonly string? SourceVersion;

        [OutputConstructor]
        private SubscriberSourceCustomLogSourceResource(
            ImmutableArray<Outputs.SubscriberSourceCustomLogSourceResourceAttribute> attributes,

            ImmutableArray<Outputs.SubscriberSourceCustomLogSourceResourceProvider> providers,

            string? sourceName,

            string? sourceVersion)
        {
            Attributes = attributes;
            Providers = providers;
            SourceName = sourceName;
            SourceVersion = sourceVersion;
        }
    }
}
