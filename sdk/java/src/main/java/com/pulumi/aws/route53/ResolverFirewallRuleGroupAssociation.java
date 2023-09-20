// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.route53.ResolverFirewallRuleGroupAssociationArgs;
import com.pulumi.aws.route53.inputs.ResolverFirewallRuleGroupAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Route 53 Resolver DNS Firewall rule group association resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53.ResolverFirewallRuleGroup;
 * import com.pulumi.aws.route53.ResolverFirewallRuleGroupAssociation;
 * import com.pulumi.aws.route53.ResolverFirewallRuleGroupAssociationArgs;
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
 *         var exampleResolverFirewallRuleGroup = new ResolverFirewallRuleGroup(&#34;exampleResolverFirewallRuleGroup&#34;);
 * 
 *         var exampleResolverFirewallRuleGroupAssociation = new ResolverFirewallRuleGroupAssociation(&#34;exampleResolverFirewallRuleGroupAssociation&#34;, ResolverFirewallRuleGroupAssociationArgs.builder()        
 *             .firewallRuleGroupId(exampleResolverFirewallRuleGroup.id())
 *             .priority(100)
 *             .vpcId(aws_vpc.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * In TODO v1.5.0 and later, use an `import` block to import Route 53 Resolver DNS Firewall rule group associations using the Route 53 Resolver DNS Firewall rule group association ID. For exampleterraform import {
 * 
 *  to = aws_route53_resolver_firewall_rule_group_association.example
 * 
 *  id = &#34;rslvr-frgassoc-0123456789abcdef&#34; } Using `TODO import`, import Route 53 Resolver DNS Firewall rule group associations using the Route 53 Resolver DNS Firewall rule group association ID. For exampleconsole % TODO import aws_route53_resolver_firewall_rule_group_association.example rslvr-frgassoc-0123456789abcdef
 * 
 */
@ResourceType(type="aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation")
public class ResolverFirewallRuleGroupAssociation extends com.pulumi.resources.CustomResource {
    /**
     * The ARN (Amazon Resource Name) of the firewall rule group association.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN (Amazon Resource Name) of the firewall rule group association.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The unique identifier of the firewall rule group.
     * 
     */
    @Export(name="firewallRuleGroupId", refs={String.class}, tree="[0]")
    private Output<String> firewallRuleGroupId;

    /**
     * @return The unique identifier of the firewall rule group.
     * 
     */
    public Output<String> firewallRuleGroupId() {
        return this.firewallRuleGroupId;
    }
    /**
     * If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections. Valid values: `ENABLED`, `DISABLED`.
     * 
     */
    @Export(name="mutationProtection", refs={String.class}, tree="[0]")
    private Output<String> mutationProtection;

    /**
     * @return If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections. Valid values: `ENABLED`, `DISABLED`.
     * 
     */
    public Output<String> mutationProtection() {
        return this.mutationProtection;
    }
    /**
     * A name that lets you identify the rule group association, to manage and use it.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name that lets you identify the rule group association, to manage and use it.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC. DNS Firewall filters VPC traffic starting from the rule group with the lowest numeric priority setting.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output<Integer> priority;

    /**
     * @return The setting that determines the processing order of the rule group among the rule groups that you associate with the specified VPC. DNS Firewall filters VPC traffic starting from the rule group with the lowest numeric priority setting.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
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
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
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
     * The unique identifier of the VPC that you want to associate with the rule group.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The unique identifier of the VPC that you want to associate with the rule group.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResolverFirewallRuleGroupAssociation(String name) {
        this(name, ResolverFirewallRuleGroupAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResolverFirewallRuleGroupAssociation(String name, ResolverFirewallRuleGroupAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResolverFirewallRuleGroupAssociation(String name, ResolverFirewallRuleGroupAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation", name, args == null ? ResolverFirewallRuleGroupAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResolverFirewallRuleGroupAssociation(String name, Output<String> id, @Nullable ResolverFirewallRuleGroupAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/resolverFirewallRuleGroupAssociation:ResolverFirewallRuleGroupAssociation", name, state, makeResourceOptions(options, id));
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
    public static ResolverFirewallRuleGroupAssociation get(String name, Output<String> id, @Nullable ResolverFirewallRuleGroupAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResolverFirewallRuleGroupAssociation(name, id, state, options);
    }
}
