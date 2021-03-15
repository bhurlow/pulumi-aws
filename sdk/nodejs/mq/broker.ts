// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Amazon MQ broker resource. This resources also manages users for the broker.
 *
 * > For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).
 *
 * > **NOTE:** Amazon MQ currently places limits on **RabbitMQ** brokers. For example, a RabbitMQ broker cannot have: instances with an associated IP address of an ENI attached to the broker, an associated LDAP server to authenticate and authorize broker connections, storage type `EFS`, audit logging, or `configuration` blocks. Although this resource allows you to create RabbitMQ users, RabbitMQ users cannot have console access or groups. Also, Amazon MQ does not return information about RabbitMQ users so drift detection is not possible.
 *
 * > **NOTE:** Changes to an MQ Broker can occur when you change a parameter, such as `configuration` or `user`, and are reflected in the next maintenance window. Because of this, the provider may report a difference in its planning phase because a modification has not yet taken place. You can use the `applyImmediately` flag to instruct the service to apply the change immediately (see documentation below). Using `applyImmediately` can result in a brief downtime as the broker reboots.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.mq.Broker("example", {
 *     brokerName: "example",
 *     configuration: {
 *         id: aws_mq_configuration.test.id,
 *         revision: aws_mq_configuration.test.latest_revision,
 *     },
 *     engineType: "ActiveMQ",
 *     engineVersion: "5.15.9",
 *     hostInstanceType: "mq.t2.micro",
 *     securityGroups: [aws_security_group.test.id],
 *     users: [{
 *         username: "ExampleUser",
 *         password: "MindTheGap",
 *     }],
 * });
 * ```
 * ### High-throughput Optimized Example
 *
 * This example shows the use of EBS storage for high-throughput optimized performance.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.mq.Broker("example", {
 *     brokerName: "example",
 *     configuration: {
 *         id: aws_mq_configuration.test.id,
 *         revision: aws_mq_configuration.test.latest_revision,
 *     },
 *     engineType: "ActiveMQ",
 *     engineVersion: "5.15.9",
 *     storageType: "ebs",
 *     hostInstanceType: "mq.m5.large",
 *     securityGroups: [aws_security_group.test.id],
 *     users: [{
 *         username: "ExampleUser",
 *         password: "MindTheGap",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * MQ Brokers can be imported using their broker id, e.g.
 *
 * ```sh
 *  $ pulumi import aws:mq/broker:Broker example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
 * ```
 */
export class Broker extends pulumi.CustomResource {
    /**
     * Get an existing Broker resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BrokerState, opts?: pulumi.CustomResourceOptions): Broker {
        return new Broker(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mq/broker:Broker';

    /**
     * Returns true if the given object is an instance of Broker.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Broker {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Broker.__pulumiType;
    }

    /**
     * Specifies whether any broker modifications are applied immediately, or during the next maintenance window. Default is `false`.
     */
    public readonly applyImmediately!: pulumi.Output<boolean | undefined>;
    /**
     * ARN of the broker.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Authentication strategy used to secure the broker. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
     */
    public readonly authenticationStrategy!: pulumi.Output<string>;
    /**
     * Whether to automatically upgrade to new minor versions of brokers as Amazon MQ makes releases available.
     */
    public readonly autoMinorVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the broker.
     */
    public readonly brokerName!: pulumi.Output<string>;
    /**
     * Configuration block for broker configuration. Applies to `engineType` of `ActiveMQ` only. Detailed below.
     */
    public readonly configuration!: pulumi.Output<outputs.mq.BrokerConfiguration>;
    /**
     * Deployment mode of the broker. Valid values are `SINGLE_INSTANCE`, `ACTIVE_STANDBY_MULTI_AZ`, and `CLUSTER_MULTI_AZ`. Default is `SINGLE_INSTANCE`.
     */
    public readonly deploymentMode!: pulumi.Output<string | undefined>;
    /**
     * Configuration block containing encryption options. Detailed below.
     */
    public readonly encryptionOptions!: pulumi.Output<outputs.mq.BrokerEncryptionOptions | undefined>;
    /**
     * Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
     */
    public readonly engineType!: pulumi.Output<string>;
    /**
     * Version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions. For example, `5.15.0`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Broker's instance type. For example, `mq.t3.micro`, `mq.m5.large`.
     */
    public readonly hostInstanceType!: pulumi.Output<string>;
    /**
     * List of information about allocated brokers (both active & standby).
     * * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
     * * `instances.0.ip_address` - IP Address of the broker.
     * * `instances.0.endpoints` - Broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
     * * For `ActiveMQ`:
     * * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
     * * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
     * * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
     * * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
     * * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
     * * For `RabbitMQ`:
     * * `amqps://broker-id.mq.us-west-2.amazonaws.com:5671`
     */
    public /*out*/ readonly instances!: pulumi.Output<outputs.mq.BrokerInstance[]>;
    /**
     * Configuration block for the LDAP server used to authenticate and authorize connections to the broker. Not supported for `engineType` `RabbitMQ`. Detailed below. (Currently, AWS may not process changes to LDAP server metadata.)
     */
    public readonly ldapServerMetadata!: pulumi.Output<outputs.mq.BrokerLdapServerMetadata | undefined>;
    /**
     * Configuration block for the logging configuration of the broker. Detailed below.
     */
    public readonly logs!: pulumi.Output<outputs.mq.BrokerLogs | undefined>;
    /**
     * Configuration block for the maintenance window start time. Detailed below.
     */
    public readonly maintenanceWindowStartTime!: pulumi.Output<outputs.mq.BrokerMaintenanceWindowStartTime>;
    /**
     * Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
     */
    public readonly publiclyAccessible!: pulumi.Output<boolean | undefined>;
    /**
     * List of security group IDs assigned to the broker.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Storage type of the broker. For `engineType` `ActiveMQ`, the valid values are `efs` and `ebs`, and the AWS-default is `efs`. For `engineType` `RabbitMQ`, only `ebs` is supported. When using `ebs`, only the `mq.m5` broker instance type family is supported.
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * List of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires multiple subnets.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * Map of tags to assign to the broker.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Configuration block for broker users. For `engineType` of `RabbitMQ`, Amazon MQ does not return broker users preventing this resource from making user updates and drift detection. Detailed below.
     */
    public readonly users!: pulumi.Output<outputs.mq.BrokerUser[]>;

    /**
     * Create a Broker resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BrokerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BrokerArgs | BrokerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BrokerState | undefined;
            inputs["applyImmediately"] = state ? state.applyImmediately : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["authenticationStrategy"] = state ? state.authenticationStrategy : undefined;
            inputs["autoMinorVersionUpgrade"] = state ? state.autoMinorVersionUpgrade : undefined;
            inputs["brokerName"] = state ? state.brokerName : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["deploymentMode"] = state ? state.deploymentMode : undefined;
            inputs["encryptionOptions"] = state ? state.encryptionOptions : undefined;
            inputs["engineType"] = state ? state.engineType : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["hostInstanceType"] = state ? state.hostInstanceType : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["ldapServerMetadata"] = state ? state.ldapServerMetadata : undefined;
            inputs["logs"] = state ? state.logs : undefined;
            inputs["maintenanceWindowStartTime"] = state ? state.maintenanceWindowStartTime : undefined;
            inputs["publiclyAccessible"] = state ? state.publiclyAccessible : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["storageType"] = state ? state.storageType : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as BrokerArgs | undefined;
            if ((!args || args.brokerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'brokerName'");
            }
            if ((!args || args.engineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineType'");
            }
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if ((!args || args.hostInstanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostInstanceType'");
            }
            if ((!args || args.users === undefined) && !opts.urn) {
                throw new Error("Missing required property 'users'");
            }
            inputs["applyImmediately"] = args ? args.applyImmediately : undefined;
            inputs["authenticationStrategy"] = args ? args.authenticationStrategy : undefined;
            inputs["autoMinorVersionUpgrade"] = args ? args.autoMinorVersionUpgrade : undefined;
            inputs["brokerName"] = args ? args.brokerName : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["deploymentMode"] = args ? args.deploymentMode : undefined;
            inputs["encryptionOptions"] = args ? args.encryptionOptions : undefined;
            inputs["engineType"] = args ? args.engineType : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["hostInstanceType"] = args ? args.hostInstanceType : undefined;
            inputs["ldapServerMetadata"] = args ? args.ldapServerMetadata : undefined;
            inputs["logs"] = args ? args.logs : undefined;
            inputs["maintenanceWindowStartTime"] = args ? args.maintenanceWindowStartTime : undefined;
            inputs["publiclyAccessible"] = args ? args.publiclyAccessible : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["users"] = args ? args.users : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["instances"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Broker.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Broker resources.
 */
export interface BrokerState {
    /**
     * Specifies whether any broker modifications are applied immediately, or during the next maintenance window. Default is `false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * ARN of the broker.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Authentication strategy used to secure the broker. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
     */
    readonly authenticationStrategy?: pulumi.Input<string>;
    /**
     * Whether to automatically upgrade to new minor versions of brokers as Amazon MQ makes releases available.
     */
    readonly autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * Name of the broker.
     */
    readonly brokerName?: pulumi.Input<string>;
    /**
     * Configuration block for broker configuration. Applies to `engineType` of `ActiveMQ` only. Detailed below.
     */
    readonly configuration?: pulumi.Input<inputs.mq.BrokerConfiguration>;
    /**
     * Deployment mode of the broker. Valid values are `SINGLE_INSTANCE`, `ACTIVE_STANDBY_MULTI_AZ`, and `CLUSTER_MULTI_AZ`. Default is `SINGLE_INSTANCE`.
     */
    readonly deploymentMode?: pulumi.Input<string>;
    /**
     * Configuration block containing encryption options. Detailed below.
     */
    readonly encryptionOptions?: pulumi.Input<inputs.mq.BrokerEncryptionOptions>;
    /**
     * Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
     */
    readonly engineType?: pulumi.Input<string>;
    /**
     * Version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions. For example, `5.15.0`.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * Broker's instance type. For example, `mq.t3.micro`, `mq.m5.large`.
     */
    readonly hostInstanceType?: pulumi.Input<string>;
    /**
     * List of information about allocated brokers (both active & standby).
     * * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
     * * `instances.0.ip_address` - IP Address of the broker.
     * * `instances.0.endpoints` - Broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
     * * For `ActiveMQ`:
     * * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
     * * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
     * * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
     * * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
     * * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
     * * For `RabbitMQ`:
     * * `amqps://broker-id.mq.us-west-2.amazonaws.com:5671`
     */
    readonly instances?: pulumi.Input<pulumi.Input<inputs.mq.BrokerInstance>[]>;
    /**
     * Configuration block for the LDAP server used to authenticate and authorize connections to the broker. Not supported for `engineType` `RabbitMQ`. Detailed below. (Currently, AWS may not process changes to LDAP server metadata.)
     */
    readonly ldapServerMetadata?: pulumi.Input<inputs.mq.BrokerLdapServerMetadata>;
    /**
     * Configuration block for the logging configuration of the broker. Detailed below.
     */
    readonly logs?: pulumi.Input<inputs.mq.BrokerLogs>;
    /**
     * Configuration block for the maintenance window start time. Detailed below.
     */
    readonly maintenanceWindowStartTime?: pulumi.Input<inputs.mq.BrokerMaintenanceWindowStartTime>;
    /**
     * Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
     */
    readonly publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * List of security group IDs assigned to the broker.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Storage type of the broker. For `engineType` `ActiveMQ`, the valid values are `efs` and `ebs`, and the AWS-default is `efs`. For `engineType` `RabbitMQ`, only `ebs` is supported. When using `ebs`, only the `mq.m5` broker instance type family is supported.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * List of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires multiple subnets.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of tags to assign to the broker.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for broker users. For `engineType` of `RabbitMQ`, Amazon MQ does not return broker users preventing this resource from making user updates and drift detection. Detailed below.
     */
    readonly users?: pulumi.Input<pulumi.Input<inputs.mq.BrokerUser>[]>;
}

/**
 * The set of arguments for constructing a Broker resource.
 */
export interface BrokerArgs {
    /**
     * Specifies whether any broker modifications are applied immediately, or during the next maintenance window. Default is `false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * Authentication strategy used to secure the broker. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
     */
    readonly authenticationStrategy?: pulumi.Input<string>;
    /**
     * Whether to automatically upgrade to new minor versions of brokers as Amazon MQ makes releases available.
     */
    readonly autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * Name of the broker.
     */
    readonly brokerName: pulumi.Input<string>;
    /**
     * Configuration block for broker configuration. Applies to `engineType` of `ActiveMQ` only. Detailed below.
     */
    readonly configuration?: pulumi.Input<inputs.mq.BrokerConfiguration>;
    /**
     * Deployment mode of the broker. Valid values are `SINGLE_INSTANCE`, `ACTIVE_STANDBY_MULTI_AZ`, and `CLUSTER_MULTI_AZ`. Default is `SINGLE_INSTANCE`.
     */
    readonly deploymentMode?: pulumi.Input<string>;
    /**
     * Configuration block containing encryption options. Detailed below.
     */
    readonly encryptionOptions?: pulumi.Input<inputs.mq.BrokerEncryptionOptions>;
    /**
     * Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
     */
    readonly engineType: pulumi.Input<string>;
    /**
     * Version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions. For example, `5.15.0`.
     */
    readonly engineVersion: pulumi.Input<string>;
    /**
     * Broker's instance type. For example, `mq.t3.micro`, `mq.m5.large`.
     */
    readonly hostInstanceType: pulumi.Input<string>;
    /**
     * Configuration block for the LDAP server used to authenticate and authorize connections to the broker. Not supported for `engineType` `RabbitMQ`. Detailed below. (Currently, AWS may not process changes to LDAP server metadata.)
     */
    readonly ldapServerMetadata?: pulumi.Input<inputs.mq.BrokerLdapServerMetadata>;
    /**
     * Configuration block for the logging configuration of the broker. Detailed below.
     */
    readonly logs?: pulumi.Input<inputs.mq.BrokerLogs>;
    /**
     * Configuration block for the maintenance window start time. Detailed below.
     */
    readonly maintenanceWindowStartTime?: pulumi.Input<inputs.mq.BrokerMaintenanceWindowStartTime>;
    /**
     * Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
     */
    readonly publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * List of security group IDs assigned to the broker.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Storage type of the broker. For `engineType` `ActiveMQ`, the valid values are `efs` and `ebs`, and the AWS-default is `efs`. For `engineType` `RabbitMQ`, only `ebs` is supported. When using `ebs`, only the `mq.m5` broker instance type family is supported.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * List of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires multiple subnets.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of tags to assign to the broker.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for broker users. For `engineType` of `RabbitMQ`, Amazon MQ does not return broker users preventing this resource from making user updates and drift detection. Detailed below.
     */
    readonly users: pulumi.Input<pulumi.Input<inputs.mq.BrokerUser>[]>;
}
