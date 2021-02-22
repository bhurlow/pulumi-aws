// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cloudwatchMixins";
export * from "./compositeAlarm";
export * from "./dashboard";
export * from "./eventArchive";
export * from "./eventBus";
export * from "./eventPermission";
export * from "./eventRule";
export * from "./eventRuleMixins";
export * from "./eventTarget";
export * from "./getLogGroup";
export * from "./logDestination";
export * from "./logDestinationPolicy";
export * from "./logGroup";
export * from "./logGroupMixins";
export * from "./logMetricFilter";
export * from "./logResourcePolicy";
export * from "./logStream";
export * from "./logSubscriptionFilter";
export * from "./metricAlarm";

// Import resources to register:
import { CompositeAlarm } from "./compositeAlarm";
import { Dashboard } from "./dashboard";
import { EventArchive } from "./eventArchive";
import { EventBus } from "./eventBus";
import { EventPermission } from "./eventPermission";
import { EventRule } from "./eventRule";
import { EventTarget } from "./eventTarget";
import { LogDestination } from "./logDestination";
import { LogDestinationPolicy } from "./logDestinationPolicy";
import { LogGroup } from "./logGroup";
import { LogMetricFilter } from "./logMetricFilter";
import { LogResourcePolicy } from "./logResourcePolicy";
import { LogStream } from "./logStream";
import { LogSubscriptionFilter } from "./logSubscriptionFilter";
import { MetricAlarm } from "./metricAlarm";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:cloudwatch/compositeAlarm:CompositeAlarm":
                return new CompositeAlarm(name, <any>undefined, { urn })
            case "aws:cloudwatch/dashboard:Dashboard":
                return new Dashboard(name, <any>undefined, { urn })
            case "aws:cloudwatch/eventArchive:EventArchive":
                return new EventArchive(name, <any>undefined, { urn })
            case "aws:cloudwatch/eventBus:EventBus":
                return new EventBus(name, <any>undefined, { urn })
            case "aws:cloudwatch/eventPermission:EventPermission":
                return new EventPermission(name, <any>undefined, { urn })
            case "aws:cloudwatch/eventRule:EventRule":
                return new EventRule(name, <any>undefined, { urn })
            case "aws:cloudwatch/eventTarget:EventTarget":
                return new EventTarget(name, <any>undefined, { urn })
            case "aws:cloudwatch/logDestination:LogDestination":
                return new LogDestination(name, <any>undefined, { urn })
            case "aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy":
                return new LogDestinationPolicy(name, <any>undefined, { urn })
            case "aws:cloudwatch/logGroup:LogGroup":
                return new LogGroup(name, <any>undefined, { urn })
            case "aws:cloudwatch/logMetricFilter:LogMetricFilter":
                return new LogMetricFilter(name, <any>undefined, { urn })
            case "aws:cloudwatch/logResourcePolicy:LogResourcePolicy":
                return new LogResourcePolicy(name, <any>undefined, { urn })
            case "aws:cloudwatch/logStream:LogStream":
                return new LogStream(name, <any>undefined, { urn })
            case "aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter":
                return new LogSubscriptionFilter(name, <any>undefined, { urn })
            case "aws:cloudwatch/metricAlarm:MetricAlarm":
                return new MetricAlarm(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "cloudwatch/compositeAlarm", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/dashboard", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/eventArchive", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/eventBus", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/eventPermission", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/eventRule", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/eventTarget", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/logDestination", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/logDestinationPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/logGroup", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/logMetricFilter", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/logResourcePolicy", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/logStream", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/logSubscriptionFilter", _module)
pulumi.runtime.registerResourceModule("aws", "cloudwatch/metricAlarm", _module)
