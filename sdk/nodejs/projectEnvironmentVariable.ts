// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumi/vercel";
 *
 * const exampleProject = new vercel.Project("exampleProject", {gitRepository: {
 *     type: "github",
 *     repo: "vercel/some-repo",
 * }});
 * // An environment variable that will be created
 * // for this project for the "production" environment.
 * const exampleProjectEnvironmentVariable = new vercel.ProjectEnvironmentVariable("exampleProjectEnvironmentVariable", {
 *     projectId: exampleProject.id,
 *     key: "foo",
 *     value: "bar",
 *     targets: ["production"],
 * });
 * // An environment variable that will be created
 * // for this project for the "preview" environment when the branch is "staging".
 * const exampleGitBranch = new vercel.ProjectEnvironmentVariable("exampleGitBranch", {
 *     projectId: exampleProject.id,
 *     key: "foo",
 *     value: "bar-staging",
 *     targets: ["preview"],
 *     gitBranch: "staging",
 * });
 * ```
 *
 * ## Import
 *
 * If importing into a personal account, or with a team configured on the provider, simply use the project_id and environment variable id. - project_id can be found in the project `settings` tab in the Vercel UI. - environment variable id can be taken from the network tab on the project page.
 *
 * ```sh
 *  $ pulumi import vercel:index/projectEnvironmentVariable:ProjectEnvironmentVariable example prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/FdT2e1E5Of6Cihmt
 * ```
 *
 *  Alternatively, you can import via the team_id, project_id and environment variable id. - team_id can be found in the team `settings` tab in the Vercel UI. - project_id can be found in the project `settings` tab in the Vercel UI. - environment variable id can be taken from the network tab on the project page.
 *
 * ```sh
 *  $ pulumi import vercel:index/projectEnvironmentVariable:ProjectEnvironmentVariable example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/FdT2e1E5Of6Cihmt
 * ```
 */
export class ProjectEnvironmentVariable extends pulumi.CustomResource {
    /**
     * Get an existing ProjectEnvironmentVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectEnvironmentVariableState, opts?: pulumi.CustomResourceOptions): ProjectEnvironmentVariable {
        return new ProjectEnvironmentVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vercel:index/projectEnvironmentVariable:ProjectEnvironmentVariable';

    /**
     * Returns true if the given object is an instance of ProjectEnvironmentVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectEnvironmentVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectEnvironmentVariable.__pulumiType;
    }

    /**
     * The git branch of the Environment Variable.
     */
    public readonly gitBranch!: pulumi.Output<string | undefined>;
    /**
     * The name of the Environment Variable.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the Vercel project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
     */
    public readonly targets!: pulumi.Output<string[]>;
    /**
     * The ID of the Vercel team.Required when configuring a team resource if a default team has not been set in the provider.
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The value of the Environment Variable.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a ProjectEnvironmentVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectEnvironmentVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectEnvironmentVariableArgs | ProjectEnvironmentVariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectEnvironmentVariableState | undefined;
            resourceInputs["gitBranch"] = state ? state.gitBranch : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["targets"] = state ? state.targets : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ProjectEnvironmentVariableArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.targets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targets'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["gitBranch"] = args ? args.gitBranch : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["value"] = args?.value ? pulumi.secret(args.value) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProjectEnvironmentVariable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectEnvironmentVariable resources.
 */
export interface ProjectEnvironmentVariableState {
    /**
     * The git branch of the Environment Variable.
     */
    gitBranch?: pulumi.Input<string>;
    /**
     * The name of the Environment Variable.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the Vercel project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
     */
    targets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Vercel team.Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The value of the Environment Variable.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectEnvironmentVariable resource.
 */
export interface ProjectEnvironmentVariableArgs {
    /**
     * The git branch of the Environment Variable.
     */
    gitBranch?: pulumi.Input<string>;
    /**
     * The name of the Environment Variable.
     */
    key: pulumi.Input<string>;
    /**
     * The ID of the Vercel project.
     */
    projectId: pulumi.Input<string>;
    /**
     * The environments that the Environment Variable should be present on. Valid targets are either `production`, `preview`, or `development`.
     */
    targets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Vercel team.Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
    /**
     * The value of the Environment Variable.
     */
    value: pulumi.Input<string>;
}
