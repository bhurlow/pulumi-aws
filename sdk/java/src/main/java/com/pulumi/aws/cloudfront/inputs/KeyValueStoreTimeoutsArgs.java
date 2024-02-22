// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudfront.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KeyValueStoreTimeoutsArgs extends com.pulumi.resources.ResourceArgs {

    public static final KeyValueStoreTimeoutsArgs Empty = new KeyValueStoreTimeoutsArgs();

    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
     * 
     */
    @Import(name="create")
    private @Nullable Output<String> create;

    /**
     * @return A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
     * 
     */
    public Optional<Output<String>> create() {
        return Optional.ofNullable(this.create);
    }

    private KeyValueStoreTimeoutsArgs() {}

    private KeyValueStoreTimeoutsArgs(KeyValueStoreTimeoutsArgs $) {
        this.create = $.create;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KeyValueStoreTimeoutsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KeyValueStoreTimeoutsArgs $;

        public Builder() {
            $ = new KeyValueStoreTimeoutsArgs();
        }

        public Builder(KeyValueStoreTimeoutsArgs defaults) {
            $ = new KeyValueStoreTimeoutsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param create A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
         * 
         * @return builder
         * 
         */
        public Builder create(@Nullable Output<String> create) {
            $.create = create;
            return this;
        }

        /**
         * @param create A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
         * 
         * @return builder
         * 
         */
        public Builder create(String create) {
            return create(Output.of(create));
        }

        public KeyValueStoreTimeoutsArgs build() {
            return $;
        }
    }

}
