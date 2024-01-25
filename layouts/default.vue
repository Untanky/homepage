<template>
    <NavigationHeader :links="links" />
    <slot />
    <footer class="h-16" />
</template>

<script setup lang="ts">
    const { t } = useI18n({
        messages: {
            en: {
                home: 'Home',
                resume: 'Resumé',
                projects: 'Projects',
                blog: 'Blog',
            },
            de: {
                home: 'Start',
                resume: 'Lebenslauf',
                projects: 'Projekte',
                blog: 'Blog',
            }
        }
    });

    useHead({
        script: [
            {
                async: true,
                src: 'https://analytics.lukasgrimm.me/script.js',
                'data-website-id': '0962f0bb-97d5-4f08-ad6c-0fd2a284c41e',
            }
        ]
    });

    const links = ref([
        {
            to: '/',
            name: t('home'),
        },
        {
            to: '/resume',
            name: t('resume'),
        },
        {
            to: '/projects',
            name: t('projects'),
        },
        // {
        //     to: '/blog',
        //     name: t('blog'),
        // },
    ]);
</script>
