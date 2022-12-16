// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.secretsmanager;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.secretsmanager.SecretArgs;
import com.pulumi.aws.secretsmanager.inputs.SecretState;
import com.pulumi.aws.secretsmanager.outputs.SecretReplica;
import com.pulumi.aws.secretsmanager.outputs.SecretRotationRules;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage AWS Secrets Manager secret metadata. To manage secret rotation, see the `aws.secretsmanager.SecretRotation` resource. To manage a secret value, see the `aws.secretsmanager.SecretVersion` resource.
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
 *         var example = new Secret(&#34;example&#34;);
 * 
 *     }
 * }
 * ```
 * ### Rotation Configuration
 * 
 * To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g., RDS) or deploying a custom Lambda function.
 * 
 * &gt; **NOTE:** Configuring rotation causes the secret to rotate once as soon as you store the secret. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.
 * 
 * &gt; **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version&#39;s VersionStage field.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.secretsmanager.Secret;
 * import com.pulumi.aws.secretsmanager.SecretArgs;
 * import com.pulumi.aws.secretsmanager.inputs.SecretRotationRulesArgs;
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
 *         var rotation_example = new Secret(&#34;rotation-example&#34;, SecretArgs.builder()        
 *             .rotationLambdaArn(aws_lambda_function.example().arn())
 *             .rotationRules(SecretRotationRulesArgs.builder()
 *                 .automaticallyAfterDays(7)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_secretsmanager_secret` can be imported by using the secret Amazon Resource Name (ARN), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:secretsmanager/secret:Secret example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
 * ```
 * 
 */
@ResourceType(type="aws:secretsmanager/secret:Secret")
public class Secret extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the secret.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the secret.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Description of the secret.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the secret.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Accepts boolean value to specify whether to overwrite a secret with the same name in the destination Region.
     * 
     */
    @Export(name="forceOverwriteReplicaSecret", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceOverwriteReplicaSecret;

    /**
     * @return Accepts boolean value to specify whether to overwrite a secret with the same name in the destination Region.
     * 
     */
    public Output<Optional<Boolean>> forceOverwriteReplicaSecret() {
        return Codegen.optional(this.forceOverwriteReplicaSecret);
    }
    /**
     * ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account&#39;s default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account&#39;s default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * Friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Output<String> namePrefix() {
        return this.namePrefix;
    }
    /**
     * Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = &#34;&#34;`) _will not_ delete the policy since it could have been set by `aws.secretsmanager.SecretPolicy`. To delete the `policy`, set it to `&#34;{}&#34;` (an empty JSON document).
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = &#34;&#34;`) _will not_ delete the policy since it could have been set by `aws.secretsmanager.SecretPolicy`. To delete the `policy`, set it to `&#34;{}&#34;` (an empty JSON document).
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * Number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     * 
     */
    @Export(name="recoveryWindowInDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> recoveryWindowInDays;

    /**
     * @return Number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
     * 
     */
    public Output<Optional<Integer>> recoveryWindowInDays() {
        return Codegen.optional(this.recoveryWindowInDays);
    }
    /**
     * Configuration block to support secret replication. See details below.
     * 
     */
    @Export(name="replicas", refs={List.class,SecretReplica.class}, tree="[0,1]")
    private Output<List<SecretReplica>> replicas;

    /**
     * @return Configuration block to support secret replication. See details below.
     * 
     */
    public Output<List<SecretReplica>> replicas() {
        return this.replicas;
    }
    /**
     * Whether automatic rotation is enabled for this secret.
     * 
     * @deprecated
     * Use the aws_secretsmanager_secret_rotation resource instead
     * 
     */
    @Deprecated /* Use the aws_secretsmanager_secret_rotation resource instead */
    @Export(name="rotationEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> rotationEnabled;

    /**
     * @return Whether automatic rotation is enabled for this secret.
     * 
     */
    public Output<Boolean> rotationEnabled() {
        return this.rotationEnabled;
    }
    /**
     * ARN of the Lambda function that can rotate the secret. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     * 
     * @deprecated
     * Use the aws_secretsmanager_secret_rotation resource instead
     * 
     */
    @Deprecated /* Use the aws_secretsmanager_secret_rotation resource instead */
    @Export(name="rotationLambdaArn", refs={String.class}, tree="[0]")
    private Output<String> rotationLambdaArn;

    /**
     * @return ARN of the Lambda function that can rotate the secret. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     * 
     */
    public Output<String> rotationLambdaArn() {
        return this.rotationLambdaArn;
    }
    /**
     * Configuration block for the rotation configuration of this secret. Defined below. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     * 
     * @deprecated
     * Use the aws_secretsmanager_secret_rotation resource instead
     * 
     */
    @Deprecated /* Use the aws_secretsmanager_secret_rotation resource instead */
    @Export(name="rotationRules", refs={SecretRotationRules.class}, tree="[0]")
    private Output<SecretRotationRules> rotationRules;

    /**
     * @return Configuration block for the rotation configuration of this secret. Defined below. Use the `aws.secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
     * 
     */
    public Output<SecretRotationRules> rotationRules() {
        return this.rotationRules;
    }
    /**
     * Key-value map of user-defined tags that are attached to the secret. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of user-defined tags that are attached to the secret. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Secret(String name) {
        this(name, SecretArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Secret(String name, @Nullable SecretArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Secret(String name, @Nullable SecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:secretsmanager/secret:Secret", name, args == null ? SecretArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Secret(String name, Output<String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:secretsmanager/secret:Secret", name, state, makeResourceOptions(options, id));
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
    public static Secret get(String name, Output<String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Secret(name, id, state, options);
    }
}
