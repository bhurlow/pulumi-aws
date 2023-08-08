// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DataSync
{
    /// <summary>
    /// Manages an HDFS Location within AWS DataSync.
    /// 
    /// &gt; **NOTE:** The DataSync Agents must be available before creating this resource.
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
    ///     var example = new Aws.DataSync.LocationHdfs("example", new()
    ///     {
    ///         AgentArns = new[]
    ///         {
    ///             aws_datasync_agent.Example.Arn,
    ///         },
    ///         AuthenticationType = "SIMPLE",
    ///         SimpleUser = "example",
    ///         NameNodes = new[]
    ///         {
    ///             new Aws.DataSync.Inputs.LocationHdfsNameNodeArgs
    ///             {
    ///                 Hostname = aws_instance.Example.Private_dns,
    ///                 Port = 80,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_datasync_location_hdfs.example
    /// 
    ///  id = "arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567" } Using `pulumi import`, import `aws_datasync_location_hdfs` using the Amazon Resource Name (ARN). For exampleconsole % pulumi import aws_datasync_location_hdfs.example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
    /// </summary>
    [AwsResourceType("aws:datasync/locationHdfs:LocationHdfs")]
    public partial class LocationHdfs : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of DataSync Agent ARNs with which this location will be associated.
        /// </summary>
        [Output("agentArns")]
        public Output<ImmutableArray<string>> AgentArns { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the DataSync Location.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The type of authentication used to determine the identity of the user. Valid values are `SIMPLE` and `KERBEROS`.
        /// </summary>
        [Output("authenticationType")]
        public Output<string?> AuthenticationType { get; private set; } = null!;

        /// <summary>
        /// The size of data blocks to write into the HDFS cluster. The block size must be a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
        /// </summary>
        [Output("blockSize")]
        public Output<int?> BlockSize { get; private set; } = null!;

        /// <summary>
        /// The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Output("kerberosKeytab")]
        public Output<string?> KerberosKeytab { get; private set; } = null!;

        /// <summary>
        /// The krb5.conf file that contains the Kerberos configuration information. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Output("kerberosKrb5Conf")]
        public Output<string?> KerberosKrb5Conf { get; private set; } = null!;

        /// <summary>
        /// The Kerberos principal with access to the files and folders on the HDFS cluster. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Output("kerberosPrincipal")]
        public Output<string?> KerberosPrincipal { get; private set; } = null!;

        /// <summary>
        /// The URI of the HDFS cluster's Key Management Server (KMS).
        /// </summary>
        [Output("kmsKeyProviderUri")]
        public Output<string?> KmsKeyProviderUri { get; private set; } = null!;

        /// <summary>
        /// The NameNode that manages the HDFS namespace. The NameNode performs operations such as opening, closing, and renaming files and directories. The NameNode contains the information to map blocks of data to the DataNodes. You can use only one NameNode. See configuration below.
        /// </summary>
        [Output("nameNodes")]
        public Output<ImmutableArray<Outputs.LocationHdfsNameNode>> NameNodes { get; private set; } = null!;

        /// <summary>
        /// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster. If `qop_configuration` isn't specified, `rpc_protection` and `data_transfer_protection` default to `PRIVACY`. If you set RpcProtection or DataTransferProtection, the other parameter assumes the same value.  See configuration below.
        /// </summary>
        [Output("qopConfiguration")]
        public Output<Outputs.LocationHdfsQopConfiguration?> QopConfiguration { get; private set; } = null!;

        /// <summary>
        /// The number of DataNodes to replicate the data to when writing to the HDFS cluster. By default, data is replicated to three DataNodes.
        /// </summary>
        [Output("replicationFactor")]
        public Output<int?> ReplicationFactor { get; private set; } = null!;

        /// <summary>
        /// The user name used to identify the client on the host operating system. If `SIMPLE` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Output("simpleUser")]
        public Output<string?> SimpleUser { get; private set; } = null!;

        /// <summary>
        /// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
        /// </summary>
        [Output("subdirectory")]
        public Output<string?> Subdirectory { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a LocationHdfs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocationHdfs(string name, LocationHdfsArgs args, CustomResourceOptions? options = null)
            : base("aws:datasync/locationHdfs:LocationHdfs", name, args ?? new LocationHdfsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocationHdfs(string name, Input<string> id, LocationHdfsState? state = null, CustomResourceOptions? options = null)
            : base("aws:datasync/locationHdfs:LocationHdfs", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LocationHdfs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocationHdfs Get(string name, Input<string> id, LocationHdfsState? state = null, CustomResourceOptions? options = null)
        {
            return new LocationHdfs(name, id, state, options);
        }
    }

    public sealed class LocationHdfsArgs : global::Pulumi.ResourceArgs
    {
        [Input("agentArns", required: true)]
        private InputList<string>? _agentArns;

        /// <summary>
        /// A list of DataSync Agent ARNs with which this location will be associated.
        /// </summary>
        public InputList<string> AgentArns
        {
            get => _agentArns ?? (_agentArns = new InputList<string>());
            set => _agentArns = value;
        }

        /// <summary>
        /// The type of authentication used to determine the identity of the user. Valid values are `SIMPLE` and `KERBEROS`.
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// The size of data blocks to write into the HDFS cluster. The block size must be a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
        /// </summary>
        [Input("blockSize")]
        public Input<int>? BlockSize { get; set; }

        /// <summary>
        /// The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("kerberosKeytab")]
        public Input<string>? KerberosKeytab { get; set; }

        /// <summary>
        /// The krb5.conf file that contains the Kerberos configuration information. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("kerberosKrb5Conf")]
        public Input<string>? KerberosKrb5Conf { get; set; }

        /// <summary>
        /// The Kerberos principal with access to the files and folders on the HDFS cluster. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("kerberosPrincipal")]
        public Input<string>? KerberosPrincipal { get; set; }

        /// <summary>
        /// The URI of the HDFS cluster's Key Management Server (KMS).
        /// </summary>
        [Input("kmsKeyProviderUri")]
        public Input<string>? KmsKeyProviderUri { get; set; }

        [Input("nameNodes", required: true)]
        private InputList<Inputs.LocationHdfsNameNodeArgs>? _nameNodes;

        /// <summary>
        /// The NameNode that manages the HDFS namespace. The NameNode performs operations such as opening, closing, and renaming files and directories. The NameNode contains the information to map blocks of data to the DataNodes. You can use only one NameNode. See configuration below.
        /// </summary>
        public InputList<Inputs.LocationHdfsNameNodeArgs> NameNodes
        {
            get => _nameNodes ?? (_nameNodes = new InputList<Inputs.LocationHdfsNameNodeArgs>());
            set => _nameNodes = value;
        }

        /// <summary>
        /// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster. If `qop_configuration` isn't specified, `rpc_protection` and `data_transfer_protection` default to `PRIVACY`. If you set RpcProtection or DataTransferProtection, the other parameter assumes the same value.  See configuration below.
        /// </summary>
        [Input("qopConfiguration")]
        public Input<Inputs.LocationHdfsQopConfigurationArgs>? QopConfiguration { get; set; }

        /// <summary>
        /// The number of DataNodes to replicate the data to when writing to the HDFS cluster. By default, data is replicated to three DataNodes.
        /// </summary>
        [Input("replicationFactor")]
        public Input<int>? ReplicationFactor { get; set; }

        /// <summary>
        /// The user name used to identify the client on the host operating system. If `SIMPLE` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("simpleUser")]
        public Input<string>? SimpleUser { get; set; }

        /// <summary>
        /// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
        /// </summary>
        [Input("subdirectory")]
        public Input<string>? Subdirectory { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public LocationHdfsArgs()
        {
        }
        public static new LocationHdfsArgs Empty => new LocationHdfsArgs();
    }

    public sealed class LocationHdfsState : global::Pulumi.ResourceArgs
    {
        [Input("agentArns")]
        private InputList<string>? _agentArns;

        /// <summary>
        /// A list of DataSync Agent ARNs with which this location will be associated.
        /// </summary>
        public InputList<string> AgentArns
        {
            get => _agentArns ?? (_agentArns = new InputList<string>());
            set => _agentArns = value;
        }

        /// <summary>
        /// Amazon Resource Name (ARN) of the DataSync Location.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The type of authentication used to determine the identity of the user. Valid values are `SIMPLE` and `KERBEROS`.
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// The size of data blocks to write into the HDFS cluster. The block size must be a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
        /// </summary>
        [Input("blockSize")]
        public Input<int>? BlockSize { get; set; }

        /// <summary>
        /// The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("kerberosKeytab")]
        public Input<string>? KerberosKeytab { get; set; }

        /// <summary>
        /// The krb5.conf file that contains the Kerberos configuration information. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("kerberosKrb5Conf")]
        public Input<string>? KerberosKrb5Conf { get; set; }

        /// <summary>
        /// The Kerberos principal with access to the files and folders on the HDFS cluster. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("kerberosPrincipal")]
        public Input<string>? KerberosPrincipal { get; set; }

        /// <summary>
        /// The URI of the HDFS cluster's Key Management Server (KMS).
        /// </summary>
        [Input("kmsKeyProviderUri")]
        public Input<string>? KmsKeyProviderUri { get; set; }

        [Input("nameNodes")]
        private InputList<Inputs.LocationHdfsNameNodeGetArgs>? _nameNodes;

        /// <summary>
        /// The NameNode that manages the HDFS namespace. The NameNode performs operations such as opening, closing, and renaming files and directories. The NameNode contains the information to map blocks of data to the DataNodes. You can use only one NameNode. See configuration below.
        /// </summary>
        public InputList<Inputs.LocationHdfsNameNodeGetArgs> NameNodes
        {
            get => _nameNodes ?? (_nameNodes = new InputList<Inputs.LocationHdfsNameNodeGetArgs>());
            set => _nameNodes = value;
        }

        /// <summary>
        /// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster. If `qop_configuration` isn't specified, `rpc_protection` and `data_transfer_protection` default to `PRIVACY`. If you set RpcProtection or DataTransferProtection, the other parameter assumes the same value.  See configuration below.
        /// </summary>
        [Input("qopConfiguration")]
        public Input<Inputs.LocationHdfsQopConfigurationGetArgs>? QopConfiguration { get; set; }

        /// <summary>
        /// The number of DataNodes to replicate the data to when writing to the HDFS cluster. By default, data is replicated to three DataNodes.
        /// </summary>
        [Input("replicationFactor")]
        public Input<int>? ReplicationFactor { get; set; }

        /// <summary>
        /// The user name used to identify the client on the host operating system. If `SIMPLE` is specified for `authentication_type`, this parameter is required.
        /// </summary>
        [Input("simpleUser")]
        public Input<string>? SimpleUser { get; set; }

        /// <summary>
        /// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
        /// </summary>
        [Input("subdirectory")]
        public Input<string>? Subdirectory { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public LocationHdfsState()
        {
        }
        public static new LocationHdfsState Empty => new LocationHdfsState();
    }
}
