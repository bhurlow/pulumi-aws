// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kendra;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.kendra.QuerySuggestionsBlockListArgs;
import com.pulumi.aws.kendra.inputs.QuerySuggestionsBlockListState;
import com.pulumi.aws.kendra.outputs.QuerySuggestionsBlockListSourceS3Path;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS Kendra block list used for query suggestions for an index.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.kendra.QuerySuggestionsBlockList;
 * import com.pulumi.aws.kendra.QuerySuggestionsBlockListArgs;
 * import com.pulumi.aws.kendra.inputs.QuerySuggestionsBlockListSourceS3PathArgs;
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
 *         var example = new QuerySuggestionsBlockList(&#34;example&#34;, QuerySuggestionsBlockListArgs.builder()        
 *             .indexId(aws_kendra_index.example().id())
 *             .roleArn(aws_iam_role.example().arn())
 *             .sourceS3Path(QuerySuggestionsBlockListSourceS3PathArgs.builder()
 *                 .bucket(aws_s3_bucket.example().id())
 *                 .key(&#34;example/suggestions.txt&#34;)
 *                 .build())
 *             .tags(Map.of(&#34;Name&#34;, &#34;Example Kendra Index&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_kendra_query_suggestions_block_list` can be imported using the unique identifiers of the block list and index separated by a slash (`/`), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:kendra/querySuggestionsBlockList:QuerySuggestionsBlockList example blocklist-123456780/idx-8012925589
 * ```
 * 
 */
@ResourceType(type="aws:kendra/querySuggestionsBlockList:QuerySuggestionsBlockList")
public class QuerySuggestionsBlockList extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the block list.
     * 
     */
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    /**
     * @return ARN of the block list.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The description for a block list.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description for a block list.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The identifier of the index for a block list.
     * 
     */
    @Export(name="indexId", type=String.class, parameters={})
    private Output<String> indexId;

    /**
     * @return The identifier of the index for a block list.
     * 
     */
    public Output<String> indexId() {
        return this.indexId;
    }
    /**
     * The name for the block list.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name for the block list.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The unique indentifier of the block list.
     * 
     */
    @Export(name="querySuggestionsBlockListId", type=String.class, parameters={})
    private Output<String> querySuggestionsBlockListId;

    /**
     * @return The unique indentifier of the block list.
     * 
     */
    public Output<String> querySuggestionsBlockListId() {
        return this.querySuggestionsBlockListId;
    }
    /**
     * The IAM (Identity and Access Management) role used to access the block list text file in S3.
     * 
     */
    @Export(name="roleArn", type=String.class, parameters={})
    private Output<String> roleArn;

    /**
     * @return The IAM (Identity and Access Management) role used to access the block list text file in S3.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }
    /**
     * The S3 path where your block list text file sits in S3. Detailed below.
     * 
     */
    @Export(name="sourceS3Path", type=QuerySuggestionsBlockListSourceS3Path.class, parameters={})
    private Output<QuerySuggestionsBlockListSourceS3Path> sourceS3Path;

    /**
     * @return The S3 path where your block list text file sits in S3. Detailed below.
     * 
     */
    public Output<QuerySuggestionsBlockListSourceS3Path> sourceS3Path() {
        return this.sourceS3Path;
    }
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
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
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
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
    public QuerySuggestionsBlockList(String name) {
        this(name, QuerySuggestionsBlockListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QuerySuggestionsBlockList(String name, QuerySuggestionsBlockListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QuerySuggestionsBlockList(String name, QuerySuggestionsBlockListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kendra/querySuggestionsBlockList:QuerySuggestionsBlockList", name, args == null ? QuerySuggestionsBlockListArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QuerySuggestionsBlockList(String name, Output<String> id, @Nullable QuerySuggestionsBlockListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:kendra/querySuggestionsBlockList:QuerySuggestionsBlockList", name, state, makeResourceOptions(options, id));
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
    public static QuerySuggestionsBlockList get(String name, Output<String> id, @Nullable QuerySuggestionsBlockListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QuerySuggestionsBlockList(name, id, state, options);
    }
}
