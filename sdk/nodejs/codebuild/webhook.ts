// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages a CodeBuild webhook, which is an endpoint accepted by the CodeBuild service to trigger builds from source code repositories. Depending on the source type of the CodeBuild project, the CodeBuild service may also automatically create and delete the actual repository webhook as well.
 *
 * ## Example Usage
 * ### Bitbucket and GitHub
 *
 * When working with [Bitbucket](https://bitbucket.org) and [GitHub](https://github.com) source CodeBuild webhooks, the CodeBuild service will automatically create (on `aws.codebuild.Webhook` resource creation) and delete (on `aws.codebuild.Webhook` resource deletion) the Bitbucket/GitHub repository webhook using its granted OAuth permissions. This behavior cannot be controlled by this provider.
 *
 * > **Note:** The AWS account that this provider uses to create this resource *must* have authorized CodeBuild to access Bitbucket/GitHub's OAuth API in each applicable region. This is a manual step that must be done *before* creating webhooks with this resource. If OAuth is not configured, AWS will return an error similar to `ResourceNotFoundException: Could not find access token for server type github`. More information can be found in the CodeBuild User Guide for [Bitbucket](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-bitbucket-pull-request.html) and [GitHub](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-pull-request.html).
 *
 * > **Note:** Further managing the automatically created Bitbucket/GitHub webhook with the `bitbucketHook`/`githubRepositoryWebhook` resource is only possible with importing that resource after creation of the `aws.codebuild.Webhook` resource. The CodeBuild API does not ever provide the `secret` attribute for the `aws.codebuild.Webhook` resource in this scenario.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.codebuild.Webhook("example", {
 *     projectName: aws_codebuild_project.example.name,
 *     buildType: "BUILD",
 *     filterGroups: [{
 *         filters: [
 *             {
 *                 type: "EVENT",
 *                 pattern: "PUSH",
 *             },
 *             {
 *                 type: "BASE_REF",
 *                 pattern: "master",
 *             },
 *         ],
 *     }],
 * });
 * ```
 * ### GitHub Enterprise
 *
 * When working with [GitHub Enterprise](https://enterprise.github.com/) source CodeBuild webhooks, the GHE repository webhook must be separately managed (e.g., manually or with the `githubRepositoryWebhook` resource).
 *
 * More information creating webhooks with GitHub Enterprise can be found in the [CodeBuild User Guide](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-enterprise.html).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as github from "@pulumi/github";
 *
 * const exampleWebhook = new aws.codebuild.Webhook("exampleWebhook", {projectName: aws_codebuild_project.example.name});
 * const exampleRepositoryWebhook = new github.RepositoryWebhook("exampleRepositoryWebhook", {
 *     active: true,
 *     events: ["push"],
 *     repository: github_repository.example.name,
 *     configuration: {
 *         url: exampleWebhook.payloadUrl,
 *         secret: exampleWebhook.secret,
 *         contentType: "json",
 *         insecureSsl: false,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_codebuild_webhook.example
 *
 *  id = "MyProjectName" } Using `pulumi import`, import CodeBuild Webhooks using the CodeBuild Project name. For exampleconsole % pulumi import aws_codebuild_webhook.example MyProjectName
 */
export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookState, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codebuild/webhook:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    /**
     * A regular expression used to determine which branches get built. Default is all branches are built. We recommend using `filterGroup` over `branchFilter`.
     */
    public readonly branchFilter!: pulumi.Output<string | undefined>;
    /**
     * The type of build this webhook will trigger. Valid values for this parameter are: `BUILD`, `BUILD_BATCH`.
     */
    public readonly buildType!: pulumi.Output<string | undefined>;
    /**
     * Information about the webhook's trigger. Filter group blocks are documented below.
     */
    public readonly filterGroups!: pulumi.Output<outputs.codebuild.WebhookFilterGroup[] | undefined>;
    /**
     * The CodeBuild endpoint where webhook events are sent.
     */
    public /*out*/ readonly payloadUrl!: pulumi.Output<string>;
    /**
     * The name of the build project.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;
    /**
     * The URL to the webhook.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebhookArgs | WebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebhookState | undefined;
            resourceInputs["branchFilter"] = state ? state.branchFilter : undefined;
            resourceInputs["buildType"] = state ? state.buildType : undefined;
            resourceInputs["filterGroups"] = state ? state.filterGroups : undefined;
            resourceInputs["payloadUrl"] = state ? state.payloadUrl : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as WebhookArgs | undefined;
            if ((!args || args.projectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectName'");
            }
            resourceInputs["branchFilter"] = args ? args.branchFilter : undefined;
            resourceInputs["buildType"] = args ? args.buildType : undefined;
            resourceInputs["filterGroups"] = args ? args.filterGroups : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["payloadUrl"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Webhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Webhook resources.
 */
export interface WebhookState {
    /**
     * A regular expression used to determine which branches get built. Default is all branches are built. We recommend using `filterGroup` over `branchFilter`.
     */
    branchFilter?: pulumi.Input<string>;
    /**
     * The type of build this webhook will trigger. Valid values for this parameter are: `BUILD`, `BUILD_BATCH`.
     */
    buildType?: pulumi.Input<string>;
    /**
     * Information about the webhook's trigger. Filter group blocks are documented below.
     */
    filterGroups?: pulumi.Input<pulumi.Input<inputs.codebuild.WebhookFilterGroup>[]>;
    /**
     * The CodeBuild endpoint where webhook events are sent.
     */
    payloadUrl?: pulumi.Input<string>;
    /**
     * The name of the build project.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
     */
    secret?: pulumi.Input<string>;
    /**
     * The URL to the webhook.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    /**
     * A regular expression used to determine which branches get built. Default is all branches are built. We recommend using `filterGroup` over `branchFilter`.
     */
    branchFilter?: pulumi.Input<string>;
    /**
     * The type of build this webhook will trigger. Valid values for this parameter are: `BUILD`, `BUILD_BATCH`.
     */
    buildType?: pulumi.Input<string>;
    /**
     * Information about the webhook's trigger. Filter group blocks are documented below.
     */
    filterGroups?: pulumi.Input<pulumi.Input<inputs.codebuild.WebhookFilterGroup>[]>;
    /**
     * The name of the build project.
     */
    projectName: pulumi.Input<string>;
}
