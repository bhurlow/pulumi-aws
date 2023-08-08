// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an association with WAF Regional Web ACL.
 *
 * > **Note:** An Application Load Balancer can only be associated with one WAF Regional WebACL.
 *
 * ## Example Usage
 * ### Application Load Balancer Association
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ipset = new aws.wafregional.IpSet("ipset", {ipSetDescriptors: [{
 *     type: "IPV4",
 *     value: "192.0.7.0/24",
 * }]});
 * const fooRule = new aws.wafregional.Rule("fooRule", {
 *     metricName: "tfWAFRule",
 *     predicates: [{
 *         dataId: ipset.id,
 *         negated: false,
 *         type: "IPMatch",
 *     }],
 * });
 * const fooWebAcl = new aws.wafregional.WebAcl("fooWebAcl", {
 *     metricName: "foo",
 *     defaultAction: {
 *         type: "ALLOW",
 *     },
 *     rules: [{
 *         action: {
 *             type: "BLOCK",
 *         },
 *         priority: 1,
 *         ruleId: fooRule.id,
 *     }],
 * });
 * const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.1.0.0/16"});
 * const available = aws.getAvailabilityZones({});
 * const fooSubnet = new aws.ec2.Subnet("fooSubnet", {
 *     vpcId: fooVpc.id,
 *     cidrBlock: "10.1.1.0/24",
 *     availabilityZone: available.then(available => available.names?.[0]),
 * });
 * const bar = new aws.ec2.Subnet("bar", {
 *     vpcId: fooVpc.id,
 *     cidrBlock: "10.1.2.0/24",
 *     availabilityZone: available.then(available => available.names?.[1]),
 * });
 * const fooLoadBalancer = new aws.alb.LoadBalancer("fooLoadBalancer", {
 *     internal: true,
 *     subnets: [
 *         fooSubnet.id,
 *         bar.id,
 *     ],
 * });
 * const fooWebAclAssociation = new aws.wafregional.WebAclAssociation("fooWebAclAssociation", {
 *     resourceArn: fooLoadBalancer.arn,
 *     webAclId: fooWebAcl.id,
 * });
 * ```
 * ### API Gateway Association
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as crypto from "crypto";
 *
 * const ipset = new aws.wafregional.IpSet("ipset", {ipSetDescriptors: [{
 *     type: "IPV4",
 *     value: "192.0.7.0/24",
 * }]});
 * const fooRule = new aws.wafregional.Rule("fooRule", {
 *     metricName: "tfWAFRule",
 *     predicates: [{
 *         dataId: ipset.id,
 *         negated: false,
 *         type: "IPMatch",
 *     }],
 * });
 * const fooWebAcl = new aws.wafregional.WebAcl("fooWebAcl", {
 *     metricName: "foo",
 *     defaultAction: {
 *         type: "ALLOW",
 *     },
 *     rules: [{
 *         action: {
 *             type: "BLOCK",
 *         },
 *         priority: 1,
 *         ruleId: fooRule.id,
 *     }],
 * });
 * const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {body: JSON.stringify({
 *     openapi: "3.0.1",
 *     info: {
 *         title: "example",
 *         version: "1.0",
 *     },
 *     paths: {
 *         "/path1": {
 *             get: {
 *                 "x-amazon-apigateway-integration": {
 *                     httpMethod: "GET",
 *                     payloadFormatVersion: "1.0",
 *                     type: "HTTP_PROXY",
 *                     uri: "https://ip-ranges.amazonaws.com/ip-ranges.json",
 *                 },
 *             },
 *         },
 *     },
 * })});
 * const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {
 *     restApi: exampleRestApi.id,
 *     triggers: {
 *         redeployment: exampleRestApi.body.apply(body => JSON.stringify(body)).apply(toJSON => crypto.createHash('sha1').update(toJSON).digest('hex')),
 *     },
 * });
 * const exampleStage = new aws.apigateway.Stage("exampleStage", {
 *     deployment: exampleDeployment.id,
 *     restApi: exampleRestApi.id,
 *     stageName: "example",
 * });
 * const association = new aws.wafregional.WebAclAssociation("association", {
 *     resourceArn: exampleStage.arn,
 *     webAclId: fooWebAcl.id,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_wafregional_web_acl_association.foo
 *
 *  id = "web_acl_id:resource_arn" } Using `pulumi import`, import WAF Regional Web ACL Association using their `web_acl_id:resource_arn`. For exampleconsole % pulumi import aws_wafregional_web_acl_association.foo web_acl_id:resource_arn
 */
export class WebAclAssociation extends pulumi.CustomResource {
    /**
     * Get an existing WebAclAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAclAssociationState, opts?: pulumi.CustomResourceOptions): WebAclAssociation {
        return new WebAclAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafregional/webAclAssociation:WebAclAssociation';

    /**
     * Returns true if the given object is an instance of WebAclAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAclAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAclAssociation.__pulumiType;
    }

    /**
     * ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
     */
    public readonly resourceArn!: pulumi.Output<string>;
    /**
     * The ID of the WAF Regional WebACL to create an association.
     */
    public readonly webAclId!: pulumi.Output<string>;

    /**
     * Create a WebAclAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAclAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAclAssociationArgs | WebAclAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebAclAssociationState | undefined;
            resourceInputs["resourceArn"] = state ? state.resourceArn : undefined;
            resourceInputs["webAclId"] = state ? state.webAclId : undefined;
        } else {
            const args = argsOrState as WebAclAssociationArgs | undefined;
            if ((!args || args.resourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceArn'");
            }
            if ((!args || args.webAclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webAclId'");
            }
            resourceInputs["resourceArn"] = args ? args.resourceArn : undefined;
            resourceInputs["webAclId"] = args ? args.webAclId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebAclAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebAclAssociation resources.
 */
export interface WebAclAssociationState {
    /**
     * ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
     */
    resourceArn?: pulumi.Input<string>;
    /**
     * The ID of the WAF Regional WebACL to create an association.
     */
    webAclId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebAclAssociation resource.
 */
export interface WebAclAssociationArgs {
    /**
     * ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
     */
    resourceArn: pulumi.Input<string>;
    /**
     * The ID of the WAF Regional WebACL to create an association.
     */
    webAclId: pulumi.Input<string>;
}
