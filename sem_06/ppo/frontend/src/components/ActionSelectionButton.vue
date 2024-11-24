<template>
  <button class="btn btn-sm btn-reaction dropdown-toggle" type="button"
          id="dropdownMenuButton1" data-bs-toggle="dropdown"
          aria-expanded="false">
    {{ text }}
  </button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ActionSelectionButton',
  props: ['text'],
});
</script>

<style scoped>

</style>
