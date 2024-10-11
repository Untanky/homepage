import { FileService } from '@/models/file';
import type { ProjectPreview, Project } from '@/models/project';
import { getFrontmatter } from './parse-md';

export class ProjectService {
  constructor(private readonly fileService: FileService) {}

  async getProjectPreviews(path: string): Promise<Array<ProjectPreview>> {
    const filePaths = await this.fileService.listFiles(path);

    const results = await Promise.allSettled(filePaths.map(
      (path) => this.fileService.readFile(path)
        .then((file) => (console.log(file), getFrontmatter<ProjectPreview>(file))),
    ));

    results
      .filter((result) => result.status === 'rejected')
      .forEach((result) => console.error(result.reason));

    return results
      .filter((result) => result.status === 'fulfilled')
      .map(({ value }) => value);
  }

  async getProject(): Promise<Project> {
    throw new Error('not implemented');
  }
}
