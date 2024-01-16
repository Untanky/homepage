export type ExperienceItemModel = {
    start: Date;
    end?: Date;
    company: string;
    role: string;
    description: string;
};

export type EducationItemModel = {
    start: Date;
    end?: Date;
    school: string;
    course: string;
    grade: string;
    description: string;
};

export type SkillItemModel = {
    name: string;
    value: number;
};

export type StrengthItemModel = {
    name: string;
    description: string;
};

export type BiographyModel = {
    name: string;
    role: string;
    about: string;
    location: string;
    github: string;
    linkedIn: string;
    x: string;
};
