export default defineI18nConfig(() => ({
    legacy: false,
    locale: 'en',
    messages: {
        en: {
            welcome: 'Welcome',
            projects: {
                title: 'Projects',
                subtitle: 'Take a look at the different projects I have worked on over the years.',
            }
        },
        de: {
            welcome: 'Willkommen',
            projects: {
                title: 'Projekte',
                subtitle: '',
            }
        }
    }
}));
