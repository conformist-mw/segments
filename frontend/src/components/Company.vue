<template>
  <h1 class="my-5">Компании</h1>
  <div v-if="isLoading" class="d-flex justify-content-center">
    <div class="spinner-border" role="status">
      <span class="visually-hidden">Loading...</span>
    </div>
  </div>
  <div v-else-if="error" class="alert alert-danger" role="alert">
    {{ error }}
  </div>
  <div v-else class="row">
    <transition-group name="companies-list">
      <div
          class="col-6 d-flex align-items-stretch"
          v-for="company in companies"
          :key="company.slug"
      >
        <div class="card">
          <img :src="company.image" class="card-img-top" :alt="company.name">
          <div class="card-body d-flex flex-column justify-content-end text-center">
            <a
              href="#"
              @click="$router.push(`/company/${company.slug}/sections/`)"
              class="btn btn-outline-secondary stretched-link"
            >
              {{ company.name }}
            </a>
          </div>
        </div>
      </div>
    </transition-group>
  </div>
</template>

<script>
import $api from '../http';

export default {
  data() {
    return {
      companies: [],
      isLoading: false,
      error: false,
    };
  },
  methods: {
    fetchCompanies() {
      this.isLoading = true;
      $api.get('/companies/')
        .then((response) => {
          this.companies = response.data;
        })
        .catch((error) => {
          this.error = error.response.data.detail;
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
  mounted() {
    this.fetchCompanies();
  },
};
</script>

<style scoped>
.card {
  padding: 25px;
}
</style>
