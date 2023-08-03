import * as vercel from "@pulumiverse/vercel";

const project = new vercel.Project("pulumi-vercel", {
  rootDirectory: "src",
});

const deployment = new vercel.Deployment("production-deployment", {
  projectId: project.id,
  production: true,
  files: vercel.getProjectDirectoryOutput({ path: "./src" }).files,
});

export const url = deployment.url;
