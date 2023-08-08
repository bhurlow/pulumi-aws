// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.glue.ResourcePolicyArgs;
import com.pulumi.aws.glue.inputs.ResourcePolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Glue resource policy. Only one can exist per region.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetCallerIdentityArgs;
 * import com.pulumi.aws.inputs.GetPartitionArgs;
 * import com.pulumi.aws.inputs.GetRegionArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.glue.ResourcePolicy;
 * import com.pulumi.aws.glue.ResourcePolicyArgs;
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
 *         final var currentCallerIdentity = AwsFunctions.getCallerIdentity();
 * 
 *         final var currentPartition = AwsFunctions.getPartition();
 * 
 *         final var currentRegion = AwsFunctions.getRegion();
 * 
 *         final var glue-example-policy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .actions(&#34;glue:CreateTable&#34;)
 *                 .resources(String.format(&#34;arn:%s:glue:%s:%s:*&#34;, currentPartition.applyValue(getPartitionResult -&gt; getPartitionResult.partition()),currentRegion.applyValue(getRegionResult -&gt; getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId())))
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .identifiers(&#34;*&#34;)
 *                     .type(&#34;AWS&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var example = new ResourcePolicy(&#34;example&#34;, ResourcePolicyArgs.builder()        
 *             .policy(glue_example_policy.json())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_glue_resource_policy.Test
 * 
 *  id = &#34;12356789012&#34; } Using `pulumi import`, import Glue Resource Policy using the account ID. For exampleconsole % pulumi import aws_glue_resource_policy.Test 12356789012
 * 
 */
@ResourceType(type="aws:glue/resourcePolicy:ResourcePolicy")
public class ResourcePolicy extends com.pulumi.resources.CustomResource {
    /**
     * Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
     * 
     */
    @Export(name="enableHybrid", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enableHybrid;

    /**
     * @return Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
     * 
     */
    public Output<Optional<String>> enableHybrid() {
        return Codegen.optional(this.enableHybrid);
    }
    /**
     * The policy to be applied to the aws glue data catalog.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return The policy to be applied to the aws glue data catalog.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResourcePolicy(String name) {
        this(name, ResourcePolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourcePolicy(String name, ResourcePolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourcePolicy(String name, ResourcePolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:glue/resourcePolicy:ResourcePolicy", name, args == null ? ResourcePolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResourcePolicy(String name, Output<String> id, @Nullable ResourcePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:glue/resourcePolicy:ResourcePolicy", name, state, makeResourceOptions(options, id));
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
    public static ResourcePolicy get(String name, Output<String> id, @Nullable ResourcePolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourcePolicy(name, id, state, options);
    }
}
