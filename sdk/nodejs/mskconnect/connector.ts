// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an Amazon MSK Connect Connector resource.
 *
 * ## Example Usage
 * ### Basic configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.mskconnect.Connector("example", {
 *     kafkaconnectVersion: "2.7.1",
 *     capacity: {
 *         autoscaling: {
 *             mcuCount: 1,
 *             minWorkerCount: 1,
 *             maxWorkerCount: 2,
 *             scaleInPolicy: {
 *                 cpuUtilizationPercentage: 20,
 *             },
 *             scaleOutPolicy: {
 *                 cpuUtilizationPercentage: 80,
 *             },
 *         },
 *     },
 *     connectorConfiguration: {
 *         "connector.class": "com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector",
 *         "tasks.max": "1",
 *         topics: "example",
 *     },
 *     kafkaCluster: {
 *         apacheKafkaCluster: {
 *             bootstrapServers: aws_msk_cluster.example.bootstrap_brokers_tls,
 *             vpc: {
 *                 securityGroups: [aws_security_group.example.id],
 *                 subnets: [
 *                     aws_subnet.example1.id,
 *                     aws_subnet.example2.id,
 *                     aws_subnet.example3.id,
 *                 ],
 *             },
 *         },
 *     },
 *     kafkaClusterClientAuthentication: {
 *         authenticationType: "NONE",
 *     },
 *     kafkaClusterEncryptionInTransit: {
 *         encryptionType: "TLS",
 *     },
 *     plugins: [{
 *         customPlugin: {
 *             arn: aws_mskconnect_custom_plugin.example.arn,
 *             revision: aws_mskconnect_custom_plugin.example.latest_revision,
 *         },
 *     }],
 *     serviceExecutionRoleArn: aws_iam_role.example.arn,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_mskconnect_connector.example
 *
 *  id = "arn:aws:kafkaconnect:eu-central-1:123456789012:connector/example/264edee4-17a3-412e-bd76-6681cfc93805-3" } Using `pulumi import`, import MSK Connect Connector using the connector's `arn`. For exampleconsole % pulumi import aws_mskconnect_connector.example 'arn:aws:kafkaconnect:eu-central-1:123456789012:connector/example/264edee4-17a3-412e-bd76-6681cfc93805-3'
 */
export class Connector extends pulumi.CustomResource {
    /**
     * Get an existing Connector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectorState, opts?: pulumi.CustomResourceOptions): Connector {
        return new Connector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mskconnect/connector:Connector';

    /**
     * Returns true if the given object is an instance of Connector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connector.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the custom plugin.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Information about the capacity allocated to the connector. See below.
     */
    public readonly capacity!: pulumi.Output<outputs.mskconnect.ConnectorCapacity>;
    /**
     * A map of keys to values that represent the configuration for the connector.
     */
    public readonly connectorConfiguration!: pulumi.Output<{[key: string]: string}>;
    /**
     * A summary description of the connector.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies which Apache Kafka cluster to connect to. See below.
     */
    public readonly kafkaCluster!: pulumi.Output<outputs.mskconnect.ConnectorKafkaCluster>;
    /**
     * Details of the client authentication used by the Apache Kafka cluster. See below.
     */
    public readonly kafkaClusterClientAuthentication!: pulumi.Output<outputs.mskconnect.ConnectorKafkaClusterClientAuthentication>;
    /**
     * Details of encryption in transit to the Apache Kafka cluster. See below.
     */
    public readonly kafkaClusterEncryptionInTransit!: pulumi.Output<outputs.mskconnect.ConnectorKafkaClusterEncryptionInTransit>;
    /**
     * The version of Kafka Connect. It has to be compatible with both the Apache Kafka cluster's version and the plugins.
     */
    public readonly kafkaconnectVersion!: pulumi.Output<string>;
    /**
     * Details about log delivery. See below.
     */
    public readonly logDelivery!: pulumi.Output<outputs.mskconnect.ConnectorLogDelivery | undefined>;
    /**
     * The name of the connector.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies which plugins to use for the connector. See below.
     */
    public readonly plugins!: pulumi.Output<outputs.mskconnect.ConnectorPlugin[]>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role used by the connector to access the Amazon Web Services resources that it needs. The types of resources depends on the logic of the connector. For example, a connector that has Amazon S3 as a destination must have permissions that allow it to write to the S3 destination bucket.
     */
    public readonly serviceExecutionRoleArn!: pulumi.Output<string>;
    /**
     * The current version of the connector.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * Specifies which worker configuration to use with the connector. See below.
     */
    public readonly workerConfiguration!: pulumi.Output<outputs.mskconnect.ConnectorWorkerConfiguration | undefined>;

    /**
     * Create a Connector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectorArgs | ConnectorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectorState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["connectorConfiguration"] = state ? state.connectorConfiguration : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["kafkaCluster"] = state ? state.kafkaCluster : undefined;
            resourceInputs["kafkaClusterClientAuthentication"] = state ? state.kafkaClusterClientAuthentication : undefined;
            resourceInputs["kafkaClusterEncryptionInTransit"] = state ? state.kafkaClusterEncryptionInTransit : undefined;
            resourceInputs["kafkaconnectVersion"] = state ? state.kafkaconnectVersion : undefined;
            resourceInputs["logDelivery"] = state ? state.logDelivery : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["plugins"] = state ? state.plugins : undefined;
            resourceInputs["serviceExecutionRoleArn"] = state ? state.serviceExecutionRoleArn : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["workerConfiguration"] = state ? state.workerConfiguration : undefined;
        } else {
            const args = argsOrState as ConnectorArgs | undefined;
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            if ((!args || args.connectorConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectorConfiguration'");
            }
            if ((!args || args.kafkaCluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kafkaCluster'");
            }
            if ((!args || args.kafkaClusterClientAuthentication === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kafkaClusterClientAuthentication'");
            }
            if ((!args || args.kafkaClusterEncryptionInTransit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kafkaClusterEncryptionInTransit'");
            }
            if ((!args || args.kafkaconnectVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kafkaconnectVersion'");
            }
            if ((!args || args.plugins === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plugins'");
            }
            if ((!args || args.serviceExecutionRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceExecutionRoleArn'");
            }
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["connectorConfiguration"] = args ? args.connectorConfiguration : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["kafkaCluster"] = args ? args.kafkaCluster : undefined;
            resourceInputs["kafkaClusterClientAuthentication"] = args ? args.kafkaClusterClientAuthentication : undefined;
            resourceInputs["kafkaClusterEncryptionInTransit"] = args ? args.kafkaClusterEncryptionInTransit : undefined;
            resourceInputs["kafkaconnectVersion"] = args ? args.kafkaconnectVersion : undefined;
            resourceInputs["logDelivery"] = args ? args.logDelivery : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["plugins"] = args ? args.plugins : undefined;
            resourceInputs["serviceExecutionRoleArn"] = args ? args.serviceExecutionRoleArn : undefined;
            resourceInputs["workerConfiguration"] = args ? args.workerConfiguration : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connector resources.
 */
export interface ConnectorState {
    /**
     * The Amazon Resource Name (ARN) of the custom plugin.
     */
    arn?: pulumi.Input<string>;
    /**
     * Information about the capacity allocated to the connector. See below.
     */
    capacity?: pulumi.Input<inputs.mskconnect.ConnectorCapacity>;
    /**
     * A map of keys to values that represent the configuration for the connector.
     */
    connectorConfiguration?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A summary description of the connector.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies which Apache Kafka cluster to connect to. See below.
     */
    kafkaCluster?: pulumi.Input<inputs.mskconnect.ConnectorKafkaCluster>;
    /**
     * Details of the client authentication used by the Apache Kafka cluster. See below.
     */
    kafkaClusterClientAuthentication?: pulumi.Input<inputs.mskconnect.ConnectorKafkaClusterClientAuthentication>;
    /**
     * Details of encryption in transit to the Apache Kafka cluster. See below.
     */
    kafkaClusterEncryptionInTransit?: pulumi.Input<inputs.mskconnect.ConnectorKafkaClusterEncryptionInTransit>;
    /**
     * The version of Kafka Connect. It has to be compatible with both the Apache Kafka cluster's version and the plugins.
     */
    kafkaconnectVersion?: pulumi.Input<string>;
    /**
     * Details about log delivery. See below.
     */
    logDelivery?: pulumi.Input<inputs.mskconnect.ConnectorLogDelivery>;
    /**
     * The name of the connector.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies which plugins to use for the connector. See below.
     */
    plugins?: pulumi.Input<pulumi.Input<inputs.mskconnect.ConnectorPlugin>[]>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role used by the connector to access the Amazon Web Services resources that it needs. The types of resources depends on the logic of the connector. For example, a connector that has Amazon S3 as a destination must have permissions that allow it to write to the S3 destination bucket.
     */
    serviceExecutionRoleArn?: pulumi.Input<string>;
    /**
     * The current version of the connector.
     */
    version?: pulumi.Input<string>;
    /**
     * Specifies which worker configuration to use with the connector. See below.
     */
    workerConfiguration?: pulumi.Input<inputs.mskconnect.ConnectorWorkerConfiguration>;
}

/**
 * The set of arguments for constructing a Connector resource.
 */
export interface ConnectorArgs {
    /**
     * Information about the capacity allocated to the connector. See below.
     */
    capacity: pulumi.Input<inputs.mskconnect.ConnectorCapacity>;
    /**
     * A map of keys to values that represent the configuration for the connector.
     */
    connectorConfiguration: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A summary description of the connector.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies which Apache Kafka cluster to connect to. See below.
     */
    kafkaCluster: pulumi.Input<inputs.mskconnect.ConnectorKafkaCluster>;
    /**
     * Details of the client authentication used by the Apache Kafka cluster. See below.
     */
    kafkaClusterClientAuthentication: pulumi.Input<inputs.mskconnect.ConnectorKafkaClusterClientAuthentication>;
    /**
     * Details of encryption in transit to the Apache Kafka cluster. See below.
     */
    kafkaClusterEncryptionInTransit: pulumi.Input<inputs.mskconnect.ConnectorKafkaClusterEncryptionInTransit>;
    /**
     * The version of Kafka Connect. It has to be compatible with both the Apache Kafka cluster's version and the plugins.
     */
    kafkaconnectVersion: pulumi.Input<string>;
    /**
     * Details about log delivery. See below.
     */
    logDelivery?: pulumi.Input<inputs.mskconnect.ConnectorLogDelivery>;
    /**
     * The name of the connector.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies which plugins to use for the connector. See below.
     */
    plugins: pulumi.Input<pulumi.Input<inputs.mskconnect.ConnectorPlugin>[]>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role used by the connector to access the Amazon Web Services resources that it needs. The types of resources depends on the logic of the connector. For example, a connector that has Amazon S3 as a destination must have permissions that allow it to write to the S3 destination bucket.
     */
    serviceExecutionRoleArn: pulumi.Input<string>;
    /**
     * Specifies which worker configuration to use with the connector. See below.
     */
    workerConfiguration?: pulumi.Input<inputs.mskconnect.ConnectorWorkerConfiguration>;
}
