<template>
  <strong>Error occurred. </strong> {{ error }}
  <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"
          @click="deleteError(error)"></button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapActions } from 'vuex';

export default defineComponent({
  name: 'NotificationBody',
  props: ['error'],
  methods: {
    ...mapActions(['addError', 'deleteError']),
  },
});
</script>

<style scoped>

</style>
