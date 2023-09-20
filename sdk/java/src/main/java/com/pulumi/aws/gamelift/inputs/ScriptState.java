// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.gamelift.inputs;

import com.pulumi.aws.gamelift.inputs.ScriptStorageLocationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ScriptState extends com.pulumi.resources.ResourceArgs {

    public static final ScriptState Empty = new ScriptState();

    /**
     * GameLift Script ARN.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return GameLift Script ARN.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Name of the script
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the script
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Information indicating where your game script files are stored. See below.
     * 
     */
    @Import(name="storageLocation")
    private @Nullable Output<ScriptStorageLocationArgs> storageLocation;

    /**
     * @return Information indicating where your game script files are stored. See below.
     * 
     */
    public Optional<Output<ScriptStorageLocationArgs>> storageLocation() {
        return Optional.ofNullable(this.storageLocation);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Version that is associated with this script.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version that is associated with this script.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
     * 
     */
    @Import(name="zipFile")
    private @Nullable Output<String> zipFile;

    /**
     * @return A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
     * 
     */
    public Optional<Output<String>> zipFile() {
        return Optional.ofNullable(this.zipFile);
    }

    private ScriptState() {}

    private ScriptState(ScriptState $) {
        this.arn = $.arn;
        this.name = $.name;
        this.storageLocation = $.storageLocation;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.version = $.version;
        this.zipFile = $.zipFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ScriptState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ScriptState $;

        public Builder() {
            $ = new ScriptState();
        }

        public Builder(ScriptState defaults) {
            $ = new ScriptState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn GameLift Script ARN.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn GameLift Script ARN.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param name Name of the script
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the script
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param storageLocation Information indicating where your game script files are stored. See below.
         * 
         * @return builder
         * 
         */
        public Builder storageLocation(@Nullable Output<ScriptStorageLocationArgs> storageLocation) {
            $.storageLocation = storageLocation;
            return this;
        }

        /**
         * @param storageLocation Information indicating where your game script files are stored. See below.
         * 
         * @return builder
         * 
         */
        public Builder storageLocation(ScriptStorageLocationArgs storageLocation) {
            return storageLocation(Output.of(storageLocation));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param version Version that is associated with this script.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version that is associated with this script.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param zipFile A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
         * 
         * @return builder
         * 
         */
        public Builder zipFile(@Nullable Output<String> zipFile) {
            $.zipFile = zipFile;
            return this;
        }

        /**
         * @param zipFile A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
         * 
         * @return builder
         * 
         */
        public Builder zipFile(String zipFile) {
            return zipFile(Output.of(zipFile));
        }

        public ScriptState build() {
            return $;
        }
    }

}
