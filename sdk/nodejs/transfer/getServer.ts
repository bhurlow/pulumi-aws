// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ARN of an AWS Transfer Server for use in other
 * resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.transfer.getServer({
 *     serverId: "s-1234567",
 * }));
 * ```
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:transfer/getServer:getServer", {
        "serverId": args.serverId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServer.
 */
export interface GetServerArgs {
    /**
     * ID for an SFTP server.
     */
    serverId: string;
}

/**
 * A collection of values returned by getServer.
 */
export interface GetServerResult {
    /**
     * Amazon Resource Name (ARN) of Transfer Server.
     */
    readonly arn: string;
    /**
     * The ARN of any certificate.
     */
    readonly certificate: string;
    /**
     * The domain of the storage system that is used for file transfers.
     */
    readonly domain: string;
    /**
     * The endpoint of the Transfer Server (e.g. `s-12345678.server.transfer.REGION.amazonaws.com`).
     */
    readonly endpoint: string;
    /**
     * The type of endpoint that the server is connected to.
     */
    readonly endpointType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
     */
    readonly identityProviderType: string;
    /**
     * Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identityProviderType` of `API_GATEWAY`.
     */
    readonly invocationRole: string;
    /**
     * Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
     */
    readonly loggingRole: string;
    /**
     * The file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint.
     */
    readonly protocols: string[];
    /**
     * The name of the security policy that is attached to the server.
     */
    readonly securityPolicyName: string;
    readonly serverId: string;
    /**
     * URL of the service endpoint used to authenticate users with an `identityProviderType` of `API_GATEWAY`.
     */
    readonly url: string;
}
