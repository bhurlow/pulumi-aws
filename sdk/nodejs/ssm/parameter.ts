// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an SSM Parameter resource.
 *
 * > **Note:** `overwrite` also makes it possible to overwrite an existing SSM Parameter that's not created by the provider before. This argument has been deprecated and will be removed in v6.0.0 of the provider. For more information on how this affects the behavior of this resource, see this issue comment.
 *
 * ## Example Usage
 * ### Basic example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ssm.Parameter("foo", {
 *     type: "String",
 *     value: "bar",
 * });
 * ```
 * ### Encrypted string using default SSM KMS key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const _default = new aws.rds.Instance("default", {
 *     allocatedStorage: 10,
 *     storageType: "gp2",
 *     engine: "mysql",
 *     engineVersion: "5.7.16",
 *     instanceClass: "db.t2.micro",
 *     dbName: "mydb",
 *     username: "foo",
 *     password: _var.database_master_password,
 *     dbSubnetGroupName: "my_database_subnet_group",
 *     parameterGroupName: "default.mysql5.7",
 * });
 * const secret = new aws.ssm.Parameter("secret", {
 *     description: "The parameter description",
 *     type: "SecureString",
 *     value: _var.database_master_password,
 *     tags: {
 *         environment: "production",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ssm_parameter.my_param
 *
 *  id = "/my_path/my_paramname" } Using `pulumi import`, import SSM Parameters using the parameter store `name`. For exampleconsole % pulumi import aws_ssm_parameter.my_param /my_path/my_paramname
 */
export class Parameter extends pulumi.CustomResource {
    /**
     * Get an existing Parameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ParameterState, opts?: pulumi.CustomResourceOptions): Parameter {
        return new Parameter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/parameter:Parameter';

    /**
     * Returns true if the given object is an instance of Parameter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Parameter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Parameter.__pulumiType;
    }

    /**
     * Regular expression used to validate the parameter value.
     */
    public readonly allowedPattern!: pulumi.Output<string | undefined>;
    /**
     * ARN of the parameter.
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
     */
    public readonly dataType!: pulumi.Output<string>;
    /**
     * Description of the parameter.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the preview. This argument is not valid with a `type` of `SecureString`.
     */
    public readonly insecureValue!: pulumi.Output<string>;
    /**
     * KMS key ID or ARN for encrypting a SecureString.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
     *
     * @deprecated this attribute has been deprecated
     */
    public readonly overwrite!: pulumi.Output<boolean | undefined>;
    /**
     * Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
     */
    public readonly tier!: pulumi.Output<string | undefined>;
    /**
     * Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     *
     * The following arguments are optional:
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Value of the parameter. This value is always marked as sensitive in the plan output, regardless of `type`.
     *
     * > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * Version of the parameter.
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a Parameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ParameterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ParameterArgs | ParameterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ParameterState | undefined;
            resourceInputs["allowedPattern"] = state ? state.allowedPattern : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["dataType"] = state ? state.dataType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["insecureValue"] = state ? state.insecureValue : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overwrite"] = state ? state.overwrite : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["tier"] = state ? state.tier : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ParameterArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["allowedPattern"] = args ? args.allowedPattern : undefined;
            resourceInputs["arn"] = args ? args.arn : undefined;
            resourceInputs["dataType"] = args ? args.dataType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["insecureValue"] = args ? args.insecureValue : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overwrite"] = args ? args.overwrite : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args?.value ? pulumi.secret(args.value) : undefined;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Parameter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Parameter resources.
 */
export interface ParameterState {
    /**
     * Regular expression used to validate the parameter value.
     */
    allowedPattern?: pulumi.Input<string>;
    /**
     * ARN of the parameter.
     */
    arn?: pulumi.Input<string>;
    /**
     * Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
     */
    dataType?: pulumi.Input<string>;
    /**
     * Description of the parameter.
     */
    description?: pulumi.Input<string>;
    /**
     * Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the preview. This argument is not valid with a `type` of `SecureString`.
     */
    insecureValue?: pulumi.Input<string>;
    /**
     * KMS key ID or ARN for encrypting a SecureString.
     */
    keyId?: pulumi.Input<string>;
    /**
     * Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
     */
    name?: pulumi.Input<string>;
    /**
     * Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
     *
     * @deprecated this attribute has been deprecated
     */
    overwrite?: pulumi.Input<boolean>;
    /**
     * Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
     */
    tier?: pulumi.Input<string>;
    /**
     * Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     *
     * The following arguments are optional:
     */
    type?: pulumi.Input<string | enums.ssm.ParameterType>;
    /**
     * Value of the parameter. This value is always marked as sensitive in the plan output, regardless of `type`.
     *
     * > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
     */
    value?: pulumi.Input<string>;
    /**
     * Version of the parameter.
     */
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Parameter resource.
 */
export interface ParameterArgs {
    /**
     * Regular expression used to validate the parameter value.
     */
    allowedPattern?: pulumi.Input<string>;
    /**
     * ARN of the parameter.
     */
    arn?: pulumi.Input<string>;
    /**
     * Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
     */
    dataType?: pulumi.Input<string>;
    /**
     * Description of the parameter.
     */
    description?: pulumi.Input<string>;
    /**
     * Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the preview. This argument is not valid with a `type` of `SecureString`.
     */
    insecureValue?: pulumi.Input<string>;
    /**
     * KMS key ID or ARN for encrypting a SecureString.
     */
    keyId?: pulumi.Input<string>;
    /**
     * Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
     */
    name?: pulumi.Input<string>;
    /**
     * Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
     *
     * @deprecated this attribute has been deprecated
     */
    overwrite?: pulumi.Input<boolean>;
    /**
     * Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
     */
    tier?: pulumi.Input<string>;
    /**
     * Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     *
     * The following arguments are optional:
     */
    type: pulumi.Input<string | enums.ssm.ParameterType>;
    /**
     * Value of the parameter. This value is always marked as sensitive in the plan output, regardless of `type`.
     *
     * > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
     */
    value?: pulumi.Input<string>;
}
