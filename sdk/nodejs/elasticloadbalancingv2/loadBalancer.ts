// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer resource.
 *
 * > **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.
 *
 * ## Example Usage
 * ### Application Load Balancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.lb.LoadBalancer("test", {
 *     internal: false,
 *     loadBalancerType: "application",
 *     securityGroups: [aws_security_group.lb_sg.id],
 *     subnets: .map(subnet => subnet.id),
 *     enableDeletionProtection: true,
 *     accessLogs: {
 *         bucket: aws_s3_bucket.lb_logs.bucket,
 *         prefix: "test-lb",
 *         enabled: true,
 *     },
 *     tags: {
 *         Environment: "production",
 *     },
 * });
 * ```
 * ### Network Load Balancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.lb.LoadBalancer("test", {
 *     internal: false,
 *     loadBalancerType: "network",
 *     subnets: .map(subnet => subnet.id),
 *     enableDeletionProtection: true,
 *     tags: {
 *         Environment: "production",
 *     },
 * });
 * ```
 * ### Specifying Elastic IPs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lb.LoadBalancer("example", {
 *     loadBalancerType: "network",
 *     subnetMappings: [
 *         {
 *             subnetId: aws_subnet.example1.id,
 *             allocationId: aws_eip.example1.id,
 *         },
 *         {
 *             subnetId: aws_subnet.example2.id,
 *             allocationId: aws_eip.example2.id,
 *         },
 *     ],
 * });
 * ```
 * ### Specifying private IP addresses for an internal-facing load balancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lb.LoadBalancer("example", {
 *     loadBalancerType: "network",
 *     subnetMappings: [
 *         {
 *             subnetId: aws_subnet.example1.id,
 *             privateIpv4Address: "10.0.1.15",
 *         },
 *         {
 *             subnetId: aws_subnet.example2.id,
 *             privateIpv4Address: "10.0.2.15",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * LBs can be imported using their ARN, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:elasticloadbalancingv2/loadBalancer:LoadBalancer bar arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188
 * ```
 *
 * @deprecated aws.elasticloadbalancingv2.LoadBalancer has been deprecated in favor of aws.lb.LoadBalancer
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        pulumi.log.warn("LoadBalancer is deprecated: aws.elasticloadbalancingv2.LoadBalancer has been deprecated in favor of aws.lb.LoadBalancer")
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticloadbalancingv2/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * An Access Logs block. Access Logs documented below.
     */
    public readonly accessLogs!: pulumi.Output<outputs.elasticloadbalancingv2.LoadBalancerAccessLogs | undefined>;
    /**
     * The ARN of the load balancer (matches `id`).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ARN suffix for use with CloudWatch Metrics.
     */
    public /*out*/ readonly arnSuffix!: pulumi.Output<string>;
    /**
     * The ID of the customer owned ipv4 pool to use for this load balancer.
     */
    public readonly customerOwnedIpv4Pool!: pulumi.Output<string | undefined>;
    /**
     * Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
     */
    public readonly desyncMitigationMode!: pulumi.Output<string | undefined>;
    /**
     * The DNS name of the load balancer.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     */
    public readonly dropInvalidHeaderFields!: pulumi.Output<boolean | undefined>;
    /**
     * If true, cross-zone load balancing of the load balancer will be enabled. For `network` and `gateway` type load balancers, this feature is disabled by default (`false`). For `application` load balancer this feature is always enabled (`true`) and cannot be disabled. Defaults to `false`.
     */
    public readonly enableCrossZoneLoadBalancing!: pulumi.Output<boolean | undefined>;
    /**
     * If true, deletion of the load balancer will be disabled via the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     */
    public readonly enableDeletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     */
    public readonly enableHttp2!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
     */
    public readonly enableWafFailOpen!: pulumi.Output<boolean | undefined>;
    /**
     * The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     */
    public readonly idleTimeout!: pulumi.Output<number | undefined>;
    /**
     * If true, the LB will be internal.
     */
    public readonly internal!: pulumi.Output<boolean>;
    /**
     * The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`.
     */
    public readonly ipAddressType!: pulumi.Output<string>;
    /**
     * The type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
     */
    public readonly loadBalancerType!: pulumi.Output<string | undefined>;
    /**
     * The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the Application Load Balancer should preserve the Host header in the HTTP request and send it to the target without any change. Defaults to `false`.
     */
    public readonly preserveHostHeader!: pulumi.Output<boolean | undefined>;
    /**
     * A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * A subnet mapping block as documented below.
     */
    public readonly subnetMappings!: pulumi.Output<outputs.elasticloadbalancingv2.LoadBalancerSubnetMapping[]>;
    /**
     * A list of subnet IDs to attach to the LB. Subnets
     * cannot be updated for Load Balancers of type `network`. Changing this value
     * for load balancers of type `network` will force a recreation of the resource.
     */
    public readonly subnets!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
     */
    public /*out*/ readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated aws.elasticloadbalancingv2.LoadBalancer has been deprecated in favor of aws.lb.LoadBalancer */
    constructor(name: string, args?: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated aws.elasticloadbalancingv2.LoadBalancer has been deprecated in favor of aws.lb.LoadBalancer */
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LoadBalancer is deprecated: aws.elasticloadbalancingv2.LoadBalancer has been deprecated in favor of aws.lb.LoadBalancer")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            resourceInputs["accessLogs"] = state ? state.accessLogs : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["arnSuffix"] = state ? state.arnSuffix : undefined;
            resourceInputs["customerOwnedIpv4Pool"] = state ? state.customerOwnedIpv4Pool : undefined;
            resourceInputs["desyncMitigationMode"] = state ? state.desyncMitigationMode : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["dropInvalidHeaderFields"] = state ? state.dropInvalidHeaderFields : undefined;
            resourceInputs["enableCrossZoneLoadBalancing"] = state ? state.enableCrossZoneLoadBalancing : undefined;
            resourceInputs["enableDeletionProtection"] = state ? state.enableDeletionProtection : undefined;
            resourceInputs["enableHttp2"] = state ? state.enableHttp2 : undefined;
            resourceInputs["enableWafFailOpen"] = state ? state.enableWafFailOpen : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["internal"] = state ? state.internal : undefined;
            resourceInputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            resourceInputs["loadBalancerType"] = state ? state.loadBalancerType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["preserveHostHeader"] = state ? state.preserveHostHeader : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["subnetMappings"] = state ? state.subnetMappings : undefined;
            resourceInputs["subnets"] = state ? state.subnets : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            resourceInputs["accessLogs"] = args ? args.accessLogs : undefined;
            resourceInputs["customerOwnedIpv4Pool"] = args ? args.customerOwnedIpv4Pool : undefined;
            resourceInputs["desyncMitigationMode"] = args ? args.desyncMitigationMode : undefined;
            resourceInputs["dropInvalidHeaderFields"] = args ? args.dropInvalidHeaderFields : undefined;
            resourceInputs["enableCrossZoneLoadBalancing"] = args ? args.enableCrossZoneLoadBalancing : undefined;
            resourceInputs["enableDeletionProtection"] = args ? args.enableDeletionProtection : undefined;
            resourceInputs["enableHttp2"] = args ? args.enableHttp2 : undefined;
            resourceInputs["enableWafFailOpen"] = args ? args.enableWafFailOpen : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["internal"] = args ? args.internal : undefined;
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["loadBalancerType"] = args ? args.loadBalancerType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["preserveHostHeader"] = args ? args.preserveHostHeader : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["subnetMappings"] = args ? args.subnetMappings : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["arnSuffix"] = undefined /*out*/;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
            resourceInputs["zoneId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * An Access Logs block. Access Logs documented below.
     */
    accessLogs?: pulumi.Input<inputs.elasticloadbalancingv2.LoadBalancerAccessLogs>;
    /**
     * The ARN of the load balancer (matches `id`).
     */
    arn?: pulumi.Input<string>;
    /**
     * The ARN suffix for use with CloudWatch Metrics.
     */
    arnSuffix?: pulumi.Input<string>;
    /**
     * The ID of the customer owned ipv4 pool to use for this load balancer.
     */
    customerOwnedIpv4Pool?: pulumi.Input<string>;
    /**
     * Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
     */
    desyncMitigationMode?: pulumi.Input<string>;
    /**
     * The DNS name of the load balancer.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     */
    dropInvalidHeaderFields?: pulumi.Input<boolean>;
    /**
     * If true, cross-zone load balancing of the load balancer will be enabled. For `network` and `gateway` type load balancers, this feature is disabled by default (`false`). For `application` load balancer this feature is always enabled (`true`) and cannot be disabled. Defaults to `false`.
     */
    enableCrossZoneLoadBalancing?: pulumi.Input<boolean>;
    /**
     * If true, deletion of the load balancer will be disabled via the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     */
    enableDeletionProtection?: pulumi.Input<boolean>;
    /**
     * Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     */
    enableHttp2?: pulumi.Input<boolean>;
    /**
     * Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
     */
    enableWafFailOpen?: pulumi.Input<boolean>;
    /**
     * The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * If true, the LB will be internal.
     */
    internal?: pulumi.Input<boolean>;
    /**
     * The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * The type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
     */
    loadBalancerType?: pulumi.Input<string>;
    /**
     * The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Indicates whether the Application Load Balancer should preserve the Host header in the HTTP request and send it to the target without any change. Defaults to `false`.
     */
    preserveHostHeader?: pulumi.Input<boolean>;
    /**
     * A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A subnet mapping block as documented below.
     */
    subnetMappings?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.LoadBalancerSubnetMapping>[]>;
    /**
     * A list of subnet IDs to attach to the LB. Subnets
     * cannot be updated for Load Balancers of type `network`. Changing this value
     * for load balancers of type `network` will force a recreation of the resource.
     */
    subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    vpcId?: pulumi.Input<string>;
    /**
     * The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * An Access Logs block. Access Logs documented below.
     */
    accessLogs?: pulumi.Input<inputs.elasticloadbalancingv2.LoadBalancerAccessLogs>;
    /**
     * The ID of the customer owned ipv4 pool to use for this load balancer.
     */
    customerOwnedIpv4Pool?: pulumi.Input<string>;
    /**
     * Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
     */
    desyncMitigationMode?: pulumi.Input<string>;
    /**
     * Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     */
    dropInvalidHeaderFields?: pulumi.Input<boolean>;
    /**
     * If true, cross-zone load balancing of the load balancer will be enabled. For `network` and `gateway` type load balancers, this feature is disabled by default (`false`). For `application` load balancer this feature is always enabled (`true`) and cannot be disabled. Defaults to `false`.
     */
    enableCrossZoneLoadBalancing?: pulumi.Input<boolean>;
    /**
     * If true, deletion of the load balancer will be disabled via the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     */
    enableDeletionProtection?: pulumi.Input<boolean>;
    /**
     * Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     */
    enableHttp2?: pulumi.Input<boolean>;
    /**
     * Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
     */
    enableWafFailOpen?: pulumi.Input<boolean>;
    /**
     * The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * If true, the LB will be internal.
     */
    internal?: pulumi.Input<boolean>;
    /**
     * The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * The type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
     */
    loadBalancerType?: pulumi.Input<string>;
    /**
     * The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Indicates whether the Application Load Balancer should preserve the Host header in the HTTP request and send it to the target without any change. Defaults to `false`.
     */
    preserveHostHeader?: pulumi.Input<boolean>;
    /**
     * A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A subnet mapping block as documented below.
     */
    subnetMappings?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancingv2.LoadBalancerSubnetMapping>[]>;
    /**
     * A list of subnet IDs to attach to the LB. Subnets
     * cannot be updated for Load Balancers of type `network`. Changing this value
     * for load balancers of type `network` will force a recreation of the resource.
     */
    subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
