// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudfront;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cloudfront.FunctionArgs;
import com.pulumi.aws.cloudfront.inputs.FunctionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CloudFront Function resource. With CloudFront Functions in Amazon CloudFront, you can write lightweight functions in JavaScript for high-scale, latency-sensitive CDN customizations.
 * 
 * See [CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-functions.html)
 * 
 * &gt; **NOTE:** You cannot delete a function if it’s associated with a cache behavior. First, update your distributions to remove the function association from all cache behaviors, then delete the function.
 * 
 * ## Example Usage
 * ### Basic Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudfront.Function;
 * import com.pulumi.aws.cloudfront.FunctionArgs;
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
 *         var test = new Function(&#34;test&#34;, FunctionArgs.builder()        
 *             .runtime(&#34;cloudfront-js-2.0&#34;)
 *             .comment(&#34;my function&#34;)
 *             .publish(true)
 *             .code(Files.readString(Paths.get(String.format(&#34;%s/function.js&#34;, path.module()))))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import CloudFront Functions using the `name`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:cloudfront/function:Function test my_test_function
 * ```
 * 
 */
@ResourceType(type="aws:cloudfront/function:Function")
public class Function extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) identifying your CloudFront Function.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) identifying your CloudFront Function.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Source code of the function
     * 
     */
    @Export(name="code", refs={String.class}, tree="[0]")
    private Output<String> code;

    /**
     * @return Source code of the function
     * 
     */
    public Output<String> code() {
        return this.code;
    }
    /**
     * Comment.
     * 
     */
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> comment;

    /**
     * @return Comment.
     * 
     */
    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    /**
     * ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * ETag hash of any `LIVE` stage of the function.
     * 
     */
    @Export(name="liveStageEtag", refs={String.class}, tree="[0]")
    private Output<String> liveStageEtag;

    /**
     * @return ETag hash of any `LIVE` stage of the function.
     * 
     */
    public Output<String> liveStageEtag() {
        return this.liveStageEtag;
    }
    /**
     * Unique name for your CloudFront Function.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name for your CloudFront Function.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
     * 
     */
    @Export(name="publish", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> publish;

    /**
     * @return Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> publish() {
        return Codegen.optional(this.publish);
    }
    /**
     * Identifier of the function&#39;s runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="runtime", refs={String.class}, tree="[0]")
    private Output<String> runtime;

    /**
     * @return Identifier of the function&#39;s runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> runtime() {
        return this.runtime;
    }
    /**
     * Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Function(String name) {
        this(name, FunctionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Function(String name, FunctionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Function(String name, FunctionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudfront/function:Function", name, args == null ? FunctionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Function(String name, Output<String> id, @Nullable FunctionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cloudfront/function:Function", name, state, makeResourceOptions(options, id));
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
    public static Function get(String name, Output<String> id, @Nullable FunctionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Function(name, id, state, options);
    }
}
