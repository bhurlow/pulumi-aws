// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.quicksight.TemplateArgs;
import com.pulumi.aws.quicksight.inputs.TemplateState;
import com.pulumi.aws.quicksight.outputs.TemplatePermission;
import com.pulumi.aws.quicksight.outputs.TemplateSourceEntity;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing a QuickSight Template.
 * 
 * ## Example Usage
 * ### From Source Template
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.Template;
 * import com.pulumi.aws.quicksight.TemplateArgs;
 * import com.pulumi.aws.quicksight.inputs.TemplateSourceEntityArgs;
 * import com.pulumi.aws.quicksight.inputs.TemplateSourceEntitySourceTemplateArgs;
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
 *         var example = new Template(&#34;example&#34;, TemplateArgs.builder()        
 *             .templateId(&#34;example-id&#34;)
 *             .versionDescription(&#34;version&#34;)
 *             .sourceEntity(TemplateSourceEntityArgs.builder()
 *                 .sourceTemplate(TemplateSourceEntitySourceTemplateArgs.builder()
 *                     .arn(aws_quicksight_template.source().arn())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Definition
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.quicksight.Template;
 * import com.pulumi.aws.quicksight.TemplateArgs;
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
 *         var example = new Template(&#34;example&#34;, TemplateArgs.builder()        
 *             .definition(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .templateId(&#34;example-id&#34;)
 *             .versionDescription(&#34;version&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * A QuickSight Template can be imported using the AWS account ID and template ID separated by a comma (`,`) e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:quicksight/template:Template example 123456789012,example-id
 * ```
 * 
 */
@ResourceType(type="aws:quicksight/template:Template")
public class Template extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the resource.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the resource.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * AWS account ID.
     * 
     */
    @Export(name="awsAccountId", refs={String.class}, tree="[0]")
    private Output<String> awsAccountId;

    /**
     * @return AWS account ID.
     * 
     */
    public Output<String> awsAccountId() {
        return this.awsAccountId;
    }
    /**
     * The time that the template was created.
     * 
     */
    @Export(name="createdTime", refs={String.class}, tree="[0]")
    private Output<String> createdTime;

    /**
     * @return The time that the template was created.
     * 
     */
    public Output<String> createdTime() {
        return this.createdTime;
    }
    /**
     * The time that the template was last updated.
     * 
     */
    @Export(name="lastUpdatedTime", refs={String.class}, tree="[0]")
    private Output<String> lastUpdatedTime;

    /**
     * @return The time that the template was last updated.
     * 
     */
    public Output<String> lastUpdatedTime() {
        return this.lastUpdatedTime;
    }
    /**
     * Display name for the template.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name for the template.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A set of resource permissions on the template. Maximum of 64 items. See permissions.
     * 
     */
    @Export(name="permissions", refs={List.class,TemplatePermission.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TemplatePermission>> permissions;

    /**
     * @return A set of resource permissions on the template. Maximum of 64 items. See permissions.
     * 
     */
    public Output<Optional<List<TemplatePermission>>> permissions() {
        return Codegen.optional(this.permissions);
    }
    /**
     * The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
     * 
     */
    @Export(name="sourceEntity", refs={TemplateSourceEntity.class}, tree="[0]")
    private Output</* @Nullable */ TemplateSourceEntity> sourceEntity;

    /**
     * @return The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
     * 
     */
    public Output<Optional<TemplateSourceEntity>> sourceEntity() {
        return Codegen.optional(this.sourceEntity);
    }
    /**
     * Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
     * 
     */
    @Export(name="sourceEntityArn", refs={String.class}, tree="[0]")
    private Output<String> sourceEntityArn;

    /**
     * @return Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
     * 
     */
    public Output<String> sourceEntityArn() {
        return this.sourceEntityArn;
    }
    /**
     * The template creation status.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The template creation status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
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
     * Identifier for the template.
     * 
     */
    @Export(name="templateId", refs={String.class}, tree="[0]")
    private Output<String> templateId;

    /**
     * @return Identifier for the template.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }
    /**
     * A description of the current template version being created/updated.
     * 
     */
    @Export(name="versionDescription", refs={String.class}, tree="[0]")
    private Output<String> versionDescription;

    /**
     * @return A description of the current template version being created/updated.
     * 
     */
    public Output<String> versionDescription() {
        return this.versionDescription;
    }
    /**
     * The version number of the template version.
     * 
     */
    @Export(name="versionNumber", refs={Integer.class}, tree="[0]")
    private Output<Integer> versionNumber;

    /**
     * @return The version number of the template version.
     * 
     */
    public Output<Integer> versionNumber() {
        return this.versionNumber;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Template(String name) {
        this(name, TemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Template(String name, TemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Template(String name, TemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/template:Template", name, args == null ? TemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Template(String name, Output<String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:quicksight/template:Template", name, state, makeResourceOptions(options, id));
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
    public static Template get(String name, Output<String> id, @Nullable TemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Template(name, id, state, options);
    }
}
