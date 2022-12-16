// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.VpcEndpointConnectionAccepterArgs;
import com.pulumi.aws.ec2.inputs.VpcEndpointConnectionAccepterState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource to accept a pending VPC Endpoint Connection accept request to VPC Endpoint Service.
 * 
 * ## Example Usage
 * ### Accept cross-account request
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.VpcEndpointService;
 * import com.pulumi.aws.ec2.VpcEndpointServiceArgs;
 * import com.pulumi.aws.ec2.VpcEndpoint;
 * import com.pulumi.aws.ec2.VpcEndpointArgs;
 * import com.pulumi.aws.ec2.VpcEndpointConnectionAccepter;
 * import com.pulumi.aws.ec2.VpcEndpointConnectionAccepterArgs;
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
 *         var exampleVpcEndpointService = new VpcEndpointService(&#34;exampleVpcEndpointService&#34;, VpcEndpointServiceArgs.builder()        
 *             .acceptanceRequired(false)
 *             .networkLoadBalancerArns(aws_lb.example().arn())
 *             .build());
 * 
 *         var exampleVpcEndpoint = new VpcEndpoint(&#34;exampleVpcEndpoint&#34;, VpcEndpointArgs.builder()        
 *             .vpcId(aws_vpc.test_alternate().id())
 *             .serviceName(aws_vpc_endpoint_service.test().service_name())
 *             .vpcEndpointType(&#34;Interface&#34;)
 *             .privateDnsEnabled(false)
 *             .securityGroupIds(aws_security_group.test().id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(&#34;aws.alternate&#34;)
 *                 .build());
 * 
 *         var exampleVpcEndpointConnectionAccepter = new VpcEndpointConnectionAccepter(&#34;exampleVpcEndpointConnectionAccepter&#34;, VpcEndpointConnectionAccepterArgs.builder()        
 *             .vpcEndpointServiceId(exampleVpcEndpointService.id())
 *             .vpcEndpointId(exampleVpcEndpoint.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Endpoint Services can be imported using ID of the connection, which is the `VPC Endpoint Service ID` and `VPC Endpoint ID` separated by underscore (`_`). e.g.
 * 
 * ```sh
 *  $ pulumi import aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter foo vpce-svc-0f97a19d3fa8220bc_vpce-010601a6db371e263
 * ```
 * 
 */
@ResourceType(type="aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter")
public class VpcEndpointConnectionAccepter extends com.pulumi.resources.CustomResource {
    /**
     * AWS VPC Endpoint ID.
     * 
     */
    @Export(name="vpcEndpointId", refs={String.class}, tree="[0]")
    private Output<String> vpcEndpointId;

    /**
     * @return AWS VPC Endpoint ID.
     * 
     */
    public Output<String> vpcEndpointId() {
        return this.vpcEndpointId;
    }
    /**
     * AWS VPC Endpoint Service ID.
     * 
     */
    @Export(name="vpcEndpointServiceId", refs={String.class}, tree="[0]")
    private Output<String> vpcEndpointServiceId;

    /**
     * @return AWS VPC Endpoint Service ID.
     * 
     */
    public Output<String> vpcEndpointServiceId() {
        return this.vpcEndpointServiceId;
    }
    /**
     * State of the VPC Endpoint.
     * 
     */
    @Export(name="vpcEndpointState", refs={String.class}, tree="[0]")
    private Output<String> vpcEndpointState;

    /**
     * @return State of the VPC Endpoint.
     * 
     */
    public Output<String> vpcEndpointState() {
        return this.vpcEndpointState;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpointConnectionAccepter(String name) {
        this(name, VpcEndpointConnectionAccepterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpointConnectionAccepter(String name, VpcEndpointConnectionAccepterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpointConnectionAccepter(String name, VpcEndpointConnectionAccepterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter", name, args == null ? VpcEndpointConnectionAccepterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcEndpointConnectionAccepter(String name, Output<String> id, @Nullable VpcEndpointConnectionAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter", name, state, makeResourceOptions(options, id));
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
    public static VpcEndpointConnectionAccepter get(String name, Output<String> id, @Nullable VpcEndpointConnectionAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpointConnectionAccepter(name, id, state, options);
    }
}
