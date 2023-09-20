// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.datasync;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.datasync.LocationObjectStorageArgs;
import com.pulumi.aws.datasync.inputs.LocationObjectStorageState;
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
 * Manages a Object Storage Location within AWS DataSync.
 * 
 * &gt; **NOTE:** The DataSync Agents must be available before creating this resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.datasync.LocationObjectStorage;
 * import com.pulumi.aws.datasync.LocationObjectStorageArgs;
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
 *         var example = new LocationObjectStorage(&#34;example&#34;, LocationObjectStorageArgs.builder()        
 *             .agentArns(aws_datasync_agent.example().arn())
 *             .serverHostname(&#34;example&#34;)
 *             .bucketName(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import `aws_datasync_location_object_storage` using the Amazon Resource Name (ARN). For example:
 * 
 * ```sh
 *  $ pulumi import aws:datasync/locationObjectStorage:LocationObjectStorage example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
 * ```
 * 
 */
@ResourceType(type="aws:datasync/locationObjectStorage:LocationObjectStorage")
public class LocationObjectStorage extends com.pulumi.resources.CustomResource {
    /**
     * The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `access_key` and `secret_key` to provide the user name and password, respectively.
     * 
     */
    @Export(name="accessKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessKey;

    /**
     * @return The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `access_key` and `secret_key` to provide the user name and password, respectively.
     * 
     */
    public Output<Optional<String>> accessKey() {
        return Codegen.optional(this.accessKey);
    }
    /**
     * A list of DataSync Agent ARNs with which this location will be associated.
     * 
     */
    @Export(name="agentArns", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> agentArns;

    /**
     * @return A list of DataSync Agent ARNs with which this location will be associated.
     * 
     */
    public Output<List<String>> agentArns() {
        return this.agentArns;
    }
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the DataSync Location.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The bucket on the self-managed object storage server that is used to read data from.
     * 
     */
    @Export(name="bucketName", refs={String.class}, tree="[0]")
    private Output<String> bucketName;

    /**
     * @return The bucket on the self-managed object storage server that is used to read data from.
     * 
     */
    public Output<String> bucketName() {
        return this.bucketName;
    }
    /**
     * The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `access_key` and `secret_key` to provide the user name and password, respectively.
     * 
     */
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secretKey;

    /**
     * @return The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `access_key` and `secret_key` to provide the user name and password, respectively.
     * 
     */
    public Output<Optional<String>> secretKey() {
        return Codegen.optional(this.secretKey);
    }
    /**
     * Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem string. The certificate can be up to 32768 bytes (before Base64 encoding).
     * 
     */
    @Export(name="serverCertificate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverCertificate;

    /**
     * @return Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem string. The certificate can be up to 32768 bytes (before Base64 encoding).
     * 
     */
    public Output<Optional<String>> serverCertificate() {
        return Codegen.optional(this.serverCertificate);
    }
    /**
     * The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
     * 
     */
    @Export(name="serverHostname", refs={String.class}, tree="[0]")
    private Output<String> serverHostname;

    /**
     * @return The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
     * 
     */
    public Output<String> serverHostname() {
        return this.serverHostname;
    }
    /**
     * The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
     * 
     */
    @Export(name="serverPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> serverPort;

    /**
     * @return The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
     * 
     */
    public Output<Optional<Integer>> serverPort() {
        return Codegen.optional(this.serverPort);
    }
    /**
     * The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
     * 
     */
    @Export(name="serverProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverProtocol;

    /**
     * @return The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
     * 
     */
    public Output<Optional<String>> serverProtocol() {
        return Codegen.optional(this.serverProtocol);
    }
    /**
     * A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn&#39;t specified, it will default to /.
     * 
     */
    @Export(name="subdirectory", refs={String.class}, tree="[0]")
    private Output<String> subdirectory;

    /**
     * @return A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn&#39;t specified, it will default to /.
     * 
     */
    public Output<String> subdirectory() {
        return this.subdirectory;
    }
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The URL of the Object Storage location that was described.
     * 
     */
    @Export(name="uri", refs={String.class}, tree="[0]")
    private Output<String> uri;

    /**
     * @return The URL of the Object Storage location that was described.
     * 
     */
    public Output<String> uri() {
        return this.uri;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LocationObjectStorage(String name) {
        this(name, LocationObjectStorageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LocationObjectStorage(String name, LocationObjectStorageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LocationObjectStorage(String name, LocationObjectStorageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:datasync/locationObjectStorage:LocationObjectStorage", name, args == null ? LocationObjectStorageArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LocationObjectStorage(String name, Output<String> id, @Nullable LocationObjectStorageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:datasync/locationObjectStorage:LocationObjectStorage", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretKey",
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
    public static LocationObjectStorage get(String name, Output<String> id, @Nullable LocationObjectStorageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LocationObjectStorage(name, id, state, options);
    }
}
