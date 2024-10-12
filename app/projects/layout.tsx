import { ChevronLeftIcon } from '@heroicons/react/16/solid';
import Link from 'next/link';

export default function ProjectLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="m-4 md:w-[700px] md:mx-auto">
      <nav className="mb-4">
        <Link
          href="/"
          className="inline-flex items-center space-x-1 underline pl-1 pr-2 py-1 rounded-lg hover:bg-zinc-200 dark:hover:bg-zinc-800"
        >
          <ChevronLeftIcon className="size-5" />
          <span>Back</span>
        </Link>
      </nav>
      {children}
    </div>
  );
}
