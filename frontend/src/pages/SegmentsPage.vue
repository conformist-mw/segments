<template>
  <div class="row mt-4">
    <div class="col-4"></div>
    <div class="col-8">
      <SegmentList :segments="segments" :isLoading="isSegmentsLoading" />
    </div>
  </div>
</template>

<script>
import useSegments from '../hooks/useSegments';
import SegmentList from '../components/SegmentList.vue';

export default {
  components: {
    SegmentList,
  },
  setup() {
    const { segments, error, isSegmentsLoading } = useSegments();

    return {
      segments, error, isSegmentsLoading,
    };
  },
};
</script>

<style scoped>
</style>
