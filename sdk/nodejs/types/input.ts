// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface CustomEnvironmentBranchTracking {
    /**
     * The pattern of the branch name to track.
     */
    pattern: pulumi.Input<string>;
    /**
     * How a branch name should be matched against the pattern. Must be one of 'startsWith', 'endsWith' or 'equals'.
     */
    type: pulumi.Input<string>;
}

export interface DeploymentProjectSettings {
    /**
     * The build command for this deployment. If omitted, this value will be taken from the project or automatically detected.
     */
    buildCommand?: pulumi.Input<string>;
    /**
     * The framework that is being used for this deployment. If omitted, no framework is selected.
     */
    framework?: pulumi.Input<string>;
    /**
     * The install command for this deployment. If omitted, this value will be taken from the project or automatically detected.
     */
    installCommand?: pulumi.Input<string>;
    /**
     * The output directory of the deployment. If omitted, this value will be taken from the project or automatically detected.
     */
    outputDirectory?: pulumi.Input<string>;
    /**
     * The name of a directory or relative path to the source code of your project. When null is used it will default to the project root.
     */
    rootDirectory?: pulumi.Input<string>;
}

export interface DnsRecordSrv {
    /**
     * The TCP or UDP port on which the service is to be found.
     */
    port: pulumi.Input<number>;
    /**
     * The priority of the target host, lower value means more preferred.
     */
    priority: pulumi.Input<number>;
    /**
     * The canonical hostname of the machine providing the service, ending in a dot.
     */
    target: pulumi.Input<string>;
    /**
     * A relative weight for records with the same priority, higher value means higher chance of getting picked.
     */
    weight: pulumi.Input<number>;
}

export interface FirewallConfigIpRules {
    rules?: pulumi.Input<pulumi.Input<inputs.FirewallConfigIpRulesRule>[]>;
}

export interface FirewallConfigIpRulesRule {
    action: pulumi.Input<string>;
    /**
     * Hosts to apply these rules to
     */
    hostname: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    /**
     * IP or CIDR to block
     */
    ip: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
}

export interface FirewallConfigManagedRulesets {
    /**
     * Enable the owasp managed rulesets and select ruleset behaviors
     */
    owasp?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwasp>;
}

export interface FirewallConfigManagedRulesetsOwasp {
    /**
     * Generic Attack Detection
     */
    gen?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspGen>;
    /**
     * Java Attack Detection
     */
    java?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspJava>;
    /**
     * Local File Inclusion Rules
     */
    lfi?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspLfi>;
    /**
     * Multipart Rules
     */
    ma?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspMa>;
    /**
     * PHP Attack Detection
     */
    php?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspPhp>;
    /**
     * Remote Code Execution Rules
     */
    rce?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspRce>;
    /**
     * Remote File Inclusion Rules
     */
    rfi?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspRfi>;
    /**
     * Scanner Detection Rules
     */
    sd?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspSd>;
    /**
     * SQL Injection Rules
     */
    sqli?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspSqli>;
    /**
     * Cross Site Scripting Rules
     */
    xss?: pulumi.Input<inputs.FirewallConfigManagedRulesetsOwaspXss>;
}

export interface FirewallConfigManagedRulesetsOwaspGen {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspJava {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspLfi {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspMa {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspPhp {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspRce {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspRfi {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspSd {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspSqli {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigManagedRulesetsOwaspXss {
    action: pulumi.Input<string>;
    active?: pulumi.Input<boolean>;
}

export interface FirewallConfigRules {
    rules?: pulumi.Input<pulumi.Input<inputs.FirewallConfigRulesRule>[]>;
}

export interface FirewallConfigRulesRule {
    /**
     * Actions to take when the condition groups match a request
     */
    action: pulumi.Input<inputs.FirewallConfigRulesRuleAction>;
    /**
     * Rule is active or disabled
     */
    active?: pulumi.Input<boolean>;
    /**
     * Sets of conditions that may match a request
     */
    conditionGroups: pulumi.Input<pulumi.Input<inputs.FirewallConfigRulesRuleConditionGroup>[]>;
    description?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    /**
     * Name to identify the rule
     */
    name: pulumi.Input<string>;
}

export interface FirewallConfigRulesRuleAction {
    /**
     * Base action
     */
    action: pulumi.Input<string>;
    /**
     * Forward persistence of a rule aciton
     */
    actionDuration?: pulumi.Input<string>;
    /**
     * Behavior or a rate limiting action. Required if action is rate_limit
     */
    rateLimit?: pulumi.Input<inputs.FirewallConfigRulesRuleActionRateLimit>;
    /**
     * How to redirect a request. Required if action is redirect
     */
    redirect?: pulumi.Input<inputs.FirewallConfigRulesRuleActionRedirect>;
}

export interface FirewallConfigRulesRuleActionRateLimit {
    /**
     * Action to take when rate limit is exceeded
     */
    action: pulumi.Input<string>;
    /**
     * Rate limiting algorithm
     */
    algo: pulumi.Input<string>;
    /**
     * Keys used to bucket an individual client
     */
    keys: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * number of requests allowed in the window
     */
    limit: pulumi.Input<number>;
    /**
     * Time window in seconds
     */
    window: pulumi.Input<number>;
}

export interface FirewallConfigRulesRuleActionRedirect {
    location: pulumi.Input<string>;
    permanent: pulumi.Input<boolean>;
}

export interface FirewallConfigRulesRuleConditionGroup {
    /**
     * Conditions that must all match within a group
     */
    conditions: pulumi.Input<pulumi.Input<inputs.FirewallConfigRulesRuleConditionGroupCondition>[]>;
}

export interface FirewallConfigRulesRuleConditionGroupCondition {
    /**
     * Key within type to match against
     */
    key?: pulumi.Input<string>;
    neg?: pulumi.Input<boolean>;
    /**
     * How to comparse type to value
     */
    op: pulumi.Input<string>;
    /**
     * Request key type to match against
     */
    type: pulumi.Input<string>;
    value?: pulumi.Input<string>;
}

export interface ProjectEnvironment {
    /**
     * A comment explaining what the environment variable is for.
     */
    comment?: pulumi.Input<string>;
    /**
     * The IDs of Custom Environments that the Environment Variable should be present on. At least one of `target` or `customEnvironmentIds` must be set.
     */
    customEnvironmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The git branch of the Environment Variable.
     */
    gitBranch?: pulumi.Input<string>;
    /**
     * The ID of the Environment Variable.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the Environment Variable.
     */
    key: pulumi.Input<string>;
    /**
     * Whether the Environment Variable is sensitive or not. (May be affected by a [team-wide environment variable policy](https://vercel.com/docs/projects/environment-variables/sensitive-environment-variables#environment-variables-policy))
     */
    sensitive?: pulumi.Input<boolean>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`. At least one of `target` or `customEnvironmentIds` must be set.
     */
    targets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The value of the Environment Variable.
     */
    value: pulumi.Input<string>;
}

export interface ProjectEnvironmentVariablesVariable {
    /**
     * A comment explaining what the environment variable is for.
     */
    comment?: pulumi.Input<string>;
    /**
     * The IDs of Custom Environments that the Environment Variable should be present on. At least one of `target` or `customEnvironmentIds` must be set.
     */
    customEnvironmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The git branch of the Environment Variable.
     */
    gitBranch?: pulumi.Input<string>;
    /**
     * The ID of the Environment Variable.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the Environment Variable.
     */
    key: pulumi.Input<string>;
    /**
     * Whether the Environment Variable is sensitive or not.
     */
    sensitive?: pulumi.Input<boolean>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`. At least one of `target` or `customEnvironmentIds` must be set.
     */
    targets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The value of the Environment Variable.
     */
    value: pulumi.Input<string>;
}

export interface ProjectGitComments {
    /**
     * Whether Commit comments are enabled
     */
    onCommit: pulumi.Input<boolean>;
    /**
     * Whether Pull Request comments are enabled
     */
    onPullRequest: pulumi.Input<boolean>;
}

export interface ProjectGitRepository {
    /**
     * Deploy hooks are unique URLs that allow you to trigger a deployment of a given branch. See https://vercel.com/docs/deployments/deploy-hooks for full information.
     */
    deployHooks?: pulumi.Input<pulumi.Input<inputs.ProjectGitRepositoryDeployHook>[]>;
    /**
     * By default, every commit pushed to the main branch will trigger a Production Deployment instead of the usual Preview Deployment. You can switch to a different branch here.
     */
    productionBranch?: pulumi.Input<string>;
    /**
     * The name of the git repository. For example: `vercel/next.js`.
     */
    repo: pulumi.Input<string>;
    /**
     * The git provider of the repository. Must be either `github`, `gitlab`, or `bitbucket`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectGitRepositoryDeployHook {
    /**
     * The ID of the deploy hook.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the deploy hook.
     */
    name: pulumi.Input<string>;
    /**
     * The branch or commit hash that should be deployed.
     */
    ref: pulumi.Input<string>;
    /**
     * A URL that, when a POST request is made to, will trigger a new deployment.
     */
    url?: pulumi.Input<string>;
}

export interface ProjectMembersMember {
    /**
     * The email of the user to add to the project. Exactly one of `userId`, `email`, or `username` must be specified.
     */
    email?: pulumi.Input<string>;
    /**
     * The role that the user should have in the project. One of 'MEMBER', 'PROJECT_DEVELOPER', or 'PROJECT_VIEWER'.
     */
    role: pulumi.Input<string>;
    /**
     * The ID of the user to add to the project. Exactly one of `userId`, `email`, or `username` must be specified.
     */
    userId?: pulumi.Input<string>;
    /**
     * The username of the user to add to the project. Exactly one of `userId`, `email`, or `username` must be specified.
     */
    username?: pulumi.Input<string>;
}

export interface ProjectOidcTokenConfig {
    /**
     * When true, Vercel issued OpenID Connect (OIDC) tokens will be available on the compute environments. See https://vercel.com/docs/security/secure-backend-access/oidc for more information.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Configures the URL of the `iss` claim. `team` = `https://oidc.vercel.com/[teamSlug]` `global` = `https://oidc.vercel.com`
     */
    issuerMode?: pulumi.Input<string>;
}

export interface ProjectOptionsAllowlist {
    /**
     * The allowed paths for the OPTIONS Allowlist. Incoming requests will bypass Deployment Protection if they have the method `OPTIONS` and **start with** one of the path values.
     */
    paths: pulumi.Input<pulumi.Input<inputs.ProjectOptionsAllowlistPath>[]>;
}

export interface ProjectOptionsAllowlistPath {
    /**
     * The path prefix to compare with the incoming request path.
     */
    value: pulumi.Input<string>;
}

export interface ProjectPasswordProtection {
    /**
     * The deployment environment to protect. Must be one of `standardProtection`, `allDeployments`, or `onlyPreviewDeployments`.
     */
    deploymentType: pulumi.Input<string>;
    /**
     * The password that visitors must enter to gain access to your Preview Deployments. Drift detection is not possible for this field.
     */
    password: pulumi.Input<string>;
}

export interface ProjectResourceConfig {
    /**
     * The amount of CPU available to your Serverless Functions. Should be one of 'standard_legacy' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
     */
    functionDefaultCpuType?: pulumi.Input<string>;
    /**
     * The default timeout for Serverless Functions.
     */
    functionDefaultTimeout?: pulumi.Input<number>;
}

export interface ProjectTrustedIps {
    /**
     * The allowed IP addressses and CIDR ranges with optional descriptions.
     */
    addresses: pulumi.Input<pulumi.Input<inputs.ProjectTrustedIpsAddress>[]>;
    /**
     * The deployment environment to protect. Must be one of `standardProtection`, `allDeployments`, `onlyProductionDeployments`, or `onlyPreviewDeployments`.
     */
    deploymentType: pulumi.Input<string>;
    /**
     * Whether or not Trusted IPs is optional to access a deployment. Must be either `trustedIpRequired` or `trustedIpOptional`. `trustedIpOptional` is only available with Standalone Trusted IPs.
     */
    protectionMode?: pulumi.Input<string>;
}

export interface ProjectTrustedIpsAddress {
    /**
     * A description for the value
     */
    note?: pulumi.Input<string>;
    /**
     * The address or CIDR range that can access deployments.
     */
    value: pulumi.Input<string>;
}

export interface ProjectVercelAuthentication {
    /**
     * The deployment environment to protect. Must be one of `standardProtection`, `allDeployments`, `onlyPreviewDeployments`, or `none`.
     */
    deploymentType: pulumi.Input<string>;
}

export interface TeamConfigRemoteCaching {
    /**
     * Indicates if Remote Caching is enabled.
     */
    enabled?: pulumi.Input<boolean>;
}

export interface TeamConfigSaml {
    /**
     * The ID of the access group to use for the team.
     */
    accessGroupId?: pulumi.Input<string>;
    /**
     * Indicates if SAML is enforced for the team.
     */
    enforced: pulumi.Input<boolean>;
    /**
     * Directory groups to role or access group mappings.
     */
    roles?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface TeamMemberProject {
    /**
     * The ID of the project that the user should be granted access to.
     */
    projectId: pulumi.Input<string>;
    /**
     * The role that the user should have in the project.
     */
    role: pulumi.Input<string>;
}
