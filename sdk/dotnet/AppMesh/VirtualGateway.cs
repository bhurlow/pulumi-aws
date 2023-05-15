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
    /// Provides an AWS App Mesh virtual gateway resource.
    /// 
    /// ## Example Usage
    /// ### Access Logs and TLS
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppMesh.VirtualGateway("example", new()
    ///     {
    ///         MeshName = "example-service-mesh",
    ///         Spec = new Aws.AppMesh.Inputs.VirtualGatewaySpecArgs
    ///         {
    ///             Listener = new Aws.AppMesh.Inputs.VirtualGatewaySpecListenerArgs
    ///             {
    ///                 PortMapping = new Aws.AppMesh.Inputs.VirtualGatewaySpecListenerPortMappingArgs
    ///                 {
    ///                     Port = 8080,
    ///                     Protocol = "http",
    ///                 },
    ///                 Tls = new Aws.AppMesh.Inputs.VirtualGatewaySpecListenerTlsArgs
    ///                 {
    ///                     Certificate = new Aws.AppMesh.Inputs.VirtualGatewaySpecListenerTlsCertificateArgs
    ///                     {
    ///                         Acm = new Aws.AppMesh.Inputs.VirtualGatewaySpecListenerTlsCertificateAcmArgs
    ///                         {
    ///                             CertificateArn = aws_acm_certificate.Example.Arn,
    ///                         },
    ///                     },
    ///                     Mode = "STRICT",
    ///                 },
    ///             },
    ///             Logging = new Aws.AppMesh.Inputs.VirtualGatewaySpecLoggingArgs
    ///             {
    ///                 AccessLog = new Aws.AppMesh.Inputs.VirtualGatewaySpecLoggingAccessLogArgs
    ///                 {
    ///                     File = new Aws.AppMesh.Inputs.VirtualGatewaySpecLoggingAccessLogFileArgs
    ///                     {
    ///                         Path = "/var/log/access.log",
    ///                     },
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
    /// App Mesh virtual gateway can be imported using `mesh_name` together with the virtual gateway's `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:appmesh/virtualGateway:VirtualGateway example mesh/gw1
    /// ```
    /// </summary>
    [AwsResourceType("aws:appmesh/virtualGateway:VirtualGateway")]
    public partial class VirtualGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the virtual gateway.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Creation date of the virtual gateway.
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// Last update date of the virtual gateway.
        /// </summary>
        [Output("lastUpdatedDate")]
        public Output<string> LastUpdatedDate { get; private set; } = null!;

        /// <summary>
        /// Name of the service mesh in which to create the virtual gateway. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("meshName")]
        public Output<string> MeshName { get; private set; } = null!;

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Output("meshOwner")]
        public Output<string> MeshOwner { get; private set; } = null!;

        /// <summary>
        /// Name to use for the virtual gateway. Must be between 1 and 255 characters in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource owner's AWS account ID.
        /// </summary>
        [Output("resourceOwner")]
        public Output<string> ResourceOwner { get; private set; } = null!;

        /// <summary>
        /// Virtual gateway specification to apply.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.VirtualGatewaySpec> Spec { get; private set; } = null!;

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
        /// Create a VirtualGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualGateway(string name, VirtualGatewayArgs args, CustomResourceOptions? options = null)
            : base("aws:appmesh/virtualGateway:VirtualGateway", name, args ?? new VirtualGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualGateway(string name, Input<string> id, VirtualGatewayState? state = null, CustomResourceOptions? options = null)
            : base("aws:appmesh/virtualGateway:VirtualGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualGateway Get(string name, Input<string> id, VirtualGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualGateway(name, id, state, options);
        }
    }

    public sealed class VirtualGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the service mesh in which to create the virtual gateway. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName", required: true)]
        public Input<string> MeshName { get; set; } = null!;

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// Name to use for the virtual gateway. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Virtual gateway specification to apply.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Inputs.VirtualGatewaySpecArgs> Spec { get; set; } = null!;

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

        public VirtualGatewayArgs()
        {
        }
        public static new VirtualGatewayArgs Empty => new VirtualGatewayArgs();
    }

    public sealed class VirtualGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the virtual gateway.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Creation date of the virtual gateway.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// Last update date of the virtual gateway.
        /// </summary>
        [Input("lastUpdatedDate")]
        public Input<string>? LastUpdatedDate { get; set; }

        /// <summary>
        /// Name of the service mesh in which to create the virtual gateway. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("meshName")]
        public Input<string>? MeshName { get; set; }

        /// <summary>
        /// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
        /// </summary>
        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        /// <summary>
        /// Name to use for the virtual gateway. Must be between 1 and 255 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource owner's AWS account ID.
        /// </summary>
        [Input("resourceOwner")]
        public Input<string>? ResourceOwner { get; set; }

        /// <summary>
        /// Virtual gateway specification to apply.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.VirtualGatewaySpecGetArgs>? Spec { get; set; }

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

        public VirtualGatewayState()
        {
        }
        public static new VirtualGatewayState Empty => new VirtualGatewayState();
    }
}
