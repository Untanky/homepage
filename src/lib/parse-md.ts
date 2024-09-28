import { compile, run } from '@mdx-js/mdx';
import { ReactElement } from 'react';
import * as runtime from 'react/jsx-runtime';

type Props = {
  components?: never;
};

const parseMd = async (content: string): Promise<string> => {
  const code = await compile(content, { outputFormat: 'function-body' });

  return String(code);
};

export const mdAsComponent = async (content: string): Promise<(props: Props) => ReactElement> => {
  const code = await parseMd(content);

  // @ts-expect-error typing of runtimne is weird.
  const result = await run(code, {
    ...runtime,
    baseUrl: import.meta.url,
  });

  return result.default;
};
