export type Experience = {
  startDate: Date;
  endDate?: Date;
  company: string;
  role: string;
  description: string;
};

export type Education = {
  startDate: Date;
  endDate: Date;
  school: string;
  course: string;
  grade: string;
  description: string;
};

export type Biography = {
  description: string;
  links: Array<{
    label: string;
    href: string;
    type: 'github' | 'linkedin' | 'location';
  }>;
};

export type Skill = {
  label: string;
  value: number;
};

export type Strength = {
  label: string;
  description: string;
};

export type Resume = {
  biography: Biography;
  experienceList: Experience[];
  educationList: Education[];
  skills: Skill[];
  strengths: Strength[];
};
