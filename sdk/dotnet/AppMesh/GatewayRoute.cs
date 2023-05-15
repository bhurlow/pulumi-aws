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
    /// Provides an AWS App Mesh gateway route resource.
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
    ///     var example = new Aws.AppMesh.GatewayRoute("example", new()
    ///     {
    ///         MeshName = "example-service-mesh",
    ///         VirtualGatewayName = aws_appmesh_virtual_gateway.Example.Name,
    ///         Spec = new Aws.AppMesh.Inputs.GatewayRouteSpecArgs
    ///         {
    ///             HttpRoute = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteArgs
    ///             {
    ///                 Action = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteActionArgs
    ///                 {
    ///                     Target = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteActionTargetArgs
    ///                     {
    ///                         VirtualService = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs
    ///                         {
    ///                             VirtualServiceName = aws_appmesh_virtual_service.Example.Name,
    ///                         },
    ///                     },
    ///                 },
    ///                 Match = new Aws.AppMesh.Inputs.GatewayRouteSpecHttpRouteMatchArgs
    ///                 {
    ///                     Prefix = "/",
    ///                 },
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "Environment", "test" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// App Mesh gateway routes can be imported using `mesh_name` and `virtual_gateway_name` together with the gateway route's `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:appmesh/gatewayRoute:GatewayRoute example mesh/gw1/example-gateway-route
    /// ```
    /// </summary>
    [AwsResourceType("aws:appmesh/gatewayRoute:GatewayRoute")]
    public partial class GatewayRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the gateway route.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Creation date of the gateway route.
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// Last update date of the gateway route.
        /// </summary>
        [Output("lastUpdatedDate")]
        public Output<string> LastUpdatedDate { get; private set; } = null!;

        /// <summary>
        /// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("meshName")]
        public Output<string> MeshName { get; private set; } = null!;

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Output("meshOwner")]
        public Output<string> MeshOwner { get; private set; } = null!;

        /// <summary>
        /// Name to use for the gateway route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource owner's AWS account ID.
        /// </summary>
        [Output("resourceOwner")]
        public Output<string> ResourceOwner { get; private set; } = null!;

        /// <summary>
        /// Gateway route specification to apply.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.GatewayRouteSpec> Spec { get; private set; } = null!;

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
        /// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("virtualGatewayName")]
        public Output<string> VirtualGatewayName { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayRoute(string name, GatewayRouteArgs args, CustomResourceOptions? options = null)
            : base("aws:appmesh/gatewayRoute:GatewayRoute", name, args ?? new GatewayRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayRoute(string name, Input<string> id, GatewayRouteState? state = null, CustomResourceOptions? options = null)
            : base("aws:appmesh/gatewayRoute:GatewayRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayRoute Get(string name, Input<string> id, GatewayRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayRoute(name, id, state, options);
        }
    }

    public sealed class GatewayRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName", required: true)]
        public Input<string> MeshName { get; set; } = null!;

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// Name to use for the gateway route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Gateway route specification to apply.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Inputs.GatewayRouteSpecArgs> Spec { get; set; } = null!;

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

        /// <summary>
        /// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("virtualGatewayName", required: true)]
        public Input<string> VirtualGatewayName { get; set; } = null!;

        public GatewayRouteArgs()
        {
        }
        public static new GatewayRouteArgs Empty => new GatewayRouteArgs();
    }

    public sealed class GatewayRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the gateway route.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Creation date of the gateway route.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// Last update date of the gateway route.
        /// </summary>
        [Input("lastUpdatedDate")]
        public Input<string>? LastUpdatedDate { get; set; }

        /// <summary>
        /// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName")]
        public Input<string>? MeshName { get; set; }

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// Name to use for the gateway route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource owner's AWS account ID.
        /// </summary>
        [Input("resourceOwner")]
        public Input<string>? ResourceOwner { get; set; }

        /// <summary>
        /// Gateway route specification to apply.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.GatewayRouteSpecGetArgs>? Spec { get; set; }

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

        /// <summary>
        /// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("virtualGatewayName")]
        public Input<string>? VirtualGatewayName { get; set; }

        public GatewayRouteState()
        {
        }
        public static new GatewayRouteState Empty => new GatewayRouteState();
    }
}
