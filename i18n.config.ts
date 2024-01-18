export default defineI18nConfig(() => ({
    legacy: false,
    locale: 'en',
    messages: {
        en: {
            welcome: 'Welcome',
            landing: {
                experience: 'Experience',
                education: 'Education',
                aboutMe: 'About me',
                seeResume: 'See Resumé',
                exploreProjects: 'Explore Projects',
            },
            resume: {
                biography: 'Biography',
                experience: 'Experience',
                education: 'Education',
                strengths: 'Strengths',
                skills: 'Skills',
                languages: 'Languages'
            },
            projects: {
                title: 'Projects',
                subtitle: 'Take a look at the different projects I have worked on over the years.',
                noContent: 'No projects found',
                learnMore: 'Learn More'
            },
            blog: {
                title: 'Blog',
                subtitle: 'Checkout my personal blog. Here I write about stuff that interests me. Mostly I want to write about my technical projects but maybe in the future I will write about other topics that interest me.',
                noContent: 'No blog posts found',
                learnMore: 'Learn More'
            }
        },
        de: {
            welcome: 'Willkommen',
            landing: {
                experience: 'Erfahrung',
                education: 'Ausbildung',
                aboutMe: 'Über mich',
                seeResume: 'Lebenslauf',
                exploreProjects: 'Projekte ansehen',
            },
            resume: {
                biography: 'Biografie',
                experience: 'Erfahrung',
                education: 'Ausbildung',
                strengths: 'Stärken',
                skills: 'Skills',
                languages: 'Sprachen'
            },
            projects: {
                title: 'Projekte',
                subtitle: 'Schaue dir meine unterschiedlichen Projekte an. ',
                noContent: 'Keine Projekte gefunden',
                learnMore: "Mehr erfahren"
            },
            blog: {
                title: 'Blog',
                subtitle: 'Take a look at the different projects I have worked on over the years.',
                noContent: 'Keine Blogposts gefunden',
                learnMore: 'Lesen'
            }
        }
    }
}));
