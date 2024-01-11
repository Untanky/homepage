<template>
    <ol class="ml-6 space-y-2 list-disc">
        <li v-for="education in educations">
            <ResumeEducationItem :education="education" />
        </li>
    </ol>
</template>

<script setup lang="ts">
    import type { EducationItemModel } from '../models';
    
    defineProps<{ educations: EducationItemModel[] }>();
</script>