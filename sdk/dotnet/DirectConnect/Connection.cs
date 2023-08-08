// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectConnect
{
    /// <summary>
    /// Provides a Connection of Direct Connect.
    /// 
    /// ## Example Usage
    /// ### Create a connection
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var hoge = new Aws.DirectConnect.Connection("hoge", new()
    ///     {
    ///         Bandwidth = "1Gbps",
    ///         Location = "EqDC2",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Request a MACsec-capable connection
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.DirectConnect.Connection("example", new()
    ///     {
    ///         Bandwidth = "10Gbps",
    ///         Location = "EqDA2",
    ///         RequestMacsec = true,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Configure encryption mode for MACsec-capable connections
    /// 
    /// &gt; **NOTE:** You can only specify the `encryption_mode` argument once the connection is in an `Available` state.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.DirectConnect.Connection("example", new()
    ///     {
    ///         Bandwidth = "10Gbps",
    ///         EncryptionMode = "must_encrypt",
    ///         Location = "EqDC2",
    ///         RequestMacsec = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_dx_connection.test_connection
    /// 
    ///  id = "dxcon-ffre0ec3" } Using `pulumi import`, import Direct Connect connections using the connection `id`. For exampleconsole % pulumi import aws_dx_connection.test_connection dxcon-ffre0ec3
    /// </summary>
    [AwsResourceType("aws:directconnect/connection:Connection")]
    public partial class Connection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the connection.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Direct Connect endpoint on which the physical connection terminates.
        /// </summary>
        [Output("awsDevice")]
        public Output<string> AwsDevice { get; private set; } = null!;

        /// <summary>
        /// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `no_encrypt`, `should_encrypt`, and `must_encrypt`.
        /// </summary>
        [Output("encryptionMode")]
        public Output<string> EncryptionMode { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
        /// </summary>
        [Output("hasLogicalRedundancy")]
        public Output<string> HasLogicalRedundancy { get; private set; } = null!;

        /// <summary>
        /// Boolean value representing if jumbo frames have been enabled for this connection.
        /// </summary>
        [Output("jumboFrameCapable")]
        public Output<bool> JumboFrameCapable { get; private set; } = null!;

        /// <summary>
        /// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Boolean value indicating whether the connection supports MAC Security (MACsec).
        /// </summary>
        [Output("macsecCapable")]
        public Output<bool> MacsecCapable { get; private set; } = null!;

        /// <summary>
        /// The name of the connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the connection.
        /// </summary>
        [Output("ownerAccountId")]
        public Output<string> OwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// The name of the AWS Direct Connect service provider associated with the connection.
        /// </summary>
        [Output("partnerName")]
        public Output<string> PartnerName { get; private set; } = null!;

        /// <summary>
        /// The MAC Security (MACsec) port link status of the connection.
        /// </summary>
        [Output("portEncryptionStatus")]
        public Output<string> PortEncryptionStatus { get; private set; } = null!;

        /// <summary>
        /// The name of the service provider associated with the connection.
        /// </summary>
        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;

        /// <summary>
        /// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
        /// 
        /// &gt; **NOTE:** Changing the value of `request_macsec` will cause the resource to be destroyed and re-created.
        /// </summary>
        [Output("requestMacsec")]
        public Output<bool?> RequestMacsec { get; private set; } = null!;

        /// <summary>
        /// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
        /// </summary>
        [Output("skipDestroy")]
        public Output<bool?> SkipDestroy { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Output("vlanId")]
        public Output<int> VlanId { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs args, CustomResourceOptions? options = null)
            : base("aws:directconnect/connection:Connection", name, args ?? new ConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
            : base("aws:directconnect/connection:Connection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, state, options);
        }
    }

    public sealed class ConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        /// </summary>
        [Input("bandwidth", required: true)]
        public Input<string> Bandwidth { get; set; } = null!;

        /// <summary>
        /// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `no_encrypt`, `should_encrypt`, and `must_encrypt`.
        /// </summary>
        [Input("encryptionMode")]
        public Input<string>? EncryptionMode { get; set; }

        /// <summary>
        /// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the service provider associated with the connection.
        /// </summary>
        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        /// <summary>
        /// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
        /// 
        /// &gt; **NOTE:** Changing the value of `request_macsec` will cause the resource to be destroyed and re-created.
        /// </summary>
        [Input("requestMacsec")]
        public Input<bool>? RequestMacsec { get; set; }

        /// <summary>
        /// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
        /// </summary>
        [Input("skipDestroy")]
        public Input<bool>? SkipDestroy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ConnectionArgs()
        {
        }
        public static new ConnectionArgs Empty => new ConnectionArgs();
    }

    public sealed class ConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the connection.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Direct Connect endpoint on which the physical connection terminates.
        /// </summary>
        [Input("awsDevice")]
        public Input<string>? AwsDevice { get; set; }

        /// <summary>
        /// The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `no_encrypt`, `should_encrypt`, and `must_encrypt`.
        /// </summary>
        [Input("encryptionMode")]
        public Input<string>? EncryptionMode { get; set; }

        /// <summary>
        /// Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
        /// </summary>
        [Input("hasLogicalRedundancy")]
        public Input<string>? HasLogicalRedundancy { get; set; }

        /// <summary>
        /// Boolean value representing if jumbo frames have been enabled for this connection.
        /// </summary>
        [Input("jumboFrameCapable")]
        public Input<bool>? JumboFrameCapable { get; set; }

        /// <summary>
        /// The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Boolean value indicating whether the connection supports MAC Security (MACsec).
        /// </summary>
        [Input("macsecCapable")]
        public Input<bool>? MacsecCapable { get; set; }

        /// <summary>
        /// The name of the connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the AWS account that owns the connection.
        /// </summary>
        [Input("ownerAccountId")]
        public Input<string>? OwnerAccountId { get; set; }

        /// <summary>
        /// The name of the AWS Direct Connect service provider associated with the connection.
        /// </summary>
        [Input("partnerName")]
        public Input<string>? PartnerName { get; set; }

        /// <summary>
        /// The MAC Security (MACsec) port link status of the connection.
        /// </summary>
        [Input("portEncryptionStatus")]
        public Input<string>? PortEncryptionStatus { get; set; }

        /// <summary>
        /// The name of the service provider associated with the connection.
        /// </summary>
        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        /// <summary>
        /// Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
        /// 
        /// &gt; **NOTE:** Changing the value of `request_macsec` will cause the resource to be destroyed and re-created.
        /// </summary>
        [Input("requestMacsec")]
        public Input<bool>? RequestMacsec { get; set; }

        /// <summary>
        /// Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
        /// </summary>
        [Input("skipDestroy")]
        public Input<bool>? SkipDestroy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public ConnectionState()
        {
        }
        public static new ConnectionState Empty => new ConnectionState();
    }
}
