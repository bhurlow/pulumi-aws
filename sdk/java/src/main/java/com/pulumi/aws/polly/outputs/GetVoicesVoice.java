// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.polly.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetVoicesVoice {
    /**
     * @return Additional codes for languages available for the specified voice in addition to its default language.
     * 
     */
    private List<String> additionalLanguageCodes;
    /**
     * @return Gender of the voice.
     * 
     */
    private String gender;
    /**
     * @return Amazon Polly assigned voice ID.
     * 
     */
    private String id;
    /**
     * @return Language identification tag for filtering the list of voices returned. If not specified, all available voices are returned.
     * 
     */
    private String languageCode;
    /**
     * @return Human readable name of the language in English.
     * 
     */
    private String languageName;
    /**
     * @return Name of the voice.
     * 
     */
    private String name;
    /**
     * @return Specifies which engines are supported by a given voice.
     * 
     */
    private List<String> supportedEngines;

    private GetVoicesVoice() {}
    /**
     * @return Additional codes for languages available for the specified voice in addition to its default language.
     * 
     */
    public List<String> additionalLanguageCodes() {
        return this.additionalLanguageCodes;
    }
    /**
     * @return Gender of the voice.
     * 
     */
    public String gender() {
        return this.gender;
    }
    /**
     * @return Amazon Polly assigned voice ID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Language identification tag for filtering the list of voices returned. If not specified, all available voices are returned.
     * 
     */
    public String languageCode() {
        return this.languageCode;
    }
    /**
     * @return Human readable name of the language in English.
     * 
     */
    public String languageName() {
        return this.languageName;
    }
    /**
     * @return Name of the voice.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Specifies which engines are supported by a given voice.
     * 
     */
    public List<String> supportedEngines() {
        return this.supportedEngines;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVoicesVoice defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> additionalLanguageCodes;
        private String gender;
        private String id;
        private String languageCode;
        private String languageName;
        private String name;
        private List<String> supportedEngines;
        public Builder() {}
        public Builder(GetVoicesVoice defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.additionalLanguageCodes = defaults.additionalLanguageCodes;
    	      this.gender = defaults.gender;
    	      this.id = defaults.id;
    	      this.languageCode = defaults.languageCode;
    	      this.languageName = defaults.languageName;
    	      this.name = defaults.name;
    	      this.supportedEngines = defaults.supportedEngines;
        }

        @CustomType.Setter
        public Builder additionalLanguageCodes(List<String> additionalLanguageCodes) {
            this.additionalLanguageCodes = Objects.requireNonNull(additionalLanguageCodes);
            return this;
        }
        public Builder additionalLanguageCodes(String... additionalLanguageCodes) {
            return additionalLanguageCodes(List.of(additionalLanguageCodes));
        }
        @CustomType.Setter
        public Builder gender(String gender) {
            this.gender = Objects.requireNonNull(gender);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder languageCode(String languageCode) {
            this.languageCode = Objects.requireNonNull(languageCode);
            return this;
        }
        @CustomType.Setter
        public Builder languageName(String languageName) {
            this.languageName = Objects.requireNonNull(languageName);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder supportedEngines(List<String> supportedEngines) {
            this.supportedEngines = Objects.requireNonNull(supportedEngines);
            return this;
        }
        public Builder supportedEngines(String... supportedEngines) {
            return supportedEngines(List.of(supportedEngines));
        }
        public GetVoicesVoice build() {
            final var _resultValue = new GetVoicesVoice();
            _resultValue.additionalLanguageCodes = additionalLanguageCodes;
            _resultValue.gender = gender;
            _resultValue.id = id;
            _resultValue.languageCode = languageCode;
            _resultValue.languageName = languageName;
            _resultValue.name = name;
            _resultValue.supportedEngines = supportedEngines;
            return _resultValue;
        }
    }
}
