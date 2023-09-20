// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an Amazon OpenSearch Domain.
 *
 * ## Elasticsearch vs. OpenSearch
 *
 * Amazon OpenSearch Service is the successor to Amazon Elasticsearch Service and supports OpenSearch and legacy Elasticsearch OSS (up to 7.10, the final open source version of the software).
 *
 * OpenSearch Domain configurations are similar in many ways to Elasticsearch Domain configurations. However, there are important differences including these:
 *
 * * OpenSearch has `engineVersion` while Elasticsearch has `elasticsearchVersion`
 * * Versions are specified differently - _e.g._, `Elasticsearch_7.10` with OpenSearch vs. `7.10` for Elasticsearch.
 * * `instanceType` argument values end in `search` for OpenSearch vs. `elasticsearch` for Elasticsearch (_e.g._, `t2.micro.search` vs. `t2.micro.elasticsearch`).
 * * The AWS-managed service-linked role for OpenSearch is called `AWSServiceRoleForAmazonOpenSearchService` instead of `AWSServiceRoleForAmazonElasticsearchService` for Elasticsearch.
 *
 * There are also some potentially unexpected similarities in configurations:
 *
 * * ARNs for both are prefaced with `arn:aws:es:`.
 * * Both OpenSearch and Elasticsearch use assume role policies that refer to the `Principal` `Service` as `es.amazonaws.com`.
 * * IAM policy actions, such as those you will find in `accessPolicies`, are prefaced with `es:` for both.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.Domain("example", {
 *     clusterConfig: {
 *         instanceType: "r4.large.search",
 *     },
 *     engineVersion: "Elasticsearch_7.10",
 *     tags: {
 *         Domain: "TestDomain",
 *     },
 * });
 * ```
 * ### Access Policy
 *
 * > See also: `aws.opensearch.DomainPolicy` resource
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const domain = config.get("domain") || "tf-test";
 * const currentRegion = aws.getRegion({});
 * const currentCallerIdentity = aws.getCallerIdentity({});
 * const examplePolicyDocument = Promise.all([currentRegion, currentCallerIdentity]).then(([currentRegion, currentCallerIdentity]) => aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "*",
 *             identifiers: ["*"],
 *         }],
 *         actions: ["es:*"],
 *         resources: [`arn:aws:es:${currentRegion.name}:${currentCallerIdentity.accountId}:domain/${domain}/*`],
 *         conditions: [{
 *             test: "IpAddress",
 *             variable: "aws:SourceIp",
 *             values: ["66.193.100.22/32"],
 *         }],
 *     }],
 * }));
 * const exampleDomain = new aws.opensearch.Domain("exampleDomain", {accessPolicies: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json)});
 * ```
 * ### Log publishing to CloudWatch Logs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
 * const examplePolicyDocument = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["es.amazonaws.com"],
 *         }],
 *         actions: [
 *             "logs:PutLogEvents",
 *             "logs:PutLogEventsBatch",
 *             "logs:CreateLogStream",
 *         ],
 *         resources: ["arn:aws:logs:*"],
 *     }],
 * });
 * const exampleLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("exampleLogResourcePolicy", {
 *     policyName: "example",
 *     policyDocument: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
 * });
 * // .. other configuration ...
 * const exampleDomain = new aws.opensearch.Domain("exampleDomain", {logPublishingOptions: [{
 *     cloudwatchLogGroupArn: exampleLogGroup.arn,
 *     logType: "INDEX_SLOW_LOGS",
 * }]});
 * ```
 * ### VPC based OpenSearch
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const vpc = config.requireObject("vpc");
 * const domain = config.get("domain") || "tf-test";
 * const exampleVpc = aws.ec2.getVpc({
 *     tags: {
 *         Name: vpc,
 *     },
 * });
 * const exampleSubnets = exampleVpc.then(exampleVpc => aws.ec2.getSubnets({
 *     filters: [{
 *         name: "vpc-id",
 *         values: [exampleVpc.id],
 *     }],
 *     tags: {
 *         Tier: "private",
 *     },
 * }));
 * const currentRegion = aws.getRegion({});
 * const currentCallerIdentity = aws.getCallerIdentity({});
 * const exampleSecurityGroup = new aws.ec2.SecurityGroup("exampleSecurityGroup", {
 *     description: "Managed by Pulumi",
 *     vpcId: exampleVpc.then(exampleVpc => exampleVpc.id),
 *     ingress: [{
 *         fromPort: 443,
 *         toPort: 443,
 *         protocol: "tcp",
 *         cidrBlocks: [exampleVpc.then(exampleVpc => exampleVpc.cidrBlock)],
 *     }],
 * });
 * const exampleServiceLinkedRole = new aws.iam.ServiceLinkedRole("exampleServiceLinkedRole", {awsServiceName: "opensearchservice.amazonaws.com"});
 * const examplePolicyDocument = Promise.all([currentRegion, currentCallerIdentity]).then(([currentRegion, currentCallerIdentity]) => aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "*",
 *             identifiers: ["*"],
 *         }],
 *         actions: ["es:*"],
 *         resources: [`arn:aws:es:${currentRegion.name}:${currentCallerIdentity.accountId}:domain/${domain}/*`],
 *     }],
 * }));
 * const exampleDomain = new aws.opensearch.Domain("exampleDomain", {
 *     engineVersion: "OpenSearch_1.0",
 *     clusterConfig: {
 *         instanceType: "m4.large.search",
 *         zoneAwarenessEnabled: true,
 *     },
 *     vpcOptions: {
 *         subnetIds: [
 *             exampleSubnets.then(exampleSubnets => exampleSubnets.ids?.[0]),
 *             exampleSubnets.then(exampleSubnets => exampleSubnets.ids?.[1]),
 *         ],
 *         securityGroupIds: [exampleSecurityGroup.id],
 *     },
 *     advancedOptions: {
 *         "rest.action.multi.allow_explicit_index": "true",
 *     },
 *     accessPolicies: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
 *     tags: {
 *         Domain: "TestDomain",
 *     },
 * }, {
 *     dependsOn: [exampleServiceLinkedRole],
 * });
 * ```
 * ### Enabling fine-grained access control on an existing domain
 *
 * This example shows two configurations: one to create a domain without fine-grained access control and the second to modify the domain to enable fine-grained access control. For more information, see [Enabling fine-grained access control](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html).
 * ### First apply
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.Domain("example", {
 *     advancedSecurityOptions: {
 *         anonymousAuthEnabled: true,
 *         enabled: false,
 *         internalUserDatabaseEnabled: true,
 *         masterUserOptions: {
 *             masterUserName: "example",
 *             masterUserPassword: "Barbarbarbar1!",
 *         },
 *     },
 *     clusterConfig: {
 *         instanceType: "r5.large.search",
 *     },
 *     domainEndpointOptions: {
 *         enforceHttps: true,
 *         tlsSecurityPolicy: "Policy-Min-TLS-1-2-2019-07",
 *     },
 *     ebsOptions: {
 *         ebsEnabled: true,
 *         volumeSize: 10,
 *     },
 *     encryptAtRest: {
 *         enabled: true,
 *     },
 *     engineVersion: "Elasticsearch_7.1",
 *     nodeToNodeEncryption: {
 *         enabled: true,
 *     },
 * });
 * ```
 * ### Second apply
 *
 * Notice that the only change is `advanced_security_options.0.enabled` is now set to `true`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.Domain("example", {
 *     advancedSecurityOptions: {
 *         anonymousAuthEnabled: true,
 *         enabled: true,
 *         internalUserDatabaseEnabled: true,
 *         masterUserOptions: {
 *             masterUserName: "example",
 *             masterUserPassword: "Barbarbarbar1!",
 *         },
 *     },
 *     clusterConfig: {
 *         instanceType: "r5.large.search",
 *     },
 *     domainEndpointOptions: {
 *         enforceHttps: true,
 *         tlsSecurityPolicy: "Policy-Min-TLS-1-2-2019-07",
 *     },
 *     ebsOptions: {
 *         ebsEnabled: true,
 *         volumeSize: 10,
 *     },
 *     encryptAtRest: {
 *         enabled: true,
 *     },
 *     engineVersion: "Elasticsearch_7.1",
 *     nodeToNodeEncryption: {
 *         enabled: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import OpenSearch domains using the `domain_name`. For example:
 *
 * ```sh
 *  $ pulumi import aws:opensearch/domain:Domain example domain_name
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opensearch/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * IAM policy document specifying the access policies for the domain.
     */
    public readonly accessPolicies!: pulumi.Output<string>;
    /**
     * Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your OpenSearch domain on every apply.
     */
    public readonly advancedOptions!: pulumi.Output<{[key: string]: string}>;
    /**
     * Configuration block for [fine-grained access control](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html). Detailed below.
     */
    public readonly advancedSecurityOptions!: pulumi.Output<outputs.opensearch.DomainAdvancedSecurityOptions>;
    /**
     * ARN of the domain.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block for the Auto-Tune options of the domain. Detailed below.
     */
    public readonly autoTuneOptions!: pulumi.Output<outputs.opensearch.DomainAutoTuneOptions>;
    /**
     * Configuration block for the cluster of the domain. Detailed below.
     */
    public readonly clusterConfig!: pulumi.Output<outputs.opensearch.DomainClusterConfig>;
    /**
     * Configuration block for authenticating dashboard with Cognito. Detailed below.
     */
    public readonly cognitoOptions!: pulumi.Output<outputs.opensearch.DomainCognitoOptions | undefined>;
    /**
     * Domain-specific endpoint for Dashboard without https scheme.
     */
    public /*out*/ readonly dashboardEndpoint!: pulumi.Output<string>;
    /**
     * Configuration block for domain endpoint HTTP(S) related options. Detailed below.
     */
    public readonly domainEndpointOptions!: pulumi.Output<outputs.opensearch.DomainDomainEndpointOptions>;
    /**
     * Unique identifier for the domain.
     */
    public /*out*/ readonly domainId!: pulumi.Output<string>;
    /**
     * Name of the domain.
     *
     * The following arguments are optional:
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/opensearch-service/pricing/). Detailed below.
     */
    public readonly ebsOptions!: pulumi.Output<outputs.opensearch.DomainEbsOptions>;
    /**
     * Configuration block for encrypt at rest options. Only available for [certain instance types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html). Detailed below.
     */
    public readonly encryptAtRest!: pulumi.Output<outputs.opensearch.DomainEncryptAtRest>;
    /**
     * Domain-specific endpoint used to submit index, search, and data upload requests.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Either `Elasticsearch_X.Y` or `OpenSearch_X.Y` to specify the engine version for the Amazon OpenSearch Service domain. For example, `OpenSearch_1.0` or `Elasticsearch_7.9`.
     * See [Creating and managing Amazon OpenSearch Service domains](http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).
     * Defaults to the lastest version of OpenSearch.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * (**Deprecated**) Domain-specific endpoint for kibana without https scheme. Use the `dashboardEndpoint` attribute instead.
     *
     * @deprecated use 'dashboard_endpoint' attribute instead
     */
    public /*out*/ readonly kibanaEndpoint!: pulumi.Output<string>;
    /**
     * Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
     */
    public readonly logPublishingOptions!: pulumi.Output<outputs.opensearch.DomainLogPublishingOption[] | undefined>;
    /**
     * Configuration block for node-to-node encryption options. Detailed below.
     */
    public readonly nodeToNodeEncryption!: pulumi.Output<outputs.opensearch.DomainNodeToNodeEncryption>;
    /**
     * Configuration to add Off Peak update options. ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html)). Detailed below.
     */
    public readonly offPeakWindowOptions!: pulumi.Output<outputs.opensearch.DomainOffPeakWindowOptions>;
    /**
     * Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running OpenSearch 5.3 and later, Amazon OpenSearch takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions, OpenSearch takes daily automated snapshots.
     */
    public readonly snapshotOptions!: pulumi.Output<outputs.opensearch.DomainSnapshotOptions | undefined>;
    /**
     * Software update options for the domain. Detailed below.
     */
    public readonly softwareUpdateOptions!: pulumi.Output<outputs.opensearch.DomainSoftwareUpdateOptions>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     * * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnetIds` were created inside.
     * * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html)). Detailed below.
     */
    public readonly vpcOptions!: pulumi.Output<outputs.opensearch.DomainVpcOptions | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            resourceInputs["accessPolicies"] = state ? state.accessPolicies : undefined;
            resourceInputs["advancedOptions"] = state ? state.advancedOptions : undefined;
            resourceInputs["advancedSecurityOptions"] = state ? state.advancedSecurityOptions : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["autoTuneOptions"] = state ? state.autoTuneOptions : undefined;
            resourceInputs["clusterConfig"] = state ? state.clusterConfig : undefined;
            resourceInputs["cognitoOptions"] = state ? state.cognitoOptions : undefined;
            resourceInputs["dashboardEndpoint"] = state ? state.dashboardEndpoint : undefined;
            resourceInputs["domainEndpointOptions"] = state ? state.domainEndpointOptions : undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["ebsOptions"] = state ? state.ebsOptions : undefined;
            resourceInputs["encryptAtRest"] = state ? state.encryptAtRest : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["kibanaEndpoint"] = state ? state.kibanaEndpoint : undefined;
            resourceInputs["logPublishingOptions"] = state ? state.logPublishingOptions : undefined;
            resourceInputs["nodeToNodeEncryption"] = state ? state.nodeToNodeEncryption : undefined;
            resourceInputs["offPeakWindowOptions"] = state ? state.offPeakWindowOptions : undefined;
            resourceInputs["snapshotOptions"] = state ? state.snapshotOptions : undefined;
            resourceInputs["softwareUpdateOptions"] = state ? state.softwareUpdateOptions : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcOptions"] = state ? state.vpcOptions : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            resourceInputs["accessPolicies"] = args ? args.accessPolicies : undefined;
            resourceInputs["advancedOptions"] = args ? args.advancedOptions : undefined;
            resourceInputs["advancedSecurityOptions"] = args ? args.advancedSecurityOptions : undefined;
            resourceInputs["autoTuneOptions"] = args ? args.autoTuneOptions : undefined;
            resourceInputs["clusterConfig"] = args ? args.clusterConfig : undefined;
            resourceInputs["cognitoOptions"] = args ? args.cognitoOptions : undefined;
            resourceInputs["domainEndpointOptions"] = args ? args.domainEndpointOptions : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["ebsOptions"] = args ? args.ebsOptions : undefined;
            resourceInputs["encryptAtRest"] = args ? args.encryptAtRest : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["logPublishingOptions"] = args ? args.logPublishingOptions : undefined;
            resourceInputs["nodeToNodeEncryption"] = args ? args.nodeToNodeEncryption : undefined;
            resourceInputs["offPeakWindowOptions"] = args ? args.offPeakWindowOptions : undefined;
            resourceInputs["snapshotOptions"] = args ? args.snapshotOptions : undefined;
            resourceInputs["softwareUpdateOptions"] = args ? args.softwareUpdateOptions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcOptions"] = args ? args.vpcOptions : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["dashboardEndpoint"] = undefined /*out*/;
            resourceInputs["domainId"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["kibanaEndpoint"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * IAM policy document specifying the access policies for the domain.
     */
    accessPolicies?: pulumi.Input<string>;
    /**
     * Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your OpenSearch domain on every apply.
     */
    advancedOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for [fine-grained access control](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html). Detailed below.
     */
    advancedSecurityOptions?: pulumi.Input<inputs.opensearch.DomainAdvancedSecurityOptions>;
    /**
     * ARN of the domain.
     */
    arn?: pulumi.Input<string>;
    /**
     * Configuration block for the Auto-Tune options of the domain. Detailed below.
     */
    autoTuneOptions?: pulumi.Input<inputs.opensearch.DomainAutoTuneOptions>;
    /**
     * Configuration block for the cluster of the domain. Detailed below.
     */
    clusterConfig?: pulumi.Input<inputs.opensearch.DomainClusterConfig>;
    /**
     * Configuration block for authenticating dashboard with Cognito. Detailed below.
     */
    cognitoOptions?: pulumi.Input<inputs.opensearch.DomainCognitoOptions>;
    /**
     * Domain-specific endpoint for Dashboard without https scheme.
     */
    dashboardEndpoint?: pulumi.Input<string>;
    /**
     * Configuration block for domain endpoint HTTP(S) related options. Detailed below.
     */
    domainEndpointOptions?: pulumi.Input<inputs.opensearch.DomainDomainEndpointOptions>;
    /**
     * Unique identifier for the domain.
     */
    domainId?: pulumi.Input<string>;
    /**
     * Name of the domain.
     *
     * The following arguments are optional:
     */
    domainName?: pulumi.Input<string>;
    /**
     * Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/opensearch-service/pricing/). Detailed below.
     */
    ebsOptions?: pulumi.Input<inputs.opensearch.DomainEbsOptions>;
    /**
     * Configuration block for encrypt at rest options. Only available for [certain instance types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html). Detailed below.
     */
    encryptAtRest?: pulumi.Input<inputs.opensearch.DomainEncryptAtRest>;
    /**
     * Domain-specific endpoint used to submit index, search, and data upload requests.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Either `Elasticsearch_X.Y` or `OpenSearch_X.Y` to specify the engine version for the Amazon OpenSearch Service domain. For example, `OpenSearch_1.0` or `Elasticsearch_7.9`.
     * See [Creating and managing Amazon OpenSearch Service domains](http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).
     * Defaults to the lastest version of OpenSearch.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * (**Deprecated**) Domain-specific endpoint for kibana without https scheme. Use the `dashboardEndpoint` attribute instead.
     *
     * @deprecated use 'dashboard_endpoint' attribute instead
     */
    kibanaEndpoint?: pulumi.Input<string>;
    /**
     * Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
     */
    logPublishingOptions?: pulumi.Input<pulumi.Input<inputs.opensearch.DomainLogPublishingOption>[]>;
    /**
     * Configuration block for node-to-node encryption options. Detailed below.
     */
    nodeToNodeEncryption?: pulumi.Input<inputs.opensearch.DomainNodeToNodeEncryption>;
    /**
     * Configuration to add Off Peak update options. ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html)). Detailed below.
     */
    offPeakWindowOptions?: pulumi.Input<inputs.opensearch.DomainOffPeakWindowOptions>;
    /**
     * Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running OpenSearch 5.3 and later, Amazon OpenSearch takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions, OpenSearch takes daily automated snapshots.
     */
    snapshotOptions?: pulumi.Input<inputs.opensearch.DomainSnapshotOptions>;
    /**
     * Software update options for the domain. Detailed below.
     */
    softwareUpdateOptions?: pulumi.Input<inputs.opensearch.DomainSoftwareUpdateOptions>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     * * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnetIds` were created inside.
     * * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html)). Detailed below.
     */
    vpcOptions?: pulumi.Input<inputs.opensearch.DomainVpcOptions>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * IAM policy document specifying the access policies for the domain.
     */
    accessPolicies?: pulumi.Input<string>;
    /**
     * Key-value string pairs to specify advanced configuration options. Note that the values for these configuration options must be strings (wrapped in quotes) or they may be wrong and cause a perpetual diff, causing the provider to want to recreate your OpenSearch domain on every apply.
     */
    advancedOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for [fine-grained access control](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html). Detailed below.
     */
    advancedSecurityOptions?: pulumi.Input<inputs.opensearch.DomainAdvancedSecurityOptions>;
    /**
     * Configuration block for the Auto-Tune options of the domain. Detailed below.
     */
    autoTuneOptions?: pulumi.Input<inputs.opensearch.DomainAutoTuneOptions>;
    /**
     * Configuration block for the cluster of the domain. Detailed below.
     */
    clusterConfig?: pulumi.Input<inputs.opensearch.DomainClusterConfig>;
    /**
     * Configuration block for authenticating dashboard with Cognito. Detailed below.
     */
    cognitoOptions?: pulumi.Input<inputs.opensearch.DomainCognitoOptions>;
    /**
     * Configuration block for domain endpoint HTTP(S) related options. Detailed below.
     */
    domainEndpointOptions?: pulumi.Input<inputs.opensearch.DomainDomainEndpointOptions>;
    /**
     * Name of the domain.
     *
     * The following arguments are optional:
     */
    domainName?: pulumi.Input<string>;
    /**
     * Configuration block for EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/opensearch-service/pricing/). Detailed below.
     */
    ebsOptions?: pulumi.Input<inputs.opensearch.DomainEbsOptions>;
    /**
     * Configuration block for encrypt at rest options. Only available for [certain instance types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html). Detailed below.
     */
    encryptAtRest?: pulumi.Input<inputs.opensearch.DomainEncryptAtRest>;
    /**
     * Either `Elasticsearch_X.Y` or `OpenSearch_X.Y` to specify the engine version for the Amazon OpenSearch Service domain. For example, `OpenSearch_1.0` or `Elasticsearch_7.9`.
     * See [Creating and managing Amazon OpenSearch Service domains](http://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).
     * Defaults to the lastest version of OpenSearch.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Configuration block for publishing slow and application logs to CloudWatch Logs. This block can be declared multiple times, for each log_type, within the same resource. Detailed below.
     */
    logPublishingOptions?: pulumi.Input<pulumi.Input<inputs.opensearch.DomainLogPublishingOption>[]>;
    /**
     * Configuration block for node-to-node encryption options. Detailed below.
     */
    nodeToNodeEncryption?: pulumi.Input<inputs.opensearch.DomainNodeToNodeEncryption>;
    /**
     * Configuration to add Off Peak update options. ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html)). Detailed below.
     */
    offPeakWindowOptions?: pulumi.Input<inputs.opensearch.DomainOffPeakWindowOptions>;
    /**
     * Configuration block for snapshot related options. Detailed below. DEPRECATED. For domains running OpenSearch 5.3 and later, Amazon OpenSearch takes hourly automated snapshots, making this setting irrelevant. For domains running earlier versions, OpenSearch takes daily automated snapshots.
     */
    snapshotOptions?: pulumi.Input<inputs.opensearch.DomainSnapshotOptions>;
    /**
     * Software update options for the domain. Detailed below.
     */
    softwareUpdateOptions?: pulumi.Input<inputs.opensearch.DomainSoftwareUpdateOptions>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for VPC related options. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html)). Detailed below.
     */
    vpcOptions?: pulumi.Input<inputs.opensearch.DomainVpcOptions>;
}
