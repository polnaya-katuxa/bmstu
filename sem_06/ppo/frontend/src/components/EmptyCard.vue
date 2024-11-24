<template>
  <div class="card">
    <div class="card-body">
      {{ text }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'EmptyCard',
  props: ['text'],
});
</script>

<style scoped>

</style>
