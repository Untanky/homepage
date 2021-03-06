<template>
  <div class="project">
    <card>
      <template slot="header">
        <h2 class="project-title">
          {{ project.title }}
        </h2>
      </template>

      <div class="project-image-container">
        <div
          class="project-links project-image"
          :style="`background-image: url(${project.imageUrl})`"
        >
          <a
            class="project-link github"
            :href="project.gitHubUrl"
          >
            <img
              src="https://upload.wikimedia.org/wikipedia/commons/9/91/Octicons-mark-github.svg"
              alt="GitHub"
            />
          </a>
          <a
            class="project-link demo"
            :href="project.demoUrl"
          >
            <img
              src="https://upload.wikimedia.org/wikipedia/commons/c/c4/Globe_icon.svg"
              alt="Demo"
            />
          </a>
        </div>
      </div>

      <template slot="footer">
        <div class="tag-list">
          <badge
            v-for="(tag, index, key) in project.tags"
            :key="key"
          >
            {{ tag }}
          </badge>
        </div>
      </template>
    </card>
  </div>
</template>

<script>
import Card from './Card.vue';
import Badge from './Badge.vue';

export default {
  name: 'Project',
  components: {
    Card,
    Badge,
  },
  props: {
    project: {
      type: Object,
      required: true,
    },
  },
  computed: {
    backgroundImage() { return `background-image: url(${this.project.imageUrl})`; },
  },
};
</script>

<style lang="scss">
@import "../assets/scss/variables";

.project {
  .project-title {
    margin: 0;
    text-align: left;
  }

  .project-image-container {
    position: relative;
    margin: -1em -1em -1em -1em;

    .project-image {
      width: 100%;
      height: 500px;
      background-color: transparent;
      background-repeat: no-repeat;
      background-position: center top;
      background-size: cover;
      background-blend-mode: luminosity;
      transition: 0.5s linear background-color;

      &:hover {
        background-color: rgba(1, 1, 1, 0.65);
        background-repeat: no-repeat;
        background-position: center top;
        background-size: cover;
        background-blend-mode: luminosity;
        transition: 0.5s linear background-color;

        .project-link {
          display: inline-block;
          width: 64px;
          height: 64px;
          align-self: center;
          background-color: transparent;
          background-blend-mode: overlay;
          background: $primary-secondary-gradient;
          // border: 1px solid black;
          border-radius: 100%;
          box-shadow: $shadow;
          transition: 0.25s linear background-color, 0.25s linear box-shadow;

          img {
            width: 32px;
            height: 32px;
            padding: 1em;
          }

          &:not(:last-child) {
            margin-right: 1em;
          }

          &:hover {
            background-color: rgb(163, 163, 163);
            box-shadow: $hover-shadow;
            transition: 0.25s linear background-color, 0.25s linear box-shadow;
          }
        }
      }
    }

    .project-links {
      display: flex;
      justify-content: center;
      align-content: center;

      .project-link {
        display: none;
      }
    }
  }

  footer {
    text-align: left;
    float: none;

    .tag-list {
      margin-top: -0.25em;

      .badge {
        margin: 0.25em 0.25em 0 0;
      }
    }
  }
}
</style>
