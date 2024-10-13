import { GitHubService } from '@/lib/github';
import { renderMarkdown } from '@/lib/parse-md';
import { ProjectService } from '@/lib/project';
import { Octokit } from 'octokit';
import { VFile } from 'vfile';

const gitHubService = new GitHubService('untanky', 'content', new Octokit({
  userAgent: `untanky-homepage/${process.env.VERSION}`,
  auth: process.env.GITHUB_TOKEN,
}));

const projectService = new ProjectService(gitHubService);

export default async function ProjectPage({ params }: { params: { slug: string } }) {
  const project = await projectService.getProject(`project/${params.slug}.md`);

  const { Content: ProjectContent } = await renderMarkdown(new VFile(project.content));

  // Render the MDX content, supplying the ClientComponent as a component
  return (
    <article className="prose prose-zinc dark:prose-invert">
      <h1>{project.title}</h1>
      <ProjectContent />
    </article>
  );
};
