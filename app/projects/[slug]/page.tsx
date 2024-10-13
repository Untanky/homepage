import { renderMarkdown } from '@/lib/parse-md';
import { VFile } from 'vfile';
import { projectService } from '../project-service';

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
