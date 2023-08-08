// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.SubnetCidrReservationArgs;
import com.pulumi.aws.ec2.inputs.SubnetCidrReservationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a subnet CIDR reservation resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.SubnetCidrReservation;
 * import com.pulumi.aws.ec2.SubnetCidrReservationArgs;
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
 *         var example = new SubnetCidrReservation(&#34;example&#34;, SubnetCidrReservationArgs.builder()        
 *             .cidrBlock(&#34;10.0.0.16/28&#34;)
 *             .reservationType(&#34;prefix&#34;)
 *             .subnetId(aws_subnet.example().id())
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
 *  to = aws_ec2_subnet_cidr_reservation.example
 * 
 *  id = &#34;subnet-01llsxvsxabqiymcz:scr-4mnvz6wb7otksjcs9&#34; } Using `pulumi import`, import Existing CIDR reservations using `SUBNET_ID:RESERVATION_ID`. For exampleconsole % pulumi import aws_ec2_subnet_cidr_reservation.example subnet-01llsxvsxabqiymcz:scr-4mnvz6wb7otksjcs9
 * 
 */
@ResourceType(type="aws:ec2/subnetCidrReservation:SubnetCidrReservation")
public class SubnetCidrReservation extends com.pulumi.resources.CustomResource {
    /**
     * The CIDR block for the reservation.
     * 
     */
    @Export(name="cidrBlock", refs={String.class}, tree="[0]")
    private Output<String> cidrBlock;

    /**
     * @return The CIDR block for the reservation.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * A brief description of the reservation.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A brief description of the reservation.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * ID of the AWS account that owns this CIDR reservation.
     * 
     */
    @Export(name="ownerId", refs={String.class}, tree="[0]")
    private Output<String> ownerId;

    /**
     * @return ID of the AWS account that owns this CIDR reservation.
     * 
     */
    public Output<String> ownerId() {
        return this.ownerId;
    }
    /**
     * The type of reservation to create. Valid values: `explicit`, `prefix`
     * 
     */
    @Export(name="reservationType", refs={String.class}, tree="[0]")
    private Output<String> reservationType;

    /**
     * @return The type of reservation to create. Valid values: `explicit`, `prefix`
     * 
     */
    public Output<String> reservationType() {
        return this.reservationType;
    }
    /**
     * The ID of the subnet to create the reservation for.
     * 
     */
    @Export(name="subnetId", refs={String.class}, tree="[0]")
    private Output<String> subnetId;

    /**
     * @return The ID of the subnet to create the reservation for.
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SubnetCidrReservation(String name) {
        this(name, SubnetCidrReservationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SubnetCidrReservation(String name, SubnetCidrReservationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SubnetCidrReservation(String name, SubnetCidrReservationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/subnetCidrReservation:SubnetCidrReservation", name, args == null ? SubnetCidrReservationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SubnetCidrReservation(String name, Output<String> id, @Nullable SubnetCidrReservationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/subnetCidrReservation:SubnetCidrReservation", name, state, makeResourceOptions(options, id));
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
    public static SubnetCidrReservation get(String name, Output<String> id, @Nullable SubnetCidrReservationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SubnetCidrReservation(name, id, state, options);
    }
}
