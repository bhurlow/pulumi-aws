// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BucketArgs, BucketState } from "./bucket";
export type Bucket = import("./bucket").Bucket;
export const Bucket: typeof import("./bucket").Bucket = null as any;
utilities.lazyLoad(exports, ["Bucket"], () => require("./bucket"));

export { BucketAccessKeyArgs, BucketAccessKeyState } from "./bucketAccessKey";
export type BucketAccessKey = import("./bucketAccessKey").BucketAccessKey;
export const BucketAccessKey: typeof import("./bucketAccessKey").BucketAccessKey = null as any;
utilities.lazyLoad(exports, ["BucketAccessKey"], () => require("./bucketAccessKey"));

export { CertificateArgs, CertificateState } from "./certificate";
export type Certificate = import("./certificate").Certificate;
export const Certificate: typeof import("./certificate").Certificate = null as any;
utilities.lazyLoad(exports, ["Certificate"], () => require("./certificate"));

export { ContainerServiceArgs, ContainerServiceState } from "./containerService";
export type ContainerService = import("./containerService").ContainerService;
export const ContainerService: typeof import("./containerService").ContainerService = null as any;
utilities.lazyLoad(exports, ["ContainerService"], () => require("./containerService"));

export { ContainerServiceDeploymentVersionArgs, ContainerServiceDeploymentVersionState } from "./containerServiceDeploymentVersion";
export type ContainerServiceDeploymentVersion = import("./containerServiceDeploymentVersion").ContainerServiceDeploymentVersion;
export const ContainerServiceDeploymentVersion: typeof import("./containerServiceDeploymentVersion").ContainerServiceDeploymentVersion = null as any;
utilities.lazyLoad(exports, ["ContainerServiceDeploymentVersion"], () => require("./containerServiceDeploymentVersion"));

export { DatabaseArgs, DatabaseState } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { DiskArgs, DiskState } from "./disk";
export type Disk = import("./disk").Disk;
export const Disk: typeof import("./disk").Disk = null as any;
utilities.lazyLoad(exports, ["Disk"], () => require("./disk"));

export { Disk_attachmentArgs, Disk_attachmentState } from "./disk_attachment";
export type Disk_attachment = import("./disk_attachment").Disk_attachment;
export const Disk_attachment: typeof import("./disk_attachment").Disk_attachment = null as any;
utilities.lazyLoad(exports, ["Disk_attachment"], () => require("./disk_attachment"));

export { DomainArgs, DomainState } from "./domain";
export type Domain = import("./domain").Domain;
export const Domain: typeof import("./domain").Domain = null as any;
utilities.lazyLoad(exports, ["Domain"], () => require("./domain"));

export { DomainEntryArgs, DomainEntryState } from "./domainEntry";
export type DomainEntry = import("./domainEntry").DomainEntry;
export const DomainEntry: typeof import("./domainEntry").DomainEntry = null as any;
utilities.lazyLoad(exports, ["DomainEntry"], () => require("./domainEntry"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstancePublicPortsArgs, InstancePublicPortsState } from "./instancePublicPorts";
export type InstancePublicPorts = import("./instancePublicPorts").InstancePublicPorts;
export const InstancePublicPorts: typeof import("./instancePublicPorts").InstancePublicPorts = null as any;
utilities.lazyLoad(exports, ["InstancePublicPorts"], () => require("./instancePublicPorts"));

export { KeyPairArgs, KeyPairState } from "./keyPair";
export type KeyPair = import("./keyPair").KeyPair;
export const KeyPair: typeof import("./keyPair").KeyPair = null as any;
utilities.lazyLoad(exports, ["KeyPair"], () => require("./keyPair"));

export { LbArgs, LbState } from "./lb";
export type Lb = import("./lb").Lb;
export const Lb: typeof import("./lb").Lb = null as any;
utilities.lazyLoad(exports, ["Lb"], () => require("./lb"));

export { LbAttachmentArgs, LbAttachmentState } from "./lbAttachment";
export type LbAttachment = import("./lbAttachment").LbAttachment;
export const LbAttachment: typeof import("./lbAttachment").LbAttachment = null as any;
utilities.lazyLoad(exports, ["LbAttachment"], () => require("./lbAttachment"));

export { LbCertificateArgs, LbCertificateState } from "./lbCertificate";
export type LbCertificate = import("./lbCertificate").LbCertificate;
export const LbCertificate: typeof import("./lbCertificate").LbCertificate = null as any;
utilities.lazyLoad(exports, ["LbCertificate"], () => require("./lbCertificate"));

export { LbCertificateAttachmentArgs, LbCertificateAttachmentState } from "./lbCertificateAttachment";
export type LbCertificateAttachment = import("./lbCertificateAttachment").LbCertificateAttachment;
export const LbCertificateAttachment: typeof import("./lbCertificateAttachment").LbCertificateAttachment = null as any;
utilities.lazyLoad(exports, ["LbCertificateAttachment"], () => require("./lbCertificateAttachment"));

export { LbHttpsRedirectionPolicyArgs, LbHttpsRedirectionPolicyState } from "./lbHttpsRedirectionPolicy";
export type LbHttpsRedirectionPolicy = import("./lbHttpsRedirectionPolicy").LbHttpsRedirectionPolicy;
export const LbHttpsRedirectionPolicy: typeof import("./lbHttpsRedirectionPolicy").LbHttpsRedirectionPolicy = null as any;
utilities.lazyLoad(exports, ["LbHttpsRedirectionPolicy"], () => require("./lbHttpsRedirectionPolicy"));

export { LbStickinessPolicyArgs, LbStickinessPolicyState } from "./lbStickinessPolicy";
export type LbStickinessPolicy = import("./lbStickinessPolicy").LbStickinessPolicy;
export const LbStickinessPolicy: typeof import("./lbStickinessPolicy").LbStickinessPolicy = null as any;
utilities.lazyLoad(exports, ["LbStickinessPolicy"], () => require("./lbStickinessPolicy"));

export { StaticIpArgs, StaticIpState } from "./staticIp";
export type StaticIp = import("./staticIp").StaticIp;
export const StaticIp: typeof import("./staticIp").StaticIp = null as any;
utilities.lazyLoad(exports, ["StaticIp"], () => require("./staticIp"));

export { StaticIpAttachmentArgs, StaticIpAttachmentState } from "./staticIpAttachment";
export type StaticIpAttachment = import("./staticIpAttachment").StaticIpAttachment;
export const StaticIpAttachment: typeof import("./staticIpAttachment").StaticIpAttachment = null as any;
utilities.lazyLoad(exports, ["StaticIpAttachment"], () => require("./staticIpAttachment"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:lightsail/bucket:Bucket":
                return new Bucket(name, <any>undefined, { urn })
            case "aws:lightsail/bucketAccessKey:BucketAccessKey":
                return new BucketAccessKey(name, <any>undefined, { urn })
            case "aws:lightsail/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "aws:lightsail/containerService:ContainerService":
                return new ContainerService(name, <any>undefined, { urn })
            case "aws:lightsail/containerServiceDeploymentVersion:ContainerServiceDeploymentVersion":
                return new ContainerServiceDeploymentVersion(name, <any>undefined, { urn })
            case "aws:lightsail/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "aws:lightsail/disk:Disk":
                return new Disk(name, <any>undefined, { urn })
            case "aws:lightsail/disk_attachment:Disk_attachment":
                return new Disk_attachment(name, <any>undefined, { urn })
            case "aws:lightsail/domain:Domain":
                return new Domain(name, <any>undefined, { urn })
            case "aws:lightsail/domainEntry:DomainEntry":
                return new DomainEntry(name, <any>undefined, { urn })
            case "aws:lightsail/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "aws:lightsail/instancePublicPorts:InstancePublicPorts":
                return new InstancePublicPorts(name, <any>undefined, { urn })
            case "aws:lightsail/keyPair:KeyPair":
                return new KeyPair(name, <any>undefined, { urn })
            case "aws:lightsail/lb:Lb":
                return new Lb(name, <any>undefined, { urn })
            case "aws:lightsail/lbAttachment:LbAttachment":
                return new LbAttachment(name, <any>undefined, { urn })
            case "aws:lightsail/lbCertificate:LbCertificate":
                return new LbCertificate(name, <any>undefined, { urn })
            case "aws:lightsail/lbCertificateAttachment:LbCertificateAttachment":
                return new LbCertificateAttachment(name, <any>undefined, { urn })
            case "aws:lightsail/lbHttpsRedirectionPolicy:LbHttpsRedirectionPolicy":
                return new LbHttpsRedirectionPolicy(name, <any>undefined, { urn })
            case "aws:lightsail/lbStickinessPolicy:LbStickinessPolicy":
                return new LbStickinessPolicy(name, <any>undefined, { urn })
            case "aws:lightsail/staticIp:StaticIp":
                return new StaticIp(name, <any>undefined, { urn })
            case "aws:lightsail/staticIpAttachment:StaticIpAttachment":
                return new StaticIpAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "lightsail/bucket", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/bucketAccessKey", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/certificate", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/containerService", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/containerServiceDeploymentVersion", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/database", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/disk", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/disk_attachment", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/domain", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/domainEntry", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/instance", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/instancePublicPorts", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/keyPair", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/lb", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/lbAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/lbCertificate", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/lbCertificateAttachment", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/lbHttpsRedirectionPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/lbStickinessPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/staticIp", _module)
pulumi.runtime.registerResourceModule("aws", "lightsail/staticIpAttachment", _module)
