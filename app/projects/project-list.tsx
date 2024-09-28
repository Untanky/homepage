import { ProjectPreview } from '@/models/project';
import Link from 'next/link';

export const ProjectList = ({ projects }: { projects: ProjectPreview[] }) => {
  return (
    <ul className="mt-4 space-y-4">
      {projects.map((project, index) => (
        <li key={index}>
          <Link href={project.link} className="block p-4 border border-zinc-500 rounded-xl">
            <h2 className="text-xl font-semibold underline decoration-rose-500">{project.title}</h2>
            <p className="text-zinc-700 dark:text-zinc-300">{project.description}</p>
            <ul className="flex space-x-2 mt-2">
              {project.stack.map((it, index) => (
                <li key={index} className="text-sm font-light bg-rose-700 px-2.5 py-px rounded-full">
                  {it}
                </li>
              ))}
            </ul>
          </Link>
        </li>
      ))}
    </ul>
  );
};
