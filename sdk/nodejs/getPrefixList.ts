// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * `aws.ec2.getPrefixList` provides details about a specific AWS prefix list (PL)
 * in the current region.
 *
 * This can be used both to validate a prefix list given in a variable
 * and to obtain the CIDR blocks (IP address ranges) for the associated
 * AWS service. The latter may be useful e.g., for adding network ACL
 * rules.
 *
 * The aws.ec2.ManagedPrefixList data source is normally more appropriate to use given it can return customer-managed prefix list info, as well as additional attributes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const privateS3VpcEndpoint = new aws.ec2.VpcEndpoint("privateS3VpcEndpoint", {
 *     vpcId: aws_vpc.foo.id,
 *     serviceName: "com.amazonaws.us-west-2.s3",
 * });
 * const privateS3PrefixList = aws.ec2.getPrefixListOutput({
 *     prefixListId: privateS3VpcEndpoint.prefixListId,
 * });
 * const bar = new aws.ec2.NetworkAcl("bar", {vpcId: aws_vpc.foo.id});
 * const privateS3NetworkAclRule = new aws.ec2.NetworkAclRule("privateS3NetworkAclRule", {
 *     networkAclId: bar.id,
 *     ruleNumber: 200,
 *     egress: false,
 *     protocol: "tcp",
 *     ruleAction: "allow",
 *     cidrBlock: privateS3PrefixList.apply(privateS3PrefixList => privateS3PrefixList.cidrBlocks?[0]),
 *     fromPort: 443,
 *     toPort: 443,
 * });
 * ```
 * ### Filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.ec2.getPrefixList({
 *     filters: [{
 *         name: "prefix-list-id",
 *         values: ["pl-68a54001"],
 *     }],
 * }));
 * ```
 */
/** @deprecated aws.getPrefixList has been deprecated in favor of aws.ec2.getPrefixList */
export function getPrefixList(args?: GetPrefixListArgs, opts?: pulumi.InvokeOptions): Promise<GetPrefixListResult> {
    pulumi.log.warn("getPrefixList is deprecated: aws.getPrefixList has been deprecated in favor of aws.ec2.getPrefixList")
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws:index/getPrefixList:getPrefixList", {
        "filters": args.filters,
        "name": args.name,
        "prefixListId": args.prefixListId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrefixList.
 */
export interface GetPrefixListArgs {
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    filters?: inputs.GetPrefixListFilter[];
    /**
     * Name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
     */
    name?: string;
    /**
     * ID of the prefix list to select.
     */
    prefixListId?: string;
}

/**
 * A collection of values returned by getPrefixList.
 */
export interface GetPrefixListResult {
    /**
     * List of CIDR blocks for the AWS service associated with the prefix list.
     */
    readonly cidrBlocks: string[];
    readonly filters?: outputs.GetPrefixListFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the selected prefix list.
     */
    readonly name: string;
    readonly prefixListId?: string;
}

export function getPrefixListOutput(args?: GetPrefixListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrefixListResult> {
    return pulumi.output(args).apply(a => getPrefixList(a, opts))
}

/**
 * A collection of arguments for invoking getPrefixList.
 */
export interface GetPrefixListOutputArgs {
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetPrefixListFilterArgs>[]>;
    /**
     * Name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the prefix list to select.
     */
    prefixListId?: pulumi.Input<string>;
}
