// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apigateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.apigateway.BasePathMappingArgs;
import com.pulumi.aws.apigateway.inputs.BasePathMappingState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Connects a custom domain name registered via `aws.apigateway.DomainName`
 * with a deployed API so that its methods can be called via the
 * custom domain name.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.apigateway.Stage;
 * import com.pulumi.aws.apigateway.StageArgs;
 * import com.pulumi.aws.apigateway.DomainName;
 * import com.pulumi.aws.apigateway.DomainNameArgs;
 * import com.pulumi.aws.apigateway.BasePathMapping;
 * import com.pulumi.aws.apigateway.BasePathMappingArgs;
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
 *         var exampleStage = new Stage(&#34;exampleStage&#34;, StageArgs.builder()        
 *             .deployment(aws_api_gateway_deployment.example().id())
 *             .restApi(aws_api_gateway_rest_api.example().id())
 *             .stageName(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleDomainName = new DomainName(&#34;exampleDomainName&#34;, DomainNameArgs.builder()        
 *             .domainName(&#34;example.com&#34;)
 *             .certificateName(&#34;example-api&#34;)
 *             .certificateBody(Files.readString(Paths.get(String.format(&#34;%s/example.com/example.crt&#34;, path.module()))))
 *             .certificateChain(Files.readString(Paths.get(String.format(&#34;%s/example.com/ca.crt&#34;, path.module()))))
 *             .certificatePrivateKey(Files.readString(Paths.get(String.format(&#34;%s/example.com/example.key&#34;, path.module()))))
 *             .build());
 * 
 *         var exampleBasePathMapping = new BasePathMapping(&#34;exampleBasePathMapping&#34;, BasePathMappingArgs.builder()        
 *             .restApi(aws_api_gateway_rest_api.example().id())
 *             .stageName(exampleStage.stageName())
 *             .domainName(exampleDomainName.domainName())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * For an empty `base_path` or, in other words, a root path (`/`)terraform import {
 * 
 *  to = aws_api_gateway_base_path_mapping.example
 * 
 *  id = &#34;example.com/&#34; } For a non-root `base_path`terraform import {
 * 
 *  to = aws_api_gateway_base_path_mapping.example
 * 
 *  id = &#34;example.com/base-path&#34; } Using `pulumi import`, import `aws_api_gateway_base_path_mapping` using the domain name and base path. For exampleFor an empty `base_path` or, in other words, a root path (`/`)console % pulumi import aws_api_gateway_base_path_mapping.example example.com/ For a non-root `base_path`console % pulumi import aws_api_gateway_base_path_mapping.example example.com/base-path
 * 
 */
@ResourceType(type="aws:apigateway/basePathMapping:BasePathMapping")
public class BasePathMapping extends com.pulumi.resources.CustomResource {
    /**
     * Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     * 
     */
    @Export(name="basePath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> basePath;

    /**
     * @return Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
     * 
     */
    public Output<Optional<String>> basePath() {
        return Codegen.optional(this.basePath);
    }
    /**
     * Already-registered domain name to connect the API to.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return Already-registered domain name to connect the API to.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * ID of the API to connect.
     * 
     */
    @Export(name="restApi", refs={String.class}, tree="[0]")
    private Output<String> restApi;

    /**
     * @return ID of the API to connect.
     * 
     */
    public Output<String> restApi() {
        return this.restApi;
    }
    /**
     * Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     * 
     */
    @Export(name="stageName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stageName;

    /**
     * @return Name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
     * 
     */
    public Output<Optional<String>> stageName() {
        return Codegen.optional(this.stageName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BasePathMapping(String name) {
        this(name, BasePathMappingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BasePathMapping(String name, BasePathMappingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BasePathMapping(String name, BasePathMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/basePathMapping:BasePathMapping", name, args == null ? BasePathMappingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BasePathMapping(String name, Output<String> id, @Nullable BasePathMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/basePathMapping:BasePathMapping", name, state, makeResourceOptions(options, id));
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
    public static BasePathMapping get(String name, Output<String> id, @Nullable BasePathMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BasePathMapping(name, id, state, options);
    }
}
