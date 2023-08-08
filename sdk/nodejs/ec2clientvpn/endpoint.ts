// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
 * [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2clientvpn.Endpoint("example", {
 *     description: "clientvpn-example",
 *     serverCertificateArn: aws_acm_certificate.cert.arn,
 *     clientCidrBlock: "10.0.0.0/16",
 *     authenticationOptions: [{
 *         type: "certificate-authentication",
 *         rootCertificateChainArn: aws_acm_certificate.root_cert.arn,
 *     }],
 *     connectionLogOptions: {
 *         enabled: true,
 *         cloudwatchLogGroup: aws_cloudwatch_log_group.lg.name,
 *         cloudwatchLogStream: aws_cloudwatch_log_stream.ls.name,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ec2_client_vpn_endpoint.example
 *
 *  id = "cvpn-endpoint-0ac3a1abbccddd666" } Using `pulumi import`, import AWS Client VPN endpoints using the `id` value found via `aws ec2 describe-client-vpn-endpoints`. For exampleconsole % pulumi import aws_ec2_client_vpn_endpoint.example cvpn-endpoint-0ac3a1abbccddd666
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2clientvpn/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * The ARN of the Client VPN endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Information about the authentication method to be used to authenticate clients.
     */
    public readonly authenticationOptions!: pulumi.Output<outputs.ec2clientvpn.EndpointAuthenticationOption[]>;
    /**
     * The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
     */
    public readonly clientCidrBlock!: pulumi.Output<string>;
    /**
     * The options for managing connection authorization for new client connections.
     */
    public readonly clientConnectOptions!: pulumi.Output<outputs.ec2clientvpn.EndpointClientConnectOptions>;
    /**
     * Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
     */
    public readonly clientLoginBannerOptions!: pulumi.Output<outputs.ec2clientvpn.EndpointClientLoginBannerOptions>;
    /**
     * Information about the client connection logging options.
     */
    public readonly connectionLogOptions!: pulumi.Output<outputs.ec2clientvpn.EndpointConnectionLogOptions>;
    /**
     * A brief description of the Client VPN endpoint.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The DNS name to be used by clients when establishing their VPN session.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the connecting device is used.
     */
    public readonly dnsServers!: pulumi.Output<string[] | undefined>;
    /**
     * The IDs of one or more security groups to apply to the target network. You must also specify the ID of the VPC that contains the security groups.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * Specify whether to enable the self-service portal for the Client VPN endpoint. Values can be `enabled` or `disabled`. Default value is `disabled`.
     */
    public readonly selfServicePortal!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the ACM server certificate.
     */
    public readonly serverCertificateArn!: pulumi.Output<string>;
    /**
     * The maximum session duration is a trigger by which end-users are required to re-authenticate prior to establishing a VPN session. Default value is `24` - Valid values: `8 | 10 | 12 | 24`
     */
    public readonly sessionTimeoutHours!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
     */
    public readonly splitTunnel!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The transport protocol to be used by the VPN session. Default value is `udp`.
     */
    public readonly transportProtocol!: pulumi.Output<string | undefined>;
    /**
     * The ID of the VPC to associate with the Client VPN endpoint. If no security group IDs are specified in the request, the default security group for the VPC is applied.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The port number for the Client VPN endpoint. Valid values are `443` and `1194`. Default value is `443`.
     */
    public readonly vpnPort!: pulumi.Output<number | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["authenticationOptions"] = state ? state.authenticationOptions : undefined;
            resourceInputs["clientCidrBlock"] = state ? state.clientCidrBlock : undefined;
            resourceInputs["clientConnectOptions"] = state ? state.clientConnectOptions : undefined;
            resourceInputs["clientLoginBannerOptions"] = state ? state.clientLoginBannerOptions : undefined;
            resourceInputs["connectionLogOptions"] = state ? state.connectionLogOptions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["dnsServers"] = state ? state.dnsServers : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["selfServicePortal"] = state ? state.selfServicePortal : undefined;
            resourceInputs["serverCertificateArn"] = state ? state.serverCertificateArn : undefined;
            resourceInputs["sessionTimeoutHours"] = state ? state.sessionTimeoutHours : undefined;
            resourceInputs["splitTunnel"] = state ? state.splitTunnel : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["transportProtocol"] = state ? state.transportProtocol : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpnPort"] = state ? state.vpnPort : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if ((!args || args.authenticationOptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authenticationOptions'");
            }
            if ((!args || args.clientCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientCidrBlock'");
            }
            if ((!args || args.connectionLogOptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionLogOptions'");
            }
            if ((!args || args.serverCertificateArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverCertificateArn'");
            }
            resourceInputs["authenticationOptions"] = args ? args.authenticationOptions : undefined;
            resourceInputs["clientCidrBlock"] = args ? args.clientCidrBlock : undefined;
            resourceInputs["clientConnectOptions"] = args ? args.clientConnectOptions : undefined;
            resourceInputs["clientLoginBannerOptions"] = args ? args.clientLoginBannerOptions : undefined;
            resourceInputs["connectionLogOptions"] = args ? args.connectionLogOptions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsServers"] = args ? args.dnsServers : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["selfServicePortal"] = args ? args.selfServicePortal : undefined;
            resourceInputs["serverCertificateArn"] = args ? args.serverCertificateArn : undefined;
            resourceInputs["sessionTimeoutHours"] = args ? args.sessionTimeoutHours : undefined;
            resourceInputs["splitTunnel"] = args ? args.splitTunnel : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transportProtocol"] = args ? args.transportProtocol : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vpnPort"] = args ? args.vpnPort : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Endpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    /**
     * The ARN of the Client VPN endpoint.
     */
    arn?: pulumi.Input<string>;
    /**
     * Information about the authentication method to be used to authenticate clients.
     */
    authenticationOptions?: pulumi.Input<pulumi.Input<inputs.ec2clientvpn.EndpointAuthenticationOption>[]>;
    /**
     * The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
     */
    clientCidrBlock?: pulumi.Input<string>;
    /**
     * The options for managing connection authorization for new client connections.
     */
    clientConnectOptions?: pulumi.Input<inputs.ec2clientvpn.EndpointClientConnectOptions>;
    /**
     * Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
     */
    clientLoginBannerOptions?: pulumi.Input<inputs.ec2clientvpn.EndpointClientLoginBannerOptions>;
    /**
     * Information about the client connection logging options.
     */
    connectionLogOptions?: pulumi.Input<inputs.ec2clientvpn.EndpointConnectionLogOptions>;
    /**
     * A brief description of the Client VPN endpoint.
     */
    description?: pulumi.Input<string>;
    /**
     * The DNS name to be used by clients when establishing their VPN session.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the connecting device is used.
     */
    dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of one or more security groups to apply to the target network. You must also specify the ID of the VPC that contains the security groups.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify whether to enable the self-service portal for the Client VPN endpoint. Values can be `enabled` or `disabled`. Default value is `disabled`.
     */
    selfServicePortal?: pulumi.Input<string>;
    /**
     * The ARN of the ACM server certificate.
     */
    serverCertificateArn?: pulumi.Input<string>;
    /**
     * The maximum session duration is a trigger by which end-users are required to re-authenticate prior to establishing a VPN session. Default value is `24` - Valid values: `8 | 10 | 12 | 24`
     */
    sessionTimeoutHours?: pulumi.Input<number>;
    /**
     * Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
     */
    splitTunnel?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The transport protocol to be used by the VPN session. Default value is `udp`.
     */
    transportProtocol?: pulumi.Input<string>;
    /**
     * The ID of the VPC to associate with the Client VPN endpoint. If no security group IDs are specified in the request, the default security group for the VPC is applied.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The port number for the Client VPN endpoint. Valid values are `443` and `1194`. Default value is `443`.
     */
    vpnPort?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * Information about the authentication method to be used to authenticate clients.
     */
    authenticationOptions: pulumi.Input<pulumi.Input<inputs.ec2clientvpn.EndpointAuthenticationOption>[]>;
    /**
     * The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
     */
    clientCidrBlock: pulumi.Input<string>;
    /**
     * The options for managing connection authorization for new client connections.
     */
    clientConnectOptions?: pulumi.Input<inputs.ec2clientvpn.EndpointClientConnectOptions>;
    /**
     * Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
     */
    clientLoginBannerOptions?: pulumi.Input<inputs.ec2clientvpn.EndpointClientLoginBannerOptions>;
    /**
     * Information about the client connection logging options.
     */
    connectionLogOptions: pulumi.Input<inputs.ec2clientvpn.EndpointConnectionLogOptions>;
    /**
     * A brief description of the Client VPN endpoint.
     */
    description?: pulumi.Input<string>;
    /**
     * Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the connecting device is used.
     */
    dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of one or more security groups to apply to the target network. You must also specify the ID of the VPC that contains the security groups.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify whether to enable the self-service portal for the Client VPN endpoint. Values can be `enabled` or `disabled`. Default value is `disabled`.
     */
    selfServicePortal?: pulumi.Input<string>;
    /**
     * The ARN of the ACM server certificate.
     */
    serverCertificateArn: pulumi.Input<string>;
    /**
     * The maximum session duration is a trigger by which end-users are required to re-authenticate prior to establishing a VPN session. Default value is `24` - Valid values: `8 | 10 | 12 | 24`
     */
    sessionTimeoutHours?: pulumi.Input<number>;
    /**
     * Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
     */
    splitTunnel?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The transport protocol to be used by the VPN session. Default value is `udp`.
     */
    transportProtocol?: pulumi.Input<string>;
    /**
     * The ID of the VPC to associate with the Client VPN endpoint. If no security group IDs are specified in the request, the default security group for the VPC is applied.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The port number for the Client VPN endpoint. Valid values are `443` and `1194`. Default value is `443`.
     */
    vpnPort?: pulumi.Input<number>;
}
