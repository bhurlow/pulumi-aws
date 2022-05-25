// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Connect.Outputs
{

    [OutputType]
    public sealed class GetUserHierarchyStructureHierarchyStructureResult
    {
        /// <summary>
        /// A block that defines the details of level five. The level block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelFifeResult> LevelFives;
        /// <summary>
        /// A block that defines the details of level four. The level block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelFourResult> LevelFours;
        /// <summary>
        /// A block that defines the details of level one. The level block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelOneResult> LevelOnes;
        /// <summary>
        /// A block that defines the details of level three. The level block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelThreeResult> LevelThrees;
        /// <summary>
        /// A block that defines the details of level two. The level block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelTwoResult> LevelTwos;

        [OutputConstructor]
        private GetUserHierarchyStructureHierarchyStructureResult(
            ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelFifeResult> levelFives,

            ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelFourResult> levelFours,

            ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelOneResult> levelOnes,

            ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelThreeResult> levelThrees,

            ImmutableArray<Outputs.GetUserHierarchyStructureHierarchyStructureLevelTwoResult> levelTwos)
        {
            LevelFives = levelFives;
            LevelFours = levelFours;
            LevelOnes = levelOnes;
            LevelThrees = levelThrees;
            LevelTwos = levelTwos;
        }
    }
}
