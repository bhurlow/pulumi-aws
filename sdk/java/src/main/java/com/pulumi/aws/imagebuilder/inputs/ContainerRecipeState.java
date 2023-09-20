// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.imagebuilder.inputs;

import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeComponentArgs;
import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeInstanceConfigurationArgs;
import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeTargetRepositoryArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerRecipeState extends com.pulumi.resources.ResourceArgs {

    public static final ContainerRecipeState Empty = new ContainerRecipeState();

    /**
     * (Required) Amazon Resource Name (ARN) of the container recipe.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return (Required) Amazon Resource Name (ARN) of the container recipe.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Ordered configuration block(s) with components for the container recipe. Detailed below.
     * 
     */
    @Import(name="components")
    private @Nullable Output<List<ContainerRecipeComponentArgs>> components;

    /**
     * @return Ordered configuration block(s) with components for the container recipe. Detailed below.
     * 
     */
    public Optional<Output<List<ContainerRecipeComponentArgs>>> components() {
        return Optional.ofNullable(this.components);
    }

    /**
     * The type of the container to create. Valid values: `DOCKER`.
     * 
     */
    @Import(name="containerType")
    private @Nullable Output<String> containerType;

    /**
     * @return The type of the container to create. Valid values: `DOCKER`.
     * 
     */
    public Optional<Output<String>> containerType() {
        return Optional.ofNullable(this.containerType);
    }

    /**
     * Date the container recipe was created.
     * 
     */
    @Import(name="dateCreated")
    private @Nullable Output<String> dateCreated;

    /**
     * @return Date the container recipe was created.
     * 
     */
    public Optional<Output<String>> dateCreated() {
        return Optional.ofNullable(this.dateCreated);
    }

    /**
     * The description of the container recipe.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the container recipe.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The Dockerfile template used to build the image as an inline data blob.
     * 
     */
    @Import(name="dockerfileTemplateData")
    private @Nullable Output<String> dockerfileTemplateData;

    /**
     * @return The Dockerfile template used to build the image as an inline data blob.
     * 
     */
    public Optional<Output<String>> dockerfileTemplateData() {
        return Optional.ofNullable(this.dockerfileTemplateData);
    }

    /**
     * The Amazon S3 URI for the Dockerfile that will be used to build the container image.
     * 
     */
    @Import(name="dockerfileTemplateUri")
    private @Nullable Output<String> dockerfileTemplateUri;

    /**
     * @return The Amazon S3 URI for the Dockerfile that will be used to build the container image.
     * 
     */
    public Optional<Output<String>> dockerfileTemplateUri() {
        return Optional.ofNullable(this.dockerfileTemplateUri);
    }

    /**
     * Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
     * 
     */
    @Import(name="encrypted")
    private @Nullable Output<Boolean> encrypted;

    /**
     * @return Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
     * 
     */
    public Optional<Output<Boolean>> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    /**
     * Configuration block used to configure an instance for building and testing container images. Detailed below.
     * 
     */
    @Import(name="instanceConfiguration")
    private @Nullable Output<ContainerRecipeInstanceConfigurationArgs> instanceConfiguration;

    /**
     * @return Configuration block used to configure an instance for building and testing container images. Detailed below.
     * 
     */
    public Optional<Output<ContainerRecipeInstanceConfigurationArgs>> instanceConfiguration() {
        return Optional.ofNullable(this.instanceConfiguration);
    }

    /**
     * The KMS key used to encrypt the container image.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return The KMS key used to encrypt the container image.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * The name of the container recipe.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the container recipe.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Owner of the container recipe.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return Owner of the container recipe.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * The base image for the container recipe.
     * 
     */
    @Import(name="parentImage")
    private @Nullable Output<String> parentImage;

    /**
     * @return The base image for the container recipe.
     * 
     */
    public Optional<Output<String>> parentImage() {
        return Optional.ofNullable(this.parentImage);
    }

    /**
     * Platform of the container recipe.
     * 
     */
    @Import(name="platform")
    private @Nullable Output<String> platform;

    /**
     * @return Platform of the container recipe.
     * 
     */
    public Optional<Output<String>> platform() {
        return Optional.ofNullable(this.platform);
    }

    /**
     * Specifies the operating system platform when you use a custom base image.
     * 
     */
    @Import(name="platformOverride")
    private @Nullable Output<String> platformOverride;

    /**
     * @return Specifies the operating system platform when you use a custom base image.
     * 
     */
    public Optional<Output<String>> platformOverride() {
        return Optional.ofNullable(this.platformOverride);
    }

    /**
     * Key-value map of resource tags for the container recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags for the container recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The destination repository for the container image. Detailed below.
     * 
     */
    @Import(name="targetRepository")
    private @Nullable Output<ContainerRecipeTargetRepositoryArgs> targetRepository;

    /**
     * @return The destination repository for the container image. Detailed below.
     * 
     */
    public Optional<Output<ContainerRecipeTargetRepositoryArgs>> targetRepository() {
        return Optional.ofNullable(this.targetRepository);
    }

    /**
     * Version of the container recipe.
     * 
     * The following attributes are optional:
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version of the container recipe.
     * 
     * The following attributes are optional:
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * The working directory to be used during build and test workflows.
     * 
     */
    @Import(name="workingDirectory")
    private @Nullable Output<String> workingDirectory;

    /**
     * @return The working directory to be used during build and test workflows.
     * 
     */
    public Optional<Output<String>> workingDirectory() {
        return Optional.ofNullable(this.workingDirectory);
    }

    private ContainerRecipeState() {}

    private ContainerRecipeState(ContainerRecipeState $) {
        this.arn = $.arn;
        this.components = $.components;
        this.containerType = $.containerType;
        this.dateCreated = $.dateCreated;
        this.description = $.description;
        this.dockerfileTemplateData = $.dockerfileTemplateData;
        this.dockerfileTemplateUri = $.dockerfileTemplateUri;
        this.encrypted = $.encrypted;
        this.instanceConfiguration = $.instanceConfiguration;
        this.kmsKeyId = $.kmsKeyId;
        this.name = $.name;
        this.owner = $.owner;
        this.parentImage = $.parentImage;
        this.platform = $.platform;
        this.platformOverride = $.platformOverride;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.targetRepository = $.targetRepository;
        this.version = $.version;
        this.workingDirectory = $.workingDirectory;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerRecipeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerRecipeState $;

        public Builder() {
            $ = new ContainerRecipeState();
        }

        public Builder(ContainerRecipeState defaults) {
            $ = new ContainerRecipeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn (Required) Amazon Resource Name (ARN) of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn (Required) Amazon Resource Name (ARN) of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param components Ordered configuration block(s) with components for the container recipe. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder components(@Nullable Output<List<ContainerRecipeComponentArgs>> components) {
            $.components = components;
            return this;
        }

        /**
         * @param components Ordered configuration block(s) with components for the container recipe. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder components(List<ContainerRecipeComponentArgs> components) {
            return components(Output.of(components));
        }

        /**
         * @param components Ordered configuration block(s) with components for the container recipe. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder components(ContainerRecipeComponentArgs... components) {
            return components(List.of(components));
        }

        /**
         * @param containerType The type of the container to create. Valid values: `DOCKER`.
         * 
         * @return builder
         * 
         */
        public Builder containerType(@Nullable Output<String> containerType) {
            $.containerType = containerType;
            return this;
        }

        /**
         * @param containerType The type of the container to create. Valid values: `DOCKER`.
         * 
         * @return builder
         * 
         */
        public Builder containerType(String containerType) {
            return containerType(Output.of(containerType));
        }

        /**
         * @param dateCreated Date the container recipe was created.
         * 
         * @return builder
         * 
         */
        public Builder dateCreated(@Nullable Output<String> dateCreated) {
            $.dateCreated = dateCreated;
            return this;
        }

        /**
         * @param dateCreated Date the container recipe was created.
         * 
         * @return builder
         * 
         */
        public Builder dateCreated(String dateCreated) {
            return dateCreated(Output.of(dateCreated));
        }

        /**
         * @param description The description of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dockerfileTemplateData The Dockerfile template used to build the image as an inline data blob.
         * 
         * @return builder
         * 
         */
        public Builder dockerfileTemplateData(@Nullable Output<String> dockerfileTemplateData) {
            $.dockerfileTemplateData = dockerfileTemplateData;
            return this;
        }

        /**
         * @param dockerfileTemplateData The Dockerfile template used to build the image as an inline data blob.
         * 
         * @return builder
         * 
         */
        public Builder dockerfileTemplateData(String dockerfileTemplateData) {
            return dockerfileTemplateData(Output.of(dockerfileTemplateData));
        }

        /**
         * @param dockerfileTemplateUri The Amazon S3 URI for the Dockerfile that will be used to build the container image.
         * 
         * @return builder
         * 
         */
        public Builder dockerfileTemplateUri(@Nullable Output<String> dockerfileTemplateUri) {
            $.dockerfileTemplateUri = dockerfileTemplateUri;
            return this;
        }

        /**
         * @param dockerfileTemplateUri The Amazon S3 URI for the Dockerfile that will be used to build the container image.
         * 
         * @return builder
         * 
         */
        public Builder dockerfileTemplateUri(String dockerfileTemplateUri) {
            return dockerfileTemplateUri(Output.of(dockerfileTemplateUri));
        }

        /**
         * @param encrypted Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable Output<Boolean> encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param encrypted Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(Boolean encrypted) {
            return encrypted(Output.of(encrypted));
        }

        /**
         * @param instanceConfiguration Configuration block used to configure an instance for building and testing container images. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder instanceConfiguration(@Nullable Output<ContainerRecipeInstanceConfigurationArgs> instanceConfiguration) {
            $.instanceConfiguration = instanceConfiguration;
            return this;
        }

        /**
         * @param instanceConfiguration Configuration block used to configure an instance for building and testing container images. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder instanceConfiguration(ContainerRecipeInstanceConfigurationArgs instanceConfiguration) {
            return instanceConfiguration(Output.of(instanceConfiguration));
        }

        /**
         * @param kmsKeyId The KMS key used to encrypt the container image.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId The KMS key used to encrypt the container image.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        /**
         * @param name The name of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param owner Owner of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner Owner of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param parentImage The base image for the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder parentImage(@Nullable Output<String> parentImage) {
            $.parentImage = parentImage;
            return this;
        }

        /**
         * @param parentImage The base image for the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder parentImage(String parentImage) {
            return parentImage(Output.of(parentImage));
        }

        /**
         * @param platform Platform of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder platform(@Nullable Output<String> platform) {
            $.platform = platform;
            return this;
        }

        /**
         * @param platform Platform of the container recipe.
         * 
         * @return builder
         * 
         */
        public Builder platform(String platform) {
            return platform(Output.of(platform));
        }

        /**
         * @param platformOverride Specifies the operating system platform when you use a custom base image.
         * 
         * @return builder
         * 
         */
        public Builder platformOverride(@Nullable Output<String> platformOverride) {
            $.platformOverride = platformOverride;
            return this;
        }

        /**
         * @param platformOverride Specifies the operating system platform when you use a custom base image.
         * 
         * @return builder
         * 
         */
        public Builder platformOverride(String platformOverride) {
            return platformOverride(Output.of(platformOverride));
        }

        /**
         * @param tags Key-value map of resource tags for the container recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags for the container recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param targetRepository The destination repository for the container image. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder targetRepository(@Nullable Output<ContainerRecipeTargetRepositoryArgs> targetRepository) {
            $.targetRepository = targetRepository;
            return this;
        }

        /**
         * @param targetRepository The destination repository for the container image. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder targetRepository(ContainerRecipeTargetRepositoryArgs targetRepository) {
            return targetRepository(Output.of(targetRepository));
        }

        /**
         * @param version Version of the container recipe.
         * 
         * The following attributes are optional:
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the container recipe.
         * 
         * The following attributes are optional:
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param workingDirectory The working directory to be used during build and test workflows.
         * 
         * @return builder
         * 
         */
        public Builder workingDirectory(@Nullable Output<String> workingDirectory) {
            $.workingDirectory = workingDirectory;
            return this;
        }

        /**
         * @param workingDirectory The working directory to be used during build and test workflows.
         * 
         * @return builder
         * 
         */
        public Builder workingDirectory(String workingDirectory) {
            return workingDirectory(Output.of(workingDirectory));
        }

        public ContainerRecipeState build() {
            return $;
        }
    }

}
