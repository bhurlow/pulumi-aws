// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about multiple EC2 Local Gateway Virtual Interface Groups, such as identifiers. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const all = pulumi.output(aws.ec2.getLocalGatewayVirtualInterfaceGroups());
 * ```
 */
export function getLocalGatewayVirtualInterfaceGroups(args?: GetLocalGatewayVirtualInterfaceGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayVirtualInterfaceGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGatewayVirtualInterfaceGroups:getLocalGatewayVirtualInterfaceGroups", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGatewayVirtualInterfaceGroups.
 */
export interface GetLocalGatewayVirtualInterfaceGroupsArgs {
    /**
     * One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGatewayVirtualInterfaceGroups.html) for supported filters. Detailed below.
     */
    filters?: inputs.ec2.GetLocalGatewayVirtualInterfaceGroupsFilter[];
    /**
     * Key-value map of resource tags, each pair of which must exactly match a pair on the desired local gateway route table.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGatewayVirtualInterfaceGroups.
 */
export interface GetLocalGatewayVirtualInterfaceGroupsResult {
    readonly filters?: outputs.ec2.GetLocalGatewayVirtualInterfaceGroupsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of EC2 Local Gateway Virtual Interface Group identifiers.
     */
    readonly ids: string[];
    /**
     * Set of EC2 Local Gateway Virtual Interface identifiers.
     */
    readonly localGatewayVirtualInterfaceIds: string[];
    readonly tags: {[key: string]: string};
}
