export type User = {
    id: string;
    name: string;
    surname: string;
    phone?: string;
    city?: string;
    email: string;
    avatar_url?: string;
};

export type School = {
    id: string;
    owner_id: string;
    name: string;
    description: string;
};

export type Course = {
    id: string;
    schoolID: string;
    name: string;
    description: string;
    level: number;
    price: number;
    language: string;
    image_url?: string;
    status: string;
};

export type Test = {
    id: string;
    task_url: string;
    options: string[];
    answer: string;
    level: number;
    score: number;
};

export type Lesson = {
    id: string;
    title: string;
    score: number;
    type: string;
    theory_url: string;
    video_url: string;
    tests: Test[];
    markdown: string;
};

export type LessonBrief = {
    id: string,
    title: string,
    type: string,
};

export type TestStat = {
    id: string,
    test_id: string,
    user_id: string,
    score: number,
};

export type LessonStat = {
    id: string,
    lesson_id: string,
    user_id: string,
    score: number,
    tests: TestStat[],
};

export type CourseBrief = {
    id: string;
    name: string;
    description: string;
    price: number;
    language: string;
    image_url: string;
    lessons: LessonBrief[];
};
