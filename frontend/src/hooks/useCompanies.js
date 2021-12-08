import { onMounted, ref } from 'vue';
import $api from '../http';

export default function useCompanies() {
  const companies = ref([]);
  const error = ref('');
  const isLoading = ref(true);

  const fetchCompanies = () => {
    $api.get('/companies/')
      .then((response) => {
        companies.value = response.data;
      })
      .catch((err) => {
        error.value = err.response.data.detail;
      })
      .finally(() => {
        isLoading.value = false;
      });
  };

  onMounted(fetchCompanies);

  return {
    companies, error, isLoading,
  };
}
