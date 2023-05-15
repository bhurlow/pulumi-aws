// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an App Runner Service.
 *
 * ## Example Usage
 * ### Service with a Code Repository Source
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apprunner.Service("example", {
 *     serviceName: "example",
 *     sourceConfiguration: {
 *         authenticationConfiguration: {
 *             connectionArn: aws_apprunner_connection.example.arn,
 *         },
 *         codeRepository: {
 *             codeConfiguration: {
 *                 codeConfigurationValues: {
 *                     buildCommand: "python setup.py develop",
 *                     port: "8000",
 *                     runtime: "PYTHON_3",
 *                     startCommand: "python runapp.py",
 *                 },
 *                 configurationSource: "API",
 *             },
 *             repositoryUrl: "https://github.com/example/my-example-python-app",
 *             sourceCodeVersion: {
 *                 type: "BRANCH",
 *                 value: "main",
 *             },
 *         },
 *     },
 *     networkConfiguration: {
 *         egressConfiguration: {
 *             egressType: "VPC",
 *             vpcConnectorArn: aws_apprunner_vpc_connector.connector.arn,
 *         },
 *     },
 *     tags: {
 *         Name: "example-apprunner-service",
 *     },
 * });
 * ```
 * ### Service with an Image Repository Source
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apprunner.Service("example", {
 *     serviceName: "example",
 *     sourceConfiguration: {
 *         autoDeploymentsEnabled: false,
 *         imageRepository: {
 *             imageConfiguration: {
 *                 port: "8000",
 *             },
 *             imageIdentifier: "public.ecr.aws/aws-containers/hello-app-runner:latest",
 *             imageRepositoryType: "ECR_PUBLIC",
 *         },
 *     },
 *     tags: {
 *         Name: "example-apprunner-service",
 *     },
 * });
 * ```
 * ### Service with Observability Configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleObservabilityConfiguration = new aws.apprunner.ObservabilityConfiguration("exampleObservabilityConfiguration", {
 *     observabilityConfigurationName: "example",
 *     traceConfiguration: {
 *         vendor: "AWSXRAY",
 *     },
 * });
 * const exampleService = new aws.apprunner.Service("exampleService", {
 *     serviceName: "example",
 *     observabilityConfiguration: {
 *         observabilityConfigurationArn: exampleObservabilityConfiguration.arn,
 *         observabilityEnabled: true,
 *     },
 *     sourceConfiguration: {
 *         imageRepository: {
 *             imageConfiguration: {
 *                 port: "8000",
 *             },
 *             imageIdentifier: "public.ecr.aws/aws-containers/hello-app-runner:latest",
 *             imageRepositoryType: "ECR_PUBLIC",
 *         },
 *         autoDeploymentsEnabled: false,
 *     },
 *     tags: {
 *         Name: "example-apprunner-service",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * App Runner Services can be imported by using the `arn`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:apprunner/service:Service example arn:aws:apprunner:us-east-1:1234567890:service/example/0a03292a89764e5882c41d8f991c82fe
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apprunner/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * ARN of the App Runner service.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
     */
    public readonly autoScalingConfigurationArn!: pulumi.Output<string>;
    /**
     * An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
     */
    public readonly encryptionConfiguration!: pulumi.Output<outputs.apprunner.ServiceEncryptionConfiguration | undefined>;
    /**
     * Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
     */
    public readonly healthCheckConfiguration!: pulumi.Output<outputs.apprunner.ServiceHealthCheckConfiguration>;
    /**
     * The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
     */
    public readonly instanceConfiguration!: pulumi.Output<outputs.apprunner.ServiceInstanceConfiguration>;
    /**
     * Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
     */
    public readonly networkConfiguration!: pulumi.Output<outputs.apprunner.ServiceNetworkConfiguration>;
    /**
     * The observability configuration of your service. See Observability Configuration below for more details.
     */
    public readonly observabilityConfiguration!: pulumi.Output<outputs.apprunner.ServiceObservabilityConfiguration | undefined>;
    /**
     * An alphanumeric ID that App Runner generated for this service. Unique within the AWS Region.
     */
    public /*out*/ readonly serviceId!: pulumi.Output<string>;
    /**
     * Name of the service.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.
     */
    public /*out*/ readonly serviceUrl!: pulumi.Output<string>;
    /**
     * The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
     */
    public readonly sourceConfiguration!: pulumi.Output<outputs.apprunner.ServiceSourceConfiguration>;
    /**
     * Current state of the App Runner service.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["autoScalingConfigurationArn"] = state ? state.autoScalingConfigurationArn : undefined;
            resourceInputs["encryptionConfiguration"] = state ? state.encryptionConfiguration : undefined;
            resourceInputs["healthCheckConfiguration"] = state ? state.healthCheckConfiguration : undefined;
            resourceInputs["instanceConfiguration"] = state ? state.instanceConfiguration : undefined;
            resourceInputs["networkConfiguration"] = state ? state.networkConfiguration : undefined;
            resourceInputs["observabilityConfiguration"] = state ? state.observabilityConfiguration : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["serviceUrl"] = state ? state.serviceUrl : undefined;
            resourceInputs["sourceConfiguration"] = state ? state.sourceConfiguration : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.sourceConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceConfiguration'");
            }
            resourceInputs["autoScalingConfigurationArn"] = args ? args.autoScalingConfigurationArn : undefined;
            resourceInputs["encryptionConfiguration"] = args ? args.encryptionConfiguration : undefined;
            resourceInputs["healthCheckConfiguration"] = args ? args.healthCheckConfiguration : undefined;
            resourceInputs["instanceConfiguration"] = args ? args.instanceConfiguration : undefined;
            resourceInputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            resourceInputs["observabilityConfiguration"] = args ? args.observabilityConfiguration : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["sourceConfiguration"] = args ? args.sourceConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["serviceUrl"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * ARN of the App Runner service.
     */
    arn?: pulumi.Input<string>;
    /**
     * ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
     */
    autoScalingConfigurationArn?: pulumi.Input<string>;
    /**
     * An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
     */
    encryptionConfiguration?: pulumi.Input<inputs.apprunner.ServiceEncryptionConfiguration>;
    /**
     * Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
     */
    healthCheckConfiguration?: pulumi.Input<inputs.apprunner.ServiceHealthCheckConfiguration>;
    /**
     * The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
     */
    instanceConfiguration?: pulumi.Input<inputs.apprunner.ServiceInstanceConfiguration>;
    /**
     * Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
     */
    networkConfiguration?: pulumi.Input<inputs.apprunner.ServiceNetworkConfiguration>;
    /**
     * The observability configuration of your service. See Observability Configuration below for more details.
     */
    observabilityConfiguration?: pulumi.Input<inputs.apprunner.ServiceObservabilityConfiguration>;
    /**
     * An alphanumeric ID that App Runner generated for this service. Unique within the AWS Region.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * Name of the service.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.
     */
    serviceUrl?: pulumi.Input<string>;
    /**
     * The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
     */
    sourceConfiguration?: pulumi.Input<inputs.apprunner.ServiceSourceConfiguration>;
    /**
     * Current state of the App Runner service.
     */
    status?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * ARN of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner associates the latest revision of a default auto scaling configuration.
     */
    autoScalingConfigurationArn?: pulumi.Input<string>;
    /**
     * An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default, App Runner uses an AWS managed CMK. See Encryption Configuration below for more details.
     */
    encryptionConfiguration?: pulumi.Input<inputs.apprunner.ServiceEncryptionConfiguration>;
    /**
     * Settings of the health check that AWS App Runner performs to monitor the health of your service. See Health Check Configuration below for more details.
     */
    healthCheckConfiguration?: pulumi.Input<inputs.apprunner.ServiceHealthCheckConfiguration>;
    /**
     * The runtime configuration of instances (scaling units) of the App Runner service. See Instance Configuration below for more details.
     */
    instanceConfiguration?: pulumi.Input<inputs.apprunner.ServiceInstanceConfiguration>;
    /**
     * Configuration settings related to network traffic of the web application that the App Runner service runs. See Network Configuration below for more details.
     */
    networkConfiguration?: pulumi.Input<inputs.apprunner.ServiceNetworkConfiguration>;
    /**
     * The observability configuration of your service. See Observability Configuration below for more details.
     */
    observabilityConfiguration?: pulumi.Input<inputs.apprunner.ServiceObservabilityConfiguration>;
    /**
     * Name of the service.
     */
    serviceName: pulumi.Input<string>;
    /**
     * The source to deploy to the App Runner service. Can be a code or an image repository. See Source Configuration below for more details.
     */
    sourceConfiguration: pulumi.Input<inputs.apprunner.ServiceSourceConfiguration>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
