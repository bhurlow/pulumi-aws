// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Mq.Outputs
{

    [OutputType]
    public sealed class GetBrokerUserResult
    {
        public readonly bool ConsoleAccess;
        public readonly ImmutableArray<string> Groups;
        public readonly bool ReplicationUser;
        public readonly string Username;

        [OutputConstructor]
        private GetBrokerUserResult(
            bool consoleAccess,

            ImmutableArray<string> groups,

            bool replicationUser,

            string username)
        {
            ConsoleAccess = consoleAccess;
            Groups = groups;
            ReplicationUser = replicationUser;
            Username = username;
        }
    }
}
