// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS Kendra Data Source.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     description: "example",
 *     languageCode: "en",
 *     type: "CUSTOM",
 *     tags: {
 *         hello: "world",
 *     },
 * });
 * ```
 * ### S3 Connector
 * ### With Schedule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "S3",
 *     roleArn: aws_iam_role.example.arn,
 *     schedule: "cron(9 10 1 * ? *)",
 *     configuration: {
 *         s3Configuration: {
 *             bucketName: aws_s3_bucket.example.id,
 *         },
 *     },
 * });
 * ```
 * ### With Access Control List
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "S3",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         s3Configuration: {
 *             bucketName: aws_s3_bucket.example.id,
 *             accessControlListConfiguration: {
 *                 keyPath: `s3://${aws_s3_bucket.example.id}/path-1`,
 *             },
 *         },
 *     },
 * });
 * ```
 * ### With Documents Metadata Configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "S3",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         s3Configuration: {
 *             bucketName: aws_s3_bucket.example.id,
 *             exclusionPatterns: ["example"],
 *             inclusionPatterns: ["hello"],
 *             inclusionPrefixes: ["world"],
 *             documentsMetadataConfiguration: {
 *                 s3Prefix: "example",
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Web Crawler Connector
 * ### With Seed URLs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 * ### With Site Maps
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             urls: {
 *                 siteMapsConfiguration: {
 *                     siteMaps: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 * ### With Web Crawler Mode
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     webCrawlerMode: "SUBDOMAINS",
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 * ### With Authentication Configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             authenticationConfiguration: {
 *                 basicAuthentications: [{
 *                     credentials: aws_secretsmanager_secret.example.arn,
 *                     host: "a.example.com",
 *                     port: 443,
 *                 }],
 *             },
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * }, {
 *     dependsOn: [aws_secretsmanager_secret_version.example],
 * });
 * ```
 * ### With Crawl Depth
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             crawlDepth: 3,
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 * ### With Max Links Per Page
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             maxLinksPerPage: 100,
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 * ### With Max Urls Per Minute Crawl Rate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             maxUrlsPerMinuteCrawlRate: 300,
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 * ### With Proxy Configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             proxyConfiguration: {
 *                 credentials: aws_secretsmanager_secret.example.arn,
 *                 host: "a.example.com",
 *                 port: 443,
 *             },
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * }, {
 *     dependsOn: [aws_secretsmanager_secret_version.example],
 * });
 * ```
 * ### With URL Exclusion and Inclusion Patterns
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.kendra.DataSource("example", {
 *     indexId: aws_kendra_index.example.id,
 *     type: "WEBCRAWLER",
 *     roleArn: aws_iam_role.example.arn,
 *     configuration: {
 *         webCrawlerConfiguration: {
 *             urlExclusionPatterns: ["example"],
 *             urlInclusionPatterns: ["hello"],
 *             urls: {
 *                 seedUrlConfiguration: {
 *                     seedUrls: ["REPLACE_WITH_YOUR_URL"],
 *                 },
 *             },
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Kendra Data Source using the unique identifiers of the data_source and index separated by a slash (`/`). For example:
 *
 * ```sh
 *  $ pulumi import aws:kendra/dataSource:DataSource example 1045d08d-66ef-4882-b3ed-dfb7df183e90/b34dfdf7-1f2b-4704-9581-79e00296845f
 * ```
 */
export class DataSource extends pulumi.CustomResource {
    /**
     * Get an existing DataSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataSourceState, opts?: pulumi.CustomResourceOptions): DataSource {
        return new DataSource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:kendra/dataSource:DataSource';

    /**
     * Returns true if the given object is an instance of DataSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSource.__pulumiType;
    }

    /**
     * ARN of the Data Source.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A block with the configuration information to connect to your Data Source repository. You can't specify the `configuration` block when the `type` parameter is set to `CUSTOM`. Detailed below.
     */
    public readonly configuration!: pulumi.Output<outputs.kendra.DataSourceConfiguration | undefined>;
    /**
     * The Unix timestamp of when the Data Source was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A block with the configuration information for altering document metadata and content during the document ingestion process. For more information on how to create, modify and delete document metadata, or make other content alterations when you ingest documents into Amazon Kendra, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html). Detailed below.
     */
    public readonly customDocumentEnrichmentConfiguration!: pulumi.Output<outputs.kendra.DataSourceCustomDocumentEnrichmentConfiguration | undefined>;
    /**
     * The unique identifiers of the Data Source.
     */
    public /*out*/ readonly dataSourceId!: pulumi.Output<string>;
    /**
     * A description for the Data Source connector.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When the Status field value is `FAILED`, the ErrorMessage field contains a description of the error that caused the Data Source to fail.
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * The identifier of the index for your Amazon Kendra data source.
     */
    public readonly indexId!: pulumi.Output<string>;
    /**
     * The code for a language. This allows you to support a language for all documents when creating the Data Source connector. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
     */
    public readonly languageCode!: pulumi.Output<string>;
    /**
     * A name for your data source connector.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of a role with permission to access the data source connector. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html). You can't specify the `roleArn` parameter when the `type` parameter is set to `CUSTOM`. The `roleArn` parameter is required for all other data sources.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * Sets the frequency for Amazon Kendra to check the documents in your Data Source repository and update the index. If you don't set a schedule Amazon Kendra will not periodically update the index. You can call the `StartDataSourceSyncJob` API to update the index.
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * The current status of the Data Source. When the status is `ACTIVE` the Data Source is ready to use. When the status is `FAILED`, the `errorMessage` field contains the reason that the Data Source failed.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of data source repository. For an updated list of values, refer to [Valid Values for Type](https://docs.aws.amazon.com/kendra/latest/dg/API_CreateDataSource.html#Kendra-CreateDataSource-request-Type).
     *
     * The following arguments are optional:
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The Unix timestamp of when the Data Source was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a DataSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataSourceArgs | DataSourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataSourceState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["customDocumentEnrichmentConfiguration"] = state ? state.customDocumentEnrichmentConfiguration : undefined;
            resourceInputs["dataSourceId"] = state ? state.dataSourceId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["errorMessage"] = state ? state.errorMessage : undefined;
            resourceInputs["indexId"] = state ? state.indexId : undefined;
            resourceInputs["languageCode"] = state ? state.languageCode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as DataSourceArgs | undefined;
            if ((!args || args.indexId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'indexId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["customDocumentEnrichmentConfiguration"] = args ? args.customDocumentEnrichmentConfiguration : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["indexId"] = args ? args.indexId : undefined;
            resourceInputs["languageCode"] = args ? args.languageCode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dataSourceId"] = undefined /*out*/;
            resourceInputs["errorMessage"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DataSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSource resources.
 */
export interface DataSourceState {
    /**
     * ARN of the Data Source.
     */
    arn?: pulumi.Input<string>;
    /**
     * A block with the configuration information to connect to your Data Source repository. You can't specify the `configuration` block when the `type` parameter is set to `CUSTOM`. Detailed below.
     */
    configuration?: pulumi.Input<inputs.kendra.DataSourceConfiguration>;
    /**
     * The Unix timestamp of when the Data Source was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A block with the configuration information for altering document metadata and content during the document ingestion process. For more information on how to create, modify and delete document metadata, or make other content alterations when you ingest documents into Amazon Kendra, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html). Detailed below.
     */
    customDocumentEnrichmentConfiguration?: pulumi.Input<inputs.kendra.DataSourceCustomDocumentEnrichmentConfiguration>;
    /**
     * The unique identifiers of the Data Source.
     */
    dataSourceId?: pulumi.Input<string>;
    /**
     * A description for the Data Source connector.
     */
    description?: pulumi.Input<string>;
    /**
     * When the Status field value is `FAILED`, the ErrorMessage field contains a description of the error that caused the Data Source to fail.
     */
    errorMessage?: pulumi.Input<string>;
    /**
     * The identifier of the index for your Amazon Kendra data source.
     */
    indexId?: pulumi.Input<string>;
    /**
     * The code for a language. This allows you to support a language for all documents when creating the Data Source connector. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
     */
    languageCode?: pulumi.Input<string>;
    /**
     * A name for your data source connector.
     */
    name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of a role with permission to access the data source connector. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html). You can't specify the `roleArn` parameter when the `type` parameter is set to `CUSTOM`. The `roleArn` parameter is required for all other data sources.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Sets the frequency for Amazon Kendra to check the documents in your Data Source repository and update the index. If you don't set a schedule Amazon Kendra will not periodically update the index. You can call the `StartDataSourceSyncJob` API to update the index.
     */
    schedule?: pulumi.Input<string>;
    /**
     * The current status of the Data Source. When the status is `ACTIVE` the Data Source is ready to use. When the status is `FAILED`, the `errorMessage` field contains the reason that the Data Source failed.
     */
    status?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of data source repository. For an updated list of values, refer to [Valid Values for Type](https://docs.aws.amazon.com/kendra/latest/dg/API_CreateDataSource.html#Kendra-CreateDataSource-request-Type).
     *
     * The following arguments are optional:
     */
    type?: pulumi.Input<string>;
    /**
     * The Unix timestamp of when the Data Source was last updated.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * A block with the configuration information to connect to your Data Source repository. You can't specify the `configuration` block when the `type` parameter is set to `CUSTOM`. Detailed below.
     */
    configuration?: pulumi.Input<inputs.kendra.DataSourceConfiguration>;
    /**
     * A block with the configuration information for altering document metadata and content during the document ingestion process. For more information on how to create, modify and delete document metadata, or make other content alterations when you ingest documents into Amazon Kendra, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html). Detailed below.
     */
    customDocumentEnrichmentConfiguration?: pulumi.Input<inputs.kendra.DataSourceCustomDocumentEnrichmentConfiguration>;
    /**
     * A description for the Data Source connector.
     */
    description?: pulumi.Input<string>;
    /**
     * The identifier of the index for your Amazon Kendra data source.
     */
    indexId: pulumi.Input<string>;
    /**
     * The code for a language. This allows you to support a language for all documents when creating the Data Source connector. English is supported by default. For more information on supported languages, including their codes, see [Adding documents in languages other than English](https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
     */
    languageCode?: pulumi.Input<string>;
    /**
     * A name for your data source connector.
     */
    name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of a role with permission to access the data source connector. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html). You can't specify the `roleArn` parameter when the `type` parameter is set to `CUSTOM`. The `roleArn` parameter is required for all other data sources.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Sets the frequency for Amazon Kendra to check the documents in your Data Source repository and update the index. If you don't set a schedule Amazon Kendra will not periodically update the index. You can call the `StartDataSourceSyncJob` API to update the index.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of data source repository. For an updated list of values, refer to [Valid Values for Type](https://docs.aws.amazon.com/kendra/latest/dg/API_CreateDataSource.html#Kendra-CreateDataSource-request-Type).
     *
     * The following arguments are optional:
     */
    type: pulumi.Input<string>;
}
