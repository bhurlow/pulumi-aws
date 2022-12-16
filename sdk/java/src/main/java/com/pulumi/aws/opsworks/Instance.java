// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opsworks;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.opsworks.InstanceArgs;
import com.pulumi.aws.opsworks.inputs.InstanceState;
import com.pulumi.aws.opsworks.outputs.InstanceEbsBlockDevice;
import com.pulumi.aws.opsworks.outputs.InstanceEphemeralBlockDevice;
import com.pulumi.aws.opsworks.outputs.InstanceRootBlockDevice;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an OpsWorks instance resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.opsworks.Instance;
 * import com.pulumi.aws.opsworks.InstanceArgs;
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
 *         var my_instance = new Instance(&#34;my-instance&#34;, InstanceArgs.builder()        
 *             .stackId(aws_opsworks_stack.main().id())
 *             .layerIds(aws_opsworks_custom_layer.my-layer().id())
 *             .instanceType(&#34;t2.micro&#34;)
 *             .os(&#34;Amazon Linux 2015.09&#34;)
 *             .state(&#34;stopped&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Block devices
 * 
 * Each of the `*_block_device` attributes controls a portion of the AWS
 * Instance&#39;s &#34;Block Device Mapping&#34;. It&#39;s a good idea to familiarize yourself with [AWS&#39;s Block Device
 * Mapping docs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
 * to understand the implications of using these attributes.
 * 
 * ### `ebs_block_device`
 * 
 * * `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination. Default is `true`.
 * * `device_name` - (Required) Name of the device to mount.
 * * `iops` - (Optional) Amount of provisioned [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `io1`.
 * * `snapshot_id` - (Optional) Snapshot ID to mount.
 * * `volume_size` - (Optional) Size of the volume in gigabytes.
 * * `volume_type` - (Optional) Type of volume. Valid values are `standard`, `gp2`, or `io1`. Default is `standard`.
 * 
 * Modifying any `ebs_block_device` currently requires resource replacement.
 * 
 * ### `ephemeral_block_device`
 * 
 * * `device_name` - Name of the block device to mount on the instance.
 * * `virtual_name` - The [Instance Store Device Name](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames) (e.g., `ephemeral0`).
 * 
 * Each AWS Instance type has a different set of Instance Store block devices
 * available for attachment. AWS [publishes a
 * list](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
 * of which ephemeral devices are available on each type. The devices are always
 * identified by the `virtual_name` in the format `ephemeral{0..N}`.
 * 
 * ### `root_block_device`
 * 
 * * `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination. Default is `true`.
 * * `iops` - (Optional) Amount of provisioned [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `io1`.
 * * `volume_size` - (Optional) Size of the volume in gigabytes.
 * * `volume_type` - (Optional) Type of volume. Valid values are `standard`, `gp2`, or `io1`. Default is `standard`.
 * 
 * Modifying any of the `root_block_device` settings requires resource
 * replacement.
 * 
 * &gt; **NOTE:** Currently, changes to `*_block_device` configuration of _existing_
 * resources cannot be automatically detected by this provider. After making updates
 * to block device configuration, resource recreation can be manually triggered by
 * using the [`up` command with the --replace argument](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
 * 
 * ## Import
 * 
 * Opsworks Instances can be imported using the `instance id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:opsworks/instance:Instance my_instance 4d6d1710-ded9-42a1-b08e-b043ad7af1e2
 * ```
 * 
 */
@ResourceType(type="aws:opsworks/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * OpsWorks agent to install. Default is `INHERIT`.
     * 
     */
    @Export(name="agentVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> agentVersion;

    /**
     * @return OpsWorks agent to install. Default is `INHERIT`.
     * 
     */
    public Output<Optional<String>> agentVersion() {
        return Codegen.optional(this.agentVersion);
    }
    /**
     * AMI to use for the instance.  If an AMI is specified, `os` must be `Custom`.
     * 
     */
    @Export(name="amiId", refs={String.class}, tree="[0]")
    private Output<String> amiId;

    /**
     * @return AMI to use for the instance.  If an AMI is specified, `os` must be `Custom`.
     * 
     */
    public Output<String> amiId() {
        return this.amiId;
    }
    /**
     * Machine architecture for created instances.  Valid values are `x86_64` or `i386`. The default is `x86_64`.
     * 
     */
    @Export(name="architecture", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> architecture;

    /**
     * @return Machine architecture for created instances.  Valid values are `x86_64` or `i386`. The default is `x86_64`.
     * 
     */
    public Output<Optional<String>> architecture() {
        return Codegen.optional(this.architecture);
    }
    /**
     * Creates load-based or time-based instances.  Valid values are `load`, `timer`.
     * 
     */
    @Export(name="autoScalingType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> autoScalingType;

    /**
     * @return Creates load-based or time-based instances.  Valid values are `load`, `timer`.
     * 
     */
    public Output<Optional<String>> autoScalingType() {
        return Codegen.optional(this.autoScalingType);
    }
    /**
     * Name of the availability zone where instances will be created by default.
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return Name of the availability zone where instances will be created by default.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Time that the instance was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Time that the instance was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Whether to delete EBS volume on deletion. Default is `true`.
     * 
     */
    @Export(name="deleteEbs", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteEbs;

    /**
     * @return Whether to delete EBS volume on deletion. Default is `true`.
     * 
     */
    public Output<Optional<Boolean>> deleteEbs() {
        return Codegen.optional(this.deleteEbs);
    }
    /**
     * Whether to delete the Elastic IP on deletion.
     * 
     */
    @Export(name="deleteEip", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteEip;

    /**
     * @return Whether to delete the Elastic IP on deletion.
     * 
     */
    public Output<Optional<Boolean>> deleteEip() {
        return Codegen.optional(this.deleteEip);
    }
    /**
     * Configuration block for additional EBS block devices to attach to the instance. See Block Devices below.
     * 
     */
    @Export(name="ebsBlockDevices", refs={List.class,InstanceEbsBlockDevice.class}, tree="[0,1]")
    private Output<List<InstanceEbsBlockDevice>> ebsBlockDevices;

    /**
     * @return Configuration block for additional EBS block devices to attach to the instance. See Block Devices below.
     * 
     */
    public Output<List<InstanceEbsBlockDevice>> ebsBlockDevices() {
        return this.ebsBlockDevices;
    }
    /**
     * Whether the launched EC2 instance will be EBS-optimized.
     * 
     */
    @Export(name="ebsOptimized", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ebsOptimized;

    /**
     * @return Whether the launched EC2 instance will be EBS-optimized.
     * 
     */
    public Output<Optional<Boolean>> ebsOptimized() {
        return Codegen.optional(this.ebsOptimized);
    }
    /**
     * EC2 instance ID.
     * 
     */
    @Export(name="ec2InstanceId", refs={String.class}, tree="[0]")
    private Output<String> ec2InstanceId;

    /**
     * @return EC2 instance ID.
     * 
     */
    public Output<String> ec2InstanceId() {
        return this.ec2InstanceId;
    }
    /**
     * ECS cluster&#39;s ARN for container instances.
     * 
     */
    @Export(name="ecsClusterArn", refs={String.class}, tree="[0]")
    private Output<String> ecsClusterArn;

    /**
     * @return ECS cluster&#39;s ARN for container instances.
     * 
     */
    public Output<String> ecsClusterArn() {
        return this.ecsClusterArn;
    }
    /**
     * Instance Elastic IP address.
     * 
     */
    @Export(name="elasticIp", refs={String.class}, tree="[0]")
    private Output<String> elasticIp;

    /**
     * @return Instance Elastic IP address.
     * 
     */
    public Output<String> elasticIp() {
        return this.elasticIp;
    }
    /**
     * Configuration block for ephemeral (also known as &#34;Instance Store&#34;) volumes on the instance. See Block Devices below.
     * 
     */
    @Export(name="ephemeralBlockDevices", refs={List.class,InstanceEphemeralBlockDevice.class}, tree="[0,1]")
    private Output<List<InstanceEphemeralBlockDevice>> ephemeralBlockDevices;

    /**
     * @return Configuration block for ephemeral (also known as &#34;Instance Store&#34;) volumes on the instance. See Block Devices below.
     * 
     */
    public Output<List<InstanceEphemeralBlockDevice>> ephemeralBlockDevices() {
        return this.ephemeralBlockDevices;
    }
    /**
     * Instance&#39;s host name.
     * 
     */
    @Export(name="hostname", refs={String.class}, tree="[0]")
    private Output<String> hostname;

    /**
     * @return Instance&#39;s host name.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * For registered instances, infrastructure class: ec2 or on-premises.
     * 
     */
    @Export(name="infrastructureClass", refs={String.class}, tree="[0]")
    private Output<String> infrastructureClass;

    /**
     * @return For registered instances, infrastructure class: ec2 or on-premises.
     * 
     */
    public Output<String> infrastructureClass() {
        return this.infrastructureClass;
    }
    /**
     * Controls where to install OS and package updates when the instance boots.  Default is `true`.
     * 
     */
    @Export(name="installUpdatesOnBoot", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> installUpdatesOnBoot;

    /**
     * @return Controls where to install OS and package updates when the instance boots.  Default is `true`.
     * 
     */
    public Output<Optional<Boolean>> installUpdatesOnBoot() {
        return Codegen.optional(this.installUpdatesOnBoot);
    }
    /**
     * ARN of the instance&#39;s IAM profile.
     * 
     */
    @Export(name="instanceProfileArn", refs={String.class}, tree="[0]")
    private Output<String> instanceProfileArn;

    /**
     * @return ARN of the instance&#39;s IAM profile.
     * 
     */
    public Output<String> instanceProfileArn() {
        return this.instanceProfileArn;
    }
    /**
     * Type of instance to start.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceType;

    /**
     * @return Type of instance to start.
     * 
     */
    public Output<Optional<String>> instanceType() {
        return Codegen.optional(this.instanceType);
    }
    /**
     * ID of the last service error.
     * 
     */
    @Export(name="lastServiceErrorId", refs={String.class}, tree="[0]")
    private Output<String> lastServiceErrorId;

    /**
     * @return ID of the last service error.
     * 
     */
    public Output<String> lastServiceErrorId() {
        return this.lastServiceErrorId;
    }
    /**
     * List of the layers the instance will belong to.
     * 
     */
    @Export(name="layerIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> layerIds;

    /**
     * @return List of the layers the instance will belong to.
     * 
     */
    public Output<List<String>> layerIds() {
        return this.layerIds;
    }
    /**
     * Name of operating system that will be installed.
     * 
     */
    @Export(name="os", refs={String.class}, tree="[0]")
    private Output<String> os;

    /**
     * @return Name of operating system that will be installed.
     * 
     */
    public Output<String> os() {
        return this.os;
    }
    /**
     * Instance&#39;s platform.
     * 
     */
    @Export(name="platform", refs={String.class}, tree="[0]")
    private Output<String> platform;

    /**
     * @return Instance&#39;s platform.
     * 
     */
    public Output<String> platform() {
        return this.platform;
    }
    /**
     * Private DNS name assigned to the instance. Can only be used inside the Amazon EC2, and only available if you&#39;ve enabled DNS hostnames for your VPC.
     * 
     */
    @Export(name="privateDns", refs={String.class}, tree="[0]")
    private Output<String> privateDns;

    /**
     * @return Private DNS name assigned to the instance. Can only be used inside the Amazon EC2, and only available if you&#39;ve enabled DNS hostnames for your VPC.
     * 
     */
    public Output<String> privateDns() {
        return this.privateDns;
    }
    /**
     * Private IP address assigned to the instance.
     * 
     */
    @Export(name="privateIp", refs={String.class}, tree="[0]")
    private Output<String> privateIp;

    /**
     * @return Private IP address assigned to the instance.
     * 
     */
    public Output<String> privateIp() {
        return this.privateIp;
    }
    /**
     * Public DNS name assigned to the instance. For EC2-VPC, this is only available if you&#39;ve enabled DNS hostnames for your VPC.
     * 
     */
    @Export(name="publicDns", refs={String.class}, tree="[0]")
    private Output<String> publicDns;

    /**
     * @return Public DNS name assigned to the instance. For EC2-VPC, this is only available if you&#39;ve enabled DNS hostnames for your VPC.
     * 
     */
    public Output<String> publicDns() {
        return this.publicDns;
    }
    /**
     * Public IP address assigned to the instance, if applicable.
     * 
     */
    @Export(name="publicIp", refs={String.class}, tree="[0]")
    private Output<String> publicIp;

    /**
     * @return Public IP address assigned to the instance, if applicable.
     * 
     */
    public Output<String> publicIp() {
        return this.publicIp;
    }
    /**
     * For registered instances, who performed the registration.
     * 
     */
    @Export(name="registeredBy", refs={String.class}, tree="[0]")
    private Output<String> registeredBy;

    /**
     * @return For registered instances, who performed the registration.
     * 
     */
    public Output<String> registeredBy() {
        return this.registeredBy;
    }
    /**
     * Instance&#39;s reported AWS OpsWorks Stacks agent version.
     * 
     */
    @Export(name="reportedAgentVersion", refs={String.class}, tree="[0]")
    private Output<String> reportedAgentVersion;

    /**
     * @return Instance&#39;s reported AWS OpsWorks Stacks agent version.
     * 
     */
    public Output<String> reportedAgentVersion() {
        return this.reportedAgentVersion;
    }
    /**
     * For registered instances, the reported operating system family.
     * 
     */
    @Export(name="reportedOsFamily", refs={String.class}, tree="[0]")
    private Output<String> reportedOsFamily;

    /**
     * @return For registered instances, the reported operating system family.
     * 
     */
    public Output<String> reportedOsFamily() {
        return this.reportedOsFamily;
    }
    /**
     * For registered instances, the reported operating system name.
     * 
     */
    @Export(name="reportedOsName", refs={String.class}, tree="[0]")
    private Output<String> reportedOsName;

    /**
     * @return For registered instances, the reported operating system name.
     * 
     */
    public Output<String> reportedOsName() {
        return this.reportedOsName;
    }
    /**
     * For registered instances, the reported operating system version.
     * 
     */
    @Export(name="reportedOsVersion", refs={String.class}, tree="[0]")
    private Output<String> reportedOsVersion;

    /**
     * @return For registered instances, the reported operating system version.
     * 
     */
    public Output<String> reportedOsVersion() {
        return this.reportedOsVersion;
    }
    /**
     * Configuration block for the root block device of the instance. See Block Devices below.
     * 
     */
    @Export(name="rootBlockDevices", refs={List.class,InstanceRootBlockDevice.class}, tree="[0,1]")
    private Output<List<InstanceRootBlockDevice>> rootBlockDevices;

    /**
     * @return Configuration block for the root block device of the instance. See Block Devices below.
     * 
     */
    public Output<List<InstanceRootBlockDevice>> rootBlockDevices() {
        return this.rootBlockDevices;
    }
    /**
     * Name of the type of root device instances will have by default. Valid values are `ebs` or `instance-store`.
     * 
     */
    @Export(name="rootDeviceType", refs={String.class}, tree="[0]")
    private Output<String> rootDeviceType;

    /**
     * @return Name of the type of root device instances will have by default. Valid values are `ebs` or `instance-store`.
     * 
     */
    public Output<String> rootDeviceType() {
        return this.rootDeviceType;
    }
    /**
     * Root device volume ID.
     * 
     */
    @Export(name="rootDeviceVolumeId", refs={String.class}, tree="[0]")
    private Output<String> rootDeviceVolumeId;

    /**
     * @return Root device volume ID.
     * 
     */
    public Output<String> rootDeviceVolumeId() {
        return this.rootDeviceVolumeId;
    }
    /**
     * Associated security groups.
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroupIds;

    /**
     * @return Associated security groups.
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * SSH key&#39;s Deep Security Agent (DSA) fingerprint.
     * 
     */
    @Export(name="sshHostDsaKeyFingerprint", refs={String.class}, tree="[0]")
    private Output<String> sshHostDsaKeyFingerprint;

    /**
     * @return SSH key&#39;s Deep Security Agent (DSA) fingerprint.
     * 
     */
    public Output<String> sshHostDsaKeyFingerprint() {
        return this.sshHostDsaKeyFingerprint;
    }
    /**
     * SSH key&#39;s RSA fingerprint.
     * 
     */
    @Export(name="sshHostRsaKeyFingerprint", refs={String.class}, tree="[0]")
    private Output<String> sshHostRsaKeyFingerprint;

    /**
     * @return SSH key&#39;s RSA fingerprint.
     * 
     */
    public Output<String> sshHostRsaKeyFingerprint() {
        return this.sshHostRsaKeyFingerprint;
    }
    /**
     * Name of the SSH keypair that instances will have by default.
     * 
     */
    @Export(name="sshKeyName", refs={String.class}, tree="[0]")
    private Output<String> sshKeyName;

    /**
     * @return Name of the SSH keypair that instances will have by default.
     * 
     */
    public Output<String> sshKeyName() {
        return this.sshKeyName;
    }
    /**
     * Identifier of the stack the instance will belong to.
     * 
     */
    @Export(name="stackId", refs={String.class}, tree="[0]")
    private Output<String> stackId;

    /**
     * @return Identifier of the stack the instance will belong to.
     * 
     */
    public Output<String> stackId() {
        return this.stackId;
    }
    /**
     * Desired state of the instance. Valid values are `running` or `stopped`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return Desired state of the instance. Valid values are `running` or `stopped`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * Instance status. Will be one of `booting`, `connection_lost`, `online`, `pending`, `rebooting`, `requested`, `running_setup`, `setup_failed`, `shutting_down`, `start_failed`, `stop_failed`, `stopped`, `stopping`, `terminated`, or `terminating`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Instance status. Will be one of `booting`, `connection_lost`, `online`, `pending`, `rebooting`, `requested`, `running_setup`, `setup_failed`, `shutting_down`, `start_failed`, `stop_failed`, `stopped`, `stopping`, `terminated`, or `terminating`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Subnet ID to attach to.
     * 
     */
    @Export(name="subnetId", refs={String.class}, tree="[0]")
    private Output<String> subnetId;

    /**
     * @return Subnet ID to attach to.
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }
    /**
     * Instance tenancy to use. Valid values are `default`, `dedicated` or `host`.
     * 
     */
    @Export(name="tenancy", refs={String.class}, tree="[0]")
    private Output<String> tenancy;

    /**
     * @return Instance tenancy to use. Valid values are `default`, `dedicated` or `host`.
     * 
     */
    public Output<String> tenancy() {
        return this.tenancy;
    }
    /**
     * Keyword to choose what virtualization mode created instances will use. Valid values are `paravirtual` or `hvm`.
     * 
     */
    @Export(name="virtualizationType", refs={String.class}, tree="[0]")
    private Output<String> virtualizationType;

    /**
     * @return Keyword to choose what virtualization mode created instances will use. Valid values are `paravirtual` or `hvm`.
     * 
     */
    public Output<String> virtualizationType() {
        return this.virtualizationType;
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
        super("aws:opsworks/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:opsworks/instance:Instance", name, state, makeResourceOptions(options, id));
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
