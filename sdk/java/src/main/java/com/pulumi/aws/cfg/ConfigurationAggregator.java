// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cfg;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cfg.ConfigurationAggregatorArgs;
import com.pulumi.aws.cfg.inputs.ConfigurationAggregatorState;
import com.pulumi.aws.cfg.outputs.ConfigurationAggregatorAccountAggregationSource;
import com.pulumi.aws.cfg.outputs.ConfigurationAggregatorOrganizationAggregationSource;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an AWS Config Configuration Aggregator
 * 
 * ## Example Usage
 * ### Account Based Aggregation
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cfg.ConfigurationAggregator;
 * import com.pulumi.aws.cfg.ConfigurationAggregatorArgs;
 * import com.pulumi.aws.cfg.inputs.ConfigurationAggregatorAccountAggregationSourceArgs;
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
 *         var account = new ConfigurationAggregator(&#34;account&#34;, ConfigurationAggregatorArgs.builder()        
 *             .accountAggregationSource(ConfigurationAggregatorAccountAggregationSourceArgs.builder()
 *                 .accountIds(&#34;123456789012&#34;)
 *                 .regions(&#34;us-west-2&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Organization Based Aggregation
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.iam.RolePolicyAttachment;
 * import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
 * import com.pulumi.aws.cfg.ConfigurationAggregator;
 * import com.pulumi.aws.cfg.ConfigurationAggregatorArgs;
 * import com.pulumi.aws.cfg.inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;config.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var organizationRole = new Role(&#34;organizationRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var organizationRolePolicyAttachment = new RolePolicyAttachment(&#34;organizationRolePolicyAttachment&#34;, RolePolicyAttachmentArgs.builder()        
 *             .role(organizationRole.name())
 *             .policyArn(&#34;arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations&#34;)
 *             .build());
 * 
 *         var organizationConfigurationAggregator = new ConfigurationAggregator(&#34;organizationConfigurationAggregator&#34;, ConfigurationAggregatorArgs.builder()        
 *             .organizationAggregationSource(ConfigurationAggregatorOrganizationAggregationSourceArgs.builder()
 *                 .allRegions(true)
 *                 .roleArn(organizationRole.arn())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(organizationRolePolicyAttachment)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_config_configuration_aggregator.example
 * 
 *  id = &#34;foo&#34; } Using `pulumi import`, import Configuration Aggregators using the name. For exampleconsole % pulumi import aws_config_configuration_aggregator.example foo
 * 
 */
@ResourceType(type="aws:cfg/configurationAggregator:ConfigurationAggregator")
public class ConfigurationAggregator extends com.pulumi.resources.CustomResource {
    /**
     * The account(s) to aggregate config data from as documented below.
     * 
     */
    @Export(name="accountAggregationSource", refs={ConfigurationAggregatorAccountAggregationSource.class}, tree="[0]")
    private Output</* @Nullable */ ConfigurationAggregatorAccountAggregationSource> accountAggregationSource;

    /**
     * @return The account(s) to aggregate config data from as documented below.
     * 
     */
    public Output<Optional<ConfigurationAggregatorAccountAggregationSource>> accountAggregationSource() {
        return Codegen.optional(this.accountAggregationSource);
    }
    /**
     * The ARN of the aggregator
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the aggregator
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The name of the configuration aggregator.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the configuration aggregator.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization to aggregate config data from as documented below.
     * 
     */
    @Export(name="organizationAggregationSource", refs={ConfigurationAggregatorOrganizationAggregationSource.class}, tree="[0]")
    private Output</* @Nullable */ ConfigurationAggregatorOrganizationAggregationSource> organizationAggregationSource;

    /**
     * @return The organization to aggregate config data from as documented below.
     * 
     */
    public Output<Optional<ConfigurationAggregatorOrganizationAggregationSource>> organizationAggregationSource() {
        return Codegen.optional(this.organizationAggregationSource);
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * Either `account_aggregation_source` or `organization_aggregation_source` must be specified.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * Either `account_aggregation_source` or `organization_aggregation_source` must be specified.
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
    public ConfigurationAggregator(String name) {
        this(name, ConfigurationAggregatorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConfigurationAggregator(String name, @Nullable ConfigurationAggregatorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConfigurationAggregator(String name, @Nullable ConfigurationAggregatorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cfg/configurationAggregator:ConfigurationAggregator", name, args == null ? ConfigurationAggregatorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConfigurationAggregator(String name, Output<String> id, @Nullable ConfigurationAggregatorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cfg/configurationAggregator:ConfigurationAggregator", name, state, makeResourceOptions(options, id));
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
    public static ConfigurationAggregator get(String name, Output<String> id, @Nullable ConfigurationAggregatorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConfigurationAggregator(name, id, state, options);
    }
}
