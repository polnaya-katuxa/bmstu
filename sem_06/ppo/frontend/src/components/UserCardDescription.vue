<template>
  <p>{{ text }}</p>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'UserCardDescription',
  props: ['text'],
});
</script>

<style scoped>

</style>
