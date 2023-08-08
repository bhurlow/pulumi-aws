// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.iot;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.iot.AuthorizerArgs;
import com.pulumi.aws.iot.inputs.AuthorizerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages an AWS IoT Authorizer.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iot.Authorizer;
 * import com.pulumi.aws.iot.AuthorizerArgs;
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
 *         var example = new Authorizer(&#34;example&#34;, AuthorizerArgs.builder()        
 *             .authorizerFunctionArn(aws_lambda_function.example().arn())
 *             .signingDisabled(false)
 *             .status(&#34;ACTIVE&#34;)
 *             .tokenKeyName(&#34;Token-Header&#34;)
 *             .tokenSigningPublicKeys(Map.of(&#34;Key1&#34;, Files.readString(Paths.get(&#34;test-fixtures/iot-authorizer-signing-key.pem&#34;))))
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
 *  to = aws_iot_authorizer.example
 * 
 *  id = &#34;example&#34; } Using `pulumi import`, import IOT Authorizers using the name. For exampleconsole % pulumi import aws_iot_authorizer.example example
 * 
 */
@ResourceType(type="aws:iot/authorizer:Authorizer")
public class Authorizer extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the authorizer.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the authorizer.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ARN of the authorizer&#39;s Lambda function.
     * 
     */
    @Export(name="authorizerFunctionArn", refs={String.class}, tree="[0]")
    private Output<String> authorizerFunctionArn;

    /**
     * @return The ARN of the authorizer&#39;s Lambda function.
     * 
     */
    public Output<String> authorizerFunctionArn() {
        return this.authorizerFunctionArn;
    }
    /**
     * Specifies whether the HTTP caching is enabled or not. Default: `false`.
     * 
     */
    @Export(name="enableCachingForHttp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableCachingForHttp;

    /**
     * @return Specifies whether the HTTP caching is enabled or not. Default: `false`.
     * 
     */
    public Output<Optional<Boolean>> enableCachingForHttp() {
        return Codegen.optional(this.enableCachingForHttp);
    }
    /**
     * The name of the authorizer.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the authorizer.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies whether AWS IoT validates the token signature in an authorization request. Default: `false`.
     * 
     */
    @Export(name="signingDisabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> signingDisabled;

    /**
     * @return Specifies whether AWS IoT validates the token signature in an authorization request. Default: `false`.
     * 
     */
    public Output<Optional<Boolean>> signingDisabled() {
        return Codegen.optional(this.signingDisabled);
    }
    /**
     * The status of Authorizer request at creation. Valid values: `ACTIVE`, `INACTIVE`. Default: `ACTIVE`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    /**
     * @return The status of Authorizer request at creation. Valid values: `ACTIVE`, `INACTIVE`. Default: `ACTIVE`.
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    /**
     * The name of the token key used to extract the token from the HTTP headers. This value is required if signing is enabled in your authorizer.
     * 
     */
    @Export(name="tokenKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tokenKeyName;

    /**
     * @return The name of the token key used to extract the token from the HTTP headers. This value is required if signing is enabled in your authorizer.
     * 
     */
    public Output<Optional<String>> tokenKeyName() {
        return Codegen.optional(this.tokenKeyName);
    }
    /**
     * The public keys used to verify the digital signature returned by your custom authentication service. This value is required if signing is enabled in your authorizer.
     * 
     */
    @Export(name="tokenSigningPublicKeys", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tokenSigningPublicKeys;

    /**
     * @return The public keys used to verify the digital signature returned by your custom authentication service. This value is required if signing is enabled in your authorizer.
     * 
     */
    public Output<Optional<Map<String,String>>> tokenSigningPublicKeys() {
        return Codegen.optional(this.tokenSigningPublicKeys);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Authorizer(String name) {
        this(name, AuthorizerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Authorizer(String name, AuthorizerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Authorizer(String name, AuthorizerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iot/authorizer:Authorizer", name, args == null ? AuthorizerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Authorizer(String name, Output<String> id, @Nullable AuthorizerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:iot/authorizer:Authorizer", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tokenSigningPublicKeys"
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
    public static Authorizer get(String name, Output<String> id, @Nullable AuthorizerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Authorizer(name, id, state, options);
    }
}
