// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mq.Inputs
{

    public sealed class BrokerUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable access to the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) for the user. Applies to `engine_type` of `ActiveMQ` only.
        /// </summary>
        [Input("consoleAccess")]
        public Input<bool>? ConsoleAccess { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of groups (20 maximum) to which the ActiveMQ user belongs. Applies to `engine_type` of `ActiveMQ` only.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Password of the user. It must be 12 to 250 characters long, at least 4 unique characters, and must not contain commas.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Username of the user.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public BrokerUserArgs()
        {
        }
    }
}
