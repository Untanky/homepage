import { loadFile } from '@/lib/load-file';
import { getFrontmatter } from '@/lib/parse-md';
import { ProjectPreview } from '../../src/models/project';
import { ProjectList } from './project-list';

const getProjects = async (): Promise<ProjectPreview[]> => {
  const fileContent = await loadFile('/Users/lukas/dev/homepage/app/projects/[slug]/text.md');

  console.log(await getFrontmatter(fileContent));

  return Promise.resolve([
    {
      id: 'website-overhaul',
      title: 'Website Overhaul',
      description: 'Today I want to present the website overhaul I have been working on.',
      stack: ['TypeScript', 'NextJS'],
      link: '/projects/website-overhaul',
    },
    {
      id: 'git-charged',
      title: 'GitCharged',
      description: 'GitCharged improves the UX of the Git CLI and adds some cool, opinioned features to Git.',
      stack: ['Zig', 'Git'],
      link: '/projects/git-charged',
    },
  ]);
};

export default async function ProjectsListPage() {
  const projects = await getProjects();

  return (
    <main className="m-4">
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
