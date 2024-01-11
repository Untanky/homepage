<template>
    <ol class="ml-8 space-y-2">
        <li v-for="experience in experiences">
            <ResumeExperienceItem :experience="experience" />
        </li>
    </ol>
</template>

<script setup lang="ts">
    import type { ExperienceItemModel } from '../models';
    
    defineProps<{ experiences: ExperienceItemModel[] }>();
</script>