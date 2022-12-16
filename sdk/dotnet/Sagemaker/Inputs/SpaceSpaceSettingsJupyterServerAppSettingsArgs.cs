// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker.Inputs
{

    public sealed class SpaceSpaceSettingsJupyterServerAppSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("codeRepositories")]
        private InputList<Inputs.SpaceSpaceSettingsJupyterServerAppSettingsCodeRepositoryArgs>? _codeRepositories;

        /// <summary>
        /// A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterServer application. see Code Repository below.
        /// </summary>
        public InputList<Inputs.SpaceSpaceSettingsJupyterServerAppSettingsCodeRepositoryArgs> CodeRepositories
        {
            get => _codeRepositories ?? (_codeRepositories = new InputList<Inputs.SpaceSpaceSettingsJupyterServerAppSettingsCodeRepositoryArgs>());
            set => _codeRepositories = value;
        }

        /// <summary>
        /// The default instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. see Default Resource Spec below.
        /// </summary>
        [Input("defaultResourceSpec", required: true)]
        public Input<Inputs.SpaceSpaceSettingsJupyterServerAppSettingsDefaultResourceSpecArgs> DefaultResourceSpec { get; set; } = null!;

        [Input("lifecycleConfigArns")]
        private InputList<string>? _lifecycleConfigArns;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lifecycle Configurations.
        /// </summary>
        public InputList<string> LifecycleConfigArns
        {
            get => _lifecycleConfigArns ?? (_lifecycleConfigArns = new InputList<string>());
            set => _lifecycleConfigArns = value;
        }

        public SpaceSpaceSettingsJupyterServerAppSettingsArgs()
        {
        }
        public static new SpaceSpaceSettingsJupyterServerAppSettingsArgs Empty => new SpaceSpaceSettingsJupyterServerAppSettingsArgs();
    }
}
