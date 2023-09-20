// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import {Application, ApplicationVersion} from "./index";

/**
 * Provides an Elastic Beanstalk Environment Resource. Elastic Beanstalk allows
 * you to deploy and manage applications in the AWS cloud without worrying about
 * the infrastructure that runs those applications.
 *
 * Environments are often things such as `development`, `integration`, or
 * `production`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const tftest = new aws.elasticbeanstalk.Application("tftest", {description: "tf-test-desc"});
 * const tfenvtest = new aws.elasticbeanstalk.Environment("tfenvtest", {
 *     application: tftest.name,
 *     solutionStackName: "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
 * });
 * ```
 * ## Option Settings
 *
 * Some options can be stack-specific, check [AWS Docs](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html)
 * for supported options and examples.
 *
 * The `setting` and `allSettings` mappings support the following format:
 *
 * * `namespace` - unique namespace identifying the option's associated AWS resource
 * * `name` - name of the configuration option
 * * `value` - value for the configuration option
 * * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)
 *
 * ### Example With Options
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const tftest = new aws.elasticbeanstalk.Application("tftest", {description: "tf-test-desc"});
 * const tfenvtest = new aws.elasticbeanstalk.Environment("tfenvtest", {
 *     application: tftest.name,
 *     solutionStackName: "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
 *     settings: [
 *         {
 *             namespace: "aws:ec2:vpc",
 *             name: "VPCId",
 *             value: "vpc-xxxxxxxx",
 *         },
 *         {
 *             namespace: "aws:ec2:vpc",
 *             name: "Subnets",
 *             value: "subnet-xxxxxxxx",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Elastic Beanstalk Environments using the `id`. For example:
 *
 * ```sh
 *  $ pulumi import aws:elasticbeanstalk/environment:Environment prodenv e-rpqsewtp2j
 * ```
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentState, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticbeanstalk/environment:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * List of all option settings configured in this Environment. These
     * are a combination of default settings and their overrides from `setting` in
     * the configuration.
     */
    public /*out*/ readonly allSettings!: pulumi.Output<outputs.elasticbeanstalk.EnvironmentAllSetting[]>;
    /**
     * Name of the application that contains the version
     * to be deployed
     */
    public readonly application!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The autoscaling groups used by this Environment.
     */
    public /*out*/ readonly autoscalingGroups!: pulumi.Output<string[]>;
    /**
     * Fully qualified DNS name for this Environment.
     */
    public /*out*/ readonly cname!: pulumi.Output<string>;
    /**
     * Prefix to use for the fully qualified DNS name of
     * the Environment.
     */
    public readonly cnamePrefix!: pulumi.Output<string>;
    /**
     * Short description of the Environment
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The URL to the Load Balancer for this Environment
     */
    public /*out*/ readonly endpointUrl!: pulumi.Output<string>;
    /**
     * Instances used by this Environment.
     */
    public /*out*/ readonly instances!: pulumi.Output<string[]>;
    /**
     * Launch configurations in use by this Environment.
     */
    public /*out*/ readonly launchConfigurations!: pulumi.Output<string[]>;
    /**
     * Elastic load balancers in use by this Environment.
     */
    public /*out*/ readonly loadBalancers!: pulumi.Output<string[]>;
    /**
     * A unique name for this Environment. This name is used
     * in the application URL
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
     * to use in deployment
     */
    public readonly platformArn!: pulumi.Output<string>;
    /**
     * The time between polling the AWS API to
     * check if changes have been applied. Use this to adjust the rate of API calls
     * for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
     * use the default behavior, which is an exponential backoff
     */
    public readonly pollInterval!: pulumi.Output<string | undefined>;
    /**
     * SQS queues in use by this Environment.
     */
    public /*out*/ readonly queues!: pulumi.Output<string[]>;
    /**
     * Option settings to configure the new Environment. These
     * override specific values that are set as defaults. The format is detailed
     * below in Option Settings
     */
    public readonly settings!: pulumi.Output<outputs.elasticbeanstalk.EnvironmentSetting[] | undefined>;
    /**
     * A solution stack to base your environment
     * off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
     */
    public readonly solutionStackName!: pulumi.Output<string>;
    /**
     * A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of the Elastic Beanstalk Configuration
     * template to use in deployment
     */
    public readonly templateName!: pulumi.Output<string | undefined>;
    /**
     * Elastic Beanstalk Environment tier. Valid values are `Worker`
     * or `WebServer`. If tier is left blank `WebServer` will be used.
     */
    public readonly tier!: pulumi.Output<string | undefined>;
    /**
     * Autoscaling triggers in use by this Environment.
     */
    public /*out*/ readonly triggers!: pulumi.Output<string[]>;
    /**
     * The name of the Elastic Beanstalk Application Version
     * to use in deployment.
     */
    public readonly version!: pulumi.Output<ApplicationVersion>;
    /**
     * The maximum
     * [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
     * wait for an Elastic Beanstalk Environment to be in a ready state before timing
     * out.
     */
    public readonly waitForReadyTimeout!: pulumi.Output<string | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentArgs | EnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnvironmentState | undefined;
            resourceInputs["allSettings"] = state ? state.allSettings : undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["autoscalingGroups"] = state ? state.autoscalingGroups : undefined;
            resourceInputs["cname"] = state ? state.cname : undefined;
            resourceInputs["cnamePrefix"] = state ? state.cnamePrefix : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endpointUrl"] = state ? state.endpointUrl : undefined;
            resourceInputs["instances"] = state ? state.instances : undefined;
            resourceInputs["launchConfigurations"] = state ? state.launchConfigurations : undefined;
            resourceInputs["loadBalancers"] = state ? state.loadBalancers : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platformArn"] = state ? state.platformArn : undefined;
            resourceInputs["pollInterval"] = state ? state.pollInterval : undefined;
            resourceInputs["queues"] = state ? state.queues : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["solutionStackName"] = state ? state.solutionStackName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
            resourceInputs["tier"] = state ? state.tier : undefined;
            resourceInputs["triggers"] = state ? state.triggers : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["waitForReadyTimeout"] = state ? state.waitForReadyTimeout : undefined;
        } else {
            const args = argsOrState as EnvironmentArgs | undefined;
            if ((!args || args.application === undefined) && !opts.urn) {
                throw new Error("Missing required property 'application'");
            }
            resourceInputs["application"] = args ? args.application : undefined;
            resourceInputs["cnamePrefix"] = args ? args.cnamePrefix : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["platformArn"] = args ? args.platformArn : undefined;
            resourceInputs["pollInterval"] = args ? args.pollInterval : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["solutionStackName"] = args ? args.solutionStackName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["waitForReadyTimeout"] = args ? args.waitForReadyTimeout : undefined;
            resourceInputs["allSettings"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["autoscalingGroups"] = undefined /*out*/;
            resourceInputs["cname"] = undefined /*out*/;
            resourceInputs["endpointUrl"] = undefined /*out*/;
            resourceInputs["instances"] = undefined /*out*/;
            resourceInputs["launchConfigurations"] = undefined /*out*/;
            resourceInputs["loadBalancers"] = undefined /*out*/;
            resourceInputs["queues"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["triggers"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Environment resources.
 */
export interface EnvironmentState {
    /**
     * List of all option settings configured in this Environment. These
     * are a combination of default settings and their overrides from `setting` in
     * the configuration.
     */
    allSettings?: pulumi.Input<pulumi.Input<inputs.elasticbeanstalk.EnvironmentAllSetting>[]>;
    /**
     * Name of the application that contains the version
     * to be deployed
     */
    application?: pulumi.Input<string | Application>;
    arn?: pulumi.Input<string>;
    /**
     * The autoscaling groups used by this Environment.
     */
    autoscalingGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Fully qualified DNS name for this Environment.
     */
    cname?: pulumi.Input<string>;
    /**
     * Prefix to use for the fully qualified DNS name of
     * the Environment.
     */
    cnamePrefix?: pulumi.Input<string>;
    /**
     * Short description of the Environment
     */
    description?: pulumi.Input<string>;
    /**
     * The URL to the Load Balancer for this Environment
     */
    endpointUrl?: pulumi.Input<string>;
    /**
     * Instances used by this Environment.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Launch configurations in use by this Environment.
     */
    launchConfigurations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Elastic load balancers in use by this Environment.
     */
    loadBalancers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique name for this Environment. This name is used
     * in the application URL
     */
    name?: pulumi.Input<string>;
    /**
     * The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
     * to use in deployment
     */
    platformArn?: pulumi.Input<string>;
    /**
     * The time between polling the AWS API to
     * check if changes have been applied. Use this to adjust the rate of API calls
     * for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
     * use the default behavior, which is an exponential backoff
     */
    pollInterval?: pulumi.Input<string>;
    /**
     * SQS queues in use by this Environment.
     */
    queues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Option settings to configure the new Environment. These
     * override specific values that are set as defaults. The format is detailed
     * below in Option Settings
     */
    settings?: pulumi.Input<pulumi.Input<inputs.elasticbeanstalk.EnvironmentSetting>[]>;
    /**
     * A solution stack to base your environment
     * off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
     */
    solutionStackName?: pulumi.Input<string>;
    /**
     * A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Elastic Beanstalk Configuration
     * template to use in deployment
     */
    templateName?: pulumi.Input<string>;
    /**
     * Elastic Beanstalk Environment tier. Valid values are `Worker`
     * or `WebServer`. If tier is left blank `WebServer` will be used.
     */
    tier?: pulumi.Input<string>;
    /**
     * Autoscaling triggers in use by this Environment.
     */
    triggers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Elastic Beanstalk Application Version
     * to use in deployment.
     */
    version?: pulumi.Input<ApplicationVersion>;
    /**
     * The maximum
     * [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
     * wait for an Elastic Beanstalk Environment to be in a ready state before timing
     * out.
     */
    waitForReadyTimeout?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Name of the application that contains the version
     * to be deployed
     */
    application: pulumi.Input<string | Application>;
    /**
     * Prefix to use for the fully qualified DNS name of
     * the Environment.
     */
    cnamePrefix?: pulumi.Input<string>;
    /**
     * Short description of the Environment
     */
    description?: pulumi.Input<string>;
    /**
     * A unique name for this Environment. This name is used
     * in the application URL
     */
    name?: pulumi.Input<string>;
    /**
     * The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
     * to use in deployment
     */
    platformArn?: pulumi.Input<string>;
    /**
     * The time between polling the AWS API to
     * check if changes have been applied. Use this to adjust the rate of API calls
     * for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
     * use the default behavior, which is an exponential backoff
     */
    pollInterval?: pulumi.Input<string>;
    /**
     * Option settings to configure the new Environment. These
     * override specific values that are set as defaults. The format is detailed
     * below in Option Settings
     */
    settings?: pulumi.Input<pulumi.Input<inputs.elasticbeanstalk.EnvironmentSetting>[]>;
    /**
     * A solution stack to base your environment
     * off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
     */
    solutionStackName?: pulumi.Input<string>;
    /**
     * A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the Elastic Beanstalk Configuration
     * template to use in deployment
     */
    templateName?: pulumi.Input<string>;
    /**
     * Elastic Beanstalk Environment tier. Valid values are `Worker`
     * or `WebServer`. If tier is left blank `WebServer` will be used.
     */
    tier?: pulumi.Input<string>;
    /**
     * The name of the Elastic Beanstalk Application Version
     * to use in deployment.
     */
    version?: pulumi.Input<ApplicationVersion>;
    /**
     * The maximum
     * [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
     * wait for an Elastic Beanstalk Environment to be in a ready state before timing
     * out.
     */
    waitForReadyTimeout?: pulumi.Input<string>;
}
