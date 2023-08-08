// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.athena;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.athena.WorkgroupArgs;
import com.pulumi.aws.athena.inputs.WorkgroupState;
import com.pulumi.aws.athena.outputs.WorkgroupConfiguration;
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
 * Provides an Athena Workgroup.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.athena.Workgroup;
 * import com.pulumi.aws.athena.WorkgroupArgs;
 * import com.pulumi.aws.athena.inputs.WorkgroupConfigurationArgs;
 * import com.pulumi.aws.athena.inputs.WorkgroupConfigurationResultConfigurationArgs;
 * import com.pulumi.aws.athena.inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs;
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
 *         var example = new Workgroup(&#34;example&#34;, WorkgroupArgs.builder()        
 *             .configuration(WorkgroupConfigurationArgs.builder()
 *                 .enforceWorkgroupConfiguration(true)
 *                 .publishCloudwatchMetricsEnabled(true)
 *                 .resultConfiguration(WorkgroupConfigurationResultConfigurationArgs.builder()
 *                     .outputLocation(String.format(&#34;s3://%s/output/&#34;, aws_s3_bucket.example().bucket()))
 *                     .encryptionConfiguration(WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs.builder()
 *                         .encryptionOption(&#34;SSE_KMS&#34;)
 *                         .kmsKeyArn(aws_kms_key.example().arn())
 *                         .build())
 *                     .build())
 *                 .build())
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
 *  to = aws_athena_workgroup.example
 * 
 *  id = &#34;example&#34; } Using `pulumi import`, import Athena Workgroups using their name. For exampleconsole % pulumi import aws_athena_workgroup.example example
 * 
 */
@ResourceType(type="aws:athena/workgroup:Workgroup")
public class Workgroup extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the workgroup
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the workgroup
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Configuration block with various settings for the workgroup. Documented below.
     * 
     */
    @Export(name="configuration", refs={WorkgroupConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ WorkgroupConfiguration> configuration;

    /**
     * @return Configuration block with various settings for the workgroup. Documented below.
     * 
     */
    public Output<Optional<WorkgroupConfiguration>> configuration() {
        return Codegen.optional(this.configuration);
    }
    /**
     * Description of the workgroup.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the workgroup.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Option to delete the workgroup and its contents even if the workgroup contains any named queries.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return Option to delete the workgroup and its contents even if the workgroup contains any named queries.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * Name of the workgroup.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the workgroup.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * Key-value map of resource tags for the workgroup. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags for the workgroup. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public Workgroup(String name) {
        this(name, WorkgroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Workgroup(String name, @Nullable WorkgroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Workgroup(String name, @Nullable WorkgroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:athena/workgroup:Workgroup", name, args == null ? WorkgroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Workgroup(String name, Output<String> id, @Nullable WorkgroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:athena/workgroup:Workgroup", name, state, makeResourceOptions(options, id));
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
    public static Workgroup get(String name, Output<String> id, @Nullable WorkgroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Workgroup(name, id, state, options);
    }
}
