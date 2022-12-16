// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.wafv2.RegexPatternSetArgs;
import com.pulumi.aws.wafv2.inputs.RegexPatternSetState;
import com.pulumi.aws.wafv2.outputs.RegexPatternSetRegularExpression;
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
 * Provides an AWS WAFv2 Regex Pattern Set Resource
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.wafv2.RegexPatternSet;
 * import com.pulumi.aws.wafv2.RegexPatternSetArgs;
 * import com.pulumi.aws.wafv2.inputs.RegexPatternSetRegularExpressionArgs;
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
 *         var example = new RegexPatternSet(&#34;example&#34;, RegexPatternSetArgs.builder()        
 *             .description(&#34;Example regex pattern set&#34;)
 *             .regularExpressions(            
 *                 RegexPatternSetRegularExpressionArgs.builder()
 *                     .regexString(&#34;one&#34;)
 *                     .build(),
 *                 RegexPatternSetRegularExpressionArgs.builder()
 *                     .regexString(&#34;two&#34;)
 *                     .build())
 *             .scope(&#34;REGIONAL&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Tag1&#34;, &#34;Value1&#34;),
 *                 Map.entry(&#34;Tag2&#34;, &#34;Value2&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * WAFv2 Regex Pattern Sets can be imported using `ID/name/scope` e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:wafv2/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
 * ```
 * 
 */
@ResourceType(type="aws:wafv2/regexPatternSet:RegexPatternSet")
public class RegexPatternSet extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) that identifies the cluster.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) that identifies the cluster.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * A friendly description of the regular expression pattern set.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A friendly description of the regular expression pattern set.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="lockToken", refs={String.class}, tree="[0]")
    private Output<String> lockToken;

    public Output<String> lockToken() {
        return this.lockToken;
    }
    /**
     * A friendly name of the regular expression pattern set.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A friendly name of the regular expression pattern set.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
     * 
     */
    @Export(name="regularExpressions", refs={List.class,RegexPatternSetRegularExpression.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RegexPatternSetRegularExpression>> regularExpressions;

    /**
     * @return One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
     * 
     */
    public Output<Optional<List<RegexPatternSetRegularExpression>>> regularExpressions() {
        return Codegen.optional(this.regularExpressions);
    }
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     * 
     */
    @Export(name="scope", refs={String.class}, tree="[0]")
    private Output<String> scope;

    /**
     * @return Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }
    /**
     * An array of key:value pairs to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return An array of key:value pairs to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegexPatternSet(String name) {
        this(name, RegexPatternSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegexPatternSet(String name, RegexPatternSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegexPatternSet(String name, RegexPatternSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:wafv2/regexPatternSet:RegexPatternSet", name, args == null ? RegexPatternSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RegexPatternSet(String name, Output<String> id, @Nullable RegexPatternSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:wafv2/regexPatternSet:RegexPatternSet", name, state, makeResourceOptions(options, id));
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
    public static RegexPatternSet get(String name, Output<String> id, @Nullable RegexPatternSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegexPatternSet(name, id, state, options);
    }
}
