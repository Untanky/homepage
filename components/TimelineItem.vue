<template>
  <div class="timeline-item">
    <span class="timeline" />
    <div class="timeline-info">
      <slot />
    </div>
    <div class="time-wrapper">
      <div class="time">
        {{ startDate }}
      </div>
      <div class="time">
        {{ endDate }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TimelineItem',
  props: {
    startDate: {
      type: String,
      required: true,
      default: '',
    },
    endDate: {
      type: String,
      required: true,
      default: '',
    },
  },
};
</script>

<style lang="scss" scoped>
@import '../assets/scss/responsive';
@import '../assets/scss/variables';

.timeline-item {
  position: relative;
  padding-top: 8px;
  margin-left: 2em;

  &:not(:first-child) {
    margin-top: 60px;
  }

  > .timeline { /* line */
    width: 2px;
    height: 100%;
    background: rgb(100, 100, 100);
    left: -30px;
    top: 0;
    position: absolute;
    &:before, &:after {/* circles */
      content: '';
      width: 8px;
      height: 8px;
      border-radius: 50%;
      border: 2px solid white;
      position: absolute;
      background: $primary;
      left: -5px;
      top: 0;
    }

    &:after {
      top: 100%;
    }
  }

  .time-wrapper {
    .time {
      position: absolute;
      left: -40px;
      width: 90px;
      text-align: left;
      font-size: 10px;
      font-weight: bold;
      padding: 0 0.5em;
      border-radius: 16px;
      &:first-child {
        top: -20px;
      }
      &:last-child {
        bottom: -30px;
      }
    }
  }
}

.timeline-info {
  background: $primary-secondary-gradient;
  padding: 1em;
  border-radius: 12px;
  // margin-top: 3px;
  // margin-right: 1em;
  box-shadow: $shadow;
}
</style>
