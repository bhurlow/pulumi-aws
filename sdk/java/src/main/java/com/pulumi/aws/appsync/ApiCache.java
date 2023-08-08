// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appsync;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.appsync.ApiCacheArgs;
import com.pulumi.aws.appsync.inputs.ApiCacheState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an AppSync API Cache.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.appsync.GraphQLApi;
 * import com.pulumi.aws.appsync.GraphQLApiArgs;
 * import com.pulumi.aws.appsync.ApiCache;
 * import com.pulumi.aws.appsync.ApiCacheArgs;
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
 *         var exampleGraphQLApi = new GraphQLApi(&#34;exampleGraphQLApi&#34;, GraphQLApiArgs.builder()        
 *             .authenticationType(&#34;API_KEY&#34;)
 *             .build());
 * 
 *         var exampleApiCache = new ApiCache(&#34;exampleApiCache&#34;, ApiCacheArgs.builder()        
 *             .apiId(exampleGraphQLApi.id())
 *             .apiCachingBehavior(&#34;FULL_REQUEST_CACHING&#34;)
 *             .type(&#34;LARGE&#34;)
 *             .ttl(900)
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
 *  to = aws_appsync_api_cache.example
 * 
 *  id = &#34;xxxxx&#34; } Using `pulumi import`, import `aws_appsync_api_cache` using the AppSync API ID. For exampleconsole % pulumi import aws_appsync_api_cache.example xxxxx
 * 
 */
@ResourceType(type="aws:appsync/apiCache:ApiCache")
public class ApiCache extends com.pulumi.resources.CustomResource {
    /**
     * Caching behavior. Valid values are `FULL_REQUEST_CACHING` and `PER_RESOLVER_CACHING`.
     * 
     */
    @Export(name="apiCachingBehavior", refs={String.class}, tree="[0]")
    private Output<String> apiCachingBehavior;

    /**
     * @return Caching behavior. Valid values are `FULL_REQUEST_CACHING` and `PER_RESOLVER_CACHING`.
     * 
     */
    public Output<String> apiCachingBehavior() {
        return this.apiCachingBehavior;
    }
    /**
     * GraphQL API ID.
     * 
     */
    @Export(name="apiId", refs={String.class}, tree="[0]")
    private Output<String> apiId;

    /**
     * @return GraphQL API ID.
     * 
     */
    public Output<String> apiId() {
        return this.apiId;
    }
    /**
     * At-rest encryption flag for cache. You cannot update this setting after creation.
     * 
     */
    @Export(name="atRestEncryptionEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> atRestEncryptionEnabled;

    /**
     * @return At-rest encryption flag for cache. You cannot update this setting after creation.
     * 
     */
    public Output<Optional<Boolean>> atRestEncryptionEnabled() {
        return Codegen.optional(this.atRestEncryptionEnabled);
    }
    /**
     * Transit encryption flag when connecting to cache. You cannot update this setting after creation.
     * 
     */
    @Export(name="transitEncryptionEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> transitEncryptionEnabled;

    /**
     * @return Transit encryption flag when connecting to cache. You cannot update this setting after creation.
     * 
     */
    public Output<Optional<Boolean>> transitEncryptionEnabled() {
        return Codegen.optional(this.transitEncryptionEnabled);
    }
    /**
     * TTL in seconds for cache entries.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output<Integer> ttl;

    /**
     * @return TTL in seconds for cache entries.
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }
    /**
     * Cache instance type. Valid values are `SMALL`, `MEDIUM`, `LARGE`, `XLARGE`, `LARGE_2X`, `LARGE_4X`, `LARGE_8X`, `LARGE_12X`, `T2_SMALL`, `T2_MEDIUM`, `R4_LARGE`, `R4_XLARGE`, `R4_2XLARGE`, `R4_4XLARGE`, `R4_8XLARGE`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Cache instance type. Valid values are `SMALL`, `MEDIUM`, `LARGE`, `XLARGE`, `LARGE_2X`, `LARGE_4X`, `LARGE_8X`, `LARGE_12X`, `T2_SMALL`, `T2_MEDIUM`, `R4_LARGE`, `R4_XLARGE`, `R4_2XLARGE`, `R4_4XLARGE`, `R4_8XLARGE`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApiCache(String name) {
        this(name, ApiCacheArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApiCache(String name, ApiCacheArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApiCache(String name, ApiCacheArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appsync/apiCache:ApiCache", name, args == null ? ApiCacheArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApiCache(String name, Output<String> id, @Nullable ApiCacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appsync/apiCache:ApiCache", name, state, makeResourceOptions(options, id));
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
    public static ApiCache get(String name, Output<String> id, @Nullable ApiCacheState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApiCache(name, id, state, options);
    }
}
