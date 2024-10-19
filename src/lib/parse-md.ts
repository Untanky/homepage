import { compile, type CompileOptions, run } from '@mdx-js/mdx';
import { MDXContent } from 'mdx/types.js';
import * as runtime from 'react/jsx-runtime';
import remarkFrontmatter from 'remark-frontmatter';
import remarkMdxFrontmatter from 'remark-mdx-frontmatter';
import { VFile } from 'vfile';
import { matter } from 'vfile-matter';

export type MarkdownFrontmatter = {
  id?: string;
  title?: string;
};

type RenderMarkdownResult = {
  Content: MDXContent;
};

const compileOptions: CompileOptions = {
  outputFormat: 'function-body',
  remarkPlugins: [remarkFrontmatter, remarkMdxFrontmatter],
};

const compileMarkdownContent = async (file: VFile): Promise<VFile> => {
  const code = await compile(file, compileOptions);

  return code;
};

export const getFrontmatter = async <T extends Record<string, unknown> = never>(file: VFile): Promise<T> => {
  matter(file);

  return file.data.matter as T;
};

export const renderMarkdown = async (file: VFile): Promise<RenderMarkdownResult> => {
  const code = await compileMarkdownContent(file);

  const result = await run(code, {
    ...runtime,
    baseUrl: import.meta.url,
  });

  return {
    Content: result.default,
  };
};
