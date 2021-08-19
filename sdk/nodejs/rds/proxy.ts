// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an RDS DB proxy resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.rds.Proxy("example", {
 *     debugLogging: false,
 *     engineFamily: "MYSQL",
 *     idleClientTimeout: 1800,
 *     requireTls: true,
 *     roleArn: aws_iam_role.example.arn,
 *     vpcSecurityGroupIds: [aws_security_group.example.id],
 *     vpcSubnetIds: [aws_subnet.example.id],
 *     auths: [{
 *         authScheme: "SECRETS",
 *         description: "example",
 *         iamAuth: "DISABLED",
 *         secretArn: aws_secretsmanager_secret.example.arn,
 *     }],
 *     tags: {
 *         Name: "example",
 *         Key: "value",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * DB proxies can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:rds/proxy:Proxy example example
 * ```
 */
export class Proxy extends pulumi.CustomResource {
    /**
     * Get an existing Proxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyState, opts?: pulumi.CustomResourceOptions): Proxy {
        return new Proxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/proxy:Proxy';

    /**
     * Returns true if the given object is an instance of Proxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Proxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Proxy.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) for the proxy.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
     */
    public readonly auths!: pulumi.Output<outputs.rds.ProxyAuth[]>;
    /**
     * Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
     */
    public readonly debugLogging!: pulumi.Output<boolean | undefined>;
    /**
     * The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
     */
    public readonly engineFamily!: pulumi.Output<string>;
    /**
     * The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
     */
    public readonly idleClientTimeout!: pulumi.Output<number>;
    /**
     * The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
     */
    public readonly requireTls!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * One or more VPC security group IDs to associate with the new proxy.
     */
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[]>;
    /**
     * One or more VPC subnet IDs to associate with the new proxy.
     */
    public readonly vpcSubnetIds!: pulumi.Output<string[]>;

    /**
     * Create a Proxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyArgs | ProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["auths"] = state ? state.auths : undefined;
            inputs["debugLogging"] = state ? state.debugLogging : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["engineFamily"] = state ? state.engineFamily : undefined;
            inputs["idleClientTimeout"] = state ? state.idleClientTimeout : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["requireTls"] = state ? state.requireTls : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["vpcSecurityGroupIds"] = state ? state.vpcSecurityGroupIds : undefined;
            inputs["vpcSubnetIds"] = state ? state.vpcSubnetIds : undefined;
        } else {
            const args = argsOrState as ProxyArgs | undefined;
            if ((!args || args.auths === undefined) && !opts.urn) {
                throw new Error("Missing required property 'auths'");
            }
            if ((!args || args.engineFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineFamily'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.vpcSubnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcSubnetIds'");
            }
            inputs["auths"] = args ? args.auths : undefined;
            inputs["debugLogging"] = args ? args.debugLogging : undefined;
            inputs["engineFamily"] = args ? args.engineFamily : undefined;
            inputs["idleClientTimeout"] = args ? args.idleClientTimeout : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["requireTls"] = args ? args.requireTls : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            inputs["vpcSubnetIds"] = args ? args.vpcSubnetIds : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Proxy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Proxy resources.
 */
export interface ProxyState {
    /**
     * The Amazon Resource Name (ARN) for the proxy.
     */
    arn?: pulumi.Input<string>;
    /**
     * Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
     */
    auths?: pulumi.Input<pulumi.Input<inputs.rds.ProxyAuth>[]>;
    /**
     * Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
     */
    debugLogging?: pulumi.Input<boolean>;
    /**
     * The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
     */
    engineFamily?: pulumi.Input<string>;
    /**
     * The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
     */
    idleClientTimeout?: pulumi.Input<number>;
    /**
     * The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
     */
    name?: pulumi.Input<string>;
    /**
     * A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
     */
    requireTls?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * One or more VPC security group IDs to associate with the new proxy.
     */
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more VPC subnet IDs to associate with the new proxy.
     */
    vpcSubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Proxy resource.
 */
export interface ProxyArgs {
    /**
     * Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
     */
    auths: pulumi.Input<pulumi.Input<inputs.rds.ProxyAuth>[]>;
    /**
     * Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
     */
    debugLogging?: pulumi.Input<boolean>;
    /**
     * The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
     */
    engineFamily: pulumi.Input<string>;
    /**
     * The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
     */
    idleClientTimeout?: pulumi.Input<number>;
    /**
     * The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
     */
    name?: pulumi.Input<string>;
    /**
     * A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
     */
    requireTls?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
     */
    roleArn: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * One or more VPC security group IDs to associate with the new proxy.
     */
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more VPC subnet IDs to associate with the new proxy.
     */
    vpcSubnetIds: pulumi.Input<pulumi.Input<string>[]>;
}
