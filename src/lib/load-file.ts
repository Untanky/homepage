import type { VFile } from 'vfile';
import { read } from 'to-vfile';

export const loadFile = async (path: string): Promise<VFile> => {
  return read(path);
};

export const loadJsonFile = async <Result>(path: string): Promise<Result> => {
  const content = await loadFile(path);

  return JSON.parse(String(content)) as Result;
};
