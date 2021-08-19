// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Secrets Manager secret metadata. To manage secret rotation, see the `aws.secretsmanager.SecretRotation` resource. To manage a secret value, see the `aws.secretsmanager.SecretVersion` resource.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.secretsmanager.Secret("example", {});
 * ```
 * ### Rotation Configuration
 *
 * To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g. RDS) or deploying a custom Lambda function.
 *
 * > **NOTE:** Configuring rotation causes the secret to rotate once as soon as you store the secret. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.
 *
 * > **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const rotation_example = new aws.secretsmanager.Secret("rotation-example", {
 *     rotationLambdaArn: aws_lambda_function.example.arn,
 *     rotationRules: {
 *         automaticallyAfterDays: 7,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * `aws_secretsmanager_secret` can be imported by using the secret Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:secretsmanager/secret:Secret example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
 * ```
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:secretsmanager/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Secret.__pulumiType;
    }

    /**
     * ARN of the secret.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the secret.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly forceOverwriteReplicaSecret!: pulumi.Output<boolean | undefined>;
    /**
     * ARN, Key ID, or Alias.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * Number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     */
    public readonly recoveryWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * Configuration block to support secret replication. See details below.
     */
    public readonly replicas!: pulumi.Output<outputs.secretsmanager.SecretReplica[]>;
    /**
     * Whether automatic rotation is enabled for this secret.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    public /*out*/ readonly rotationEnabled!: pulumi.Output<boolean>;
    /**
     * ARN of the Lambda function that can rotate the secret. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    public readonly rotationLambdaArn!: pulumi.Output<string>;
    /**
     * Configuration block for the rotation configuration of this secret. Defined below. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    public readonly rotationRules!: pulumi.Output<outputs.secretsmanager.SecretRotationRules>;
    /**
     * Key-value map of user-defined tags that are attached to the secret. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["forceOverwriteReplicaSecret"] = state ? state.forceOverwriteReplicaSecret : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["recoveryWindowInDays"] = state ? state.recoveryWindowInDays : undefined;
            inputs["replicas"] = state ? state.replicas : undefined;
            inputs["rotationEnabled"] = state ? state.rotationEnabled : undefined;
            inputs["rotationLambdaArn"] = state ? state.rotationLambdaArn : undefined;
            inputs["rotationRules"] = state ? state.rotationRules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["forceOverwriteReplicaSecret"] = args ? args.forceOverwriteReplicaSecret : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["recoveryWindowInDays"] = args ? args.recoveryWindowInDays : undefined;
            inputs["replicas"] = args ? args.replicas : undefined;
            inputs["rotationLambdaArn"] = args ? args.rotationLambdaArn : undefined;
            inputs["rotationRules"] = args ? args.rotationRules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["rotationEnabled"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Secret.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * ARN of the secret.
     */
    arn?: pulumi.Input<string>;
    /**
     * Description of the secret.
     */
    description?: pulumi.Input<string>;
    forceOverwriteReplicaSecret?: pulumi.Input<boolean>;
    /**
     * ARN, Key ID, or Alias.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
     */
    policy?: pulumi.Input<string>;
    /**
     * Number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     */
    recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * Configuration block to support secret replication. See details below.
     */
    replicas?: pulumi.Input<pulumi.Input<inputs.secretsmanager.SecretReplica>[]>;
    /**
     * Whether automatic rotation is enabled for this secret.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    rotationEnabled?: pulumi.Input<boolean>;
    /**
     * ARN of the Lambda function that can rotate the secret. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    rotationLambdaArn?: pulumi.Input<string>;
    /**
     * Configuration block for the rotation configuration of this secret. Defined below. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    rotationRules?: pulumi.Input<inputs.secretsmanager.SecretRotationRules>;
    /**
     * Key-value map of user-defined tags that are attached to the secret. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * Description of the secret.
     */
    description?: pulumi.Input<string>;
    forceOverwriteReplicaSecret?: pulumi.Input<boolean>;
    /**
     * ARN, Key ID, or Alias.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
     */
    policy?: pulumi.Input<string>;
    /**
     * Number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     */
    recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * Configuration block to support secret replication. See details below.
     */
    replicas?: pulumi.Input<pulumi.Input<inputs.secretsmanager.SecretReplica>[]>;
    /**
     * ARN of the Lambda function that can rotate the secret. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    rotationLambdaArn?: pulumi.Input<string>;
    /**
     * Configuration block for the rotation configuration of this secret. Defined below. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     *
     * @deprecated Use the aws_secretsmanager_secret_rotation resource instead
     */
    rotationRules?: pulumi.Input<inputs.secretsmanager.SecretRotationRules>;
    /**
     * Key-value map of user-defined tags that are attached to the secret. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
