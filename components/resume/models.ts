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
