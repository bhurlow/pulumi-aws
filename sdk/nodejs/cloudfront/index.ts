// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cachePolicy";
export * from "./distribution";
export * from "./fieldLevelEncryptionConfig";
export * from "./fieldLevelEncryptionProfile";
export * from "./function";
export * from "./getCachePolicy";
export * from "./getDistribution";
export * from "./getFunction";
export * from "./getLogDeliveryCanonicalUserId";
export * from "./getOriginAccessIdentities";
export * from "./getOriginAccessIdentity";
export * from "./getOriginRequestPolicy";
export * from "./getRealtimeLogConfig";
export * from "./getResponseHeadersPolicy";
export * from "./keyGroup";
export * from "./monitoringSubscription";
export * from "./originAccessIdentity";
export * from "./originRequestPolicy";
export * from "./publicKey";
export * from "./realtimeLogConfig";
export * from "./responseHeadersPolicy";

// Import resources to register:
import { CachePolicy } from "./cachePolicy";
import { Distribution } from "./distribution";
import { FieldLevelEncryptionConfig } from "./fieldLevelEncryptionConfig";
import { FieldLevelEncryptionProfile } from "./fieldLevelEncryptionProfile";
import { Function } from "./function";
import { KeyGroup } from "./keyGroup";
import { MonitoringSubscription } from "./monitoringSubscription";
import { OriginAccessIdentity } from "./originAccessIdentity";
import { OriginRequestPolicy } from "./originRequestPolicy";
import { PublicKey } from "./publicKey";
import { RealtimeLogConfig } from "./realtimeLogConfig";
import { ResponseHeadersPolicy } from "./responseHeadersPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:cloudfront/cachePolicy:CachePolicy":
                return new CachePolicy(name, <any>undefined, { urn })
            case "aws:cloudfront/distribution:Distribution":
                return new Distribution(name, <any>undefined, { urn })
            case "aws:cloudfront/fieldLevelEncryptionConfig:FieldLevelEncryptionConfig":
                return new FieldLevelEncryptionConfig(name, <any>undefined, { urn })
            case "aws:cloudfront/fieldLevelEncryptionProfile:FieldLevelEncryptionProfile":
                return new FieldLevelEncryptionProfile(name, <any>undefined, { urn })
            case "aws:cloudfront/function:Function":
                return new Function(name, <any>undefined, { urn })
            case "aws:cloudfront/keyGroup:KeyGroup":
                return new KeyGroup(name, <any>undefined, { urn })
            case "aws:cloudfront/monitoringSubscription:MonitoringSubscription":
                return new MonitoringSubscription(name, <any>undefined, { urn })
            case "aws:cloudfront/originAccessIdentity:OriginAccessIdentity":
                return new OriginAccessIdentity(name, <any>undefined, { urn })
            case "aws:cloudfront/originRequestPolicy:OriginRequestPolicy":
                return new OriginRequestPolicy(name, <any>undefined, { urn })
            case "aws:cloudfront/publicKey:PublicKey":
                return new PublicKey(name, <any>undefined, { urn })
            case "aws:cloudfront/realtimeLogConfig:RealtimeLogConfig":
                return new RealtimeLogConfig(name, <any>undefined, { urn })
            case "aws:cloudfront/responseHeadersPolicy:ResponseHeadersPolicy":
                return new ResponseHeadersPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "cloudfront/cachePolicy", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/distribution", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/fieldLevelEncryptionConfig", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/fieldLevelEncryptionProfile", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/function", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/keyGroup", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/monitoringSubscription", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/originAccessIdentity", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/originRequestPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/publicKey", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/realtimeLogConfig", _module)
pulumi.runtime.registerResourceModule("aws", "cloudfront/responseHeadersPolicy", _module)
