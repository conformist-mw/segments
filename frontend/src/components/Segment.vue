<template>
  <div class="text-dark mb-2">
    <div class="row">
      <div class="col-6">
        <p>
          <ul class="list-group list-group-flush">
            <li class="list-group-item d-flex justify-content-between">
              <span>Фактура: </span><strong>{{ segment.color.type.name }}</strong>
            </li>
            <li class="list-group-item d-flex justify-content-between">
              <span>Цвет: </span><strong>{{ segment.color.name }}</strong>
            </li>
            <li class="list-group-item d-flex justify-content-between">
              <span>Расположение: </span><span>{{ segment.rack }}</span>
            </li>
            <li class="list-group-item d-flex justify-content-between">
              <span>Добавлено: </span><span>{{ dateAdded }}</span>
            </li>
          </ul>
        </p>
      </div>
      <div class="col-6 text-right">
        <p>
          <ul class="list-group list-group-flush">
            <li class="list-group-item d-flex justify-content-between">
              <span>ширина: </span><span class="badge bg-success">{{ segment.width }}см</span>
            </li>
            <li class="list-group-item d-flex justify-content-between">
              <span>длина: </span><span class="badge bg-primary">{{ segment.height }}см</span>
            </li>
            <li class="list-group-item d-flex justify-content-between">
              <span>площадь: </span><span class="badge bg-secondary">{{ segment.square }} м²</span>
            </li>
            <li v-if="!segment.active" class="list-group-item d-flex justify-content-between">
              <span>Удалено: </span><span>{{ dateAdded }}</span>
            </li>
            <li
              v-if="!segment.active && segment.order_number"
              class="list-group-item d-flex justify-content-between"
            >
              <span>Номер заказа: </span><span>{{ segment.order_number.name }}</span>
            </li>
          </ul>
        </p>
      </div>
    </div>
    <div class="d-grid gap-2">
      <button
        @click="$router.push(`${$route.fullPath}${segment.id}`)"
        class="btn btn-outline-secondary"
      >Редактировать</button>
    </div>
    <hr class="mb-3">
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
    dateDeleted() {
      return moment(String(this.segment.deleted)).format('DD.MM.YY | hh:mm');
    },
  },
};
</script>

<style scoped>
</style>
