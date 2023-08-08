// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.secretsmanager;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.secretsmanager.SecretPolicyArgs;
import com.pulumi.aws.secretsmanager.inputs.SecretPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage AWS Secrets Manager secret policy.
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.secretsmanager.Secret;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.secretsmanager.SecretPolicy;
 * import com.pulumi.aws.secretsmanager.SecretPolicyArgs;
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
 *         var exampleSecret = new Secret(&#34;exampleSecret&#34;);
 * 
 *         final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .sid(&#34;EnableAnotherAWSAccountToReadTheSecret&#34;)
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;AWS&#34;)
 *                     .identifiers(&#34;arn:aws:iam::123456789012:root&#34;)
 *                     .build())
 *                 .actions(&#34;secretsmanager:GetSecretValue&#34;)
 *                 .resources(&#34;*&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleSecretPolicy = new SecretPolicy(&#34;exampleSecretPolicy&#34;, SecretPolicyArgs.builder()        
 *             .secretArn(exampleSecret.arn())
 *             .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
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
 *  to = aws_secretsmanager_secret_policy.example
 * 
 *  id = &#34;arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456&#34; } Using `pulumi import`, import `aws_secretsmanager_secret_policy` using the secret Amazon Resource Name (ARN). For exampleconsole % pulumi import aws_secretsmanager_secret_policy.example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
 * 
 */
@ResourceType(type="aws:secretsmanager/secretPolicy:SecretPolicy")
public class SecretPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
     * 
     */
    @Export(name="blockPublicPolicy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> blockPublicPolicy;

    /**
     * @return Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
     * 
     */
    public Output<Optional<Boolean>> blockPublicPolicy() {
        return Codegen.optional(this.blockPublicPolicy);
    }
    /**
     * Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Unlike `aws.secretsmanager.Secret`, where `policy` can be set to `&#34;{}&#34;` to delete the policy, `&#34;{}&#34;` is not a valid policy since `policy` is required.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Unlike `aws.secretsmanager.Secret`, where `policy` can be set to `&#34;{}&#34;` to delete the policy, `&#34;{}&#34;` is not a valid policy since `policy` is required.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * Secret ARN.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="secretArn", refs={String.class}, tree="[0]")
    private Output<String> secretArn;

    /**
     * @return Secret ARN.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> secretArn() {
        return this.secretArn;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretPolicy(String name) {
        this(name, SecretPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretPolicy(String name, SecretPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretPolicy(String name, SecretPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:secretsmanager/secretPolicy:SecretPolicy", name, args == null ? SecretPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretPolicy(String name, Output<String> id, @Nullable SecretPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:secretsmanager/secretPolicy:SecretPolicy", name, state, makeResourceOptions(options, id));
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
    public static SecretPolicy get(String name, Output<String> id, @Nullable SecretPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretPolicy(name, id, state, options);
    }
}
