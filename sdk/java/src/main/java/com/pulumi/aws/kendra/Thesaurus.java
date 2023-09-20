// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kendra;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.kendra.ThesaurusArgs;
import com.pulumi.aws.kendra.inputs.ThesaurusState;
import com.pulumi.aws.kendra.outputs.ThesaurusSourceS3Path;
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
 * Resource for managing an AWS Kendra Thesaurus.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.kendra.Thesaurus;
 * import com.pulumi.aws.kendra.ThesaurusArgs;
 * import com.pulumi.aws.kendra.inputs.ThesaurusSourceS3PathArgs;
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
 *         var example = new Thesaurus(&#34;example&#34;, ThesaurusArgs.builder()        
 *             .indexId(aws_kendra_index.example().id())
 *             .roleArn(aws_iam_role.example().arn())
 *             .sourceS3Path(ThesaurusSourceS3PathArgs.builder()
 *                 .bucket(aws_s3_bucket.example().id())
 *                 .key(aws_s3_object.example().key())
 *                 .build())
 *             .tags(Map.of(&#34;Name&#34;, &#34;Example Kendra Thesaurus&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import `aws_kendra_thesaurus` using the unique identifiers of the thesaurus and index separated by a slash (`/`). For example:
 * 
 * ```sh
 *  $ pulumi import aws:kendra/thesaurus:Thesaurus example thesaurus-123456780/idx-8012925589
 * ```
 * 
 */
@ResourceType(type="aws:kendra/thesaurus:Thesaurus")
public class Thesaurus extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the thesaurus.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the thesaurus.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The description for a thesaurus.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description for a thesaurus.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The identifier of the index for a thesaurus.
     * 
     */
    @Export(name="indexId", refs={String.class}, tree="[0]")
    private Output<String> indexId;

    /**
     * @return The identifier of the index for a thesaurus.
     * 
     */
    public Output<String> indexId() {
        return this.indexId;
    }
    /**
     * The name for the thesaurus.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name for the thesaurus.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output<String> roleArn;

    /**
     * @return The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }
    /**
     * The S3 path where your thesaurus file sits in S3. Detailed below.
     * 
     * The `source_s3_path` configuration block supports the following arguments:
     * 
     */
    @Export(name="sourceS3Path", refs={ThesaurusSourceS3Path.class}, tree="[0]")
    private Output<ThesaurusSourceS3Path> sourceS3Path;

    /**
     * @return The S3 path where your thesaurus file sits in S3. Detailed below.
     * 
     * The `source_s3_path` configuration block supports the following arguments:
     * 
     */
    public Output<ThesaurusSourceS3Path> sourceS3Path() {
        return this.sourceS3Path;
    }
    /**
     * The current status of the thesaurus.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The current status of the thesaurus.
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
    @Export(name="thesaurusId", refs={String.class}, tree="[0]")
    private Output<String> thesaurusId;

    public Output<String> thesaurusId() {
        return this.thesaurusId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Thesaurus(String name) {
        this(name, ThesaurusArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Thesaurus(String name, ThesaurusArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Thesaurus(String name, ThesaurusArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kendra/thesaurus:Thesaurus", name, args == null ? ThesaurusArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Thesaurus(String name, Output<String> id, @Nullable ThesaurusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kendra/thesaurus:Thesaurus", name, state, makeResourceOptions(options, id));
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
    public static Thesaurus get(String name, Output<String> id, @Nullable ThesaurusState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Thesaurus(name, id, state, options);
    }
}
