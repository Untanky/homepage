<template>
    <div class="text-lg font-semibold">
        {{ skill.name }}
        <div class="w-full h-4 bg-rose-100 dark:bg-rose-900 bg-opacity-10 rounded p-px">
            <div :style="`width: ${skill.value * 100}%`" class="h-full bg-rose-500 rounded"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import type { SkillItemModel } from '../models';

    defineProps<{ skill: SkillItemModel}>();
</script>
