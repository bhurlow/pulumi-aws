// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight.inputs;

import com.pulumi.aws.quicksight.inputs.ThemeConfigurationSheetTileBorderArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ThemeConfigurationSheetTileArgs extends com.pulumi.resources.ResourceArgs {

    public static final ThemeConfigurationSheetTileArgs Empty = new ThemeConfigurationSheetTileArgs();

    /**
     * The border around a tile. See border.
     * 
     */
    @Import(name="border")
    private @Nullable Output<ThemeConfigurationSheetTileBorderArgs> border;

    /**
     * @return The border around a tile. See border.
     * 
     */
    public Optional<Output<ThemeConfigurationSheetTileBorderArgs>> border() {
        return Optional.ofNullable(this.border);
    }

    private ThemeConfigurationSheetTileArgs() {}

    private ThemeConfigurationSheetTileArgs(ThemeConfigurationSheetTileArgs $) {
        this.border = $.border;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ThemeConfigurationSheetTileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ThemeConfigurationSheetTileArgs $;

        public Builder() {
            $ = new ThemeConfigurationSheetTileArgs();
        }

        public Builder(ThemeConfigurationSheetTileArgs defaults) {
            $ = new ThemeConfigurationSheetTileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param border The border around a tile. See border.
         * 
         * @return builder
         * 
         */
        public Builder border(@Nullable Output<ThemeConfigurationSheetTileBorderArgs> border) {
            $.border = border;
            return this;
        }

        /**
         * @param border The border around a tile. See border.
         * 
         * @return builder
         * 
         */
        public Builder border(ThemeConfigurationSheetTileBorderArgs border) {
            return border(Output.of(border));
        }

        public ThemeConfigurationSheetTileArgs build() {
            return $;
        }
    }

}
