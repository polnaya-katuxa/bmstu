<template>
  <p class="lead fw-normal mb-1">{{ text }}</p>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'ProfileAccentText',
  props: ['text'],
});
</script>

<style scoped>

</style>
