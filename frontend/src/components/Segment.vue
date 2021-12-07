<template>
  <div class="card text-dark bg-light mb-2">
    <div class="card-header d-flex justify-content-between fs-5">
      <div class="badge bg-info">
        {{ segment.color.type.name }} - {{ segment.color.name }}
      </div>
      <div class="badge bg-success">
        {{ segment.width }}см x {{ segment.height }}см | {{ segment.square }} м²
      </div>
    </div>
    <div class="card-body d-flex justify-content-between">
      <button type="button" class="btn btn-outline-secondary" disabled>{{ dateAdded }}</button>
      <button type="button" class="btn btn-outline-primary">{{ segment.rack.name }}</button>
      <button
          @click="$router.push(`${$route.fullPath}${segment.id}`)"
        class="btn btn-outline-secondary"
      >Редактировать</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

import moment from 'moment';
import useVuelidate from '@vuelidate/core';
import { required } from '@vuelidate/validators';

export default {
  setup() {
    return { v$: useVuelidate() };
  },
  components: {},
  data() {
    return {
      display: false,
      submitted: false,
      deletedSegment: {
        description: '',
        hasDefect: false,
        orderNumber: '',
        segmentId: '',
      },
    };
  },
  validations() {
    return {
      deletedSegment: {
        orderNumber: { required },
      },
    };
  },
  methods: {
    toggleDialog() {
      this.display = !this.display;
    },
    deleteSegment(isFormValid, segmentId) {
      this.submitted = true;

      if (!isFormValid) {
        return;
      }
      const url = `http://127.0.0.1:8000/api/segments/${segmentId}/`;
      axios.patch(url, {
        order_number: this.deletedSegment.orderNumber,
      }).then((response) => {
        console.log(response);
      }).then((error) => {
        console.log(error);
      });
      this.toggleDialog();
    },
  },
  props: {
    segment: {
      type: Object,
      required: true,
    },
  },
  computed: {
    dateAdded() {
      return moment(String(this.segment.created)).format('DD.MM.YY | hh:mm');
    },
  },
};
</script>

<style scoped>
</style>
