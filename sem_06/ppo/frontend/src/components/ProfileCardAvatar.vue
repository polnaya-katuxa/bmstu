<template>
  <img v-bind:src="pic"
       alt="No image" class="img-fluid img-thumbnail mt-4 mb-2"
       style="width: 150px; min-height: 150px !important; z-index: 1">
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ProfileCardAvatar',
  props: ['pic'],
});
</script>

<style scoped>

</style>
