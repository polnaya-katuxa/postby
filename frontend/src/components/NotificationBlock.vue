<template>
  <div class="notification-block">
    <div v-for="error in errors" :key="error"
         class="alert alert-danger alert-dismissible fade show" role="alert">
      <strong>Error occurred. </strong> {{ error }}
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"
              @click="deleteError(error)"></button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from 'vuex';

export default defineComponent({
  name: 'NotificationsBlock',
  methods: {
    ...mapActions(['addError', 'deleteError']),
  },
  computed: {
    ...mapGetters(['errors']),
  },
});
</script>

<style scoped>
.notification-block {
  position: fixed;
  width: 23vw;
  bottom: 10px;
  right: 17px;
}
</style>
