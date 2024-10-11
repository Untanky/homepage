import type { VFile } from 'vfile';

export interface FileService {
  listFiles(path: string): Promise<Array<string>>;
  readFile(path: string): Promise<VFile>;
}
