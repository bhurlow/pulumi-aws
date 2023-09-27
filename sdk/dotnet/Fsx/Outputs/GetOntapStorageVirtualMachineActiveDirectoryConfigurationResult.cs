// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fsx.Outputs
{

    [OutputType]
    public sealed class GetOntapStorageVirtualMachineActiveDirectoryConfigurationResult
    {
        /// <summary>
        /// The NetBIOS name of the AD computer object to which the SVM is joined.
        /// </summary>
        public readonly string NetbiosName;
        public readonly ImmutableArray<Outputs.GetOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationResult> SelfManagedActiveDirectoryConfigurations;

        [OutputConstructor]
        private GetOntapStorageVirtualMachineActiveDirectoryConfigurationResult(
            string netbiosName,

            ImmutableArray<Outputs.GetOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationResult> selfManagedActiveDirectoryConfigurations)
        {
            NetbiosName = netbiosName;
            SelfManagedActiveDirectoryConfigurations = selfManagedActiveDirectoryConfigurations;
        }
    }
}
