import { compile, type CompileOptions, run } from '@mdx-js/mdx';
import type { ReactElement } from 'react';
import * as runtime from 'react/jsx-runtime';
import remarkFrontmatter from 'remark-frontmatter';
import remarkMdxFrontmatter from 'remark-mdx-frontmatter';
import { VFile } from 'vfile';
import { matter } from 'vfile-matter';

export type MarkdownFrontmatter = {
  id?: string;
  title?: string;
};

type Props = {
  components?: never;
};

type RenderMarkdownResult<T extends Record<string, unknown> = never> = {
  frontmatter: T;
  Content: (props: Props) => ReactElement;
};

const compileOptions: CompileOptions = {
  outputFormat: 'function-body',
  remarkPlugins: [remarkFrontmatter, remarkMdxFrontmatter],
};

const compileMarkdownContent = async (content: string): Promise<VFile> => {
  const code = await compile(content, compileOptions);

  return code;
};

export const getFrontmatter = async <T extends Record<string, unknown> = never>(content: string): Promise<T> => {
  const file = new VFile(content);
  matter(file);

  return file.data.matter as T;
};

export const renderMarkdown = async <T extends Record<string, unknown> = never>(content: string): Promise<RenderMarkdownResult<T>> => {
  const code = await compileMarkdownContent(content);

  // @ts-expect-error typing of runtime is weird.
  const result = await run(code, {
    ...runtime,
    baseUrl: import.meta.url,
  });

  return {
    frontmatter: result.frontmatter as T,
    Content: result.default,
  };
};
