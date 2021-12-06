<template>
  <h1 class="my-5">Цеха</h1>
  <div v-if="isLoading" class="d-flex justify-content-center">
    <div class="spinner-border" role="status">
      <span class="visually-hidden">Loading...</span>
    </div>
  </div>
  <div v-else-if="error" class="alert alert-danger" role="alert">
    {{ error }}
  </div>
  <div v-else class="row">
    <transition-group name="section-list">
      <div
        class="col-6 d-flex align-items-stretch"
        v-for="section in sections"
        :key="section.slug"
      >
        <div class="card">
          <div class="card-header text-center">
            {{ section.name }} — {{ $route.fullPath }}
          </div>
          <ul class="list-group list-group-flush">
            <li class="list-group-item d-flex justify-content-between">
              <span>Количество стеллажей</span>
              <span class="badge bg-dark">{{ section.racks_count }}</span>
            </li>
            <li class="list-group-item d-flex justify-content-between">
              <span>Количество отрезов</span>
              <span class="badge bg-dark">{{ section.segments_count }}</span>
            </li>
            <li class="list-group-item d-flex justify-content-between">
              <span>Общая площадь отрезов</span>
              <span class="badge bg-dark">{{ section.square_sum }} м²</span>
            </li>
          </ul>
          <div class="card-body d-flex flex-column justify-content-end text-center">
            <a
              href="#"
              @click="$router.push(`${$route.fullPath}${section.slug}/segments/`)"
              class="btn btn-outline-secondary stretched-link"
            >
              {{ section.name }}
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
      sections: [],
      isLoading: false,
      error: false,
    };
  },
  methods: {
    fetchSections() {
      this.isLoading = true;
      $api.get(`/companies/${this.$route.params.slug}/sections/`)
        .then((response) => {
          this.sections = response.data;
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
    this.fetchSections();
  },
};
</script>

<style scoped>
.card {
  width: 100%;
}
</style>
