<template>
    <NavigationHeader :links="links" />
    <slot></slot>
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
        {
            to: '/blog',
            name: t('blog'),
        },
    ]);
</script>
