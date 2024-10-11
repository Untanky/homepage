import { FileService } from '@/models/file';
import { Octokit } from 'octokit';
import { VFile } from 'vfile';

export class GitHubService implements FileService {
  constructor(
    private readonly owner: string,
    private readonly repo: string,
    private readonly octokit: Octokit,
  ) {}

  async listFiles(path: string): Promise<Array<string>> {
    const result = await this.octokit.rest.repos.getContent({ owner: this.owner, repo: this.repo, path: path });

    if (result.status !== 200) {
      throw new Error('request was not successful');
    }

    if (!Array.isArray(result.data)) {
      throw new Error('not an array');
    }

    return result.data
      .filter((entry) => entry.type === 'file')
      .map((file) => file.path);
  }

  async readFile(path: string): Promise<VFile> {
    const result = await this.octokit.rest.repos.getContent({ owner: this.owner, repo: this.repo, path: path });

    if (result.status !== 200) {
      throw new Error('request was not successful');
    }

    if (Array.isArray(result.data)) {
      throw new Error('not a file');
    }

    if (result.data.type === 'submodule' || result.data.type === 'symlink') {
      throw new Error('not a file');
    }

    return new VFile({
      path,
      value: Buffer.from(result.data.content, 'base64').toString('utf-8'),
    });
  }
}
