// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * `aws.vpc.getSecurityGroupRule` provides details about a specific security group rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.vpc.getSecurityGroupRule({
 *     securityGroupRuleId: _var.security_group_rule_id,
 * });
 * ```
 */
export function getSecurityGroupRule(args?: GetSecurityGroupRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupRuleResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:vpc/getSecurityGroupRule:getSecurityGroupRule", {
        "filters": args.filters,
        "securityGroupRuleId": args.securityGroupRuleId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroupRule.
 */
export interface GetSecurityGroupRuleArgs {
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    filters?: inputs.vpc.GetSecurityGroupRuleFilter[];
    /**
     * ID of the security group rule to select.
     */
    securityGroupRuleId?: string;
}

/**
 * A collection of values returned by getSecurityGroupRule.
 */
export interface GetSecurityGroupRuleResult {
    /**
     * The Amazon Resource Name (ARN) of the security group rule.
     */
    readonly arn: string;
    /**
     * The destination IPv4 CIDR range.
     */
    readonly cidrIpv4: string;
    /**
     * The destination IPv6 CIDR range.
     */
    readonly cidrIpv6: string;
    /**
     * The security group rule description.
     */
    readonly description: string;
    readonly filters?: outputs.vpc.GetSecurityGroupRuleFilter[];
    /**
     * The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
     */
    readonly fromPort: number;
    readonly id: string;
    /**
     * The IP protocol name or number. Use `-1` to specify all protocols.
     */
    readonly ipProtocol: string;
    /**
     * Indicates whether the security group rule is an outbound rule.
     */
    readonly isEgress: boolean;
    /**
     * The ID of the destination prefix list.
     */
    readonly prefixListId: string;
    /**
     * The destination security group that is referenced in the rule.
     */
    readonly referencedSecurityGroupId: string;
    /**
     * The ID of the security group.
     */
    readonly securityGroupId: string;
    readonly securityGroupRuleId: string;
    /**
     * A map of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * (Optional) The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
     */
    readonly toPort: number;
}
/**
 * `aws.vpc.getSecurityGroupRule` provides details about a specific security group rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.vpc.getSecurityGroupRule({
 *     securityGroupRuleId: _var.security_group_rule_id,
 * });
 * ```
 */
export function getSecurityGroupRuleOutput(args?: GetSecurityGroupRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupRuleResult> {
    return pulumi.output(args).apply((a: any) => getSecurityGroupRule(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityGroupRule.
 */
export interface GetSecurityGroupRuleOutputArgs {
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.vpc.GetSecurityGroupRuleFilterArgs>[]>;
    /**
     * ID of the security group rule to select.
     */
    securityGroupRuleId?: pulumi.Input<string>;
}
