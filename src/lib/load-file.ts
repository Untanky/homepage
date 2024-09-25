import { readFile } from 'fs/promises';

export const loadJsonFile = async <Result>(path: string): Promise<Result> => {
  const buffer = await readFile(path);

  return JSON.parse(buffer.toString()) as Result;
};
