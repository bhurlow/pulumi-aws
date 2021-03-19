// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclLoggingConfigurationRedactedFieldArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Redact all query arguments.
        /// </summary>
        [Input("allQueryArguments")]
        public Input<Inputs.WebAclLoggingConfigurationRedactedFieldAllQueryArgumentsArgs>? AllQueryArguments { get; set; }

        /// <summary>
        /// Redact the request body, which immediately follows the request headers.
        /// </summary>
        [Input("body")]
        public Input<Inputs.WebAclLoggingConfigurationRedactedFieldBodyArgs>? Body { get; set; }

        /// <summary>
        /// Redact the HTTP method. Must be specified as an empty configuration block `{}`. The method indicates the type of operation that the request is asking the origin to perform.
        /// </summary>
        [Input("method")]
        public Input<Inputs.WebAclLoggingConfigurationRedactedFieldMethodArgs>? Method { get; set; }

        /// <summary>
        /// Redact the query string. Must be specified as an empty configuration block `{}`. This is the part of a URL that appears after a `?` character, if any.
        /// </summary>
        [Input("queryString")]
        public Input<Inputs.WebAclLoggingConfigurationRedactedFieldQueryStringArgs>? QueryString { get; set; }

        /// <summary>
        /// Redact a single header. See Single Header below for details.
        /// </summary>
        [Input("singleHeader")]
        public Input<Inputs.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs>? SingleHeader { get; set; }

        /// <summary>
        /// Redact a single query argument. See Single Query Argument below for details.
        /// </summary>
        [Input("singleQueryArgument")]
        public Input<Inputs.WebAclLoggingConfigurationRedactedFieldSingleQueryArgumentArgs>? SingleQueryArgument { get; set; }

        /// <summary>
        /// Redact the request URI path. Must be specified as an empty configuration block `{}`. This is the part of a web request that identifies a resource, for example, `/images/daily-ad.jpg`.
        /// </summary>
        [Input("uriPath")]
        public Input<Inputs.WebAclLoggingConfigurationRedactedFieldUriPathArgs>? UriPath { get; set; }

        public WebAclLoggingConfigurationRedactedFieldArgs()
        {
        }
    }
}
