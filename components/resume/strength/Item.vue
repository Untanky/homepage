<template>
    <span class="text-lg font-medium mr-2">
        {{ strength.name }}
    </span>
    <span class="text-zinc-700 dark:text-zinc-300">
        {{ strength.description }}
    </span>
</template>

<script setup lang="ts">
    import type { StrengthItemModel } from '../models';

    defineProps<{ strength: StrengthItemModel }>();
</script>
