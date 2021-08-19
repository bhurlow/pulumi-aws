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
    /// Provides an AWS App Mesh route resource.
    /// 
    /// ## Example Usage
    /// ### HTTP Routing
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serviceb = new Aws.AppMesh.Route("serviceb", new Aws.AppMesh.RouteArgs
    ///         {
    ///             MeshName = aws_appmesh_mesh.Simple.Id,
    ///             VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
    ///             Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
    ///             {
    ///                 HttpRoute = new Aws.AppMesh.Inputs.RouteSpecHttpRouteArgs
    ///                 {
    ///                     Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
    ///                     {
    ///                         Prefix = "/",
    ///                     },
    ///                     Action = new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionArgs
    ///                     {
    ///                         WeightedTargets = 
    ///                         {
    ///                             new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
    ///                             {
    ///                                 VirtualNode = aws_appmesh_virtual_node.Serviceb1.Name,
    ///                                 Weight = 90,
    ///                             },
    ///                             new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
    ///                             {
    ///                                 VirtualNode = aws_appmesh_virtual_node.Serviceb2.Name,
    ///                                 Weight = 10,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### HTTP Header Routing
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serviceb = new Aws.AppMesh.Route("serviceb", new Aws.AppMesh.RouteArgs
    ///         {
    ///             MeshName = aws_appmesh_mesh.Simple.Id,
    ///             VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
    ///             Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
    ///             {
    ///                 HttpRoute = new Aws.AppMesh.Inputs.RouteSpecHttpRouteArgs
    ///                 {
    ///                     Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
    ///                     {
    ///                         Method = "POST",
    ///                         Prefix = "/",
    ///                         Scheme = "https",
    ///                         Headers = 
    ///                         {
    ///                             new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchHeaderArgs
    ///                             {
    ///                                 Name = "clientRequestId",
    ///                                 Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchHeaderMatchArgs
    ///                                 {
    ///                                     Prefix = "123",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     Action = new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionArgs
    ///                     {
    ///                         WeightedTargets = 
    ///                         {
    ///                             new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
    ///                             {
    ///                                 VirtualNode = aws_appmesh_virtual_node.Serviceb.Name,
    ///                                 Weight = 100,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Retry Policy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serviceb = new Aws.AppMesh.Route("serviceb", new Aws.AppMesh.RouteArgs
    ///         {
    ///             MeshName = aws_appmesh_mesh.Simple.Id,
    ///             VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
    ///             Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
    ///             {
    ///                 HttpRoute = new Aws.AppMesh.Inputs.RouteSpecHttpRouteArgs
    ///                 {
    ///                     Match = new Aws.AppMesh.Inputs.RouteSpecHttpRouteMatchArgs
    ///                     {
    ///                         Prefix = "/",
    ///                     },
    ///                     RetryPolicy = new Aws.AppMesh.Inputs.RouteSpecHttpRouteRetryPolicyArgs
    ///                     {
    ///                         HttpRetryEvents = 
    ///                         {
    ///                             "server-error",
    ///                         },
    ///                         MaxRetries = 1,
    ///                         PerRetryTimeout = new Aws.AppMesh.Inputs.RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs
    ///                         {
    ///                             Unit = "s",
    ///                             Value = 15,
    ///                         },
    ///                     },
    ///                     Action = new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionArgs
    ///                     {
    ///                         WeightedTargets = 
    ///                         {
    ///                             new Aws.AppMesh.Inputs.RouteSpecHttpRouteActionWeightedTargetArgs
    ///                             {
    ///                                 VirtualNode = aws_appmesh_virtual_node.Serviceb.Name,
    ///                                 Weight = 100,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### TCP Routing
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serviceb = new Aws.AppMesh.Route("serviceb", new Aws.AppMesh.RouteArgs
    ///         {
    ///             MeshName = aws_appmesh_mesh.Simple.Id,
    ///             VirtualRouterName = aws_appmesh_virtual_router.Serviceb.Name,
    ///             Spec = new Aws.AppMesh.Inputs.RouteSpecArgs
    ///             {
    ///                 TcpRoute = new Aws.AppMesh.Inputs.RouteSpecTcpRouteArgs
    ///                 {
    ///                     Action = new Aws.AppMesh.Inputs.RouteSpecTcpRouteActionArgs
    ///                     {
    ///                         WeightedTargets = 
    ///                         {
    ///                             new Aws.AppMesh.Inputs.RouteSpecTcpRouteActionWeightedTargetArgs
    ///                             {
    ///                                 VirtualNode = aws_appmesh_virtual_node.Serviceb1.Name,
    ///                                 Weight = 100,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// App Mesh virtual routes can be imported using `mesh_name` and `virtual_router_name` together with the route's `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:appmesh/route:Route serviceb simpleapp/serviceB/serviceB-route
    /// ```
    /// 
    ///  [1]/docs/providers/aws/index.html
    /// </summary>
    [AwsResourceType("aws:appmesh/route:Route")]
    public partial class Route : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the route.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The creation date of the route.
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// The last update date of the route.
        /// </summary>
        [Output("lastUpdatedDate")]
        public Output<string> LastUpdatedDate { get; private set; } = null!;

        /// <summary>
        /// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("meshName")]
        public Output<string> MeshName { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
        /// </summary>
        [Output("meshOwner")]
        public Output<string> MeshOwner { get; private set; } = null!;

        /// <summary>
        /// The name to use for the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource owner's AWS account ID.
        /// </summary>
        [Output("resourceOwner")]
        public Output<string> ResourceOwner { get; private set; } = null!;

        /// <summary>
        /// The route specification to apply.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.RouteSpec> Spec { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("virtualRouterName")]
        public Output<string> VirtualRouterName { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("aws:appmesh/route:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
            : base("aws:appmesh/route:Route", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
        {
            return new Route(name, id, state, options);
        }
    }

    public sealed class RouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName", required: true)]
        public Input<string> MeshName { get; set; } = null!;

        /// <summary>
        /// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// The name to use for the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The route specification to apply.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Inputs.RouteSpecArgs> Spec { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("virtualRouterName", required: true)]
        public Input<string> VirtualRouterName { get; set; } = null!;

        public RouteArgs()
        {
        }
    }

    public sealed class RouteState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the route.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The creation date of the route.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// The last update date of the route.
        /// </summary>
        [Input("lastUpdatedDate")]
        public Input<string>? LastUpdatedDate { get; set; }

        /// <summary>
        /// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName")]
        public Input<string>? MeshName { get; set; }

        /// <summary>
        /// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// The name to use for the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource owner's AWS account ID.
        /// </summary>
        [Input("resourceOwner")]
        public Input<string>? ResourceOwner { get; set; }

        /// <summary>
        /// The route specification to apply.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.RouteSpecGetArgs>? Spec { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("virtualRouterName")]
        public Input<string>? VirtualRouterName { get; set; }

        public RouteState()
        {
        }
    }
}
