// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an RDS DB proxy endpoint resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy-endpoints.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.rds.ProxyEndpoint("example", {
 *     dbProxyName: aws_db_proxy.test.name,
 *     dbProxyEndpointName: "example",
 *     vpcSubnetIds: aws_subnet.test.map(__item => __item.id),
 *     targetRole: "READ_ONLY",
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import DB proxy endpoints using the `DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME`. For example:
 *
 * ```sh
 *  $ pulumi import aws:rds/proxyEndpoint:ProxyEndpoint example example/example
 * ```
 */
export class ProxyEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing ProxyEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyEndpointState, opts?: pulumi.CustomResourceOptions): ProxyEndpoint {
        return new ProxyEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/proxyEndpoint:ProxyEndpoint';

    /**
     * Returns true if the given object is an instance of ProxyEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProxyEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProxyEndpoint.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) for the proxy endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
     */
    public readonly dbProxyEndpointName!: pulumi.Output<string>;
    /**
     * The name of the DB proxy associated with the DB proxy endpoint that you create.
     */
    public readonly dbProxyName!: pulumi.Output<string>;
    /**
     * The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Indicates whether this endpoint is the default endpoint for the associated DB proxy.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
     */
    public readonly targetRole!: pulumi.Output<string | undefined>;
    /**
     * The VPC ID of the DB proxy endpoint.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * One or more VPC security group IDs to associate with the new proxy.
     */
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[]>;
    /**
     * One or more VPC subnet IDs to associate with the new proxy.
     */
    public readonly vpcSubnetIds!: pulumi.Output<string[]>;

    /**
     * Create a ProxyEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyEndpointArgs | ProxyEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyEndpointState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["dbProxyEndpointName"] = state ? state.dbProxyEndpointName : undefined;
            resourceInputs["dbProxyName"] = state ? state.dbProxyName : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["targetRole"] = state ? state.targetRole : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcSecurityGroupIds"] = state ? state.vpcSecurityGroupIds : undefined;
            resourceInputs["vpcSubnetIds"] = state ? state.vpcSubnetIds : undefined;
        } else {
            const args = argsOrState as ProxyEndpointArgs | undefined;
            if ((!args || args.dbProxyEndpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbProxyEndpointName'");
            }
            if ((!args || args.dbProxyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbProxyName'");
            }
            if ((!args || args.vpcSubnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcSubnetIds'");
            }
            resourceInputs["dbProxyEndpointName"] = args ? args.dbProxyEndpointName : undefined;
            resourceInputs["dbProxyName"] = args ? args.dbProxyName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetRole"] = args ? args.targetRole : undefined;
            resourceInputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            resourceInputs["vpcSubnetIds"] = args ? args.vpcSubnetIds : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["isDefault"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProxyEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProxyEndpoint resources.
 */
export interface ProxyEndpointState {
    /**
     * The Amazon Resource Name (ARN) for the proxy endpoint.
     */
    arn?: pulumi.Input<string>;
    /**
     * The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
     */
    dbProxyEndpointName?: pulumi.Input<string>;
    /**
     * The name of the DB proxy associated with the DB proxy endpoint that you create.
     */
    dbProxyName?: pulumi.Input<string>;
    /**
     * The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Indicates whether this endpoint is the default endpoint for the associated DB proxy.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
     */
    targetRole?: pulumi.Input<string>;
    /**
     * The VPC ID of the DB proxy endpoint.
     */
    vpcId?: pulumi.Input<string>;
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
 * The set of arguments for constructing a ProxyEndpoint resource.
 */
export interface ProxyEndpointArgs {
    /**
     * The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
     */
    dbProxyEndpointName: pulumi.Input<string>;
    /**
     * The name of the DB proxy associated with the DB proxy endpoint that you create.
     */
    dbProxyName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
     */
    targetRole?: pulumi.Input<string>;
    /**
     * One or more VPC security group IDs to associate with the new proxy.
     */
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more VPC subnet IDs to associate with the new proxy.
     */
    vpcSubnetIds: pulumi.Input<pulumi.Input<string>[]>;
}
