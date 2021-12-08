<template>
  <h1>Редактируем отрез</h1>
  <div v-if="isLoading" class="d-flex justify-content-center">
    <div class="spinner-border" role="status">
      <span class="visually-hidden">Loading...</span>
    </div>
  </div>
  <div v-else class="row">
    <div class="col-7 mx-auto">
      <form action="">
        <div class="row">
          <div class="col">
            <div class="input-group mb-3">
              <span class="input-group-text">Фактура</span>
              <input type="text" class="form-control" :value="segment.color.type.name" disabled>
            </div>
          </div>
          <div class="col">
            <div class="input-group mb-3">
              <span class="input-group-text">Цвет</span>
              <input type="text" class="form-control" :value="segment.color.name" disabled>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <div class="input-group mb-3">
              <span class="input-group-text">Ширина</span>
              <input type="text" class="form-control" :value="segment.width" disabled>
            </div>
          </div>
          <div class="col">
            <div class="input-group mb-3">
              <span class="input-group-text">Высота</span>
              <input type="text" class="form-control" :value="segment.height" disabled>
            </div>
          </div>
          <div class="col">
            <div class="input-group mb-3">
              <span class="input-group-text">Площадь</span>
              <input type="text" class="form-control" :value="segment.square" disabled>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="col">
            <div class="form-check form-switch">
              <input
                class="form-check-input"
                type="checkbox"
                id="to_delete"
                :checked="segment.active"
              >
              <label class="form-check-label" for="to_delete">Активный</label>
              <div class="form-text">Отрез будет отмечен как удалённый</div>
            </div>
          </div>
          <div class="col">
            <div class="form-check form-switch">
              <input
                class="form-check-input"
                type="checkbox"
                id="has_defect"
                :checked="segment.defect"
              >
              <label class="form-check-label" for="has_defect">Есть дефект?</label>
              <div class="form-text">Потребуется добавить описание дефекта</div>
            </div>
          </div>
        </div>
        <div class="row mt-3">
          <div class="col">
            <div class="input-group mb-3">
              <span class="input-group-text">Номер заказа</span>
              <input type="text" class="form-control" :value="segment.order_number">
            </div>
          </div>
          <div class="col">
            <div class="input-group mb-3">
              <label class="input-group-text" for="rack">Расположение</label>
              <select class="form-select" id="rack">
                <option
                  v-for="rack in segment.racks"
                  v-bind:key="rack.id"
                  value="rack.id"
                  :selected="rack.id === segment.rack.id"
                >
                  {{ rack.name }}
                </option>
              </select>
            </div>
          </div>
        </div>
        <div class="input-group">
          <span class="input-group-text">Описание</span>
          <textarea class="form-control" :value="segment.description"></textarea>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import useSegmentEdit from '../hooks/useSegmentEdit';

export default {
  setup() {
    const { segment, error, isLoading } = useSegmentEdit();

    return { segment, error, isLoading };
  },
};
</script>

<style scoped>
.form-control:disabled {
  background-color: #f7f7f7;
}
</style>
