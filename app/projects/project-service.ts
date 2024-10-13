import { GitHubService } from '@/lib/github';
import { ProjectService } from '@/lib/project';
import { Octokit } from 'octokit';

const octokit = new Octokit({
  userAgent: `untanky-homepage/${process.env.VERSION}`,
  auth: process.env.GITHUB_TOKEN,
});

const githubService = new GitHubService('untanky', 'content', octokit);

export const projectService = new ProjectService(githubService);
