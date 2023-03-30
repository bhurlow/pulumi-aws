// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cognito.IdentityPoolRoleAttachmentArgs;
import com.pulumi.aws.cognito.inputs.IdentityPoolRoleAttachmentState;
import com.pulumi.aws.cognito.outputs.IdentityPoolRoleAttachmentRoleMapping;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an AWS Cognito Identity Pool Roles Attachment.
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cognito.IdentityPool;
 * import com.pulumi.aws.cognito.IdentityPoolArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.iam.RolePolicy;
 * import com.pulumi.aws.iam.RolePolicyArgs;
 * import com.pulumi.aws.cognito.IdentityPoolRoleAttachment;
 * import com.pulumi.aws.cognito.IdentityPoolRoleAttachmentArgs;
 * import com.pulumi.aws.cognito.inputs.IdentityPoolRoleAttachmentRoleMappingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var mainIdentityPool = new IdentityPool(&#34;mainIdentityPool&#34;, IdentityPoolArgs.builder()        
 *             .identityPoolName(&#34;identity pool&#34;)
 *             .allowUnauthenticatedIdentities(false)
 *             .supportedLoginProviders(Map.of(&#34;graph.facebook.com&#34;, &#34;7346241598935555&#34;))
 *             .build());
 * 
 *         final var authenticatedPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Federated&#34;)
 *                     .identifiers(&#34;cognito-identity.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRoleWithWebIdentity&#34;)
 *                 .conditions(                
 *                     GetPolicyDocumentStatementConditionArgs.builder()
 *                         .test(&#34;StringEquals&#34;)
 *                         .variable(&#34;cognito-identity.amazonaws.com:aud&#34;)
 *                         .values(mainIdentityPool.id())
 *                         .build(),
 *                     GetPolicyDocumentStatementConditionArgs.builder()
 *                         .test(&#34;ForAnyValue:StringLike&#34;)
 *                         .variable(&#34;cognito-identity.amazonaws.com:amr&#34;)
 *                         .values(&#34;authenticated&#34;)
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *         var authenticatedRole = new Role(&#34;authenticatedRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(authenticatedPolicyDocument.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult).applyValue(authenticatedPolicyDocument -&gt; authenticatedPolicyDocument.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json())))
 *             .build());
 * 
 *         final var authenticatedRolePolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .effect(&#34;Allow&#34;)
 *             .actions(            
 *                 &#34;mobileanalytics:PutEvents&#34;,
 *                 &#34;cognito-sync:*&#34;,
 *                 &#34;cognito-identity:*&#34;)
 *             .resources(&#34;*&#34;)
 *             .build());
 * 
 *         var authenticatedRolePolicy = new RolePolicy(&#34;authenticatedRolePolicy&#34;, RolePolicyArgs.builder()        
 *             .role(authenticatedRole.id())
 *             .policy(authenticatedRolePolicyPolicyDocument.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var mainIdentityPoolRoleAttachment = new IdentityPoolRoleAttachment(&#34;mainIdentityPoolRoleAttachment&#34;, IdentityPoolRoleAttachmentArgs.builder()        
 *             .identityPoolId(mainIdentityPool.id())
 *             .roleMappings(IdentityPoolRoleAttachmentRoleMappingArgs.builder()
 *                 .identityProvider(&#34;graph.facebook.com&#34;)
 *                 .ambiguousRoleResolution(&#34;AuthenticatedRole&#34;)
 *                 .type(&#34;Rules&#34;)
 *                 .mappingRules(IdentityPoolRoleAttachmentRoleMappingMappingRuleArgs.builder()
 *                     .claim(&#34;isAdmin&#34;)
 *                     .matchType(&#34;Equals&#34;)
 *                     .roleArn(authenticatedRole.arn())
 *                     .value(&#34;paid&#34;)
 *                     .build())
 *                 .build())
 *             .roles(Map.of(&#34;authenticated&#34;, authenticatedRole.arn()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cognito Identity Pool Roles Attachment can be imported using the Identity Pool ID, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment example us-west-2:b64805ad-cb56-40ba-9ffc-f5d8207e6d42
 * ```
 * 
 */
@ResourceType(type="aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment")
public class IdentityPoolRoleAttachment extends com.pulumi.resources.CustomResource {
    /**
     * An identity pool ID in the format `REGION_GUID`.
     * 
     */
    @Export(name="identityPoolId", refs={String.class}, tree="[0]")
    private Output<String> identityPoolId;

    /**
     * @return An identity pool ID in the format `REGION_GUID`.
     * 
     */
    public Output<String> identityPoolId() {
        return this.identityPoolId;
    }
    /**
     * A List of Role Mapping.
     * 
     */
    @Export(name="roleMappings", refs={List.class,IdentityPoolRoleAttachmentRoleMapping.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IdentityPoolRoleAttachmentRoleMapping>> roleMappings;

    /**
     * @return A List of Role Mapping.
     * 
     */
    public Output<Optional<List<IdentityPoolRoleAttachmentRoleMapping>>> roleMappings() {
        return Codegen.optional(this.roleMappings);
    }
    /**
     * The map of roles associated with this pool. For a given role, the key will be either &#34;authenticated&#34; or &#34;unauthenticated&#34; and the value will be the Role ARN.
     * 
     */
    @Export(name="roles", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> roles;

    /**
     * @return The map of roles associated with this pool. For a given role, the key will be either &#34;authenticated&#34; or &#34;unauthenticated&#34; and the value will be the Role ARN.
     * 
     */
    public Output<Map<String,String>> roles() {
        return this.roles;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IdentityPoolRoleAttachment(String name) {
        this(name, IdentityPoolRoleAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IdentityPoolRoleAttachment(String name, IdentityPoolRoleAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IdentityPoolRoleAttachment(String name, IdentityPoolRoleAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, args == null ? IdentityPoolRoleAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IdentityPoolRoleAttachment(String name, Output<String> id, @Nullable IdentityPoolRoleAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, state, makeResourceOptions(options, id));
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
    public static IdentityPoolRoleAttachment get(String name, Output<String> id, @Nullable IdentityPoolRoleAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IdentityPoolRoleAttachment(name, id, state, options);
    }
}
