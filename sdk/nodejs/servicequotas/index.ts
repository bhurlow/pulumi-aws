// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { GetServiceQuotaArgs, GetServiceQuotaResult, GetServiceQuotaOutputArgs } from "./getServiceQuota";
export const getServiceQuota: typeof import("./getServiceQuota").getServiceQuota = null as any;
export const getServiceQuotaOutput: typeof import("./getServiceQuota").getServiceQuotaOutput = null as any;
utilities.lazyLoad(exports, ["getServiceQuota","getServiceQuotaOutput"], () => require("./getServiceQuota"));

export { GetTemplatesArgs, GetTemplatesResult, GetTemplatesOutputArgs } from "./getTemplates";
export const getTemplates: typeof import("./getTemplates").getTemplates = null as any;
export const getTemplatesOutput: typeof import("./getTemplates").getTemplatesOutput = null as any;
utilities.lazyLoad(exports, ["getTemplates","getTemplatesOutput"], () => require("./getTemplates"));

export { ServiceQuotaArgs, ServiceQuotaState } from "./serviceQuota";
export type ServiceQuota = import("./serviceQuota").ServiceQuota;
export const ServiceQuota: typeof import("./serviceQuota").ServiceQuota = null as any;
utilities.lazyLoad(exports, ["ServiceQuota"], () => require("./serviceQuota"));

export { TemplateArgs, TemplateState } from "./template";
export type Template = import("./template").Template;
export const Template: typeof import("./template").Template = null as any;
utilities.lazyLoad(exports, ["Template"], () => require("./template"));

export { TemplateAssociationArgs, TemplateAssociationState } from "./templateAssociation";
export type TemplateAssociation = import("./templateAssociation").TemplateAssociation;
export const TemplateAssociation: typeof import("./templateAssociation").TemplateAssociation = null as any;
utilities.lazyLoad(exports, ["TemplateAssociation"], () => require("./templateAssociation"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:servicequotas/serviceQuota:ServiceQuota":
                return new ServiceQuota(name, <any>undefined, { urn })
            case "aws:servicequotas/template:Template":
                return new Template(name, <any>undefined, { urn })
            case "aws:servicequotas/templateAssociation:TemplateAssociation":
                return new TemplateAssociation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "servicequotas/serviceQuota", _module)
pulumi.runtime.registerResourceModule("aws", "servicequotas/template", _module)
pulumi.runtime.registerResourceModule("aws", "servicequotas/templateAssociation", _module)
