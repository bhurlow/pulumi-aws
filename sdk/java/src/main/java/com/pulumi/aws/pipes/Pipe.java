// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.pipes.PipeArgs;
import com.pulumi.aws.pipes.inputs.PipeState;
import com.pulumi.aws.pipes.outputs.PipeSourceParameters;
import com.pulumi.aws.pipes.outputs.PipeTargetParameters;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS EventBridge Pipes Pipe.
 * 
 * You can find out more about EventBridge Pipes in the [User Guide](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html).
 * 
 * &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetCallerIdentityArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.sqs.Queue;
 * import com.pulumi.aws.iam.RolePolicy;
 * import com.pulumi.aws.iam.RolePolicyArgs;
 * import com.pulumi.aws.pipes.Pipe;
 * import com.pulumi.aws.pipes.PipeArgs;
 * import com.pulumi.aws.pipes.inputs.PipeSourceParametersArgs;
 * import com.pulumi.aws.pipes.inputs.PipeTargetParametersArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         final var main = AwsFunctions.getCallerIdentity();
 * 
 *         var test = new Role(&#34;test&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Version&#34;, &#34;2012-10-17&#34;),
 *                     jsonProperty(&#34;Statement&#34;, jsonObject(
 *                         jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                         jsonProperty(&#34;Action&#34;, &#34;sts:AssumeRole&#34;),
 *                         jsonProperty(&#34;Principal&#34;, jsonObject(
 *                             jsonProperty(&#34;Service&#34;, &#34;pipes.amazonaws.com&#34;)
 *                         )),
 *                         jsonProperty(&#34;Condition&#34;, jsonObject(
 *                             jsonProperty(&#34;StringEquals&#34;, jsonObject(
 *                                 jsonProperty(&#34;aws:SourceAccount&#34;, main.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.accountId()))
 *                             ))
 *                         ))
 *                     ))
 *                 )))
 *             .build());
 * 
 *         var sourceQueue = new Queue(&#34;sourceQueue&#34;);
 * 
 *         var sourceRolePolicy = new RolePolicy(&#34;sourceRolePolicy&#34;, RolePolicyArgs.builder()        
 *             .role(test.id())
 *             .policy(sourceQueue.arn().applyValue(arn -&gt; serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Version&#34;, &#34;2012-10-17&#34;),
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                         jsonProperty(&#34;Action&#34;, jsonArray(
 *                             &#34;sqs:DeleteMessage&#34;, 
 *                             &#34;sqs:GetQueueAttributes&#34;, 
 *                             &#34;sqs:ReceiveMessage&#34;
 *                         )),
 *                         jsonProperty(&#34;Resource&#34;, jsonArray(arn))
 *                     )))
 *                 ))))
 *             .build());
 * 
 *         var targetQueue = new Queue(&#34;targetQueue&#34;);
 * 
 *         var targetRolePolicy = new RolePolicy(&#34;targetRolePolicy&#34;, RolePolicyArgs.builder()        
 *             .role(test.id())
 *             .policy(targetQueue.arn().applyValue(arn -&gt; serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Version&#34;, &#34;2012-10-17&#34;),
 *                     jsonProperty(&#34;Statement&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;Effect&#34;, &#34;Allow&#34;),
 *                         jsonProperty(&#34;Action&#34;, jsonArray(&#34;sqs:SendMessage&#34;)),
 *                         jsonProperty(&#34;Resource&#34;, jsonArray(arn))
 *                     )))
 *                 ))))
 *             .build());
 * 
 *         var example = new Pipe(&#34;example&#34;, PipeArgs.builder()        
 *             .roleArn(aws_iam_role.example().arn())
 *             .source(sourceQueue.arn())
 *             .target(targetQueue.arn())
 *             .sourceParameters()
 *             .targetParameters()
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     sourceRolePolicy,
 *                     targetRolePolicy)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Pipes can be imported using the `name`. For example
 * 
 * ```sh
 *  $ pulumi import aws:pipes/pipe:Pipe example my-pipe
 * ```
 * 
 */
@ResourceType(type="aws:pipes/pipe:Pipe")
public class Pipe extends com.pulumi.resources.CustomResource {
    /**
     * ARN of this pipe.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of this pipe.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * A description of the pipe. At most 512 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the pipe. At most 512 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The state the pipe should be in. One of: `RUNNING`, `STOPPED`.
     * 
     */
    @Export(name="desiredState", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> desiredState;

    /**
     * @return The state the pipe should be in. One of: `RUNNING`, `STOPPED`.
     * 
     */
    public Output<Optional<String>> desiredState() {
        return Codegen.optional(this.desiredState);
    }
    /**
     * Enrichment resource of the pipe (typically an ARN). Read more about enrichment in the [User Guide](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html#pipes-enrichment).
     * 
     */
    @Export(name="enrichment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enrichment;

    /**
     * @return Enrichment resource of the pipe (typically an ARN). Read more about enrichment in the [User Guide](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html#pipes-enrichment).
     * 
     */
    public Output<Optional<String>> enrichment() {
        return Codegen.optional(this.enrichment);
    }
    /**
     * Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
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
     * ARN of the role that allows the pipe to send data to the target.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output<String> roleArn;

    /**
     * @return ARN of the role that allows the pipe to send data to the target.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }
    /**
     * Source resource of the pipe (typically an ARN).
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return Source resource of the pipe (typically an ARN).
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * Parameters required to set up a source for the pipe. Detailed below.
     * 
     */
    @Export(name="sourceParameters", refs={PipeSourceParameters.class}, tree="[0]")
    private Output<PipeSourceParameters> sourceParameters;

    /**
     * @return Parameters required to set up a source for the pipe. Detailed below.
     * 
     */
    public Output<PipeSourceParameters> sourceParameters() {
        return this.sourceParameters;
    }
    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Target resource of the pipe (typically an ARN).
     * 
     */
    @Export(name="target", refs={String.class}, tree="[0]")
    private Output<String> target;

    /**
     * @return Target resource of the pipe (typically an ARN).
     * 
     */
    public Output<String> target() {
        return this.target;
    }
    /**
     * Parameters required to set up a target for your pipe. Detailed below.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="targetParameters", refs={PipeTargetParameters.class}, tree="[0]")
    private Output<PipeTargetParameters> targetParameters;

    /**
     * @return Parameters required to set up a target for your pipe. Detailed below.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<PipeTargetParameters> targetParameters() {
        return this.targetParameters;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Pipe(String name) {
        this(name, PipeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Pipe(String name, PipeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Pipe(String name, PipeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:pipes/pipe:Pipe", name, args == null ? PipeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Pipe(String name, Output<String> id, @Nullable PipeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:pipes/pipe:Pipe", name, state, makeResourceOptions(options, id));
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
    public static Pipe get(String name, Output<String> id, @Nullable PipeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Pipe(name, id, state, options);
    }
}
