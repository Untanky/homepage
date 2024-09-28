import { loadFile } from '@/lib/load-file';
import { renderMarkdown } from '../../../src/lib/parse-md';

export default async function ProjectPage({ params }: { params: { slug: string } }) {
  const fileContent = await loadFile('/Users/lukas/dev/homepage/app/projects/[slug]/text.md');

  const { frontmatter, Content: ProjectContent } = await renderMarkdown(fileContent);

  // Render the MDX content, supplying the ClientComponent as a component
  return (
    <main className="prose prose-zinc dark:prose-invert m-4">
      <ProjectContent />
    </main>
  );
};
