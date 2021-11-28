<template>
  <div class="container">
    <div class="grid">
      <div class="col-4">
        <h2>Create segment</h2>
      </div>
      <div class="col-8">
        <div v-if="error" class="error">
          <i class="pi pi-exclamation-triangle"></i> Ошибка!
        </div>
        <ProgressSpinner v-else-if="isLoading" class="block right-auto top-50" />
        <SegmentList v-else :segments="segments" />
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
// import Button from 'primevue/button';
import ProgressSpinner from 'primevue/progressspinner';
import SegmentList from '@/components/SegmentList.vue';

export default {
  components: {
    ProgressSpinner,
    SegmentList,
    // Button,
  },
  data() {
    return {
      segments: [],
      isLoading: false,
      error: false,
    };
  },
  methods: {
    fetchSegments() {
      this.isLoading = true;
      axios.get('http://127.0.0.1:8000/api/segments/')
        .then((response) => {
          this.error = false;
          this.isLoading = false;
          this.segments = response.data.results;
        })
        .catch(() => {
          this.error = true;
        });
    },
  },
  mounted() {
    this.fetchSegments();
  },
};
</script>

<style>
.container {
  width: 1170px;
  margin: 0 auto;
}
</style>
