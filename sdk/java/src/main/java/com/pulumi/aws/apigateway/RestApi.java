// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apigateway;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.apigateway.RestApiArgs;
import com.pulumi.aws.apigateway.inputs.RestApiState;
import com.pulumi.aws.apigateway.outputs.RestApiEndpointConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an API Gateway REST API. The REST API can be configured via [importing an OpenAPI specification](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html) in the `body` argument (with other arguments serving as overrides) or via other provider resources to manage the resources (`aws.apigateway.Resource` resource), methods (`aws.apigateway.Method` resource), integrations (`aws.apigateway.Integration` resource), etc. of the REST API. Once the REST API is configured, the `aws.apigateway.Deployment` resource can be used along with the `aws.apigateway.Stage` resource to publish the REST API.
 * 
 * &gt; **Note:** Amazon API Gateway Version 1 resources are used for creating and deploying REST APIs. To create and deploy WebSocket and HTTP APIs, use Amazon API Gateway Version 2 resources.
 * 
 * !&gt; **WARN:** When importing Open API Specifications with the `body` argument, by default the API Gateway REST API will be replaced with the Open API Specification thus removing any existing methods, resources, integrations, or endpoints. Endpoint mutations are asynchronous operations, and race conditions with DNS are possible. To overcome this limitation, use the `put_rest_api_mode` attribute and set it to `merge`.
 * 
 * ## Example Usage
 * ### OpenAPI Specification
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.apigateway.RestApi;
 * import com.pulumi.aws.apigateway.RestApiArgs;
 * import com.pulumi.aws.apigateway.inputs.RestApiEndpointConfigurationArgs;
 * import com.pulumi.aws.apigateway.Deployment;
 * import com.pulumi.aws.apigateway.DeploymentArgs;
 * import com.pulumi.aws.apigateway.Stage;
 * import com.pulumi.aws.apigateway.StageArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var exampleRestApi = new RestApi(&#34;exampleRestApi&#34;, RestApiArgs.builder()        
 *             .body(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;openapi&#34;, &#34;3.0.1&#34;),
 *                     jsonProperty(&#34;info&#34;, jsonObject(
 *                         jsonProperty(&#34;title&#34;, &#34;example&#34;),
 *                         jsonProperty(&#34;version&#34;, &#34;1.0&#34;)
 *                     )),
 *                     jsonProperty(&#34;paths&#34;, jsonObject(
 *                         jsonProperty(&#34;/path1&#34;, jsonObject(
 *                             jsonProperty(&#34;get&#34;, jsonObject(
 *                                 jsonProperty(&#34;x-amazon-apigateway-integration&#34;, jsonObject(
 *                                     jsonProperty(&#34;httpMethod&#34;, &#34;GET&#34;),
 *                                     jsonProperty(&#34;payloadFormatVersion&#34;, &#34;1.0&#34;),
 *                                     jsonProperty(&#34;type&#34;, &#34;HTTP_PROXY&#34;),
 *                                     jsonProperty(&#34;uri&#34;, &#34;https://ip-ranges.amazonaws.com/ip-ranges.json&#34;)
 *                                 ))
 *                             ))
 *                         ))
 *                     ))
 *                 )))
 *             .endpointConfiguration(RestApiEndpointConfigurationArgs.builder()
 *                 .types(&#34;REGIONAL&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleDeployment = new Deployment(&#34;exampleDeployment&#34;, DeploymentArgs.builder()        
 *             .restApi(exampleRestApi.id())
 *             .triggers(Map.of(&#34;redeployment&#34;, exampleRestApi.body().applyValue(body -&gt; serializeJson(
 *                 body)).applyValue(toJSON -&gt; computeSHA1(toJSON))))
 *             .build());
 * 
 *         var exampleStage = new Stage(&#34;exampleStage&#34;, StageArgs.builder()        
 *             .deployment(exampleDeployment.id())
 *             .restApi(exampleRestApi.id())
 *             .stageName(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Resources
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.apigateway.RestApi;
 * import com.pulumi.aws.apigateway.Resource;
 * import com.pulumi.aws.apigateway.ResourceArgs;
 * import com.pulumi.aws.apigateway.Method;
 * import com.pulumi.aws.apigateway.MethodArgs;
 * import com.pulumi.aws.apigateway.Integration;
 * import com.pulumi.aws.apigateway.IntegrationArgs;
 * import com.pulumi.aws.apigateway.Deployment;
 * import com.pulumi.aws.apigateway.DeploymentArgs;
 * import com.pulumi.aws.apigateway.Stage;
 * import com.pulumi.aws.apigateway.StageArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var exampleRestApi = new RestApi(&#34;exampleRestApi&#34;);
 * 
 *         var exampleResource = new Resource(&#34;exampleResource&#34;, ResourceArgs.builder()        
 *             .parentId(exampleRestApi.rootResourceId())
 *             .pathPart(&#34;example&#34;)
 *             .restApi(exampleRestApi.id())
 *             .build());
 * 
 *         var exampleMethod = new Method(&#34;exampleMethod&#34;, MethodArgs.builder()        
 *             .authorization(&#34;NONE&#34;)
 *             .httpMethod(&#34;GET&#34;)
 *             .resourceId(exampleResource.id())
 *             .restApi(exampleRestApi.id())
 *             .build());
 * 
 *         var exampleIntegration = new Integration(&#34;exampleIntegration&#34;, IntegrationArgs.builder()        
 *             .httpMethod(exampleMethod.httpMethod())
 *             .resourceId(exampleResource.id())
 *             .restApi(exampleRestApi.id())
 *             .type(&#34;MOCK&#34;)
 *             .build());
 * 
 *         var exampleDeployment = new Deployment(&#34;exampleDeployment&#34;, DeploymentArgs.builder()        
 *             .restApi(exampleRestApi.id())
 *             .triggers(Map.of(&#34;redeployment&#34;, Output.tuple(exampleResource.id(), exampleMethod.id(), exampleIntegration.id()).applyValue(values -&gt; {
 *                 var exampleResourceId = values.t1;
 *                 var exampleMethodId = values.t2;
 *                 var exampleIntegrationId = values.t3;
 *                 return serializeJson(
 *                     jsonArray(
 *                         exampleResourceId, 
 *                         exampleMethodId, 
 *                         exampleIntegrationId
 *                     ));
 *             }).applyValue(toJSON -&gt; computeSHA1(toJSON))))
 *             .build());
 * 
 *         var exampleStage = new Stage(&#34;exampleStage&#34;, StageArgs.builder()        
 *             .deployment(exampleDeployment.id())
 *             .restApi(exampleRestApi.id())
 *             .stageName(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_api_gateway_rest_api` can be imported by using the REST API ID, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:apigateway/restApi:RestApi example 12345abcde
 * ```
 * 
 */
@ResourceType(type="aws:apigateway/restApi:RestApi")
public class RestApi extends com.pulumi.resources.CustomResource {
    /**
     * Source of the API key for requests. Valid values are `HEADER` (default) and `AUTHORIZER`. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-api-key-source` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-api-key-source.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="apiKeySource", refs={String.class}, tree="[0]")
    private Output<String> apiKeySource;

    /**
     * @return Source of the API key for requests. Valid values are `HEADER` (default) and `AUTHORIZER`. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-api-key-source` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-api-key-source.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<String> apiKeySource() {
        return this.apiKeySource;
    }
    /**
     * ARN
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-binary-media-types` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-binary-media-types.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="binaryMediaTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> binaryMediaTypes;

    /**
     * @return List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-binary-media-types` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-binary-media-types.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<List<String>> binaryMediaTypes() {
        return this.binaryMediaTypes;
    }
    /**
     * OpenAPI specification that defines the set of routes and integrations to create as part of the REST API. This configuration, and any updates to it, will replace all REST API configuration except values overridden in this resource configuration and other resource updates applied after this resource but before any `aws.apigateway.Deployment` creation. More information about REST API OpenAPI support can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html).
     * 
     */
    @Export(name="body", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> body;

    /**
     * @return OpenAPI specification that defines the set of routes and integrations to create as part of the REST API. This configuration, and any updates to it, will replace all REST API configuration except values overridden in this resource configuration and other resource updates applied after this resource but before any `aws.apigateway.Deployment` creation. More information about REST API OpenAPI support can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html).
     * 
     */
    public Output<Optional<String>> body() {
        return Codegen.optional(this.body);
    }
    /**
     * Creation date of the REST API
     * 
     */
    @Export(name="createdDate", refs={String.class}, tree="[0]")
    private Output<String> createdDate;

    /**
     * @return Creation date of the REST API
     * 
     */
    public Output<String> createdDate() {
        return this.createdDate;
    }
    /**
     * Description of the REST API. If importing an OpenAPI specification via the `body` argument, this corresponds to the `info.description` field. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description of the REST API. If importing an OpenAPI specification via the `body` argument, this corresponds to the `info.description` field. If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to `false`. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-endpoint-configuration` extension `disableExecuteApiEndpoint` property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-endpoint-configuration.html). If the argument value is `true` and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="disableExecuteApiEndpoint", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disableExecuteApiEndpoint;

    /**
     * @return Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to `false`. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-endpoint-configuration` extension `disableExecuteApiEndpoint` property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-endpoint-configuration.html). If the argument value is `true` and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<Boolean> disableExecuteApiEndpoint() {
        return this.disableExecuteApiEndpoint;
    }
    /**
     * Configuration block defining API endpoint configuration including endpoint type. Defined below.
     * 
     */
    @Export(name="endpointConfiguration", refs={RestApiEndpointConfiguration.class}, tree="[0]")
    private Output<RestApiEndpointConfiguration> endpointConfiguration;

    /**
     * @return Configuration block defining API endpoint configuration including endpoint type. Defined below.
     * 
     */
    public Output<RestApiEndpointConfiguration> endpointConfiguration() {
        return this.endpointConfiguration;
    }
    /**
     * Execution ARN part to be used in `lambda_permission`&#39;s `source_arn`
     * when allowing API Gateway to invoke a Lambda function,
     * e.g., `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
     * 
     */
    @Export(name="executionArn", refs={String.class}, tree="[0]")
    private Output<String> executionArn;

    /**
     * @return Execution ARN part to be used in `lambda_permission`&#39;s `source_arn`
     * when allowing API Gateway to invoke a Lambda function,
     * e.g., `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
     * 
     */
    public Output<String> executionArn() {
        return this.executionArn;
    }
    /**
     * Minimum response size to compress for the REST API. Integer between `-1` and `10485760` (10MB). Setting a value greater than `-1` will enable compression, `-1` disables compression (default). If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-minimum-compression-size` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-openapi-minimum-compression-size.html). If the argument value (_except_ `-1`) is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="minimumCompressionSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minimumCompressionSize;

    /**
     * @return Minimum response size to compress for the REST API. Integer between `-1` and `10485760` (10MB). Setting a value greater than `-1` will enable compression, `-1` disables compression (default). If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-minimum-compression-size` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-openapi-minimum-compression-size.html). If the argument value (_except_ `-1`) is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<Optional<Integer>> minimumCompressionSize() {
        return Codegen.optional(this.minimumCompressionSize);
    }
    /**
     * Name of the REST API. If importing an OpenAPI specification via the `body` argument, this corresponds to the `info.title` field. If the argument value is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the REST API. If importing an OpenAPI specification via the `body` argument, this corresponds to the `info.title` field. If the argument value is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Map of customizations for importing the specification in the `body` argument. For example, to exclude DocumentationParts from an imported API, set `ignore` equal to `documentation`. Additional documentation, including other parameters such as `basepath`, can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html).
     * 
     */
    @Export(name="parameters", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> parameters;

    /**
     * @return Map of customizations for importing the specification in the `body` argument. For example, to exclude DocumentationParts from an imported API, set `ignore` equal to `documentation`. Additional documentation, including other parameters such as `basepath`, can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html).
     * 
     */
    public Output<Optional<Map<String,String>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * JSON formatted policy document that controls access to the API Gateway. For more information about building AWS IAM policy documents with Pulumi, see the AWS IAM Policy Document Guide. The provider will only perform drift detection of its value when present in a configuration. We recommend using the `aws.apigateway.RestApiPolicy` resource instead. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-policy` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/openapi-extensions-policy.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return JSON formatted policy document that controls access to the API Gateway. For more information about building AWS IAM policy documents with Pulumi, see the AWS IAM Policy Document Guide. The provider will only perform drift detection of its value when present in a configuration. We recommend using the `aws.apigateway.RestApiPolicy` resource instead. If importing an OpenAPI specification via the `body` argument, this corresponds to the [`x-amazon-apigateway-policy` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/openapi-extensions-policy.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * Mode of the PutRestApi operation when importing an OpenAPI specification via the `body` argument (create or update operation). Valid values are `merge` and `overwrite`. If unspecificed, defaults to `overwrite` (for backwards compatibility). This corresponds to the [`x-amazon-apigateway-put-integration-method` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-put-integration-method.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    @Export(name="putRestApiMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> putRestApiMode;

    /**
     * @return Mode of the PutRestApi operation when importing an OpenAPI specification via the `body` argument (create or update operation). Valid values are `merge` and `overwrite`. If unspecificed, defaults to `overwrite` (for backwards compatibility). This corresponds to the [`x-amazon-apigateway-put-integration-method` extension](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-put-integration-method.html). If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     * 
     */
    public Output<Optional<String>> putRestApiMode() {
        return Codegen.optional(this.putRestApiMode);
    }
    /**
     * Resource ID of the REST API&#39;s root
     * 
     */
    @Export(name="rootResourceId", refs={String.class}, tree="[0]")
    private Output<String> rootResourceId;

    /**
     * @return Resource ID of the REST API&#39;s root
     * 
     */
    public Output<String> rootResourceId() {
        return this.rootResourceId;
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
    public RestApi(String name) {
        this(name, RestApiArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RestApi(String name, @Nullable RestApiArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RestApi(String name, @Nullable RestApiArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/restApi:RestApi", name, args == null ? RestApiArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RestApi(String name, Output<String> id, @Nullable RestApiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:apigateway/restApi:RestApi", name, state, makeResourceOptions(options, id));
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
    public static RestApi get(String name, Output<String> id, @Nullable RestApiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RestApi(name, id, state, options);
    }
}
