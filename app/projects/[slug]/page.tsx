import { compile, run } from '@mdx-js/mdx';
import React from 'react';
import * as runtime from 'react/jsx-runtime';

export default async function ProjectPage({ params }: { params: { slug: string } }) {
  const mdxSource = `# Hello, ${params.slug}
  
  Test
  
  ## My second test
  
  Foo
  
  Bar
  
  Baz
  
  Test
  
  Lol
  
  Amazing`;

  const code = String(
    await compile(mdxSource, { outputFormat: 'function-body' }),
  );

  // @ts-expect-error typing of runtimne is weird.
  const { default: MDXContent } = await run(code, {
    ...runtime,
    baseUrl: import.meta.url,
  });

  // Render the MDX content, supplying the ClientComponent as a component
  return (
    <main className="prose prose-zinc dark:prose-invert m-4">
      <MDXContent />
    </main>
  );
};
