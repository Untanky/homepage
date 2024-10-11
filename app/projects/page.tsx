import { GitHubService } from '@/lib/github';
import { ProjectService } from '@/lib/project';
import { Octokit } from 'octokit';
import { ProjectList } from './project-list';

const gitHubService = new GitHubService('untanky', 'content', new Octokit({
  userAgent: `untanky-homepage/${process.env.VERSION}`,
  auth: process.env.GITHUB_TOKEN,
}));

const projectService = new ProjectService(gitHubService);

export default async function ProjectsListPage() {
  const projects = await projectService.getProjectPreviews('project');

  return (
    <main className="m-4 md:w-[700px] md:mx-auto">
      <h1 className="text-4xl font-bold">
        Lukas&apos; Projects
      </h1>
      <p className="text-zinc-700 dark:text-zinc-300">
        Learn more about my projects.
      </p>
      <ProjectList projects={projects} />
    </main>
  );
};
