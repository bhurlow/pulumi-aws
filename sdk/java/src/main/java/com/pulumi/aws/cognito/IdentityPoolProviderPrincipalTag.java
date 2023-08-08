// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cognito.IdentityPoolProviderPrincipalTagArgs;
import com.pulumi.aws.cognito.inputs.IdentityPoolProviderPrincipalTagState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an AWS Cognito Identity Principal Mapping.
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_cognito_identity_pool_provider_principal_tag.example
 * 
 *  id = &#34;us-west-2_abc123:CorpAD&#34; } Using `pulumi import`, import Cognito Identity Pool Roles Attachment using the Identity Pool ID and provider name. For exampleconsole % pulumi import aws_cognito_identity_pool_provider_principal_tag.example us-west-2_abc123:CorpAD
 * 
 */
@ResourceType(type="aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag")
public class IdentityPoolProviderPrincipalTag extends com.pulumi.resources.CustomResource {
    /**
     * An identity pool ID.
     * 
     */
    @Export(name="identityPoolId", refs={String.class}, tree="[0]")
    private Output<String> identityPoolId;

    /**
     * @return An identity pool ID.
     * 
     */
    public Output<String> identityPoolId() {
        return this.identityPoolId;
    }
    /**
     * The name of the identity provider.
     * 
     */
    @Export(name="identityProviderName", refs={String.class}, tree="[0]")
    private Output<String> identityProviderName;

    /**
     * @return The name of the identity provider.
     * 
     */
    public Output<String> identityProviderName() {
        return this.identityProviderName;
    }
    /**
     * String to string map of variables.
     * 
     */
    @Export(name="principalTags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> principalTags;

    /**
     * @return String to string map of variables.
     * 
     */
    public Output<Optional<Map<String,String>>> principalTags() {
        return Codegen.optional(this.principalTags);
    }
    /**
     * use default (username and clientID) attribute mappings.
     * 
     */
    @Export(name="useDefaults", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useDefaults;

    /**
     * @return use default (username and clientID) attribute mappings.
     * 
     */
    public Output<Optional<Boolean>> useDefaults() {
        return Codegen.optional(this.useDefaults);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IdentityPoolProviderPrincipalTag(String name) {
        this(name, IdentityPoolProviderPrincipalTagArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IdentityPoolProviderPrincipalTag(String name, IdentityPoolProviderPrincipalTagArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IdentityPoolProviderPrincipalTag(String name, IdentityPoolProviderPrincipalTagArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag", name, args == null ? IdentityPoolProviderPrincipalTagArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IdentityPoolProviderPrincipalTag(String name, Output<String> id, @Nullable IdentityPoolProviderPrincipalTagState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/identityPoolProviderPrincipalTag:IdentityPoolProviderPrincipalTag", name, state, makeResourceOptions(options, id));
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
    public static IdentityPoolProviderPrincipalTag get(String name, Output<String> id, @Nullable IdentityPoolProviderPrincipalTagState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IdentityPoolProviderPrincipalTag(name, id, state, options);
    }
}
