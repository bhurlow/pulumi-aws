// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.serverlessrepository;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.serverlessrepository.CloudFormationStackArgs;
import com.pulumi.aws.serverlessrepository.inputs.CloudFormationStackState;
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
 * Deploys an Application CloudFormation Stack from the Serverless Application Repository.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetPartitionArgs;
 * import com.pulumi.aws.inputs.GetRegionArgs;
 * import com.pulumi.aws.serverlessrepository.CloudFormationStack;
 * import com.pulumi.aws.serverlessrepository.CloudFormationStackArgs;
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
 *         final var currentPartition = AwsFunctions.getPartition();
 * 
 *         final var currentRegion = AwsFunctions.getRegion();
 * 
 *         var postgres_rotator = new CloudFormationStack(&#34;postgres-rotator&#34;, CloudFormationStackArgs.builder()        
 *             .applicationId(&#34;arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser&#34;)
 *             .capabilities(            
 *                 &#34;CAPABILITY_IAM&#34;,
 *                 &#34;CAPABILITY_RESOURCE_POLICY&#34;)
 *             .parameters(Map.ofEntries(
 *                 Map.entry(&#34;endpoint&#34;, String.format(&#34;secretsmanager.%s.%s&#34;, currentRegion.applyValue(getRegionResult -&gt; getRegionResult.name()),currentPartition.applyValue(getPartitionResult -&gt; getPartitionResult.dnsSuffix()))),
 *                 Map.entry(&#34;functionName&#34;, &#34;func-postgres-rotator&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Serverless Application Repository Stack using the CloudFormation Stack name (with or without the `serverlessrepo-` prefix) or the CloudFormation Stack ID. For example:
 * 
 * ```sh
 *  $ pulumi import aws:serverlessrepository/cloudFormationStack:CloudFormationStack example serverlessrepo-postgres-rotator
 * ```
 * 
 */
@ResourceType(type="aws:serverlessrepository/cloudFormationStack:CloudFormationStack")
public class CloudFormationStack extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the application from the Serverless Application Repository.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return The ARN of the application from the Serverless Application Repository.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
     * 
     */
    @Export(name="capabilities", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> capabilities;

    /**
     * @return A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
     * 
     */
    public Output<List<String>> capabilities() {
        return this.capabilities;
    }
    /**
     * The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A map of outputs from the stack.
     * 
     */
    @Export(name="outputs", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> outputs;

    /**
     * @return A map of outputs from the stack.
     * 
     */
    public Output<Map<String,String>> outputs() {
        return this.outputs;
    }
    /**
     * A map of Parameter structures that specify input parameters for the stack.
     * 
     */
    @Export(name="parameters", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> parameters;

    /**
     * @return A map of Parameter structures that specify input parameters for the stack.
     * 
     */
    public Output<Map<String,String>> parameters() {
        return this.parameters;
    }
    /**
     * The version of the application to deploy. If not supplied, deploys the latest version.
     * 
     */
    @Export(name="semanticVersion", refs={String.class}, tree="[0]")
    private Output<String> semanticVersion;

    /**
     * @return The version of the application to deploy. If not supplied, deploys the latest version.
     * 
     */
    public Output<String> semanticVersion() {
        return this.semanticVersion;
    }
    /**
     * A list of tags to associate with this stack. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A list of tags to associate with this stack. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudFormationStack(String name) {
        this(name, CloudFormationStackArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudFormationStack(String name, CloudFormationStackArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudFormationStack(String name, CloudFormationStackArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:serverlessrepository/cloudFormationStack:CloudFormationStack", name, args == null ? CloudFormationStackArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CloudFormationStack(String name, Output<String> id, @Nullable CloudFormationStackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:serverlessrepository/cloudFormationStack:CloudFormationStack", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static CloudFormationStack get(String name, Output<String> id, @Nullable CloudFormationStackState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudFormationStack(name, id, state, options);
    }
}
