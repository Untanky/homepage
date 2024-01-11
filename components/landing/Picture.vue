<template>
    <div class="relative md:min-h-[400px] rounded-xl bg-[url('/profile.jpg')] bg-cover bg-center">
        <component :is="as" class="absolute bottom-0 left-0 right-0 text-center p-3 rounded-b-xl backdrop-blur-lg text-white text-4xl md:text-6xl font-bold">
            {{ subtitle }}
        </component>
    </div>
</template>

<script setup lang="ts">
    withDefaults(defineProps<{
        subtitle: string;
        as: string;
    }>(), { as: 'div' });
</script>
