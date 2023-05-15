// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for managing QuickSight Data Source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const _default = new aws.quicksight.DataSource("default", {
 *     dataSourceId: "example-id",
 *     parameters: {
 *         s3: {
 *             manifestFileLocation: {
 *                 bucket: "my-bucket",
 *                 key: "path/to/manifest.json",
 *             },
 *         },
 *     },
 *     type: "S3",
 * });
 * ```
 *
 * ## Import
 *
 * A QuickSight data source can be imported using the AWS account ID, and data source ID separated by a slash (`/`) e.g.,
 *
 * ```sh
 *  $ pulumi import aws:quicksight/dataSource:DataSource example 123456789123/my-data-source-id
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
    public static readonly __pulumiType = 'aws:quicksight/dataSource:DataSource';

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
     * Amazon Resource Name (ARN) of the data source
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    public readonly awsAccountId!: pulumi.Output<string>;
    /**
     * The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
     */
    public readonly credentials!: pulumi.Output<outputs.quicksight.DataSourceCredentials | undefined>;
    /**
     * An identifier for the data source.
     */
    public readonly dataSourceId!: pulumi.Output<string>;
    /**
     * A name for the data source, maximum of 128 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parameters used to connect to this data source (exactly one).
     */
    public readonly parameters!: pulumi.Output<outputs.quicksight.DataSourceParameters>;
    /**
     * A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
     */
    public readonly permissions!: pulumi.Output<outputs.quicksight.DataSourcePermission[] | undefined>;
    /**
     * Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
     */
    public readonly sslProperties!: pulumi.Output<outputs.quicksight.DataSourceSslProperties | undefined>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
     */
    public readonly vpcConnectionProperties!: pulumi.Output<outputs.quicksight.DataSourceVpcConnectionProperties | undefined>;

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
            resourceInputs["awsAccountId"] = state ? state.awsAccountId : undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["dataSourceId"] = state ? state.dataSourceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["sslProperties"] = state ? state.sslProperties : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vpcConnectionProperties"] = state ? state.vpcConnectionProperties : undefined;
        } else {
            const args = argsOrState as DataSourceArgs | undefined;
            if ((!args || args.dataSourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataSourceId'");
            }
            if ((!args || args.parameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameters'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["credentials"] = args ? args.credentials : undefined;
            resourceInputs["dataSourceId"] = args ? args.dataSourceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["sslProperties"] = args ? args.sslProperties : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vpcConnectionProperties"] = args ? args.vpcConnectionProperties : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSource resources.
 */
export interface DataSourceState {
    /**
     * Amazon Resource Name (ARN) of the data source
     */
    arn?: pulumi.Input<string>;
    /**
     * The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    awsAccountId?: pulumi.Input<string>;
    /**
     * The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
     */
    credentials?: pulumi.Input<inputs.quicksight.DataSourceCredentials>;
    /**
     * An identifier for the data source.
     */
    dataSourceId?: pulumi.Input<string>;
    /**
     * A name for the data source, maximum of 128 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The parameters used to connect to this data source (exactly one).
     */
    parameters?: pulumi.Input<inputs.quicksight.DataSourceParameters>;
    /**
     * A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.DataSourcePermission>[]>;
    /**
     * Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
     */
    sslProperties?: pulumi.Input<inputs.quicksight.DataSourceSslProperties>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
     */
    type?: pulumi.Input<string>;
    /**
     * Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
     */
    vpcConnectionProperties?: pulumi.Input<inputs.quicksight.DataSourceVpcConnectionProperties>;
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    awsAccountId?: pulumi.Input<string>;
    /**
     * The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
     */
    credentials?: pulumi.Input<inputs.quicksight.DataSourceCredentials>;
    /**
     * An identifier for the data source.
     */
    dataSourceId: pulumi.Input<string>;
    /**
     * A name for the data source, maximum of 128 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The parameters used to connect to this data source (exactly one).
     */
    parameters: pulumi.Input<inputs.quicksight.DataSourceParameters>;
    /**
     * A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.DataSourcePermission>[]>;
    /**
     * Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
     */
    sslProperties?: pulumi.Input<inputs.quicksight.DataSourceSslProperties>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.
     */
    type: pulumi.Input<string>;
    /**
     * Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
     */
    vpcConnectionProperties?: pulumi.Input<inputs.quicksight.DataSourceVpcConnectionProperties>;
}
