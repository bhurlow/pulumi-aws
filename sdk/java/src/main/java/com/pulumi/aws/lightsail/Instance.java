// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lightsail;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.inputs.InstanceState;
import com.pulumi.aws.lightsail.outputs.InstanceAddOn;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Lightsail Instance. Amazon Lightsail is a service to provide easy virtual private servers
 * with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
 * for more information.
 * 
 * &gt; **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see [&#34;Regions and Availability Zones in Amazon Lightsail&#34;](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.Instance;
 * import com.pulumi.aws.lightsail.InstanceArgs;
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
 *         var gitlabTest = new Instance(&#34;gitlabTest&#34;, InstanceArgs.builder()        
 *             .availabilityZone(&#34;us-east-1b&#34;)
 *             .blueprintId(&#34;amazon_linux_2&#34;)
 *             .bundleId(&#34;nano_1_0&#34;)
 *             .keyPairName(&#34;some_key_name&#34;)
 *             .tags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Example With User Data
 * 
 * Lightsail user data is handled differently than ec2 user data. Lightsail user data only accepts a single lined string. The below example shows installing apache and creating the index page.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.Instance;
 * import com.pulumi.aws.lightsail.InstanceArgs;
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
 *         var custom = new Instance(&#34;custom&#34;, InstanceArgs.builder()        
 *             .availabilityZone(&#34;us-east-1b&#34;)
 *             .blueprintId(&#34;amazon_linux_2&#34;)
 *             .bundleId(&#34;nano_1_0&#34;)
 *             .userData(&#34;sudo yum install -y httpd &amp;&amp; sudo systemctl start httpd &amp;&amp; sudo systemctl enable httpd &amp;&amp; echo &#39;&lt;h1&gt;Deployed via Pulumi&lt;/h1&gt;&#39; | sudo tee /var/www/html/index.html&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Enable Auto Snapshots
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lightsail.Instance;
 * import com.pulumi.aws.lightsail.InstanceArgs;
 * import com.pulumi.aws.lightsail.inputs.InstanceAddOnArgs;
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
 *         var test = new Instance(&#34;test&#34;, InstanceArgs.builder()        
 *             .addOn(InstanceAddOnArgs.builder()
 *                 .snapshotTime(&#34;06:00&#34;)
 *                 .status(&#34;Enabled&#34;)
 *                 .type(&#34;AutoSnapshot&#34;)
 *                 .build())
 *             .availabilityZone(&#34;us-east-1b&#34;)
 *             .blueprintId(&#34;amazon_linux_2&#34;)
 *             .bundleId(&#34;nano_1_0&#34;)
 *             .tags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Availability Zones
 * 
 * Lightsail currently supports the following Availability Zones (e.g., `us-east-1a`):
 * 
 * - `ap-northeast-1{a,c,d}`
 * - `ap-northeast-2{a,c}`
 * - `ap-south-1{a,b}`
 * - `ap-southeast-1{a,b,c}`
 * - `ap-southeast-2{a,b,c}`
 * - `ca-central-1{a,b}`
 * - `eu-central-1{a,b,c}`
 * - `eu-west-1{a,b,c}`
 * - `eu-west-2{a,b,c}`
 * - `eu-west-3{a,b,c}`
 * - `us-east-1{a,b,c,d,e,f}`
 * - `us-east-2{a,b,c}`
 * - `us-west-2{a,b,c}`
 * 
 * ## Bundles
 * 
 * Lightsail currently supports the following Bundle IDs (e.g., an instance in `ap-northeast-1` would use `small_2_0`):
 * 
 * ### Prefix
 * 
 * A Bundle ID starts with one of the below size prefixes:
 * 
 * - `nano_`
 * - `micro_`
 * - `small_`
 * - `medium_`
 * - `large_`
 * - `xlarge_`
 * - `2xlarge_`
 * 
 * ### Suffix
 * 
 * A Bundle ID ends with one of the following suffixes depending on Availability Zone:
 * 
 * - ap-northeast-1: `2_0`
 * - ap-northeast-2: `2_0`
 * - ap-south-1: `2_1`
 * - ap-southeast-1: `2_0`
 * - ap-southeast-2: `2_2`
 * - ca-central-1: `2_0`
 * - eu-central-1: `2_0`
 * - eu-west-1: `2_0`
 * - eu-west-2: `2_0`
 * - eu-west-3: `2_0`
 * - us-east-1: `2_0`
 * - us-east-2: `2_0`
 * - us-west-2: `2_0`
 * 
 * ## Import
 * 
 * Using `pulumi import`, import Lightsail Instances using their name. For example:
 * 
 * ```sh
 *  $ pulumi import aws:lightsail/instance:Instance gitlab_test &#39;custom_gitlab&#39;
 * ```
 * 
 */
@ResourceType(type="aws:lightsail/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * The add on configuration for the instance. Detailed below.
     * 
     */
    @Export(name="addOn", refs={InstanceAddOn.class}, tree="[0]")
    private Output</* @Nullable */ InstanceAddOn> addOn;

    /**
     * @return The add on configuration for the instance. Detailed below.
     * 
     */
    public Output<Optional<InstanceAddOn>> addOn() {
        return Codegen.optional(this.addOn);
    }
    /**
     * The ARN of the Lightsail instance (matches `id`).
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the Lightsail instance (matches `id`).
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The Availability Zone in which to create your
     * instance (see list below)
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return The Availability Zone in which to create your
     * instance (see list below)
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: `aws lightsail get-blueprints`
     * 
     */
    @Export(name="blueprintId", refs={String.class}, tree="[0]")
    private Output<String> blueprintId;

    /**
     * @return The ID for a virtual private server image. A list of available blueprint IDs can be obtained using the AWS CLI command: `aws lightsail get-blueprints`
     * 
     */
    public Output<String> blueprintId() {
        return this.blueprintId;
    }
    /**
     * The bundle of specification information (see list below)
     * 
     */
    @Export(name="bundleId", refs={String.class}, tree="[0]")
    private Output<String> bundleId;

    /**
     * @return The bundle of specification information (see list below)
     * 
     */
    public Output<String> bundleId() {
        return this.bundleId;
    }
    /**
     * The number of vCPUs the instance has.
     * 
     */
    @Export(name="cpuCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> cpuCount;

    /**
     * @return The number of vCPUs the instance has.
     * 
     */
    public Output<Integer> cpuCount() {
        return this.cpuCount;
    }
    /**
     * The timestamp when the instance was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The timestamp when the instance was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
     * 
     */
    @Export(name="ipAddressType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipAddressType;

    /**
     * @return The IP address type of the Lightsail Instance. Valid Values: `dualstack` | `ipv4`.
     * 
     */
    public Output<Optional<String>> ipAddressType() {
        return Codegen.optional(this.ipAddressType);
    }
    /**
     * List of IPv6 addresses for the Lightsail instance.
     * 
     */
    @Export(name="ipv6Addresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> ipv6Addresses;

    /**
     * @return List of IPv6 addresses for the Lightsail instance.
     * 
     */
    public Output<List<String>> ipv6Addresses() {
        return this.ipv6Addresses;
    }
    /**
     * A Boolean value indicating whether this instance has a static IP assigned to it.
     * 
     */
    @Export(name="isStaticIp", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isStaticIp;

    /**
     * @return A Boolean value indicating whether this instance has a static IP assigned to it.
     * 
     */
    public Output<Boolean> isStaticIp() {
        return this.isStaticIp;
    }
    /**
     * The name of your key pair. Created in the
     * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
     * 
     */
    @Export(name="keyPairName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyPairName;

    /**
     * @return The name of your key pair. Created in the
     * Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
     * 
     */
    public Output<Optional<String>> keyPairName() {
        return Codegen.optional(this.keyPairName);
    }
    /**
     * The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The private IP address of the instance.
     * 
     */
    @Export(name="privateIpAddress", refs={String.class}, tree="[0]")
    private Output<String> privateIpAddress;

    /**
     * @return The private IP address of the instance.
     * 
     */
    public Output<String> privateIpAddress() {
        return this.privateIpAddress;
    }
    /**
     * The public IP address of the instance.
     * 
     */
    @Export(name="publicIpAddress", refs={String.class}, tree="[0]")
    private Output<String> publicIpAddress;

    /**
     * @return The public IP address of the instance.
     * 
     */
    public Output<String> publicIpAddress() {
        return this.publicIpAddress;
    }
    /**
     * The amount of RAM in GB on the instance (e.g., 1.0).
     * 
     */
    @Export(name="ramSize", refs={Double.class}, tree="[0]")
    private Output<Double> ramSize;

    /**
     * @return The amount of RAM in GB on the instance (e.g., 1.0).
     * 
     */
    public Output<Double> ramSize() {
        return this.ramSize;
    }
    /**
     * A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Single lined launch script as a string to configure server with additional user data
     * 
     */
    @Export(name="userData", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userData;

    /**
     * @return Single lined launch script as a string to configure server with additional user data
     * 
     */
    public Output<Optional<String>> userData() {
        return Codegen.optional(this.userData);
    }
    /**
     * The user name for connecting to the instance (e.g., ec2-user).
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The user name for connecting to the instance (e.g., ec2-user).
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lightsail/instance:Instance", name, state, makeResourceOptions(options, id));
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
