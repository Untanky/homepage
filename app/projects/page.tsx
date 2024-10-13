import { ProjectList } from './project-list';
import { projectService } from './project-service';

export default async function ProjectsListPage() {
  const projects = await projectService.getProjectPreviews('project');

  return (
    <main>
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
