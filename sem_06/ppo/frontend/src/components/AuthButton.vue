<template>
  <button type="submit" class="btn btn-pink col">{{ text }}</button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'AuthButton',
  props: ['text'],
});
</script>

<style scoped>

</style>
