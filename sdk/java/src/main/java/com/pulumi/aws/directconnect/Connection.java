// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.directconnect;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.directconnect.ConnectionArgs;
import com.pulumi.aws.directconnect.inputs.ConnectionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Connection of Direct Connect.
 * 
 * ## Example Usage
 * ### Create a connection
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.directconnect.Connection;
 * import com.pulumi.aws.directconnect.ConnectionArgs;
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
 *         var hoge = new Connection(&#34;hoge&#34;, ConnectionArgs.builder()        
 *             .bandwidth(&#34;1Gbps&#34;)
 *             .location(&#34;EqDC2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Request a MACsec-capable connection
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.directconnect.Connection;
 * import com.pulumi.aws.directconnect.ConnectionArgs;
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
 *         var example = new Connection(&#34;example&#34;, ConnectionArgs.builder()        
 *             .bandwidth(&#34;10Gbps&#34;)
 *             .location(&#34;EqDA2&#34;)
 *             .requestMacsec(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Configure encryption mode for MACsec-capable connections
 * 
 * &gt; **NOTE:** You can only specify the `encryption_mode` argument once the connection is in an `Available` state.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.directconnect.Connection;
 * import com.pulumi.aws.directconnect.ConnectionArgs;
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
 *         var example = new Connection(&#34;example&#34;, ConnectionArgs.builder()        
 *             .bandwidth(&#34;10Gbps&#34;)
 *             .encryptionMode(&#34;must_encrypt&#34;)
 *             .location(&#34;EqDC2&#34;)
 *             .requestMacsec(true)
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
 *  to = aws_dx_connection.test_connection
 * 
 *  id = &#34;dxcon-ffre0ec3&#34; } Using `pulumi import`, import Direct Connect connections using the connection `id`. For exampleconsole % pulumi import aws_dx_connection.test_connection dxcon-ffre0ec3
 * 
 */
@ResourceType(type="aws:directconnect/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the connection.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the connection.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The Direct Connect endpoint on which the physical connection terminates.
     * 
     */
    @Export(name="awsDevice", refs={String.class}, tree="[0]")
    private Output<String> awsDevice;

    /**
     * @return The Direct Connect endpoint on which the physical connection terminates.
     * 
     */
    public Output<String> awsDevice() {
        return this.awsDevice;
    }
    /**
     * The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
     * 
     */
    @Export(name="bandwidth", refs={String.class}, tree="[0]")
    private Output<String> bandwidth;

    /**
     * @return The bandwidth of the connection. Valid values for dedicated connections: 1Gbps, 10Gbps. Valid values for hosted connections: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
     * 
     */
    public Output<String> bandwidth() {
        return this.bandwidth;
    }
    /**
     * The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `no_encrypt`, `should_encrypt`, and `must_encrypt`.
     * 
     */
    @Export(name="encryptionMode", refs={String.class}, tree="[0]")
    private Output<String> encryptionMode;

    /**
     * @return The connection MAC Security (MACsec) encryption mode. MAC Security (MACsec) is only available on dedicated connections. Valid values are `no_encrypt`, `should_encrypt`, and `must_encrypt`.
     * 
     */
    public Output<String> encryptionMode() {
        return this.encryptionMode;
    }
    /**
     * Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
     * 
     */
    @Export(name="hasLogicalRedundancy", refs={String.class}, tree="[0]")
    private Output<String> hasLogicalRedundancy;

    /**
     * @return Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6).
     * 
     */
    public Output<String> hasLogicalRedundancy() {
        return this.hasLogicalRedundancy;
    }
    /**
     * Boolean value representing if jumbo frames have been enabled for this connection.
     * 
     */
    @Export(name="jumboFrameCapable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> jumboFrameCapable;

    /**
     * @return Boolean value representing if jumbo frames have been enabled for this connection.
     * 
     */
    public Output<Boolean> jumboFrameCapable() {
        return this.jumboFrameCapable;
    }
    /**
     * The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Boolean value indicating whether the connection supports MAC Security (MACsec).
     * 
     */
    @Export(name="macsecCapable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> macsecCapable;

    /**
     * @return Boolean value indicating whether the connection supports MAC Security (MACsec).
     * 
     */
    public Output<Boolean> macsecCapable() {
        return this.macsecCapable;
    }
    /**
     * The name of the connection.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the connection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the AWS account that owns the connection.
     * 
     */
    @Export(name="ownerAccountId", refs={String.class}, tree="[0]")
    private Output<String> ownerAccountId;

    /**
     * @return The ID of the AWS account that owns the connection.
     * 
     */
    public Output<String> ownerAccountId() {
        return this.ownerAccountId;
    }
    /**
     * The name of the AWS Direct Connect service provider associated with the connection.
     * 
     */
    @Export(name="partnerName", refs={String.class}, tree="[0]")
    private Output<String> partnerName;

    /**
     * @return The name of the AWS Direct Connect service provider associated with the connection.
     * 
     */
    public Output<String> partnerName() {
        return this.partnerName;
    }
    /**
     * The MAC Security (MACsec) port link status of the connection.
     * 
     */
    @Export(name="portEncryptionStatus", refs={String.class}, tree="[0]")
    private Output<String> portEncryptionStatus;

    /**
     * @return The MAC Security (MACsec) port link status of the connection.
     * 
     */
    public Output<String> portEncryptionStatus() {
        return this.portEncryptionStatus;
    }
    /**
     * The name of the service provider associated with the connection.
     * 
     */
    @Export(name="providerName", refs={String.class}, tree="[0]")
    private Output<String> providerName;

    /**
     * @return The name of the service provider associated with the connection.
     * 
     */
    public Output<String> providerName() {
        return this.providerName;
    }
    /**
     * Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
     * 
     * &gt; **NOTE:** Changing the value of `request_macsec` will cause the resource to be destroyed and re-created.
     * 
     */
    @Export(name="requestMacsec", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> requestMacsec;

    /**
     * @return Boolean value indicating whether you want the connection to support MAC Security (MACsec). MAC Security (MACsec) is only available on dedicated connections. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for more information about MAC Security (MACsec) prerequisites. Default value: `false`.
     * 
     * &gt; **NOTE:** Changing the value of `request_macsec` will cause the resource to be destroyed and re-created.
     * 
     */
    public Output<Optional<Boolean>> requestMacsec() {
        return Codegen.optional(this.requestMacsec);
    }
    /**
     * Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
     * 
     */
    @Export(name="skipDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipDestroy;

    /**
     * @return Set to true if you do not wish the connection to be deleted at destroy time, and instead just removed from the state.
     * 
     */
    public Output<Optional<Boolean>> skipDestroy() {
        return Codegen.optional(this.skipDestroy);
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
     * The VLAN ID.
     * 
     */
    @Export(name="vlanId", refs={Integer.class}, tree="[0]")
    private Output<Integer> vlanId;

    /**
     * @return The VLAN ID.
     * 
     */
    public Output<Integer> vlanId() {
        return this.vlanId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:directconnect/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:directconnect/connection:Connection", name, state, makeResourceOptions(options, id));
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
    public static Connection get(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connection(name, id, state, options);
    }
}
