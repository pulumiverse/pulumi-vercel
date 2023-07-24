import * as pulumi from "@pulumi/pulumi";
import * as vercel from "@pulumi/vercel";

const project = new vercel.Project("my-project", {});
new vercel.Deployment("my-deployment", {
  files: vercel.getProjectDirectoryOutput({ path: "./src" }).files,
  projectId: project.id,
});
