<template>
  <Card class="mb-5 shadow-5 transition-all">
    <template #content>
      <div class="flex justify-content-between">
        <div>
          <Chip
            :label="segment.color.type.name"
            icon="pi pi-circle"
            class="color-type"
          />
          <Chip :label="segment.color.name" icon="pi pi-circle-fill" class="ml-2" />
        </div>
        <div>
          <Tag
            :value="`${segment.width}см x ${segment.height}см`"
            severity="success"
            icon="pi pi-stop"
          ></Tag>
        </div>
        <div>
          <Tag severity="warning" :value="`${segment.square} м²`" icon="pi pi-times" />
        </div>
      </div>
    </template>
    <template #footer>
      <div class="flex justify-content-between">
        <div>
          Расположение: <span class="rack">{{ segment.rack.name }}</span>
        </div>
        <div>
          Добавлено: {{ dateAdded }}
        </div>
        <div>
          <Button
            @click="toggleDialog"
            class="p-button-icon-only p-button-danger p-button-outlined">
            <i class="pi pi-trash"></i>
          </Button>

        </div>
      </div>
    </template>
  </Card>
  <Dialog v-model:visible="display" :modal="true">
    <template #header>
      <h3>Удалить отрезок</h3>
    </template>
    <form @submit.prevent="deleteSegment(!v$.$invalid, segment.id)">
      <div class="p-fluid">
        <div class="p-field">
          <label
            :class="{'p-error':v$.deletedSegment.orderNumber.$invalid && submitted }"
          >Номер заказа</label>
          <InputText
            v-model="v$.deletedSegment.orderNumber.$model"
            :class="{'p-invalid':v$.deletedSegment.orderNumber.$invalid && submitted }"
            type="text"
          />
          <div v-if="v$.$invalid && submitted">
            Это поле обязательно!
          </div>
        </div>
        <div class="p-field mt-3">
          <label><Checkbox v-model="deletedSegment.hasDefect" :binary="true" /> Есть дефект?</label>
        </div>
        <div class="p-field mt-3">
          <label>Описание отреза</label>
          <InputText v-model="deletedSegment.description" type="text" />
        </div>
      </div>
      <Button
        type="submit"
        label="Удалить"
        icon="pi pi-trash"
        style="float: right"
        class="p-button-danger p-button-outlined mt-5"
      />
    </form>
  </Dialog>
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
.color-type {
  background-color: var(--surface-200);
}
.rack {
  border-bottom: 1px dashed var(--surface-600);
}
label {
  margin-bottom: 0.5rem;
  display: inline-block;
}
</style>
