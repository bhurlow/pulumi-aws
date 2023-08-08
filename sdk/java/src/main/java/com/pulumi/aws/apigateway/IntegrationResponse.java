// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apigateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.apigateway.IntegrationResponseArgs;
import com.pulumi.aws.apigateway.inputs.IntegrationResponseState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an HTTP Method Integration Response for an API Gateway Resource.
 * 
 * &gt; **Note:** Depends on having `aws.apigateway.Integration` inside your rest api. To ensure this
 * you might need to add an explicit `depends_on` for clean runs.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.apigateway.RestApi;
 * import com.pulumi.aws.apigateway.RestApiArgs;
 * import com.pulumi.aws.apigateway.Resource;
 * import com.pulumi.aws.apigateway.ResourceArgs;
 * import com.pulumi.aws.apigateway.Method;
 * import com.pulumi.aws.apigateway.MethodArgs;
 * import com.pulumi.aws.apigateway.Integration;
 * import com.pulumi.aws.apigateway.IntegrationArgs;
 * import com.pulumi.aws.apigateway.MethodResponse;
 * import com.pulumi.aws.apigateway.MethodResponseArgs;
 * import com.pulumi.aws.apigateway.IntegrationResponse;
 * import com.pulumi.aws.apigateway.IntegrationResponseArgs;
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
 *         var myDemoAPI = new RestApi(&#34;myDemoAPI&#34;, RestApiArgs.builder()        
 *             .description(&#34;This is my API for demonstration purposes&#34;)
 *             .build());
 * 
 *         var myDemoResource = new Resource(&#34;myDemoResource&#34;, ResourceArgs.builder()        
 *             .restApi(myDemoAPI.id())
 *             .parentId(myDemoAPI.rootResourceId())
 *             .pathPart(&#34;mydemoresource&#34;)
 *             .build());
 * 
 *         var myDemoMethod = new Method(&#34;myDemoMethod&#34;, MethodArgs.builder()        
 *             .restApi(myDemoAPI.id())
 *             .resourceId(myDemoResource.id())
 *             .httpMethod(&#34;GET&#34;)
 *             .authorization(&#34;NONE&#34;)
 *             .build());
 * 
 *         var myDemoIntegration = new Integration(&#34;myDemoIntegration&#34;, IntegrationArgs.builder()        
 *             .restApi(myDemoAPI.id())
 *             .resourceId(myDemoResource.id())
 *             .httpMethod(myDemoMethod.httpMethod())
 *             .type(&#34;MOCK&#34;)
 *             .build());
 * 
 *         var response200 = new MethodResponse(&#34;response200&#34;, MethodResponseArgs.builder()        
 *             .restApi(myDemoAPI.id())
 *             .resourceId(myDemoResource.id())
 *             .httpMethod(myDemoMethod.httpMethod())
 *             .statusCode(&#34;200&#34;)
 *             .build());
 * 
 *         var myDemoIntegrationResponse = new IntegrationResponse(&#34;myDemoIntegrationResponse&#34;, IntegrationResponseArgs.builder()        
 *             .restApi(myDemoAPI.id())
 *             .resourceId(myDemoResource.id())
 *             .httpMethod(myDemoMethod.httpMethod())
 *             .statusCode(response200.statusCode())
 *             .responseTemplates(Map.of(&#34;application/xml&#34;, &#34;&#34;&#34;
 * #set($inputRoot = $input.path(&#39;$&#39;))
 * &lt;?xml version=&#34;1.0&#34; encoding=&#34;UTF-8&#34;?&gt;
 * &lt;message&gt;
 *     $inputRoot.body
 * &lt;/message&gt;
 *             &#34;&#34;&#34;))
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
 *  to = aws_api_gateway_integration_response.example
 * 
 *  id = &#34;12345abcde/67890fghij/GET/200&#34; } Using `pulumi import`, import `aws_api_gateway_integration_response` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`. For exampleconsole % pulumi import aws_api_gateway_integration_response.example 12345abcde/67890fghij/GET/200
 * 
 */
@ResourceType(type="aws:apigateway/integrationResponse:IntegrationResponse")
public class IntegrationResponse extends com.pulumi.resources.CustomResource {
    /**
     * How to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
     * 
     */
    @Export(name="contentHandling", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentHandling;

    /**
     * @return How to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
     * 
     */
    public Output<Optional<String>> contentHandling() {
        return Codegen.optional(this.contentHandling);
    }
    /**
     * HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`).
     * 
     */
    @Export(name="httpMethod", refs={String.class}, tree="[0]")
    private Output<String> httpMethod;

    /**
     * @return HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`).
     * 
     */
    public Output<String> httpMethod() {
        return this.httpMethod;
    }
    /**
     * API resource ID.
     * 
     */
    @Export(name="resourceId", refs={String.class}, tree="[0]")
    private Output<String> resourceId;

    /**
     * @return API resource ID.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * Map of response parameters that can be read from the backend response. For example: `response_parameters = { &#34;method.response.header.X-Some-Header&#34; = &#34;integration.response.header.X-Some-Other-Header&#34; }`.
     * 
     */
    @Export(name="responseParameters", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> responseParameters;

    /**
     * @return Map of response parameters that can be read from the backend response. For example: `response_parameters = { &#34;method.response.header.X-Some-Header&#34; = &#34;integration.response.header.X-Some-Other-Header&#34; }`.
     * 
     */
    public Output<Optional<Map<String,String>>> responseParameters() {
        return Codegen.optional(this.responseParameters);
    }
    /**
     * Map of templates used to transform the integration response body.
     * 
     */
    @Export(name="responseTemplates", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> responseTemplates;

    /**
     * @return Map of templates used to transform the integration response body.
     * 
     */
    public Output<Optional<Map<String,String>>> responseTemplates() {
        return Codegen.optional(this.responseTemplates);
    }
    /**
     * ID of the associated REST API.
     * 
     */
    @Export(name="restApi", refs={String.class}, tree="[0]")
    private Output<String> restApi;

    /**
     * @return ID of the associated REST API.
     * 
     */
    public Output<String> restApi() {
        return this.restApi;
    }
    /**
     * Regular expression pattern used to choose an integration response based on the response from the backend. Omit configuring this to make the integration the default one. If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched. For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
     * 
     */
    @Export(name="selectionPattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> selectionPattern;

    /**
     * @return Regular expression pattern used to choose an integration response based on the response from the backend. Omit configuring this to make the integration the default one. If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched. For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
     * 
     */
    public Output<Optional<String>> selectionPattern() {
        return Codegen.optional(this.selectionPattern);
    }
    /**
     * HTTP status code.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="statusCode", refs={String.class}, tree="[0]")
    private Output<String> statusCode;

    /**
     * @return HTTP status code.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> statusCode() {
        return this.statusCode;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IntegrationResponse(String name) {
        this(name, IntegrationResponseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IntegrationResponse(String name, IntegrationResponseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IntegrationResponse(String name, IntegrationResponseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/integrationResponse:IntegrationResponse", name, args == null ? IntegrationResponseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IntegrationResponse(String name, Output<String> id, @Nullable IntegrationResponseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/integrationResponse:IntegrationResponse", name, state, makeResourceOptions(options, id));
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
    public static IntegrationResponse get(String name, Output<String> id, @Nullable IntegrationResponseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IntegrationResponse(name, id, state, options);
    }
}
