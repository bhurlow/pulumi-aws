// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./certificate";
export * from "./certificateAuthority";
export * from "./certificateAuthorityCertificate";
export * from "./getCertificate";
export * from "./getCertificateAuthority";
export * from "./policy";

// Import resources to register:
import { Certificate } from "./certificate";
import { CertificateAuthority } from "./certificateAuthority";
import { CertificateAuthorityCertificate } from "./certificateAuthorityCertificate";
import { Policy } from "./policy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:acmpca/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "aws:acmpca/certificateAuthority:CertificateAuthority":
                return new CertificateAuthority(name, <any>undefined, { urn })
            case "aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate":
                return new CertificateAuthorityCertificate(name, <any>undefined, { urn })
            case "aws:acmpca/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "acmpca/certificate", _module)
pulumi.runtime.registerResourceModule("aws", "acmpca/certificateAuthority", _module)
pulumi.runtime.registerResourceModule("aws", "acmpca/certificateAuthorityCertificate", _module)
pulumi.runtime.registerResourceModule("aws", "acmpca/policy", _module)
