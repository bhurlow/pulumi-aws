// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Amplify Branch resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.amplify.App("example", {});
 * const master = new aws.amplify.Branch("master", {
 *     appId: example.id,
 *     branchName: "master",
 *     framework: "React",
 *     stage: "PRODUCTION",
 *     environmentVariables: {
 *         REACT_APP_API_SERVER: "https://api.example.com",
 *     },
 * });
 * ```
 * ### Basic Authentication
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.amplify.App("example", {});
 * const master = new aws.amplify.Branch("master", {
 *     appId: example.id,
 *     branchName: "master",
 *     basicAuthConfig: [{
 *         enableBasicAuth: true,
 *         username: "username",
 *         password: "password",
 *     }],
 * });
 * ```
 * ### Notifications
 *
 * Amplify Console uses CloudWatch Events and SNS for email notifications.  To implement the same functionality, you need to set `enableNotification` in a `aws.amplify.Branch` resource, as well as creating a CloudWatch Events Rule, a SNS topic, and SNS subscriptions.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.amplify.App("example", {});
 * const master = new aws.amplify.Branch("master", {
 *     appId: example.id,
 *     branchName: "master",
 *     enableNotification: true,
 * });
 * // CloudWatch Events Rule for Amplify notifications
 * const amplifyAppMasterEventRule = new aws.cloudwatch.EventRule("amplifyAppMasterEventRule", {
 *     description: pulumi.interpolate`AWS Amplify build notifications for :  App: ${aws_amplify_app.app.id} Branch: ${master.branchName}`,
 *     eventPattern: pulumi.all([example.id, master.branchName]).apply(([id, branchName]) => JSON.stringify({
 *         detail: {
 *             appId: [id],
 *             branchName: [branchName],
 *             jobStatus: [
 *                 "SUCCEED",
 *                 "FAILED",
 *                 "STARTED",
 *             ],
 *         },
 *         "detail-type": ["Amplify Deployment Status Change"],
 *         source: ["aws.amplify"],
 *     })),
 * });
 * const amplifyAppMasterTopic = new aws.sns.Topic("amplifyAppMasterTopic", {});
 * const amplifyAppMasterEventTarget = new aws.cloudwatch.EventTarget("amplifyAppMasterEventTarget", {
 *     rule: amplifyAppMasterEventRule.name,
 *     arn: amplifyAppMasterTopic.arn,
 *     inputTransformer: {
 *         inputPaths: {
 *             jobId: `$.detail.jobId`,
 *             appId: `$.detail.appId`,
 *             region: `$.region`,
 *             branch: `$.detail.branchName`,
 *             status: `$.detail.jobStatus`,
 *         },
 *         inputTemplate: "\"Build notification from the AWS Amplify Console for app: https://<branch>.<appId>.amplifyapp.com/. Your build status is <status>. Go to https://console.aws.amazon.com/amplify/home?region=<region>#<appId>/<branch>/<jobId> to view details on your build. \"",
 *     },
 * });
 * // SNS Topic for Amplify notifications
 * const amplifyAppMasterPolicyDocument = pulumi.all([master.arn, amplifyAppMasterTopic.arn]).apply(([masterArn, amplifyAppMasterTopicArn]) => aws.iam.getPolicyDocument({
 *     statements: [{
 *         sid: `Allow_Publish_Events ${masterArn}`,
 *         effect: "Allow",
 *         actions: ["SNS:Publish"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["events.amazonaws.com"],
 *         }],
 *         resources: [amplifyAppMasterTopicArn],
 *     }],
 * }));
 * const amplifyAppMasterTopicPolicy = new aws.sns.TopicPolicy("amplifyAppMasterTopicPolicy", {
 *     arn: amplifyAppMasterTopic.arn,
 *     policy: amplifyAppMasterPolicyDocument.apply(amplifyAppMasterPolicyDocument => amplifyAppMasterPolicyDocument.json),
 * });
 * ```
 *
 * ## Import
 *
 * Amplify branch can be imported using `app_id` and `branch_name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:amplify/branch:Branch master d2ypk4k47z8u6/master
 * ```
 */
export class Branch extends pulumi.CustomResource {
    /**
     * Get an existing Branch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchState, opts?: pulumi.CustomResourceOptions): Branch {
        return new Branch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:amplify/branch:Branch';

    /**
     * Returns true if the given object is an instance of Branch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Branch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Branch.__pulumiType;
    }

    /**
     * The unique ID for an Amplify app.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) for the branch.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of custom resources that are linked to this branch.
     */
    public /*out*/ readonly associatedResources!: pulumi.Output<string[]>;
    /**
     * The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
     */
    public readonly backendEnvironmentArn!: pulumi.Output<string | undefined>;
    /**
     * The basic authorization credentials for the branch.
     */
    public readonly basicAuthCredentials!: pulumi.Output<string | undefined>;
    /**
     * The name for the branch.
     */
    public readonly branchName!: pulumi.Output<string>;
    /**
     * The custom domains for the branch.
     */
    public /*out*/ readonly customDomains!: pulumi.Output<string[]>;
    /**
     * The description for the branch.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination branch if the branch is a pull request branch.
     */
    public /*out*/ readonly destinationBranch!: pulumi.Output<string>;
    /**
     * The display name for a branch. This is used as the default domain prefix.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Enables auto building for the branch.
     */
    public readonly enableAutoBuild!: pulumi.Output<boolean | undefined>;
    /**
     * Enables basic authorization for the branch.
     */
    public readonly enableBasicAuth!: pulumi.Output<boolean | undefined>;
    /**
     * Enables notifications for the branch.
     */
    public readonly enableNotification!: pulumi.Output<boolean | undefined>;
    /**
     * Enables performance mode for the branch.
     */
    public readonly enablePerformanceMode!: pulumi.Output<boolean | undefined>;
    /**
     * Enables pull request previews for this branch.
     */
    public readonly enablePullRequestPreview!: pulumi.Output<boolean | undefined>;
    /**
     * The environment variables for the branch.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The framework for the branch.
     */
    public readonly framework!: pulumi.Output<string | undefined>;
    /**
     * The Amplify environment name for the pull request.
     */
    public readonly pullRequestEnvironmentName!: pulumi.Output<string | undefined>;
    /**
     * The source branch if the branch is a pull request branch.
     */
    public /*out*/ readonly sourceBranch!: pulumi.Output<string>;
    /**
     * Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
     */
    public readonly stage!: pulumi.Output<string | undefined>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The content Time To Live (TTL) for the website in seconds.
     */
    public readonly ttl!: pulumi.Output<string | undefined>;

    /**
     * Create a Branch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchArgs | BranchState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BranchState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["associatedResources"] = state ? state.associatedResources : undefined;
            inputs["backendEnvironmentArn"] = state ? state.backendEnvironmentArn : undefined;
            inputs["basicAuthCredentials"] = state ? state.basicAuthCredentials : undefined;
            inputs["branchName"] = state ? state.branchName : undefined;
            inputs["customDomains"] = state ? state.customDomains : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destinationBranch"] = state ? state.destinationBranch : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enableAutoBuild"] = state ? state.enableAutoBuild : undefined;
            inputs["enableBasicAuth"] = state ? state.enableBasicAuth : undefined;
            inputs["enableNotification"] = state ? state.enableNotification : undefined;
            inputs["enablePerformanceMode"] = state ? state.enablePerformanceMode : undefined;
            inputs["enablePullRequestPreview"] = state ? state.enablePullRequestPreview : undefined;
            inputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            inputs["framework"] = state ? state.framework : undefined;
            inputs["pullRequestEnvironmentName"] = state ? state.pullRequestEnvironmentName : undefined;
            inputs["sourceBranch"] = state ? state.sourceBranch : undefined;
            inputs["stage"] = state ? state.stage : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as BranchArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.branchName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'branchName'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["backendEnvironmentArn"] = args ? args.backendEnvironmentArn : undefined;
            inputs["basicAuthCredentials"] = args ? args.basicAuthCredentials : undefined;
            inputs["branchName"] = args ? args.branchName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enableAutoBuild"] = args ? args.enableAutoBuild : undefined;
            inputs["enableBasicAuth"] = args ? args.enableBasicAuth : undefined;
            inputs["enableNotification"] = args ? args.enableNotification : undefined;
            inputs["enablePerformanceMode"] = args ? args.enablePerformanceMode : undefined;
            inputs["enablePullRequestPreview"] = args ? args.enablePullRequestPreview : undefined;
            inputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            inputs["framework"] = args ? args.framework : undefined;
            inputs["pullRequestEnvironmentName"] = args ? args.pullRequestEnvironmentName : undefined;
            inputs["stage"] = args ? args.stage : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["associatedResources"] = undefined /*out*/;
            inputs["customDomains"] = undefined /*out*/;
            inputs["destinationBranch"] = undefined /*out*/;
            inputs["sourceBranch"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Branch.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Branch resources.
 */
export interface BranchState {
    /**
     * The unique ID for an Amplify app.
     */
    appId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the branch.
     */
    arn?: pulumi.Input<string>;
    /**
     * A list of custom resources that are linked to this branch.
     */
    associatedResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
     */
    backendEnvironmentArn?: pulumi.Input<string>;
    /**
     * The basic authorization credentials for the branch.
     */
    basicAuthCredentials?: pulumi.Input<string>;
    /**
     * The name for the branch.
     */
    branchName?: pulumi.Input<string>;
    /**
     * The custom domains for the branch.
     */
    customDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description for the branch.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination branch if the branch is a pull request branch.
     */
    destinationBranch?: pulumi.Input<string>;
    /**
     * The display name for a branch. This is used as the default domain prefix.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Enables auto building for the branch.
     */
    enableAutoBuild?: pulumi.Input<boolean>;
    /**
     * Enables basic authorization for the branch.
     */
    enableBasicAuth?: pulumi.Input<boolean>;
    /**
     * Enables notifications for the branch.
     */
    enableNotification?: pulumi.Input<boolean>;
    /**
     * Enables performance mode for the branch.
     */
    enablePerformanceMode?: pulumi.Input<boolean>;
    /**
     * Enables pull request previews for this branch.
     */
    enablePullRequestPreview?: pulumi.Input<boolean>;
    /**
     * The environment variables for the branch.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The framework for the branch.
     */
    framework?: pulumi.Input<string>;
    /**
     * The Amplify environment name for the pull request.
     */
    pullRequestEnvironmentName?: pulumi.Input<string>;
    /**
     * The source branch if the branch is a pull request branch.
     */
    sourceBranch?: pulumi.Input<string>;
    /**
     * Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
     */
    stage?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The content Time To Live (TTL) for the website in seconds.
     */
    ttl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Branch resource.
 */
export interface BranchArgs {
    /**
     * The unique ID for an Amplify app.
     */
    appId: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
     */
    backendEnvironmentArn?: pulumi.Input<string>;
    /**
     * The basic authorization credentials for the branch.
     */
    basicAuthCredentials?: pulumi.Input<string>;
    /**
     * The name for the branch.
     */
    branchName: pulumi.Input<string>;
    /**
     * The description for the branch.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name for a branch. This is used as the default domain prefix.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Enables auto building for the branch.
     */
    enableAutoBuild?: pulumi.Input<boolean>;
    /**
     * Enables basic authorization for the branch.
     */
    enableBasicAuth?: pulumi.Input<boolean>;
    /**
     * Enables notifications for the branch.
     */
    enableNotification?: pulumi.Input<boolean>;
    /**
     * Enables performance mode for the branch.
     */
    enablePerformanceMode?: pulumi.Input<boolean>;
    /**
     * Enables pull request previews for this branch.
     */
    enablePullRequestPreview?: pulumi.Input<boolean>;
    /**
     * The environment variables for the branch.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The framework for the branch.
     */
    framework?: pulumi.Input<string>;
    /**
     * The Amplify environment name for the pull request.
     */
    pullRequestEnvironmentName?: pulumi.Input<string>;
    /**
     * Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
     */
    stage?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The content Time To Live (TTL) for the website in seconds.
     */
    ttl?: pulumi.Input<string>;
}
