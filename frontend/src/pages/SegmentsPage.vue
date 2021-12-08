<template>
  <h1 class="my-5">Отрезки!</h1>
  <div class="row">
    <div class="col-4"></div>
    <div class="col-8">
      <SegmentList :segments="segments" />
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
