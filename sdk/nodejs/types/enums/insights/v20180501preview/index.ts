// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApplicationType = {
    Web: "web",
    Other: "other",
} as const;

/**
 * Type of application being monitored.
 */
export type ApplicationType = (typeof ApplicationType)[keyof typeof ApplicationType];

export const FlowType = {
    Bluefield: "Bluefield",
} as const;

/**
 * Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
 */
export type FlowType = (typeof FlowType)[keyof typeof FlowType];

export const IngestionMode = {
    ApplicationInsights: "ApplicationInsights",
    ApplicationInsightsWithDiagnosticSettings: "ApplicationInsightsWithDiagnosticSettings",
    LogAnalytics: "LogAnalytics",
} as const;

/**
 * Indicates the flow of the ingestion.
 */
export type IngestionMode = (typeof IngestionMode)[keyof typeof IngestionMode];

export const PublicNetworkAccessType = {
    /**
     * Enables connectivity to Application Insights through public DNS.
     */
    Enabled: "Enabled",
    /**
     * Disables public connectivity to Application Insights through public DNS.
     */
    Disabled: "Disabled",
} as const;

/**
 * The network access type for accessing Application Insights query.
 */
export type PublicNetworkAccessType = (typeof PublicNetworkAccessType)[keyof typeof PublicNetworkAccessType];

export const RequestSource = {
    Rest: "rest",
} as const;

/**
 * Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
 */
export type RequestSource = (typeof RequestSource)[keyof typeof RequestSource];

export const WebTestKind = {
    Ping: "ping",
    Multistep: "multistep",
} as const;

/**
 * The kind of WebTest that this web test watches. Choices are ping and multistep.
 */
export type WebTestKind = (typeof WebTestKind)[keyof typeof WebTestKind];

export const WebTestKindEnum = {
    Ping: "ping",
    Multistep: "multistep",
    Basic: "basic",
    Standard: "standard",
} as const;

/**
 * The kind of web test this is, valid choices are ping, multistep, basic, and standard.
 */
export type WebTestKindEnum = (typeof WebTestKindEnum)[keyof typeof WebTestKindEnum];
