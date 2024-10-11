export type ProjectPreview = {
  id: string;
  title: string;
  description: string;
  stack: string[];
  link: string;
};

export type Project = {
  id: string;
  title: string;
  description: string;
  stack: string[];
  version?: number;
  createdAt?: Date;
  editedAt?: Date;
  content: string;
};
