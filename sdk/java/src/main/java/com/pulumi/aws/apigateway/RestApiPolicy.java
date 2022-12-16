// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apigateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.apigateway.RestApiPolicyArgs;
import com.pulumi.aws.apigateway.inputs.RestApiPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an API Gateway REST API Policy.
 * 
 * &gt; **Note:** Amazon API Gateway Version 1 resources are used for creating and deploying REST APIs. To create and deploy WebSocket and HTTP APIs, use Amazon API Gateway Version 2 resources.
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.apigateway.RestApi;
 * import com.pulumi.aws.apigateway.RestApiPolicy;
 * import com.pulumi.aws.apigateway.RestApiPolicyArgs;
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
 *         var testRestApi = new RestApi(&#34;testRestApi&#34;);
 * 
 *         var testRestApiPolicy = new RestApiPolicy(&#34;testRestApiPolicy&#34;, RestApiPolicyArgs.builder()        
 *             .restApiId(testRestApi.id())
 *             .policy(testRestApi.executionArn().applyValue(executionArn -&gt; &#34;&#34;&#34;
 * {
 *   &#34;Version&#34;: &#34;2012-10-17&#34;,
 *   &#34;Statement&#34;: [
 *     {
 *       &#34;Effect&#34;: &#34;Allow&#34;,
 *       &#34;Principal&#34;: {
 *         &#34;AWS&#34;: &#34;*&#34;
 *       },
 *       &#34;Action&#34;: &#34;execute-api:Invoke&#34;,
 *       &#34;Resource&#34;: &#34;%s&#34;,
 *       &#34;Condition&#34;: {
 *         &#34;IpAddress&#34;: {
 *           &#34;aws:SourceIp&#34;: &#34;123.123.123.123/32&#34;
 *         }
 *       }
 *     }
 *   ]
 * }
 * &#34;, executionArn)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_api_gateway_rest_api_policy` can be imported by using the REST API ID, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:apigateway/restApiPolicy:RestApiPolicy example 12345abcde
 * ```
 * 
 */
@ResourceType(type="aws:apigateway/restApiPolicy:RestApiPolicy")
public class RestApiPolicy extends com.pulumi.resources.CustomResource {
    /**
     * JSON formatted policy document that controls access to the API Gateway.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return JSON formatted policy document that controls access to the API Gateway.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * ID of the REST API.
     * 
     */
    @Export(name="restApiId", refs={String.class}, tree="[0]")
    private Output<String> restApiId;

    /**
     * @return ID of the REST API.
     * 
     */
    public Output<String> restApiId() {
        return this.restApiId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RestApiPolicy(String name) {
        this(name, RestApiPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RestApiPolicy(String name, RestApiPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RestApiPolicy(String name, RestApiPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/restApiPolicy:RestApiPolicy", name, args == null ? RestApiPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RestApiPolicy(String name, Output<String> id, @Nullable RestApiPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/restApiPolicy:RestApiPolicy", name, state, makeResourceOptions(options, id));
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
    public static RestApiPolicy get(String name, Output<String> id, @Nullable RestApiPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RestApiPolicy(name, id, state, options);
    }
}
