// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectConnect
{
    public static class GetLocations
    {
        /// <summary>
        /// Retrieve information about the AWS Direct Connect locations in the current AWS Region.
        /// These are the locations that can be specified when configuring `aws.directconnect.Connection` or `aws.directconnect.LinkAggregationGroup` resources.
        /// 
        /// &gt; **Note:** This data source is different from the `aws.directconnect.getLocation` data source which retrieves information about a specific AWS Direct Connect location in the current AWS Region.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var available = Aws.DirectConnect.GetLocations.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLocationsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLocationsResult>("aws:directconnect/getLocations:getLocations", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetLocationsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Code for the locations.
        /// </summary>
        public readonly ImmutableArray<string> LocationCodes;

        [OutputConstructor]
        private GetLocationsResult(
            string id,

            ImmutableArray<string> locationCodes)
        {
            Id = id;
            LocationCodes = locationCodes;
        }
    }
}
