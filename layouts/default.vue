<template>
    <NavigationHeader :links="links" />
    <slot />
</template>

<script setup lang="ts">
const links = ref([
    {
        to: '/',
        name: 'Home',
    },
    {
        to: '/resume',
        name: 'Resumé',
    },
    {
        to: '/projects',
        name: 'Projects',
    },
    {
        to: '/blog',
        name: 'Blog',
    },
]);
</script>
