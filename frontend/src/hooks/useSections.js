import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import $api from '../http';

export default function useSections() {
  const sections = ref([]);
  const error = ref('');
  const isLoading = ref(true);
  const route = useRoute();

  const fetchSections = () => {
    $api.get(`/companies/${route.params.companySlug}/sections/`)
      .then((response) => {
        sections.value = response.data;
      })
      .catch((err) => {
        error.value = err.response.data.detail;
      })
      .finally(() => {
        isLoading.value = false;
      });
  };

  onMounted(fetchSections);

  return {
    sections, error, isLoading,
  };
}
