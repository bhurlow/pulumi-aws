// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Inputs
{

    public sealed class ProviderEndpointArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessanalyzer")]
        public Input<string>? Accessanalyzer { get; set; }

        [Input("account")]
        public Input<string>? Account { get; set; }

        [Input("acm")]
        public Input<string>? Acm { get; set; }

        [Input("acmpca")]
        public Input<string>? Acmpca { get; set; }

        [Input("alexaforbusiness")]
        public Input<string>? Alexaforbusiness { get; set; }

        [Input("amg")]
        public Input<string>? Amg { get; set; }

        [Input("amp")]
        public Input<string>? Amp { get; set; }

        [Input("amplify")]
        public Input<string>? Amplify { get; set; }

        [Input("amplifybackend")]
        public Input<string>? Amplifybackend { get; set; }

        [Input("amplifyuibuilder")]
        public Input<string>? Amplifyuibuilder { get; set; }

        [Input("apigateway")]
        public Input<string>? Apigateway { get; set; }

        [Input("apigatewaymanagementapi")]
        public Input<string>? Apigatewaymanagementapi { get; set; }

        [Input("apigatewayv2")]
        public Input<string>? Apigatewayv2 { get; set; }

        [Input("appautoscaling")]
        public Input<string>? Appautoscaling { get; set; }

        [Input("appconfig")]
        public Input<string>? Appconfig { get; set; }

        [Input("appconfigdata")]
        public Input<string>? Appconfigdata { get; set; }

        [Input("appflow")]
        public Input<string>? Appflow { get; set; }

        [Input("appintegrations")]
        public Input<string>? Appintegrations { get; set; }

        [Input("appintegrationsservice")]
        public Input<string>? Appintegrationsservice { get; set; }

        [Input("applicationautoscaling")]
        public Input<string>? Applicationautoscaling { get; set; }

        [Input("applicationcostprofiler")]
        public Input<string>? Applicationcostprofiler { get; set; }

        [Input("applicationdiscovery")]
        public Input<string>? Applicationdiscovery { get; set; }

        [Input("applicationdiscoveryservice")]
        public Input<string>? Applicationdiscoveryservice { get; set; }

        [Input("applicationinsights")]
        public Input<string>? Applicationinsights { get; set; }

        [Input("appmesh")]
        public Input<string>? Appmesh { get; set; }

        [Input("appregistry")]
        public Input<string>? Appregistry { get; set; }

        [Input("apprunner")]
        public Input<string>? Apprunner { get; set; }

        [Input("appstream")]
        public Input<string>? Appstream { get; set; }

        [Input("appsync")]
        public Input<string>? Appsync { get; set; }

        [Input("athena")]
        public Input<string>? Athena { get; set; }

        [Input("auditmanager")]
        public Input<string>? Auditmanager { get; set; }

        [Input("augmentedairuntime")]
        public Input<string>? Augmentedairuntime { get; set; }

        [Input("autoscaling")]
        public Input<string>? Autoscaling { get; set; }

        [Input("autoscalingplans")]
        public Input<string>? Autoscalingplans { get; set; }

        [Input("backup")]
        public Input<string>? Backup { get; set; }

        [Input("backupgateway")]
        public Input<string>? Backupgateway { get; set; }

        [Input("batch")]
        public Input<string>? Batch { get; set; }

        [Input("beanstalk")]
        public Input<string>? Beanstalk { get; set; }

        [Input("billingconductor")]
        public Input<string>? Billingconductor { get; set; }

        [Input("braket")]
        public Input<string>? Braket { get; set; }

        [Input("budgets")]
        public Input<string>? Budgets { get; set; }

        [Input("ce")]
        public Input<string>? Ce { get; set; }

        [Input("chime")]
        public Input<string>? Chime { get; set; }

        [Input("chimesdkidentity")]
        public Input<string>? Chimesdkidentity { get; set; }

        [Input("chimesdkmeetings")]
        public Input<string>? Chimesdkmeetings { get; set; }

        [Input("chimesdkmessaging")]
        public Input<string>? Chimesdkmessaging { get; set; }

        [Input("cloud9")]
        public Input<string>? Cloud9 { get; set; }

        [Input("cloudcontrol")]
        public Input<string>? Cloudcontrol { get; set; }

        [Input("cloudcontrolapi")]
        public Input<string>? Cloudcontrolapi { get; set; }

        [Input("clouddirectory")]
        public Input<string>? Clouddirectory { get; set; }

        [Input("cloudformation")]
        public Input<string>? Cloudformation { get; set; }

        [Input("cloudfront")]
        public Input<string>? Cloudfront { get; set; }

        [Input("cloudhsm")]
        public Input<string>? Cloudhsm { get; set; }

        [Input("cloudhsmv2")]
        public Input<string>? Cloudhsmv2 { get; set; }

        [Input("cloudsearch")]
        public Input<string>? Cloudsearch { get; set; }

        [Input("cloudsearchdomain")]
        public Input<string>? Cloudsearchdomain { get; set; }

        [Input("cloudtrail")]
        public Input<string>? Cloudtrail { get; set; }

        [Input("cloudwatch")]
        public Input<string>? Cloudwatch { get; set; }

        [Input("cloudwatchevents")]
        public Input<string>? Cloudwatchevents { get; set; }

        [Input("cloudwatchevidently")]
        public Input<string>? Cloudwatchevidently { get; set; }

        [Input("cloudwatchlog")]
        public Input<string>? Cloudwatchlog { get; set; }

        [Input("cloudwatchlogs")]
        public Input<string>? Cloudwatchlogs { get; set; }

        [Input("cloudwatchrum")]
        public Input<string>? Cloudwatchrum { get; set; }

        [Input("codeartifact")]
        public Input<string>? Codeartifact { get; set; }

        [Input("codebuild")]
        public Input<string>? Codebuild { get; set; }

        [Input("codecommit")]
        public Input<string>? Codecommit { get; set; }

        [Input("codedeploy")]
        public Input<string>? Codedeploy { get; set; }

        [Input("codeguruprofiler")]
        public Input<string>? Codeguruprofiler { get; set; }

        [Input("codegurureviewer")]
        public Input<string>? Codegurureviewer { get; set; }

        [Input("codepipeline")]
        public Input<string>? Codepipeline { get; set; }

        [Input("codestar")]
        public Input<string>? Codestar { get; set; }

        [Input("codestarconnections")]
        public Input<string>? Codestarconnections { get; set; }

        [Input("codestarnotifications")]
        public Input<string>? Codestarnotifications { get; set; }

        [Input("cognitoidentity")]
        public Input<string>? Cognitoidentity { get; set; }

        [Input("cognitoidentityprovider")]
        public Input<string>? Cognitoidentityprovider { get; set; }

        [Input("cognitoidp")]
        public Input<string>? Cognitoidp { get; set; }

        [Input("cognitosync")]
        public Input<string>? Cognitosync { get; set; }

        [Input("comprehend")]
        public Input<string>? Comprehend { get; set; }

        [Input("comprehendmedical")]
        public Input<string>? Comprehendmedical { get; set; }

        [Input("computeoptimizer")]
        public Input<string>? Computeoptimizer { get; set; }

        [Input("config")]
        public Input<string>? Config { get; set; }

        [Input("configservice")]
        public Input<string>? Configservice { get; set; }

        [Input("connect")]
        public Input<string>? Connect { get; set; }

        [Input("connectcontactlens")]
        public Input<string>? Connectcontactlens { get; set; }

        [Input("connectparticipant")]
        public Input<string>? Connectparticipant { get; set; }

        [Input("connectwisdomservice")]
        public Input<string>? Connectwisdomservice { get; set; }

        [Input("controltower")]
        public Input<string>? Controltower { get; set; }

        [Input("costandusagereportservice")]
        public Input<string>? Costandusagereportservice { get; set; }

        [Input("costexplorer")]
        public Input<string>? Costexplorer { get; set; }

        [Input("cur")]
        public Input<string>? Cur { get; set; }

        [Input("customerprofiles")]
        public Input<string>? Customerprofiles { get; set; }

        [Input("databasemigration")]
        public Input<string>? Databasemigration { get; set; }

        [Input("databasemigrationservice")]
        public Input<string>? Databasemigrationservice { get; set; }

        [Input("databrew")]
        public Input<string>? Databrew { get; set; }

        [Input("dataexchange")]
        public Input<string>? Dataexchange { get; set; }

        [Input("datapipeline")]
        public Input<string>? Datapipeline { get; set; }

        [Input("datasync")]
        public Input<string>? Datasync { get; set; }

        [Input("dax")]
        public Input<string>? Dax { get; set; }

        [Input("deploy")]
        public Input<string>? Deploy { get; set; }

        [Input("detective")]
        public Input<string>? Detective { get; set; }

        [Input("devicefarm")]
        public Input<string>? Devicefarm { get; set; }

        [Input("devopsguru")]
        public Input<string>? Devopsguru { get; set; }

        [Input("directconnect")]
        public Input<string>? Directconnect { get; set; }

        [Input("directoryservice")]
        public Input<string>? Directoryservice { get; set; }

        [Input("discovery")]
        public Input<string>? Discovery { get; set; }

        [Input("dlm")]
        public Input<string>? Dlm { get; set; }

        [Input("dms")]
        public Input<string>? Dms { get; set; }

        [Input("docdb")]
        public Input<string>? Docdb { get; set; }

        [Input("drs")]
        public Input<string>? Drs { get; set; }

        [Input("ds")]
        public Input<string>? Ds { get; set; }

        [Input("dynamodb")]
        public Input<string>? Dynamodb { get; set; }

        [Input("dynamodbstreams")]
        public Input<string>? Dynamodbstreams { get; set; }

        [Input("ebs")]
        public Input<string>? Ebs { get; set; }

        [Input("ec2")]
        public Input<string>? Ec2 { get; set; }

        [Input("ec2instanceconnect")]
        public Input<string>? Ec2instanceconnect { get; set; }

        [Input("ecr")]
        public Input<string>? Ecr { get; set; }

        [Input("ecrpublic")]
        public Input<string>? Ecrpublic { get; set; }

        [Input("ecs")]
        public Input<string>? Ecs { get; set; }

        [Input("efs")]
        public Input<string>? Efs { get; set; }

        [Input("eks")]
        public Input<string>? Eks { get; set; }

        [Input("elasticache")]
        public Input<string>? Elasticache { get; set; }

        [Input("elasticbeanstalk")]
        public Input<string>? Elasticbeanstalk { get; set; }

        [Input("elasticinference")]
        public Input<string>? Elasticinference { get; set; }

        [Input("elasticloadbalancing")]
        public Input<string>? Elasticloadbalancing { get; set; }

        [Input("elasticloadbalancingv2")]
        public Input<string>? Elasticloadbalancingv2 { get; set; }

        [Input("elasticsearch")]
        public Input<string>? Elasticsearch { get; set; }

        [Input("elasticsearchservice")]
        public Input<string>? Elasticsearchservice { get; set; }

        [Input("elastictranscoder")]
        public Input<string>? Elastictranscoder { get; set; }

        [Input("elb")]
        public Input<string>? Elb { get; set; }

        [Input("elbv2")]
        public Input<string>? Elbv2 { get; set; }

        [Input("emr")]
        public Input<string>? Emr { get; set; }

        [Input("emrcontainers")]
        public Input<string>? Emrcontainers { get; set; }

        [Input("emrserverless")]
        public Input<string>? Emrserverless { get; set; }

        [Input("es")]
        public Input<string>? Es { get; set; }

        [Input("eventbridge")]
        public Input<string>? Eventbridge { get; set; }

        [Input("events")]
        public Input<string>? Events { get; set; }

        [Input("evidently")]
        public Input<string>? Evidently { get; set; }

        [Input("finspace")]
        public Input<string>? Finspace { get; set; }

        [Input("finspacedata")]
        public Input<string>? Finspacedata { get; set; }

        [Input("firehose")]
        public Input<string>? Firehose { get; set; }

        [Input("fis")]
        public Input<string>? Fis { get; set; }

        [Input("fms")]
        public Input<string>? Fms { get; set; }

        [Input("forecast")]
        public Input<string>? Forecast { get; set; }

        [Input("forecastquery")]
        public Input<string>? Forecastquery { get; set; }

        [Input("forecastqueryservice")]
        public Input<string>? Forecastqueryservice { get; set; }

        [Input("forecastservice")]
        public Input<string>? Forecastservice { get; set; }

        [Input("frauddetector")]
        public Input<string>? Frauddetector { get; set; }

        [Input("fsx")]
        public Input<string>? Fsx { get; set; }

        [Input("gamelift")]
        public Input<string>? Gamelift { get; set; }

        [Input("glacier")]
        public Input<string>? Glacier { get; set; }

        [Input("globalaccelerator")]
        public Input<string>? Globalaccelerator { get; set; }

        [Input("glue")]
        public Input<string>? Glue { get; set; }

        [Input("gluedatabrew")]
        public Input<string>? Gluedatabrew { get; set; }

        [Input("grafana")]
        public Input<string>? Grafana { get; set; }

        [Input("greengrass")]
        public Input<string>? Greengrass { get; set; }

        [Input("greengrassv2")]
        public Input<string>? Greengrassv2 { get; set; }

        [Input("groundstation")]
        public Input<string>? Groundstation { get; set; }

        [Input("guardduty")]
        public Input<string>? Guardduty { get; set; }

        [Input("health")]
        public Input<string>? Health { get; set; }

        [Input("healthlake")]
        public Input<string>? Healthlake { get; set; }

        [Input("honeycode")]
        public Input<string>? Honeycode { get; set; }

        [Input("iam")]
        public Input<string>? Iam { get; set; }

        [Input("identitystore")]
        public Input<string>? Identitystore { get; set; }

        [Input("imagebuilder")]
        public Input<string>? Imagebuilder { get; set; }

        [Input("inspector")]
        public Input<string>? Inspector { get; set; }

        [Input("inspector2")]
        public Input<string>? Inspector2 { get; set; }

        [Input("inspectorv2")]
        public Input<string>? Inspectorv2 { get; set; }

        [Input("iot")]
        public Input<string>? Iot { get; set; }

        [Input("iot1clickdevices")]
        public Input<string>? Iot1clickdevices { get; set; }

        [Input("iot1clickdevicesservice")]
        public Input<string>? Iot1clickdevicesservice { get; set; }

        [Input("iot1clickprojects")]
        public Input<string>? Iot1clickprojects { get; set; }

        [Input("iotanalytics")]
        public Input<string>? Iotanalytics { get; set; }

        [Input("iotdata")]
        public Input<string>? Iotdata { get; set; }

        [Input("iotdataplane")]
        public Input<string>? Iotdataplane { get; set; }

        [Input("iotdeviceadvisor")]
        public Input<string>? Iotdeviceadvisor { get; set; }

        [Input("iotevents")]
        public Input<string>? Iotevents { get; set; }

        [Input("ioteventsdata")]
        public Input<string>? Ioteventsdata { get; set; }

        [Input("iotfleethub")]
        public Input<string>? Iotfleethub { get; set; }

        [Input("iotjobsdata")]
        public Input<string>? Iotjobsdata { get; set; }

        [Input("iotjobsdataplane")]
        public Input<string>? Iotjobsdataplane { get; set; }

        [Input("iotsecuretunneling")]
        public Input<string>? Iotsecuretunneling { get; set; }

        [Input("iotsitewise")]
        public Input<string>? Iotsitewise { get; set; }

        [Input("iotthingsgraph")]
        public Input<string>? Iotthingsgraph { get; set; }

        [Input("iottwinmaker")]
        public Input<string>? Iottwinmaker { get; set; }

        [Input("iotwireless")]
        public Input<string>? Iotwireless { get; set; }

        [Input("ivs")]
        public Input<string>? Ivs { get; set; }

        [Input("ivschat")]
        public Input<string>? Ivschat { get; set; }

        [Input("kafka")]
        public Input<string>? Kafka { get; set; }

        [Input("kafkaconnect")]
        public Input<string>? Kafkaconnect { get; set; }

        [Input("kendra")]
        public Input<string>? Kendra { get; set; }

        [Input("keyspaces")]
        public Input<string>? Keyspaces { get; set; }

        [Input("kinesis")]
        public Input<string>? Kinesis { get; set; }

        [Input("kinesisanalytics")]
        public Input<string>? Kinesisanalytics { get; set; }

        [Input("kinesisanalyticsv2")]
        public Input<string>? Kinesisanalyticsv2 { get; set; }

        [Input("kinesisvideo")]
        public Input<string>? Kinesisvideo { get; set; }

        [Input("kinesisvideoarchivedmedia")]
        public Input<string>? Kinesisvideoarchivedmedia { get; set; }

        [Input("kinesisvideomedia")]
        public Input<string>? Kinesisvideomedia { get; set; }

        [Input("kinesisvideosignaling")]
        public Input<string>? Kinesisvideosignaling { get; set; }

        [Input("kinesisvideosignalingchannels")]
        public Input<string>? Kinesisvideosignalingchannels { get; set; }

        [Input("kms")]
        public Input<string>? Kms { get; set; }

        [Input("lakeformation")]
        public Input<string>? Lakeformation { get; set; }

        [Input("lambda")]
        public Input<string>? Lambda { get; set; }

        [Input("lex")]
        public Input<string>? Lex { get; set; }

        [Input("lexmodelbuilding")]
        public Input<string>? Lexmodelbuilding { get; set; }

        [Input("lexmodelbuildingservice")]
        public Input<string>? Lexmodelbuildingservice { get; set; }

        [Input("lexmodels")]
        public Input<string>? Lexmodels { get; set; }

        [Input("lexmodelsv2")]
        public Input<string>? Lexmodelsv2 { get; set; }

        [Input("lexruntime")]
        public Input<string>? Lexruntime { get; set; }

        [Input("lexruntimeservice")]
        public Input<string>? Lexruntimeservice { get; set; }

        [Input("lexruntimev2")]
        public Input<string>? Lexruntimev2 { get; set; }

        [Input("lexv2models")]
        public Input<string>? Lexv2models { get; set; }

        [Input("lexv2runtime")]
        public Input<string>? Lexv2runtime { get; set; }

        [Input("licensemanager")]
        public Input<string>? Licensemanager { get; set; }

        [Input("lightsail")]
        public Input<string>? Lightsail { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("locationservice")]
        public Input<string>? Locationservice { get; set; }

        [Input("logs")]
        public Input<string>? Logs { get; set; }

        [Input("lookoutequipment")]
        public Input<string>? Lookoutequipment { get; set; }

        [Input("lookoutforvision")]
        public Input<string>? Lookoutforvision { get; set; }

        [Input("lookoutmetrics")]
        public Input<string>? Lookoutmetrics { get; set; }

        [Input("lookoutvision")]
        public Input<string>? Lookoutvision { get; set; }

        [Input("machinelearning")]
        public Input<string>? Machinelearning { get; set; }

        [Input("macie")]
        public Input<string>? Macie { get; set; }

        [Input("macie2")]
        public Input<string>? Macie2 { get; set; }

        [Input("managedblockchain")]
        public Input<string>? Managedblockchain { get; set; }

        [Input("managedgrafana")]
        public Input<string>? Managedgrafana { get; set; }

        [Input("marketplacecatalog")]
        public Input<string>? Marketplacecatalog { get; set; }

        [Input("marketplacecommerceanalytics")]
        public Input<string>? Marketplacecommerceanalytics { get; set; }

        [Input("marketplaceentitlement")]
        public Input<string>? Marketplaceentitlement { get; set; }

        [Input("marketplaceentitlementservice")]
        public Input<string>? Marketplaceentitlementservice { get; set; }

        [Input("marketplacemetering")]
        public Input<string>? Marketplacemetering { get; set; }

        [Input("mediaconnect")]
        public Input<string>? Mediaconnect { get; set; }

        [Input("mediaconvert")]
        public Input<string>? Mediaconvert { get; set; }

        [Input("medialive")]
        public Input<string>? Medialive { get; set; }

        [Input("mediapackage")]
        public Input<string>? Mediapackage { get; set; }

        [Input("mediapackagevod")]
        public Input<string>? Mediapackagevod { get; set; }

        [Input("mediastore")]
        public Input<string>? Mediastore { get; set; }

        [Input("mediastoredata")]
        public Input<string>? Mediastoredata { get; set; }

        [Input("mediatailor")]
        public Input<string>? Mediatailor { get; set; }

        [Input("memorydb")]
        public Input<string>? Memorydb { get; set; }

        [Input("meteringmarketplace")]
        public Input<string>? Meteringmarketplace { get; set; }

        [Input("mgh")]
        public Input<string>? Mgh { get; set; }

        [Input("mgn")]
        public Input<string>? Mgn { get; set; }

        [Input("migrationhub")]
        public Input<string>? Migrationhub { get; set; }

        [Input("migrationhubconfig")]
        public Input<string>? Migrationhubconfig { get; set; }

        [Input("migrationhubrefactorspaces")]
        public Input<string>? Migrationhubrefactorspaces { get; set; }

        [Input("migrationhubstrategy")]
        public Input<string>? Migrationhubstrategy { get; set; }

        [Input("migrationhubstrategyrecommendations")]
        public Input<string>? Migrationhubstrategyrecommendations { get; set; }

        [Input("mobile")]
        public Input<string>? Mobile { get; set; }

        [Input("mq")]
        public Input<string>? Mq { get; set; }

        [Input("msk")]
        public Input<string>? Msk { get; set; }

        [Input("mturk")]
        public Input<string>? Mturk { get; set; }

        [Input("mwaa")]
        public Input<string>? Mwaa { get; set; }

        [Input("neptune")]
        public Input<string>? Neptune { get; set; }

        [Input("networkfirewall")]
        public Input<string>? Networkfirewall { get; set; }

        [Input("networkmanager")]
        public Input<string>? Networkmanager { get; set; }

        [Input("nimble")]
        public Input<string>? Nimble { get; set; }

        [Input("nimblestudio")]
        public Input<string>? Nimblestudio { get; set; }

        [Input("opensearch")]
        public Input<string>? Opensearch { get; set; }

        [Input("opensearchservice")]
        public Input<string>? Opensearchservice { get; set; }

        [Input("opsworks")]
        public Input<string>? Opsworks { get; set; }

        [Input("opsworkscm")]
        public Input<string>? Opsworkscm { get; set; }

        [Input("organizations")]
        public Input<string>? Organizations { get; set; }

        [Input("outposts")]
        public Input<string>? Outposts { get; set; }

        [Input("panorama")]
        public Input<string>? Panorama { get; set; }

        [Input("personalize")]
        public Input<string>? Personalize { get; set; }

        [Input("personalizeevents")]
        public Input<string>? Personalizeevents { get; set; }

        [Input("personalizeruntime")]
        public Input<string>? Personalizeruntime { get; set; }

        [Input("pi")]
        public Input<string>? Pi { get; set; }

        [Input("pinpoint")]
        public Input<string>? Pinpoint { get; set; }

        [Input("pinpointemail")]
        public Input<string>? Pinpointemail { get; set; }

        [Input("pinpointsmsvoice")]
        public Input<string>? Pinpointsmsvoice { get; set; }

        [Input("polly")]
        public Input<string>? Polly { get; set; }

        [Input("pricing")]
        public Input<string>? Pricing { get; set; }

        [Input("prometheus")]
        public Input<string>? Prometheus { get; set; }

        [Input("prometheusservice")]
        public Input<string>? Prometheusservice { get; set; }

        [Input("proton")]
        public Input<string>? Proton { get; set; }

        [Input("qldb")]
        public Input<string>? Qldb { get; set; }

        [Input("qldbsession")]
        public Input<string>? Qldbsession { get; set; }

        [Input("quicksight")]
        public Input<string>? Quicksight { get; set; }

        [Input("ram")]
        public Input<string>? Ram { get; set; }

        [Input("rbin")]
        public Input<string>? Rbin { get; set; }

        [Input("rds")]
        public Input<string>? Rds { get; set; }

        [Input("rdsdata")]
        public Input<string>? Rdsdata { get; set; }

        [Input("rdsdataservice")]
        public Input<string>? Rdsdataservice { get; set; }

        [Input("recyclebin")]
        public Input<string>? Recyclebin { get; set; }

        [Input("redshift")]
        public Input<string>? Redshift { get; set; }

        [Input("redshiftdata")]
        public Input<string>? Redshiftdata { get; set; }

        [Input("redshiftdataapiservice")]
        public Input<string>? Redshiftdataapiservice { get; set; }

        [Input("redshiftserverless")]
        public Input<string>? Redshiftserverless { get; set; }

        [Input("rekognition")]
        public Input<string>? Rekognition { get; set; }

        [Input("resiliencehub")]
        public Input<string>? Resiliencehub { get; set; }

        [Input("resourceexplorer2")]
        public Input<string>? Resourceexplorer2 { get; set; }

        [Input("resourcegroups")]
        public Input<string>? Resourcegroups { get; set; }

        [Input("resourcegroupstagging")]
        public Input<string>? Resourcegroupstagging { get; set; }

        [Input("resourcegroupstaggingapi")]
        public Input<string>? Resourcegroupstaggingapi { get; set; }

        [Input("robomaker")]
        public Input<string>? Robomaker { get; set; }

        [Input("rolesanywhere")]
        public Input<string>? Rolesanywhere { get; set; }

        [Input("route53")]
        public Input<string>? Route53 { get; set; }

        [Input("route53domains")]
        public Input<string>? Route53domains { get; set; }

        [Input("route53recoverycluster")]
        public Input<string>? Route53recoverycluster { get; set; }

        [Input("route53recoverycontrolconfig")]
        public Input<string>? Route53recoverycontrolconfig { get; set; }

        [Input("route53recoveryreadiness")]
        public Input<string>? Route53recoveryreadiness { get; set; }

        [Input("route53resolver")]
        public Input<string>? Route53resolver { get; set; }

        [Input("rum")]
        public Input<string>? Rum { get; set; }

        [Input("s3")]
        public Input<string>? S3 { get; set; }

        [Input("s3api")]
        public Input<string>? S3api { get; set; }

        [Input("s3control")]
        public Input<string>? S3control { get; set; }

        [Input("s3outposts")]
        public Input<string>? S3outposts { get; set; }

        [Input("sagemaker")]
        public Input<string>? Sagemaker { get; set; }

        [Input("sagemakera2iruntime")]
        public Input<string>? Sagemakera2iruntime { get; set; }

        [Input("sagemakeredge")]
        public Input<string>? Sagemakeredge { get; set; }

        [Input("sagemakeredgemanager")]
        public Input<string>? Sagemakeredgemanager { get; set; }

        [Input("sagemakerfeaturestoreruntime")]
        public Input<string>? Sagemakerfeaturestoreruntime { get; set; }

        [Input("sagemakerruntime")]
        public Input<string>? Sagemakerruntime { get; set; }

        [Input("savingsplans")]
        public Input<string>? Savingsplans { get; set; }

        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        [Input("schemas")]
        public Input<string>? Schemas { get; set; }

        [Input("sdb")]
        public Input<string>? Sdb { get; set; }

        [Input("secretsmanager")]
        public Input<string>? Secretsmanager { get; set; }

        [Input("securityhub")]
        public Input<string>? Securityhub { get; set; }

        [Input("serverlessapplicationrepository")]
        public Input<string>? Serverlessapplicationrepository { get; set; }

        [Input("serverlessapprepo")]
        public Input<string>? Serverlessapprepo { get; set; }

        [Input("serverlessrepo")]
        public Input<string>? Serverlessrepo { get; set; }

        [Input("servicecatalog")]
        public Input<string>? Servicecatalog { get; set; }

        [Input("servicecatalogappregistry")]
        public Input<string>? Servicecatalogappregistry { get; set; }

        [Input("servicediscovery")]
        public Input<string>? Servicediscovery { get; set; }

        [Input("servicequotas")]
        public Input<string>? Servicequotas { get; set; }

        [Input("ses")]
        public Input<string>? Ses { get; set; }

        [Input("sesv2")]
        public Input<string>? Sesv2 { get; set; }

        [Input("sfn")]
        public Input<string>? Sfn { get; set; }

        [Input("shield")]
        public Input<string>? Shield { get; set; }

        [Input("signer")]
        public Input<string>? Signer { get; set; }

        [Input("simpledb")]
        public Input<string>? Simpledb { get; set; }

        [Input("sms")]
        public Input<string>? Sms { get; set; }

        [Input("snowball")]
        public Input<string>? Snowball { get; set; }

        [Input("snowdevicemanagement")]
        public Input<string>? Snowdevicemanagement { get; set; }

        [Input("sns")]
        public Input<string>? Sns { get; set; }

        [Input("sqs")]
        public Input<string>? Sqs { get; set; }

        [Input("ssm")]
        public Input<string>? Ssm { get; set; }

        [Input("ssmcontacts")]
        public Input<string>? Ssmcontacts { get; set; }

        [Input("ssmincidents")]
        public Input<string>? Ssmincidents { get; set; }

        [Input("sso")]
        public Input<string>? Sso { get; set; }

        [Input("ssoadmin")]
        public Input<string>? Ssoadmin { get; set; }

        [Input("ssooidc")]
        public Input<string>? Ssooidc { get; set; }

        [Input("stepfunctions")]
        public Input<string>? Stepfunctions { get; set; }

        [Input("storagegateway")]
        public Input<string>? Storagegateway { get; set; }

        [Input("sts")]
        public Input<string>? Sts { get; set; }

        [Input("support")]
        public Input<string>? Support { get; set; }

        [Input("swf")]
        public Input<string>? Swf { get; set; }

        [Input("synthetics")]
        public Input<string>? Synthetics { get; set; }

        [Input("textract")]
        public Input<string>? Textract { get; set; }

        [Input("timestreamquery")]
        public Input<string>? Timestreamquery { get; set; }

        [Input("timestreamwrite")]
        public Input<string>? Timestreamwrite { get; set; }

        [Input("transcribe")]
        public Input<string>? Transcribe { get; set; }

        [Input("transcribeservice")]
        public Input<string>? Transcribeservice { get; set; }

        [Input("transcribestreaming")]
        public Input<string>? Transcribestreaming { get; set; }

        [Input("transcribestreamingservice")]
        public Input<string>? Transcribestreamingservice { get; set; }

        [Input("transfer")]
        public Input<string>? Transfer { get; set; }

        [Input("translate")]
        public Input<string>? Translate { get; set; }

        [Input("voiceid")]
        public Input<string>? Voiceid { get; set; }

        [Input("waf")]
        public Input<string>? Waf { get; set; }

        [Input("wafregional")]
        public Input<string>? Wafregional { get; set; }

        [Input("wafv2")]
        public Input<string>? Wafv2 { get; set; }

        [Input("wellarchitected")]
        public Input<string>? Wellarchitected { get; set; }

        [Input("wisdom")]
        public Input<string>? Wisdom { get; set; }

        [Input("workdocs")]
        public Input<string>? Workdocs { get; set; }

        [Input("worklink")]
        public Input<string>? Worklink { get; set; }

        [Input("workmail")]
        public Input<string>? Workmail { get; set; }

        [Input("workmailmessageflow")]
        public Input<string>? Workmailmessageflow { get; set; }

        [Input("workspaces")]
        public Input<string>? Workspaces { get; set; }

        [Input("workspacesweb")]
        public Input<string>? Workspacesweb { get; set; }

        [Input("xray")]
        public Input<string>? Xray { get; set; }

        public ProviderEndpointArgs()
        {
        }
        public static new ProviderEndpointArgs Empty => new ProviderEndpointArgs();
    }
}
