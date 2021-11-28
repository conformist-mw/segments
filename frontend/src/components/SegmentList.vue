<template>
  <div class="segments-list" v-if="segments.length > 0">
    <Segment
      v-for="segment in segments"
      :segment="segment"
      :key="segment.id"
    />
  </div>
  <h2 v-else style="color: red">
    Таких отрезов нет.
  </h2>
</template>

<script>
import Segment from '@/components/Segment.vue';

export default {
  components: {
    Segment,
  },
  props: {
    segments: {
      type: Array,
      required: true,
    },
  },
};
</script>

<style scoped>
.segments-list {
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 10px 15px;
}
</style>
