import { CodeBracketIcon, DocumentTextIcon } from '@heroicons/react/20/solid';
import { GitHubIcon } from '@/components/github-icon';
import Link from 'next/link';
import { LinkedInIcon } from '@/components/linkedin-icon';

export default function Home() {
  return (
    <main className="mx-4 sm:w-[400px] sm:mx-auto">
      <h1 className="text-4xl font-bold text-center mt-4">
        Lukas Grimm
      </h1>
      <h2 className="text-2xl font-semibold text-center text-zinc-700 dark:text-zinc-300">
        Software Engineer
      </h2>
      <section className="mt-2">
        <h3 className="sr-only">
          About
        </h3>
        <p>
          Lukas is a Software Engineer from Berlin, Germany. He likes to work with modern technologies and knows a lot about Web Development, Cloud and DevOps.
        </p>
      </section>
      <section className="mt-3">
        <h3 className="sr-only">
          General Information
        </h3>
        <div className="grid grid-cols-2 divide-x-2 divide-rose-500 dark:divide-rose-500">
          <div className="pr-2">
            <span className="block text-zinc-700 dark:text-zinc-300">
              Experience
            </span>
            <span className="block text-xl">
              4 years
            </span>
          </div>
          <div className="pl-2">
            <span className="block text-zinc-700 dark:text-zinc-300">
              Education
            </span>
            <span className="block text-xl">
              B.Sc. Computer Science
            </span>
          </div>
        </div>
      </section>
      <section className="flex flex-col mt-1 space-y-3">
        <h3 className="sr-only">
          Links
        </h3>
        <Link
          className="flex items-center space-x-2 bg-rose-500 text-white dark:text-black px-4 py-2 border border-rose-400 dark:border-rose-600 rounded-xl hover:bg-rose-600 hover:dark:bg-rose-400 transition-colors"
          href="/resume"
        >
          <DocumentTextIcon className="size-5" />
          <span>See Resumé</span>
        </Link>
        <Link
          className="flex items-center space-x-2 px-4 py-2 border border-rose-500 rounded-xl hover:bg-rose-50 hover:dark:bg-rose-950 transition-colors"
          href="/projects"
        >
          <CodeBracketIcon className="size-5" />
          <span>See Projects</span>
        </Link>
      </section>
      <section className="flex flex-col space-y-3">
        <h3 className="sr-only">
          Social Media Links
        </h3>
        <Link
          className="flex items-center space-x-2 px-4 py-2 border border-rose-500 rounded-xl hover:bg-rose-50 hover:dark:bg-rose-950 transition-colors"
          href="https://github.com/untanky"
        >
          <GitHubIcon className="size-5" />
          <span>GitHub</span>
        </Link>
        <Link
          className="flex items-center space-x-2 px-4 py-2 border border-rose-500 rounded-xl hover:bg-rose-50 hover:dark:bg-rose-950 transition-colors"
          href="https://linkedin.com/in/lukasgrimm"
        >
          <LinkedInIcon className="size-5" />
          <span>LinkedIn</span>
        </Link>
      </section>
    </main>
  );
}
