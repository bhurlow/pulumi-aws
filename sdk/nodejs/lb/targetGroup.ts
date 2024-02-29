// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a Target Group resource for use with Load Balancer resources.
 *
 * > **Note:** `aws.alb.TargetGroup` is known as `aws.lb.TargetGroup`. The functionality is identical.
 *
 * ## Example Usage
 * ### Instance Target Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
 * const test = new aws.lb.TargetGroup("test", {
 *     port: 80,
 *     protocol: "HTTP",
 *     vpcId: main.id,
 * });
 * ```
 * ### IP Target Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
 * const ip_example = new aws.lb.TargetGroup("ip-example", {
 *     port: 80,
 *     protocol: "HTTP",
 *     targetType: "ip",
 *     vpcId: main.id,
 * });
 * ```
 * ### Lambda Target Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const lambda_example = new aws.lb.TargetGroup("lambda-example", {targetType: "lambda"});
 * ```
 * ### ALB Target Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const alb_example = new aws.lb.TargetGroup("alb-example", {
 *     targetType: "alb",
 *     port: 80,
 *     protocol: "TCP",
 *     vpcId: aws_vpc.main.id,
 * });
 * ```
 * ### Target group with unhealthy connection termination disabled
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const tcp_example = new aws.lb.TargetGroup("tcp-example", {
 *     port: 25,
 *     protocol: "TCP",
 *     vpcId: aws_vpc.main.id,
 *     targetHealthStates: [{
 *         enableUnhealthyConnectionTermination: false,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Target Groups using their ARN. For example:
 *
 * ```sh
 *  $ pulumi import aws:lb/targetGroup:TargetGroup app_front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:targetgroup/app-front-end/20cfe21448b66314
 * ```
 */
export class TargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing TargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetGroupState, opts?: pulumi.CustomResourceOptions): TargetGroup {
        return new TargetGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lb/targetGroup:TargetGroup';

    /**
     * Returns true if the given object is an instance of TargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetGroup.__pulumiType;
    }

    /**
     * ARN of the Target Group (matches `id`).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ARN suffix for use with CloudWatch Metrics.
     */
    public /*out*/ readonly arnSuffix!: pulumi.Output<string>;
    /**
     * Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
     */
    public readonly connectionTermination!: pulumi.Output<boolean>;
    /**
     * Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
     */
    public readonly deregistrationDelay!: pulumi.Output<number | undefined>;
    /**
     * Health Check configuration block. Detailed below.
     */
    public readonly healthCheck!: pulumi.Output<outputs.lb.TargetGroupHealthCheck>;
    /**
     * The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
     */
    public readonly ipAddressType!: pulumi.Output<string>;
    /**
     * Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
     */
    public readonly lambdaMultiValueHeadersEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * ARNs of the Load Balancers associated with the Target Group.
     */
    public /*out*/ readonly loadBalancerArns!: pulumi.Output<string[]>;
    /**
     * Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin`, `leastOutstandingRequests`, or `weightedRandom`. The default is `roundRobin`.
     */
    public readonly loadBalancingAlgorithmType!: pulumi.Output<string>;
    /**
     * Determines whether to enable target anomaly mitigation.  Target anomaly mitigation is only supported by the `weightedRandom` load balancing algorithm type.  See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#automatic-target-weights) for more information.  The value is `"on"` or `"off"`. The default is `"off"`.
     */
    public readonly loadBalancingAnomalyMitigation!: pulumi.Output<string>;
    /**
     * Indicates whether cross zone load balancing is enabled. The value is `"true"`, `"false"` or `"useLoadBalancerConfiguration"`. The default is `"useLoadBalancerConfiguration"`.
     */
    public readonly loadBalancingCrossZoneEnabled!: pulumi.Output<string>;
    /**
     * Name of the target group. If omitted, this provider will assign a random, unique name. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * Port on which targets receive traffic, unless overridden when registering a specific target. Required when `targetType` is `instance`, `ip` or `alb`. Does not apply when `targetType` is `lambda`.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
     */
    public readonly preserveClientIp!: pulumi.Output<string>;
    /**
     * Protocol to use for routing traffic to the targets.
     * Should be one of `GENEVE`, `HTTP`, `HTTPS`, `TCP`, `TCP_UDP`, `TLS`, or `UDP`.
     * Required when `targetType` is `instance`, `ip`, or `alb`.
     * Does not apply when `targetType` is `lambda`.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
     */
    public readonly protocolVersion!: pulumi.Output<string>;
    /**
     * Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
     */
    public readonly proxyProtocolV2!: pulumi.Output<boolean | undefined>;
    /**
     * Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
     */
    public readonly slowStart!: pulumi.Output<number | undefined>;
    /**
     * Stickiness configuration block. Detailed below.
     */
    public readonly stickiness!: pulumi.Output<outputs.lb.TargetGroupStickiness>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Target failover block. Only applicable for Gateway Load Balancer target groups. See targetFailover for more information.
     */
    public readonly targetFailovers!: pulumi.Output<outputs.lb.TargetGroupTargetFailover[]>;
    /**
     * Target health state block. Only applicable for Network Load Balancer target groups when `protocol` is `TCP` or `TLS`. See targetHealthState for more information.
     */
    public readonly targetHealthStates!: pulumi.Output<outputs.lb.TargetGroupTargetHealthState[]>;
    /**
     * Type of target that you must specify when registering targets with this target group.
     * See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values.
     * The default is `instance`.
     *
     * Note that you can't specify targets for a target group using both instance IDs and IP addresses.
     *
     * If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
     *
     * Network Load Balancers do not support the `lambda` target type.
     *
     * Application Load Balancers do not support the `alb` target type.
     */
    public readonly targetType!: pulumi.Output<string | undefined>;
    /**
     * Identifier of the VPC in which to create the target group. Required when `targetType` is `instance`, `ip` or `alb`. Does not apply when `targetType` is `lambda`.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a TargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TargetGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetGroupArgs | TargetGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TargetGroupState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["arnSuffix"] = state ? state.arnSuffix : undefined;
            resourceInputs["connectionTermination"] = state ? state.connectionTermination : undefined;
            resourceInputs["deregistrationDelay"] = state ? state.deregistrationDelay : undefined;
            resourceInputs["healthCheck"] = state ? state.healthCheck : undefined;
            resourceInputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            resourceInputs["lambdaMultiValueHeadersEnabled"] = state ? state.lambdaMultiValueHeadersEnabled : undefined;
            resourceInputs["loadBalancerArns"] = state ? state.loadBalancerArns : undefined;
            resourceInputs["loadBalancingAlgorithmType"] = state ? state.loadBalancingAlgorithmType : undefined;
            resourceInputs["loadBalancingAnomalyMitigation"] = state ? state.loadBalancingAnomalyMitigation : undefined;
            resourceInputs["loadBalancingCrossZoneEnabled"] = state ? state.loadBalancingCrossZoneEnabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["preserveClientIp"] = state ? state.preserveClientIp : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["protocolVersion"] = state ? state.protocolVersion : undefined;
            resourceInputs["proxyProtocolV2"] = state ? state.proxyProtocolV2 : undefined;
            resourceInputs["slowStart"] = state ? state.slowStart : undefined;
            resourceInputs["stickiness"] = state ? state.stickiness : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["targetFailovers"] = state ? state.targetFailovers : undefined;
            resourceInputs["targetHealthStates"] = state ? state.targetHealthStates : undefined;
            resourceInputs["targetType"] = state ? state.targetType : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as TargetGroupArgs | undefined;
            resourceInputs["connectionTermination"] = args ? args.connectionTermination : undefined;
            resourceInputs["deregistrationDelay"] = args ? args.deregistrationDelay : undefined;
            resourceInputs["healthCheck"] = args ? args.healthCheck : undefined;
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["lambdaMultiValueHeadersEnabled"] = args ? args.lambdaMultiValueHeadersEnabled : undefined;
            resourceInputs["loadBalancingAlgorithmType"] = args ? args.loadBalancingAlgorithmType : undefined;
            resourceInputs["loadBalancingAnomalyMitigation"] = args ? args.loadBalancingAnomalyMitigation : undefined;
            resourceInputs["loadBalancingCrossZoneEnabled"] = args ? args.loadBalancingCrossZoneEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["preserveClientIp"] = args ? args.preserveClientIp : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["protocolVersion"] = args ? args.protocolVersion : undefined;
            resourceInputs["proxyProtocolV2"] = args ? args.proxyProtocolV2 : undefined;
            resourceInputs["slowStart"] = args ? args.slowStart : undefined;
            resourceInputs["stickiness"] = args ? args.stickiness : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetFailovers"] = args ? args.targetFailovers : undefined;
            resourceInputs["targetHealthStates"] = args ? args.targetHealthStates : undefined;
            resourceInputs["targetType"] = args ? args.targetType : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["arnSuffix"] = undefined /*out*/;
            resourceInputs["loadBalancerArns"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "aws:elasticloadbalancingv2/targetGroup:TargetGroup" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(TargetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetGroup resources.
 */
export interface TargetGroupState {
    /**
     * ARN of the Target Group (matches `id`).
     */
    arn?: pulumi.Input<string>;
    /**
     * ARN suffix for use with CloudWatch Metrics.
     */
    arnSuffix?: pulumi.Input<string>;
    /**
     * Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
     */
    connectionTermination?: pulumi.Input<boolean>;
    /**
     * Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
     */
    deregistrationDelay?: pulumi.Input<number>;
    /**
     * Health Check configuration block. Detailed below.
     */
    healthCheck?: pulumi.Input<inputs.lb.TargetGroupHealthCheck>;
    /**
     * The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
     */
    lambdaMultiValueHeadersEnabled?: pulumi.Input<boolean>;
    /**
     * ARNs of the Load Balancers associated with the Target Group.
     */
    loadBalancerArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin`, `leastOutstandingRequests`, or `weightedRandom`. The default is `roundRobin`.
     */
    loadBalancingAlgorithmType?: pulumi.Input<string>;
    /**
     * Determines whether to enable target anomaly mitigation.  Target anomaly mitigation is only supported by the `weightedRandom` load balancing algorithm type.  See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#automatic-target-weights) for more information.  The value is `"on"` or `"off"`. The default is `"off"`.
     */
    loadBalancingAnomalyMitigation?: pulumi.Input<string>;
    /**
     * Indicates whether cross zone load balancing is enabled. The value is `"true"`, `"false"` or `"useLoadBalancerConfiguration"`. The default is `"useLoadBalancerConfiguration"`.
     */
    loadBalancingCrossZoneEnabled?: pulumi.Input<string>;
    /**
     * Name of the target group. If omitted, this provider will assign a random, unique name. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Port on which targets receive traffic, unless overridden when registering a specific target. Required when `targetType` is `instance`, `ip` or `alb`. Does not apply when `targetType` is `lambda`.
     */
    port?: pulumi.Input<number>;
    /**
     * Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
     */
    preserveClientIp?: pulumi.Input<string>;
    /**
     * Protocol to use for routing traffic to the targets.
     * Should be one of `GENEVE`, `HTTP`, `HTTPS`, `TCP`, `TCP_UDP`, `TLS`, or `UDP`.
     * Required when `targetType` is `instance`, `ip`, or `alb`.
     * Does not apply when `targetType` is `lambda`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
     */
    protocolVersion?: pulumi.Input<string>;
    /**
     * Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
     */
    proxyProtocolV2?: pulumi.Input<boolean>;
    /**
     * Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
     */
    slowStart?: pulumi.Input<number>;
    /**
     * Stickiness configuration block. Detailed below.
     */
    stickiness?: pulumi.Input<inputs.lb.TargetGroupStickiness>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Target failover block. Only applicable for Gateway Load Balancer target groups. See targetFailover for more information.
     */
    targetFailovers?: pulumi.Input<pulumi.Input<inputs.lb.TargetGroupTargetFailover>[]>;
    /**
     * Target health state block. Only applicable for Network Load Balancer target groups when `protocol` is `TCP` or `TLS`. See targetHealthState for more information.
     */
    targetHealthStates?: pulumi.Input<pulumi.Input<inputs.lb.TargetGroupTargetHealthState>[]>;
    /**
     * Type of target that you must specify when registering targets with this target group.
     * See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values.
     * The default is `instance`.
     *
     * Note that you can't specify targets for a target group using both instance IDs and IP addresses.
     *
     * If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
     *
     * Network Load Balancers do not support the `lambda` target type.
     *
     * Application Load Balancers do not support the `alb` target type.
     */
    targetType?: pulumi.Input<string>;
    /**
     * Identifier of the VPC in which to create the target group. Required when `targetType` is `instance`, `ip` or `alb`. Does not apply when `targetType` is `lambda`.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetGroup resource.
 */
export interface TargetGroupArgs {
    /**
     * Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
     */
    connectionTermination?: pulumi.Input<boolean>;
    /**
     * Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
     */
    deregistrationDelay?: pulumi.Input<number>;
    /**
     * Health Check configuration block. Detailed below.
     */
    healthCheck?: pulumi.Input<inputs.lb.TargetGroupHealthCheck>;
    /**
     * The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`. Default is `false`.
     */
    lambdaMultiValueHeadersEnabled?: pulumi.Input<boolean>;
    /**
     * Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `roundRobin`, `leastOutstandingRequests`, or `weightedRandom`. The default is `roundRobin`.
     */
    loadBalancingAlgorithmType?: pulumi.Input<string>;
    /**
     * Determines whether to enable target anomaly mitigation.  Target anomaly mitigation is only supported by the `weightedRandom` load balancing algorithm type.  See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#automatic-target-weights) for more information.  The value is `"on"` or `"off"`. The default is `"off"`.
     */
    loadBalancingAnomalyMitigation?: pulumi.Input<string>;
    /**
     * Indicates whether cross zone load balancing is enabled. The value is `"true"`, `"false"` or `"useLoadBalancerConfiguration"`. The default is `"useLoadBalancerConfiguration"`.
     */
    loadBalancingCrossZoneEnabled?: pulumi.Input<string>;
    /**
     * Name of the target group. If omitted, this provider will assign a random, unique name. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Port on which targets receive traffic, unless overridden when registering a specific target. Required when `targetType` is `instance`, `ip` or `alb`. Does not apply when `targetType` is `lambda`.
     */
    port?: pulumi.Input<number>;
    /**
     * Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
     */
    preserveClientIp?: pulumi.Input<string>;
    /**
     * Protocol to use for routing traffic to the targets.
     * Should be one of `GENEVE`, `HTTP`, `HTTPS`, `TCP`, `TCP_UDP`, `TLS`, or `UDP`.
     * Required when `targetType` is `instance`, `ip`, or `alb`.
     * Does not apply when `targetType` is `lambda`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
     */
    protocolVersion?: pulumi.Input<string>;
    /**
     * Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
     */
    proxyProtocolV2?: pulumi.Input<boolean>;
    /**
     * Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
     */
    slowStart?: pulumi.Input<number>;
    /**
     * Stickiness configuration block. Detailed below.
     */
    stickiness?: pulumi.Input<inputs.lb.TargetGroupStickiness>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Target failover block. Only applicable for Gateway Load Balancer target groups. See targetFailover for more information.
     */
    targetFailovers?: pulumi.Input<pulumi.Input<inputs.lb.TargetGroupTargetFailover>[]>;
    /**
     * Target health state block. Only applicable for Network Load Balancer target groups when `protocol` is `TCP` or `TLS`. See targetHealthState for more information.
     */
    targetHealthStates?: pulumi.Input<pulumi.Input<inputs.lb.TargetGroupTargetHealthState>[]>;
    /**
     * Type of target that you must specify when registering targets with this target group.
     * See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values.
     * The default is `instance`.
     *
     * Note that you can't specify targets for a target group using both instance IDs and IP addresses.
     *
     * If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
     *
     * Network Load Balancers do not support the `lambda` target type.
     *
     * Application Load Balancers do not support the `alb` target type.
     */
    targetType?: pulumi.Input<string>;
    /**
     * Identifier of the VPC in which to create the target group. Required when `targetType` is `instance`, `ip` or `alb`. Does not apply when `targetType` is `lambda`.
     */
    vpcId?: pulumi.Input<string>;
}
