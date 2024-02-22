// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.aws.route53.inputs.RecordState;
import com.pulumi.aws.route53.outputs.RecordAlias;
import com.pulumi.aws.route53.outputs.RecordCidrRoutingPolicy;
import com.pulumi.aws.route53.outputs.RecordFailoverRoutingPolicy;
import com.pulumi.aws.route53.outputs.RecordGeolocationRoutingPolicy;
import com.pulumi.aws.route53.outputs.RecordGeoproximityRoutingPolicy;
import com.pulumi.aws.route53.outputs.RecordLatencyRoutingPolicy;
import com.pulumi.aws.route53.outputs.RecordWeightedRoutingPolicy;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Route53 record resource.
 * 
 * ## Example Usage
 * ### Simple routing policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53.Record;
 * import com.pulumi.aws.route53.RecordArgs;
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
 *         var www = new Record(&#34;www&#34;, RecordArgs.builder()        
 *             .zoneId(aws_route53_zone.primary().zone_id())
 *             .name(&#34;www.example.com&#34;)
 *             .type(&#34;A&#34;)
 *             .ttl(300)
 *             .records(aws_eip.lb().public_ip())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Weighted routing policy
 * 
 * Other routing policies are configured similarly. See [Amazon Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html) for details.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53.Record;
 * import com.pulumi.aws.route53.RecordArgs;
 * import com.pulumi.aws.route53.inputs.RecordWeightedRoutingPolicyArgs;
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
 *         var www_dev = new Record(&#34;www-dev&#34;, RecordArgs.builder()        
 *             .zoneId(aws_route53_zone.primary().zone_id())
 *             .name(&#34;www&#34;)
 *             .type(&#34;CNAME&#34;)
 *             .ttl(5)
 *             .weightedRoutingPolicies(RecordWeightedRoutingPolicyArgs.builder()
 *                 .weight(10)
 *                 .build())
 *             .setIdentifier(&#34;dev&#34;)
 *             .records(&#34;dev.example.com&#34;)
 *             .build());
 * 
 *         var www_live = new Record(&#34;www-live&#34;, RecordArgs.builder()        
 *             .zoneId(aws_route53_zone.primary().zone_id())
 *             .name(&#34;www&#34;)
 *             .type(&#34;CNAME&#34;)
 *             .ttl(5)
 *             .weightedRoutingPolicies(RecordWeightedRoutingPolicyArgs.builder()
 *                 .weight(90)
 *                 .build())
 *             .setIdentifier(&#34;live&#34;)
 *             .records(&#34;live.example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Geoproximity routing policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53.Record;
 * import com.pulumi.aws.route53.RecordArgs;
 * import com.pulumi.aws.route53.inputs.RecordGeoproximityRoutingPolicyArgs;
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
 *         var www = new Record(&#34;www&#34;, RecordArgs.builder()        
 *             .zoneId(aws_route53_zone.primary().zone_id())
 *             .name(&#34;www.example.com&#34;)
 *             .type(&#34;CNAME&#34;)
 *             .ttl(300)
 *             .geoproximityRoutingPolicy(RecordGeoproximityRoutingPolicyArgs.builder()
 *                 .coordinates(RecordGeoproximityRoutingPolicyCoordinateArgs.builder()
 *                     .latitude(&#34;49.22&#34;)
 *                     .longitude(&#34;-74.01&#34;)
 *                     .build())
 *                 .build())
 *             .setIdentifier(&#34;dev&#34;)
 *             .records(&#34;dev.example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Alias record
 * 
 * See [related part of Amazon Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html)
 * to understand differences between alias and non-alias records.
 * 
 * TTL for all alias records is [60 seconds](https://aws.amazon.com/route53/faqs/#dns_failover_do_i_need_to_adjust),
 * you cannot change this, therefore `ttl` has to be omitted in alias records.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.elb.LoadBalancer;
 * import com.pulumi.aws.elb.LoadBalancerArgs;
 * import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
 * import com.pulumi.aws.route53.Record;
 * import com.pulumi.aws.route53.RecordArgs;
 * import com.pulumi.aws.route53.inputs.RecordAliasArgs;
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
 *         var main = new LoadBalancer(&#34;main&#34;, LoadBalancerArgs.builder()        
 *             .availabilityZones(&#34;us-east-1c&#34;)
 *             .listeners(LoadBalancerListenerArgs.builder()
 *                 .instancePort(80)
 *                 .instanceProtocol(&#34;http&#34;)
 *                 .lbPort(80)
 *                 .lbProtocol(&#34;http&#34;)
 *                 .build())
 *             .build());
 * 
 *         var www = new Record(&#34;www&#34;, RecordArgs.builder()        
 *             .zoneId(aws_route53_zone.primary().zone_id())
 *             .name(&#34;example.com&#34;)
 *             .type(&#34;A&#34;)
 *             .aliases(RecordAliasArgs.builder()
 *                 .name(main.dnsName())
 *                 .zoneId(main.zoneId())
 *                 .evaluateTargetHealth(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### NS and SOA Record Management
 * 
 * When creating Route 53 zones, the `NS` and `SOA` records for the zone are automatically created. Enabling the `allow_overwrite` argument will allow managing these records in a single deployment without the requirement for `import`.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53.Zone;
 * import com.pulumi.aws.route53.Record;
 * import com.pulumi.aws.route53.RecordArgs;
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
 *         var exampleZone = new Zone(&#34;exampleZone&#34;);
 * 
 *         var exampleRecord = new Record(&#34;exampleRecord&#34;, RecordArgs.builder()        
 *             .allowOverwrite(true)
 *             .name(&#34;test.example.com&#34;)
 *             .ttl(172800)
 *             .type(&#34;NS&#34;)
 *             .zoneId(exampleZone.zoneId())
 *             .records(            
 *                 exampleZone.nameServers().applyValue(nameServers -&gt; nameServers[0]),
 *                 exampleZone.nameServers().applyValue(nameServers -&gt; nameServers[1]),
 *                 exampleZone.nameServers().applyValue(nameServers -&gt; nameServers[2]),
 *                 exampleZone.nameServers().applyValue(nameServers -&gt; nameServers[3]))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * If the record also contains a set identifier, append it:
 * 
 * If the record name is the empty string, it can be omitted:
 * 
 * __Using `pulumi import` to import__ Route53 Records using the ID of the record, record name, record type, and set identifier. For example:
 * 
 * Using the ID of the record, which is the zone identifier, record name, and record type, separated by underscores (`_`):
 * 
 * ```sh
 *  $ pulumi import aws:route53/record:Record myrecord Z4KAPRWWNC7JR_dev.example.com_NS
 * ```
 *  If the record also contains a set identifier, append it:
 * 
 * ```sh
 *  $ pulumi import aws:route53/record:Record myrecord Z4KAPRWWNC7JR_dev.example.com_NS_dev
 * ```
 * 
 */
@ResourceType(type="aws:route53/record:Record")
public class Record extends com.pulumi.resources.CustomResource {
    /**
     * An alias block. Conflicts with `ttl` &amp; `records`.
     * Documented below.
     * 
     */
    @Export(name="aliases", refs={List.class,RecordAlias.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RecordAlias>> aliases;

    /**
     * @return An alias block. Conflicts with `ttl` &amp; `records`.
     * Documented below.
     * 
     */
    public Output<Optional<List<RecordAlias>>> aliases() {
        return Codegen.optional(this.aliases);
    }
    /**
     * Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
     * 
     * Exactly one of `records` or `alias` must be specified: this determines whether it&#39;s an alias record.
     * 
     */
    @Export(name="allowOverwrite", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowOverwrite;

    /**
     * @return Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
     * 
     * Exactly one of `records` or `alias` must be specified: this determines whether it&#39;s an alias record.
     * 
     */
    public Output<Boolean> allowOverwrite() {
        return this.allowOverwrite;
    }
    /**
     * A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
     * 
     */
    @Export(name="cidrRoutingPolicy", refs={RecordCidrRoutingPolicy.class}, tree="[0]")
    private Output</* @Nullable */ RecordCidrRoutingPolicy> cidrRoutingPolicy;

    /**
     * @return A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
     * 
     */
    public Output<Optional<RecordCidrRoutingPolicy>> cidrRoutingPolicy() {
        return Codegen.optional(this.cidrRoutingPolicy);
    }
    /**
     * A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
     * 
     */
    @Export(name="failoverRoutingPolicies", refs={List.class,RecordFailoverRoutingPolicy.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RecordFailoverRoutingPolicy>> failoverRoutingPolicies;

    /**
     * @return A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
     * 
     */
    public Output<Optional<List<RecordFailoverRoutingPolicy>>> failoverRoutingPolicies() {
        return Codegen.optional(this.failoverRoutingPolicies);
    }
    /**
     * [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
     * 
     */
    @Export(name="fqdn", refs={String.class}, tree="[0]")
    private Output<String> fqdn;

    /**
     * @return [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
     * 
     */
    public Output<String> fqdn() {
        return this.fqdn;
    }
    /**
     * A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
     * 
     */
    @Export(name="geolocationRoutingPolicies", refs={List.class,RecordGeolocationRoutingPolicy.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RecordGeolocationRoutingPolicy>> geolocationRoutingPolicies;

    /**
     * @return A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
     * 
     */
    public Output<Optional<List<RecordGeolocationRoutingPolicy>>> geolocationRoutingPolicies() {
        return Codegen.optional(this.geolocationRoutingPolicies);
    }
    /**
     * A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
     * 
     */
    @Export(name="geoproximityRoutingPolicy", refs={RecordGeoproximityRoutingPolicy.class}, tree="[0]")
    private Output</* @Nullable */ RecordGeoproximityRoutingPolicy> geoproximityRoutingPolicy;

    /**
     * @return A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
     * 
     */
    public Output<Optional<RecordGeoproximityRoutingPolicy>> geoproximityRoutingPolicy() {
        return Codegen.optional(this.geoproximityRoutingPolicy);
    }
    /**
     * The health check the record should be associated with.
     * 
     */
    @Export(name="healthCheckId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckId;

    /**
     * @return The health check the record should be associated with.
     * 
     */
    public Output<Optional<String>> healthCheckId() {
        return Codegen.optional(this.healthCheckId);
    }
    /**
     * A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
     * 
     */
    @Export(name="latencyRoutingPolicies", refs={List.class,RecordLatencyRoutingPolicy.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RecordLatencyRoutingPolicy>> latencyRoutingPolicies;

    /**
     * @return A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
     * 
     */
    public Output<Optional<List<RecordLatencyRoutingPolicy>>> latencyRoutingPolicies() {
        return Codegen.optional(this.latencyRoutingPolicies);
    }
    /**
     * Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
     * 
     */
    @Export(name="multivalueAnswerRoutingPolicy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> multivalueAnswerRoutingPolicy;

    /**
     * @return Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
     * 
     */
    public Output<Optional<Boolean>> multivalueAnswerRoutingPolicy() {
        return Codegen.optional(this.multivalueAnswerRoutingPolicy);
    }
    /**
     * The name of the record.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the record.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\&#34;\&#34;` inside the provider configuration string (e.g., `&#34;first255characters\&#34;\&#34;morecharacters&#34;`).
     * 
     */
    @Export(name="records", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> records;

    /**
     * @return A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\&#34;\&#34;` inside the provider configuration string (e.g., `&#34;first255characters\&#34;\&#34;morecharacters&#34;`).
     * 
     */
    public Output<Optional<List<String>>> records() {
        return Codegen.optional(this.records);
    }
    /**
     * Unique identifier to differentiate records with routing policies from one another. Required if using `cidr_routing_policy`, `failover_routing_policy`, `geolocation_routing_policy`,`geoproximity_routing_policy`, `latency_routing_policy`, `multivalue_answer_routing_policy`, or `weighted_routing_policy`.
     * 
     */
    @Export(name="setIdentifier", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> setIdentifier;

    /**
     * @return Unique identifier to differentiate records with routing policies from one another. Required if using `cidr_routing_policy`, `failover_routing_policy`, `geolocation_routing_policy`,`geoproximity_routing_policy`, `latency_routing_policy`, `multivalue_answer_routing_policy`, or `weighted_routing_policy`.
     * 
     */
    public Output<Optional<String>> setIdentifier() {
        return Codegen.optional(this.setIdentifier);
    }
    /**
     * The TTL of the record.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return The TTL of the record.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
     * 
     */
    @Export(name="weightedRoutingPolicies", refs={List.class,RecordWeightedRoutingPolicy.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RecordWeightedRoutingPolicy>> weightedRoutingPolicies;

    /**
     * @return A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
     * 
     */
    public Output<Optional<List<RecordWeightedRoutingPolicy>>> weightedRoutingPolicies() {
        return Codegen.optional(this.weightedRoutingPolicies);
    }
    /**
     * The ID of the hosted zone to contain this record.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The ID of the hosted zone to contain this record.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Record(String name) {
        this(name, RecordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Record(String name, RecordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Record(String name, RecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/record:Record", name, args == null ? RecordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Record(String name, Output<String> id, @Nullable RecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53/record:Record", name, state, makeResourceOptions(options, id));
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
    public static Record get(String name, Output<String> id, @Nullable RecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Record(name, id, state, options);
    }
}
