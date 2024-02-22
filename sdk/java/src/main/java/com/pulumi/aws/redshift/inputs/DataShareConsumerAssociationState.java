// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.redshift.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DataShareConsumerAssociationState extends com.pulumi.resources.ResourceArgs {

    public static final DataShareConsumerAssociationState Empty = new DataShareConsumerAssociationState();

    /**
     * Whether to allow write operations for a datashare.
     * 
     */
    @Import(name="allowWrites")
    private @Nullable Output<Boolean> allowWrites;

    /**
     * @return Whether to allow write operations for a datashare.
     * 
     */
    public Optional<Output<Boolean>> allowWrites() {
        return Optional.ofNullable(this.allowWrites);
    }

    /**
     * Whether the datashare is associated with the entire account. Conflicts with `consumer_arn` and `consumer_region`.
     * 
     */
    @Import(name="associateEntireAccount")
    private @Nullable Output<Boolean> associateEntireAccount;

    /**
     * @return Whether the datashare is associated with the entire account. Conflicts with `consumer_arn` and `consumer_region`.
     * 
     */
    public Optional<Output<Boolean>> associateEntireAccount() {
        return Optional.ofNullable(this.associateEntireAccount);
    }

    /**
     * Amazon Resource Name (ARN) of the consumer that is associated with the datashare. Conflicts with `associate_entire_account` and `consumer_region`.
     * 
     */
    @Import(name="consumerArn")
    private @Nullable Output<String> consumerArn;

    /**
     * @return Amazon Resource Name (ARN) of the consumer that is associated with the datashare. Conflicts with `associate_entire_account` and `consumer_region`.
     * 
     */
    public Optional<Output<String>> consumerArn() {
        return Optional.ofNullable(this.consumerArn);
    }

    /**
     * From a datashare consumer account, associates a datashare with all existing and future namespaces in the specified AWS Region. Conflicts with `associate_entire_account` and `consumer_arn`.
     * 
     */
    @Import(name="consumerRegion")
    private @Nullable Output<String> consumerRegion;

    /**
     * @return From a datashare consumer account, associates a datashare with all existing and future namespaces in the specified AWS Region. Conflicts with `associate_entire_account` and `consumer_arn`.
     * 
     */
    public Optional<Output<String>> consumerRegion() {
        return Optional.ofNullable(this.consumerRegion);
    }

    /**
     * Amazon Resource Name (ARN) of the datashare that the consumer is to use with the account or the namespace.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="dataShareArn")
    private @Nullable Output<String> dataShareArn;

    /**
     * @return Amazon Resource Name (ARN) of the datashare that the consumer is to use with the account or the namespace.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> dataShareArn() {
        return Optional.ofNullable(this.dataShareArn);
    }

    /**
     * Identifier of a datashare to show its managing entity.
     * 
     */
    @Import(name="managedBy")
    private @Nullable Output<String> managedBy;

    /**
     * @return Identifier of a datashare to show its managing entity.
     * 
     */
    public Optional<Output<String>> managedBy() {
        return Optional.ofNullable(this.managedBy);
    }

    /**
     * Amazon Resource Name (ARN) of the producer.
     * 
     */
    @Import(name="producerArn")
    private @Nullable Output<String> producerArn;

    /**
     * @return Amazon Resource Name (ARN) of the producer.
     * 
     */
    public Optional<Output<String>> producerArn() {
        return Optional.ofNullable(this.producerArn);
    }

    private DataShareConsumerAssociationState() {}

    private DataShareConsumerAssociationState(DataShareConsumerAssociationState $) {
        this.allowWrites = $.allowWrites;
        this.associateEntireAccount = $.associateEntireAccount;
        this.consumerArn = $.consumerArn;
        this.consumerRegion = $.consumerRegion;
        this.dataShareArn = $.dataShareArn;
        this.managedBy = $.managedBy;
        this.producerArn = $.producerArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DataShareConsumerAssociationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DataShareConsumerAssociationState $;

        public Builder() {
            $ = new DataShareConsumerAssociationState();
        }

        public Builder(DataShareConsumerAssociationState defaults) {
            $ = new DataShareConsumerAssociationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowWrites Whether to allow write operations for a datashare.
         * 
         * @return builder
         * 
         */
        public Builder allowWrites(@Nullable Output<Boolean> allowWrites) {
            $.allowWrites = allowWrites;
            return this;
        }

        /**
         * @param allowWrites Whether to allow write operations for a datashare.
         * 
         * @return builder
         * 
         */
        public Builder allowWrites(Boolean allowWrites) {
            return allowWrites(Output.of(allowWrites));
        }

        /**
         * @param associateEntireAccount Whether the datashare is associated with the entire account. Conflicts with `consumer_arn` and `consumer_region`.
         * 
         * @return builder
         * 
         */
        public Builder associateEntireAccount(@Nullable Output<Boolean> associateEntireAccount) {
            $.associateEntireAccount = associateEntireAccount;
            return this;
        }

        /**
         * @param associateEntireAccount Whether the datashare is associated with the entire account. Conflicts with `consumer_arn` and `consumer_region`.
         * 
         * @return builder
         * 
         */
        public Builder associateEntireAccount(Boolean associateEntireAccount) {
            return associateEntireAccount(Output.of(associateEntireAccount));
        }

        /**
         * @param consumerArn Amazon Resource Name (ARN) of the consumer that is associated with the datashare. Conflicts with `associate_entire_account` and `consumer_region`.
         * 
         * @return builder
         * 
         */
        public Builder consumerArn(@Nullable Output<String> consumerArn) {
            $.consumerArn = consumerArn;
            return this;
        }

        /**
         * @param consumerArn Amazon Resource Name (ARN) of the consumer that is associated with the datashare. Conflicts with `associate_entire_account` and `consumer_region`.
         * 
         * @return builder
         * 
         */
        public Builder consumerArn(String consumerArn) {
            return consumerArn(Output.of(consumerArn));
        }

        /**
         * @param consumerRegion From a datashare consumer account, associates a datashare with all existing and future namespaces in the specified AWS Region. Conflicts with `associate_entire_account` and `consumer_arn`.
         * 
         * @return builder
         * 
         */
        public Builder consumerRegion(@Nullable Output<String> consumerRegion) {
            $.consumerRegion = consumerRegion;
            return this;
        }

        /**
         * @param consumerRegion From a datashare consumer account, associates a datashare with all existing and future namespaces in the specified AWS Region. Conflicts with `associate_entire_account` and `consumer_arn`.
         * 
         * @return builder
         * 
         */
        public Builder consumerRegion(String consumerRegion) {
            return consumerRegion(Output.of(consumerRegion));
        }

        /**
         * @param dataShareArn Amazon Resource Name (ARN) of the datashare that the consumer is to use with the account or the namespace.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder dataShareArn(@Nullable Output<String> dataShareArn) {
            $.dataShareArn = dataShareArn;
            return this;
        }

        /**
         * @param dataShareArn Amazon Resource Name (ARN) of the datashare that the consumer is to use with the account or the namespace.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder dataShareArn(String dataShareArn) {
            return dataShareArn(Output.of(dataShareArn));
        }

        /**
         * @param managedBy Identifier of a datashare to show its managing entity.
         * 
         * @return builder
         * 
         */
        public Builder managedBy(@Nullable Output<String> managedBy) {
            $.managedBy = managedBy;
            return this;
        }

        /**
         * @param managedBy Identifier of a datashare to show its managing entity.
         * 
         * @return builder
         * 
         */
        public Builder managedBy(String managedBy) {
            return managedBy(Output.of(managedBy));
        }

        /**
         * @param producerArn Amazon Resource Name (ARN) of the producer.
         * 
         * @return builder
         * 
         */
        public Builder producerArn(@Nullable Output<String> producerArn) {
            $.producerArn = producerArn;
            return this;
        }

        /**
         * @param producerArn Amazon Resource Name (ARN) of the producer.
         * 
         * @return builder
         * 
         */
        public Builder producerArn(String producerArn) {
            return producerArn(Output.of(producerArn));
        }

        public DataShareConsumerAssociationState build() {
            return $;
        }
    }

}
