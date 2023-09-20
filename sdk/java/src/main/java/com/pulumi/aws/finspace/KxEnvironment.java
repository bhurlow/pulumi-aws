// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.finspace;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.finspace.KxEnvironmentArgs;
import com.pulumi.aws.finspace.inputs.KxEnvironmentState;
import com.pulumi.aws.finspace.outputs.KxEnvironmentCustomDnsConfiguration;
import com.pulumi.aws.finspace.outputs.KxEnvironmentTransitGatewayConfiguration;
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
 * Resource for managing an AWS FinSpace Kx Environment.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.finspace.KxEnvironment;
 * import com.pulumi.aws.finspace.KxEnvironmentArgs;
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
 *         var exampleKey = new Key(&#34;exampleKey&#34;, KeyArgs.builder()        
 *             .description(&#34;Sample KMS Key&#34;)
 *             .deletionWindowInDays(7)
 *             .build());
 * 
 *         var exampleKxEnvironment = new KxEnvironment(&#34;exampleKxEnvironment&#34;, KxEnvironmentArgs.builder()        
 *             .kmsKeyId(exampleKey.arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Transit Gateway Configuration
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.kms.Key;
 * import com.pulumi.aws.kms.KeyArgs;
 * import com.pulumi.aws.ec2transitgateway.TransitGateway;
 * import com.pulumi.aws.ec2transitgateway.TransitGatewayArgs;
 * import com.pulumi.aws.finspace.KxEnvironment;
 * import com.pulumi.aws.finspace.KxEnvironmentArgs;
 * import com.pulumi.aws.finspace.inputs.KxEnvironmentTransitGatewayConfigurationArgs;
 * import com.pulumi.aws.finspace.inputs.KxEnvironmentCustomDnsConfigurationArgs;
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
 *         var exampleKey = new Key(&#34;exampleKey&#34;, KeyArgs.builder()        
 *             .description(&#34;Sample KMS Key&#34;)
 *             .deletionWindowInDays(7)
 *             .build());
 * 
 *         var exampleTransitGateway = new TransitGateway(&#34;exampleTransitGateway&#34;, TransitGatewayArgs.builder()        
 *             .description(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleEnv = new KxEnvironment(&#34;exampleEnv&#34;, KxEnvironmentArgs.builder()        
 *             .description(&#34;Environment description&#34;)
 *             .kmsKeyId(exampleKey.arn())
 *             .transitGatewayConfiguration(KxEnvironmentTransitGatewayConfigurationArgs.builder()
 *                 .transitGatewayId(exampleTransitGateway.id())
 *                 .routableCidrSpace(&#34;100.64.0.0/26&#34;)
 *                 .build())
 *             .customDnsConfigurations(KxEnvironmentCustomDnsConfigurationArgs.builder()
 *                 .customDnsServerName(&#34;example.finspace.amazonaws.com&#34;)
 *                 .customDnsServerIp(&#34;10.0.0.76&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import an AWS FinSpace Kx Environment using the `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:finspace/kxEnvironment:KxEnvironment example n3ceo7wqxoxcti5tujqwzs
 * ```
 * 
 */
@ResourceType(type="aws:finspace/kxEnvironment:KxEnvironment")
public class KxEnvironment extends com.pulumi.resources.CustomResource {
    /**
     * Amazon Resource Name (ARN) identifier of the KX environment.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) identifier of the KX environment.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * AWS Availability Zone IDs that this environment is available in. Important when selecting VPC subnets to use in cluster creation.
     * 
     */
    @Export(name="availabilityZones", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> availabilityZones;

    /**
     * @return AWS Availability Zone IDs that this environment is available in. Important when selecting VPC subnets to use in cluster creation.
     * 
     */
    public Output<List<String>> availabilityZones() {
        return this.availabilityZones;
    }
    /**
     * Timestamp at which the environment is created in FinSpace. Value determined as epoch time in seconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000.
     * 
     */
    @Export(name="createdTimestamp", refs={String.class}, tree="[0]")
    private Output<String> createdTimestamp;

    /**
     * @return Timestamp at which the environment is created in FinSpace. Value determined as epoch time in seconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000.
     * 
     */
    public Output<String> createdTimestamp() {
        return this.createdTimestamp;
    }
    /**
     * List of DNS server name and server IP. This is used to set up Route-53 outbound resolvers. Defined below.
     * 
     */
    @Export(name="customDnsConfigurations", refs={List.class,KxEnvironmentCustomDnsConfiguration.class}, tree="[0,1]")
    private Output</* @Nullable */ List<KxEnvironmentCustomDnsConfiguration>> customDnsConfigurations;

    /**
     * @return List of DNS server name and server IP. This is used to set up Route-53 outbound resolvers. Defined below.
     * 
     */
    public Output<Optional<List<KxEnvironmentCustomDnsConfiguration>>> customDnsConfigurations() {
        return Codegen.optional(this.customDnsConfigurations);
    }
    /**
     * Description for the KX environment.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description for the KX environment.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Unique identifier for the AWS environment infrastructure account.
     * 
     */
    @Export(name="infrastructureAccountId", refs={String.class}, tree="[0]")
    private Output<String> infrastructureAccountId;

    /**
     * @return Unique identifier for the AWS environment infrastructure account.
     * 
     */
    public Output<String> infrastructureAccountId() {
        return this.infrastructureAccountId;
    }
    /**
     * KMS key ID to encrypt your data in the FinSpace environment.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="kmsKeyId", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyId;

    /**
     * @return KMS key ID to encrypt your data in the FinSpace environment.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> kmsKeyId() {
        return this.kmsKeyId;
    }
    /**
     * Last timestamp at which the environment was updated in FinSpace. Value determined as epoch time in seconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000.
     * 
     */
    @Export(name="lastModifiedTimestamp", refs={String.class}, tree="[0]")
    private Output<String> lastModifiedTimestamp;

    /**
     * @return Last timestamp at which the environment was updated in FinSpace. Value determined as epoch time in seconds. For example, the value for Monday, November 1, 2021 12:00:00 PM UTC is specified as 1635768000.
     * 
     */
    public Output<String> lastModifiedTimestamp() {
        return this.lastModifiedTimestamp;
    }
    /**
     * Name of the KX environment that you want to create.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the KX environment that you want to create.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Status of environment creation
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of environment creation
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
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
     * Transit gateway and network configuration that is used to connect the KX environment to an internal network. Defined below.
     * 
     */
    @Export(name="transitGatewayConfiguration", refs={KxEnvironmentTransitGatewayConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ KxEnvironmentTransitGatewayConfiguration> transitGatewayConfiguration;

    /**
     * @return Transit gateway and network configuration that is used to connect the KX environment to an internal network. Defined below.
     * 
     */
    public Output<Optional<KxEnvironmentTransitGatewayConfiguration>> transitGatewayConfiguration() {
        return Codegen.optional(this.transitGatewayConfiguration);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KxEnvironment(String name) {
        this(name, KxEnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KxEnvironment(String name, KxEnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KxEnvironment(String name, KxEnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:finspace/kxEnvironment:KxEnvironment", name, args == null ? KxEnvironmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KxEnvironment(String name, Output<String> id, @Nullable KxEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:finspace/kxEnvironment:KxEnvironment", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
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
    public static KxEnvironment get(String name, Output<String> id, @Nullable KxEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KxEnvironment(name, id, state, options);
    }
}
