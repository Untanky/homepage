<template>
  <div class="introduction">
    <div class="picture">
      <img
        class="profile-picture"
        alt="profile"
        :src="profileImage"
      >
    </div>
    <div class="info">
      <h1 class="title">
        Lukas Grimm
      </h1>
      <h3 class="subtitle">
        {{ $t('homepageContent.subTitle') }}
      </h3>
      <div
        v-html="$t('homepageContent.introduction')"
      />
    </div>
  </div>
</template>

<script>
import ProfileImage from '../assets/img/profile.jpg';

export default {
  name: 'Introduction',
  data() {
    return {
      profileImage: ProfileImage,
    };
  },
};
</script>

<style lang="scss" scoped>
@import "../assets/scss/variables";
@import "../assets/scss/responsive";
@import "../assets/scss/landing-page";

.introduction {
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  max-width: $landing-page-width;
  background: $primary-secondary-gradient;
  justify-content: center;
  align-content: center;
  padding: 1em;
  border-radius: $border-radius;
  margin: 0 auto;
  box-shadow: $shadow;

  @include tablet-up {
    padding: 2em;
    flex-wrap: nowrap;
  }

  .picture {
    box-sizing: border-box;
    // align-self: center;
    height: 264px;

    .profile-picture {
      width: 256px;
      height: 256px;
      border-radius: 100%;
      box-shadow: $shadow;
    }
  }

  .info {
    display: flex;
    flex-direction: column;
    align-content: center;
    justify-content: center;
    text-align: left;

    @include tablet-up {
      padding-left: 2em;
    }

    .title, .subtitle {
      text-align: center;
      margin: 0.25em 0;

      @include tablet-up {
        text-align: left;
      }
    }

    .title {
      font-weight: 800;
    }

    p {
      text-align: justify;
      margin: 0.25em 0;
    }
  }
}
</style>
