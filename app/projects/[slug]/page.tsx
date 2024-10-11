import { GitHubService } from '@/lib/github';
import { renderMarkdown } from '@/lib/parse-md';
import { ProjectPreview } from '@/models/project';
import { Octokit } from 'octokit';

const gitHubService = new GitHubService('untanky', 'content', new Octokit({
  userAgent: `untanky-homepage/${process.env.VERSION}`,
  auth: process.env.GITHUB_TOKEN,
}));

export default async function ProjectPage({ params }: { params: { slug: string } }) {
  const file = await gitHubService.readFile(`project/${params.slug}.md`);

  const { frontmatter, Content: ProjectContent } = await renderMarkdown<ProjectPreview>(file);

  // Render the MDX content, supplying the ClientComponent as a component
  return (
    <main className="prose prose-zinc dark:prose-invert m-4 md:w-[700px] md:mx-auto">
      <h1>{frontmatter.title}</h1>
      <ProjectContent />
    </main>
  );
};
