// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a DynamoDB table resource.
 *
 * > **Note:** It is recommended to use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) for `readCapacity` and/or `writeCapacity` if there's `autoscaling policy` attached to the table.
 *
 * > **Note:** When using aws.dynamodb.TableReplica with this resource, use `lifecycle` `ignoreChanges` for `replica`, _e.g._, `lifecycle { ignoreChanges = [replica] }`.
 *
 * ## DynamoDB Table attributes
 *
 * Only define attributes on the table object that are going to be used as:
 *
 * * Table hash key or range key
 * * LSI or GSI hash key or range key
 *
 * The DynamoDB API expects attribute structure (name and type) to be passed along when creating or updating GSI/LSIs or creating the initial table. In these cases it expects the Hash / Range keys to be provided. Because these get re-used in numerous places (i.e the table's range key could be a part of one or more GSIs), they are stored on the table object to prevent duplication and increase consistency. If you add attributes here that are not used in these scenarios it can cause an infinite loop in planning.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * The following dynamodb table description models the table and GSI shown in the [AWS SDK example documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const basic_dynamodb_table = new aws.dynamodb.Table("basic-dynamodb-table", {
 *     attributes: [
 *         {
 *             name: "UserId",
 *             type: "S",
 *         },
 *         {
 *             name: "GameTitle",
 *             type: "S",
 *         },
 *         {
 *             name: "TopScore",
 *             type: "N",
 *         },
 *     ],
 *     billingMode: "PROVISIONED",
 *     globalSecondaryIndexes: [{
 *         hashKey: "GameTitle",
 *         name: "GameTitleIndex",
 *         nonKeyAttributes: ["UserId"],
 *         projectionType: "INCLUDE",
 *         rangeKey: "TopScore",
 *         readCapacity: 10,
 *         writeCapacity: 10,
 *     }],
 *     hashKey: "UserId",
 *     rangeKey: "GameTitle",
 *     readCapacity: 20,
 *     tags: {
 *         Environment: "production",
 *         Name: "dynamodb-table-1",
 *     },
 *     ttl: {
 *         attributeName: "TimeToExist",
 *         enabled: false,
 *     },
 *     writeCapacity: 20,
 * });
 * ```
 * ### Global Tables
 *
 * This resource implements support for [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) via `replica` configuration blocks. For working with [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html), see the `aws.dynamodb.GlobalTable` resource.
 *
 * > **Note:** aws.dynamodb.TableReplica is an alternate way of configuring Global Tables. Do not use `replica` configuration blocks of `aws.dynamodb.Table` together with aws_dynamodb_table_replica.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.dynamodb.Table("example", {
 *     attributes: [{
 *         name: "TestTableHashKey",
 *         type: "S",
 *     }],
 *     billingMode: "PAY_PER_REQUEST",
 *     hashKey: "TestTableHashKey",
 *     replicas: [
 *         {
 *             regionName: "us-east-2",
 *         },
 *         {
 *             regionName: "us-west-2",
 *         },
 *     ],
 *     streamEnabled: true,
 *     streamViewType: "NEW_AND_OLD_IMAGES",
 * });
 * ```
 *
 * ## Import
 *
 * DynamoDB tables can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:dynamodb/table:Table basic-dynamodb-table GameScores
 * ```
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:dynamodb/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * ARN of the table
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Set of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. See below.
     */
    public readonly attributes!: pulumi.Output<outputs.dynamodb.TableAttribute[]>;
    /**
     * Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
     */
    public readonly billingMode!: pulumi.Output<string | undefined>;
    /**
     * Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
     */
    public readonly globalSecondaryIndexes!: pulumi.Output<outputs.dynamodb.TableGlobalSecondaryIndex[] | undefined>;
    /**
     * Name of the hash key in the index; must be defined as an attribute in the resource.
     */
    public readonly hashKey!: pulumi.Output<string>;
    /**
     * Describe an LSI on the table; these can only be allocated *at creation* so you cannot change this definition after you have created the resource. See below.
     */
    public readonly localSecondaryIndexes!: pulumi.Output<outputs.dynamodb.TableLocalSecondaryIndex[] | undefined>;
    /**
     * Name of the index
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether to enable Point In Time Recovery for the replica. Default is `false`.
     */
    public readonly pointInTimeRecovery!: pulumi.Output<outputs.dynamodb.TablePointInTimeRecovery>;
    /**
     * Name of the range key.
     */
    public readonly rangeKey!: pulumi.Output<string | undefined>;
    /**
     * Number of read units for this index. Must be set if billingMode is set to PROVISIONED.
     */
    public readonly readCapacity!: pulumi.Output<number>;
    /**
     * Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. See below.
     */
    public readonly replicas!: pulumi.Output<outputs.dynamodb.TableReplica[] | undefined>;
    /**
     * Time of the point-in-time recovery point to restore.
     */
    public readonly restoreDateTime!: pulumi.Output<string | undefined>;
    /**
     * Name of the table to restore. Must match the name of an existing table.
     */
    public readonly restoreSourceName!: pulumi.Output<string | undefined>;
    /**
     * If set, restores table to the most recent point-in-time recovery point.
     */
    public readonly restoreToLatestTime!: pulumi.Output<boolean | undefined>;
    /**
     * Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. See below.
     */
    public readonly serverSideEncryption!: pulumi.Output<outputs.dynamodb.TableServerSideEncryption>;
    /**
     * ARN of the Table Stream. Only available when `streamEnabled = true`
     */
    public /*out*/ readonly streamArn!: pulumi.Output<string>;
    /**
     * Whether Streams are enabled.
     */
    public readonly streamEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
     */
    public /*out*/ readonly streamLabel!: pulumi.Output<string>;
    /**
     * When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
     */
    public readonly streamViewType!: pulumi.Output<string>;
    /**
     * Storage class of the table. Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS`.
     */
    public readonly tableClass!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to populate on the created table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Configuration block for TTL. See below.
     */
    public readonly ttl!: pulumi.Output<outputs.dynamodb.TableTtl>;
    /**
     * Number of write units for this index. Must be set if billingMode is set to PROVISIONED.
     */
    public readonly writeCapacity!: pulumi.Output<number>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["billingMode"] = state ? state.billingMode : undefined;
            resourceInputs["globalSecondaryIndexes"] = state ? state.globalSecondaryIndexes : undefined;
            resourceInputs["hashKey"] = state ? state.hashKey : undefined;
            resourceInputs["localSecondaryIndexes"] = state ? state.localSecondaryIndexes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pointInTimeRecovery"] = state ? state.pointInTimeRecovery : undefined;
            resourceInputs["rangeKey"] = state ? state.rangeKey : undefined;
            resourceInputs["readCapacity"] = state ? state.readCapacity : undefined;
            resourceInputs["replicas"] = state ? state.replicas : undefined;
            resourceInputs["restoreDateTime"] = state ? state.restoreDateTime : undefined;
            resourceInputs["restoreSourceName"] = state ? state.restoreSourceName : undefined;
            resourceInputs["restoreToLatestTime"] = state ? state.restoreToLatestTime : undefined;
            resourceInputs["serverSideEncryption"] = state ? state.serverSideEncryption : undefined;
            resourceInputs["streamArn"] = state ? state.streamArn : undefined;
            resourceInputs["streamEnabled"] = state ? state.streamEnabled : undefined;
            resourceInputs["streamLabel"] = state ? state.streamLabel : undefined;
            resourceInputs["streamViewType"] = state ? state.streamViewType : undefined;
            resourceInputs["tableClass"] = state ? state.tableClass : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["writeCapacity"] = state ? state.writeCapacity : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["billingMode"] = args ? args.billingMode : undefined;
            resourceInputs["globalSecondaryIndexes"] = args ? args.globalSecondaryIndexes : undefined;
            resourceInputs["hashKey"] = args ? args.hashKey : undefined;
            resourceInputs["localSecondaryIndexes"] = args ? args.localSecondaryIndexes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pointInTimeRecovery"] = args ? args.pointInTimeRecovery : undefined;
            resourceInputs["rangeKey"] = args ? args.rangeKey : undefined;
            resourceInputs["readCapacity"] = args ? args.readCapacity : undefined;
            resourceInputs["replicas"] = args ? args.replicas : undefined;
            resourceInputs["restoreDateTime"] = args ? args.restoreDateTime : undefined;
            resourceInputs["restoreSourceName"] = args ? args.restoreSourceName : undefined;
            resourceInputs["restoreToLatestTime"] = args ? args.restoreToLatestTime : undefined;
            resourceInputs["serverSideEncryption"] = args ? args.serverSideEncryption : undefined;
            resourceInputs["streamEnabled"] = args ? args.streamEnabled : undefined;
            resourceInputs["streamViewType"] = args ? args.streamViewType : undefined;
            resourceInputs["tableClass"] = args ? args.tableClass : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["writeCapacity"] = args ? args.writeCapacity : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["streamArn"] = undefined /*out*/;
            resourceInputs["streamLabel"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * ARN of the table
     */
    arn?: pulumi.Input<string>;
    /**
     * Set of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. See below.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableAttribute>[]>;
    /**
     * Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
     */
    billingMode?: pulumi.Input<string>;
    /**
     * Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
     */
    globalSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableGlobalSecondaryIndex>[]>;
    /**
     * Name of the hash key in the index; must be defined as an attribute in the resource.
     */
    hashKey?: pulumi.Input<string>;
    /**
     * Describe an LSI on the table; these can only be allocated *at creation* so you cannot change this definition after you have created the resource. See below.
     */
    localSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableLocalSecondaryIndex>[]>;
    /**
     * Name of the index
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to enable Point In Time Recovery for the replica. Default is `false`.
     */
    pointInTimeRecovery?: pulumi.Input<inputs.dynamodb.TablePointInTimeRecovery>;
    /**
     * Name of the range key.
     */
    rangeKey?: pulumi.Input<string>;
    /**
     * Number of read units for this index. Must be set if billingMode is set to PROVISIONED.
     */
    readCapacity?: pulumi.Input<number>;
    /**
     * Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. See below.
     */
    replicas?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableReplica>[]>;
    /**
     * Time of the point-in-time recovery point to restore.
     */
    restoreDateTime?: pulumi.Input<string>;
    /**
     * Name of the table to restore. Must match the name of an existing table.
     */
    restoreSourceName?: pulumi.Input<string>;
    /**
     * If set, restores table to the most recent point-in-time recovery point.
     */
    restoreToLatestTime?: pulumi.Input<boolean>;
    /**
     * Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. See below.
     */
    serverSideEncryption?: pulumi.Input<inputs.dynamodb.TableServerSideEncryption>;
    /**
     * ARN of the Table Stream. Only available when `streamEnabled = true`
     */
    streamArn?: pulumi.Input<string>;
    /**
     * Whether Streams are enabled.
     */
    streamEnabled?: pulumi.Input<boolean>;
    /**
     * Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
     */
    streamLabel?: pulumi.Input<string>;
    /**
     * When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
     */
    streamViewType?: pulumi.Input<string>;
    /**
     * Storage class of the table. Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS`.
     */
    tableClass?: pulumi.Input<string>;
    /**
     * A map of tags to populate on the created table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for TTL. See below.
     */
    ttl?: pulumi.Input<inputs.dynamodb.TableTtl>;
    /**
     * Number of write units for this index. Must be set if billingMode is set to PROVISIONED.
     */
    writeCapacity?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * Set of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. See below.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableAttribute>[]>;
    /**
     * Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
     */
    billingMode?: pulumi.Input<string>;
    /**
     * Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
     */
    globalSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableGlobalSecondaryIndex>[]>;
    /**
     * Name of the hash key in the index; must be defined as an attribute in the resource.
     */
    hashKey?: pulumi.Input<string>;
    /**
     * Describe an LSI on the table; these can only be allocated *at creation* so you cannot change this definition after you have created the resource. See below.
     */
    localSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableLocalSecondaryIndex>[]>;
    /**
     * Name of the index
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to enable Point In Time Recovery for the replica. Default is `false`.
     */
    pointInTimeRecovery?: pulumi.Input<inputs.dynamodb.TablePointInTimeRecovery>;
    /**
     * Name of the range key.
     */
    rangeKey?: pulumi.Input<string>;
    /**
     * Number of read units for this index. Must be set if billingMode is set to PROVISIONED.
     */
    readCapacity?: pulumi.Input<number>;
    /**
     * Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. See below.
     */
    replicas?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableReplica>[]>;
    /**
     * Time of the point-in-time recovery point to restore.
     */
    restoreDateTime?: pulumi.Input<string>;
    /**
     * Name of the table to restore. Must match the name of an existing table.
     */
    restoreSourceName?: pulumi.Input<string>;
    /**
     * If set, restores table to the most recent point-in-time recovery point.
     */
    restoreToLatestTime?: pulumi.Input<boolean>;
    /**
     * Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. See below.
     */
    serverSideEncryption?: pulumi.Input<inputs.dynamodb.TableServerSideEncryption>;
    /**
     * Whether Streams are enabled.
     */
    streamEnabled?: pulumi.Input<boolean>;
    /**
     * When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
     */
    streamViewType?: pulumi.Input<string>;
    /**
     * Storage class of the table. Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS`.
     */
    tableClass?: pulumi.Input<string>;
    /**
     * A map of tags to populate on the created table. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for TTL. See below.
     */
    ttl?: pulumi.Input<inputs.dynamodb.TableTtl>;
    /**
     * Number of write units for this index. Must be set if billingMode is set to PROVISIONED.
     */
    writeCapacity?: pulumi.Input<number>;
}
