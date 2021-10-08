// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Inputs
{

    public sealed class ClusterMasterInstanceFleetLaunchSpecificationsGetArgs : Pulumi.ResourceArgs
    {
        [Input("onDemandSpecifications")]
        private InputList<Inputs.ClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecificationGetArgs>? _onDemandSpecifications;

        /// <summary>
        /// Configuration block for on demand instances launch specifications.
        /// </summary>
        public InputList<Inputs.ClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecificationGetArgs> OnDemandSpecifications
        {
            get => _onDemandSpecifications ?? (_onDemandSpecifications = new InputList<Inputs.ClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecificationGetArgs>());
            set => _onDemandSpecifications = value;
        }

        [Input("spotSpecifications")]
        private InputList<Inputs.ClusterMasterInstanceFleetLaunchSpecificationsSpotSpecificationGetArgs>? _spotSpecifications;

        /// <summary>
        /// Configuration block for spot instances launch specifications.
        /// </summary>
        public InputList<Inputs.ClusterMasterInstanceFleetLaunchSpecificationsSpotSpecificationGetArgs> SpotSpecifications
        {
            get => _spotSpecifications ?? (_spotSpecifications = new InputList<Inputs.ClusterMasterInstanceFleetLaunchSpecificationsSpotSpecificationGetArgs>());
            set => _spotSpecifications = value;
        }

        public ClusterMasterInstanceFleetLaunchSpecificationsGetArgs()
        {
        }
    }
}
