export default defineI18nConfig(() => ({
    legacy: false,
    locale: 'de',
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
                subtitle: '',
            }
        }
    }
}));
