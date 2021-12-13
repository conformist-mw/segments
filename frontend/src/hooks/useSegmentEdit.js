import { useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';
import $api from '../http';

export default function useSegmentEdit() {
  const route = useRoute();
  const segment = ref({});
  const error = ref('');
  const isLoading = ref(true);

  const fetchSegment = () => {
    $api.get(route.fullPath)
      .then((response) => {
        segment.value = response.data;
      })
      .catch((err) => {
        error.value = err.response.data;
      })
      .finally(() => {
        isLoading.value = false;
      });
  };

  const saveSegment = () => {
    $api.patch(`${route.fullPath}/`, segment.value)
      .then((response) => {
        segment.value = response.data;
      })
      .catch((err) => {
        error.value = err.response.data;
      })
      .finally(() => {
        isLoading.value = false;
      });
  };

  onMounted(fetchSegment);

  return {
    segment, error, isLoading, saveSegment,
  };
}
