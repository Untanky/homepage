<template>
    <h3 class="text-lg font-medium">{{ education.school }} - {{ education.course }} - Grade: {{ education.grade }}</h3>
    <i>{{ formatDate(education.start) }} - {{ education.end ? formatDate(education.end) : 'Now' }}</i>
    <p class="text-zinc-700 dark:text-zinc-300">{{ education.description }}</p>
</template>

<script setup lang="ts">
    import type { EducationItemModel } from '../models';
    
    defineProps<{ education: EducationItemModel }>();
</script>
