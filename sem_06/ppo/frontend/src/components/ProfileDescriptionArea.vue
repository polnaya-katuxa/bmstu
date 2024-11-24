<template>
  <div class="p-4" style="background-color: rgba(248, 242, 255, 0.99);">
    <p class="font-italic mb-1">{{ text }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ProfileDescriptionArea',
  props: ['text'],
});
</script>

<style scoped>

</style>
