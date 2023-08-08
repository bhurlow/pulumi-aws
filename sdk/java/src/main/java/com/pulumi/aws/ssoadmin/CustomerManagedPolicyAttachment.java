// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ssoadmin;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ssoadmin.CustomerManagedPolicyAttachmentArgs;
import com.pulumi.aws.ssoadmin.inputs.CustomerManagedPolicyAttachmentState;
import com.pulumi.aws.ssoadmin.outputs.CustomerManagedPolicyAttachmentCustomerManagedPolicyReference;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a customer managed policy attachment for a Single Sign-On (SSO) Permission Set resource
 * 
 * &gt; **NOTE:** Creating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_ssoadmin_customer_managed_policy_attachment.example
 * 
 *  id = &#34;TestPolicy,/,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72&#34; } Using `pulumi import`, import SSO Managed Policy Attachments using the `name`, `path`, `permission_set_arn`, and `instance_arn` separated by a comma (`,`). For exampleconsole % pulumi import aws_ssoadmin_customer_managed_policy_attachment.example TestPolicy,/,arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
 * 
 */
@ResourceType(type="aws:ssoadmin/customerManagedPolicyAttachment:CustomerManagedPolicyAttachment")
public class CustomerManagedPolicyAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the name and path of a customer managed policy. See below.
     * 
     */
    @Export(name="customerManagedPolicyReference", refs={CustomerManagedPolicyAttachmentCustomerManagedPolicyReference.class}, tree="[0]")
    private Output<CustomerManagedPolicyAttachmentCustomerManagedPolicyReference> customerManagedPolicyReference;

    /**
     * @return Specifies the name and path of a customer managed policy. See below.
     * 
     */
    public Output<CustomerManagedPolicyAttachmentCustomerManagedPolicyReference> customerManagedPolicyReference() {
        return this.customerManagedPolicyReference;
    }
    /**
     * The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
     * 
     */
    @Export(name="instanceArn", refs={String.class}, tree="[0]")
    private Output<String> instanceArn;

    /**
     * @return The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
     * 
     */
    public Output<String> instanceArn() {
        return this.instanceArn;
    }
    /**
     * The Amazon Resource Name (ARN) of the Permission Set.
     * 
     */
    @Export(name="permissionSetArn", refs={String.class}, tree="[0]")
    private Output<String> permissionSetArn;

    /**
     * @return The Amazon Resource Name (ARN) of the Permission Set.
     * 
     */
    public Output<String> permissionSetArn() {
        return this.permissionSetArn;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomerManagedPolicyAttachment(String name) {
        this(name, CustomerManagedPolicyAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomerManagedPolicyAttachment(String name, CustomerManagedPolicyAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomerManagedPolicyAttachment(String name, CustomerManagedPolicyAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ssoadmin/customerManagedPolicyAttachment:CustomerManagedPolicyAttachment", name, args == null ? CustomerManagedPolicyAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomerManagedPolicyAttachment(String name, Output<String> id, @Nullable CustomerManagedPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ssoadmin/customerManagedPolicyAttachment:CustomerManagedPolicyAttachment", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CustomerManagedPolicyAttachment get(String name, Output<String> id, @Nullable CustomerManagedPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomerManagedPolicyAttachment(name, id, state, options);
    }
}
