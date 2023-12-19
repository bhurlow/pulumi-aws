// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cloudfront.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionState extends com.pulumi.resources.ResourceArgs {

    public static final FunctionState Empty = new FunctionState();

    /**
     * Amazon Resource Name (ARN) identifying your CloudFront Function.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) identifying your CloudFront Function.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Source code of the function
     * 
     */
    @Import(name="code")
    private @Nullable Output<String> code;

    /**
     * @return Source code of the function
     * 
     */
    public Optional<Output<String>> code() {
        return Optional.ofNullable(this.code);
    }

    /**
     * Comment.
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return Comment.
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * ETag hash of any `LIVE` stage of the function.
     * 
     */
    @Import(name="liveStageEtag")
    private @Nullable Output<String> liveStageEtag;

    /**
     * @return ETag hash of any `LIVE` stage of the function.
     * 
     */
    public Optional<Output<String>> liveStageEtag() {
        return Optional.ofNullable(this.liveStageEtag);
    }

    /**
     * Unique name for your CloudFront Function.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Unique name for your CloudFront Function.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
     * 
     */
    @Import(name="publish")
    private @Nullable Output<Boolean> publish;

    /**
     * @return Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> publish() {
        return Optional.ofNullable(this.publish);
    }

    /**
     * Identifier of the function&#39;s runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="runtime")
    private @Nullable Output<String> runtime;

    /**
     * @return Identifier of the function&#39;s runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> runtime() {
        return Optional.ofNullable(this.runtime);
    }

    /**
     * Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private FunctionState() {}

    private FunctionState(FunctionState $) {
        this.arn = $.arn;
        this.code = $.code;
        this.comment = $.comment;
        this.etag = $.etag;
        this.liveStageEtag = $.liveStageEtag;
        this.name = $.name;
        this.publish = $.publish;
        this.runtime = $.runtime;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionState $;

        public Builder() {
            $ = new FunctionState();
        }

        public Builder(FunctionState defaults) {
            $ = new FunctionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) identifying your CloudFront Function.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) identifying your CloudFront Function.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param code Source code of the function
         * 
         * @return builder
         * 
         */
        public Builder code(@Nullable Output<String> code) {
            $.code = code;
            return this;
        }

        /**
         * @param code Source code of the function
         * 
         * @return builder
         * 
         */
        public Builder code(String code) {
            return code(Output.of(code));
        }

        /**
         * @param comment Comment.
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment Comment.
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param etag ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag ETag hash of the function. This is the value for the `DEVELOPMENT` stage of the function.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param liveStageEtag ETag hash of any `LIVE` stage of the function.
         * 
         * @return builder
         * 
         */
        public Builder liveStageEtag(@Nullable Output<String> liveStageEtag) {
            $.liveStageEtag = liveStageEtag;
            return this;
        }

        /**
         * @param liveStageEtag ETag hash of any `LIVE` stage of the function.
         * 
         * @return builder
         * 
         */
        public Builder liveStageEtag(String liveStageEtag) {
            return liveStageEtag(Output.of(liveStageEtag));
        }

        /**
         * @param name Unique name for your CloudFront Function.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Unique name for your CloudFront Function.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param publish Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder publish(@Nullable Output<Boolean> publish) {
            $.publish = publish;
            return this;
        }

        /**
         * @param publish Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder publish(Boolean publish) {
            return publish(Output.of(publish));
        }

        /**
         * @param runtime Identifier of the function&#39;s runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder runtime(@Nullable Output<String> runtime) {
            $.runtime = runtime;
            return this;
        }

        /**
         * @param runtime Identifier of the function&#39;s runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder runtime(String runtime) {
            return runtime(Output.of(runtime));
        }

        /**
         * @param status Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public FunctionState build() {
            return $;
        }
    }

}
