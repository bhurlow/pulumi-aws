// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getHostedZoneId";
export * from "./getListener";
export * from "./getLoadBalancer";
export * from "./getTargetGroup";
export * from "./listener";
export * from "./listenerCertificate";
export * from "./listenerRule";
export * from "./loadBalancer";
export * from "./targetGroup";
export * from "./targetGroupAttachment";

// Import resources to register:
import { Listener } from "./listener";
import { ListenerCertificate } from "./listenerCertificate";
import { ListenerRule } from "./listenerRule";
import { LoadBalancer } from "./loadBalancer";
import { TargetGroup } from "./targetGroup";
import { TargetGroupAttachment } from "./targetGroupAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:lb/listener:Listener":
                return new Listener(name, <any>undefined, { urn })
            case "aws:lb/listenerCertificate:ListenerCertificate":
                return new ListenerCertificate(name, <any>undefined, { urn })
            case "aws:lb/listenerRule:ListenerRule":
                return new ListenerRule(name, <any>undefined, { urn })
            case "aws:lb/loadBalancer:LoadBalancer":
                return new LoadBalancer(name, <any>undefined, { urn })
            case "aws:lb/targetGroup:TargetGroup":
                return new TargetGroup(name, <any>undefined, { urn })
            case "aws:lb/targetGroupAttachment:TargetGroupAttachment":
                return new TargetGroupAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "lb/listener", _module)
pulumi.runtime.registerResourceModule("aws", "lb/listenerCertificate", _module)
pulumi.runtime.registerResourceModule("aws", "lb/listenerRule", _module)
pulumi.runtime.registerResourceModule("aws", "lb/loadBalancer", _module)
pulumi.runtime.registerResourceModule("aws", "lb/targetGroup", _module)
pulumi.runtime.registerResourceModule("aws", "lb/targetGroupAttachment", _module)
