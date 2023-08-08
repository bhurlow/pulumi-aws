// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing AWS Audit Manager Account Registration.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.auditmanager.AccountRegistration("example", {});
 * ```
 * ### Deregister On Destroy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.auditmanager.AccountRegistration("example", {deregisterOnDestroy: true});
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_auditmanager_account_registration.example
 *
 *  id = "us-east-1" } Using `pulumi import`, import Audit Manager Account Registration resources using the `id`. For exampleconsole % pulumi import aws_auditmanager_account_registration.example us-east-1
 */
export class AccountRegistration extends pulumi.CustomResource {
    /**
     * Get an existing AccountRegistration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountRegistrationState, opts?: pulumi.CustomResourceOptions): AccountRegistration {
        return new AccountRegistration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:auditmanager/accountRegistration:AccountRegistration';

    /**
     * Returns true if the given object is an instance of AccountRegistration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountRegistration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountRegistration.__pulumiType;
    }

    /**
     * Identifier for the delegated administrator account.
     */
    public readonly delegatedAdminAccount!: pulumi.Output<string | undefined>;
    /**
     * Flag to deregister AuditManager in the account upon destruction. Defaults to `false` (ie. AuditManager will remain active in the account, even if this resource is removed).
     */
    public readonly deregisterOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * KMS key identifier.
     */
    public readonly kmsKey!: pulumi.Output<string | undefined>;
    /**
     * Status of the account registration request.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AccountRegistration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountRegistrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountRegistrationArgs | AccountRegistrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountRegistrationState | undefined;
            resourceInputs["delegatedAdminAccount"] = state ? state.delegatedAdminAccount : undefined;
            resourceInputs["deregisterOnDestroy"] = state ? state.deregisterOnDestroy : undefined;
            resourceInputs["kmsKey"] = state ? state.kmsKey : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AccountRegistrationArgs | undefined;
            resourceInputs["delegatedAdminAccount"] = args ? args.delegatedAdminAccount : undefined;
            resourceInputs["deregisterOnDestroy"] = args ? args.deregisterOnDestroy : undefined;
            resourceInputs["kmsKey"] = args ? args.kmsKey : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountRegistration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountRegistration resources.
 */
export interface AccountRegistrationState {
    /**
     * Identifier for the delegated administrator account.
     */
    delegatedAdminAccount?: pulumi.Input<string>;
    /**
     * Flag to deregister AuditManager in the account upon destruction. Defaults to `false` (ie. AuditManager will remain active in the account, even if this resource is removed).
     */
    deregisterOnDestroy?: pulumi.Input<boolean>;
    /**
     * KMS key identifier.
     */
    kmsKey?: pulumi.Input<string>;
    /**
     * Status of the account registration request.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountRegistration resource.
 */
export interface AccountRegistrationArgs {
    /**
     * Identifier for the delegated administrator account.
     */
    delegatedAdminAccount?: pulumi.Input<string>;
    /**
     * Flag to deregister AuditManager in the account upon destruction. Defaults to `false` (ie. AuditManager will remain active in the account, even if this resource is removed).
     */
    deregisterOnDestroy?: pulumi.Input<boolean>;
    /**
     * KMS key identifier.
     */
    kmsKey?: pulumi.Input<string>;
}
