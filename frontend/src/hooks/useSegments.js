import { useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';
import $api from '../http';

export default function useSegments() {
  const route = useRoute();
  const segments = ref([]);
  const error = ref('');
  const isSegmentsLoading = ref(true);

  const fetchSegments = () => {
    $api.get(route.fullPath)
      .then((response) => {
        segments.value = response.data.results;
      })
      .catch((err) => {
        error.value = err.response.data.detail;
      })
      .finally(() => {
        isSegmentsLoading.value = false;
      });
  };

  onMounted(fetchSegments);

  return {
    segments, error, isSegmentsLoading,
  };
}
