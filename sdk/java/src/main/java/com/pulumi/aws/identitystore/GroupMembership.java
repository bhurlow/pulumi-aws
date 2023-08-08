// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.identitystore;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.identitystore.GroupMembershipArgs;
import com.pulumi.aws.identitystore.inputs.GroupMembershipState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS IdentityStore Group Membership.
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_identitystore_group_membership.example
 * 
 *  id = &#34;d-0000000000/00000000-0000-0000-0000-000000000000&#34; } Using `pulumi import`, import `aws_identitystore_group_membership` using the `identity_store_id/membership_id`. For exampleconsole % pulumi import aws_identitystore_group_membership.example d-0000000000/00000000-0000-0000-0000-000000000000
 * 
 */
@ResourceType(type="aws:identitystore/groupMembership:GroupMembership")
public class GroupMembership extends com.pulumi.resources.CustomResource {
    /**
     * The identifier for a group in the Identity Store.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The identifier for a group in the Identity Store.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * Identity Store ID associated with the Single Sign-On Instance.
     * 
     */
    @Export(name="identityStoreId", refs={String.class}, tree="[0]")
    private Output<String> identityStoreId;

    /**
     * @return Identity Store ID associated with the Single Sign-On Instance.
     * 
     */
    public Output<String> identityStoreId() {
        return this.identityStoreId;
    }
    /**
     * The identifier for a user in the Identity Store.
     * 
     */
    @Export(name="memberId", refs={String.class}, tree="[0]")
    private Output<String> memberId;

    /**
     * @return The identifier for a user in the Identity Store.
     * 
     */
    public Output<String> memberId() {
        return this.memberId;
    }
    /**
     * The identifier of the newly created group membership in the Identity Store.
     * 
     */
    @Export(name="membershipId", refs={String.class}, tree="[0]")
    private Output<String> membershipId;

    /**
     * @return The identifier of the newly created group membership in the Identity Store.
     * 
     */
    public Output<String> membershipId() {
        return this.membershipId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupMembership(String name) {
        this(name, GroupMembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupMembership(String name, GroupMembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupMembership(String name, GroupMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:identitystore/groupMembership:GroupMembership", name, args == null ? GroupMembershipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupMembership(String name, Output<String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:identitystore/groupMembership:GroupMembership", name, state, makeResourceOptions(options, id));
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
    public static GroupMembership get(String name, Output<String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupMembership(name, id, state, options);
    }
}
