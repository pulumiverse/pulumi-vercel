// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Access Group Project Resource.
 *
 * An Access Group Project resource defines the relationship between a `vercel.AccessGroup` and a `vercel.Project`.
 *
 * For more detailed information, please see the [Vercel documentation](https://vercel.com/docs/accounts/team-members-and-roles/access-groups).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vercel from "@pulumiverse/vercel";
 *
 * const example = new vercel.Project("example", {name: "example-project"});
 * const exampleAccessGroup = new vercel.AccessGroup("example", {name: "example-access-group"});
 * const exampleAccessGroupProject = new vercel.AccessGroupProject("example", {
 *     projectId: example.id,
 *     accessGroupId: exampleAccessGroup.id,
 *     role: "ADMIN",
 * });
 * ```
 *
 * ## Import
 *
 * If importing into a personal account, or with a team configured on
 *
 * the provider, use the access_group_id and project_id.
 *
 * ```sh
 * $ pulumi import vercel:index/accessGroupProject:AccessGroupProject example ag_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
 * ```
 *
 * If importing to a team, use the team_id, access_group_id and project_id.
 *
 * ```sh
 * $ pulumi import vercel:index/accessGroupProject:AccessGroupProject example team_xxxxxxxxxxxxxxxxxxxxxxxx/ag_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
 * ```
 */
export class AccessGroupProject extends pulumi.CustomResource {
    /**
     * Get an existing AccessGroupProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessGroupProjectState, opts?: pulumi.CustomResourceOptions): AccessGroupProject {
        return new AccessGroupProject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vercel:index/accessGroupProject:AccessGroupProject';

    /**
     * Returns true if the given object is an instance of AccessGroupProject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessGroupProject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessGroupProject.__pulumiType;
    }

    /**
     * The ID of the Access Group.
     */
    public readonly accessGroupId!: pulumi.Output<string>;
    /**
     * The Project ID to assign to the access group.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The project role to assign to the access group. Must be either `ADMIN`, `PROJECT_DEVELOPER`, or `PROJECT_VIEWER`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The ID of the team the access group project should exist under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a AccessGroupProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessGroupProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessGroupProjectArgs | AccessGroupProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessGroupProjectState | undefined;
            resourceInputs["accessGroupId"] = state ? state.accessGroupId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as AccessGroupProjectArgs | undefined;
            if ((!args || args.accessGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessGroupId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["accessGroupId"] = args ? args.accessGroupId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessGroupProject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessGroupProject resources.
 */
export interface AccessGroupProjectState {
    /**
     * The ID of the Access Group.
     */
    accessGroupId?: pulumi.Input<string>;
    /**
     * The Project ID to assign to the access group.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The project role to assign to the access group. Must be either `ADMIN`, `PROJECT_DEVELOPER`, or `PROJECT_VIEWER`.
     */
    role?: pulumi.Input<string>;
    /**
     * The ID of the team the access group project should exist under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessGroupProject resource.
 */
export interface AccessGroupProjectArgs {
    /**
     * The ID of the Access Group.
     */
    accessGroupId: pulumi.Input<string>;
    /**
     * The Project ID to assign to the access group.
     */
    projectId: pulumi.Input<string>;
    /**
     * The project role to assign to the access group. Must be either `ADMIN`, `PROJECT_DEVELOPER`, or `PROJECT_VIEWER`.
     */
    role: pulumi.Input<string>;
    /**
     * The ID of the team the access group project should exist under. Required when configuring a team resource if a default team has not been set in the provider.
     */
    teamId?: pulumi.Input<string>;
}
