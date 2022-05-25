// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Connect.Inputs
{

    public sealed class UserHierarchyGroupHierarchyPathGetArgs : Pulumi.ResourceArgs
    {
        [Input("levelFives")]
        private InputList<Inputs.UserHierarchyGroupHierarchyPathLevelFifeGetArgs>? _levelFives;

        /// <summary>
        /// A block that defines the details of level five. The level block is documented below.
        /// </summary>
        public InputList<Inputs.UserHierarchyGroupHierarchyPathLevelFifeGetArgs> LevelFives
        {
            get => _levelFives ?? (_levelFives = new InputList<Inputs.UserHierarchyGroupHierarchyPathLevelFifeGetArgs>());
            set => _levelFives = value;
        }

        [Input("levelFours")]
        private InputList<Inputs.UserHierarchyGroupHierarchyPathLevelFourGetArgs>? _levelFours;

        /// <summary>
        /// A block that defines the details of level four. The level block is documented below.
        /// </summary>
        public InputList<Inputs.UserHierarchyGroupHierarchyPathLevelFourGetArgs> LevelFours
        {
            get => _levelFours ?? (_levelFours = new InputList<Inputs.UserHierarchyGroupHierarchyPathLevelFourGetArgs>());
            set => _levelFours = value;
        }

        [Input("levelOnes")]
        private InputList<Inputs.UserHierarchyGroupHierarchyPathLevelOneGetArgs>? _levelOnes;

        /// <summary>
        /// A block that defines the details of level one. The level block is documented below.
        /// </summary>
        public InputList<Inputs.UserHierarchyGroupHierarchyPathLevelOneGetArgs> LevelOnes
        {
            get => _levelOnes ?? (_levelOnes = new InputList<Inputs.UserHierarchyGroupHierarchyPathLevelOneGetArgs>());
            set => _levelOnes = value;
        }

        [Input("levelThrees")]
        private InputList<Inputs.UserHierarchyGroupHierarchyPathLevelThreeGetArgs>? _levelThrees;

        /// <summary>
        /// A block that defines the details of level three. The level block is documented below.
        /// </summary>
        public InputList<Inputs.UserHierarchyGroupHierarchyPathLevelThreeGetArgs> LevelThrees
        {
            get => _levelThrees ?? (_levelThrees = new InputList<Inputs.UserHierarchyGroupHierarchyPathLevelThreeGetArgs>());
            set => _levelThrees = value;
        }

        [Input("levelTwos")]
        private InputList<Inputs.UserHierarchyGroupHierarchyPathLevelTwoGetArgs>? _levelTwos;

        /// <summary>
        /// A block that defines the details of level two. The level block is documented below.
        /// </summary>
        public InputList<Inputs.UserHierarchyGroupHierarchyPathLevelTwoGetArgs> LevelTwos
        {
            get => _levelTwos ?? (_levelTwos = new InputList<Inputs.UserHierarchyGroupHierarchyPathLevelTwoGetArgs>());
            set => _levelTwos = value;
        }

        public UserHierarchyGroupHierarchyPathGetArgs()
        {
        }
    }
}
