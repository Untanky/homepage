import { readFile } from 'fs/promises';

export const loadFile = async (path: string): Promise<string> => {
  const buffer = await readFile(path);

  return buffer.toString();
};

export const loadJsonFile = async <Result>(path: string): Promise<Result> => {
  const content = await loadFile(path);

  return JSON.parse(content) as Result;
};
