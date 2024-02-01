// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class V2modelsIntentTimeoutsArgs extends com.pulumi.resources.ResourceArgs {

    public static final V2modelsIntentTimeoutsArgs Empty = new V2modelsIntentTimeoutsArgs();

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

    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
     * 
     */
    @Import(name="delete")
    private @Nullable Output<String> delete;

    /**
     * @return A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
     * 
     */
    public Optional<Output<String>> delete() {
        return Optional.ofNullable(this.delete);
    }

    /**
     * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
     * 
     */
    @Import(name="update")
    private @Nullable Output<String> update;

    /**
     * @return A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
     * 
     */
    public Optional<Output<String>> update() {
        return Optional.ofNullable(this.update);
    }

    private V2modelsIntentTimeoutsArgs() {}

    private V2modelsIntentTimeoutsArgs(V2modelsIntentTimeoutsArgs $) {
        this.create = $.create;
        this.delete = $.delete;
        this.update = $.update;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(V2modelsIntentTimeoutsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private V2modelsIntentTimeoutsArgs $;

        public Builder() {
            $ = new V2modelsIntentTimeoutsArgs();
        }

        public Builder(V2modelsIntentTimeoutsArgs defaults) {
            $ = new V2modelsIntentTimeoutsArgs(Objects.requireNonNull(defaults));
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

        /**
         * @param delete A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
         * 
         * @return builder
         * 
         */
        public Builder delete(@Nullable Output<String> delete) {
            $.delete = delete;
            return this;
        }

        /**
         * @param delete A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
         * 
         * @return builder
         * 
         */
        public Builder delete(String delete) {
            return delete(Output.of(delete));
        }

        /**
         * @param update A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
         * 
         * @return builder
         * 
         */
        public Builder update(@Nullable Output<String> update) {
            $.update = update;
            return this;
        }

        /**
         * @param update A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).
         * 
         * @return builder
         * 
         */
        public Builder update(String update) {
            return update(Output.of(update));
        }

        public V2modelsIntentTimeoutsArgs build() {
            return $;
        }
    }

}
