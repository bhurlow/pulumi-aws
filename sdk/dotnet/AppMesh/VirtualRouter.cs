// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh
{
    /// <summary>
    /// Provides an AWS App Mesh virtual router resource.
    /// 
    /// ## Breaking Changes
    /// 
    /// Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `aws.appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
    /// 
    /// * Remove service `service_names` from the `spec` argument.
    /// AWS has created a `aws.appmesh.VirtualService` resource for each of service names.
    /// These resource can be imported using `import`.
    /// 
    /// * Add a `listener` configuration block to the `spec` argument.
    /// 
    /// The state associated with existing resources will automatically be migrated.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var serviceb = new Aws.AppMesh.VirtualRouter("serviceb", new()
    ///     {
    ///         MeshName = aws_appmesh_mesh.Simple.Id,
    ///         Spec = new Aws.AppMesh.Inputs.VirtualRouterSpecArgs
    ///         {
    ///             Listener = new Aws.AppMesh.Inputs.VirtualRouterSpecListenerArgs
    ///             {
    ///                 PortMapping = new Aws.AppMesh.Inputs.VirtualRouterSpecListenerPortMappingArgs
    ///                 {
    ///                     Port = 8080,
    ///                     Protocol = "http",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// App Mesh virtual routers can be imported using `mesh_name` together with the virtual router's `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:appmesh/virtualRouter:VirtualRouter serviceb simpleapp/serviceB
    /// ```
    /// </summary>
    [AwsResourceType("aws:appmesh/virtualRouter:VirtualRouter")]
    public partial class VirtualRouter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the virtual router.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Creation date of the virtual router.
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// Last update date of the virtual router.
        /// </summary>
        [Output("lastUpdatedDate")]
        public Output<string> LastUpdatedDate { get; private set; } = null!;

        /// <summary>
        /// Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("meshName")]
        public Output<string> MeshName { get; private set; } = null!;

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Output("meshOwner")]
        public Output<string> MeshOwner { get; private set; } = null!;

        /// <summary>
        /// Name to use for the virtual router. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource owner's AWS account ID.
        /// </summary>
        [Output("resourceOwner")]
        public Output<string> ResourceOwner { get; private set; } = null!;

        /// <summary>
        /// Virtual router specification to apply.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.VirtualRouterSpec> Spec { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualRouter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualRouter(string name, VirtualRouterArgs args, CustomResourceOptions? options = null)
            : base("aws:appmesh/virtualRouter:VirtualRouter", name, args ?? new VirtualRouterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualRouter(string name, Input<string> id, VirtualRouterState? state = null, CustomResourceOptions? options = null)
            : base("aws:appmesh/virtualRouter:VirtualRouter", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualRouter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualRouter Get(string name, Input<string> id, VirtualRouterState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualRouter(name, id, state, options);
        }
    }

    public sealed class VirtualRouterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName", required: true)]
        public Input<string> MeshName { get; set; } = null!;

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// Name to use for the virtual router. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Virtual router specification to apply.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Inputs.VirtualRouterSpecArgs> Spec { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VirtualRouterArgs()
        {
        }
        public static new VirtualRouterArgs Empty => new VirtualRouterArgs();
    }

    public sealed class VirtualRouterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the virtual router.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Creation date of the virtual router.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// Last update date of the virtual router.
        /// </summary>
        [Input("lastUpdatedDate")]
        public Input<string>? LastUpdatedDate { get; set; }

        /// <summary>
        /// Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName")]
        public Input<string>? MeshName { get; set; }

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// Name to use for the virtual router. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource owner's AWS account ID.
        /// </summary>
        [Input("resourceOwner")]
        public Input<string>? ResourceOwner { get; set; }

        /// <summary>
        /// Virtual router specification to apply.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.VirtualRouterSpecGetArgs>? Spec { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public VirtualRouterState()
        {
        }
        public static new VirtualRouterState Empty => new VirtualRouterState();
    }
}
