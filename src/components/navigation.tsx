import { CodeBracketIcon, HomeIcon } from '@heroicons/react/20/solid';
import Link from 'next/link';

export const Navigation = () => {
  return (
    <nav className="flex space-x-2 border border-zinc-500 p-1 rounded-xl mb-2">
      <Link href="/" className="flex items-center space-x-2 pl-1.5 pr-2.5 py-1 rounded-lg hover:bg-zinc-200 dark:hover:bg-zinc-800">
        <HomeIcon className="size-5" />
        <span>Home</span>
      </Link>
      <Link href="/projects" className="flex items-center space-x-2 pl-1.5 pr-2.5 py-1 rounded-lg hover:bg-zinc-200 dark:hover:bg-zinc-800">
        <CodeBracketIcon className="size-5" />
        <span>Projects</span>
      </Link>
    </nav>
  );
};
