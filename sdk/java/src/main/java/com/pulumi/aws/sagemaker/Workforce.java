// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.sagemaker.WorkforceArgs;
import com.pulumi.aws.sagemaker.inputs.WorkforceState;
import com.pulumi.aws.sagemaker.outputs.WorkforceCognitoConfig;
import com.pulumi.aws.sagemaker.outputs.WorkforceOidcConfig;
import com.pulumi.aws.sagemaker.outputs.WorkforceSourceIpConfig;
import com.pulumi.aws.sagemaker.outputs.WorkforceWorkforceVpcConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a SageMaker Workforce resource.
 * 
 * ## Example Usage
 * ### Cognito Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cognito.UserPool;
 * import com.pulumi.aws.cognito.UserPoolClient;
 * import com.pulumi.aws.cognito.UserPoolClientArgs;
 * import com.pulumi.aws.cognito.UserPoolDomain;
 * import com.pulumi.aws.cognito.UserPoolDomainArgs;
 * import com.pulumi.aws.sagemaker.Workforce;
 * import com.pulumi.aws.sagemaker.WorkforceArgs;
 * import com.pulumi.aws.sagemaker.inputs.WorkforceCognitoConfigArgs;
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
 *         var exampleUserPool = new UserPool(&#34;exampleUserPool&#34;);
 * 
 *         var exampleUserPoolClient = new UserPoolClient(&#34;exampleUserPoolClient&#34;, UserPoolClientArgs.builder()        
 *             .generateSecret(true)
 *             .userPoolId(exampleUserPool.id())
 *             .build());
 * 
 *         var exampleUserPoolDomain = new UserPoolDomain(&#34;exampleUserPoolDomain&#34;, UserPoolDomainArgs.builder()        
 *             .domain(&#34;example&#34;)
 *             .userPoolId(exampleUserPool.id())
 *             .build());
 * 
 *         var exampleWorkforce = new Workforce(&#34;exampleWorkforce&#34;, WorkforceArgs.builder()        
 *             .workforceName(&#34;example&#34;)
 *             .cognitoConfig(WorkforceCognitoConfigArgs.builder()
 *                 .clientId(exampleUserPoolClient.id())
 *                 .userPool(exampleUserPoolDomain.userPoolId())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Oidc Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sagemaker.Workforce;
 * import com.pulumi.aws.sagemaker.WorkforceArgs;
 * import com.pulumi.aws.sagemaker.inputs.WorkforceOidcConfigArgs;
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
 *         var example = new Workforce(&#34;example&#34;, WorkforceArgs.builder()        
 *             .oidcConfig(WorkforceOidcConfigArgs.builder()
 *                 .authorizationEndpoint(&#34;https://example.com&#34;)
 *                 .clientId(&#34;example&#34;)
 *                 .clientSecret(&#34;example&#34;)
 *                 .issuer(&#34;https://example.com&#34;)
 *                 .jwksUri(&#34;https://example.com&#34;)
 *                 .logoutEndpoint(&#34;https://example.com&#34;)
 *                 .tokenEndpoint(&#34;https://example.com&#34;)
 *                 .userInfoEndpoint(&#34;https://example.com&#34;)
 *                 .build())
 *             .workforceName(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SageMaker Workforces can be imported using the `workforce_name`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:sagemaker/workforce:Workforce example example
 * ```
 * 
 */
@ResourceType(type="aws:sagemaker/workforce:Workforce")
public class Workforce extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Workforce.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) assigned by AWS to this Workforce.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with `oidc_config`. see Cognito Config details below.
     * 
     */
    @Export(name="cognitoConfig", refs={WorkforceCognitoConfig.class}, tree="[0]")
    private Output</* @Nullable */ WorkforceCognitoConfig> cognitoConfig;

    /**
     * @return Use this parameter to configure an Amazon Cognito private workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. Conflicts with `oidc_config`. see Cognito Config details below.
     * 
     */
    public Output<Optional<WorkforceCognitoConfig>> cognitoConfig() {
        return Codegen.optional(this.cognitoConfig);
    }
    /**
     * Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with `cognito_config`. see OIDC Config details below.
     * 
     */
    @Export(name="oidcConfig", refs={WorkforceOidcConfig.class}, tree="[0]")
    private Output</* @Nullable */ WorkforceOidcConfig> oidcConfig;

    /**
     * @return Use this parameter to configure a private workforce using your own OIDC Identity Provider. Conflicts with `cognito_config`. see OIDC Config details below.
     * 
     */
    public Output<Optional<WorkforceOidcConfig>> oidcConfig() {
        return Codegen.optional(this.oidcConfig);
    }
    /**
     * A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn&#39;t restricted to specific IP addresses. see Source Ip Config details below.
     * 
     */
    @Export(name="sourceIpConfig", refs={WorkforceSourceIpConfig.class}, tree="[0]")
    private Output<WorkforceSourceIpConfig> sourceIpConfig;

    /**
     * @return A list of IP address ranges Used to create an allow list of IP addresses for a private workforce. By default, a workforce isn&#39;t restricted to specific IP addresses. see Source Ip Config details below.
     * 
     */
    public Output<WorkforceSourceIpConfig> sourceIpConfig() {
        return this.sourceIpConfig;
    }
    /**
     * The subdomain for your OIDC Identity Provider.
     * * `workforce_vpc_config.0.vpc_endpoint_id` - The IDs for the VPC service endpoints of your VPC workforce.
     * 
     */
    @Export(name="subdomain", refs={String.class}, tree="[0]")
    private Output<String> subdomain;

    /**
     * @return The subdomain for your OIDC Identity Provider.
     * * `workforce_vpc_config.0.vpc_endpoint_id` - The IDs for the VPC service endpoints of your VPC workforce.
     * 
     */
    public Output<String> subdomain() {
        return this.subdomain;
    }
    /**
     * The name of the Workforce (must be unique).
     * 
     */
    @Export(name="workforceName", refs={String.class}, tree="[0]")
    private Output<String> workforceName;

    /**
     * @return The name of the Workforce (must be unique).
     * 
     */
    public Output<String> workforceName() {
        return this.workforceName;
    }
    /**
     * configure a workforce using VPC. see Workforce VPC Config details below.
     * 
     */
    @Export(name="workforceVpcConfig", refs={WorkforceWorkforceVpcConfig.class}, tree="[0]")
    private Output</* @Nullable */ WorkforceWorkforceVpcConfig> workforceVpcConfig;

    /**
     * @return configure a workforce using VPC. see Workforce VPC Config details below.
     * 
     */
    public Output<Optional<WorkforceWorkforceVpcConfig>> workforceVpcConfig() {
        return Codegen.optional(this.workforceVpcConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Workforce(String name) {
        this(name, WorkforceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Workforce(String name, WorkforceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Workforce(String name, WorkforceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/workforce:Workforce", name, args == null ? WorkforceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Workforce(String name, Output<String> id, @Nullable WorkforceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/workforce:Workforce", name, state, makeResourceOptions(options, id));
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
    public static Workforce get(String name, Output<String> id, @Nullable WorkforceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Workforce(name, id, state, options);
    }
}
