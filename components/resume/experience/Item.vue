<template>
    <h3 class="text-lg font-medium">{{ experience.role }} - {{ experience.company }}</h3>
    <i>{{ formatDate(experience.start) }} - {{ experience.end ? formatDate(experience.end) : 'Now' }}</i>
    <p class="text-zinc-700 dark:text-zinc-300">{{ experience.description }}</p>
</template>

<script setup lang="ts">
    import type { ExperienceItemModel } from '../models';
    
    defineProps<{ experience: ExperienceItemModel }>();
</script>
