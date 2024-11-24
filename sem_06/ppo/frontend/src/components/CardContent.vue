<template>
  <div class="styled-scrollbars" style="overflow: auto">
    <div v-html="text"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'CardContent',
  props: ['text'],
});
</script>

<style scoped>

</style>
