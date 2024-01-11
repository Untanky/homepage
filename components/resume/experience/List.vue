<template>
    <ol class="ml-6 space-y-2 list-disc">
        <li v-for="experience in experiences">
            <ResumeExperienceItem :experience="experience" />
        </li>
    </ol>
</template>

<script setup lang="ts">
    import type { ExperienceItemModel } from '../models';
    
    defineProps<{ experiences: ExperienceItemModel[] }>();
</script>