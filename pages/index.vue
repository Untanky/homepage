<template>
    <div class="max-w-[820px] mx-4 md:mx-auto grid grid-cols-6 gap-4">
        <div class="relative col-span-3 rounded-xl row-span-3">
            <img
                class="rounded-xl"
                src="/profile.jpg"
                alt="Profile Picture"
            >
            <h1 class="absolute bottom-0 left-0 right-0 text-center p-3 rounded-b-xl backdrop-blur-lg text-white text-6xl font-bold">
                {{ content.name }}
            </h1>
        </div>
        <h2 class="border border-rose-200 p-4 rounded-xl text-4xl font-medium col-span-3">
            {{ content.role }}
        </h2>
        <div class="border border-rose-200 p-4 rounded-xl">
            <div class="text-gray-700">Experience</div>
            <div class="text-xl">
                {{ content.experience }}
            </div>
        </div>
        <div class="border border-rose-200 p-4 rounded-xl col-span-2">
            <div class="text-gray-700">Education</div>
            <div class="text-xl">
                {{ content.education }}
            </div>
        </div>
        <div class="relative col-span-3 rounded-xl row-span-2">
            <img
                class="rounded-xl"
                src="/map.png"
                alt="Profile Picture"
            >
            <div class="absolute bottom-0 left-0 right-0 text-center p-3 rounded-b-xl backdrop-blur-lg text-white text-xl">
                {{ content.location }}
            </div>
        </div>
        <div class="border border-rose-200 p-4 rounded-xl col-span-3">
            <div class="text-gray-700 mb-2">About me</div>
            <div>
                {{ content.about }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const content = await queryContent('/biography').findOne()
console.log(content);
</script>
