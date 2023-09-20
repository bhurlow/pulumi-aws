// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx.inputs;

import com.pulumi.aws.fsx.inputs.OntapStorageVirtualMachineActiveDirectoryConfigurationArgs;
import com.pulumi.aws.fsx.inputs.OntapStorageVirtualMachineEndpointArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OntapStorageVirtualMachineState extends com.pulumi.resources.ResourceArgs {

    public static final OntapStorageVirtualMachineState Empty = new OntapStorageVirtualMachineState();

    /**
     * Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
     * 
     */
    @Import(name="activeDirectoryConfiguration")
    private @Nullable Output<OntapStorageVirtualMachineActiveDirectoryConfigurationArgs> activeDirectoryConfiguration;

    /**
     * @return Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
     * 
     */
    public Optional<Output<OntapStorageVirtualMachineActiveDirectoryConfigurationArgs>> activeDirectoryConfiguration() {
        return Optional.ofNullable(this.activeDirectoryConfiguration);
    }

    /**
     * Amazon Resource Name of the storage virtual machine.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name of the storage virtual machine.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
     * 
     */
    @Import(name="endpoints")
    private @Nullable Output<List<OntapStorageVirtualMachineEndpointArgs>> endpoints;

    /**
     * @return The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
     * 
     */
    public Optional<Output<List<OntapStorageVirtualMachineEndpointArgs>>> endpoints() {
        return Optional.ofNullable(this.endpoints);
    }

    /**
     * The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
     * 
     */
    @Import(name="fileSystemId")
    private @Nullable Output<String> fileSystemId;

    /**
     * @return The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
     * 
     */
    public Optional<Output<String>> fileSystemId() {
        return Optional.ofNullable(this.fileSystemId);
    }

    /**
     * The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
     * 
     */
    @Import(name="rootVolumeSecurityStyle")
    private @Nullable Output<String> rootVolumeSecurityStyle;

    /**
     * @return Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
     * 
     */
    public Optional<Output<String>> rootVolumeSecurityStyle() {
        return Optional.ofNullable(this.rootVolumeSecurityStyle);
    }

    /**
     * Describes the SVM&#39;s subtype, e.g. `DEFAULT`
     * 
     */
    @Import(name="subtype")
    private @Nullable Output<String> subtype;

    /**
     * @return Describes the SVM&#39;s subtype, e.g. `DEFAULT`
     * 
     */
    public Optional<Output<String>> subtype() {
        return Optional.ofNullable(this.subtype);
    }

    @Import(name="svmAdminPassword")
    private @Nullable Output<String> svmAdminPassword;

    public Optional<Output<String>> svmAdminPassword() {
        return Optional.ofNullable(this.svmAdminPassword);
    }

    /**
     * A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The SVM&#39;s UUID (universally unique identifier).
     * 
     */
    @Import(name="uuid")
    private @Nullable Output<String> uuid;

    /**
     * @return The SVM&#39;s UUID (universally unique identifier).
     * 
     */
    public Optional<Output<String>> uuid() {
        return Optional.ofNullable(this.uuid);
    }

    private OntapStorageVirtualMachineState() {}

    private OntapStorageVirtualMachineState(OntapStorageVirtualMachineState $) {
        this.activeDirectoryConfiguration = $.activeDirectoryConfiguration;
        this.arn = $.arn;
        this.endpoints = $.endpoints;
        this.fileSystemId = $.fileSystemId;
        this.name = $.name;
        this.rootVolumeSecurityStyle = $.rootVolumeSecurityStyle;
        this.subtype = $.subtype;
        this.svmAdminPassword = $.svmAdminPassword;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.uuid = $.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OntapStorageVirtualMachineState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OntapStorageVirtualMachineState $;

        public Builder() {
            $ = new OntapStorageVirtualMachineState();
        }

        public Builder(OntapStorageVirtualMachineState defaults) {
            $ = new OntapStorageVirtualMachineState(Objects.requireNonNull(defaults));
        }

        /**
         * @param activeDirectoryConfiguration Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder activeDirectoryConfiguration(@Nullable Output<OntapStorageVirtualMachineActiveDirectoryConfigurationArgs> activeDirectoryConfiguration) {
            $.activeDirectoryConfiguration = activeDirectoryConfiguration;
            return this;
        }

        /**
         * @param activeDirectoryConfiguration Configuration block that Amazon FSx uses to join the FSx ONTAP Storage Virtual Machine(SVM) to your Microsoft Active Directory (AD) directory. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder activeDirectoryConfiguration(OntapStorageVirtualMachineActiveDirectoryConfigurationArgs activeDirectoryConfiguration) {
            return activeDirectoryConfiguration(Output.of(activeDirectoryConfiguration));
        }

        /**
         * @param arn Amazon Resource Name of the storage virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name of the storage virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param endpoints The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(@Nullable Output<List<OntapStorageVirtualMachineEndpointArgs>> endpoints) {
            $.endpoints = endpoints;
            return this;
        }

        /**
         * @param endpoints The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(List<OntapStorageVirtualMachineEndpointArgs> endpoints) {
            return endpoints(Output.of(endpoints));
        }

        /**
         * @param endpoints The endpoints that are used to access data or to manage the storage virtual machine using the NetApp ONTAP CLI, REST API, or NetApp SnapMirror. See Endpoints below.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(OntapStorageVirtualMachineEndpointArgs... endpoints) {
            return endpoints(List.of(endpoints));
        }

        /**
         * @param fileSystemId The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
         * 
         * @return builder
         * 
         */
        public Builder fileSystemId(@Nullable Output<String> fileSystemId) {
            $.fileSystemId = fileSystemId;
            return this;
        }

        /**
         * @param fileSystemId The ID of the Amazon FSx ONTAP File System that this SVM will be created on.
         * 
         * @return builder
         * 
         */
        public Builder fileSystemId(String fileSystemId) {
            return fileSystemId(Output.of(fileSystemId));
        }

        /**
         * @param name The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the SVM. You can use a maximum of 47 alphanumeric characters, plus the underscore (_) special character.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param rootVolumeSecurityStyle Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
         * 
         * @return builder
         * 
         */
        public Builder rootVolumeSecurityStyle(@Nullable Output<String> rootVolumeSecurityStyle) {
            $.rootVolumeSecurityStyle = rootVolumeSecurityStyle;
            return this;
        }

        /**
         * @param rootVolumeSecurityStyle Specifies the root volume security style, Valid values are `UNIX`, `NTFS`, and `MIXED`. All volumes created under this SVM will inherit the root security style unless the security style is specified on the volume. Default value is `UNIX`.
         * 
         * @return builder
         * 
         */
        public Builder rootVolumeSecurityStyle(String rootVolumeSecurityStyle) {
            return rootVolumeSecurityStyle(Output.of(rootVolumeSecurityStyle));
        }

        /**
         * @param subtype Describes the SVM&#39;s subtype, e.g. `DEFAULT`
         * 
         * @return builder
         * 
         */
        public Builder subtype(@Nullable Output<String> subtype) {
            $.subtype = subtype;
            return this;
        }

        /**
         * @param subtype Describes the SVM&#39;s subtype, e.g. `DEFAULT`
         * 
         * @return builder
         * 
         */
        public Builder subtype(String subtype) {
            return subtype(Output.of(subtype));
        }

        public Builder svmAdminPassword(@Nullable Output<String> svmAdminPassword) {
            $.svmAdminPassword = svmAdminPassword;
            return this;
        }

        public Builder svmAdminPassword(String svmAdminPassword) {
            return svmAdminPassword(Output.of(svmAdminPassword));
        }

        /**
         * @param tags A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the storage virtual machine. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param uuid The SVM&#39;s UUID (universally unique identifier).
         * 
         * @return builder
         * 
         */
        public Builder uuid(@Nullable Output<String> uuid) {
            $.uuid = uuid;
            return this;
        }

        /**
         * @param uuid The SVM&#39;s UUID (universally unique identifier).
         * 
         * @return builder
         * 
         */
        public Builder uuid(String uuid) {
            return uuid(Output.of(uuid));
        }

        public OntapStorageVirtualMachineState build() {
            return $;
        }
    }

}
