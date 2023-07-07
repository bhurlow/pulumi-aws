// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcEndpointArgs;
import com.pulumi.aws.ec2.inputs.VpcEndpointState;
import com.pulumi.aws.ec2.outputs.VpcEndpointDnsEntry;
import com.pulumi.aws.ec2.outputs.VpcEndpointDnsOptions;
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
 * Provides a VPC Endpoint resource.
 * 
 * &gt; **NOTE on VPC Endpoints and VPC Endpoint Associations:** The provider provides both standalone VPC Endpoint Associations for
 * Route Tables - (an association between a VPC endpoint and a single `route_table_id`),
 * Security Groups - (an association between a VPC endpoint and a single `security_group_id`),
 * and Subnets - (an association between a VPC endpoint and a single `subnet_id`) and
 * a VPC Endpoint resource with `route_table_ids` and `subnet_ids` attributes.
 * Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
 * Doing so will cause a conflict of associations and will overwrite the association.
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcEndpoint;
 * import com.pulumi.aws.ec2.VpcEndpointArgs;
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
 *         var s3 = new VpcEndpoint(&#34;s3&#34;, VpcEndpointArgs.builder()        
 *             .vpcId(aws_vpc.main().id())
 *             .serviceName(&#34;com.amazonaws.us-west-2.s3&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Basic w/ Tags
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcEndpoint;
 * import com.pulumi.aws.ec2.VpcEndpointArgs;
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
 *         var s3 = new VpcEndpoint(&#34;s3&#34;, VpcEndpointArgs.builder()        
 *             .vpcId(aws_vpc.main().id())
 *             .serviceName(&#34;com.amazonaws.us-west-2.s3&#34;)
 *             .tags(Map.of(&#34;Environment&#34;, &#34;test&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Interface Endpoint Type
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcEndpoint;
 * import com.pulumi.aws.ec2.VpcEndpointArgs;
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
 *         var ec2 = new VpcEndpoint(&#34;ec2&#34;, VpcEndpointArgs.builder()        
 *             .vpcId(aws_vpc.main().id())
 *             .serviceName(&#34;com.amazonaws.us-west-2.ec2&#34;)
 *             .vpcEndpointType(&#34;Interface&#34;)
 *             .securityGroupIds(aws_security_group.sg1().id())
 *             .privateDnsEnabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gateway Load Balancer Endpoint Type
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetCallerIdentityArgs;
 * import com.pulumi.aws.ec2.VpcEndpointService;
 * import com.pulumi.aws.ec2.VpcEndpointServiceArgs;
 * import com.pulumi.aws.ec2.VpcEndpoint;
 * import com.pulumi.aws.ec2.VpcEndpointArgs;
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
 *         final var current = AwsFunctions.getCallerIdentity();
 * 
 *         var exampleVpcEndpointService = new VpcEndpointService(&#34;exampleVpcEndpointService&#34;, VpcEndpointServiceArgs.builder()        
 *             .acceptanceRequired(false)
 *             .allowedPrincipals(current.applyValue(getCallerIdentityResult -&gt; getCallerIdentityResult.arn()))
 *             .gatewayLoadBalancerArns(aws_lb.example().arn())
 *             .build());
 * 
 *         var exampleVpcEndpoint = new VpcEndpoint(&#34;exampleVpcEndpoint&#34;, VpcEndpointArgs.builder()        
 *             .serviceName(exampleVpcEndpointService.serviceName())
 *             .subnetIds(aws_subnet.example().id())
 *             .vpcEndpointType(exampleVpcEndpointService.serviceType())
 *             .vpcId(aws_vpc.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Endpoints can be imported using the `vpc endpoint id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/vpcEndpoint:VpcEndpoint endpoint1 vpce-3ecf2a57
 * ```
 * 
 */
@ResourceType(type="aws:ec2/vpcEndpoint:VpcEndpoint")
public class VpcEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the VPC endpoint.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the VPC endpoint.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
     * 
     */
    @Export(name="autoAccept", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoAccept;

    /**
     * @return Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
     * 
     */
    public Output<Optional<Boolean>> autoAccept() {
        return Codegen.optional(this.autoAccept);
    }
    /**
     * The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
     * 
     */
    @Export(name="cidrBlocks", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> cidrBlocks;

    /**
     * @return The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
     * 
     */
    public Output<List<String>> cidrBlocks() {
        return this.cidrBlocks;
    }
    /**
     * The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
     * 
     */
    @Export(name="dnsEntries", refs={List.class,VpcEndpointDnsEntry.class}, tree="[0,1]")
    private Output<List<VpcEndpointDnsEntry>> dnsEntries;

    /**
     * @return The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
     * 
     */
    public Output<List<VpcEndpointDnsEntry>> dnsEntries() {
        return this.dnsEntries;
    }
    /**
     * The DNS options for the endpoint. See dns_options below.
     * 
     */
    @Export(name="dnsOptions", refs={VpcEndpointDnsOptions.class}, tree="[0]")
    private Output<VpcEndpointDnsOptions> dnsOptions;

    /**
     * @return The DNS options for the endpoint. See dns_options below.
     * 
     */
    public Output<VpcEndpointDnsOptions> dnsOptions() {
        return this.dnsOptions;
    }
    /**
     * The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
     * 
     */
    @Export(name="ipAddressType", refs={String.class}, tree="[0]")
    private Output<String> ipAddressType;

    /**
     * @return The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
     * 
     */
    public Output<String> ipAddressType() {
        return this.ipAddressType;
    }
    /**
     * One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
     * 
     */
    @Export(name="networkInterfaceIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> networkInterfaceIds;

    /**
     * @return One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
     * 
     */
    public Output<List<String>> networkInterfaceIds() {
        return this.networkInterfaceIds;
    }
    /**
     * The ID of the AWS account that owns the VPC endpoint.
     * 
     */
    @Export(name="ownerId", refs={String.class}, tree="[0]")
    private Output<String> ownerId;

    /**
     * @return The ID of the AWS account that owns the VPC endpoint.
     * 
     */
    public Output<String> ownerId() {
        return this.ownerId;
    }
    /**
     * A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
     * 
     */
    @Export(name="prefixListId", refs={String.class}, tree="[0]")
    private Output<String> prefixListId;

    /**
     * @return The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
     * 
     */
    public Output<String> prefixListId() {
        return this.prefixListId;
    }
    /**
     * Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
     * Defaults to `false`.
     * 
     */
    @Export(name="privateDnsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> privateDnsEnabled;

    /**
     * @return Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
     * Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> privateDnsEnabled() {
        return Codegen.optional(this.privateDnsEnabled);
    }
    /**
     * Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
     * 
     */
    @Export(name="requesterManaged", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> requesterManaged;

    /**
     * @return Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
     * 
     */
    public Output<Boolean> requesterManaged() {
        return this.requesterManaged;
    }
    /**
     * One or more route table IDs. Applicable for endpoints of type `Gateway`.
     * 
     */
    @Export(name="routeTableIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> routeTableIds;

    /**
     * @return One or more route table IDs. Applicable for endpoints of type `Gateway`.
     * 
     */
    public Output<List<String>> routeTableIds() {
        return this.routeTableIds;
    }
    /**
     * The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
     * If no security groups are specified, the VPC&#39;s [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroupIds;

    /**
     * @return The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
     * If no security groups are specified, the VPC&#39;s [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * The service name. For AWS services the service name is usually in the form `com.amazonaws.&lt;region&gt;.&lt;service&gt;` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.&lt;region&gt;.notebook`).
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service name. For AWS services the service name is usually in the form `com.amazonaws.&lt;region&gt;.&lt;service&gt;` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.&lt;region&gt;.notebook`).
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * The state of the VPC endpoint.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The state of the VPC endpoint.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
     * 
     */
    @Export(name="subnetIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> subnetIds;

    /**
     * @return The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
     * 
     */
    public Output<List<String>> subnetIds() {
        return this.subnetIds;
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
     * 
     */
    @Export(name="vpcEndpointType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcEndpointType;

    /**
     * @return The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
     * 
     */
    public Output<Optional<String>> vpcEndpointType() {
        return Codegen.optional(this.vpcEndpointType);
    }
    /**
     * The ID of the VPC in which the endpoint will be used.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC in which the endpoint will be used.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpoint(String name) {
        this(name, VpcEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpoint(String name, VpcEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpoint(String name, VpcEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcEndpoint:VpcEndpoint", name, args == null ? VpcEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcEndpoint(String name, Output<String> id, @Nullable VpcEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcEndpoint:VpcEndpoint", name, state, makeResourceOptions(options, id));
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
    public static VpcEndpoint get(String name, Output<String> id, @Nullable VpcEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpoint(name, id, state, options);
    }
}
