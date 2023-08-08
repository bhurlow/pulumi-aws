// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ResourceExplorer.Inputs
{

    public sealed class ViewFiltersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The string that contains the search keywords, prefixes, and operators to control the results that can be returned by a search operation. For more details, see [Search query syntax](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html).
        /// </summary>
        [Input("filterString", required: true)]
        public Input<string> FilterString { get; set; } = null!;

        public ViewFiltersArgs()
        {
        }
        public static new ViewFiltersArgs Empty => new ViewFiltersArgs();
    }
}
