<template>
<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <div class="container">
    <div>
      <a class="navbar-brand" href="/">
        <font-awesome-icon icon="clone" /> Отрезы
      </a>
      <a class="navbar-brand" href="/admin">
        <font-awesome-icon icon="user-shield" /> Admin
      </a>
    </div>
    <div class="d-flex">
      <span>
        <font-awesome-icon icon="user" /> Username
      </span>
    </div>
  </div>
</nav>
</template>

<script>
import { faClone, faUserShield, faUser } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';

library.add(faClone, faUserShield, faUser);
export default {};
</script>

<style scoped>

</style>
