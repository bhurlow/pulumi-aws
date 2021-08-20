// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Inputs
{

    public sealed class EventConnectionAuthParametersOauthOauthHttpParametersArgs : Pulumi.ResourceArgs
    {
        [Input("bodies")]
        private InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersBodyArgs>? _bodies;

        /// <summary>
        /// Contains additional body string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
        /// </summary>
        public InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersBodyArgs> Bodies
        {
            get => _bodies ?? (_bodies = new InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersBodyArgs>());
            set => _bodies = value;
        }

        [Input("headers")]
        private InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersHeaderArgs>? _headers;

        /// <summary>
        /// Contains additional header parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
        /// </summary>
        public InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersHeaderArgs>());
            set => _headers = value;
        }

        [Input("queryStrings")]
        private InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersQueryStringArgs>? _queryStrings;

        /// <summary>
        /// Contains additional query string parameters for the connection. You can include up to 100 additional body string parameters per request. Each additional parameter counts towards the event payload size, which cannot exceed 64 KB. Each parameter can contain the following:
        /// </summary>
        public InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersQueryStringArgs> QueryStrings
        {
            get => _queryStrings ?? (_queryStrings = new InputList<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersQueryStringArgs>());
            set => _queryStrings = value;
        }

        public EventConnectionAuthParametersOauthOauthHttpParametersArgs()
        {
        }
    }
}