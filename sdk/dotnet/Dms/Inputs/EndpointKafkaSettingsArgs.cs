// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms.Inputs
{

    public sealed class EndpointKafkaSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kafka broker location. Specify in the form broker-hostname-or-ip:port.
        /// </summary>
        [Input("broker", required: true)]
        public Input<string> Broker { get; set; } = null!;

        /// <summary>
        /// Shows detailed control information for table definition, column definition, and table and column changes in the Kafka message output. The default is `false`.
        /// </summary>
        [Input("includeControlDetails")]
        public Input<bool>? IncludeControlDetails { get; set; }

        /// <summary>
        /// Include NULL and empty columns for records migrated to the endpoint. The default is `false`.
        /// </summary>
        [Input("includeNullAndEmpty")]
        public Input<bool>? IncludeNullAndEmpty { get; set; }

        /// <summary>
        /// Shows the partition value within the Kafka message output unless the partition type is `schema-table-type`. The default is `false`.
        /// </summary>
        [Input("includePartitionValue")]
        public Input<bool>? IncludePartitionValue { get; set; }

        /// <summary>
        /// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table`, `drop-table`, `add-column`, `drop-column`, and `rename-column`. The default is `false`.
        /// </summary>
        [Input("includeTableAlterOperations")]
        public Input<bool>? IncludeTableAlterOperations { get; set; }

        /// <summary>
        /// Provides detailed transaction information from the source database. This information includes a commit timestamp, a log position, and values for `transaction_id`, previous `transaction_id`, and `transaction_record_id` (the record offset within a transaction). The default is `false`.
        /// </summary>
        [Input("includeTransactionDetails")]
        public Input<bool>? IncludeTransactionDetails { get; set; }

        /// <summary>
        /// The output format for the records created on the endpoint. The message format is `JSON` (default) or `JSON_UNFORMATTED` (a single line with no tab).
        /// </summary>
        [Input("messageFormat")]
        public Input<string>? MessageFormat { get; set; }

        /// <summary>
        /// The maximum size in bytes for records created on the endpoint The default is `1,000,000`.
        /// </summary>
        [Input("messageMaxBytes")]
        public Input<int>? MessageMaxBytes { get; set; }

        /// <summary>
        /// Set this optional parameter to true to avoid adding a '0x' prefix to raw data in hexadecimal format. For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka target. Use the `no_hex_prefix` endpoint setting to enable migration of RAW data type columns without adding the `'0x'` prefix.
        /// </summary>
        [Input("noHexPrefix")]
        public Input<bool>? NoHexPrefix { get; set; }

        /// <summary>
        /// Prefixes schema and table names to partition values, when the partition type is `primary-key-type`. Doing this increases data distribution among Kafka partitions. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same partition, which causes throttling. The default is `false`.
        /// </summary>
        [Input("partitionIncludeSchemaTable")]
        public Input<bool>? PartitionIncludeSchemaTable { get; set; }

        /// <summary>
        /// The secure password you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
        /// </summary>
        [Input("saslPassword")]
        public Input<string>? SaslPassword { get; set; }

        /// <summary>
        /// The secure user name you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
        /// </summary>
        [Input("saslUsername")]
        public Input<string>? SaslUsername { get; set; }

        /// <summary>
        /// Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS). Options include `ssl-encryption`, `ssl-authentication`, and `sasl-ssl`. `sasl-ssl` requires `sasl_username` and `sasl_password`.
        /// </summary>
        [Input("securityProtocol")]
        public Input<string>? SecurityProtocol { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the private certificate authority (CA) cert that AWS DMS uses to securely connect to your Kafka target endpoint.
        /// </summary>
        [Input("sslCaCertificateArn")]
        public Input<string>? SslCaCertificateArn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the client certificate used to securely connect to a Kafka target endpoint.
        /// </summary>
        [Input("sslClientCertificateArn")]
        public Input<string>? SslClientCertificateArn { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the client private key used to securely connect to a Kafka target endpoint.
        /// </summary>
        [Input("sslClientKeyArn")]
        public Input<string>? SslClientKeyArn { get; set; }

        /// <summary>
        /// The password for the client private key used to securely connect to a Kafka target endpoint.
        /// </summary>
        [Input("sslClientKeyPassword")]
        public Input<string>? SslClientKeyPassword { get; set; }

        /// <summary>
        /// Kafka topic for migration. Defaults to `kafka-default-topic`.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public EndpointKafkaSettingsArgs()
        {
        }
    }
}
